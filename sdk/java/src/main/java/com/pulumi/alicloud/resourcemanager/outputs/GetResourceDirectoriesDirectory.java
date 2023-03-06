// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetResourceDirectoriesDirectory {
    /**
     * @return The ID of resource directory.
     * 
     */
    private String id;
    /**
     * @return The ID of the master account.
     * 
     */
    private String masterAccountId;
    /**
     * @return The name of the master account.
     * 
     */
    private String masterAccountName;
    /**
     * @return The ID of the resource directory.
     * 
     */
    private String resourceDirectoryId;
    /**
     * @return The ID of the root folder.
     * 
     */
    private String rootFolderId;
    /**
     * @return (Available in 1.120.0+.) The status of the control policy.
     * 
     */
    private String status;

    private GetResourceDirectoriesDirectory() {}
    /**
     * @return The ID of resource directory.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the master account.
     * 
     */
    public String masterAccountId() {
        return this.masterAccountId;
    }
    /**
     * @return The name of the master account.
     * 
     */
    public String masterAccountName() {
        return this.masterAccountName;
    }
    /**
     * @return The ID of the resource directory.
     * 
     */
    public String resourceDirectoryId() {
        return this.resourceDirectoryId;
    }
    /**
     * @return The ID of the root folder.
     * 
     */
    public String rootFolderId() {
        return this.rootFolderId;
    }
    /**
     * @return (Available in 1.120.0+.) The status of the control policy.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetResourceDirectoriesDirectory defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String masterAccountId;
        private String masterAccountName;
        private String resourceDirectoryId;
        private String rootFolderId;
        private String status;
        public Builder() {}
        public Builder(GetResourceDirectoriesDirectory defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.masterAccountId = defaults.masterAccountId;
    	      this.masterAccountName = defaults.masterAccountName;
    	      this.resourceDirectoryId = defaults.resourceDirectoryId;
    	      this.rootFolderId = defaults.rootFolderId;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder masterAccountId(String masterAccountId) {
            this.masterAccountId = Objects.requireNonNull(masterAccountId);
            return this;
        }
        @CustomType.Setter
        public Builder masterAccountName(String masterAccountName) {
            this.masterAccountName = Objects.requireNonNull(masterAccountName);
            return this;
        }
        @CustomType.Setter
        public Builder resourceDirectoryId(String resourceDirectoryId) {
            this.resourceDirectoryId = Objects.requireNonNull(resourceDirectoryId);
            return this;
        }
        @CustomType.Setter
        public Builder rootFolderId(String rootFolderId) {
            this.rootFolderId = Objects.requireNonNull(rootFolderId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetResourceDirectoriesDirectory build() {
            final var o = new GetResourceDirectoriesDirectory();
            o.id = id;
            o.masterAccountId = masterAccountId;
            o.masterAccountName = masterAccountName;
            o.resourceDirectoryId = resourceDirectoryId;
            o.rootFolderId = rootFolderId;
            o.status = status;
            return o;
        }
    }
}