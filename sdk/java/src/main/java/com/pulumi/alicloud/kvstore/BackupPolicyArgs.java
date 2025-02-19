// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackupPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackupPolicyArgs Empty = new BackupPolicyArgs();

    /**
     * Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     * 
     */
    @Import(name="backupPeriods")
    private @Nullable Output<List<String>> backupPeriods;

    /**
     * @return Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     * 
     */
    public Optional<Output<List<String>>> backupPeriods() {
        return Optional.ofNullable(this.backupPeriods);
    }

    /**
     * Backup time, in the format of HH:mmZ- HH:mm Z
     * 
     */
    @Import(name="backupTime")
    private @Nullable Output<String> backupTime;

    /**
     * @return Backup time, in the format of HH:mmZ- HH:mm Z
     * 
     */
    public Optional<Output<String>> backupTime() {
        return Optional.ofNullable(this.backupTime);
    }

    /**
     * The id of ApsaraDB for Redis or Memcache intance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The id of ApsaraDB for Redis or Memcache intance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    private BackupPolicyArgs() {}

    private BackupPolicyArgs(BackupPolicyArgs $) {
        this.backupPeriods = $.backupPeriods;
        this.backupTime = $.backupTime;
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackupPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackupPolicyArgs $;

        public Builder() {
            $ = new BackupPolicyArgs();
        }

        public Builder(BackupPolicyArgs defaults) {
            $ = new BackupPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupPeriods Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
         * 
         * @return builder
         * 
         */
        public Builder backupPeriods(@Nullable Output<List<String>> backupPeriods) {
            $.backupPeriods = backupPeriods;
            return this;
        }

        /**
         * @param backupPeriods Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
         * 
         * @return builder
         * 
         */
        public Builder backupPeriods(List<String> backupPeriods) {
            return backupPeriods(Output.of(backupPeriods));
        }

        /**
         * @param backupPeriods Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
         * 
         * @return builder
         * 
         */
        public Builder backupPeriods(String... backupPeriods) {
            return backupPeriods(List.of(backupPeriods));
        }

        /**
         * @param backupTime Backup time, in the format of HH:mmZ- HH:mm Z
         * 
         * @return builder
         * 
         */
        public Builder backupTime(@Nullable Output<String> backupTime) {
            $.backupTime = backupTime;
            return this;
        }

        /**
         * @param backupTime Backup time, in the format of HH:mmZ- HH:mm Z
         * 
         * @return builder
         * 
         */
        public Builder backupTime(String backupTime) {
            return backupTime(Output.of(backupTime));
        }

        /**
         * @param instanceId The id of ApsaraDB for Redis or Memcache intance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The id of ApsaraDB for Redis or Memcache intance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        public BackupPolicyArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
