package processor;

import model.Transaction;
import model.TronTransaction;
import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.api.java.tuple.Tuple3;
import org.apache.flink.util.Collector;

public class TronFlatmapProcessor implements FlatMapFunction<TronTransaction, Tuple3<String,String,Long>> {
    @Override
    public void flatMap(TronTransaction tronTransaction, Collector<Tuple3<String, String, Long>> collector) throws Exception {
         String From =tronTransaction.From;
         String payoff = "支付";
         String obtain ="获得";
         String to = tronTransaction.To;
         collector.collect(new Tuple3<>(From,payoff,Long.getLong(tronTransaction.amount)));
         collector.collect(new Tuple3<>(to,obtain,Long.getLong(tronTransaction.amount)));
    }
}
