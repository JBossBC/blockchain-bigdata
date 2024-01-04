package processor;

import model.TronTransaction;
import org.apache.flink.api.java.functions.KeySelector;
import org.apache.flink.api.java.tuple.Tuple3;

public class TronKeySelectProcessor implements KeySelector<Tuple3<String,String,Long>,String> {


    @Override
    public String getKey(Tuple3<String, String, Long> tronInfo) throws Exception {
        return tronInfo.f0+tronInfo.f1;
    }
}
