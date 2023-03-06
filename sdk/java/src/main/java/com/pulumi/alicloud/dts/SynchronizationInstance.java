// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dts;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dts.SynchronizationInstanceArgs;
import com.pulumi.alicloud.dts.inputs.SynchronizationInstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DTS Synchronization Instance resource.
 * 
 * For information about DTS Synchronization Instance and how to use it, see [What is Synchronization Instance](https://www.alibabacloud.com/help/en/doc-detail/130744.html).
 * 
 * &gt; **NOTE:** Available in v1.138.0+.
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
 * import com.pulumi.alicloud.dts.SynchronizationInstance;
 * import com.pulumi.alicloud.dts.SynchronizationInstanceArgs;
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
 *         var default_ = new SynchronizationInstance(&#34;default&#34;, SynchronizationInstanceArgs.builder()        
 *             .destinationEndpointEngineName(&#34;ADB30&#34;)
 *             .destinationEndpointRegion(&#34;cn-hangzhou&#34;)
 *             .instanceClass(&#34;small&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .sourceEndpointEngineName(&#34;PolarDB&#34;)
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
 * DTS Synchronization Instance can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dts/synchronizationInstance:SynchronizationInstance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dts/synchronizationInstance:SynchronizationInstance")
public class SynchronizationInstance extends com.pulumi.resources.CustomResource {
    /**
     * Whether to automatically renew when it expires. Valid values: `true`, `false`.
     * 
     */
    @Export(name="autoPay", type=String.class, parameters={})
    private Output</* @Nullable */ String> autoPay;

    /**
     * @return Whether to automatically renew when it expires. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<String>> autoPay() {
        return Codegen.optional(this.autoPay);
    }
    /**
     * Whether to automatically start the task after the purchase completed. Valid values: `true`, `false`.
     * 
     */
    @Export(name="autoStart", type=String.class, parameters={})
    private Output</* @Nullable */ String> autoStart;

    /**
     * @return Whether to automatically start the task after the purchase completed. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<String>> autoStart() {
        return Codegen.optional(this.autoStart);
    }
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
     * The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     * 
     */
    @Export(name="instanceClass", type=String.class, parameters={})
    private Output<String> instanceClass;

    /**
     * @return The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
     * 
     */
    public Output<String> instanceClass() {
        return this.instanceClass;
    }
    /**
     * The duration of prepaid instance purchase. this parameter is required When `payment_type` equals `Subscription`.
     * 
     */
    @Export(name="paymentDuration", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> paymentDuration;

    /**
     * @return The duration of prepaid instance purchase. this parameter is required When `payment_type` equals `Subscription`.
     * 
     */
    public Output<Optional<Integer>> paymentDuration() {
        return Codegen.optional(this.paymentDuration);
    }
    /**
     * The payment duration unit. Valid values: `Month`, `Year`. When `payment_type` is `Subscription`, this parameter is valid and must be passed in.
     * 
     */
    @Export(name="paymentDurationUnit", type=String.class, parameters={})
    private Output</* @Nullable */ String> paymentDurationUnit;

    /**
     * @return The payment duration unit. Valid values: `Month`, `Year`. When `payment_type` is `Subscription`, this parameter is valid and must be passed in.
     * 
     */
    public Output<Optional<String>> paymentDurationUnit() {
        return Codegen.optional(this.paymentDurationUnit);
    }
    /**
     * The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The number of instances purchased.
     * 
     */
    @Export(name="quantity", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> quantity;

    /**
     * @return The number of instances purchased.
     * 
     */
    public Output<Optional<Integer>> quantity() {
        return Codegen.optional(this.quantity);
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
     * The sync architecture. Valid values: `oneway`, `bidirectional`.
     * 
     */
    @Export(name="syncArchitecture", type=String.class, parameters={})
    private Output</* @Nullable */ String> syncArchitecture;

    /**
     * @return The sync architecture. Valid values: `oneway`, `bidirectional`.
     * 
     */
    public Output<Optional<String>> syncArchitecture() {
        return Codegen.optional(this.syncArchitecture);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SynchronizationInstance(String name) {
        this(name, SynchronizationInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SynchronizationInstance(String name, SynchronizationInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SynchronizationInstance(String name, SynchronizationInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/synchronizationInstance:SynchronizationInstance", name, args == null ? SynchronizationInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SynchronizationInstance(String name, Output<String> id, @Nullable SynchronizationInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/synchronizationInstance:SynchronizationInstance", name, state, makeResourceOptions(options, id));
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
    public static SynchronizationInstance get(String name, Output<String> id, @Nullable SynchronizationInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SynchronizationInstance(name, id, state, options);
    }
}