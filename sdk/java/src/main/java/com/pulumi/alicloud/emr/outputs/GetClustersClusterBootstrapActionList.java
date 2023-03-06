// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClustersClusterBootstrapActionList {
    /**
     * @return Parameters of the boot operation.
     * 
     */
    private String arg;
    /**
     * @return The internal name of the service.
     * 
     */
    private String name;
    /**
     * @return Boot operation script path.
     * 
     */
    private String path;

    private GetClustersClusterBootstrapActionList() {}
    /**
     * @return Parameters of the boot operation.
     * 
     */
    public String arg() {
        return this.arg;
    }
    /**
     * @return The internal name of the service.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Boot operation script path.
     * 
     */
    public String path() {
        return this.path;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClustersClusterBootstrapActionList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arg;
        private String name;
        private String path;
        public Builder() {}
        public Builder(GetClustersClusterBootstrapActionList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arg = defaults.arg;
    	      this.name = defaults.name;
    	      this.path = defaults.path;
        }

        @CustomType.Setter
        public Builder arg(String arg) {
            this.arg = Objects.requireNonNull(arg);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public GetClustersClusterBootstrapActionList build() {
            final var o = new GetClustersClusterBootstrapActionList();
            o.arg = arg;
            o.name = name;
            o.path = path;
            return o;
        }
    }
}