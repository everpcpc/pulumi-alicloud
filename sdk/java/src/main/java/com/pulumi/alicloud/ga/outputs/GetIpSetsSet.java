// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIpSetsSet {
    /**
     * @return The ID of an acceleration region.
     * 
     */
    private String accelerateRegionId;
    /**
     * @return The bandwidth allocated to the acceleration region.
     * 
     */
    private Integer bandwidth;
    /**
     * @return The ID of the Ip Set.
     * 
     */
    private String id;
    /**
     * @return The list of accelerated IP addresses in the acceleration region.
     * 
     */
    private List<String> ipAddressLists;
    /**
     * @return Accelerated area ID.
     * 
     */
    private String ipSetId;
    /**
     * @return The IP protocol used by the GA instance.
     * 
     */
    private String ipVersion;
    /**
     * @return The status of the acceleration region.
     * 
     */
    private String status;

    private GetIpSetsSet() {}
    /**
     * @return The ID of an acceleration region.
     * 
     */
    public String accelerateRegionId() {
        return this.accelerateRegionId;
    }
    /**
     * @return The bandwidth allocated to the acceleration region.
     * 
     */
    public Integer bandwidth() {
        return this.bandwidth;
    }
    /**
     * @return The ID of the Ip Set.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of accelerated IP addresses in the acceleration region.
     * 
     */
    public List<String> ipAddressLists() {
        return this.ipAddressLists;
    }
    /**
     * @return Accelerated area ID.
     * 
     */
    public String ipSetId() {
        return this.ipSetId;
    }
    /**
     * @return The IP protocol used by the GA instance.
     * 
     */
    public String ipVersion() {
        return this.ipVersion;
    }
    /**
     * @return The status of the acceleration region.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpSetsSet defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accelerateRegionId;
        private Integer bandwidth;
        private String id;
        private List<String> ipAddressLists;
        private String ipSetId;
        private String ipVersion;
        private String status;
        public Builder() {}
        public Builder(GetIpSetsSet defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accelerateRegionId = defaults.accelerateRegionId;
    	      this.bandwidth = defaults.bandwidth;
    	      this.id = defaults.id;
    	      this.ipAddressLists = defaults.ipAddressLists;
    	      this.ipSetId = defaults.ipSetId;
    	      this.ipVersion = defaults.ipVersion;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder accelerateRegionId(String accelerateRegionId) {
            this.accelerateRegionId = Objects.requireNonNull(accelerateRegionId);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidth(Integer bandwidth) {
            this.bandwidth = Objects.requireNonNull(bandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddressLists(List<String> ipAddressLists) {
            this.ipAddressLists = Objects.requireNonNull(ipAddressLists);
            return this;
        }
        public Builder ipAddressLists(String... ipAddressLists) {
            return ipAddressLists(List.of(ipAddressLists));
        }
        @CustomType.Setter
        public Builder ipSetId(String ipSetId) {
            this.ipSetId = Objects.requireNonNull(ipSetId);
            return this;
        }
        @CustomType.Setter
        public Builder ipVersion(String ipVersion) {
            this.ipVersion = Objects.requireNonNull(ipVersion);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetIpSetsSet build() {
            final var o = new GetIpSetsSet();
            o.accelerateRegionId = accelerateRegionId;
            o.bandwidth = bandwidth;
            o.id = id;
            o.ipAddressLists = ipAddressLists;
            o.ipSetId = ipSetId;
            o.ipVersion = ipVersion;
            o.status = status;
            return o;
        }
    }
}
