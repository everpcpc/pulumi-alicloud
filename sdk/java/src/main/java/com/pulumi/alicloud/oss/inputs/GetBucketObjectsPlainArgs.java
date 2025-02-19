// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBucketObjectsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBucketObjectsPlainArgs Empty = new GetBucketObjectsPlainArgs();

    /**
     * Name of the bucket that contains the objects to find.
     * 
     */
    @Import(name="bucketName", required=true)
    private String bucketName;

    /**
     * @return Name of the bucket that contains the objects to find.
     * 
     */
    public String bucketName() {
        return this.bucketName;
    }

    /**
     * Filter results by the given key prefix (such as &#34;path/to/folder/logs-&#34;).
     * 
     */
    @Import(name="keyPrefix")
    private @Nullable String keyPrefix;

    /**
     * @return Filter results by the given key prefix (such as &#34;path/to/folder/logs-&#34;).
     * 
     */
    public Optional<String> keyPrefix() {
        return Optional.ofNullable(this.keyPrefix);
    }

    /**
     * A regex string to filter results by key.
     * 
     */
    @Import(name="keyRegex")
    private @Nullable String keyRegex;

    /**
     * @return A regex string to filter results by key.
     * 
     */
    public Optional<String> keyRegex() {
        return Optional.ofNullable(this.keyRegex);
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

    private GetBucketObjectsPlainArgs() {}

    private GetBucketObjectsPlainArgs(GetBucketObjectsPlainArgs $) {
        this.bucketName = $.bucketName;
        this.keyPrefix = $.keyPrefix;
        this.keyRegex = $.keyRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBucketObjectsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBucketObjectsPlainArgs $;

        public Builder() {
            $ = new GetBucketObjectsPlainArgs();
        }

        public Builder(GetBucketObjectsPlainArgs defaults) {
            $ = new GetBucketObjectsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucketName Name of the bucket that contains the objects to find.
         * 
         * @return builder
         * 
         */
        public Builder bucketName(String bucketName) {
            $.bucketName = bucketName;
            return this;
        }

        /**
         * @param keyPrefix Filter results by the given key prefix (such as &#34;path/to/folder/logs-&#34;).
         * 
         * @return builder
         * 
         */
        public Builder keyPrefix(@Nullable String keyPrefix) {
            $.keyPrefix = keyPrefix;
            return this;
        }

        /**
         * @param keyRegex A regex string to filter results by key.
         * 
         * @return builder
         * 
         */
        public Builder keyRegex(@Nullable String keyRegex) {
            $.keyRegex = keyRegex;
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

        public GetBucketObjectsPlainArgs build() {
            $.bucketName = Objects.requireNonNull($.bucketName, "expected parameter 'bucketName' to be non-null");
            return $;
        }
    }

}
