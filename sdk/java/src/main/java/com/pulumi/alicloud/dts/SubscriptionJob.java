// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dts;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dts.SubscriptionJobArgs;
import com.pulumi.alicloud.dts.inputs.SubscriptionJobState;
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
 * Provides a DTS Subscription Job resource.
 * 
 * For information about DTS Subscription Job and how to use it, see [What is Subscription Job](https://help.aliyun.com/document_detail/254791.html).
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.Database;
 * import com.pulumi.alicloud.rds.DatabaseArgs;
 * import com.pulumi.alicloud.rds.Account;
 * import com.pulumi.alicloud.rds.AccountArgs;
 * import com.pulumi.alicloud.rds.AccountPrivilege;
 * import com.pulumi.alicloud.rds.AccountPrivilegeArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.dts.SubscriptionJob;
 * import com.pulumi.alicloud.dts.SubscriptionJobArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
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
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;dtsSubscriptionJob&#34;);
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;Rds&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .instanceType(&#34;rds.mysql.s1.small&#34;)
 *             .instanceStorage(&#34;10&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .instanceName(name)
 *             .build());
 * 
 *         for (var i = 0; i &lt; 2; i++) {
 *             new Database(&#34;db-&#34; + i, DatabaseArgs.builder()            
 *                 .instanceId(instance.id())
 *                 .description(&#34;from terraform&#34;)
 *                 .build());
 * 
 *         
 * }
 *         var account = new Account(&#34;account&#34;, AccountArgs.builder()        
 *             .instanceId(instance.id())
 *             .password(&#34;Test12345&#34;)
 *             .description(&#34;from terraform&#34;)
 *             .build());
 * 
 *         var privilege = new AccountPrivilege(&#34;privilege&#34;, AccountPrivilegeArgs.builder()        
 *             .instanceId(instance.id())
 *             .accountName(account.name())
 *             .privilege(&#34;ReadWrite&#34;)
 *             .dbNames(db.stream().map(element -&gt; element.name()).collect(toList()))
 *             .build());
 * 
 *         final var default1Networks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;default-NODELETING&#34;)
 *             .build());
 * 
 *         final var default1Switches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(data.alicloud_vpcs().default().ids()[0])
 *             .build());
 * 
 *         var defaultSubscriptionJob = new SubscriptionJob(&#34;defaultSubscriptionJob&#34;, SubscriptionJobArgs.builder()        
 *             .dtsJobName(name)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .sourceEndpointEngineName(&#34;MySQL&#34;)
 *             .sourceEndpointRegion(&#34;cn-hangzhou&#34;)
 *             .sourceEndpointInstanceType(&#34;RDS&#34;)
 *             .sourceEndpointInstanceId(instance.id())
 *             .sourceEndpointDatabaseName(&#34;tfaccountpri_0&#34;)
 *             .sourceEndpointUserName(&#34;tftestprivilege&#34;)
 *             .sourceEndpointPassword(&#34;Test12345&#34;)
 *             .dbList(&#34;&#34;&#34;
 *         {&#34;dtstestdata&#34;: {&#34;name&#34;: &#34;tfaccountpri_0&#34;, &#34;all&#34;: true}}
 *             &#34;&#34;&#34;)
 *             .subscriptionInstanceNetworkType(&#34;vpc&#34;)
 *             .subscriptionInstanceVpcId(default1Networks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .subscriptionInstanceVswitchId(default1Switches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]))
 *             .status(&#34;Normal&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DTS Subscription Job can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dts/subscriptionJob:SubscriptionJob example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dts/subscriptionJob:SubscriptionJob")
public class SubscriptionJob extends com.pulumi.resources.CustomResource {
    /**
     * Subscription start time in Unix timestamp format.
     * 
     */
    @Export(name="checkpoint", type=String.class, parameters={})
    private Output<String> checkpoint;

    /**
     * @return Subscription start time in Unix timestamp format.
     * 
     */
    public Output<String> checkpoint() {
        return this.checkpoint;
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
     * Subscription object, in the format of JSON strings. For detailed definitions, please refer to the description of migration, synchronization or subscription objects [document](https://help.aliyun.com/document_detail/209545.html).
     * 
     */
    @Export(name="dbList", type=String.class, parameters={})
    private Output</* @Nullable */ String> dbList;

    /**
     * @return Subscription object, in the format of JSON strings. For detailed definitions, please refer to the description of migration, synchronization or subscription objects [document](https://help.aliyun.com/document_detail/209545.html).
     * 
     */
    public Output<Optional<String>> dbList() {
        return Codegen.optional(this.dbList);
    }
    /**
     * This parameter decides whether to monitor the delay status. Valid values: `true`, `false`.
     * 
     */
    @Export(name="delayNotice", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> delayNotice;

    /**
     * @return This parameter decides whether to monitor the delay status. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> delayNotice() {
        return Codegen.optional(this.delayNotice);
    }
    /**
     * The mobile phone number of the contact who delayed the alarm. Multiple mobile phone numbers separated by English commas `,`. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.
     * 
     */
    @Export(name="delayPhone", type=String.class, parameters={})
    private Output</* @Nullable */ String> delayPhone;

    /**
     * @return The mobile phone number of the contact who delayed the alarm. Multiple mobile phone numbers separated by English commas `,`. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.
     * 
     */
    public Output<Optional<String>> delayPhone() {
        return Codegen.optional(this.delayPhone);
    }
    /**
     * When `delay_notice` is set to `true`, this parameter must be passed in. The threshold for triggering the delay alarm. The unit is second and needs to be an integer. The threshold can be set according to business needs. It is recommended to set it above 10 seconds to avoid delay fluctuations caused by network and database load.
     * 
     */
    @Export(name="delayRuleTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> delayRuleTime;

    /**
     * @return When `delay_notice` is set to `true`, this parameter must be passed in. The threshold for triggering the delay alarm. The unit is second and needs to be an integer. The threshold can be set according to business needs. It is recommended to set it above 10 seconds to avoid delay fluctuations caused by network and database load.
     * 
     */
    public Output<Optional<String>> delayRuleTime() {
        return Codegen.optional(this.delayRuleTime);
    }
    /**
     * The destination endpoint engine name. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`.
     * 
     */
    @Export(name="destinationEndpointEngineName", type=String.class, parameters={})
    private Output</* @Nullable */ String> destinationEndpointEngineName;

    /**
     * @return The destination endpoint engine name. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`.
     * 
     */
    public Output<Optional<String>> destinationEndpointEngineName() {
        return Codegen.optional(this.destinationEndpointEngineName);
    }
    /**
     * The destination region. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     * 
     */
    @Export(name="destinationRegion", type=String.class, parameters={})
    private Output</* @Nullable */ String> destinationRegion;

    /**
     * @return The destination region. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
     * 
     */
    public Output<Optional<String>> destinationRegion() {
        return Codegen.optional(this.destinationRegion);
    }
    /**
     * The ID of subscription instance.
     * 
     */
    @Export(name="dtsInstanceId", type=String.class, parameters={})
    private Output<String> dtsInstanceId;

    /**
     * @return The ID of subscription instance.
     * 
     */
    public Output<String> dtsInstanceId() {
        return this.dtsInstanceId;
    }
    /**
     * The name of subscription task.
     * 
     */
    @Export(name="dtsJobName", type=String.class, parameters={})
    private Output</* @Nullable */ String> dtsJobName;

    /**
     * @return The name of subscription task.
     * 
     */
    public Output<Optional<String>> dtsJobName() {
        return Codegen.optional(this.dtsJobName);
    }
    /**
     * This parameter decides whether to monitor abnormal status. Valid values: `true`, `false`.
     * 
     */
    @Export(name="errorNotice", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> errorNotice;

    /**
     * @return This parameter decides whether to monitor abnormal status. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> errorNotice() {
        return Codegen.optional(this.errorNotice);
    }
    /**
     * The mobile phone number of the contact for abnormal alarm. Multiple mobile phone numbers separated by English commas `,`. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.
     * 
     */
    @Export(name="errorPhone", type=String.class, parameters={})
    private Output</* @Nullable */ String> errorPhone;

    /**
     * @return The mobile phone number of the contact for abnormal alarm. Multiple mobile phone numbers separated by English commas `,`. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.
     * 
     */
    public Output<Optional<String>> errorPhone() {
        return Codegen.optional(this.errorPhone);
    }
    /**
     * The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`.
     * 
     */
    @Export(name="instanceClass", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceClass;

    /**
     * @return The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`.
     * 
     */
    public Output<Optional<String>> instanceClass() {
        return Codegen.optional(this.instanceClass);
    }
    /**
     * The duration of prepaid instance purchase. When `payment_type` is `Subscription`, this parameter is valid and must be passed in.
     * 
     */
    @Export(name="paymentDuration", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> paymentDuration;

    /**
     * @return The duration of prepaid instance purchase. When `payment_type` is `Subscription`, this parameter is valid and must be passed in.
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
     * DTS reserves parameters, the format is a JSON string, you can pass in this parameter to complete the source and target database information (such as the data storage format of the target Kafka database, the instance ID of the cloud enterprise network CEN). For more information, please refer to the parameter description of the [Reserve parameter](https://help.aliyun.com/document_detail/176470.html).
     * 
     */
    @Export(name="reserve", type=String.class, parameters={})
    private Output</* @Nullable */ String> reserve;

    /**
     * @return DTS reserves parameters, the format is a JSON string, you can pass in this parameter to complete the source and target database information (such as the data storage format of the target Kafka database, the instance ID of the cloud enterprise network CEN). For more information, please refer to the parameter description of the [Reserve parameter](https://help.aliyun.com/document_detail/176470.html).
     * 
     */
    public Output<Optional<String>> reserve() {
        return Codegen.optional(this.reserve);
    }
    /**
     * To subscribe to the name of the database.
     * 
     */
    @Export(name="sourceEndpointDatabaseName", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointDatabaseName;

    /**
     * @return To subscribe to the name of the database.
     * 
     */
    public Output<Optional<String>> sourceEndpointDatabaseName() {
        return Codegen.optional(this.sourceEndpointDatabaseName);
    }
    /**
     * The source database type value is MySQL or Oracle. Valid values: `MySQL`, `Oracle`.
     * 
     */
    @Export(name="sourceEndpointEngineName", type=String.class, parameters={})
    private Output<String> sourceEndpointEngineName;

    /**
     * @return The source database type value is MySQL or Oracle. Valid values: `MySQL`, `Oracle`.
     * 
     */
    public Output<String> sourceEndpointEngineName() {
        return this.sourceEndpointEngineName;
    }
    /**
     * The ID of source instance. Only when the type of source database instance was RDS MySQL, PolarDB-X 1.0, PolarDB MySQL, this parameter can be available and must be set.
     * 
     */
    @Export(name="sourceEndpointInstanceId", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointInstanceId;

    /**
     * @return The ID of source instance. Only when the type of source database instance was RDS MySQL, PolarDB-X 1.0, PolarDB MySQL, this parameter can be available and must be set.
     * 
     */
    public Output<Optional<String>> sourceEndpointInstanceId() {
        return Codegen.optional(this.sourceEndpointInstanceId);
    }
    /**
     * The type of source instance. Valid values: `RDS`, `PolarDB`, `DRDS`, `LocalInstance`, `ECS`, `Express`, `CEN`, `dg`.
     * 
     */
    @Export(name="sourceEndpointInstanceType", type=String.class, parameters={})
    private Output<String> sourceEndpointInstanceType;

    /**
     * @return The type of source instance. Valid values: `RDS`, `PolarDB`, `DRDS`, `LocalInstance`, `ECS`, `Express`, `CEN`, `dg`.
     * 
     */
    public Output<String> sourceEndpointInstanceType() {
        return this.sourceEndpointInstanceType;
    }
    /**
     * The IP of source endpoint.
     * 
     */
    @Export(name="sourceEndpointIp", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointIp;

    /**
     * @return The IP of source endpoint.
     * 
     */
    public Output<Optional<String>> sourceEndpointIp() {
        return Codegen.optional(this.sourceEndpointIp);
    }
    /**
     * The SID of Oracle Database. When the source database is self-built Oracle and the Oracle database is a non-RAC instance, this parameter is available and must be passed in.
     * 
     */
    @Export(name="sourceEndpointOracleSid", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointOracleSid;

    /**
     * @return The SID of Oracle Database. When the source database is self-built Oracle and the Oracle database is a non-RAC instance, this parameter is available and must be passed in.
     * 
     */
    public Output<Optional<String>> sourceEndpointOracleSid() {
        return Codegen.optional(this.sourceEndpointOracleSid);
    }
    /**
     * The Alibaba Cloud account ID to which the source instance belongs. This parameter is only available when configuring data subscriptions across Alibaba Cloud accounts and must be passed in.
     * 
     */
    @Export(name="sourceEndpointOwnerId", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointOwnerId;

    /**
     * @return The Alibaba Cloud account ID to which the source instance belongs. This parameter is only available when configuring data subscriptions across Alibaba Cloud accounts and must be passed in.
     * 
     */
    public Output<Optional<String>> sourceEndpointOwnerId() {
        return Codegen.optional(this.sourceEndpointOwnerId);
    }
    /**
     * The password of source database instance account.
     * 
     */
    @Export(name="sourceEndpointPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointPassword;

    /**
     * @return The password of source database instance account.
     * 
     */
    public Output<Optional<String>> sourceEndpointPassword() {
        return Codegen.optional(this.sourceEndpointPassword);
    }
    /**
     * The port of source database.
     * 
     */
    @Export(name="sourceEndpointPort", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointPort;

    /**
     * @return The port of source database.
     * 
     */
    public Output<Optional<String>> sourceEndpointPort() {
        return Codegen.optional(this.sourceEndpointPort);
    }
    /**
     * The region of source database.
     * 
     */
    @Export(name="sourceEndpointRegion", type=String.class, parameters={})
    private Output<String> sourceEndpointRegion;

    /**
     * @return The region of source database.
     * 
     */
    public Output<String> sourceEndpointRegion() {
        return this.sourceEndpointRegion;
    }
    /**
     * Both the authorization roles. When the source instance and configure subscriptions task of the Alibaba Cloud account is not the same as the need to pass the parameter, to specify the source of the authorization roles, to allow configuration subscription task of the Alibaba Cloud account to access the source of the source instance information.
     * 
     */
    @Export(name="sourceEndpointRole", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointRole;

    /**
     * @return Both the authorization roles. When the source instance and configure subscriptions task of the Alibaba Cloud account is not the same as the need to pass the parameter, to specify the source of the authorization roles, to allow configuration subscription task of the Alibaba Cloud account to access the source of the source instance information.
     * 
     */
    public Output<Optional<String>> sourceEndpointRole() {
        return Codegen.optional(this.sourceEndpointRole);
    }
    /**
     * The username of source database instance account.
     * 
     */
    @Export(name="sourceEndpointUserName", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceEndpointUserName;

    /**
     * @return The username of source database instance account.
     * 
     */
    public Output<Optional<String>> sourceEndpointUserName() {
        return Codegen.optional(this.sourceEndpointUserName);
    }
    /**
     * The status of the task. Valid values: `Normal`, `Abnormal`. When a task created, it is in this state of `NotStarted`. You can specify this state to `Normal` to start the job, and specify this state of `Abnormal` to stop the job. **Note: We treat the state `Starting` as the state of `Normal`, and consider the two states to be consistent on the user side.**
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the task. Valid values: `Normal`, `Abnormal`. When a task created, it is in this state of `NotStarted`. You can specify this state to `Normal` to start the job, and specify this state of `Abnormal` to stop the job. **Note: We treat the state `Starting` as the state of `Normal`, and consider the two states to be consistent on the user side.**
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Whether to subscribe the DDL type of data. Valid values: `true`, `false`.
     * 
     */
    @Export(name="subscriptionDataTypeDdl", type=Boolean.class, parameters={})
    private Output<Boolean> subscriptionDataTypeDdl;

    /**
     * @return Whether to subscribe the DDL type of data. Valid values: `true`, `false`.
     * 
     */
    public Output<Boolean> subscriptionDataTypeDdl() {
        return this.subscriptionDataTypeDdl;
    }
    /**
     * Whether to subscribe the DML type of data. Valid values: `true`, `false`.
     * 
     */
    @Export(name="subscriptionDataTypeDml", type=Boolean.class, parameters={})
    private Output<Boolean> subscriptionDataTypeDml;

    /**
     * @return Whether to subscribe the DML type of data. Valid values: `true`, `false`.
     * 
     */
    public Output<Boolean> subscriptionDataTypeDml() {
        return this.subscriptionDataTypeDml;
    }
    /**
     * Subscription task type of network value: classic: classic Network. Virtual Private Cloud (vpc): a vpc. Valid values: `classic`, `vpc`.
     * 
     */
    @Export(name="subscriptionInstanceNetworkType", type=String.class, parameters={})
    private Output</* @Nullable */ String> subscriptionInstanceNetworkType;

    /**
     * @return Subscription task type of network value: classic: classic Network. Virtual Private Cloud (vpc): a vpc. Valid values: `classic`, `vpc`.
     * 
     */
    public Output<Optional<String>> subscriptionInstanceNetworkType() {
        return Codegen.optional(this.subscriptionInstanceNetworkType);
    }
    /**
     * The ID of subscription vpc instance. When the value of `subscription_instance_network_type` is vpc, this parameter is available and must be passed in.
     * 
     */
    @Export(name="subscriptionInstanceVpcId", type=String.class, parameters={})
    private Output</* @Nullable */ String> subscriptionInstanceVpcId;

    /**
     * @return The ID of subscription vpc instance. When the value of `subscription_instance_network_type` is vpc, this parameter is available and must be passed in.
     * 
     */
    public Output<Optional<String>> subscriptionInstanceVpcId() {
        return Codegen.optional(this.subscriptionInstanceVpcId);
    }
    /**
     * The ID of subscription VSwitch instance. When the value of `subscription_instance_network_type` is vpc, this parameter is available and must be passed in.
     * 
     */
    @Export(name="subscriptionInstanceVswitchId", type=String.class, parameters={})
    private Output</* @Nullable */ String> subscriptionInstanceVswitchId;

    /**
     * @return The ID of subscription VSwitch instance. When the value of `subscription_instance_network_type` is vpc, this parameter is available and must be passed in.
     * 
     */
    public Output<Optional<String>> subscriptionInstanceVswitchId() {
        return Codegen.optional(this.subscriptionInstanceVswitchId);
    }
    /**
     * The sync architecture. Valid values: `bidirectional`, `oneway`.
     * 
     */
    @Export(name="syncArchitecture", type=String.class, parameters={})
    private Output</* @Nullable */ String> syncArchitecture;

    /**
     * @return The sync architecture. Valid values: `bidirectional`, `oneway`.
     * 
     */
    public Output<Optional<String>> syncArchitecture() {
        return Codegen.optional(this.syncArchitecture);
    }
    /**
     * The synchronization direction. Valid values: `Forward`, `Reverse`. When the topology type of the data synchronization instance is bidirectional, it can be passed in to reverse to start the reverse synchronization link.
     * 
     */
    @Export(name="synchronizationDirection", type=String.class, parameters={})
    private Output</* @Nullable */ String> synchronizationDirection;

    /**
     * @return The synchronization direction. Valid values: `Forward`, `Reverse`. When the topology type of the data synchronization instance is bidirectional, it can be passed in to reverse to start the reverse synchronization link.
     * 
     */
    public Output<Optional<String>> synchronizationDirection() {
        return Codegen.optional(this.synchronizationDirection);
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
    public SubscriptionJob(String name) {
        this(name, SubscriptionJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SubscriptionJob(String name, SubscriptionJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SubscriptionJob(String name, SubscriptionJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/subscriptionJob:SubscriptionJob", name, args == null ? SubscriptionJobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SubscriptionJob(String name, Output<String> id, @Nullable SubscriptionJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/subscriptionJob:SubscriptionJob", name, state, makeResourceOptions(options, id));
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
    public static SubscriptionJob get(String name, Output<String> id, @Nullable SubscriptionJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SubscriptionJob(name, id, state, options);
    }
}
