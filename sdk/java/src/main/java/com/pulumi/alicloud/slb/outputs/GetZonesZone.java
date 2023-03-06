// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.alicloud.slb.outputs.GetZonesZoneSupportedResource;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetZonesZone {
    /**
     * @return ID of the zone. It is same as `master_zone_id`.
     * 
     */
    private String id;
    /**
     * @return The primary zone.
     * 
     */
    private String masterZoneId;
    /**
     * @return The secondary zone.
     * 
     */
    private String slaveZoneId;
    /**
     * @return (Deprecated from 1.157.0) A list of slb slave zone ids in which the slb master zone.
     * It has been deprecated from v1.157.0 and use `slave_zone_id` instead.
     * 
     * @deprecated
     * the attribute slb_slave_zone_ids has been deprecated from version 1.157.0 and use slave_zone_id instead.
     * 
     */
    @Deprecated /* the attribute slb_slave_zone_ids has been deprecated from version 1.157.0 and use slave_zone_id instead. */
    private List<String> slbSlaveZoneIds;
    /**
     * @return (Available in 1.154.0+)A list of available resource which the slb master zone supported.
     * 
     */
    private List<GetZonesZoneSupportedResource> supportedResources;

    private GetZonesZone() {}
    /**
     * @return ID of the zone. It is same as `master_zone_id`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The primary zone.
     * 
     */
    public String masterZoneId() {
        return this.masterZoneId;
    }
    /**
     * @return The secondary zone.
     * 
     */
    public String slaveZoneId() {
        return this.slaveZoneId;
    }
    /**
     * @return (Deprecated from 1.157.0) A list of slb slave zone ids in which the slb master zone.
     * It has been deprecated from v1.157.0 and use `slave_zone_id` instead.
     * 
     * @deprecated
     * the attribute slb_slave_zone_ids has been deprecated from version 1.157.0 and use slave_zone_id instead.
     * 
     */
    @Deprecated /* the attribute slb_slave_zone_ids has been deprecated from version 1.157.0 and use slave_zone_id instead. */
    public List<String> slbSlaveZoneIds() {
        return this.slbSlaveZoneIds;
    }
    /**
     * @return (Available in 1.154.0+)A list of available resource which the slb master zone supported.
     * 
     */
    public List<GetZonesZoneSupportedResource> supportedResources() {
        return this.supportedResources;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZonesZone defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String masterZoneId;
        private String slaveZoneId;
        private List<String> slbSlaveZoneIds;
        private List<GetZonesZoneSupportedResource> supportedResources;
        public Builder() {}
        public Builder(GetZonesZone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.masterZoneId = defaults.masterZoneId;
    	      this.slaveZoneId = defaults.slaveZoneId;
    	      this.slbSlaveZoneIds = defaults.slbSlaveZoneIds;
    	      this.supportedResources = defaults.supportedResources;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder masterZoneId(String masterZoneId) {
            this.masterZoneId = Objects.requireNonNull(masterZoneId);
            return this;
        }
        @CustomType.Setter
        public Builder slaveZoneId(String slaveZoneId) {
            this.slaveZoneId = Objects.requireNonNull(slaveZoneId);
            return this;
        }
        @CustomType.Setter
        public Builder slbSlaveZoneIds(List<String> slbSlaveZoneIds) {
            this.slbSlaveZoneIds = Objects.requireNonNull(slbSlaveZoneIds);
            return this;
        }
        public Builder slbSlaveZoneIds(String... slbSlaveZoneIds) {
            return slbSlaveZoneIds(List.of(slbSlaveZoneIds));
        }
        @CustomType.Setter
        public Builder supportedResources(List<GetZonesZoneSupportedResource> supportedResources) {
            this.supportedResources = Objects.requireNonNull(supportedResources);
            return this;
        }
        public Builder supportedResources(GetZonesZoneSupportedResource... supportedResources) {
            return supportedResources(List.of(supportedResources));
        }
        public GetZonesZone build() {
            final var o = new GetZonesZone();
            o.id = id;
            o.masterZoneId = masterZoneId;
            o.slaveZoneId = slaveZoneId;
            o.slbSlaveZoneIds = slbSlaveZoneIds;
            o.supportedResources = supportedResources;
            return o;
        }
    }
}