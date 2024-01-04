package eventbus

import (
	"bufio"
	"container/list"
	"fmt"
	"sync"
	"time"
)

const (
	_default_concurrent_number = 15
	_default_memthreshold      = 500 * 1024 * 1024
	_default_failed_sleep_time = 30 * time.Millisecond
)

type Event interface {
	CallBack() any
}

type EventBus struct {
	// 0 代表开始,1代表停止,2代表终止
	stats        int32
	cond         *sync.Cond
	drain        int32
	memThreshold uint64
	concurrentN  int32
	//容忍的最大消费失败次数
	tolerateConsumeFailed int32
	//消费失败重试的过度策略
	retryExcess func()
	channel     chan Event
	produce     func() Event
	consume     func(Event) bool
	failedPool  *list.List
	final       chan any
}

type EventBusConfig func(*EventBus)

func WithMaxMemThreshold(memThreshold uint64) EventBusConfig {
	return func(eb *EventBus) {
		eb.memThreshold = memThreshold
	}
}
func WithMaxConcurrentNumber(maxConcurrentNumber int32) EventBusConfig {
	return func(eb *EventBus) {
		eb.concurrentN = maxConcurrentNumber
	}
}

func NewEventBus(produce func() Event, consume func(Event) bool, configs ...EventBusConfig) (eventbus *EventBus) {
	eventbus = new(EventBus)
	eventbus.cond = sync.NewCond(&sync.RWMutex{})
	eventbus.produce = produce
	eventbus.consume = consume
	eventbus.memThreshold = _default_memthreshold
	eventbus.concurrentN = _default_concurrent_number
	eventbus.retryExcess = func() {
		time.Sleep(_default_failed_sleep_time)
	}
	eventbus.failedPool = list.New()
	for _, value := range configs {
		value(eventbus)
	}
	eventbus.final = make(chan any)
	eventbus.channel = make(chan Event, eventbus.concurrentN)
	return eventbus
}

func (eventbus *EventBus) Run() (err error) {
	defer func() {
		if panicError := recover(); err != nil {
			err = fmt.Errorf("事件总线出现异常:%v", panicError)
		}
	}()
	producer := make(chan any, eventbus.concurrentN)
	produceOnce := make(chan any)
	onceExecute := sync.OnceFunc(func() {
		produceOnce <- struct{}{}
	})
	go func() {
		for {
			producer <- struct{}{}
			go func() {
				defer func() {
					<-producer
				}()
				eventbus.cond.L.Lock()
				for eventbus.stats != 0 {
					eventbus.cond.Wait()
				}
				eventbus.cond.L.Unlock()
				eventbus.channel <- eventbus.produce()
				onceExecute()
			}()
		}
	}()
	<-produceOnce
	close(produceOnce)
	consumer := make(chan any, eventbus.concurrentN)
	go func() {
		consumer <- struct{}{}
		for {
			defer func() {
				<-consumer
			}()
			eventbus.cond.L.Lock()
			for eventbus.stats != 0 {
				eventbus.cond.Wait()
			}
			eventbus.cond.L.Unlock()
			e := <-eventbus.channel
			var retry = 0
			for !eventbus.consume(e) && retry < int(eventbus.tolerateConsumeFailed) {
				eventbus.retryExcess()
				retry++
			}
			if retry >= int(eventbus.tolerateConsumeFailed) {
				eventbus.failedPool.PushBack(e)
				return
			}

		}
	}()
	<-eventbus.final
	return err
}
