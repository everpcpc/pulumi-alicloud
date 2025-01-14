// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerBackupPlanDetailArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerBackupPlanDetailArgs Empty = new ServerBackupPlanDetailArgs();

    /**
     * Whether to turn on application consistency. The application consistency snapshot backs up memory data and ongoing database transactions at the time of snapshot creation to ensure the consistency of application system data and database transactions. By applying consistent snapshots, there is no data damage or loss, so as to avoid log rollback during database startup and ensure that the application is in a consistent startup state. Valid values: `true`, `false`.
     * 
     */
    @Import(name="appConsistent", required=true)
    private Output<Boolean> appConsistent;

    /**
     * @return Whether to turn on application consistency. The application consistency snapshot backs up memory data and ongoing database transactions at the time of snapshot creation to ensure the consistency of application system data and database transactions. By applying consistent snapshots, there is no data damage or loss, so as to avoid log rollback during database startup and ensure that the application is in a consistent startup state. Valid values: `true`, `false`.
     * 
     */
    public Output<Boolean> appConsistent() {
        return this.appConsistent;
    }

    /**
     * Only vaild when DoCopy is true. The destination region ID when replicating to another region. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    @Import(name="destinationRegionId")
    private @Nullable Output<String> destinationRegionId;

    /**
     * @return Only vaild when DoCopy is true. The destination region ID when replicating to another region. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    public Optional<Output<String>> destinationRegionId() {
        return Optional.ofNullable(this.destinationRegionId);
    }

    /**
     * Only vaild when DoCopy is true. The retention days of the destination backup. When not specified, the destination backup will be saved permanently. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    @Import(name="destinationRetention")
    private @Nullable Output<Integer> destinationRetention;

    /**
     * @return Only vaild when DoCopy is true. The retention days of the destination backup. When not specified, the destination backup will be saved permanently. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    public Optional<Output<Integer>> destinationRetention() {
        return Optional.ofNullable(this.destinationRetention);
    }

    /**
     * The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
     * 
     */
    @Import(name="diskIdLists")
    private @Nullable Output<List<String>> diskIdLists;

    /**
     * @return The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
     * 
     */
    public Optional<Output<List<String>>> diskIdLists() {
        return Optional.ofNullable(this.diskIdLists);
    }

    /**
     * Whether replicate to another region. Valid values: `true`, `false`.
     * 
     */
    @Import(name="doCopy")
    private @Nullable Output<Boolean> doCopy;

    /**
     * @return Whether replicate to another region. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> doCopy() {
        return Optional.ofNullable(this.doCopy);
    }

    /**
     * Only the Linux system is valid. Whether to use the Linux FsFreeze mechanism to ensure that the file system is read-only consistent before creating a storage snapshot. The default is True. Valid values: `true`, `false`.
     * 
     */
    @Import(name="enableFsFreeze")
    private @Nullable Output<Boolean> enableFsFreeze;

    /**
     * @return Only the Linux system is valid. Whether to use the Linux FsFreeze mechanism to ensure that the file system is read-only consistent before creating a storage snapshot. The default is True. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> enableFsFreeze() {
        return Optional.ofNullable(this.enableFsFreeze);
    }

    /**
     * Only vaild for the linux system when AppConsistent is true. The application thaw script path (e.g. /tmp/postscript.sh). The postscript.sh script must meet the following conditions: in terms of permissions, only the root user as the owner has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    @Import(name="postScriptPath")
    private @Nullable Output<String> postScriptPath;

    /**
     * @return Only vaild for the linux system when AppConsistent is true. The application thaw script path (e.g. /tmp/postscript.sh). The postscript.sh script must meet the following conditions: in terms of permissions, only the root user as the owner has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    public Optional<Output<String>> postScriptPath() {
        return Optional.ofNullable(this.postScriptPath);
    }

    /**
     * Only vaild for the linux system when AppConsistent is true. Apply the freeze script path (e.g. /tmp/prescript.sh). prescript.sh scripts must meet the following conditions: in terms of permissions, only root, as the owner, has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    @Import(name="preScriptPath")
    private @Nullable Output<String> preScriptPath;

    /**
     * @return Only vaild for the linux system when AppConsistent is true. Apply the freeze script path (e.g. /tmp/prescript.sh). prescript.sh scripts must meet the following conditions: in terms of permissions, only root, as the owner, has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    public Optional<Output<String>> preScriptPath() {
        return Optional.ofNullable(this.preScriptPath);
    }

    /**
     * Whether to turn on file system consistency. If SnapshotGroup is true, when AppConsistent is true but the relevant conditions are not met or AppConsistent is false, the resulting snapshot will be a file system consistency snapshot. The file system consistency ensures that the file system memory and disk information are synchronized at the time of snapshot creation, and the file system write operation is frozen to make the file system in a consistent state. The file system consistency snapshot can prevent the operating system from performing disk inspection and repair operations such as CHKDSK or fsck after restart. Valid values: `true`, `false`.
     * 
     */
    @Import(name="snapshotGroup", required=true)
    private Output<Boolean> snapshotGroup;

    /**
     * @return Whether to turn on file system consistency. If SnapshotGroup is true, when AppConsistent is true but the relevant conditions are not met or AppConsistent is false, the resulting snapshot will be a file system consistency snapshot. The file system consistency ensures that the file system memory and disk information are synchronized at the time of snapshot creation, and the file system write operation is frozen to make the file system in a consistent state. The file system consistency snapshot can prevent the operating system from performing disk inspection and repair operations such as CHKDSK or fsck after restart. Valid values: `true`, `false`.
     * 
     */
    public Output<Boolean> snapshotGroup() {
        return this.snapshotGroup;
    }

    /**
     * Only the Linux system is valid, and the IO freeze timeout period. The default is 30 seconds.
     * 
     */
    @Import(name="timeoutInSeconds")
    private @Nullable Output<Integer> timeoutInSeconds;

    /**
     * @return Only the Linux system is valid, and the IO freeze timeout period. The default is 30 seconds.
     * 
     */
    public Optional<Output<Integer>> timeoutInSeconds() {
        return Optional.ofNullable(this.timeoutInSeconds);
    }

    private ServerBackupPlanDetailArgs() {}

    private ServerBackupPlanDetailArgs(ServerBackupPlanDetailArgs $) {
        this.appConsistent = $.appConsistent;
        this.destinationRegionId = $.destinationRegionId;
        this.destinationRetention = $.destinationRetention;
        this.diskIdLists = $.diskIdLists;
        this.doCopy = $.doCopy;
        this.enableFsFreeze = $.enableFsFreeze;
        this.postScriptPath = $.postScriptPath;
        this.preScriptPath = $.preScriptPath;
        this.snapshotGroup = $.snapshotGroup;
        this.timeoutInSeconds = $.timeoutInSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerBackupPlanDetailArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerBackupPlanDetailArgs $;

        public Builder() {
            $ = new ServerBackupPlanDetailArgs();
        }

        public Builder(ServerBackupPlanDetailArgs defaults) {
            $ = new ServerBackupPlanDetailArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appConsistent Whether to turn on application consistency. The application consistency snapshot backs up memory data and ongoing database transactions at the time of snapshot creation to ensure the consistency of application system data and database transactions. By applying consistent snapshots, there is no data damage or loss, so as to avoid log rollback during database startup and ensure that the application is in a consistent startup state. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder appConsistent(Output<Boolean> appConsistent) {
            $.appConsistent = appConsistent;
            return this;
        }

        /**
         * @param appConsistent Whether to turn on application consistency. The application consistency snapshot backs up memory data and ongoing database transactions at the time of snapshot creation to ensure the consistency of application system data and database transactions. By applying consistent snapshots, there is no data damage or loss, so as to avoid log rollback during database startup and ensure that the application is in a consistent startup state. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder appConsistent(Boolean appConsistent) {
            return appConsistent(Output.of(appConsistent));
        }

        /**
         * @param destinationRegionId Only vaild when DoCopy is true. The destination region ID when replicating to another region. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
         * 
         * @return builder
         * 
         */
        public Builder destinationRegionId(@Nullable Output<String> destinationRegionId) {
            $.destinationRegionId = destinationRegionId;
            return this;
        }

        /**
         * @param destinationRegionId Only vaild when DoCopy is true. The destination region ID when replicating to another region. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
         * 
         * @return builder
         * 
         */
        public Builder destinationRegionId(String destinationRegionId) {
            return destinationRegionId(Output.of(destinationRegionId));
        }

        /**
         * @param destinationRetention Only vaild when DoCopy is true. The retention days of the destination backup. When not specified, the destination backup will be saved permanently. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
         * 
         * @return builder
         * 
         */
        public Builder destinationRetention(@Nullable Output<Integer> destinationRetention) {
            $.destinationRetention = destinationRetention;
            return this;
        }

        /**
         * @param destinationRetention Only vaild when DoCopy is true. The retention days of the destination backup. When not specified, the destination backup will be saved permanently. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
         * 
         * @return builder
         * 
         */
        public Builder destinationRetention(Integer destinationRetention) {
            return destinationRetention(Output.of(destinationRetention));
        }

        /**
         * @param diskIdLists The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
         * 
         * @return builder
         * 
         */
        public Builder diskIdLists(@Nullable Output<List<String>> diskIdLists) {
            $.diskIdLists = diskIdLists;
            return this;
        }

        /**
         * @param diskIdLists The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
         * 
         * @return builder
         * 
         */
        public Builder diskIdLists(List<String> diskIdLists) {
            return diskIdLists(Output.of(diskIdLists));
        }

        /**
         * @param diskIdLists The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
         * 
         * @return builder
         * 
         */
        public Builder diskIdLists(String... diskIdLists) {
            return diskIdLists(List.of(diskIdLists));
        }

        /**
         * @param doCopy Whether replicate to another region. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder doCopy(@Nullable Output<Boolean> doCopy) {
            $.doCopy = doCopy;
            return this;
        }

        /**
         * @param doCopy Whether replicate to another region. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder doCopy(Boolean doCopy) {
            return doCopy(Output.of(doCopy));
        }

        /**
         * @param enableFsFreeze Only the Linux system is valid. Whether to use the Linux FsFreeze mechanism to ensure that the file system is read-only consistent before creating a storage snapshot. The default is True. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableFsFreeze(@Nullable Output<Boolean> enableFsFreeze) {
            $.enableFsFreeze = enableFsFreeze;
            return this;
        }

        /**
         * @param enableFsFreeze Only the Linux system is valid. Whether to use the Linux FsFreeze mechanism to ensure that the file system is read-only consistent before creating a storage snapshot. The default is True. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableFsFreeze(Boolean enableFsFreeze) {
            return enableFsFreeze(Output.of(enableFsFreeze));
        }

        /**
         * @param postScriptPath Only vaild for the linux system when AppConsistent is true. The application thaw script path (e.g. /tmp/postscript.sh). The postscript.sh script must meet the following conditions: in terms of permissions, only the root user as the owner has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
         * 
         * @return builder
         * 
         */
        public Builder postScriptPath(@Nullable Output<String> postScriptPath) {
            $.postScriptPath = postScriptPath;
            return this;
        }

        /**
         * @param postScriptPath Only vaild for the linux system when AppConsistent is true. The application thaw script path (e.g. /tmp/postscript.sh). The postscript.sh script must meet the following conditions: in terms of permissions, only the root user as the owner has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
         * 
         * @return builder
         * 
         */
        public Builder postScriptPath(String postScriptPath) {
            return postScriptPath(Output.of(postScriptPath));
        }

        /**
         * @param preScriptPath Only vaild for the linux system when AppConsistent is true. Apply the freeze script path (e.g. /tmp/prescript.sh). prescript.sh scripts must meet the following conditions: in terms of permissions, only root, as the owner, has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
         * 
         * @return builder
         * 
         */
        public Builder preScriptPath(@Nullable Output<String> preScriptPath) {
            $.preScriptPath = preScriptPath;
            return this;
        }

        /**
         * @param preScriptPath Only vaild for the linux system when AppConsistent is true. Apply the freeze script path (e.g. /tmp/prescript.sh). prescript.sh scripts must meet the following conditions: in terms of permissions, only root, as the owner, has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
         * 
         * @return builder
         * 
         */
        public Builder preScriptPath(String preScriptPath) {
            return preScriptPath(Output.of(preScriptPath));
        }

        /**
         * @param snapshotGroup Whether to turn on file system consistency. If SnapshotGroup is true, when AppConsistent is true but the relevant conditions are not met or AppConsistent is false, the resulting snapshot will be a file system consistency snapshot. The file system consistency ensures that the file system memory and disk information are synchronized at the time of snapshot creation, and the file system write operation is frozen to make the file system in a consistent state. The file system consistency snapshot can prevent the operating system from performing disk inspection and repair operations such as CHKDSK or fsck after restart. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotGroup(Output<Boolean> snapshotGroup) {
            $.snapshotGroup = snapshotGroup;
            return this;
        }

        /**
         * @param snapshotGroup Whether to turn on file system consistency. If SnapshotGroup is true, when AppConsistent is true but the relevant conditions are not met or AppConsistent is false, the resulting snapshot will be a file system consistency snapshot. The file system consistency ensures that the file system memory and disk information are synchronized at the time of snapshot creation, and the file system write operation is frozen to make the file system in a consistent state. The file system consistency snapshot can prevent the operating system from performing disk inspection and repair operations such as CHKDSK or fsck after restart. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotGroup(Boolean snapshotGroup) {
            return snapshotGroup(Output.of(snapshotGroup));
        }

        /**
         * @param timeoutInSeconds Only the Linux system is valid, and the IO freeze timeout period. The default is 30 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutInSeconds(@Nullable Output<Integer> timeoutInSeconds) {
            $.timeoutInSeconds = timeoutInSeconds;
            return this;
        }

        /**
         * @param timeoutInSeconds Only the Linux system is valid, and the IO freeze timeout period. The default is 30 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutInSeconds(Integer timeoutInSeconds) {
            return timeoutInSeconds(Output.of(timeoutInSeconds));
        }

        public ServerBackupPlanDetailArgs build() {
            $.appConsistent = Objects.requireNonNull($.appConsistent, "expected parameter 'appConsistent' to be non-null");
            $.snapshotGroup = Objects.requireNonNull($.snapshotGroup, "expected parameter 'snapshotGroup' to be non-null");
            return $;
        }
    }

}
