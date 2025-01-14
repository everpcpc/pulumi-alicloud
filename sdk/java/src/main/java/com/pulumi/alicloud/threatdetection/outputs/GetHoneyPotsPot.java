// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetHoneyPotsPot {
    /**
     * @return Honeypot ID.
     * 
     */
    private String honeypotId;
    /**
     * @return The image ID of the honeypot.
     * 
     */
    private String honeypotImageId;
    /**
     * @return Honeypot mirror name.
     * 
     */
    private String honeypotImageName;
    /**
     * @return Honeypot custom name.
     * 
     */
    private String honeypotName;
    /**
     * @return Honeypot ID. The value is the same as `honeypot_id`.
     * 
     */
    private String id;
    /**
     * @return The ID of the honeypot management node.
     * 
     */
    private String nodeId;
    /**
     * @return The custom parameter ID of honeypot.
     * 
     */
    private String presetId;
    /**
     * @return Honeypot status.
     * 
     */
    private List<String> states;
    /**
     * @return The status of the resource
     * 
     */
    private String status;

    private GetHoneyPotsPot() {}
    /**
     * @return Honeypot ID.
     * 
     */
    public String honeypotId() {
        return this.honeypotId;
    }
    /**
     * @return The image ID of the honeypot.
     * 
     */
    public String honeypotImageId() {
        return this.honeypotImageId;
    }
    /**
     * @return Honeypot mirror name.
     * 
     */
    public String honeypotImageName() {
        return this.honeypotImageName;
    }
    /**
     * @return Honeypot custom name.
     * 
     */
    public String honeypotName() {
        return this.honeypotName;
    }
    /**
     * @return Honeypot ID. The value is the same as `honeypot_id`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the honeypot management node.
     * 
     */
    public String nodeId() {
        return this.nodeId;
    }
    /**
     * @return The custom parameter ID of honeypot.
     * 
     */
    public String presetId() {
        return this.presetId;
    }
    /**
     * @return Honeypot status.
     * 
     */
    public List<String> states() {
        return this.states;
    }
    /**
     * @return The status of the resource
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHoneyPotsPot defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String honeypotId;
        private String honeypotImageId;
        private String honeypotImageName;
        private String honeypotName;
        private String id;
        private String nodeId;
        private String presetId;
        private List<String> states;
        private String status;
        public Builder() {}
        public Builder(GetHoneyPotsPot defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.honeypotId = defaults.honeypotId;
    	      this.honeypotImageId = defaults.honeypotImageId;
    	      this.honeypotImageName = defaults.honeypotImageName;
    	      this.honeypotName = defaults.honeypotName;
    	      this.id = defaults.id;
    	      this.nodeId = defaults.nodeId;
    	      this.presetId = defaults.presetId;
    	      this.states = defaults.states;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder honeypotId(String honeypotId) {
            this.honeypotId = Objects.requireNonNull(honeypotId);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotImageId(String honeypotImageId) {
            this.honeypotImageId = Objects.requireNonNull(honeypotImageId);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotImageName(String honeypotImageName) {
            this.honeypotImageName = Objects.requireNonNull(honeypotImageName);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotName(String honeypotName) {
            this.honeypotName = Objects.requireNonNull(honeypotName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder nodeId(String nodeId) {
            this.nodeId = Objects.requireNonNull(nodeId);
            return this;
        }
        @CustomType.Setter
        public Builder presetId(String presetId) {
            this.presetId = Objects.requireNonNull(presetId);
            return this;
        }
        @CustomType.Setter
        public Builder states(List<String> states) {
            this.states = Objects.requireNonNull(states);
            return this;
        }
        public Builder states(String... states) {
            return states(List.of(states));
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetHoneyPotsPot build() {
            final var o = new GetHoneyPotsPot();
            o.honeypotId = honeypotId;
            o.honeypotImageId = honeypotImageId;
            o.honeypotImageName = honeypotImageName;
            o.honeypotName = honeypotName;
            o.id = id;
            o.nodeId = nodeId;
            o.presetId = presetId;
            o.states = states;
            o.status = status;
            return o;
        }
    }
}
