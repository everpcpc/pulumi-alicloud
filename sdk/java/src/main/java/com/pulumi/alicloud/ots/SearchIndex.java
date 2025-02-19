// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ots.SearchIndexArgs;
import com.pulumi.alicloud.ots.inputs.SearchIndexState;
import com.pulumi.alicloud.ots.outputs.SearchIndexSchema;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an OTS search index resource.
 * 
 * For information about OTS search index and how to use it, see [Search index overview](https://www.alibabacloud.com/help/en/tablestore/latest/search-index-overview).
 * 
 * &gt; **NOTE:** Available in v1.187.0+.
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * OTS search index can be imported using id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ots/searchIndex:SearchIndex index1 &lt;instance_name&gt;:&lt;table_name&gt;:&lt;index_name&gt;:&lt;index_type&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ots/searchIndex:SearchIndex")
public class SearchIndex extends com.pulumi.resources.CustomResource {
    /**
     * The search index create time.
     * 
     */
    @Export(name="createTime", type=Integer.class, parameters={})
    private Output<Integer> createTime;

    /**
     * @return The search index create time.
     * 
     */
    public Output<Integer> createTime() {
        return this.createTime;
    }
    /**
     * The timestamp for sync phase.
     * 
     */
    @Export(name="currentSyncTimestamp", type=Integer.class, parameters={})
    private Output<Integer> currentSyncTimestamp;

    /**
     * @return The timestamp for sync phase.
     * 
     */
    public Output<Integer> currentSyncTimestamp() {
        return this.currentSyncTimestamp;
    }
    /**
     * The index id of the search index which could not be changed.
     * 
     */
    @Export(name="indexId", type=String.class, parameters={})
    private Output<String> indexId;

    /**
     * @return The index id of the search index which could not be changed.
     * 
     */
    public Output<String> indexId() {
        return this.indexId;
    }
    /**
     * The index name of the OTS Table. If changed, a new index would be created.
     * 
     */
    @Export(name="indexName", type=String.class, parameters={})
    private Output<String> indexName;

    /**
     * @return The index name of the OTS Table. If changed, a new index would be created.
     * 
     */
    public Output<String> indexName() {
        return this.indexName;
    }
    /**
     * The name of the OTS instance in which table will located.
     * 
     */
    @Export(name="instanceName", type=String.class, parameters={})
    private Output<String> instanceName;

    /**
     * @return The name of the OTS instance in which table will located.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * The schema of the search index. If changed, a new index would be created.
     * 
     */
    @Export(name="schemas", type=List.class, parameters={SearchIndexSchema.class})
    private Output<List<SearchIndexSchema>> schemas;

    /**
     * @return The schema of the search index. If changed, a new index would be created.
     * 
     */
    public Output<List<SearchIndexSchema>> schemas() {
        return this.schemas;
    }
    /**
     * The search index sync phase. possible values: `Full`, `Incr`.
     * 
     */
    @Export(name="syncPhase", type=String.class, parameters={})
    private Output<String> syncPhase;

    /**
     * @return The search index sync phase. possible values: `Full`, `Incr`.
     * 
     */
    public Output<String> syncPhase() {
        return this.syncPhase;
    }
    /**
     * The name of the OTS table. If changed, a new table would be created.
     * 
     */
    @Export(name="tableName", type=String.class, parameters={})
    private Output<String> tableName;

    /**
     * @return The name of the OTS table. If changed, a new table would be created.
     * 
     */
    public Output<String> tableName() {
        return this.tableName;
    }
    /**
     * The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
     * If the retention period exceeds the TTL value, OTS automatically deletes expired data.
     * 
     */
    @Export(name="timeToLive", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> timeToLive;

    /**
     * @return The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
     * If the retention period exceeds the TTL value, OTS automatically deletes expired data.
     * 
     */
    public Output<Optional<Integer>> timeToLive() {
        return Codegen.optional(this.timeToLive);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SearchIndex(String name) {
        this(name, SearchIndexArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SearchIndex(String name, SearchIndexArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SearchIndex(String name, SearchIndexArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ots/searchIndex:SearchIndex", name, args == null ? SearchIndexArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SearchIndex(String name, Output<String> id, @Nullable SearchIndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ots/searchIndex:SearchIndex", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static SearchIndex get(String name, Output<String> id, @Nullable SearchIndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SearchIndex(name, id, state, options);
    }
}
