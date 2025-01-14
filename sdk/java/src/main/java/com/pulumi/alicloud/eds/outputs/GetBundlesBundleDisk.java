// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBundlesBundleDisk {
    /**
     * @return The disk size attribute of the bundle.
     * 
     */
    private String diskSize;
    /**
     * @return The disk type attribute of the bundle.
     * 
     */
    private String diskType;

    private GetBundlesBundleDisk() {}
    /**
     * @return The disk size attribute of the bundle.
     * 
     */
    public String diskSize() {
        return this.diskSize;
    }
    /**
     * @return The disk type attribute of the bundle.
     * 
     */
    public String diskType() {
        return this.diskType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBundlesBundleDisk defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String diskSize;
        private String diskType;
        public Builder() {}
        public Builder(GetBundlesBundleDisk defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.diskSize = defaults.diskSize;
    	      this.diskType = defaults.diskType;
        }

        @CustomType.Setter
        public Builder diskSize(String diskSize) {
            this.diskSize = Objects.requireNonNull(diskSize);
            return this;
        }
        @CustomType.Setter
        public Builder diskType(String diskType) {
            this.diskType = Objects.requireNonNull(diskType);
            return this;
        }
        public GetBundlesBundleDisk build() {
            final var o = new GetBundlesBundleDisk();
            o.diskSize = diskSize;
            o.diskType = diskType;
            return o;
        }
    }
}
