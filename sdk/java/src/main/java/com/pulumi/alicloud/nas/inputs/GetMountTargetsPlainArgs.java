// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetMountTargetsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMountTargetsPlainArgs Empty = new GetMountTargetsPlainArgs();

    /**
     * Filter results by a specific AccessGroupName.
     * 
     */
    @Import(name="accessGroupName")
    private @Nullable String accessGroupName;

    /**
     * @return Filter results by a specific AccessGroupName.
     * 
     */
    public Optional<String> accessGroupName() {
        return Optional.ofNullable(this.accessGroupName);
    }

    /**
     * The ID of the FileSystem that owns the MountTarget.
     * 
     */
    @Import(name="fileSystemId", required=true)
    private String fileSystemId;

    /**
     * @return The ID of the FileSystem that owns the MountTarget.
     * 
     */
    public String fileSystemId() {
        return this.fileSystemId;
    }

    /**
     * A list of MountTargetDomain.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of MountTargetDomain.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Field `mount_target_domain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
     * 
     * @deprecated
     * Field &#39;mount_target_domain&#39; has been deprecated from provider version 1.53.0. New field &#39;ids&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it. */
    @Import(name="mountTargetDomain")
    private @Nullable String mountTargetDomain;

    /**
     * @return Field `mount_target_domain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
     * 
     * @deprecated
     * Field &#39;mount_target_domain&#39; has been deprecated from provider version 1.53.0. New field &#39;ids&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it. */
    public Optional<String> mountTargetDomain() {
        return Optional.ofNullable(this.mountTargetDomain);
    }

    /**
     * Filter results by a specific NetworkType.
     * 
     */
    @Import(name="networkType")
    private @Nullable String networkType;

    /**
     * @return Filter results by a specific NetworkType.
     * 
     */
    public Optional<String> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Field `type` has been deprecated from provider version 1.95.0. New field `network_type` replaces it.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.95.0. New field &#39;network_type&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it. */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return Field `type` has been deprecated from provider version 1.95.0. New field `network_type` replaces it.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.95.0. New field &#39;network_type&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it. */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Filter results by a specific VpcId.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return Filter results by a specific VpcId.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * Filter results by a specific VSwitchId.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable String vswitchId;

    /**
     * @return Filter results by a specific VSwitchId.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetMountTargetsPlainArgs() {}

    private GetMountTargetsPlainArgs(GetMountTargetsPlainArgs $) {
        this.accessGroupName = $.accessGroupName;
        this.fileSystemId = $.fileSystemId;
        this.ids = $.ids;
        this.mountTargetDomain = $.mountTargetDomain;
        this.networkType = $.networkType;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.type = $.type;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMountTargetsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMountTargetsPlainArgs $;

        public Builder() {
            $ = new GetMountTargetsPlainArgs();
        }

        public Builder(GetMountTargetsPlainArgs defaults) {
            $ = new GetMountTargetsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessGroupName Filter results by a specific AccessGroupName.
         * 
         * @return builder
         * 
         */
        public Builder accessGroupName(@Nullable String accessGroupName) {
            $.accessGroupName = accessGroupName;
            return this;
        }

        /**
         * @param fileSystemId The ID of the FileSystem that owns the MountTarget.
         * 
         * @return builder
         * 
         */
        public Builder fileSystemId(String fileSystemId) {
            $.fileSystemId = fileSystemId;
            return this;
        }

        /**
         * @param ids A list of MountTargetDomain.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of MountTargetDomain.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param mountTargetDomain Field `mount_target_domain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;mount_target_domain&#39; has been deprecated from provider version 1.53.0. New field &#39;ids&#39; replaces it.
         * 
         */
        @Deprecated /* Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it. */
        public Builder mountTargetDomain(@Nullable String mountTargetDomain) {
            $.mountTargetDomain = mountTargetDomain;
            return this;
        }

        /**
         * @param networkType Filter results by a specific NetworkType.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable String networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param status Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param type Field `type` has been deprecated from provider version 1.95.0. New field `network_type` replaces it.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;type&#39; has been deprecated from provider version 1.95.0. New field &#39;network_type&#39; replaces it.
         * 
         */
        @Deprecated /* Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it. */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        /**
         * @param vpcId Filter results by a specific VpcId.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vswitchId Filter results by a specific VSwitchId.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable String vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        public GetMountTargetsPlainArgs build() {
            $.fileSystemId = Objects.requireNonNull($.fileSystemId, "expected parameter 'fileSystemId' to be non-null");
            return $;
        }
    }

}
