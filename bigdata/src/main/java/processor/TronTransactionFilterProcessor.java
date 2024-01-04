package processor;

import model.TronTransaction;
import org.apache.flink.api.common.state.StateTtlConfig;
import org.apache.flink.api.common.state.ValueState;
import org.apache.flink.api.common.state.ValueStateDescriptor;
import org.apache.flink.api.common.time.Time;
import org.apache.flink.api.common.typeinfo.TypeInformation;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.api.functions.ProcessFunction;
import org.apache.flink.util.Collector;
import org.roaringbitmap.longlong.Roaring64Bitmap;

import java.util.Objects;

public class TronTransactionFilterProcessor extends ProcessFunction<TronTransaction,TronTransaction> {
       private  ValueState<Roaring64Bitmap> status;
    @Override
    public void open(Configuration parameters) throws Exception {
    ValueStateDescriptor  valueDescriptor=new ValueStateDescriptor<>("bitmap_tron_transaction", TypeInformation.of(Roaring64Bitmap.class));
    StateTtlConfig stateTtlConfig = StateTtlConfig.newBuilder(Time.days(15)).setStateVisibility(StateTtlConfig.StateVisibility.NeverReturnExpired)
            .updateTtlOnCreateAndWrite().cleanupFullSnapshot().build();
    valueDescriptor.enableTimeToLive(stateTtlConfig);
    this.status = getRuntimeContext().getState(valueDescriptor);
    }

    @Override
    public void processElement(TronTransaction tronTransaction, ProcessFunction<TronTransaction, TronTransaction>.Context context, Collector<TronTransaction> collector) throws Exception {
           if (tronTransaction.revert ||!tronTransaction.confirmed){
               return;
           }
           if (Objects.isNull(this.status.value())){
               this.status.update(new Roaring64Bitmap());
           }
           // must cant is null
          Roaring64Bitmap roaing= this.status.value();
           if (roaing.contains(Long.parseLong(tronTransaction.Hash))){
              return;
           }
           roaing.add(Long.parseLong(tronTransaction.Hash));
           this.status.update(roaing);
           collector.collect(tronTransaction);
    }
}
