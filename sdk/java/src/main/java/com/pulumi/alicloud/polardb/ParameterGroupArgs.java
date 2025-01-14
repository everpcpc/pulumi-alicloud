// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.alicloud.polardb.inputs.ParameterGroupParameterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ParameterGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ParameterGroupArgs Empty = new ParameterGroupArgs();

    /**
     * The type of the database engine. Only `MySQL` is supported.
     * 
     */
    @Import(name="dbType", required=true)
    private Output<String> dbType;

    /**
     * @return The type of the database engine. Only `MySQL` is supported.
     * 
     */
    public Output<String> dbType() {
        return this.dbType;
    }

    /**
     * The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
     * 
     */
    @Import(name="dbVersion", required=true)
    private Output<String> dbVersion;

    /**
     * @return The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
     * 
     */
    public Output<String> dbVersion() {
        return this.dbVersion;
    }

    /**
     * The description of the parameter template. It must be 0 to 200 characters in length.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the parameter template. It must be 0 to 200 characters in length.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The parameter template. See the following `Block parameters`.
     * 
     */
    @Import(name="parameters", required=true)
    private Output<List<ParameterGroupParameterArgs>> parameters;

    /**
     * @return The parameter template. See the following `Block parameters`.
     * 
     */
    public Output<List<ParameterGroupParameterArgs>> parameters() {
        return this.parameters;
    }

    private ParameterGroupArgs() {}

    private ParameterGroupArgs(ParameterGroupArgs $) {
        this.dbType = $.dbType;
        this.dbVersion = $.dbVersion;
        this.description = $.description;
        this.name = $.name;
        this.parameters = $.parameters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ParameterGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ParameterGroupArgs $;

        public Builder() {
            $ = new ParameterGroupArgs();
        }

        public Builder(ParameterGroupArgs defaults) {
            $ = new ParameterGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbType The type of the database engine. Only `MySQL` is supported.
         * 
         * @return builder
         * 
         */
        public Builder dbType(Output<String> dbType) {
            $.dbType = dbType;
            return this;
        }

        /**
         * @param dbType The type of the database engine. Only `MySQL` is supported.
         * 
         * @return builder
         * 
         */
        public Builder dbType(String dbType) {
            return dbType(Output.of(dbType));
        }

        /**
         * @param dbVersion The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
         * 
         * @return builder
         * 
         */
        public Builder dbVersion(Output<String> dbVersion) {
            $.dbVersion = dbVersion;
            return this;
        }

        /**
         * @param dbVersion The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
         * 
         * @return builder
         * 
         */
        public Builder dbVersion(String dbVersion) {
            return dbVersion(Output.of(dbVersion));
        }

        /**
         * @param description The description of the parameter template. It must be 0 to 200 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the parameter template. It must be 0 to 200 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parameters The parameter template. See the following `Block parameters`.
         * 
         * @return builder
         * 
         */
        public Builder parameters(Output<List<ParameterGroupParameterArgs>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters The parameter template. See the following `Block parameters`.
         * 
         * @return builder
         * 
         */
        public Builder parameters(List<ParameterGroupParameterArgs> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param parameters The parameter template. See the following `Block parameters`.
         * 
         * @return builder
         * 
         */
        public Builder parameters(ParameterGroupParameterArgs... parameters) {
            return parameters(List.of(parameters));
        }

        public ParameterGroupArgs build() {
            $.dbType = Objects.requireNonNull($.dbType, "expected parameter 'dbType' to be non-null");
            $.dbVersion = Objects.requireNonNull($.dbVersion, "expected parameter 'dbVersion' to be non-null");
            $.parameters = Objects.requireNonNull($.parameters, "expected parameter 'parameters' to be non-null");
            return $;
        }
    }

}
