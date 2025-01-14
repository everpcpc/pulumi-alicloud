// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram.inputs;

import com.pulumi.alicloud.ram.inputs.GetPolicyDocumentStatementArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPolicyDocumentArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPolicyDocumentArgs Empty = new GetPolicyDocumentArgs();

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Statement of the RAM policy document. See the following `Block statement`.
     * 
     */
    @Import(name="statements")
    private @Nullable Output<List<GetPolicyDocumentStatementArgs>> statements;

    /**
     * @return Statement of the RAM policy document. See the following `Block statement`.
     * 
     */
    public Optional<Output<List<GetPolicyDocumentStatementArgs>>> statements() {
        return Optional.ofNullable(this.statements);
    }

    /**
     * Version of the RAM policy document. Valid value is `1`. Default value is `1`.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the RAM policy document. Valid value is `1`. Default value is `1`.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private GetPolicyDocumentArgs() {}

    private GetPolicyDocumentArgs(GetPolicyDocumentArgs $) {
        this.outputFile = $.outputFile;
        this.statements = $.statements;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPolicyDocumentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPolicyDocumentArgs $;

        public Builder() {
            $ = new GetPolicyDocumentArgs();
        }

        public Builder(GetPolicyDocumentArgs defaults) {
            $ = new GetPolicyDocumentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param statements Statement of the RAM policy document. See the following `Block statement`.
         * 
         * @return builder
         * 
         */
        public Builder statements(@Nullable Output<List<GetPolicyDocumentStatementArgs>> statements) {
            $.statements = statements;
            return this;
        }

        /**
         * @param statements Statement of the RAM policy document. See the following `Block statement`.
         * 
         * @return builder
         * 
         */
        public Builder statements(List<GetPolicyDocumentStatementArgs> statements) {
            return statements(Output.of(statements));
        }

        /**
         * @param statements Statement of the RAM policy document. See the following `Block statement`.
         * 
         * @return builder
         * 
         */
        public Builder statements(GetPolicyDocumentStatementArgs... statements) {
            return statements(List.of(statements));
        }

        /**
         * @param version Version of the RAM policy document. Valid value is `1`. Default value is `1`.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the RAM policy document. Valid value is `1`. Default value is `1`.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public GetPolicyDocumentArgs build() {
            return $;
        }
    }

}
