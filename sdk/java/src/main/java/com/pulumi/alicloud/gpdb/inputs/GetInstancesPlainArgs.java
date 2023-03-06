// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.gpdb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstancesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstancesPlainArgs Empty = new GetInstancesPlainArgs();

    /**
     * Instance availability zone.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable String availabilityZone;

    /**
     * @return Instance availability zone.
     * 
     */
    public Optional<String> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * The db instance categories.
     * 
     */
    @Import(name="dbInstanceCategories")
    private @Nullable String dbInstanceCategories;

    /**
     * @return The db instance categories.
     * 
     */
    public Optional<String> dbInstanceCategories() {
        return Optional.ofNullable(this.dbInstanceCategories);
    }

    /**
     * The db instance modes.
     * 
     */
    @Import(name="dbInstanceModes")
    private @Nullable String dbInstanceModes;

    /**
     * @return The db instance modes.
     * 
     */
    public Optional<String> dbInstanceModes() {
        return Optional.ofNullable(this.dbInstanceModes);
    }

    /**
     * The description of the instance.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return The description of the instance.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * The ids list of AnalyticDB for PostgreSQL instances.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return The ids list of AnalyticDB for PostgreSQL instances.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The network type of the instance.
     * 
     */
    @Import(name="instanceNetworkType")
    private @Nullable String instanceNetworkType;

    /**
     * @return The network type of the instance.
     * 
     */
    public Optional<String> instanceNetworkType() {
        return Optional.ofNullable(this.instanceNetworkType);
    }

    /**
     * A regex string to apply to the instance name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to apply to the instance name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The ID of the enterprise resource group to which the instance belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The ID of the enterprise resource group to which the instance belongs.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags of the instance.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return The tags of the instance.
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The vswitch id.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable String vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetInstancesPlainArgs() {}

    private GetInstancesPlainArgs(GetInstancesPlainArgs $) {
        this.availabilityZone = $.availabilityZone;
        this.dbInstanceCategories = $.dbInstanceCategories;
        this.dbInstanceModes = $.dbInstanceModes;
        this.description = $.description;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.instanceNetworkType = $.instanceNetworkType;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstancesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstancesPlainArgs $;

        public Builder() {
            $ = new GetInstancesPlainArgs();
        }

        public Builder(GetInstancesPlainArgs defaults) {
            $ = new GetInstancesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityZone Instance availability zone.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable String availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param dbInstanceCategories The db instance categories.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceCategories(@Nullable String dbInstanceCategories) {
            $.dbInstanceCategories = dbInstanceCategories;
            return this;
        }

        /**
         * @param dbInstanceModes The db instance modes.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceModes(@Nullable String dbInstanceModes) {
            $.dbInstanceModes = dbInstanceModes;
            return this;
        }

        /**
         * @param description The description of the instance.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids The ids list of AnalyticDB for PostgreSQL instances.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids The ids list of AnalyticDB for PostgreSQL instances.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceNetworkType The network type of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceNetworkType(@Nullable String instanceNetworkType) {
            $.instanceNetworkType = instanceNetworkType;
            return this;
        }

        /**
         * @param nameRegex A regex string to apply to the instance name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the enterprise resource group to which the instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param status The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags The tags of the instance.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param vswitchId The vswitch id.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable String vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        public GetInstancesPlainArgs build() {
            return $;
        }
    }

}