package sink;

import constants.Constants;
import model.TronTransaction;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.api.functions.sink.RichSinkFunction;
import org.neo4j.driver.Config;
import org.neo4j.driver.Driver;
import org.neo4j.driver.GraphDatabase;

import java.util.HashMap;
import java.util.Map;

public class neo4jSink extends RichSinkFunction<model.TronTransaction> {
    private  Driver driver;
    private final String insert="MERGE (owner:Address{address:$ownerAddress})"+
                                "MERGE (to: Address{address: $toAddress})"+
                                "MERGE (owner)-[r:transaction{chain:tron,amount:$amount}]->(to)";
    @Override
    public void open(Configuration parameters) throws Exception {
        Config.builder().with
        driver=GraphDatabase.driver(Constants.Neo4j_URL);
    }

    @Override
    public void close() throws Exception {
        driver.close();
    }

    @Override
    public void invoke(model.TronTransaction value, Context context) throws Exception {
        Map<String,Object> Parameters =new HashMap<>();
        Parameters.put("ownerAddress",value.ownerAddress);
        Parameters.put("toAddress",value.toAddress);
        Parameters.put("amount",value.amount);
        driver.session().run(insert,Parameters);
    }
}
