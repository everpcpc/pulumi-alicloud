// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RdsBackupArgs extends com.pulumi.resources.ResourceArgs {

    public static final RdsBackupArgs Empty = new RdsBackupArgs();

    /**
     * The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
     * 
     */
    @Import(name="backupMethod")
    private @Nullable Output<String> backupMethod;

    /**
     * @return The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
     * 
     */
    public Optional<Output<String>> backupMethod() {
        return Optional.ofNullable(this.backupMethod);
    }

    /**
     * The policy that you want to use for the backup task. Valid values:
     * * **db**: specifies to perform a database-level backup.
     * * **instance**: specifies to perform an instance-level backup.
     * 
     */
    @Import(name="backupStrategy")
    private @Nullable Output<String> backupStrategy;

    /**
     * @return The policy that you want to use for the backup task. Valid values:
     * * **db**: specifies to perform a database-level backup.
     * * **instance**: specifies to perform an instance-level backup.
     * 
     */
    public Optional<Output<String>> backupStrategy() {
        return Optional.ofNullable(this.backupStrategy);
    }

    /**
     * The method that you want to use for the backup task. Default value: `Auto`. Valid values:
     * * **Auto**: specifies to automatically perform a full or incremental backup.
     * * **FullBackup**: specifies to perform a full backup.
     * 
     */
    @Import(name="backupType")
    private @Nullable Output<String> backupType;

    /**
     * @return The method that you want to use for the backup task. Default value: `Auto`. Valid values:
     * * **Auto**: specifies to automatically perform a full or incremental backup.
     * * **FullBackup**: specifies to perform a full backup.
     * 
     */
    public Optional<Output<String>> backupType() {
        return Optional.ofNullable(this.backupType);
    }

    /**
     * The db instance id.
     * 
     */
    @Import(name="dbInstanceId", required=true)
    private Output<String> dbInstanceId;

    /**
     * @return The db instance id.
     * 
     */
    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }

    /**
     * The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
     * 
     */
    @Import(name="dbName")
    private @Nullable Output<String> dbName;

    /**
     * @return The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
     * 
     */
    public Optional<Output<String>> dbName() {
        return Optional.ofNullable(this.dbName);
    }

    /**
     * Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
     * 
     */
    @Import(name="removeFromState")
    private @Nullable Output<Boolean> removeFromState;

    /**
     * @return Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
     * 
     */
    public Optional<Output<Boolean>> removeFromState() {
        return Optional.ofNullable(this.removeFromState);
    }

    private RdsBackupArgs() {}

    private RdsBackupArgs(RdsBackupArgs $) {
        this.backupMethod = $.backupMethod;
        this.backupStrategy = $.backupStrategy;
        this.backupType = $.backupType;
        this.dbInstanceId = $.dbInstanceId;
        this.dbName = $.dbName;
        this.removeFromState = $.removeFromState;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RdsBackupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RdsBackupArgs $;

        public Builder() {
            $ = new RdsBackupArgs();
        }

        public Builder(RdsBackupArgs defaults) {
            $ = new RdsBackupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupMethod The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
         * 
         * @return builder
         * 
         */
        public Builder backupMethod(@Nullable Output<String> backupMethod) {
            $.backupMethod = backupMethod;
            return this;
        }

        /**
         * @param backupMethod The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
         * 
         * @return builder
         * 
         */
        public Builder backupMethod(String backupMethod) {
            return backupMethod(Output.of(backupMethod));
        }

        /**
         * @param backupStrategy The policy that you want to use for the backup task. Valid values:
         * * **db**: specifies to perform a database-level backup.
         * * **instance**: specifies to perform an instance-level backup.
         * 
         * @return builder
         * 
         */
        public Builder backupStrategy(@Nullable Output<String> backupStrategy) {
            $.backupStrategy = backupStrategy;
            return this;
        }

        /**
         * @param backupStrategy The policy that you want to use for the backup task. Valid values:
         * * **db**: specifies to perform a database-level backup.
         * * **instance**: specifies to perform an instance-level backup.
         * 
         * @return builder
         * 
         */
        public Builder backupStrategy(String backupStrategy) {
            return backupStrategy(Output.of(backupStrategy));
        }

        /**
         * @param backupType The method that you want to use for the backup task. Default value: `Auto`. Valid values:
         * * **Auto**: specifies to automatically perform a full or incremental backup.
         * * **FullBackup**: specifies to perform a full backup.
         * 
         * @return builder
         * 
         */
        public Builder backupType(@Nullable Output<String> backupType) {
            $.backupType = backupType;
            return this;
        }

        /**
         * @param backupType The method that you want to use for the backup task. Default value: `Auto`. Valid values:
         * * **Auto**: specifies to automatically perform a full or incremental backup.
         * * **FullBackup**: specifies to perform a full backup.
         * 
         * @return builder
         * 
         */
        public Builder backupType(String backupType) {
            return backupType(Output.of(backupType));
        }

        /**
         * @param dbInstanceId The db instance id.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(Output<String> dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        /**
         * @param dbInstanceId The db instance id.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(String dbInstanceId) {
            return dbInstanceId(Output.of(dbInstanceId));
        }

        /**
         * @param dbName The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dbName(@Nullable Output<String> dbName) {
            $.dbName = dbName;
            return this;
        }

        /**
         * @param dbName The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dbName(String dbName) {
            return dbName(Output.of(dbName));
        }

        /**
         * @param removeFromState Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
         * 
         * @return builder
         * 
         */
        public Builder removeFromState(@Nullable Output<Boolean> removeFromState) {
            $.removeFromState = removeFromState;
            return this;
        }

        /**
         * @param removeFromState Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
         * 
         * @return builder
         * 
         */
        public Builder removeFromState(Boolean removeFromState) {
            return removeFromState(Output.of(removeFromState));
        }

        public RdsBackupArgs build() {
            $.dbInstanceId = Objects.requireNonNull($.dbInstanceId, "expected parameter 'dbInstanceId' to be non-null");
            return $;
        }
    }

}
