// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetWafPoliciesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetWafPoliciesPlainArgs Empty = new GetWafPoliciesPlainArgs();

    /**
     * A list of Waf Policy IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Waf Policy IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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
     * The query conditions. The value is a string in the JSON format. Format: `{&#34;PolicyIds&#34;:&#34;The ID of the proteuleIds&#34;:&#34;Thection policy&#34;,&#34;R range of protection rule IDs&#34;,&#34;PolicyNameLike&#34;:&#34;The name of the protection policy&#34;,&#34;DomainNames&#34;:&#34;The protected domain names&#34;,&#34;PolicyType&#34;:&#34;default&#34;,&#34;DefenseScenes&#34;:&#34;waf_group&#34;,&#34;PolicyStatus&#34;:&#34;on&#34;,&#34;OrderBy&#34;:&#34;GmtModified&#34;,&#34;Desc&#34;:&#34;false&#34;}`.
     * 
     */
    @Import(name="queryArgs")
    private @Nullable String queryArgs;

    /**
     * @return The query conditions. The value is a string in the JSON format. Format: `{&#34;PolicyIds&#34;:&#34;The ID of the proteuleIds&#34;:&#34;Thection policy&#34;,&#34;R range of protection rule IDs&#34;,&#34;PolicyNameLike&#34;:&#34;The name of the protection policy&#34;,&#34;DomainNames&#34;:&#34;The protected domain names&#34;,&#34;PolicyType&#34;:&#34;default&#34;,&#34;DefenseScenes&#34;:&#34;waf_group&#34;,&#34;PolicyStatus&#34;:&#34;on&#34;,&#34;OrderBy&#34;:&#34;GmtModified&#34;,&#34;Desc&#34;:&#34;false&#34;}`.
     * 
     */
    public Optional<String> queryArgs() {
        return Optional.ofNullable(this.queryArgs);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetWafPoliciesPlainArgs() {}

    private GetWafPoliciesPlainArgs(GetWafPoliciesPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.queryArgs = $.queryArgs;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetWafPoliciesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWafPoliciesPlainArgs $;

        public Builder() {
            $ = new GetWafPoliciesPlainArgs();
        }

        public Builder(GetWafPoliciesPlainArgs defaults) {
            $ = new GetWafPoliciesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Waf Policy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Waf Policy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
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

        /**
         * @param queryArgs The query conditions. The value is a string in the JSON format. Format: `{&#34;PolicyIds&#34;:&#34;The ID of the proteuleIds&#34;:&#34;Thection policy&#34;,&#34;R range of protection rule IDs&#34;,&#34;PolicyNameLike&#34;:&#34;The name of the protection policy&#34;,&#34;DomainNames&#34;:&#34;The protected domain names&#34;,&#34;PolicyType&#34;:&#34;default&#34;,&#34;DefenseScenes&#34;:&#34;waf_group&#34;,&#34;PolicyStatus&#34;:&#34;on&#34;,&#34;OrderBy&#34;:&#34;GmtModified&#34;,&#34;Desc&#34;:&#34;false&#34;}`.
         * 
         * @return builder
         * 
         */
        public Builder queryArgs(@Nullable String queryArgs) {
            $.queryArgs = queryArgs;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetWafPoliciesPlainArgs build() {
            return $;
        }
    }

}
