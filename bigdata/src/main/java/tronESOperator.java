import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import program.TronProgram;


public class tronESOperator {
    public static void main(String[] args) throws Exception {
        StreamExecutionEnvironment env =StreamExecutionEnvironment.getExecutionEnvironment();
        new TronProgram(env).process();
    }
}
