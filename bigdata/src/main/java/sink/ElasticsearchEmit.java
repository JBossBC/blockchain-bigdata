package sink;

import model.TronTransaction;
import org.apache.flink.api.common.functions.RuntimeContext;
import org.apache.flink.api.connector.sink2.SinkWriter;
import org.apache.flink.connector.elasticsearch.sink.ElasticsearchEmitter;
import org.apache.flink.connector.elasticsearch.sink.RequestIndexer;
import org.apache.flink.streaming.connectors.elasticsearch.ElasticsearchSinkFunction;
import org.elasticsearch.action.index.IndexRequest;

public class ElasticsearchEmit implements ElasticsearchSinkFunction<TronTransaction> {
//    @Override
//    public void emit(TronTransaction tronTransaction, SinkWriter.Context context, RequestIndexer requestIndexer) {
//        IndexRequest request = new IndexRequest();
//        request.source(tronTransaction).index("transaction");
//        requestIndexer.add(request);
//    }

    @Override
    public void process(TronTransaction tronTransaction, RuntimeContext runtimeContext, org.apache.flink.streaming.connectors.elasticsearch.RequestIndexer requestIndexer) {
        IndexRequest request = new IndexRequest();
        request.source(tronTransaction).index("transaction");
        requestIndexer.add(request);
    }
}
