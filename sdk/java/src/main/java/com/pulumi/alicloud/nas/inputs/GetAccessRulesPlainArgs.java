// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAccessRulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAccessRulesPlainArgs Empty = new GetAccessRulesPlainArgs();

    /**
     * Filter results by a specific AccessGroupName.
     * 
     */
    @Import(name="accessGroupName", required=true)
    private String accessGroupName;

    /**
     * @return Filter results by a specific AccessGroupName.
     * 
     */
    public String accessGroupName() {
        return this.accessGroupName;
    }

    /**
     * A list of rule IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of rule IDs.
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
     * Filter results by a specific RWAccess.
     * 
     */
    @Import(name="rwAccess")
    private @Nullable String rwAccess;

    /**
     * @return Filter results by a specific RWAccess.
     * 
     */
    public Optional<String> rwAccess() {
        return Optional.ofNullable(this.rwAccess);
    }

    /**
     * Filter results by a specific SourceCidrIp.
     * 
     */
    @Import(name="sourceCidrIp")
    private @Nullable String sourceCidrIp;

    /**
     * @return Filter results by a specific SourceCidrIp.
     * 
     */
    public Optional<String> sourceCidrIp() {
        return Optional.ofNullable(this.sourceCidrIp);
    }

    /**
     * Filter results by a specific UserAccess.
     * 
     */
    @Import(name="userAccess")
    private @Nullable String userAccess;

    /**
     * @return Filter results by a specific UserAccess.
     * 
     */
    public Optional<String> userAccess() {
        return Optional.ofNullable(this.userAccess);
    }

    private GetAccessRulesPlainArgs() {}

    private GetAccessRulesPlainArgs(GetAccessRulesPlainArgs $) {
        this.accessGroupName = $.accessGroupName;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.rwAccess = $.rwAccess;
        this.sourceCidrIp = $.sourceCidrIp;
        this.userAccess = $.userAccess;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAccessRulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAccessRulesPlainArgs $;

        public Builder() {
            $ = new GetAccessRulesPlainArgs();
        }

        public Builder(GetAccessRulesPlainArgs defaults) {
            $ = new GetAccessRulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessGroupName Filter results by a specific AccessGroupName.
         * 
         * @return builder
         * 
         */
        public Builder accessGroupName(String accessGroupName) {
            $.accessGroupName = accessGroupName;
            return this;
        }

        /**
         * @param ids A list of rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of rule IDs.
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
         * @param rwAccess Filter results by a specific RWAccess.
         * 
         * @return builder
         * 
         */
        public Builder rwAccess(@Nullable String rwAccess) {
            $.rwAccess = rwAccess;
            return this;
        }

        /**
         * @param sourceCidrIp Filter results by a specific SourceCidrIp.
         * 
         * @return builder
         * 
         */
        public Builder sourceCidrIp(@Nullable String sourceCidrIp) {
            $.sourceCidrIp = sourceCidrIp;
            return this;
        }

        /**
         * @param userAccess Filter results by a specific UserAccess.
         * 
         * @return builder
         * 
         */
        public Builder userAccess(@Nullable String userAccess) {
            $.userAccess = userAccess;
            return this;
        }

        public GetAccessRulesPlainArgs build() {
            $.accessGroupName = Objects.requireNonNull($.accessGroupName, "expected parameter 'accessGroupName' to be non-null");
            return $;
        }
    }

}
