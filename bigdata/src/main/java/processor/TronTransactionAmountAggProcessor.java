package processor;

import model.TronTransaction;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.api.java.tuple.Tuple3;
import org.apache.flink.streaming.api.functions.KeyedProcessFunction;
import org.apache.flink.streaming.api.functions.windowing.ProcessWindowFunction;
import org.apache.flink.streaming.api.windowing.windows.TimeWindow;
import org.apache.flink.util.Collector;

import java.text.SimpleDateFormat;
import java.util.Iterator;
import java.util.TimeZone;

public class TronTransactionAmountAggProcessor extends ProcessWindowFunction<Tuple3<String,String,Long>, String,String, TimeWindow> {


    @Override
    public void process(String s, ProcessWindowFunction<Tuple3<String, String, Long>, String, String, TimeWindow>.Context context, Iterable<Tuple3<String, String, Long>> iterable, Collector<String> collector) throws Exception {
        long sum =0;
        Iterator<Tuple3<String, String, Long>> infos = iterable.iterator();
        while(infos.hasNext()){
            Tuple3<String, String, Long> next = infos.next();
            sum+=next.f2.longValue();
        }
        StringBuilder sb =new StringBuilder();
        TimeZone timeZone = TimeZone.getTimeZone("Asia/Shanghai");
        SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        sdf.setTimeZone(timeZone);
        String start =sdf.format(context.window().getStart());
        String end =sdf.format(context.window().getEnd());
        sb.append("[tron blockchain]");
        sb.append("从");
        sb.append(start);
        sb.append("到");
        sb.append(end);
        sb.append("\t");
        sb.append(s);
        sb.append(":\t");
        sb.append(sum);
        collector.collect(sb.toString());
    }
}
