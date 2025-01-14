// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseLogicDatabaseArgs extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseLogicDatabaseArgs Empty = new EnterpriseLogicDatabaseArgs();

    /**
     * Logical Library alias.
     * 
     */
    @Import(name="alias", required=true)
    private Output<String> alias;

    /**
     * @return Logical Library alias.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }

    /**
     * Sub-Database ID
     * 
     */
    @Import(name="databaseIds", required=true)
    private Output<List<String>> databaseIds;

    /**
     * @return Sub-Database ID
     * 
     */
    public Output<List<String>> databaseIds() {
        return this.databaseIds;
    }

    /**
     * The ID of the logical Library.
     * 
     */
    @Import(name="logicDatabaseId")
    private @Nullable Output<String> logicDatabaseId;

    /**
     * @return The ID of the logical Library.
     * 
     */
    public Optional<Output<String>> logicDatabaseId() {
        return Optional.ofNullable(this.logicDatabaseId);
    }

    private EnterpriseLogicDatabaseArgs() {}

    private EnterpriseLogicDatabaseArgs(EnterpriseLogicDatabaseArgs $) {
        this.alias = $.alias;
        this.databaseIds = $.databaseIds;
        this.logicDatabaseId = $.logicDatabaseId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseLogicDatabaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseLogicDatabaseArgs $;

        public Builder() {
            $ = new EnterpriseLogicDatabaseArgs();
        }

        public Builder(EnterpriseLogicDatabaseArgs defaults) {
            $ = new EnterpriseLogicDatabaseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias Logical Library alias.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias Logical Library alias.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param databaseIds Sub-Database ID
         * 
         * @return builder
         * 
         */
        public Builder databaseIds(Output<List<String>> databaseIds) {
            $.databaseIds = databaseIds;
            return this;
        }

        /**
         * @param databaseIds Sub-Database ID
         * 
         * @return builder
         * 
         */
        public Builder databaseIds(List<String> databaseIds) {
            return databaseIds(Output.of(databaseIds));
        }

        /**
         * @param databaseIds Sub-Database ID
         * 
         * @return builder
         * 
         */
        public Builder databaseIds(String... databaseIds) {
            return databaseIds(List.of(databaseIds));
        }

        /**
         * @param logicDatabaseId The ID of the logical Library.
         * 
         * @return builder
         * 
         */
        public Builder logicDatabaseId(@Nullable Output<String> logicDatabaseId) {
            $.logicDatabaseId = logicDatabaseId;
            return this;
        }

        /**
         * @param logicDatabaseId The ID of the logical Library.
         * 
         * @return builder
         * 
         */
        public Builder logicDatabaseId(String logicDatabaseId) {
            return logicDatabaseId(Output.of(logicDatabaseId));
        }

        public EnterpriseLogicDatabaseArgs build() {
            $.alias = Objects.requireNonNull($.alias, "expected parameter 'alias' to be non-null");
            $.databaseIds = Objects.requireNonNull($.databaseIds, "expected parameter 'databaseIds' to be non-null");
            return $;
        }
    }

}
