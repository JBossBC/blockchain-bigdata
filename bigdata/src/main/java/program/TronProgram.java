package program;

import constants.Constants;
import model.TronTransaction;

import org.apache.flink.api.common.eventtime.WatermarkStrategy;
import org.apache.flink.api.common.serialization.SimpleStringSchema;

import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.api.windowing.assigners.TumblingEventTimeWindows;
import org.apache.flink.streaming.api.windowing.time.Time;
import org.apache.flink.streaming.connectors.elasticsearch7.ElasticsearchSink;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer09;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaProducer09;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import processor.TronFlatmapProcessor;
import processor.TronKeySelectProcessor;
import processor.TronTransactionAmountAggProcessor;
import processor.TronTransactionFilterProcessor;
import sink.ElasticsearchEmit;
import sink.TronHbase;
import util.HttpHostUtil;

import java.time.Duration;
import java.util.Arrays;
import java.util.Properties;

public class TronProgram implements  IProgram{
    private FlinkKafkaConsumer09<model.TronTransaction>consumer;
    private ElasticsearchSink<TronTransaction> elasticsearchSink;
    private final StreamExecutionEnvironment env;
    private FlinkKafkaProducer09<String>Producer;
//    private final String KafkaProducerTopic ="transaction-agg";
    public TronProgram(StreamExecutionEnvironment env){
        this.env=env;
        initSource();
        initSink();
    }
    @Override
    public void initSink() {
        this.elasticsearchSink = new org.apache.flink.streaming.connectors.elasticsearch7.ElasticsearchSink.Builder<TronTransaction>(Arrays.asList(HttpHostUtil.httpHosts(Constants.ES_CLUSTER_123)), new ElasticsearchEmit()).build();
        this.Producer =new FlinkKafkaProducer09<String>(Constants.Kafka_Bootstrap_Server,Constants.KafkaProducer_Topic,new SimpleStringSchema());
    }

    @Override
    public void initSource() {
       Properties properties=new Properties();
       properties.setProperty(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG,Constants.Kafka_Bootstrap_Server);
       properties.setProperty(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG,Constants.Kafka_Offset_Reset);
       properties.setProperty(ConsumerConfig.GROUP_ID_CONFIG,Constants.Kafka_Consumer_Group_Id);
       this.consumer=new FlinkKafkaConsumer09<>(Constants.Kafka_Topic_Name, new schema.TronTransaction.TronTransactionSchema(),properties);
    }

    @Override
    public void process() {
        SingleOutputStreamOperator<TronTransaction> tron =  this.env.addSource(this.consumer).assignTimestampsAndWatermarks(WatermarkStrategy.<TronTransaction>forBoundedOutOfOrderness(Duration.ofSeconds(30)).withTimestampAssigner((tronTransaction, l) -> tronTransaction.CreateTime));
        SingleOutputStreamOperator<TronTransaction> reasonableTransaction = tron.process(new TronTransactionFilterProcessor());
        reasonableTransaction.addSink(this.elasticsearchSink).setParallelism(10).name("Elasticsearch Storage");
        reasonableTransaction.addSink(new TronHbase()).setParallelism(10).name("HBase Storage");
        reasonableTransaction.flatMap(new TronFlatmapProcessor()).keyBy(new TronKeySelectProcessor()).window(TumblingEventTimeWindows.of(Time.days(1))).process(new TronTransactionAmountAggProcessor()).addSink(Producer).name("aggregation the transaction");
    }
}
