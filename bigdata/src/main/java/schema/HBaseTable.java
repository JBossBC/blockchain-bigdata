package schema;

import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.ColumnFamilyDescriptor;
import org.apache.hadoop.hbase.client.CoprocessorDescriptor;
import org.apache.hadoop.hbase.client.Durability;
import org.apache.hadoop.hbase.client.TableDescriptor;
import org.apache.hadoop.hbase.util.Bytes;

import java.util.Collection;
import java.util.Map;
import java.util.Set;

public class HBaseTable implements TableDescriptor {
    @Override
    public int getColumnFamilyCount() {
        return 0;
    }

    @Override
    public Collection<CoprocessorDescriptor> getCoprocessorDescriptors() {
        return null;
    }

    @Override
    public Durability getDurability() {
        return null;
    }

    @Override
    public ColumnFamilyDescriptor[] getColumnFamilies() {
        return new ColumnFamilyDescriptor[0];
    }

    @Override
    public Set<byte[]> getColumnFamilyNames() {
        return null;
    }

    @Override
    public ColumnFamilyDescriptor getColumnFamily(byte[] bytes) {
        return null;
    }

    @Override
    public String getFlushPolicyClassName() {
        return null;
    }

    @Override
    public long getMaxFileSize() {
        return 0;
    }

    @Override
    public long getMemStoreFlushSize() {
        return 0;
    }

    @Override
    public int getPriority() {
        return 0;
    }

    @Override
    public int getRegionReplication() {
        return 0;
    }

    @Override
    public String getRegionSplitPolicyClassName() {
        return null;
    }

    @Override
    public TableName getTableName() {
        return null;
    }

    @Override
    public String getOwnerString() {
        return null;
    }

    @Override
    public Bytes getValue(Bytes bytes) {
        return null;
    }

    @Override
    public byte[] getValue(byte[] bytes) {
        return new byte[0];
    }

    @Override
    public String getValue(String s) {
        return null;
    }

    @Override
    public Map<Bytes, Bytes> getValues() {
        return null;
    }

    @Override
    public boolean hasCoprocessor(String s) {
        return false;
    }

    @Override
    public boolean hasColumnFamily(byte[] bytes) {
        return false;
    }

    @Override
    public boolean hasRegionMemStoreReplication() {
        return false;
    }

    @Override
    public boolean isCompactionEnabled() {
        return false;
    }

    @Override
    public boolean isSplitEnabled() {
        return false;
    }

    @Override
    public boolean isMergeEnabled() {
        return false;
    }

    @Override
    public boolean isMetaRegion() {
        return false;
    }

    @Override
    public boolean isMetaTable() {
        return false;
    }

    @Override
    public boolean isNormalizationEnabled() {
        return false;
    }

    @Override
    public int getNormalizerTargetRegionCount() {
        return 0;
    }

    @Override
    public long getNormalizerTargetRegionSize() {
        return 0;
    }

    @Override
    public boolean isReadOnly() {
        return false;
    }

    @Override
    public String toStringCustomizedValues() {
        return null;
    }
}
