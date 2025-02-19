// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetVpcFirewallsFirewallPeerVpcPeerVpcCidrTableListPeerRouteEntryList {
    /**
     * @return The target network segment of the peer VPC.
     * 
     */
    private String peerDestinationCidr;
    /**
     * @return The ID of the next-hop instance in the peer VPC.
     * 
     */
    private String peerNextHopInstanceId;

    private GetVpcFirewallsFirewallPeerVpcPeerVpcCidrTableListPeerRouteEntryList() {}
    /**
     * @return The target network segment of the peer VPC.
     * 
     */
    public String peerDestinationCidr() {
        return this.peerDestinationCidr;
    }
    /**
     * @return The ID of the next-hop instance in the peer VPC.
     * 
     */
    public String peerNextHopInstanceId() {
        return this.peerNextHopInstanceId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcFirewallsFirewallPeerVpcPeerVpcCidrTableListPeerRouteEntryList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String peerDestinationCidr;
        private String peerNextHopInstanceId;
        public Builder() {}
        public Builder(GetVpcFirewallsFirewallPeerVpcPeerVpcCidrTableListPeerRouteEntryList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.peerDestinationCidr = defaults.peerDestinationCidr;
    	      this.peerNextHopInstanceId = defaults.peerNextHopInstanceId;
        }

        @CustomType.Setter
        public Builder peerDestinationCidr(String peerDestinationCidr) {
            this.peerDestinationCidr = Objects.requireNonNull(peerDestinationCidr);
            return this;
        }
        @CustomType.Setter
        public Builder peerNextHopInstanceId(String peerNextHopInstanceId) {
            this.peerNextHopInstanceId = Objects.requireNonNull(peerNextHopInstanceId);
            return this;
        }
        public GetVpcFirewallsFirewallPeerVpcPeerVpcCidrTableListPeerRouteEntryList build() {
            final var o = new GetVpcFirewallsFirewallPeerVpcPeerVpcCidrTableListPeerRouteEntryList();
            o.peerDestinationCidr = peerDestinationCidr;
            o.peerNextHopInstanceId = peerNextHopInstanceId;
            return o;
        }
    }
}
