// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.alicloud.oss.outputs.GetBucketsBucketCorsRule;
import com.pulumi.alicloud.oss.outputs.GetBucketsBucketLifecycleRule;
import com.pulumi.alicloud.oss.outputs.GetBucketsBucketLogging;
import com.pulumi.alicloud.oss.outputs.GetBucketsBucketRefererConfig;
import com.pulumi.alicloud.oss.outputs.GetBucketsBucketServerSideEncryptionRule;
import com.pulumi.alicloud.oss.outputs.GetBucketsBucketVersioning;
import com.pulumi.alicloud.oss.outputs.GetBucketsBucketWebsite;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBucketsBucket {
    /**
     * @return Bucket access control list. Possible values: `private`, `public-read` and `public-read-write`.
     * 
     */
    private String acl;
    /**
     * @return A list of CORS rule configurations. Each element contains the following attributes:
     * 
     */
    private List<GetBucketsBucketCorsRule> corsRules;
    /**
     * @return Bucket creation date.
     * 
     */
    private String creationDate;
    /**
     * @return Internet domain name for accessing the bucket from outside.
     * 
     */
    private String extranetEndpoint;
    /**
     * @return Intranet domain name for accessing the bucket from an ECS instance in the same region.
     * 
     */
    private String intranetEndpoint;
    /**
     * @return A list CORS of lifecycle configurations. When Lifecycle is enabled, OSS automatically deletes the objects or transitions the objects (to another storage class) corresponding the lifecycle rules on a regular basis. Each element contains the following attributes:
     * 
     */
    private List<GetBucketsBucketLifecycleRule> lifecycleRules;
    /**
     * @return Region of the data center where the bucket is located.
     * 
     */
    private String location;
    /**
     * @return A list of one element containing configuration parameters used for storing access log information. It contains the following attributes:
     * 
     */
    private GetBucketsBucketLogging logging;
    /**
     * @return Bucket name.
     * 
     */
    private String name;
    /**
     * @return Bucket owner.
     * 
     */
    private String owner;
    private @Nullable String policy;
    /**
     * @return Redundancy type. Possible values: `LRS`, and `ZRS`.
     * 
     */
    private String redundancyType;
    /**
     * @return A list of one element containing referer configuration. It contains the following attributes:
     * 
     */
    private GetBucketsBucketRefererConfig refererConfig;
    /**
     * @return A configuration of default encryption for a bucket. It contains the following attributes:
     * 
     */
    private GetBucketsBucketServerSideEncryptionRule serverSideEncryptionRule;
    /**
     * @return Object storage type. Possible values: `Standard`, `IA`, `Archive` and `ColdArchive`.
     * 
     */
    private String storageClass;
    /**
     * @return A mapping of tags.
     * 
     */
    private Map<String,Object> tags;
    /**
     * @return If present , the versioning state has been set on the bucket. It contains the following attribute.
     * 
     */
    private GetBucketsBucketVersioning versioning;
    /**
     * @return A list of one element containing configuration parameters used when the bucket is used as a website. It contains the following attributes:
     * 
     */
    private GetBucketsBucketWebsite website;

    private GetBucketsBucket() {}
    /**
     * @return Bucket access control list. Possible values: `private`, `public-read` and `public-read-write`.
     * 
     */
    public String acl() {
        return this.acl;
    }
    /**
     * @return A list of CORS rule configurations. Each element contains the following attributes:
     * 
     */
    public List<GetBucketsBucketCorsRule> corsRules() {
        return this.corsRules;
    }
    /**
     * @return Bucket creation date.
     * 
     */
    public String creationDate() {
        return this.creationDate;
    }
    /**
     * @return Internet domain name for accessing the bucket from outside.
     * 
     */
    public String extranetEndpoint() {
        return this.extranetEndpoint;
    }
    /**
     * @return Intranet domain name for accessing the bucket from an ECS instance in the same region.
     * 
     */
    public String intranetEndpoint() {
        return this.intranetEndpoint;
    }
    /**
     * @return A list CORS of lifecycle configurations. When Lifecycle is enabled, OSS automatically deletes the objects or transitions the objects (to another storage class) corresponding the lifecycle rules on a regular basis. Each element contains the following attributes:
     * 
     */
    public List<GetBucketsBucketLifecycleRule> lifecycleRules() {
        return this.lifecycleRules;
    }
    /**
     * @return Region of the data center where the bucket is located.
     * 
     */
    public String location() {
        return this.location;
    }
    /**
     * @return A list of one element containing configuration parameters used for storing access log information. It contains the following attributes:
     * 
     */
    public GetBucketsBucketLogging logging() {
        return this.logging;
    }
    /**
     * @return Bucket name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Bucket owner.
     * 
     */
    public String owner() {
        return this.owner;
    }
    public Optional<String> policy() {
        return Optional.ofNullable(this.policy);
    }
    /**
     * @return Redundancy type. Possible values: `LRS`, and `ZRS`.
     * 
     */
    public String redundancyType() {
        return this.redundancyType;
    }
    /**
     * @return A list of one element containing referer configuration. It contains the following attributes:
     * 
     */
    public GetBucketsBucketRefererConfig refererConfig() {
        return this.refererConfig;
    }
    /**
     * @return A configuration of default encryption for a bucket. It contains the following attributes:
     * 
     */
    public GetBucketsBucketServerSideEncryptionRule serverSideEncryptionRule() {
        return this.serverSideEncryptionRule;
    }
    /**
     * @return Object storage type. Possible values: `Standard`, `IA`, `Archive` and `ColdArchive`.
     * 
     */
    public String storageClass() {
        return this.storageClass;
    }
    /**
     * @return A mapping of tags.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags;
    }
    /**
     * @return If present , the versioning state has been set on the bucket. It contains the following attribute.
     * 
     */
    public GetBucketsBucketVersioning versioning() {
        return this.versioning;
    }
    /**
     * @return A list of one element containing configuration parameters used when the bucket is used as a website. It contains the following attributes:
     * 
     */
    public GetBucketsBucketWebsite website() {
        return this.website;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketsBucket defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String acl;
        private List<GetBucketsBucketCorsRule> corsRules;
        private String creationDate;
        private String extranetEndpoint;
        private String intranetEndpoint;
        private List<GetBucketsBucketLifecycleRule> lifecycleRules;
        private String location;
        private GetBucketsBucketLogging logging;
        private String name;
        private String owner;
        private @Nullable String policy;
        private String redundancyType;
        private GetBucketsBucketRefererConfig refererConfig;
        private GetBucketsBucketServerSideEncryptionRule serverSideEncryptionRule;
        private String storageClass;
        private Map<String,Object> tags;
        private GetBucketsBucketVersioning versioning;
        private GetBucketsBucketWebsite website;
        public Builder() {}
        public Builder(GetBucketsBucket defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acl = defaults.acl;
    	      this.corsRules = defaults.corsRules;
    	      this.creationDate = defaults.creationDate;
    	      this.extranetEndpoint = defaults.extranetEndpoint;
    	      this.intranetEndpoint = defaults.intranetEndpoint;
    	      this.lifecycleRules = defaults.lifecycleRules;
    	      this.location = defaults.location;
    	      this.logging = defaults.logging;
    	      this.name = defaults.name;
    	      this.owner = defaults.owner;
    	      this.policy = defaults.policy;
    	      this.redundancyType = defaults.redundancyType;
    	      this.refererConfig = defaults.refererConfig;
    	      this.serverSideEncryptionRule = defaults.serverSideEncryptionRule;
    	      this.storageClass = defaults.storageClass;
    	      this.tags = defaults.tags;
    	      this.versioning = defaults.versioning;
    	      this.website = defaults.website;
        }

        @CustomType.Setter
        public Builder acl(String acl) {
            this.acl = Objects.requireNonNull(acl);
            return this;
        }
        @CustomType.Setter
        public Builder corsRules(List<GetBucketsBucketCorsRule> corsRules) {
            this.corsRules = Objects.requireNonNull(corsRules);
            return this;
        }
        public Builder corsRules(GetBucketsBucketCorsRule... corsRules) {
            return corsRules(List.of(corsRules));
        }
        @CustomType.Setter
        public Builder creationDate(String creationDate) {
            this.creationDate = Objects.requireNonNull(creationDate);
            return this;
        }
        @CustomType.Setter
        public Builder extranetEndpoint(String extranetEndpoint) {
            this.extranetEndpoint = Objects.requireNonNull(extranetEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder intranetEndpoint(String intranetEndpoint) {
            this.intranetEndpoint = Objects.requireNonNull(intranetEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder lifecycleRules(List<GetBucketsBucketLifecycleRule> lifecycleRules) {
            this.lifecycleRules = Objects.requireNonNull(lifecycleRules);
            return this;
        }
        public Builder lifecycleRules(GetBucketsBucketLifecycleRule... lifecycleRules) {
            return lifecycleRules(List.of(lifecycleRules));
        }
        @CustomType.Setter
        public Builder location(String location) {
            this.location = Objects.requireNonNull(location);
            return this;
        }
        @CustomType.Setter
        public Builder logging(GetBucketsBucketLogging logging) {
            this.logging = Objects.requireNonNull(logging);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder policy(@Nullable String policy) {
            this.policy = policy;
            return this;
        }
        @CustomType.Setter
        public Builder redundancyType(String redundancyType) {
            this.redundancyType = Objects.requireNonNull(redundancyType);
            return this;
        }
        @CustomType.Setter
        public Builder refererConfig(GetBucketsBucketRefererConfig refererConfig) {
            this.refererConfig = Objects.requireNonNull(refererConfig);
            return this;
        }
        @CustomType.Setter
        public Builder serverSideEncryptionRule(GetBucketsBucketServerSideEncryptionRule serverSideEncryptionRule) {
            this.serverSideEncryptionRule = Objects.requireNonNull(serverSideEncryptionRule);
            return this;
        }
        @CustomType.Setter
        public Builder storageClass(String storageClass) {
            this.storageClass = Objects.requireNonNull(storageClass);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder versioning(GetBucketsBucketVersioning versioning) {
            this.versioning = Objects.requireNonNull(versioning);
            return this;
        }
        @CustomType.Setter
        public Builder website(GetBucketsBucketWebsite website) {
            this.website = Objects.requireNonNull(website);
            return this;
        }
        public GetBucketsBucket build() {
            final var o = new GetBucketsBucket();
            o.acl = acl;
            o.corsRules = corsRules;
            o.creationDate = creationDate;
            o.extranetEndpoint = extranetEndpoint;
            o.intranetEndpoint = intranetEndpoint;
            o.lifecycleRules = lifecycleRules;
            o.location = location;
            o.logging = logging;
            o.name = name;
            o.owner = owner;
            o.policy = policy;
            o.redundancyType = redundancyType;
            o.refererConfig = refererConfig;
            o.serverSideEncryptionRule = serverSideEncryptionRule;
            o.storageClass = storageClass;
            o.tags = tags;
            o.versioning = versioning;
            o.website = website;
            return o;
        }
    }
}
