// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall.outputs;

import com.pulumi.alicloud.cloudfirewall.outputs.FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class FirewallVpcFirewallLocalVpcLocalVpcCidrTableList {
    /**
     * @return The list of route entries of the local VPC.See the following `Block LocalRouteEntryList`.
     * 
     */
    private List<FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList> localRouteEntryLists;
    /**
     * @return The ID of the route table of the local VPC.
     * 
     */
    private String localRouteTableId;

    private FirewallVpcFirewallLocalVpcLocalVpcCidrTableList() {}
    /**
     * @return The list of route entries of the local VPC.See the following `Block LocalRouteEntryList`.
     * 
     */
    public List<FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList> localRouteEntryLists() {
        return this.localRouteEntryLists;
    }
    /**
     * @return The ID of the route table of the local VPC.
     * 
     */
    public String localRouteTableId() {
        return this.localRouteTableId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FirewallVpcFirewallLocalVpcLocalVpcCidrTableList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList> localRouteEntryLists;
        private String localRouteTableId;
        public Builder() {}
        public Builder(FirewallVpcFirewallLocalVpcLocalVpcCidrTableList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.localRouteEntryLists = defaults.localRouteEntryLists;
    	      this.localRouteTableId = defaults.localRouteTableId;
        }

        @CustomType.Setter
        public Builder localRouteEntryLists(List<FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList> localRouteEntryLists) {
            this.localRouteEntryLists = Objects.requireNonNull(localRouteEntryLists);
            return this;
        }
        public Builder localRouteEntryLists(FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList... localRouteEntryLists) {
            return localRouteEntryLists(List.of(localRouteEntryLists));
        }
        @CustomType.Setter
        public Builder localRouteTableId(String localRouteTableId) {
            this.localRouteTableId = Objects.requireNonNull(localRouteTableId);
            return this;
        }
        public FirewallVpcFirewallLocalVpcLocalVpcCidrTableList build() {
            final var o = new FirewallVpcFirewallLocalVpcLocalVpcCidrTableList();
            o.localRouteEntryLists = localRouteEntryLists;
            o.localRouteTableId = localRouteTableId;
            return o;
        }
    }
}
