// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetContainerGroupsGroupHostAlias {
    /**
     * @return The name of the host.
     * 
     */
    private List<String> hostnames;
    /**
     * @return The IP address of the container.
     * 
     */
    private String ip;

    private GetContainerGroupsGroupHostAlias() {}
    /**
     * @return The name of the host.
     * 
     */
    public List<String> hostnames() {
        return this.hostnames;
    }
    /**
     * @return The IP address of the container.
     * 
     */
    public String ip() {
        return this.ip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerGroupsGroupHostAlias defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> hostnames;
        private String ip;
        public Builder() {}
        public Builder(GetContainerGroupsGroupHostAlias defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hostnames = defaults.hostnames;
    	      this.ip = defaults.ip;
        }

        @CustomType.Setter
        public Builder hostnames(List<String> hostnames) {
            this.hostnames = Objects.requireNonNull(hostnames);
            return this;
        }
        public Builder hostnames(String... hostnames) {
            return hostnames(List.of(hostnames));
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        public GetContainerGroupsGroupHostAlias build() {
            final var o = new GetContainerGroupsGroupHostAlias();
            o.hostnames = hostnames;
            o.ip = ip;
            return o;
        }
    }
}
