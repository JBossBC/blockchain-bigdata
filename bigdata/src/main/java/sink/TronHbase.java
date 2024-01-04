package sink;
import constants.Constants;
import model.TronTransaction;
import org.apache.flink.addons.hbase.HBaseTableSchema;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.api.functions.sink.RichSinkFunction;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.*;
import org.apache.hadoop.hbase.util.Bytes;
import schema.HBaseTable;

public   class TronHbase extends RichSinkFunction<TronTransaction> {
    public  static Connection connection;
    public static Table table;
    @Override
    public void open(Configuration parameters) throws Exception {
        org.apache.hadoop.conf.Configuration config = HBaseConfiguration.create();
        config.set("hbase.zookeeper.quorum","dev-master1");
        config.set("hbase.zookeeper.property.clientPort","2181");
        connection = ConnectionFactory.createConnection(config);
        Admin admin=connection.getAdmin();
        if (!admin.tableExists(TableName.valueOf(Constants.HBase_Table_Name))){
//            admin.createTable();
        }
        table =connection.getTable(TableName.valueOf("blockchain-transaction"));
    }
    @Override
    public void invoke(TronTransaction tran, Context context) throws Exception {
        Put put =new Put(Bytes.toBytes(tran.Hash));
        put.addColumn(Bytes.toBytes("blockchain"),Bytes.toBytes("type"),Bytes.toBytes(tran.blockchain));
        put.addColumn(Bytes.toBytes("transaction"),Bytes.toBytes("from"),Bytes.toBytes(tran.ownerAddress));
        put.addColumn(Bytes.toBytes("transaction"),Bytes.toBytes("to"),Bytes.toBytes(tran.toAddress));
        put.addColumn(Bytes.toBytes("transaction"),Bytes.toBytes("amount"),Bytes.toBytes(tran.amount));
        put.addColumn(Bytes.toBytes("transaction"),Bytes.toBytes("timestamp"),Bytes.toBytes(tran.CreateTime));
        put.addColumn(Bytes.toBytes("extra"),Bytes.toBytes("block"),Bytes.toBytes(tran.BlockNumber));
        put.addColumn(Bytes.toBytes("extra"),Bytes.toBytes("transactionIndex"),Bytes.toBytes(tran.BlockNumberIndex));
        put.addColumn(Bytes.toBytes("extra"),Bytes.toBytes("fee"),Bytes.toBytes(tran.fee));
        table.put(put);
    }
}
