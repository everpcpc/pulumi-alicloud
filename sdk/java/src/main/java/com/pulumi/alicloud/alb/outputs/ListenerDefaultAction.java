// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.alicloud.alb.outputs.ListenerDefaultActionForwardGroupConfig;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ListenerDefaultAction {
    /**
     * @return The configurations of the actions. This parameter is required if Type is set to FowardGroup.
     * 
     */
    private ListenerDefaultActionForwardGroupConfig forwardGroupConfig;
    /**
     * @return Action Type.
     * 
     */
    private String type;

    private ListenerDefaultAction() {}
    /**
     * @return The configurations of the actions. This parameter is required if Type is set to FowardGroup.
     * 
     */
    public ListenerDefaultActionForwardGroupConfig forwardGroupConfig() {
        return this.forwardGroupConfig;
    }
    /**
     * @return Action Type.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ListenerDefaultAction defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private ListenerDefaultActionForwardGroupConfig forwardGroupConfig;
        private String type;
        public Builder() {}
        public Builder(ListenerDefaultAction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.forwardGroupConfig = defaults.forwardGroupConfig;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder forwardGroupConfig(ListenerDefaultActionForwardGroupConfig forwardGroupConfig) {
            this.forwardGroupConfig = Objects.requireNonNull(forwardGroupConfig);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public ListenerDefaultAction build() {
            final var o = new ListenerDefaultAction();
            o.forwardGroupConfig = forwardGroupConfig;
            o.type = type;
            return o;
        }
    }
}
