// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRdsBackupsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRdsBackupsPlainArgs Empty = new GetRdsBackupsPlainArgs();

    /**
     * BackupMode.
     * 
     */
    @Import(name="backupMode")
    private @Nullable String backupMode;

    /**
     * @return BackupMode.
     * 
     */
    public Optional<String> backupMode() {
        return Optional.ofNullable(this.backupMode);
    }

    /**
     * Backup task status. **NOTE:** This parameter will only be returned when a task is executed. Value:
     * * **NoStart**: Not started
     * * **Checking**: check the backup
     * * **Preparing**: Prepare a backup
     * * **Waiting**: Waiting for backup
     * * **Uploading**: Upload backup
     * * **Finished**: Complete backup
     * * **Failed**: backup Failed
     * 
     */
    @Import(name="backupStatus")
    private @Nullable String backupStatus;

    /**
     * @return Backup task status. **NOTE:** This parameter will only be returned when a task is executed. Value:
     * * **NoStart**: Not started
     * * **Checking**: check the backup
     * * **Preparing**: Prepare a backup
     * * **Waiting**: Waiting for backup
     * * **Uploading**: Upload backup
     * * **Finished**: Complete backup
     * * **Failed**: backup Failed
     * 
     */
    public Optional<String> backupStatus() {
        return Optional.ofNullable(this.backupStatus);
    }

    /**
     * The db instance id.
     * 
     */
    @Import(name="dbInstanceId", required=true)
    private String dbInstanceId;

    /**
     * @return The db instance id.
     * 
     */
    public String dbInstanceId() {
        return this.dbInstanceId;
    }

    /**
     * The end time.
     * 
     */
    @Import(name="endTime")
    private @Nullable String endTime;

    /**
     * @return The end time.
     * 
     */
    public Optional<String> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * A list of Backup IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Backup IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
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
     * The start time.
     * 
     */
    @Import(name="startTime")
    private @Nullable String startTime;

    /**
     * @return The start time.
     * 
     */
    public Optional<String> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    private GetRdsBackupsPlainArgs() {}

    private GetRdsBackupsPlainArgs(GetRdsBackupsPlainArgs $) {
        this.backupMode = $.backupMode;
        this.backupStatus = $.backupStatus;
        this.dbInstanceId = $.dbInstanceId;
        this.endTime = $.endTime;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.startTime = $.startTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRdsBackupsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRdsBackupsPlainArgs $;

        public Builder() {
            $ = new GetRdsBackupsPlainArgs();
        }

        public Builder(GetRdsBackupsPlainArgs defaults) {
            $ = new GetRdsBackupsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupMode BackupMode.
         * 
         * @return builder
         * 
         */
        public Builder backupMode(@Nullable String backupMode) {
            $.backupMode = backupMode;
            return this;
        }

        /**
         * @param backupStatus Backup task status. **NOTE:** This parameter will only be returned when a task is executed. Value:
         * * **NoStart**: Not started
         * * **Checking**: check the backup
         * * **Preparing**: Prepare a backup
         * * **Waiting**: Waiting for backup
         * * **Uploading**: Upload backup
         * * **Finished**: Complete backup
         * * **Failed**: backup Failed
         * 
         * @return builder
         * 
         */
        public Builder backupStatus(@Nullable String backupStatus) {
            $.backupStatus = backupStatus;
            return this;
        }

        /**
         * @param dbInstanceId The db instance id.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(String dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        /**
         * @param endTime The end time.
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable String endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param ids A list of Backup IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Backup IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
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
         * @param startTime The start time.
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable String startTime) {
            $.startTime = startTime;
            return this;
        }

        public GetRdsBackupsPlainArgs build() {
            $.dbInstanceId = Objects.requireNonNull($.dbInstanceId, "expected parameter 'dbInstanceId' to be non-null");
            return $;
        }
    }

}
