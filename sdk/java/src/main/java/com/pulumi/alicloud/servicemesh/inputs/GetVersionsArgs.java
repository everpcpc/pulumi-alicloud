// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVersionsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVersionsArgs Empty = new GetVersionsArgs();

    /**
     * The edition of the ASM instance.
     * 
     */
    @Import(name="edition")
    private @Nullable Output<String> edition;

    /**
     * @return The edition of the ASM instance.
     * 
     */
    public Optional<Output<String>> edition() {
        return Optional.ofNullable(this.edition);
    }

    /**
     * A list of ASM versions. Its element formats as `&lt;edition&gt;:&lt;version&gt;`.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of ASM versions. Its element formats as `&lt;edition&gt;:&lt;version&gt;`.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

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

    private GetVersionsArgs() {}

    private GetVersionsArgs(GetVersionsArgs $) {
        this.edition = $.edition;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVersionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVersionsArgs $;

        public Builder() {
            $ = new GetVersionsArgs();
        }

        public Builder(GetVersionsArgs defaults) {
            $ = new GetVersionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param edition The edition of the ASM instance.
         * 
         * @return builder
         * 
         */
        public Builder edition(@Nullable Output<String> edition) {
            $.edition = edition;
            return this;
        }

        /**
         * @param edition The edition of the ASM instance.
         * 
         * @return builder
         * 
         */
        public Builder edition(String edition) {
            return edition(Output.of(edition));
        }

        /**
         * @param ids A list of ASM versions. Its element formats as `&lt;edition&gt;:&lt;version&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of ASM versions. Its element formats as `&lt;edition&gt;:&lt;version&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of ASM versions. Its element formats as `&lt;edition&gt;:&lt;version&gt;`.
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

        public GetVersionsArgs build() {
            return $;
        }
    }

}
