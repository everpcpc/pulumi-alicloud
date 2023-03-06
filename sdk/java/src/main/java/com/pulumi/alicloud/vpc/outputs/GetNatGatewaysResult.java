// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.alicloud.vpc.outputs.GetNatGatewaysGateway;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetNatGatewaysResult {
    private @Nullable Boolean dryRun;
    private @Nullable Boolean enableDetails;
    /**
     * @return A list of Nat gateways. Each element contains the following attributes:
     * 
     */
    private List<GetNatGatewaysGateway> gateways;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (Optional) A list of Nat gateways IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of Nat gateways names.
     * 
     */
    private List<String> names;
    /**
     * @return The name of the NAT gateway.
     * 
     */
    private @Nullable String natGatewayName;
    /**
     * @return The type of the NAT gateway.
     * 
     */
    private @Nullable String natType;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;
    /**
     * @return The billing method of the NAT gateway.
     * 
     */
    private @Nullable String paymentType;
    /**
     * @return The ID of the resource group.
     * 
     */
    private @Nullable String resourceGroupId;
    /**
     * @return The specification of the NAT gateway.
     * 
     */
    private @Nullable String specification;
    /**
     * @return The status of the NAT gateway.
     * 
     */
    private @Nullable String status;
    /**
     * @return The tags of NAT gateway.
     * 
     */
    private @Nullable Map<String,Object> tags;
    private Integer totalCount;
    /**
     * @return The ID of the VPC.
     * 
     */
    private @Nullable String vpcId;

    private GetNatGatewaysResult() {}
    public Optional<Boolean> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    /**
     * @return A list of Nat gateways. Each element contains the following attributes:
     * 
     */
    public List<GetNatGatewaysGateway> gateways() {
        return this.gateways;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Optional) A list of Nat gateways IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of Nat gateways names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    /**
     * @return The name of the NAT gateway.
     * 
     */
    public Optional<String> natGatewayName() {
        return Optional.ofNullable(this.natGatewayName);
    }
    /**
     * @return The type of the NAT gateway.
     * 
     */
    public Optional<String> natType() {
        return Optional.ofNullable(this.natType);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }
    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }
    /**
     * @return The billing method of the NAT gateway.
     * 
     */
    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }
    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return The specification of the NAT gateway.
     * 
     */
    public Optional<String> specification() {
        return Optional.ofNullable(this.specification);
    }
    /**
     * @return The status of the NAT gateway.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return The tags of NAT gateway.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    public Integer totalCount() {
        return this.totalCount;
    }
    /**
     * @return The ID of the VPC.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNatGatewaysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean dryRun;
        private @Nullable Boolean enableDetails;
        private List<GetNatGatewaysGateway> gateways;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String natGatewayName;
        private @Nullable String natType;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        private @Nullable String paymentType;
        private @Nullable String resourceGroupId;
        private @Nullable String specification;
        private @Nullable String status;
        private @Nullable Map<String,Object> tags;
        private Integer totalCount;
        private @Nullable String vpcId;
        public Builder() {}
        public Builder(GetNatGatewaysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dryRun = defaults.dryRun;
    	      this.enableDetails = defaults.enableDetails;
    	      this.gateways = defaults.gateways;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.natGatewayName = defaults.natGatewayName;
    	      this.natType = defaults.natType;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
    	      this.paymentType = defaults.paymentType;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.specification = defaults.specification;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.totalCount = defaults.totalCount;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder dryRun(@Nullable Boolean dryRun) {
            this.dryRun = dryRun;
            return this;
        }
        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder gateways(List<GetNatGatewaysGateway> gateways) {
            this.gateways = Objects.requireNonNull(gateways);
            return this;
        }
        public Builder gateways(GetNatGatewaysGateway... gateways) {
            return gateways(List.of(gateways));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder natGatewayName(@Nullable String natGatewayName) {
            this.natGatewayName = natGatewayName;
            return this;
        }
        @CustomType.Setter
        public Builder natType(@Nullable String natType) {
            this.natType = natType;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder pageNumber(@Nullable Integer pageNumber) {
            this.pageNumber = pageNumber;
            return this;
        }
        @CustomType.Setter
        public Builder pageSize(@Nullable Integer pageSize) {
            this.pageSize = pageSize;
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(@Nullable String paymentType) {
            this.paymentType = paymentType;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder specification(@Nullable String specification) {
            this.specification = specification;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder totalCount(Integer totalCount) {
            this.totalCount = Objects.requireNonNull(totalCount);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        public GetNatGatewaysResult build() {
            final var o = new GetNatGatewaysResult();
            o.dryRun = dryRun;
            o.enableDetails = enableDetails;
            o.gateways = gateways;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.natGatewayName = natGatewayName;
            o.natType = natType;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            o.paymentType = paymentType;
            o.resourceGroupId = resourceGroupId;
            o.specification = specification;
            o.status = status;
            o.tags = tags;
            o.totalCount = totalCount;
            o.vpcId = vpcId;
            return o;
        }
    }
}