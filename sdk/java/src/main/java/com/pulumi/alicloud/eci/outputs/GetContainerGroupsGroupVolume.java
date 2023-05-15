// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.alicloud.eci.outputs.GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPath;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetContainerGroupsGroupVolume {
    /**
     * @return The list of configuration file paths.
     * 
     */
    private List<GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPath> configFileVolumeConfigFileToPaths;
    /**
     * @return The ID of DiskVolume.
     * 
     */
    private String diskVolumeDiskId;
    /**
     * @return The type of DiskVolume.
     * 
     */
    private String diskVolumeFsType;
    /**
     * @return The name of the FlexVolume driver.
     * 
     */
    private String flexVolumeDriver;
    /**
     * @return The type of the mounted file system. The default value is determined by the script of FlexVolume.
     * 
     */
    private String flexVolumeFsType;
    /**
     * @return The list of FlexVolume objects.
     * 
     */
    private String flexVolumeOptions;
    /**
     * @return The name of the volume.
     * 
     */
    private String name;
    /**
     * @return The path to the NFS volume.
     * 
     */
    private String nfsVolumePath;
    /**
     * @return Default value: `false`.
     * 
     */
    private Boolean nfsVolumeReadOnly;
    /**
     * @return The address of the NFS server.
     * 
     */
    private String nfsVolumeServer;
    /**
     * @return The type of the volume. Currently, the following types of volumes are supported: EmptyDirVolume, NFSVolume, ConfigFileVolume, and FlexVolume.
     * 
     */
    private String type;

    private GetContainerGroupsGroupVolume() {}
    /**
     * @return The list of configuration file paths.
     * 
     */
    public List<GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPath> configFileVolumeConfigFileToPaths() {
        return this.configFileVolumeConfigFileToPaths;
    }
    /**
     * @return The ID of DiskVolume.
     * 
     */
    public String diskVolumeDiskId() {
        return this.diskVolumeDiskId;
    }
    /**
     * @return The type of DiskVolume.
     * 
     */
    public String diskVolumeFsType() {
        return this.diskVolumeFsType;
    }
    /**
     * @return The name of the FlexVolume driver.
     * 
     */
    public String flexVolumeDriver() {
        return this.flexVolumeDriver;
    }
    /**
     * @return The type of the mounted file system. The default value is determined by the script of FlexVolume.
     * 
     */
    public String flexVolumeFsType() {
        return this.flexVolumeFsType;
    }
    /**
     * @return The list of FlexVolume objects.
     * 
     */
    public String flexVolumeOptions() {
        return this.flexVolumeOptions;
    }
    /**
     * @return The name of the volume.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The path to the NFS volume.
     * 
     */
    public String nfsVolumePath() {
        return this.nfsVolumePath;
    }
    /**
     * @return Default value: `false`.
     * 
     */
    public Boolean nfsVolumeReadOnly() {
        return this.nfsVolumeReadOnly;
    }
    /**
     * @return The address of the NFS server.
     * 
     */
    public String nfsVolumeServer() {
        return this.nfsVolumeServer;
    }
    /**
     * @return The type of the volume. Currently, the following types of volumes are supported: EmptyDirVolume, NFSVolume, ConfigFileVolume, and FlexVolume.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerGroupsGroupVolume defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPath> configFileVolumeConfigFileToPaths;
        private String diskVolumeDiskId;
        private String diskVolumeFsType;
        private String flexVolumeDriver;
        private String flexVolumeFsType;
        private String flexVolumeOptions;
        private String name;
        private String nfsVolumePath;
        private Boolean nfsVolumeReadOnly;
        private String nfsVolumeServer;
        private String type;
        public Builder() {}
        public Builder(GetContainerGroupsGroupVolume defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configFileVolumeConfigFileToPaths = defaults.configFileVolumeConfigFileToPaths;
    	      this.diskVolumeDiskId = defaults.diskVolumeDiskId;
    	      this.diskVolumeFsType = defaults.diskVolumeFsType;
    	      this.flexVolumeDriver = defaults.flexVolumeDriver;
    	      this.flexVolumeFsType = defaults.flexVolumeFsType;
    	      this.flexVolumeOptions = defaults.flexVolumeOptions;
    	      this.name = defaults.name;
    	      this.nfsVolumePath = defaults.nfsVolumePath;
    	      this.nfsVolumeReadOnly = defaults.nfsVolumeReadOnly;
    	      this.nfsVolumeServer = defaults.nfsVolumeServer;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder configFileVolumeConfigFileToPaths(List<GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPath> configFileVolumeConfigFileToPaths) {
            this.configFileVolumeConfigFileToPaths = Objects.requireNonNull(configFileVolumeConfigFileToPaths);
            return this;
        }
        public Builder configFileVolumeConfigFileToPaths(GetContainerGroupsGroupVolumeConfigFileVolumeConfigFileToPath... configFileVolumeConfigFileToPaths) {
            return configFileVolumeConfigFileToPaths(List.of(configFileVolumeConfigFileToPaths));
        }
        @CustomType.Setter
        public Builder diskVolumeDiskId(String diskVolumeDiskId) {
            this.diskVolumeDiskId = Objects.requireNonNull(diskVolumeDiskId);
            return this;
        }
        @CustomType.Setter
        public Builder diskVolumeFsType(String diskVolumeFsType) {
            this.diskVolumeFsType = Objects.requireNonNull(diskVolumeFsType);
            return this;
        }
        @CustomType.Setter
        public Builder flexVolumeDriver(String flexVolumeDriver) {
            this.flexVolumeDriver = Objects.requireNonNull(flexVolumeDriver);
            return this;
        }
        @CustomType.Setter
        public Builder flexVolumeFsType(String flexVolumeFsType) {
            this.flexVolumeFsType = Objects.requireNonNull(flexVolumeFsType);
            return this;
        }
        @CustomType.Setter
        public Builder flexVolumeOptions(String flexVolumeOptions) {
            this.flexVolumeOptions = Objects.requireNonNull(flexVolumeOptions);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nfsVolumePath(String nfsVolumePath) {
            this.nfsVolumePath = Objects.requireNonNull(nfsVolumePath);
            return this;
        }
        @CustomType.Setter
        public Builder nfsVolumeReadOnly(Boolean nfsVolumeReadOnly) {
            this.nfsVolumeReadOnly = Objects.requireNonNull(nfsVolumeReadOnly);
            return this;
        }
        @CustomType.Setter
        public Builder nfsVolumeServer(String nfsVolumeServer) {
            this.nfsVolumeServer = Objects.requireNonNull(nfsVolumeServer);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetContainerGroupsGroupVolume build() {
            final var o = new GetContainerGroupsGroupVolume();
            o.configFileVolumeConfigFileToPaths = configFileVolumeConfigFileToPaths;
            o.diskVolumeDiskId = diskVolumeDiskId;
            o.diskVolumeFsType = diskVolumeFsType;
            o.flexVolumeDriver = flexVolumeDriver;
            o.flexVolumeFsType = flexVolumeFsType;
            o.flexVolumeOptions = flexVolumeOptions;
            o.name = name;
            o.nfsVolumePath = nfsVolumePath;
            o.nfsVolumeReadOnly = nfsVolumeReadOnly;
            o.nfsVolumeServer = nfsVolumeServer;
            o.type = type;
            return o;
        }
    }
}
