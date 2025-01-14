// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReplicationVaultArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReplicationVaultArgs Empty = new ReplicationVaultArgs();

    /**
     * The description of the backup vault. The description must be 0 to 255 characters in length.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the backup vault. The description must be 0 to 255 characters in length.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the region where the source vault resides.
     * 
     */
    @Import(name="replicationSourceRegionId", required=true)
    private Output<String> replicationSourceRegionId;

    /**
     * @return The ID of the region where the source vault resides.
     * 
     */
    public Output<String> replicationSourceRegionId() {
        return this.replicationSourceRegionId;
    }

    /**
     * The ID of the source vault.
     * 
     */
    @Import(name="replicationSourceVaultId", required=true)
    private Output<String> replicationSourceVaultId;

    /**
     * @return The ID of the source vault.
     * 
     */
    public Output<String> replicationSourceVaultId() {
        return this.replicationSourceVaultId;
    }

    /**
     * The name of the backup vault. The name must be 1 to 64 characters in length.
     * 
     */
    @Import(name="vaultName", required=true)
    private Output<String> vaultName;

    /**
     * @return The name of the backup vault. The name must be 1 to 64 characters in length.
     * 
     */
    public Output<String> vaultName() {
        return this.vaultName;
    }

    /**
     * The storage type of the backup vault. Valid values: `STANDARD`.
     * 
     */
    @Import(name="vaultStorageClass")
    private @Nullable Output<String> vaultStorageClass;

    /**
     * @return The storage type of the backup vault. Valid values: `STANDARD`.
     * 
     */
    public Optional<Output<String>> vaultStorageClass() {
        return Optional.ofNullable(this.vaultStorageClass);
    }

    private ReplicationVaultArgs() {}

    private ReplicationVaultArgs(ReplicationVaultArgs $) {
        this.description = $.description;
        this.replicationSourceRegionId = $.replicationSourceRegionId;
        this.replicationSourceVaultId = $.replicationSourceVaultId;
        this.vaultName = $.vaultName;
        this.vaultStorageClass = $.vaultStorageClass;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReplicationVaultArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReplicationVaultArgs $;

        public Builder() {
            $ = new ReplicationVaultArgs();
        }

        public Builder(ReplicationVaultArgs defaults) {
            $ = new ReplicationVaultArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the backup vault. The description must be 0 to 255 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the backup vault. The description must be 0 to 255 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param replicationSourceRegionId The ID of the region where the source vault resides.
         * 
         * @return builder
         * 
         */
        public Builder replicationSourceRegionId(Output<String> replicationSourceRegionId) {
            $.replicationSourceRegionId = replicationSourceRegionId;
            return this;
        }

        /**
         * @param replicationSourceRegionId The ID of the region where the source vault resides.
         * 
         * @return builder
         * 
         */
        public Builder replicationSourceRegionId(String replicationSourceRegionId) {
            return replicationSourceRegionId(Output.of(replicationSourceRegionId));
        }

        /**
         * @param replicationSourceVaultId The ID of the source vault.
         * 
         * @return builder
         * 
         */
        public Builder replicationSourceVaultId(Output<String> replicationSourceVaultId) {
            $.replicationSourceVaultId = replicationSourceVaultId;
            return this;
        }

        /**
         * @param replicationSourceVaultId The ID of the source vault.
         * 
         * @return builder
         * 
         */
        public Builder replicationSourceVaultId(String replicationSourceVaultId) {
            return replicationSourceVaultId(Output.of(replicationSourceVaultId));
        }

        /**
         * @param vaultName The name of the backup vault. The name must be 1 to 64 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder vaultName(Output<String> vaultName) {
            $.vaultName = vaultName;
            return this;
        }

        /**
         * @param vaultName The name of the backup vault. The name must be 1 to 64 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder vaultName(String vaultName) {
            return vaultName(Output.of(vaultName));
        }

        /**
         * @param vaultStorageClass The storage type of the backup vault. Valid values: `STANDARD`.
         * 
         * @return builder
         * 
         */
        public Builder vaultStorageClass(@Nullable Output<String> vaultStorageClass) {
            $.vaultStorageClass = vaultStorageClass;
            return this;
        }

        /**
         * @param vaultStorageClass The storage type of the backup vault. Valid values: `STANDARD`.
         * 
         * @return builder
         * 
         */
        public Builder vaultStorageClass(String vaultStorageClass) {
            return vaultStorageClass(Output.of(vaultStorageClass));
        }

        public ReplicationVaultArgs build() {
            $.replicationSourceRegionId = Objects.requireNonNull($.replicationSourceRegionId, "expected parameter 'replicationSourceRegionId' to be non-null");
            $.replicationSourceVaultId = Objects.requireNonNull($.replicationSourceVaultId, "expected parameter 'replicationSourceVaultId' to be non-null");
            $.vaultName = Objects.requireNonNull($.vaultName, "expected parameter 'vaultName' to be non-null");
            return $;
        }
    }

}
