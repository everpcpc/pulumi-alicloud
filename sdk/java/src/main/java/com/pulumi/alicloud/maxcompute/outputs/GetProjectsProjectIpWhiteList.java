// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProjectIpWhiteList {
    /**
     * @return Classic network IP white list.
     * 
     */
    private String ipList;
    /**
     * @return VPC network whitelist.
     * 
     */
    private String vpcIpList;

    private GetProjectsProjectIpWhiteList() {}
    /**
     * @return Classic network IP white list.
     * 
     */
    public String ipList() {
        return this.ipList;
    }
    /**
     * @return VPC network whitelist.
     * 
     */
    public String vpcIpList() {
        return this.vpcIpList;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProjectIpWhiteList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ipList;
        private String vpcIpList;
        public Builder() {}
        public Builder(GetProjectsProjectIpWhiteList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipList = defaults.ipList;
    	      this.vpcIpList = defaults.vpcIpList;
        }

        @CustomType.Setter
        public Builder ipList(String ipList) {
            this.ipList = Objects.requireNonNull(ipList);
            return this;
        }
        @CustomType.Setter
        public Builder vpcIpList(String vpcIpList) {
            this.vpcIpList = Objects.requireNonNull(vpcIpList);
            return this;
        }
        public GetProjectsProjectIpWhiteList build() {
            final var o = new GetProjectsProjectIpWhiteList();
            o.ipList = ipList;
            o.vpcIpList = vpcIpList;
            return o;
        }
    }
}
