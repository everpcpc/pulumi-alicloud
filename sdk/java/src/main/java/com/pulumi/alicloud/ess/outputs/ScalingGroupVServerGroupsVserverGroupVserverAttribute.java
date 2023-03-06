// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ScalingGroupVServerGroupsVserverGroupVserverAttribute {
    private Integer port;
    private String vserverGroupId;
    private Integer weight;

    private ScalingGroupVServerGroupsVserverGroupVserverAttribute() {}
    public Integer port() {
        return this.port;
    }
    public String vserverGroupId() {
        return this.vserverGroupId;
    }
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ScalingGroupVServerGroupsVserverGroupVserverAttribute defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer port;
        private String vserverGroupId;
        private Integer weight;
        public Builder() {}
        public Builder(ScalingGroupVServerGroupsVserverGroupVserverAttribute defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.port = defaults.port;
    	      this.vserverGroupId = defaults.vserverGroupId;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder vserverGroupId(String vserverGroupId) {
            this.vserverGroupId = Objects.requireNonNull(vserverGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public ScalingGroupVServerGroupsVserverGroupVserverAttribute build() {
            final var o = new ScalingGroupVServerGroupsVserverGroupVserverAttribute();
            o.port = port;
            o.vserverGroupId = vserverGroupId;
            o.weight = weight;
            return o;
        }
    }
}