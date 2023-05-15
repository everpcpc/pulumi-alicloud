// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.alicloud.threatdetection.outputs.GetHoneypotProbesProbeHoneypotBindListBindPortList;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetHoneypotProbesProbeHoneypotBindList {
    /**
     * @return List of listening ports. Available when `enable_details` is on.
     * 
     */
    private List<GetHoneypotProbesProbeHoneypotBindListBindPortList> bindPortLists;
    private String honeypotId;

    private GetHoneypotProbesProbeHoneypotBindList() {}
    /**
     * @return List of listening ports. Available when `enable_details` is on.
     * 
     */
    public List<GetHoneypotProbesProbeHoneypotBindListBindPortList> bindPortLists() {
        return this.bindPortLists;
    }
    public String honeypotId() {
        return this.honeypotId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHoneypotProbesProbeHoneypotBindList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetHoneypotProbesProbeHoneypotBindListBindPortList> bindPortLists;
        private String honeypotId;
        public Builder() {}
        public Builder(GetHoneypotProbesProbeHoneypotBindList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bindPortLists = defaults.bindPortLists;
    	      this.honeypotId = defaults.honeypotId;
        }

        @CustomType.Setter
        public Builder bindPortLists(List<GetHoneypotProbesProbeHoneypotBindListBindPortList> bindPortLists) {
            this.bindPortLists = Objects.requireNonNull(bindPortLists);
            return this;
        }
        public Builder bindPortLists(GetHoneypotProbesProbeHoneypotBindListBindPortList... bindPortLists) {
            return bindPortLists(List.of(bindPortLists));
        }
        @CustomType.Setter
        public Builder honeypotId(String honeypotId) {
            this.honeypotId = Objects.requireNonNull(honeypotId);
            return this;
        }
        public GetHoneypotProbesProbeHoneypotBindList build() {
            final var o = new GetHoneypotProbesProbeHoneypotBindList();
            o.bindPortLists = bindPortLists;
            o.honeypotId = honeypotId;
            return o;
        }
    }
}
