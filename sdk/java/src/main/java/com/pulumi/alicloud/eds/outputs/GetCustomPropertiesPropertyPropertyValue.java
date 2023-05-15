// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCustomPropertiesPropertyPropertyValue {
    /**
     * @return The value of an attribute.
     * 
     */
    private String propertyValue;
    /**
     * @return The value of an attribute id.
     * 
     */
    private String propertyValueId;

    private GetCustomPropertiesPropertyPropertyValue() {}
    /**
     * @return The value of an attribute.
     * 
     */
    public String propertyValue() {
        return this.propertyValue;
    }
    /**
     * @return The value of an attribute id.
     * 
     */
    public String propertyValueId() {
        return this.propertyValueId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCustomPropertiesPropertyPropertyValue defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String propertyValue;
        private String propertyValueId;
        public Builder() {}
        public Builder(GetCustomPropertiesPropertyPropertyValue defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.propertyValue = defaults.propertyValue;
    	      this.propertyValueId = defaults.propertyValueId;
        }

        @CustomType.Setter
        public Builder propertyValue(String propertyValue) {
            this.propertyValue = Objects.requireNonNull(propertyValue);
            return this;
        }
        @CustomType.Setter
        public Builder propertyValueId(String propertyValueId) {
            this.propertyValueId = Objects.requireNonNull(propertyValueId);
            return this;
        }
        public GetCustomPropertiesPropertyPropertyValue build() {
            final var o = new GetCustomPropertiesPropertyPropertyValue();
            o.propertyValue = propertyValue;
            o.propertyValueId = propertyValueId;
            return o;
        }
    }
}
