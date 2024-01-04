package model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class TronTransaction extends Transaction {

    @JsonProperty("ownerAddress")
    public String ownerAddress;

    @JsonProperty("toAddressList")
    public String[] toAddressList;
    @JsonProperty("toAddress")
    public String toAddress;

    @JsonProperty("contractType")
    public long contractType;

    @JsonProperty("confirmed")
    public boolean confirmed;

    @JsonProperty("revert")
    public boolean revert;

    @JsonProperty("contractData")
    public ContractData contractData;

    @JsonProperty("SmartCalls")
    public String smartCalls;

    @JsonProperty("Events")
    public String events;

    @JsonProperty("id")
    public String id;

    @JsonProperty("data")
    public String data;

    @JsonProperty("fee")
    public String fee;

    @JsonProperty("contractRet")
    public String contractRet;

    @JsonProperty("result")
    public String result;

    @JsonProperty("amount")
    public String amount;

    @JsonProperty("cost")
    public Cost cost;

    @JsonProperty("tokenInfo")
    public TokenInfo tokenInfo;

    @JsonProperty("tokenType")
    public String tokenType;

    // 省略 getter 和 setter 方法

    public static class ContractData {
        @JsonProperty("amount")
        public long amount;

        @JsonProperty("owner_address")
        public String ownerAddress;

        @JsonProperty("to_address")
        public String toAddress;

        // 省略 getter 和 setter 方法
    }

    public static class Cost {
        @JsonProperty("net_fee")
        public long netFee;

        @JsonProperty("energy_penalty_total")
        public long energyPenaltyTotal;

        @JsonProperty("energy_usage")
        public long energyUsage;

        @JsonProperty("fee")
        public long fee;

        @JsonProperty("energy_fee")
        public long energyFee;

        @JsonProperty("energy_usage_total")
        public long energyUsageTotal;

        @JsonProperty("origin_energy_usage")
        public long originEnergyUsage;

        @JsonProperty("net_usage")
        public long netUsage;

        // 省略 getter 和 setter 方法
    }

    public static class TokenInfo {
        @JsonProperty("tokenId")
        public String tokenId;

        @JsonProperty("tokenAbbr")
        public String tokenAbbr;

        @JsonProperty("tokenName")
        public String tokenName;

        @JsonProperty("tokenDecimal")
        public long tokenDecimal;

        @JsonProperty("tokenCanShow")
        public long tokenCanShow;

        @JsonProperty("tokenType")
        public String tokenType;

        @JsonProperty("tokenLogo")
        public String tokenLogo;

        @JsonProperty("tokenLevel")
        public String tokenLevel;

        @JsonProperty("vip")
        public boolean vip;

        // 省略 getter 和 setter 方法
    }
}
