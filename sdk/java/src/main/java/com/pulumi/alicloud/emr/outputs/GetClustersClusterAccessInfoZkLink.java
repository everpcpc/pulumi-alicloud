// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClustersClusterAccessInfoZkLink {
    /**
     * @return The access link address of ZooKeeper.
     * 
     */
    private String link;
    /**
     * @return The port of ZooKeeper.
     * 
     */
    private String port;

    private GetClustersClusterAccessInfoZkLink() {}
    /**
     * @return The access link address of ZooKeeper.
     * 
     */
    public String link() {
        return this.link;
    }
    /**
     * @return The port of ZooKeeper.
     * 
     */
    public String port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClustersClusterAccessInfoZkLink defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String link;
        private String port;
        public Builder() {}
        public Builder(GetClustersClusterAccessInfoZkLink defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.link = defaults.link;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder link(String link) {
            this.link = Objects.requireNonNull(link);
            return this;
        }
        @CustomType.Setter
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public GetClustersClusterAccessInfoZkLink build() {
            final var o = new GetClustersClusterAccessInfoZkLink();
            o.link = link;
            o.port = port;
            return o;
        }
    }
}