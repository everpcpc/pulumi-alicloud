// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.GetAclsAclAclEntry;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAclsAcl {
    /**
     * @return The entries of the Acl.
     * 
     */
    private List<GetAclsAclAclEntry> aclEntries;
    /**
     * @return The  ID of the Acl.
     * 
     */
    private String aclId;
    /**
     * @return The name of the acl.
     * 
     */
    private String aclName;
    /**
     * @return The address ip version.
     * 
     */
    private String addressIpVersion;
    /**
     * @return The ID of the Acl. Its value is same as `acl_id`.
     * 
     */
    private String id;
    /**
     * @return The status of the resource.
     * 
     */
    private String status;

    private GetAclsAcl() {}
    /**
     * @return The entries of the Acl.
     * 
     */
    public List<GetAclsAclAclEntry> aclEntries() {
        return this.aclEntries;
    }
    /**
     * @return The  ID of the Acl.
     * 
     */
    public String aclId() {
        return this.aclId;
    }
    /**
     * @return The name of the acl.
     * 
     */
    public String aclName() {
        return this.aclName;
    }
    /**
     * @return The address ip version.
     * 
     */
    public String addressIpVersion() {
        return this.addressIpVersion;
    }
    /**
     * @return The ID of the Acl. Its value is same as `acl_id`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The status of the resource.
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
        private String status;
        public Builder() {}
        public Builder(GetAclsAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aclEntries = defaults.aclEntries;
    	      this.aclId = defaults.aclId;
    	      this.aclName = defaults.aclName;
    	      this.addressIpVersion = defaults.addressIpVersion;
    	      this.id = defaults.id;
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
            o.status = status;
            return o;
        }
    }
}
