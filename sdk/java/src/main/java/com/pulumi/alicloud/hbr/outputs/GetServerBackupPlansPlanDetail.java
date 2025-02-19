// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetServerBackupPlansPlanDetail {
    /**
     * @return Whether to turn on application consistency. The application consistency snapshot backs up memory data and ongoing database transactions at the time of snapshot creation to ensure the consistency of application system data and database transactions. By applying consistent snapshots, there is no data damage or loss, so as to avoid log rollback during database startup and ensure that the application is in a consistent startup state. Valid values: `true`, `false`.
     * 
     */
    private Boolean appConsistent;
    /**
     * @return Only vaild when DoCopy is true. The destination region ID when replicating to another region. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    private String destinationRegionId;
    /**
     * @return Only vaild when DoCopy is true. The retention days of the destination backup. When not specified, the destination backup will be saved permanently. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    private Integer destinationRetention;
    /**
     * @return The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
     * 
     */
    private List<String> diskIdLists;
    /**
     * @return Whether replicate to another region. Valid values: `true`, `false`.
     * 
     */
    private Boolean doCopy;
    /**
     * @return Only the Linux system is valid. Whether to use the Linux FsFreeze mechanism to ensure that the file system is read-only consistent before creating a storage snapshot. The default is True. Valid values: `true`, `false`.
     * 
     */
    private Boolean enableFsFreeze;
    /**
     * @return Only vaild for the linux system when AppConsistent is true. The application thaw script path (e.g. /tmp/postscript.sh). The postscript.sh script must meet the following conditions: in terms of permissions, only the root user as the owner has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    private String postScriptPath;
    /**
     * @return Only vaild for the linux system when AppConsistent is true. Apply the freeze script path (e.g. /tmp/prescript.sh). prescript.sh scripts must meet the following conditions: in terms of permissions, only root, as the owner, has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    private String preScriptPath;
    /**
     * @return Whether to turn on file system consistency. If SnapshotGroup is true, when AppConsistent is true but the relevant conditions are not met or AppConsistent is false, the resulting snapshot will be a file system consistency snapshot. The file system consistency ensures that the file system memory and disk information are synchronized at the time of snapshot creation, and the file system write operation is frozen to make the file system in a consistent state. The file system consistency snapshot can prevent the operating system from performing disk inspection and repair operations such as CHKDSK or fsck after restart. Valid values: `true`, `false`.
     * 
     */
    private Boolean snapshotGroup;
    /**
     * @return Only the Linux system is valid, and the IO freeze timeout period. The default is 30 seconds.
     * 
     */
    private Integer timeoutInSeconds;

    private GetServerBackupPlansPlanDetail() {}
    /**
     * @return Whether to turn on application consistency. The application consistency snapshot backs up memory data and ongoing database transactions at the time of snapshot creation to ensure the consistency of application system data and database transactions. By applying consistent snapshots, there is no data damage or loss, so as to avoid log rollback during database startup and ensure that the application is in a consistent startup state. Valid values: `true`, `false`.
     * 
     */
    public Boolean appConsistent() {
        return this.appConsistent;
    }
    /**
     * @return Only vaild when DoCopy is true. The destination region ID when replicating to another region. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    public String destinationRegionId() {
        return this.destinationRegionId;
    }
    /**
     * @return Only vaild when DoCopy is true. The retention days of the destination backup. When not specified, the destination backup will be saved permanently. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    public Integer destinationRetention() {
        return this.destinationRetention;
    }
    /**
     * @return The list of cloud disks to be backed up in the ECS instance. When not specified, a snapshot is executed for all the disks on the ECS instance.
     * 
     */
    public List<String> diskIdLists() {
        return this.diskIdLists;
    }
    /**
     * @return Whether replicate to another region. Valid values: `true`, `false`.
     * 
     */
    public Boolean doCopy() {
        return this.doCopy;
    }
    /**
     * @return Only the Linux system is valid. Whether to use the Linux FsFreeze mechanism to ensure that the file system is read-only consistent before creating a storage snapshot. The default is True. Valid values: `true`, `false`.
     * 
     */
    public Boolean enableFsFreeze() {
        return this.enableFsFreeze;
    }
    /**
     * @return Only vaild for the linux system when AppConsistent is true. The application thaw script path (e.g. /tmp/postscript.sh). The postscript.sh script must meet the following conditions: in terms of permissions, only the root user as the owner has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    public String postScriptPath() {
        return this.postScriptPath;
    }
    /**
     * @return Only vaild for the linux system when AppConsistent is true. Apply the freeze script path (e.g. /tmp/prescript.sh). prescript.sh scripts must meet the following conditions: in terms of permissions, only root, as the owner, has read, write, and execute permissions, that is, 700 permissions. In terms of content, the script content needs to be customized according to the application itself. This indicates that this parameter must be set when creating an application consistency snapshot for a Linux instance. If the script is set incorrectly (for example, permissions, save path, or file name are set incorrectly), the resulting snapshot is a file system consistency snapshot.
     * 
     */
    public String preScriptPath() {
        return this.preScriptPath;
    }
    /**
     * @return Whether to turn on file system consistency. If SnapshotGroup is true, when AppConsistent is true but the relevant conditions are not met or AppConsistent is false, the resulting snapshot will be a file system consistency snapshot. The file system consistency ensures that the file system memory and disk information are synchronized at the time of snapshot creation, and the file system write operation is frozen to make the file system in a consistent state. The file system consistency snapshot can prevent the operating system from performing disk inspection and repair operations such as CHKDSK or fsck after restart. Valid values: `true`, `false`.
     * 
     */
    public Boolean snapshotGroup() {
        return this.snapshotGroup;
    }
    /**
     * @return Only the Linux system is valid, and the IO freeze timeout period. The default is 30 seconds.
     * 
     */
    public Integer timeoutInSeconds() {
        return this.timeoutInSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerBackupPlansPlanDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean appConsistent;
        private String destinationRegionId;
        private Integer destinationRetention;
        private List<String> diskIdLists;
        private Boolean doCopy;
        private Boolean enableFsFreeze;
        private String postScriptPath;
        private String preScriptPath;
        private Boolean snapshotGroup;
        private Integer timeoutInSeconds;
        public Builder() {}
        public Builder(GetServerBackupPlansPlanDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appConsistent = defaults.appConsistent;
    	      this.destinationRegionId = defaults.destinationRegionId;
    	      this.destinationRetention = defaults.destinationRetention;
    	      this.diskIdLists = defaults.diskIdLists;
    	      this.doCopy = defaults.doCopy;
    	      this.enableFsFreeze = defaults.enableFsFreeze;
    	      this.postScriptPath = defaults.postScriptPath;
    	      this.preScriptPath = defaults.preScriptPath;
    	      this.snapshotGroup = defaults.snapshotGroup;
    	      this.timeoutInSeconds = defaults.timeoutInSeconds;
        }

        @CustomType.Setter
        public Builder appConsistent(Boolean appConsistent) {
            this.appConsistent = Objects.requireNonNull(appConsistent);
            return this;
        }
        @CustomType.Setter
        public Builder destinationRegionId(String destinationRegionId) {
            this.destinationRegionId = Objects.requireNonNull(destinationRegionId);
            return this;
        }
        @CustomType.Setter
        public Builder destinationRetention(Integer destinationRetention) {
            this.destinationRetention = Objects.requireNonNull(destinationRetention);
            return this;
        }
        @CustomType.Setter
        public Builder diskIdLists(List<String> diskIdLists) {
            this.diskIdLists = Objects.requireNonNull(diskIdLists);
            return this;
        }
        public Builder diskIdLists(String... diskIdLists) {
            return diskIdLists(List.of(diskIdLists));
        }
        @CustomType.Setter
        public Builder doCopy(Boolean doCopy) {
            this.doCopy = Objects.requireNonNull(doCopy);
            return this;
        }
        @CustomType.Setter
        public Builder enableFsFreeze(Boolean enableFsFreeze) {
            this.enableFsFreeze = Objects.requireNonNull(enableFsFreeze);
            return this;
        }
        @CustomType.Setter
        public Builder postScriptPath(String postScriptPath) {
            this.postScriptPath = Objects.requireNonNull(postScriptPath);
            return this;
        }
        @CustomType.Setter
        public Builder preScriptPath(String preScriptPath) {
            this.preScriptPath = Objects.requireNonNull(preScriptPath);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotGroup(Boolean snapshotGroup) {
            this.snapshotGroup = Objects.requireNonNull(snapshotGroup);
            return this;
        }
        @CustomType.Setter
        public Builder timeoutInSeconds(Integer timeoutInSeconds) {
            this.timeoutInSeconds = Objects.requireNonNull(timeoutInSeconds);
            return this;
        }
        public GetServerBackupPlansPlanDetail build() {
            final var o = new GetServerBackupPlansPlanDetail();
            o.appConsistent = appConsistent;
            o.destinationRegionId = destinationRegionId;
            o.destinationRetention = destinationRetention;
            o.diskIdLists = diskIdLists;
            o.doCopy = doCopy;
            o.enableFsFreeze = enableFsFreeze;
            o.postScriptPath = postScriptPath;
            o.preScriptPath = preScriptPath;
            o.snapshotGroup = snapshotGroup;
            o.timeoutInSeconds = timeoutInSeconds;
            return o;
        }
    }
}
