// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dts;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dts.MigrationInstanceArgs;
import com.pulumi.alicloud.dts.inputs.MigrationInstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DTS Migration Instance resource.
 * 
 * For information about DTS Migration Instance and how to use it, see [What is Synchronization Instance](https://www.alibabacloud.com/help/en/doc-detail/208270.html).
 * 
 * &gt; **NOTE:** Available in v1.157.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.dts.MigrationInstance;
 * import com.pulumi.alicloud.dts.MigrationInstanceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new MigrationInstance(&#34;default&#34;, MigrationInstanceArgs.builder()        
 *             .destinationEndpointEngineName(&#34;MySQL&#34;)
 *             .destinationEndpointRegion(&#34;cn-hangzhou&#34;)
 *             .instanceClass(&#34;small&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .sourceEndpointEngineName(&#34;MySQL&#34;)
 *             .sourceEndpointRegion(&#34;cn-hangzhou&#34;)
 *             .syncArchitecture(&#34;oneway&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DTS Migration Instance can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dts/migrationInstance:MigrationInstance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dts/migrationInstance:MigrationInstance")
public class MigrationInstance extends com.pulumi.resources.CustomResource {
    /**
     * [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
     * 
     */
    @Export(name="computeUnit", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> computeUnit;

    /**
     * @return [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
     * 
     */
    public Output<Optional<Integer>> computeUnit() {
        return Codegen.optional(this.computeUnit);
    }
    /**
     * The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `source_endpoint_engine_name` equals `drds`.
     * 
     */
    @Export(name="databaseCount", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> databaseCount;

    /**
     * @return The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `source_endpoint_engine_name` equals `drds`.
     * 
     */
    public Output<Optional<Integer>> databaseCount() {
        return Codegen.optional(this.databaseCount);
    }
    /**
     * The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     * 
     */
    @Export(name="destinationEndpointEngineName", type=String.class, parameters={})
    private Output<String> destinationEndpointEngineName;

    /**
     * @return The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     * 
     */
    public Output<String> destinationEndpointEngineName() {
        return this.destinationEndpointEngineName;
    }
    /**
     * The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     * 
     */
    @Export(name="destinationEndpointRegion", type=String.class, parameters={})
    private Output<String> destinationEndpointRegion;

    /**
     * @return The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     * 
     */
    public Output<String> destinationEndpointRegion() {
        return this.destinationEndpointRegion;
    }
    /**
     * The ID of the Migration Instance.
     * 
     */
    @Export(name="dtsInstanceId", type=String.class, parameters={})
    private Output<String> dtsInstanceId;

    /**
     * @return The ID of the Migration Instance.
     * 
     */
    public Output<String> dtsInstanceId() {
        return this.dtsInstanceId;
    }
    /**
     * The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     * 
     */
    @Export(name="instanceClass", type=String.class, parameters={})
    private Output<String> instanceClass;

    /**
     * @return The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     * 
     */
    public Output<String> instanceClass() {
        return this.instanceClass;
    }
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `PayAsYouGo`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     * 
     */
    @Export(name="sourceEndpointEngineName", type=String.class, parameters={})
    private Output<String> sourceEndpointEngineName;

    /**
     * @return The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
     * 
     */
    public Output<String> sourceEndpointEngineName() {
        return this.sourceEndpointEngineName;
    }
    /**
     * The region of source instance.
     * 
     */
    @Export(name="sourceEndpointRegion", type=String.class, parameters={})
    private Output<String> sourceEndpointRegion;

    /**
     * @return The region of source instance.
     * 
     */
    public Output<String> sourceEndpointRegion() {
        return this.sourceEndpointRegion;
    }
    /**
     * The status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The sync architecture. Valid values: `oneway`.
     * 
     */
    @Export(name="syncArchitecture", type=String.class, parameters={})
    private Output</* @Nullable */ String> syncArchitecture;

    /**
     * @return The sync architecture. Valid values: `oneway`.
     * 
     */
    public Output<Optional<String>> syncArchitecture() {
        return Codegen.optional(this.syncArchitecture);
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MigrationInstance(String name) {
        this(name, MigrationInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MigrationInstance(String name, MigrationInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MigrationInstance(String name, MigrationInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/migrationInstance:MigrationInstance", name, args == null ? MigrationInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MigrationInstance(String name, Output<String> id, @Nullable MigrationInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/migrationInstance:MigrationInstance", name, state, makeResourceOptions(options, id));
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
    public static MigrationInstance get(String name, Output<String> id, @Nullable MigrationInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MigrationInstance(name, id, state, options);
    }
}
