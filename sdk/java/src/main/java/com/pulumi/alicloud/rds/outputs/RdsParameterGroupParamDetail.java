// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RdsParameterGroupParamDetail {
    /**
     * @return The name of a parameter.
     * 
     */
    private String paramName;
    /**
     * @return The value of a parameter.
     * 
     */
    private String paramValue;

    private RdsParameterGroupParamDetail() {}
    /**
     * @return The name of a parameter.
     * 
     */
    public String paramName() {
        return this.paramName;
    }
    /**
     * @return The value of a parameter.
     * 
     */
    public String paramValue() {
        return this.paramValue;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RdsParameterGroupParamDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String paramName;
        private String paramValue;
        public Builder() {}
        public Builder(RdsParameterGroupParamDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.paramName = defaults.paramName;
    	      this.paramValue = defaults.paramValue;
        }

        @CustomType.Setter
        public Builder paramName(String paramName) {
            this.paramName = Objects.requireNonNull(paramName);
            return this;
        }
        @CustomType.Setter
        public Builder paramValue(String paramValue) {
            this.paramValue = Objects.requireNonNull(paramValue);
            return this;
        }
        public RdsParameterGroupParamDetail build() {
            final var o = new RdsParameterGroupParamDetail();
            o.paramName = paramName;
            o.paramValue = paramValue;
            return o;
        }
    }
}
