// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDdosBgpInstancesInstance {
    /**
     * @return The instance&#39;s elastic defend bandwidth.
     * 
     */
    private Integer bandwidth;
    /**
     * @return The instance&#39;s base defend bandwidth.
     * 
     */
    private Integer baseBandwidth;
    /**
     * @return The instance&#39;s id.
     * 
     */
    private String id;
    /**
     * @return The instance&#39;s count of ip config.
     * 
     */
    private Integer ipCount;
    /**
     * @return The instance&#39;s IP version.
     * 
     */
    private String ipType;
    /**
     * @return The instance&#39;s remark.
     * 
     */
    private String name;
    /**
     * @return Normal defend bandwidth of the instance. The unit is Gbps.
     * 
     */
    private Integer normalBandwidth;
    /**
     * @return A region of instance.
     * 
     */
    private String region;
    /**
     * @return The instance&#39;s type.
     * 
     */
    private String type;

    private GetDdosBgpInstancesInstance() {}
    /**
     * @return The instance&#39;s elastic defend bandwidth.
     * 
     */
    public Integer bandwidth() {
        return this.bandwidth;
    }
    /**
     * @return The instance&#39;s base defend bandwidth.
     * 
     */
    public Integer baseBandwidth() {
        return this.baseBandwidth;
    }
    /**
     * @return The instance&#39;s id.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The instance&#39;s count of ip config.
     * 
     */
    public Integer ipCount() {
        return this.ipCount;
    }
    /**
     * @return The instance&#39;s IP version.
     * 
     */
    public String ipType() {
        return this.ipType;
    }
    /**
     * @return The instance&#39;s remark.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Normal defend bandwidth of the instance. The unit is Gbps.
     * 
     */
    public Integer normalBandwidth() {
        return this.normalBandwidth;
    }
    /**
     * @return A region of instance.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The instance&#39;s type.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDdosBgpInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer bandwidth;
        private Integer baseBandwidth;
        private String id;
        private Integer ipCount;
        private String ipType;
        private String name;
        private Integer normalBandwidth;
        private String region;
        private String type;
        public Builder() {}
        public Builder(GetDdosBgpInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bandwidth = defaults.bandwidth;
    	      this.baseBandwidth = defaults.baseBandwidth;
    	      this.id = defaults.id;
    	      this.ipCount = defaults.ipCount;
    	      this.ipType = defaults.ipType;
    	      this.name = defaults.name;
    	      this.normalBandwidth = defaults.normalBandwidth;
    	      this.region = defaults.region;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder bandwidth(Integer bandwidth) {
            this.bandwidth = Objects.requireNonNull(bandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder baseBandwidth(Integer baseBandwidth) {
            this.baseBandwidth = Objects.requireNonNull(baseBandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipCount(Integer ipCount) {
            this.ipCount = Objects.requireNonNull(ipCount);
            return this;
        }
        @CustomType.Setter
        public Builder ipType(String ipType) {
            this.ipType = Objects.requireNonNull(ipType);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder normalBandwidth(Integer normalBandwidth) {
            this.normalBandwidth = Objects.requireNonNull(normalBandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetDdosBgpInstancesInstance build() {
            final var o = new GetDdosBgpInstancesInstance();
            o.bandwidth = bandwidth;
            o.baseBandwidth = baseBandwidth;
            o.id = id;
            o.ipCount = ipCount;
            o.ipType = ipType;
            o.name = name;
            o.normalBandwidth = normalBandwidth;
            o.region = region;
            o.type = type;
            return o;
        }
    }
}