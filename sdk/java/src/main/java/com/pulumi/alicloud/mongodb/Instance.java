// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mongodb.InstanceArgs;
import com.pulumi.alicloud.mongodb.inputs.InstanceState;
import com.pulumi.alicloud.mongodb.outputs.InstanceReplicaSet;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * MongoDB can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:mongodb/instance:Instance example dds-bp1291daeda44194
 * ```
 * 
 */
@ResourceType(type="alicloud:mongodb/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     * 
     */
    @Export(name="accountPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> accountPassword;

    /**
     * @return Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     * 
     */
    public Output<Optional<String>> accountPassword() {
        return Codegen.optional(this.accountPassword);
    }
    /**
     * Auto renew for prepaid, true of false. Default is false.
     * &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     * 
     */
    @Export(name="autoRenew", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Auto renew for prepaid, true of false. Default is false.
     * &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     * 
     */
    @Export(name="backupPeriods", type=List.class, parameters={String.class})
    private Output<List<String>> backupPeriods;

    /**
     * @return MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     * 
     */
    public Output<List<String>> backupPeriods() {
        return this.backupPeriods;
    }
    /**
     * MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    @Export(name="backupTime", type=String.class, parameters={})
    private Output<String> backupTime;

    /**
     * @return MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    public Output<String> backupTime() {
        return this.backupTime;
    }
    /**
     * Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    @Export(name="dbInstanceClass", type=String.class, parameters={})
    private Output<String> dbInstanceClass;

    /**
     * @return Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    public Output<String> dbInstanceClass() {
        return this.dbInstanceClass;
    }
    /**
     * User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     * 
     */
    @Export(name="dbInstanceStorage", type=Integer.class, parameters={})
    private Output<Integer> dbInstanceStorage;

    /**
     * @return User-defined DB instance storage space.Unit: GB. Value range:
     * - Custom storage space.
     * - 10-GB increments.
     * 
     */
    public Output<Integer> dbInstanceStorage() {
        return this.dbInstanceStorage;
    }
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     * 
     */
    @Export(name="engineVersion", type=String.class, parameters={})
    private Output<String> engineVersion;

    /**
     * @return Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zone_id` and `secondary_zone_id` parameter values.
     * 
     */
    @Export(name="hiddenZoneId", type=String.class, parameters={})
    private Output</* @Nullable */ String> hiddenZoneId;

    /**
     * @return Configure the zone where the hidden node is located to deploy multiple zones. **NOTE:** This parameter value cannot be the same as `zone_id` and `secondary_zone_id` parameter values.
     * 
     */
    public Output<Optional<String>> hiddenZoneId() {
        return Codegen.optional(this.hiddenZoneId);
    }
    /**
     * Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
     * 
     */
    @Export(name="instanceChargeType", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceChargeType;

    /**
     * @return Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
     * 
     */
    public Output<Optional<String>> instanceChargeType() {
        return Codegen.optional(this.instanceChargeType);
    }
    /**
     * An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedPassword() {
        return Codegen.optional(this.kmsEncryptedPassword);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Export(name="kmsEncryptionContext", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    @Export(name="maintainEndTime", type=String.class, parameters={})
    private Output<String> maintainEndTime;

    /**
     * @return The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    public Output<String> maintainEndTime() {
        return this.maintainEndTime;
    }
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    @Export(name="maintainStartTime", type=String.class, parameters={})
    private Output<String> maintainStartTime;

    /**
     * @return The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     * 
     */
    public Output<String> maintainStartTime() {
        return this.maintainStartTime;
    }
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of DB instance. It a string of 2 to 256 characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     * 
     */
    @Export(name="networkType", type=String.class, parameters={})
    private Output<String> networkType;

    /**
     * @return The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     *   Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
     * 
     */
    @Export(name="orderType", type=String.class, parameters={})
    private Output</* @Nullable */ String> orderType;

    /**
     * @return The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     *   Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
     * 
     */
    public Output<Optional<String>> orderType() {
        return Codegen.optional(this.orderType);
    }
    /**
     * The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     * 
     */
    @Export(name="period", type=Integer.class, parameters={})
    private Output<Integer> period;

    /**
     * @return The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     * 
     */
    public Output<Integer> period() {
        return this.period;
    }
    /**
     * The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     * 
     */
    @Export(name="readonlyReplicas", type=Integer.class, parameters={})
    private Output<Integer> readonlyReplicas;

    /**
     * @return The number of read-only nodes in the replica set instance. Default value: 0. Valid values: 0 to 5.
     * 
     */
    public Output<Integer> readonlyReplicas() {
        return this.readonlyReplicas;
    }
    /**
     * The name of the mongo replica set
     * 
     */
    @Export(name="replicaSetName", type=String.class, parameters={})
    private Output<String> replicaSetName;

    /**
     * @return The name of the mongo replica set
     * 
     */
    public Output<String> replicaSetName() {
        return this.replicaSetName;
    }
    /**
     * Replica set instance information. The details see Block replica_sets. **NOTE:** Available in v1.140+.
     * 
     */
    @Export(name="replicaSets", type=List.class, parameters={InstanceReplicaSet.class})
    private Output<List<InstanceReplicaSet>> replicaSets;

    /**
     * @return Replica set instance information. The details see Block replica_sets. **NOTE:** Available in v1.140+.
     * 
     */
    public Output<List<InstanceReplicaSet>> replicaSets() {
        return this.replicaSets;
    }
    /**
     * Number of replica set nodes. Valid values: [1, 3, 5, 7]
     * 
     */
    @Export(name="replicationFactor", type=Integer.class, parameters={})
    private Output<Integer> replicationFactor;

    /**
     * @return Number of replica set nodes. Valid values: [1, 3, 5, 7]
     * 
     */
    public Output<Integer> replicationFactor() {
        return this.replicationFactor;
    }
    /**
     * The ID of the Resource Group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the Resource Group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Instance log backup retention days. Available in 1.42.0+.
     * 
     */
    @Export(name="retentionPeriod", type=Integer.class, parameters={})
    private Output<Integer> retentionPeriod;

    /**
     * @return Instance log backup retention days. Available in 1.42.0+.
     * 
     */
    public Output<Integer> retentionPeriod() {
        return this.retentionPeriod;
    }
    /**
     * Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zone_id` and `hidden_zone_id` parameter values.
     * 
     */
    @Export(name="secondaryZoneId", type=String.class, parameters={})
    private Output</* @Nullable */ String> secondaryZoneId;

    /**
     * @return Configure the available area where the slave node (Secondary node) is located to realize multi-available area deployment. **NOTE:** This parameter value cannot be the same as `zone_id` and `hidden_zone_id` parameter values.
     * 
     */
    public Output<Optional<String>> secondaryZoneId() {
        return Codegen.optional(this.secondaryZoneId);
    }
    /**
     * The Security Group ID of ECS.
     * 
     */
    @Export(name="securityGroupId", type=String.class, parameters={})
    private Output<String> securityGroupId;

    /**
     * @return The Security Group ID of ECS.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    @Export(name="securityIpLists", type=List.class, parameters={String.class})
    private Output<List<String>> securityIpLists;

    /**
     * @return List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     * 
     */
    public Output<List<String>> securityIpLists() {
        return this.securityIpLists;
    }
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
     * 
     */
    @Export(name="sslAction", type=String.class, parameters={})
    private Output<String> sslAction;

    /**
     * @return Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
     * 
     */
    public Output<String> sslAction() {
        return this.sslAction;
    }
    /**
     * Status of the SSL feature. `Open`: SSL is turned on; `Closed`: SSL is turned off.
     * 
     */
    @Export(name="sslStatus", type=String.class, parameters={})
    private Output<String> sslStatus;

    /**
     * @return Status of the SSL feature. `Open`: SSL is turned on; `Closed`: SSL is turned off.
     * 
     */
    public Output<String> sslStatus() {
        return this.sslStatus;
    }
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     * 
     */
    @Export(name="storageEngine", type=String.class, parameters={})
    private Output<String> storageEngine;

    /**
     * @return Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     * 
     */
    public Output<String> storageEngine() {
        return this.storageEngine;
    }
    /**
     * The storage type of the instance. Valid values: `cloud_essd1`, `cloud_essd2`, `cloud_essd3`, `local_ssd`.
     * 
     */
    @Export(name="storageType", type=String.class, parameters={})
    private Output<String> storageType;

    /**
     * @return The storage type of the instance. Valid values: `cloud_essd1`, `cloud_essd2`, `cloud_essd3`, `local_ssd`.
     * 
     */
    public Output<String> storageType() {
        return this.storageType;
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
     * The TDE(Transparent Data Encryption) status.
     * 
     */
    @Export(name="tdeStatus", type=String.class, parameters={})
    private Output</* @Nullable */ String> tdeStatus;

    /**
     * @return The TDE(Transparent Data Encryption) status.
     * 
     */
    public Output<Optional<String>> tdeStatus() {
        return Codegen.optional(this.tdeStatus);
    }
    /**
     * The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to &#34;true&#34; in the data source `alicloud.getZones`.
     * 
     */
    @Export(name="zoneId", type=String.class, parameters={})
    private Output<String> zoneId;

    /**
     * @return The Zone to launch the DB instance. it supports multiple zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * The multiple zone ID can be retrieved by setting `multi` to &#34;true&#34; in the data source `alicloud.getZones`.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
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
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountPassword"
            ))
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
