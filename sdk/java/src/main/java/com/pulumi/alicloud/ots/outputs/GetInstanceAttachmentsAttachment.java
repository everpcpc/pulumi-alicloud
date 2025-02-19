// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstanceAttachmentsAttachment {
    /**
     * @return The domain of the instance attachment.
     * 
     */
    private String domain;
    /**
     * @return The access endpoint of the instance attachment.
     * 
     */
    private String endpoint;
    /**
     * @return The resource ID, the value is same as &#34;instance_name&#34;.
     * 
     */
    private String id;
    /**
     * @return The name of OTS instance.
     * 
     */
    private String instanceName;
    /**
     * @return The region of the instance attachment.
     * 
     */
    private String region;
    /**
     * @return The ID of attaching VPC to instance.
     * 
     */
    private String vpcId;
    /**
     * @return The name of attaching VPC to instance.
     * 
     */
    private String vpcName;

    private GetInstanceAttachmentsAttachment() {}
    /**
     * @return The domain of the instance attachment.
     * 
     */
    public String domain() {
        return this.domain;
    }
    /**
     * @return The access endpoint of the instance attachment.
     * 
     */
    public String endpoint() {
        return this.endpoint;
    }
    /**
     * @return The resource ID, the value is same as &#34;instance_name&#34;.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of OTS instance.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }
    /**
     * @return The region of the instance attachment.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The ID of attaching VPC to instance.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The name of attaching VPC to instance.
     * 
     */
    public String vpcName() {
        return this.vpcName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceAttachmentsAttachment defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private String endpoint;
        private String id;
        private String instanceName;
        private String region;
        private String vpcId;
        private String vpcName;
        public Builder() {}
        public Builder(GetInstanceAttachmentsAttachment defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.endpoint = defaults.endpoint;
    	      this.id = defaults.id;
    	      this.instanceName = defaults.instanceName;
    	      this.region = defaults.region;
    	      this.vpcId = defaults.vpcId;
    	      this.vpcName = defaults.vpcName;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vpcName(String vpcName) {
            this.vpcName = Objects.requireNonNull(vpcName);
            return this;
        }
        public GetInstanceAttachmentsAttachment build() {
            final var o = new GetInstanceAttachmentsAttachment();
            o.domain = domain;
            o.endpoint = endpoint;
            o.id = id;
            o.instanceName = instanceName;
            o.region = region;
            o.vpcId = vpcId;
            o.vpcName = vpcName;
            return o;
        }
    }
}
