// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList {
    /**
     * @return The target network segment of the local VPC.
     * 
     */
    private String localDestinationCidr;
    /**
     * @return The ID of the next-hop instance in the local VPC.
     * 
     */
    private String localNextHopInstanceId;

    private FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList() {}
    /**
     * @return The target network segment of the local VPC.
     * 
     */
    public String localDestinationCidr() {
        return this.localDestinationCidr;
    }
    /**
     * @return The ID of the next-hop instance in the local VPC.
     * 
     */
    public String localNextHopInstanceId() {
        return this.localNextHopInstanceId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String localDestinationCidr;
        private String localNextHopInstanceId;
        public Builder() {}
        public Builder(FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.localDestinationCidr = defaults.localDestinationCidr;
    	      this.localNextHopInstanceId = defaults.localNextHopInstanceId;
        }

        @CustomType.Setter
        public Builder localDestinationCidr(String localDestinationCidr) {
            this.localDestinationCidr = Objects.requireNonNull(localDestinationCidr);
            return this;
        }
        @CustomType.Setter
        public Builder localNextHopInstanceId(String localNextHopInstanceId) {
            this.localNextHopInstanceId = Objects.requireNonNull(localNextHopInstanceId);
            return this;
        }
        public FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList build() {
            final var o = new FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList();
            o.localDestinationCidr = localDestinationCidr;
            o.localNextHopInstanceId = localNextHopInstanceId;
            return o;
        }
    }
}
