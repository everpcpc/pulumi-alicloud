// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.alicloud.alb.outputs.GetAclsAclAclEntry;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAclsAcl {
    /**
     * @return ACL Entries.
     * 
     */
    private List<GetAclsAclAclEntry> aclEntries;
    /**
     * @return Access Control Policy ID.
     * 
     */
    private String aclId;
    /**
     * @return The ACL Name.
     * 
     */
    private String aclName;
    /**
     * @return Address Protocol Version.
     * 
     */
    private String addressIpVersion;
    /**
     * @return The ID of the Acl.
     * 
     */
    private String id;
    /**
     * @return Resource Group to Which the Number.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The state of the ACL. Valid values:`Provisioning` , `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     * 
     */
    private String status;

    private GetAclsAcl() {}
    /**
     * @return ACL Entries.
     * 
     */
    public List<GetAclsAclAclEntry> aclEntries() {
        return this.aclEntries;
    }
    /**
     * @return Access Control Policy ID.
     * 
     */
    public String aclId() {
        return this.aclId;
    }
    /**
     * @return The ACL Name.
     * 
     */
    public String aclName() {
        return this.aclName;
    }
    /**
     * @return Address Protocol Version.
     * 
     */
    public String addressIpVersion() {
        return this.addressIpVersion;
    }
    /**
     * @return The ID of the Acl.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Resource Group to Which the Number.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The state of the ACL. Valid values:`Provisioning` , `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAclsAcl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetAclsAclAclEntry> aclEntries;
        private String aclId;
        private String aclName;
        private String addressIpVersion;
        private String id;
        private String resourceGroupId;
        private String status;
        public Builder() {}
        public Builder(GetAclsAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aclEntries = defaults.aclEntries;
    	      this.aclId = defaults.aclId;
    	      this.aclName = defaults.aclName;
    	      this.addressIpVersion = defaults.addressIpVersion;
    	      this.id = defaults.id;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder aclEntries(List<GetAclsAclAclEntry> aclEntries) {
            this.aclEntries = Objects.requireNonNull(aclEntries);
            return this;
        }
        public Builder aclEntries(GetAclsAclAclEntry... aclEntries) {
            return aclEntries(List.of(aclEntries));
        }
        @CustomType.Setter
        public Builder aclId(String aclId) {
            this.aclId = Objects.requireNonNull(aclId);
            return this;
        }
        @CustomType.Setter
        public Builder aclName(String aclName) {
            this.aclName = Objects.requireNonNull(aclName);
            return this;
        }
        @CustomType.Setter
        public Builder addressIpVersion(String addressIpVersion) {
            this.addressIpVersion = Objects.requireNonNull(addressIpVersion);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            this.resourceGroupId = Objects.requireNonNull(resourceGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetAclsAcl build() {
            final var o = new GetAclsAcl();
            o.aclEntries = aclEntries;
            o.aclId = aclId;
            o.aclName = aclName;
            o.addressIpVersion = addressIpVersion;
            o.id = id;
            o.resourceGroupId = resourceGroupId;
            o.status = status;
            return o;
        }
    }
}
