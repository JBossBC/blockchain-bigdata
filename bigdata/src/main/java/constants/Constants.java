package constants;

public interface Constants {
    String Kafka_Bootstrap_Server =  "dev-master1:9092,dev-worker1:9092,dev-worker2:9092";
    String Kafka_Offset_Reset = "earliest";
    String Kafka_Consumer_Group_Id = "tron";
    String Kafka_Topic_Name ="tron-transaction";
    String ES_CLUSTER_123 = "http://192.168.31.194:9200";
    String HBase_Table_Name ="blockchain-transaction";
    String KafkaProducer_Topic ="transaction-agg";
    String Neo4j_URL = "http://dev-master1:7687";
}
