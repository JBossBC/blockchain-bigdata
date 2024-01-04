package database

import (
	"context"
	"errors"
	"log"
	"neo4j-producer/config"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

const (
	_kafka_version  = "kafka.version"
	_kafka_assignor = "kafka.assignor"
	_kafka_oldest   = "kafka.oldest"
	_kafka_brokers  = "kafka.brokers"
	_kafka_group    = "kafka.group"
	_kafka_topics   = "kafka.topics"
)

var (
	oldest   = config.GetBool(_kafka_oldest)
	verbose  = false
	assignor = config.GetString(_kafka_assignor)
	brokers  = config.GetString(_kafka_brokers)
	group    = config.GetString(_kafka_group)
	topics   = config.GetString(_kafka_topics)
)

func init() {
	keepRunning := true
	log.Println("Starting a new Sarama consumer")
	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(config.GetString(_kafka_version))
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}
	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	consumer := Consumer{
		ready: make(chan bool),
	}
	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

type Consumer struct {
	ready chan bool
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s,timestamp = %v,topic= %s", string(message.Value), message.Timestamp, message.Topic)
			//TODO the message consumer policy
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
