// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPublicIpAddressPoolCidrBlocksBlock {
    /**
     * @return The CIDR block.
     * 
     */
    private String cidrBlock;
    /**
     * @return The time when the CIDR block was created. The time is displayed in YYYY-MM-DDThh:mm:ssZ format.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the Public Ip Address Pool Cidr Block.
     * 
     */
    private String id;
    /**
     * @return The ID of the Vpc Public IP address pool.
     * 
     */
    private String publicIpAddressPoolId;
    /**
     * @return The status of the CIDR block in the Vpc Public IP address pool. Valid values: `Created`, `Modifying`, `Deleting`.
     * 
     */
    private String status;
    /**
     * @return The number of occupied IP addresses in the CIDR block.
     * 
     */
    private Integer totalIpNum;
    /**
     * @return The total number of available IP addresses in the CIDR block.
     * 
     */
    private Integer usedIpNum;

    private GetPublicIpAddressPoolCidrBlocksBlock() {}
    /**
     * @return The CIDR block.
     * 
     */
    public String cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * @return The time when the CIDR block was created. The time is displayed in YYYY-MM-DDThh:mm:ssZ format.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the Public Ip Address Pool Cidr Block.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the Vpc Public IP address pool.
     * 
     */
    public String publicIpAddressPoolId() {
        return this.publicIpAddressPoolId;
    }
    /**
     * @return The status of the CIDR block in the Vpc Public IP address pool. Valid values: `Created`, `Modifying`, `Deleting`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The number of occupied IP addresses in the CIDR block.
     * 
     */
    public Integer totalIpNum() {
        return this.totalIpNum;
    }
    /**
     * @return The total number of available IP addresses in the CIDR block.
     * 
     */
    public Integer usedIpNum() {
        return this.usedIpNum;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPublicIpAddressPoolCidrBlocksBlock defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cidrBlock;
        private String createTime;
        private String id;
        private String publicIpAddressPoolId;
        private String status;
        private Integer totalIpNum;
        private Integer usedIpNum;
        public Builder() {}
        public Builder(GetPublicIpAddressPoolCidrBlocksBlock defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidrBlock = defaults.cidrBlock;
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.publicIpAddressPoolId = defaults.publicIpAddressPoolId;
    	      this.status = defaults.status;
    	      this.totalIpNum = defaults.totalIpNum;
    	      this.usedIpNum = defaults.usedIpNum;
        }

        @CustomType.Setter
        public Builder cidrBlock(String cidrBlock) {
            this.cidrBlock = Objects.requireNonNull(cidrBlock);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder publicIpAddressPoolId(String publicIpAddressPoolId) {
            this.publicIpAddressPoolId = Objects.requireNonNull(publicIpAddressPoolId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder totalIpNum(Integer totalIpNum) {
            this.totalIpNum = Objects.requireNonNull(totalIpNum);
            return this;
        }
        @CustomType.Setter
        public Builder usedIpNum(Integer usedIpNum) {
            this.usedIpNum = Objects.requireNonNull(usedIpNum);
            return this;
        }
        public GetPublicIpAddressPoolCidrBlocksBlock build() {
            final var o = new GetPublicIpAddressPoolCidrBlocksBlock();
            o.cidrBlock = cidrBlock;
            o.createTime = createTime;
            o.id = id;
            o.publicIpAddressPoolId = publicIpAddressPoolId;
            o.status = status;
            o.totalIpNum = totalIpNum;
            o.usedIpNum = usedIpNum;
            return o;
        }
    }
}