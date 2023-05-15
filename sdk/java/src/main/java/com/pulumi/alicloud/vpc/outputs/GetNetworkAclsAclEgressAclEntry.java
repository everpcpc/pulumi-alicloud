// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetNetworkAclsAclEgressAclEntry {
    /**
     * @return Description of the entry direction rule.
     * 
     */
    private String description;
    /**
     * @return The destination address segment.
     * 
     */
    private String destinationCidrIp;
    /**
     * @return The name of the entry direction rule entry.
     * 
     */
    private String networkAclEntryName;
    /**
     * @return The authorization policy.
     * 
     */
    private String policy;
    /**
     * @return Source port range.
     * 
     */
    private String port;
    /**
     * @return Transport layer protocol.
     * 
     */
    private String protocol;

    private GetNetworkAclsAclEgressAclEntry() {}
    /**
     * @return Description of the entry direction rule.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The destination address segment.
     * 
     */
    public String destinationCidrIp() {
        return this.destinationCidrIp;
    }
    /**
     * @return The name of the entry direction rule entry.
     * 
     */
    public String networkAclEntryName() {
        return this.networkAclEntryName;
    }
    /**
     * @return The authorization policy.
     * 
     */
    public String policy() {
        return this.policy;
    }
    /**
     * @return Source port range.
     * 
     */
    public String port() {
        return this.port;
    }
    /**
     * @return Transport layer protocol.
     * 
     */
    public String protocol() {
        return this.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNetworkAclsAclEgressAclEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String destinationCidrIp;
        private String networkAclEntryName;
        private String policy;
        private String port;
        private String protocol;
        public Builder() {}
        public Builder(GetNetworkAclsAclEgressAclEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.destinationCidrIp = defaults.destinationCidrIp;
    	      this.networkAclEntryName = defaults.networkAclEntryName;
    	      this.policy = defaults.policy;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder destinationCidrIp(String destinationCidrIp) {
            this.destinationCidrIp = Objects.requireNonNull(destinationCidrIp);
            return this;
        }
        @CustomType.Setter
        public Builder networkAclEntryName(String networkAclEntryName) {
            this.networkAclEntryName = Objects.requireNonNull(networkAclEntryName);
            return this;
        }
        @CustomType.Setter
        public Builder policy(String policy) {
            this.policy = Objects.requireNonNull(policy);
            return this;
        }
        @CustomType.Setter
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        public GetNetworkAclsAclEgressAclEntry build() {
            final var o = new GetNetworkAclsAclEgressAclEntry();
            o.description = description;
            o.destinationCidrIp = destinationCidrIp;
            o.networkAclEntryName = networkAclEntryName;
            o.policy = policy;
            o.port = port;
            o.protocol = protocol;
            return o;
        }
    }
}
