package schema;

import com.google.gson.Gson;
import org.apache.flink.api.common.serialization.DeserializationSchema;
import org.apache.flink.api.common.typeinfo.TypeInformation;

import java.io.IOException;


public class TronTransaction {
    public static class TronTransactionSchema implements DeserializationSchema<model.TronTransaction>{
        private static final long serialVersionUID= 922095187725158430L;
        private static Gson gson =new Gson();
        @Override
        public model.TronTransaction deserialize(byte[] bytes) throws IOException {
            if (bytes!=null){
                return gson.fromJson(String.valueOf(bytes),model.TronTransaction.class);
            }
            return null;
        }

        @Override
        public boolean isEndOfStream(model.TronTransaction tronTransaction) {
            return false;
        }

        @Override
        public TypeInformation<model.TronTransaction> getProducedType() {
            return TypeInformation.of(model.TronTransaction.class);
        }
    }
}
