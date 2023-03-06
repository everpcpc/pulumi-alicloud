// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cddc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDedicatedHostGroupsGroupDedicatedHostCountGroupByHostType {
    private String placeHolder;

    private GetDedicatedHostGroupsGroupDedicatedHostCountGroupByHostType() {}
    public String placeHolder() {
        return this.placeHolder;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDedicatedHostGroupsGroupDedicatedHostCountGroupByHostType defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String placeHolder;
        public Builder() {}
        public Builder(GetDedicatedHostGroupsGroupDedicatedHostCountGroupByHostType defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.placeHolder = defaults.placeHolder;
        }

        @CustomType.Setter
        public Builder placeHolder(String placeHolder) {
            this.placeHolder = Objects.requireNonNull(placeHolder);
            return this;
        }
        public GetDedicatedHostGroupsGroupDedicatedHostCountGroupByHostType build() {
            final var o = new GetDedicatedHostGroupsGroupDedicatedHostCountGroupByHostType();
            o.placeHolder = placeHolder;
            return o;
        }
    }
}