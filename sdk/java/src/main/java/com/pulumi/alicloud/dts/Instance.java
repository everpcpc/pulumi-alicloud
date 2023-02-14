// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dts;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dts.InstanceArgs;
import com.pulumi.alicloud.dts.inputs.InstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Dts Instance resource.
 * 
 * For information about Dts Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/data-transmission-service/latest/createdtsinstance).
 * 
 * &gt; **NOTE:** Available in v1.198.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.dts.Instance;
 * import com.pulumi.alicloud.dts.InstanceArgs;
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
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status(&#34;OK&#34;)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .type(&#34;sync&#34;)
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .instanceClass(&#34;large&#34;)
 *             .sourceEndpointEngineName(&#34;MySQL&#34;)
 *             .sourceRegion(&#34;cn-hangzhou&#34;)
 *             .destinationEndpointEngineName(&#34;MySQL&#34;)
 *             .region(&#34;cn-hangzhou&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Dts Instance can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dts/instance:Instance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dts/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Whether to automatically renew the fee when it expires. Valid values:
     * - **false**: No, the default value.
     * - **true**: Yes.
     * 
     */
    @Export(name="autoPay", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoPay;

    /**
     * @return Whether to automatically renew the fee when it expires. Valid values:
     * - **false**: No, the default value.
     * - **true**: Yes.
     * 
     */
    public Output<Optional<Boolean>> autoPay() {
        return Codegen.optional(this.autoPay);
    }
    /**
     * Whether to start the task automatically after the purchase is completed. Value:
     * - **false**: No, the default value.
     * - **true**: Yes.
     * 
     */
    @Export(name="autoStart", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoStart;

    /**
     * @return Whether to start the task automatically after the purchase is completed. Value:
     * - **false**: No, the default value.
     * - **true**: Yes.
     * 
     */
    public Output<Optional<Boolean>> autoStart() {
        return Codegen.optional(this.autoStart);
    }
    /**
     * Specifications of ETL. The unit is compute unit (CU),1CU = 1vCPU +4GB of memory. The value range is an integer greater than or equal to 2. **NOTE:** Enter this parameter and enable ETL to clean and convert data.
     * 
     */
    @Export(name="computeUnit", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> computeUnit;

    /**
     * @return Specifications of ETL. The unit is compute unit (CU),1CU = 1vCPU +4GB of memory. The value range is an integer greater than or equal to 2. **NOTE:** Enter this parameter and enable ETL to clean and convert data.
     * 
     */
    public Output<Optional<Integer>> computeUnit() {
        return Codegen.optional(this.computeUnit);
    }
    /**
     * Instance creation time
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Instance creation time
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The number of private custom RDS instances in the PolarDB-X. The default value is **1**. **NOTE:** This parameter is required only when **source_endpoint_engine_name** is **DRDS**.
     * 
     */
    @Export(name="databaseCount", type=Integer.class, parameters={})
    private Output<Integer> databaseCount;

    /**
     * @return The number of private custom RDS instances in the PolarDB-X. The default value is **1**. **NOTE:** This parameter is required only when **source_endpoint_engine_name** is **DRDS**.
     * 
     */
    public Output<Integer> databaseCount() {
        return this.databaseCount;
    }
    /**
     * The target database engine type.
     * - **MySQL**:MySQL databases (including RDS MySQL and self-built MySQL).
     * - **PolarDB**:PolarDB MySQL.
     * - **polardb_o**:PolarDB O engine.
     * - **polardb_pg**:PolarDB PostgreSQL.
     * - **Redis**:Redis databases (including apsaradb for Redis and user-created Redis).
     * - **DRDS**: cloud-native distributed database PolarDB-X 1.0 and 2.0.
     * - **PostgreSQL**: User-created PostgreSQL.
     * - **ODPS**:MaxCompute project.
     * - **oracle**: self-built Oracle.
     * - **mongodb**:MongoDB databases (including apsaradb for MongoDB and user-created MongoDB).
     * - **tidb**:TiDB database.
     * - **ADS**: Cloud native data warehouse AnalyticDB MySQL 2.0.
     * - **ADB30**: Cloud native data warehouse AnalyticDB MySQL 3.0.
     * - **Greenplum**: Cloud native data warehouse AnalyticDB PostgreSQL.
     * - **MSSQL**:SQL Server databases (including RDS SQL Server and self-built SQL Server).
     * - **kafka**:Kafka databases (including Kafka and self-built Kafka).
     * - **DataHub**: DataHub, an Alibaba cloud streaming data service.
     * - **clickhouse**: ClickHouse.
     * - **DB2**: self-built DB2 LUW.
     * - **as400**:AS/400.
     * - **Tablestore**: Tablestore.
     * - **NOTE:**
     * - The default value is **MySQL**.
     * - For more information about the supported source and destination databases, see [Database, Synchronization Initialization Type, and Synchronization Topology](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-synchronization-scenarios-1) and [Supported Database and Migration Type](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-migration-scenarios).
     * - This parameter or **job_id** must be passed in.
     * 
     */
    @Export(name="destinationEndpointEngineName", type=String.class, parameters={})
    private Output<String> destinationEndpointEngineName;

    /**
     * @return The target database engine type.
     * - **MySQL**:MySQL databases (including RDS MySQL and self-built MySQL).
     * - **PolarDB**:PolarDB MySQL.
     * - **polardb_o**:PolarDB O engine.
     * - **polardb_pg**:PolarDB PostgreSQL.
     * - **Redis**:Redis databases (including apsaradb for Redis and user-created Redis).
     * - **DRDS**: cloud-native distributed database PolarDB-X 1.0 and 2.0.
     * - **PostgreSQL**: User-created PostgreSQL.
     * - **ODPS**:MaxCompute project.
     * - **oracle**: self-built Oracle.
     * - **mongodb**:MongoDB databases (including apsaradb for MongoDB and user-created MongoDB).
     * - **tidb**:TiDB database.
     * - **ADS**: Cloud native data warehouse AnalyticDB MySQL 2.0.
     * - **ADB30**: Cloud native data warehouse AnalyticDB MySQL 3.0.
     * - **Greenplum**: Cloud native data warehouse AnalyticDB PostgreSQL.
     * - **MSSQL**:SQL Server databases (including RDS SQL Server and self-built SQL Server).
     * - **kafka**:Kafka databases (including Kafka and self-built Kafka).
     * - **DataHub**: DataHub, an Alibaba cloud streaming data service.
     * - **clickhouse**: ClickHouse.
     * - **DB2**: self-built DB2 LUW.
     * - **as400**:AS/400.
     * - **Tablestore**: Tablestore.
     * - **NOTE:**
     * - The default value is **MySQL**.
     * - For more information about the supported source and destination databases, see [Database, Synchronization Initialization Type, and Synchronization Topology](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-synchronization-scenarios-1) and [Supported Database and Migration Type](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-migration-scenarios).
     * - This parameter or **job_id** must be passed in.
     * 
     */
    public Output<String> destinationEndpointEngineName() {
        return this.destinationEndpointEngineName;
    }
    /**
     * The target instance region. For more information, see [List of supported regions](https://www.alibabacloud.com/help/en/data-transmission-service/latest/list-of-supported-regions). **NOTE:** This parameter or **job_id** must be passed in.
     * 
     */
    @Export(name="destinationRegion", type=String.class, parameters={})
    private Output</* @Nullable */ String> destinationRegion;

    /**
     * @return The target instance region. For more information, see [List of supported regions](https://www.alibabacloud.com/help/en/data-transmission-service/latest/list-of-supported-regions). **NOTE:** This parameter or **job_id** must be passed in.
     * 
     */
    public Output<Optional<String>> destinationRegion() {
        return Codegen.optional(this.destinationRegion);
    }
    /**
     * The ID of the subscription instance.
     * 
     */
    @Export(name="dtsInstanceId", type=String.class, parameters={})
    private Output<String> dtsInstanceId;

    /**
     * @return The ID of the subscription instance.
     * 
     */
    public Output<String> dtsInstanceId() {
        return this.dtsInstanceId;
    }
    /**
     * Assign a specified number of DU resources to DTS tasks in the DTS exclusive cluster. Valid values: **1** ~ **100**. **NOTE:** The value of this parameter must be within the range of the number of DUs available for the DTS dedicated cluster.
     * 
     */
    @Export(name="du", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> du;

    /**
     * @return Assign a specified number of DU resources to DTS tasks in the DTS exclusive cluster. Valid values: **1** ~ **100**. **NOTE:** The value of this parameter must be within the range of the number of DUs available for the DTS dedicated cluster.
     * 
     */
    public Output<Optional<Integer>> du() {
        return Codegen.optional(this.du);
    }
    /**
     * Subscription billing type, Valid values: `ONLY_CONFIGURATION_FEE`: charges only configuration fees; `CONFIGURATION_FEE_AND_DATA_FEE`: charges configuration fees and data traffic fees.
     * 
     */
    @Export(name="feeType", type=String.class, parameters={})
    private Output</* @Nullable */ String> feeType;

    /**
     * @return Subscription billing type, Valid values: `ONLY_CONFIGURATION_FEE`: charges only configuration fees; `CONFIGURATION_FEE_AND_DATA_FEE`: charges configuration fees and data traffic fees.
     * 
     */
    public Output<Optional<String>> feeType() {
        return Codegen.optional(this.feeType);
    }
    /**
     * The type of the migration or synchronization instance.
     * - The specifications of the migration instance: **xxlarge**, **xlarge**, **large**, **medium**, **small**.
     * - The types of synchronization instances: **large**, **medium**, **small**, **micro**.
     * - **NOTE:** For performance descriptions of different specifications, see [Data Migration Link Specifications](https://www.alibabacloud.com/help/en/data-transmission-service/latest/cd773b) and [Data Synchronization Link Specifications](https://www.alibabacloud.com/help/en/data-transmission-service/latest/6bce7c).
     * 
     */
    @Export(name="instanceClass", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceClass;

    /**
     * @return The type of the migration or synchronization instance.
     * - The specifications of the migration instance: **xxlarge**, **xlarge**, **large**, **medium**, **small**.
     * - The types of synchronization instances: **large**, **medium**, **small**, **micro**.
     * - **NOTE:** For performance descriptions of different specifications, see [Data Migration Link Specifications](https://www.alibabacloud.com/help/en/data-transmission-service/latest/cd773b) and [Data Synchronization Link Specifications](https://www.alibabacloud.com/help/en/data-transmission-service/latest/6bce7c).
     * 
     */
    public Output<Optional<String>> instanceClass() {
        return Codegen.optional(this.instanceClass);
    }
    /**
     * The name of Dts instance.
     * 
     */
    @Export(name="instanceName", type=String.class, parameters={})
    private Output<String> instanceName;

    /**
     * @return The name of Dts instance.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * The ID of the task obtained by calling the **ConfigureDtsJob** operation (**DtsJobId**).&gt; After you pass in this parameter, you do not need to pass the **source_region**, **destination_region**, **type**, **source_endpoint_engine_name**, or **destination_endpoint_engine_name** parameters. Even if the input is passed in, the configuration in **job_id** shall prevail.
     * 
     */
    @Export(name="jobId", type=String.class, parameters={})
    private Output</* @Nullable */ String> jobId;

    /**
     * @return The ID of the task obtained by calling the **ConfigureDtsJob** operation (**DtsJobId**).&gt; After you pass in this parameter, you do not need to pass the **source_region**, **destination_region**, **type**, **source_endpoint_engine_name**, or **destination_endpoint_engine_name** parameters. Even if the input is passed in, the configuration in **job_id** shall prevail.
     * 
     */
    public Output<Optional<String>> jobId() {
        return Codegen.optional(this.jobId);
    }
    /**
     * The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output</* @Nullable */ String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
     * 
     */
    public Output<Optional<String>> paymentType() {
        return Codegen.optional(this.paymentType);
    }
    /**
     * The billing method of the subscription instance. Value: `Year`, `Month`. **NOTE:** This parameter is valid and must be passed in only when `payment_type` is `Subscription`.
     * 
     */
    @Export(name="period", type=String.class, parameters={})
    private Output</* @Nullable */ String> period;

    /**
     * @return The billing method of the subscription instance. Value: `Year`, `Month`. **NOTE:** This parameter is valid and must be passed in only when `payment_type` is `Subscription`.
     * 
     */
    public Output<Optional<String>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Resource Group ID.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return Resource Group ID.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Source instance database engine type.
     * - **MySQL**:MySQL databases (including RDS MySQL and self-built MySQL).
     * - **PolarDB**:PolarDB MySQL.
     * - **polardb_o**:PolarDB O engine.
     * - **polardb_pg**:PolarDB PostgreSQL.
     * - **Redis**:Redis databases (including apsaradb for Redis and user-created Redis).
     * - **DRDS**: cloud-native distributed database PolarDB-X 1.0 and 2.0.
     * - **PostgreSQL**: User-created PostgreSQL.
     * - **ODPS**:MaxCompute.
     * - **oracle**: self-built Oracle.
     * - **mongodb**:MongoDB databases (including apsaradb for MongoDB and user-created MongoDB).
     * - **tidb**:TiDB database.
     * - **ADS**: Cloud native data warehouse AnalyticDB MySQL 2.0.
     * - **ADB30**: Cloud native data warehouse AnalyticDB MySQL 3.0.
     * - **Greenplum**: Cloud native data warehouse AnalyticDB PostgreSQL.
     * - **MSSQL**:SQL Server databases (including RDS SQL Server and self-built SQL Server).
     * - **kafka**:Kafka databases (including Kafka and self-built Kafka).
     * - **DataHub**: DataHub, an Alibaba cloud streaming data service.
     * - **clickhouse**: ClickHouse.
     * - **DB2**: self-built DB2 LUW.
     * - **as400**:AS/400.
     * - **Tablestore**: Tablestore.
     * - **NOTE:**
     * - The default value is **MySQL**.
     * - For more information about the supported source and destination databases, see [Database, Synchronization Initialization Type, and Synchronization Topology](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-synchronization-scenarios-1) and [Supported Database and Migration Type](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-migration-scenarios).
     * - This parameter or **job_id** must be passed in.
     * 
     */
    @Export(name="sourceEndpointEngineName", type=String.class, parameters={})
    private Output<String> sourceEndpointEngineName;

    /**
     * @return Source instance database engine type.
     * - **MySQL**:MySQL databases (including RDS MySQL and self-built MySQL).
     * - **PolarDB**:PolarDB MySQL.
     * - **polardb_o**:PolarDB O engine.
     * - **polardb_pg**:PolarDB PostgreSQL.
     * - **Redis**:Redis databases (including apsaradb for Redis and user-created Redis).
     * - **DRDS**: cloud-native distributed database PolarDB-X 1.0 and 2.0.
     * - **PostgreSQL**: User-created PostgreSQL.
     * - **ODPS**:MaxCompute.
     * - **oracle**: self-built Oracle.
     * - **mongodb**:MongoDB databases (including apsaradb for MongoDB and user-created MongoDB).
     * - **tidb**:TiDB database.
     * - **ADS**: Cloud native data warehouse AnalyticDB MySQL 2.0.
     * - **ADB30**: Cloud native data warehouse AnalyticDB MySQL 3.0.
     * - **Greenplum**: Cloud native data warehouse AnalyticDB PostgreSQL.
     * - **MSSQL**:SQL Server databases (including RDS SQL Server and self-built SQL Server).
     * - **kafka**:Kafka databases (including Kafka and self-built Kafka).
     * - **DataHub**: DataHub, an Alibaba cloud streaming data service.
     * - **clickhouse**: ClickHouse.
     * - **DB2**: self-built DB2 LUW.
     * - **as400**:AS/400.
     * - **Tablestore**: Tablestore.
     * - **NOTE:**
     * - The default value is **MySQL**.
     * - For more information about the supported source and destination databases, see [Database, Synchronization Initialization Type, and Synchronization Topology](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-synchronization-scenarios-1) and [Supported Database and Migration Type](https://www.alibabacloud.com/help/en/data-transmission-service/latest/overview-of-data-migration-scenarios).
     * - This parameter or **job_id** must be passed in.
     * 
     */
    public Output<String> sourceEndpointEngineName() {
        return this.sourceEndpointEngineName;
    }
    /**
     * The source instance region. For more information, see [List of supported regions](https://www.alibabacloud.com/help/en/data-transmission-service/latest/list-of-supported-regions). **NOTE:** This parameter or **job_id** must be passed in.
     * 
     */
    @Export(name="sourceRegion", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceRegion;

    /**
     * @return The source instance region. For more information, see [List of supported regions](https://www.alibabacloud.com/help/en/data-transmission-service/latest/list-of-supported-regions). **NOTE:** This parameter or **job_id** must be passed in.
     * 
     */
    public Output<Optional<String>> sourceRegion() {
        return Codegen.optional(this.sourceRegion);
    }
    /**
     * Instance status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Instance status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Synchronization topology, value:
     * - **oneway**: one-way synchronization, the default value.
     * - **bidirectional**: two-way synchronization.
     * 
     */
    @Export(name="syncArchitecture", type=String.class, parameters={})
    private Output</* @Nullable */ String> syncArchitecture;

    /**
     * @return Synchronization topology, value:
     * - **oneway**: one-way synchronization, the default value.
     * - **bidirectional**: two-way synchronization.
     * 
     */
    public Output<Optional<String>> syncArchitecture() {
        return Codegen.optional(this.syncArchitecture);
    }
    /**
     * The synchronization direction. Default value: `Forward`. Valid values:
     * 
     */
    @Export(name="synchronizationDirection", type=String.class, parameters={})
    private Output</* @Nullable */ String> synchronizationDirection;

    /**
     * @return The synchronization direction. Default value: `Forward`. Valid values:
     * 
     */
    public Output<Optional<String>> synchronizationDirection() {
        return Codegen.optional(this.synchronizationDirection);
    }
    /**
     * The tag value corresponding to the tag key.See the following `Block Tags`.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tag value corresponding to the tag key.See the following `Block Tags`.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The instance type. Valid values:
     * - **migration**: MIGRATION.
     * - **sync**: synchronization.
     * - **subscribe**: SUBSCRIBE.
     * - **NOTE:** This parameter or **job_id** must be passed in.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output</* @Nullable */ String> type;

    /**
     * @return The instance type. Valid values:
     * - **migration**: MIGRATION.
     * - **sync**: synchronization.
     * - **subscribe**: SUBSCRIBE.
     * - **NOTE:** This parameter or **job_id** must be passed in.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * Prepaid instance purchase duration.
     * - When **period** is **Month**, the values are: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
     * - When **Period** is **Year**, the values are 1, 2, 3, and 5.
     * - **NOTE:**
     * - This parameter is valid and must be passed in only when **payment_type** is `Subscription`.
     * - The billing method of the subscription instance. You can set the parameter `period`.
     * 
     */
    @Export(name="usedTime", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> usedTime;

    /**
     * @return Prepaid instance purchase duration.
     * - When **period** is **Month**, the values are: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
     * - When **Period** is **Year**, the values are 1, 2, 3, and 5.
     * - **NOTE:**
     * - This parameter is valid and must be passed in only when **payment_type** is `Subscription`.
     * - The billing method of the subscription instance. You can set the parameter `period`.
     * 
     */
    public Output<Optional<Integer>> usedTime() {
        return Codegen.optional(this.usedTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, @Nullable InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, @Nullable InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
