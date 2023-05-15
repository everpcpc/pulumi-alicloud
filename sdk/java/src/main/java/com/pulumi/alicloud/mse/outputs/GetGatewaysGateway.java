// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse.outputs;

import com.pulumi.alicloud.mse.outputs.GetGatewaysGatewaySlbList;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetGatewaysGateway {
    /**
     * @return The backup vswitch id.
     * 
     */
    private String backupVswitchId;
    /**
     * @return The name of the Gateway.
     * 
     */
    private String gatewayName;
    /**
     * @return Gateway unique identification.
     * 
     */
    private String gatewayUniqueId;
    /**
     * @return The ID of the Gateway.
     * 
     */
    private String id;
    /**
     * @return The payment type of the resource.
     * 
     */
    private String paymentType;
    /**
     * @return Number of Gateway Nodes.
     * 
     */
    private String replica;
    /**
     * @return A list of gateway Slb.
     * 
     */
    private List<GetGatewaysGatewaySlbList> slbLists;
    /**
     * @return Gateway Node Specifications.
     * 
     */
    private String spec;
    /**
     * @return The status of the gateway.
     * 
     */
    private String status;
    /**
     * @return The ID of the vpc.
     * 
     */
    private String vpcId;
    /**
     * @return The ID of the vswitch.
     * 
     */
    private String vswitchId;

    private GetGatewaysGateway() {}
    /**
     * @return The backup vswitch id.
     * 
     */
    public String backupVswitchId() {
        return this.backupVswitchId;
    }
    /**
     * @return The name of the Gateway.
     * 
     */
    public String gatewayName() {
        return this.gatewayName;
    }
    /**
     * @return Gateway unique identification.
     * 
     */
    public String gatewayUniqueId() {
        return this.gatewayUniqueId;
    }
    /**
     * @return The ID of the Gateway.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The payment type of the resource.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return Number of Gateway Nodes.
     * 
     */
    public String replica() {
        return this.replica;
    }
    /**
     * @return A list of gateway Slb.
     * 
     */
    public List<GetGatewaysGatewaySlbList> slbLists() {
        return this.slbLists;
    }
    /**
     * @return Gateway Node Specifications.
     * 
     */
    public String spec() {
        return this.spec;
    }
    /**
     * @return The status of the gateway.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the vpc.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The ID of the vswitch.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGatewaysGateway defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backupVswitchId;
        private String gatewayName;
        private String gatewayUniqueId;
        private String id;
        private String paymentType;
        private String replica;
        private List<GetGatewaysGatewaySlbList> slbLists;
        private String spec;
        private String status;
        private String vpcId;
        private String vswitchId;
        public Builder() {}
        public Builder(GetGatewaysGateway defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupVswitchId = defaults.backupVswitchId;
    	      this.gatewayName = defaults.gatewayName;
    	      this.gatewayUniqueId = defaults.gatewayUniqueId;
    	      this.id = defaults.id;
    	      this.paymentType = defaults.paymentType;
    	      this.replica = defaults.replica;
    	      this.slbLists = defaults.slbLists;
    	      this.spec = defaults.spec;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder backupVswitchId(String backupVswitchId) {
            this.backupVswitchId = Objects.requireNonNull(backupVswitchId);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayName(String gatewayName) {
            this.gatewayName = Objects.requireNonNull(gatewayName);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayUniqueId(String gatewayUniqueId) {
            this.gatewayUniqueId = Objects.requireNonNull(gatewayUniqueId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
            return this;
        }
        @CustomType.Setter
        public Builder replica(String replica) {
            this.replica = Objects.requireNonNull(replica);
            return this;
        }
        @CustomType.Setter
        public Builder slbLists(List<GetGatewaysGatewaySlbList> slbLists) {
            this.slbLists = Objects.requireNonNull(slbLists);
            return this;
        }
        public Builder slbLists(GetGatewaysGatewaySlbList... slbLists) {
            return slbLists(List.of(slbLists));
        }
        @CustomType.Setter
        public Builder spec(String spec) {
            this.spec = Objects.requireNonNull(spec);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        public GetGatewaysGateway build() {
            final var o = new GetGatewaysGateway();
            o.backupVswitchId = backupVswitchId;
            o.gatewayName = gatewayName;
            o.gatewayUniqueId = gatewayUniqueId;
            o.id = id;
            o.paymentType = paymentType;
            o.replica = replica;
            o.slbLists = slbLists;
            o.spec = spec;
            o.status = status;
            o.vpcId = vpcId;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
