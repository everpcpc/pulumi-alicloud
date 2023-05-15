// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAssetsAsset {
    /**
     * @return The creation time of the resource
     * 
     */
    private String createTime;
    /**
     * @return The ID of the instance.
     * 
     */
    private String id;
    /**
     * @return The UUID of the instance.
     * 
     */
    private String uuid;

    private GetAssetsAsset() {}
    /**
     * @return The creation time of the resource
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the instance.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The UUID of the instance.
     * 
     */
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAssetsAsset defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createTime;
        private String id;
        private String uuid;
        public Builder() {}
        public Builder(GetAssetsAsset defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            this.uuid = Objects.requireNonNull(uuid);
            return this;
        }
        public GetAssetsAsset build() {
            final var o = new GetAssetsAsset();
            o.createTime = createTime;
            o.id = id;
            o.uuid = uuid;
            return o;
        }
    }
}
