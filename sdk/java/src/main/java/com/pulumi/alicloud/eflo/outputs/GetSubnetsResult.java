// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eflo.outputs;

import com.pulumi.alicloud.eflo.outputs.GetSubnetsSubnet;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSubnetsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of name of Subnets.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;
    /**
     * @return Resource Group ID.
     * 
     */
    private @Nullable String resourceGroupId;
    /**
     * @return The status of the resource.
     * 
     */
    private @Nullable String status;
    /**
     * @return The Eflo subnet ID.
     * 
     */
    private @Nullable String subnetId;
    /**
     * @return The Subnet name.
     * 
     */
    private @Nullable String subnetName;
    /**
     * @return A list of Subnet Entries. Each element contains the following attributes:
     * 
     */
    private List<GetSubnetsSubnet> subnets;
    /**
     * @return Eflo subnet usage type.
     * 
     */
    private @Nullable String type;
    /**
     * @return Eflo VPD ID.
     * 
     */
    private @Nullable String vpdId;
    /**
     * @return The zone ID of the resource.
     * 
     */
    private @Nullable String zoneId;

    private GetSubnetsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of name of Subnets.
     * 
     */
    public List<String> names() {
        return this.names;
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
     * @return Resource Group ID.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return The status of the resource.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return The Eflo subnet ID.
     * 
     */
    public Optional<String> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }
    /**
     * @return The Subnet name.
     * 
     */
    public Optional<String> subnetName() {
        return Optional.ofNullable(this.subnetName);
    }
    /**
     * @return A list of Subnet Entries. Each element contains the following attributes:
     * 
     */
    public List<GetSubnetsSubnet> subnets() {
        return this.subnets;
    }
    /**
     * @return Eflo subnet usage type.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return Eflo VPD ID.
     * 
     */
    public Optional<String> vpdId() {
        return Optional.ofNullable(this.vpdId);
    }
    /**
     * @return The zone ID of the resource.
     * 
     */
    public Optional<String> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSubnetsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        private @Nullable String resourceGroupId;
        private @Nullable String status;
        private @Nullable String subnetId;
        private @Nullable String subnetName;
        private List<GetSubnetsSubnet> subnets;
        private @Nullable String type;
        private @Nullable String vpdId;
        private @Nullable String zoneId;
        public Builder() {}
        public Builder(GetSubnetsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.status = defaults.status;
    	      this.subnetId = defaults.subnetId;
    	      this.subnetName = defaults.subnetName;
    	      this.subnets = defaults.subnets;
    	      this.type = defaults.type;
    	      this.vpdId = defaults.vpdId;
    	      this.zoneId = defaults.zoneId;
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
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder subnetId(@Nullable String subnetId) {
            this.subnetId = subnetId;
            return this;
        }
        @CustomType.Setter
        public Builder subnetName(@Nullable String subnetName) {
            this.subnetName = subnetName;
            return this;
        }
        @CustomType.Setter
        public Builder subnets(List<GetSubnetsSubnet> subnets) {
            this.subnets = Objects.requireNonNull(subnets);
            return this;
        }
        public Builder subnets(GetSubnetsSubnet... subnets) {
            return subnets(List.of(subnets));
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder vpdId(@Nullable String vpdId) {
            this.vpdId = vpdId;
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(@Nullable String zoneId) {
            this.zoneId = zoneId;
            return this;
        }
        public GetSubnetsResult build() {
            final var o = new GetSubnetsResult();
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            o.resourceGroupId = resourceGroupId;
            o.status = status;
            o.subnetId = subnetId;
            o.subnetName = subnetName;
            o.subnets = subnets;
            o.type = type;
            o.vpdId = vpdId;
            o.zoneId = zoneId;
            return o;
        }
    }
}
