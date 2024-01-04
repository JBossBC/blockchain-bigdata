package model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Transaction {
    @JsonProperty("blockchain")
    public String blockchain;
    @JsonProperty("from")
    public String From;
    @JsonProperty("to")
    public String To;
    @JsonProperty("balance")
    public int  Balance;
    @JsonProperty("fee")
    public int  Fee;
    @JsonProperty("block")
    public int BlockNumber;
    // the index for the block transactions
    @JsonProperty("blockIndex")
    public int BlockNumberIndex;
    @JsonProperty("timestamp")
    public long CreateTime;
    @JsonProperty("status")
    public int Status;
    @JsonProperty("hash")
    public String Hash;
    public String getId(){
        return this.Hash;
    }
}
