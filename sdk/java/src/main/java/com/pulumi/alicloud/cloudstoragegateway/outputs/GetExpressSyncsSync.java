// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetExpressSyncsSync {
    /**
     * @return The name of the OSS Bucket.
     * 
     */
    private String bucketName;
    /**
     * @return The prefix of the OSS Bucket.
     * 
     */
    private String bucketPrefix;
    /**
     * @return The region of the OSS Bucket.
     * 
     */
    private String bucketRegion;
    /**
     * @return The description of the Express Sync.
     * 
     */
    private String description;
    /**
     * @return The ID of the Express Sync.
     * 
     */
    private String expressSyncId;
    /**
     * @return The name of the Express Sync.
     * 
     */
    private String expressSyncName;
    private String id;
    /**
     * @return The name of the message topic (Topic) corresponding to the Express Sync in the Alibaba Cloud Message Service MNS.
     * 
     */
    private String mnsTopic;

    private GetExpressSyncsSync() {}
    /**
     * @return The name of the OSS Bucket.
     * 
     */
    public String bucketName() {
        return this.bucketName;
    }
    /**
     * @return The prefix of the OSS Bucket.
     * 
     */
    public String bucketPrefix() {
        return this.bucketPrefix;
    }
    /**
     * @return The region of the OSS Bucket.
     * 
     */
    public String bucketRegion() {
        return this.bucketRegion;
    }
    /**
     * @return The description of the Express Sync.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the Express Sync.
     * 
     */
    public String expressSyncId() {
        return this.expressSyncId;
    }
    /**
     * @return The name of the Express Sync.
     * 
     */
    public String expressSyncName() {
        return this.expressSyncName;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the message topic (Topic) corresponding to the Express Sync in the Alibaba Cloud Message Service MNS.
     * 
     */
    public String mnsTopic() {
        return this.mnsTopic;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetExpressSyncsSync defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bucketName;
        private String bucketPrefix;
        private String bucketRegion;
        private String description;
        private String expressSyncId;
        private String expressSyncName;
        private String id;
        private String mnsTopic;
        public Builder() {}
        public Builder(GetExpressSyncsSync defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucketName = defaults.bucketName;
    	      this.bucketPrefix = defaults.bucketPrefix;
    	      this.bucketRegion = defaults.bucketRegion;
    	      this.description = defaults.description;
    	      this.expressSyncId = defaults.expressSyncId;
    	      this.expressSyncName = defaults.expressSyncName;
    	      this.id = defaults.id;
    	      this.mnsTopic = defaults.mnsTopic;
        }

        @CustomType.Setter
        public Builder bucketName(String bucketName) {
            this.bucketName = Objects.requireNonNull(bucketName);
            return this;
        }
        @CustomType.Setter
        public Builder bucketPrefix(String bucketPrefix) {
            this.bucketPrefix = Objects.requireNonNull(bucketPrefix);
            return this;
        }
        @CustomType.Setter
        public Builder bucketRegion(String bucketRegion) {
            this.bucketRegion = Objects.requireNonNull(bucketRegion);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder expressSyncId(String expressSyncId) {
            this.expressSyncId = Objects.requireNonNull(expressSyncId);
            return this;
        }
        @CustomType.Setter
        public Builder expressSyncName(String expressSyncName) {
            this.expressSyncName = Objects.requireNonNull(expressSyncName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder mnsTopic(String mnsTopic) {
            this.mnsTopic = Objects.requireNonNull(mnsTopic);
            return this;
        }
        public GetExpressSyncsSync build() {
            final var o = new GetExpressSyncsSync();
            o.bucketName = bucketName;
            o.bucketPrefix = bucketPrefix;
            o.bucketRegion = bucketRegion;
            o.description = description;
            o.expressSyncId = expressSyncId;
            o.expressSyncName = expressSyncName;
            o.id = id;
            o.mnsTopic = mnsTopic;
            return o;
        }
    }
}
