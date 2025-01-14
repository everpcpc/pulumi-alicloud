// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicecatalog.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLaunchOptionsOptionConstraintSummary {
    /**
     * @return Constraint type.The value is Launch, which indicates that the constraint is started.
     * 
     */
    private String constraintType;
    /**
     * @return Constraint description.
     * 
     */
    private String description;

    private GetLaunchOptionsOptionConstraintSummary() {}
    /**
     * @return Constraint type.The value is Launch, which indicates that the constraint is started.
     * 
     */
    public String constraintType() {
        return this.constraintType;
    }
    /**
     * @return Constraint description.
     * 
     */
    public String description() {
        return this.description;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLaunchOptionsOptionConstraintSummary defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String constraintType;
        private String description;
        public Builder() {}
        public Builder(GetLaunchOptionsOptionConstraintSummary defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.constraintType = defaults.constraintType;
    	      this.description = defaults.description;
        }

        @CustomType.Setter
        public Builder constraintType(String constraintType) {
            this.constraintType = Objects.requireNonNull(constraintType);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public GetLaunchOptionsOptionConstraintSummary build() {
            final var o = new GetLaunchOptionsOptionConstraintSummary();
            o.constraintType = constraintType;
            o.description = description;
            return o;
        }
    }
}
