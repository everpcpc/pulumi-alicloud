// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecondaryIndexState extends com.pulumi.resources.ResourceArgs {

    public static final SecondaryIndexState Empty = new SecondaryIndexState();

    /**
     * A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    @Import(name="definedColumns")
    private @Nullable Output<List<String>> definedColumns;

    /**
     * @return A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    public Optional<Output<List<String>>> definedColumns() {
        return Optional.ofNullable(this.definedColumns);
    }

    /**
     * whether the index contains data that already exists in the data table. When include_base_data is set to true, it means that stock data is included.
     * 
     */
    @Import(name="includeBaseData")
    private @Nullable Output<Boolean> includeBaseData;

    /**
     * @return whether the index contains data that already exists in the data table. When include_base_data is set to true, it means that stock data is included.
     * 
     */
    public Optional<Output<Boolean>> includeBaseData() {
        return Optional.ofNullable(this.includeBaseData);
    }

    /**
     * The index name of the OTS Table. If changed, a new index would be created.
     * 
     */
    @Import(name="indexName")
    private @Nullable Output<String> indexName;

    /**
     * @return The index name of the OTS Table. If changed, a new index would be created.
     * 
     */
    public Optional<Output<String>> indexName() {
        return Optional.ofNullable(this.indexName);
    }

    /**
     * The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
     * 
     */
    @Import(name="indexType")
    private @Nullable Output<String> indexType;

    /**
     * @return The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
     * 
     */
    public Optional<Output<String>> indexType() {
        return Optional.ofNullable(this.indexType);
    }

    /**
     * The name of the OTS instance in which table will located.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return The name of the OTS instance in which table will located.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    @Import(name="primaryKeys")
    private @Nullable Output<List<String>> primaryKeys;

    /**
     * @return A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    public Optional<Output<List<String>>> primaryKeys() {
        return Optional.ofNullable(this.primaryKeys);
    }

    /**
     * The name of the OTS table. If changed, a new table would be created.
     * 
     */
    @Import(name="tableName")
    private @Nullable Output<String> tableName;

    /**
     * @return The name of the OTS table. If changed, a new table would be created.
     * 
     */
    public Optional<Output<String>> tableName() {
        return Optional.ofNullable(this.tableName);
    }

    private SecondaryIndexState() {}

    private SecondaryIndexState(SecondaryIndexState $) {
        this.definedColumns = $.definedColumns;
        this.includeBaseData = $.includeBaseData;
        this.indexName = $.indexName;
        this.indexType = $.indexType;
        this.instanceName = $.instanceName;
        this.primaryKeys = $.primaryKeys;
        this.tableName = $.tableName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecondaryIndexState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecondaryIndexState $;

        public Builder() {
            $ = new SecondaryIndexState();
        }

        public Builder(SecondaryIndexState defaults) {
            $ = new SecondaryIndexState(Objects.requireNonNull(defaults));
        }

        /**
         * @param definedColumns A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
         * 
         * @return builder
         * 
         */
        public Builder definedColumns(@Nullable Output<List<String>> definedColumns) {
            $.definedColumns = definedColumns;
            return this;
        }

        /**
         * @param definedColumns A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
         * 
         * @return builder
         * 
         */
        public Builder definedColumns(List<String> definedColumns) {
            return definedColumns(Output.of(definedColumns));
        }

        /**
         * @param definedColumns A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
         * 
         * @return builder
         * 
         */
        public Builder definedColumns(String... definedColumns) {
            return definedColumns(List.of(definedColumns));
        }

        /**
         * @param includeBaseData whether the index contains data that already exists in the data table. When include_base_data is set to true, it means that stock data is included.
         * 
         * @return builder
         * 
         */
        public Builder includeBaseData(@Nullable Output<Boolean> includeBaseData) {
            $.includeBaseData = includeBaseData;
            return this;
        }

        /**
         * @param includeBaseData whether the index contains data that already exists in the data table. When include_base_data is set to true, it means that stock data is included.
         * 
         * @return builder
         * 
         */
        public Builder includeBaseData(Boolean includeBaseData) {
            return includeBaseData(Output.of(includeBaseData));
        }

        /**
         * @param indexName The index name of the OTS Table. If changed, a new index would be created.
         * 
         * @return builder
         * 
         */
        public Builder indexName(@Nullable Output<String> indexName) {
            $.indexName = indexName;
            return this;
        }

        /**
         * @param indexName The index name of the OTS Table. If changed, a new index would be created.
         * 
         * @return builder
         * 
         */
        public Builder indexName(String indexName) {
            return indexName(Output.of(indexName));
        }

        /**
         * @param indexType The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
         * 
         * @return builder
         * 
         */
        public Builder indexType(@Nullable Output<String> indexType) {
            $.indexType = indexType;
            return this;
        }

        /**
         * @param indexType The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
         * 
         * @return builder
         * 
         */
        public Builder indexType(String indexType) {
            return indexType(Output.of(indexType));
        }

        /**
         * @param instanceName The name of the OTS instance in which table will located.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The name of the OTS instance in which table will located.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param primaryKeys A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeys(@Nullable Output<List<String>> primaryKeys) {
            $.primaryKeys = primaryKeys;
            return this;
        }

        /**
         * @param primaryKeys A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeys(List<String> primaryKeys) {
            return primaryKeys(Output.of(primaryKeys));
        }

        /**
         * @param primaryKeys A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeys(String... primaryKeys) {
            return primaryKeys(List.of(primaryKeys));
        }

        /**
         * @param tableName The name of the OTS table. If changed, a new table would be created.
         * 
         * @return builder
         * 
         */
        public Builder tableName(@Nullable Output<String> tableName) {
            $.tableName = tableName;
            return this;
        }

        /**
         * @param tableName The name of the OTS table. If changed, a new table would be created.
         * 
         * @return builder
         * 
         */
        public Builder tableName(String tableName) {
            return tableName(Output.of(tableName));
        }

        public SecondaryIndexState build() {
            return $;
        }
    }

}