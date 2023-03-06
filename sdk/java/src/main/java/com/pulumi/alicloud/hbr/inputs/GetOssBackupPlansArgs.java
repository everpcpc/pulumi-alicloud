// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOssBackupPlansArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOssBackupPlansArgs Empty = new GetOssBackupPlansArgs();

    /**
     * The name of OSS bucket.
     * 
     */
    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    /**
     * @return The name of OSS bucket.
     * 
     */
    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    /**
     * A list of OssBackupPlan IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of OssBackupPlan IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by OssBackupPlan name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by OssBackupPlan name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The ID of backup vault.
     * 
     */
    @Import(name="vaultId")
    private @Nullable Output<String> vaultId;

    /**
     * @return The ID of backup vault.
     * 
     */
    public Optional<Output<String>> vaultId() {
        return Optional.ofNullable(this.vaultId);
    }

    private GetOssBackupPlansArgs() {}

    private GetOssBackupPlansArgs(GetOssBackupPlansArgs $) {
        this.bucket = $.bucket;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.vaultId = $.vaultId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOssBackupPlansArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOssBackupPlansArgs $;

        public Builder() {
            $ = new GetOssBackupPlansArgs();
        }

        public Builder(GetOssBackupPlansArgs defaults) {
            $ = new GetOssBackupPlansArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of OSS bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of OSS bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param ids A list of OssBackupPlan IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of OssBackupPlan IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of OssBackupPlan IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by OssBackupPlan name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by OssBackupPlan name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param vaultId The ID of backup vault.
         * 
         * @return builder
         * 
         */
        public Builder vaultId(@Nullable Output<String> vaultId) {
            $.vaultId = vaultId;
            return this;
        }

        /**
         * @param vaultId The ID of backup vault.
         * 
         * @return builder
         * 
         */
        public Builder vaultId(String vaultId) {
            return vaultId(Output.of(vaultId));
        }

        public GetOssBackupPlansArgs build() {
            return $;
        }
    }

}