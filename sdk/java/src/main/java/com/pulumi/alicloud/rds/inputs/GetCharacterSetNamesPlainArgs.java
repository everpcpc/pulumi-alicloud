// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCharacterSetNamesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCharacterSetNamesPlainArgs Empty = new GetCharacterSetNamesPlainArgs();

    /**
     * Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`.
     * 
     */
    @Import(name="engine", required=true)
    private String engine;

    /**
     * @return Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`.
     * 
     */
    public String engine() {
        return this.engine;
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetCharacterSetNamesPlainArgs() {}

    private GetCharacterSetNamesPlainArgs(GetCharacterSetNamesPlainArgs $) {
        this.engine = $.engine;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCharacterSetNamesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCharacterSetNamesPlainArgs $;

        public Builder() {
            $ = new GetCharacterSetNamesPlainArgs();
        }

        public Builder(GetCharacterSetNamesPlainArgs defaults) {
            $ = new GetCharacterSetNamesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param engine Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            $.engine = engine;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public GetCharacterSetNamesPlainArgs build() {
            $.engine = Objects.requireNonNull($.engine, "expected parameter 'engine' to be non-null");
            return $;
        }
    }

}
