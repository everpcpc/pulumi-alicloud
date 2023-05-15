// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetTemplateScratchesScratchPreferenceParameter {
    /**
     * @return Priority parameter key.
     * 
     */
    private String parameterKey;
    /**
     * @return Priority parameter value.
     * 
     */
    private String parameterValue;

    private GetTemplateScratchesScratchPreferenceParameter() {}
    /**
     * @return Priority parameter key.
     * 
     */
    public String parameterKey() {
        return this.parameterKey;
    }
    /**
     * @return Priority parameter value.
     * 
     */
    public String parameterValue() {
        return this.parameterValue;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTemplateScratchesScratchPreferenceParameter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String parameterKey;
        private String parameterValue;
        public Builder() {}
        public Builder(GetTemplateScratchesScratchPreferenceParameter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.parameterKey = defaults.parameterKey;
    	      this.parameterValue = defaults.parameterValue;
        }

        @CustomType.Setter
        public Builder parameterKey(String parameterKey) {
            this.parameterKey = Objects.requireNonNull(parameterKey);
            return this;
        }
        @CustomType.Setter
        public Builder parameterValue(String parameterValue) {
            this.parameterValue = Objects.requireNonNull(parameterValue);
            return this;
        }
        public GetTemplateScratchesScratchPreferenceParameter build() {
            final var o = new GetTemplateScratchesScratchPreferenceParameter();
            o.parameterKey = parameterKey;
            o.parameterValue = parameterValue;
            return o;
        }
    }
}
