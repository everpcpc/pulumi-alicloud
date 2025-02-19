// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb.inputs;

import com.pulumi.alicloud.mongodb.inputs.ShardingInstanceConfigServerListArgs;
import com.pulumi.alicloud.mongodb.inputs.ShardingInstanceMongoListArgs;
import com.pulumi.alicloud.mongodb.inputs.ShardingInstanceShardListArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ShardingInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final ShardingInstanceState Empty = new ShardingInstanceState();

    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     * 
     */
    @Import(name="accountPassword")
    private @Nullable Output<String> accountPassword;

    /**
     * @return Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     * 
     */
    public Optional<Output<String>> accountPassword() {
        return Optional.ofNullable(this.accountPassword);
    }

    /**
     * Auto renew for prepaid, true of false. Default is false.
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<Boolean> autoRenew;

    /**
     * @return Auto renew for prepaid, true of false. Default is false.
     * 
     */
    public Optional<Output<Boolean>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     * 
     */
    @Import(name="backupPeriods")
    private @Nullable Output<List<String>> backupPeriods;

    /**
     * @return MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     * 
     */
    public Optional<Output<List<String>>> backupPeriods() {
        return Optional.ofNullable(this.backupPeriods);
    }

    /**
     * MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    @Import(name="backupTime")
    private @Nullable Output<String> backupTime;

    /**
     * @return MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
     * 
     */
    public Optional<Output<String>> backupTime() {
        return Optional.ofNullable(this.backupTime);
    }

    /**
     * The node information list of config server. The details see Block `config_server_list`. **NOTE:** Available in v1.140+.
     * 
     */
    @Import(name="configServerLists")
    private @Nullable Output<List<ShardingInstanceConfigServerListArgs>> configServerLists;

    /**
     * @return The node information list of config server. The details see Block `config_server_list`. **NOTE:** Available in v1.140+.
     * 
     */
    public Optional<Output<List<ShardingInstanceConfigServerListArgs>>> configServerLists() {
        return Optional.ofNullable(this.configServerLists);
    }

    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
     * 
     */
    @Import(name="engineVersion")
    private @Nullable Output<String> engineVersion;

    /**
     * @return Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
     * 
     */
    public Optional<Output<String>> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }

    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     * 
     */
    @Import(name="instanceChargeType")
    private @Nullable Output<String> instanceChargeType;

    /**
     * @return Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     * 
     */
    public Optional<Output<String>> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }

    /**
     * An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    @Import(name="kmsEncryptedPassword")
    private @Nullable Output<String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
     * 
     */
    public Optional<Output<String>> kmsEncryptedPassword() {
        return Optional.ofNullable(this.kmsEncryptedPassword);
    }

    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Import(name="kmsEncryptionContext")
    private @Nullable Output<Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Optional<Output<Map<String,Object>>> kmsEncryptionContext() {
        return Optional.ofNullable(this.kmsEncryptionContext);
    }

    /**
     * The mongo-node count can be purchased is in range of [2, 32].
     * 
     */
    @Import(name="mongoLists")
    private @Nullable Output<List<ShardingInstanceMongoListArgs>> mongoLists;

    /**
     * @return The mongo-node count can be purchased is in range of [2, 32].
     * 
     */
    public Optional<Output<List<ShardingInstanceMongoListArgs>>> mongoLists() {
        return Optional.ofNullable(this.mongoLists);
    }

    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of DB instance. It a string of 2 to 256 characters.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     * 
     */
    @Import(name="networkType")
    private @Nullable Output<String> networkType;

    /**
     * @return The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     * 
     */
    public Optional<Output<String>> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     *   Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
     * 
     */
    @Import(name="orderType")
    private @Nullable Output<String> orderType;

    /**
     * @return The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     *   Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
     * 
     */
    public Optional<Output<String>> orderType() {
        return Optional.ofNullable(this.orderType);
    }

    /**
     * The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     * 
     */
    @Import(name="protocolType")
    private @Nullable Output<String> protocolType;

    /**
     * @return The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     * 
     */
    public Optional<Output<String>> protocolType() {
        return Optional.ofNullable(this.protocolType);
    }

    /**
     * The ID of the Resource Group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the Resource Group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Instance log backup retention days. **NOTE:** Available in 1.42.0+.
     * 
     */
    @Import(name="retentionPeriod")
    private @Nullable Output<Integer> retentionPeriod;

    /**
     * @return Instance log backup retention days. **NOTE:** Available in 1.42.0+.
     * 
     */
    public Optional<Output<Integer>> retentionPeriod() {
        return Optional.ofNullable(this.retentionPeriod);
    }

    /**
     * The Security Group ID of ECS.
     * 
     */
    @Import(name="securityGroupId")
    private @Nullable Output<String> securityGroupId;

    /**
     * @return The Security Group ID of ECS.
     * 
     */
    public Optional<Output<String>> securityGroupId() {
        return Optional.ofNullable(this.securityGroupId);
    }

    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
     * 
     */
    @Import(name="securityIpLists")
    private @Nullable Output<List<String>> securityIpLists;

    /**
     * @return List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
     * 
     */
    public Optional<Output<List<String>>> securityIpLists() {
        return Optional.ofNullable(this.securityIpLists);
    }

    /**
     * the shard-node count can be purchased is in range of [2, 32].
     * 
     */
    @Import(name="shardLists")
    private @Nullable Output<List<ShardingInstanceShardListArgs>> shardLists;

    /**
     * @return the shard-node count can be purchased is in range of [2, 32].
     * 
     */
    public Optional<Output<List<ShardingInstanceShardListArgs>>> shardLists() {
        return Optional.ofNullable(this.shardLists);
    }

    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     * 
     */
    @Import(name="storageEngine")
    private @Nullable Output<String> storageEngine;

    /**
     * @return Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     * 
     */
    public Optional<Output<String>> storageEngine() {
        return Optional.ofNullable(this.storageEngine);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
     * 
     */
    @Import(name="tdeStatus")
    private @Nullable Output<String> tdeStatus;

    /**
     * @return The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
     * 
     */
    public Optional<Output<String>> tdeStatus() {
        return Optional.ofNullable(this.tdeStatus);
    }

    /**
     * The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    /**
     * The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
     * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private ShardingInstanceState() {}

    private ShardingInstanceState(ShardingInstanceState $) {
        this.accountPassword = $.accountPassword;
        this.autoRenew = $.autoRenew;
        this.backupPeriods = $.backupPeriods;
        this.backupTime = $.backupTime;
        this.configServerLists = $.configServerLists;
        this.engineVersion = $.engineVersion;
        this.instanceChargeType = $.instanceChargeType;
        this.kmsEncryptedPassword = $.kmsEncryptedPassword;
        this.kmsEncryptionContext = $.kmsEncryptionContext;
        this.mongoLists = $.mongoLists;
        this.name = $.name;
        this.networkType = $.networkType;
        this.orderType = $.orderType;
        this.period = $.period;
        this.protocolType = $.protocolType;
        this.resourceGroupId = $.resourceGroupId;
        this.retentionPeriod = $.retentionPeriod;
        this.securityGroupId = $.securityGroupId;
        this.securityIpLists = $.securityIpLists;
        this.shardLists = $.shardLists;
        this.storageEngine = $.storageEngine;
        this.tags = $.tags;
        this.tdeStatus = $.tdeStatus;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ShardingInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ShardingInstanceState $;

        public Builder() {
            $ = new ShardingInstanceState();
        }

        public Builder(ShardingInstanceState defaults) {
            $ = new ShardingInstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountPassword Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
         * 
         * @return builder
         * 
         */
        public Builder accountPassword(@Nullable Output<String> accountPassword) {
            $.accountPassword = accountPassword;
            return this;
        }

        /**
         * @param accountPassword Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
         * 
         * @return builder
         * 
         */
        public Builder accountPassword(String accountPassword) {
            return accountPassword(Output.of(accountPassword));
        }

        /**
         * @param autoRenew Auto renew for prepaid, true of false. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<Boolean> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew Auto renew for prepaid, true of false. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(Boolean autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param backupPeriods MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
         * 
         * @return builder
         * 
         */
        public Builder backupPeriods(@Nullable Output<List<String>> backupPeriods) {
            $.backupPeriods = backupPeriods;
            return this;
        }

        /**
         * @param backupPeriods MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
         * 
         * @return builder
         * 
         */
        public Builder backupPeriods(List<String> backupPeriods) {
            return backupPeriods(Output.of(backupPeriods));
        }

        /**
         * @param backupPeriods MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
         * 
         * @return builder
         * 
         */
        public Builder backupPeriods(String... backupPeriods) {
            return backupPeriods(List.of(backupPeriods));
        }

        /**
         * @param backupTime MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
         * 
         * @return builder
         * 
         */
        public Builder backupTime(@Nullable Output<String> backupTime) {
            $.backupTime = backupTime;
            return this;
        }

        /**
         * @param backupTime MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like &#34;23:00Z-24:00Z&#34;.
         * 
         * @return builder
         * 
         */
        public Builder backupTime(String backupTime) {
            return backupTime(Output.of(backupTime));
        }

        /**
         * @param configServerLists The node information list of config server. The details see Block `config_server_list`. **NOTE:** Available in v1.140+.
         * 
         * @return builder
         * 
         */
        public Builder configServerLists(@Nullable Output<List<ShardingInstanceConfigServerListArgs>> configServerLists) {
            $.configServerLists = configServerLists;
            return this;
        }

        /**
         * @param configServerLists The node information list of config server. The details see Block `config_server_list`. **NOTE:** Available in v1.140+.
         * 
         * @return builder
         * 
         */
        public Builder configServerLists(List<ShardingInstanceConfigServerListArgs> configServerLists) {
            return configServerLists(Output.of(configServerLists));
        }

        /**
         * @param configServerLists The node information list of config server. The details see Block `config_server_list`. **NOTE:** Available in v1.140+.
         * 
         * @return builder
         * 
         */
        public Builder configServerLists(ShardingInstanceConfigServerListArgs... configServerLists) {
            return configServerLists(List.of(configServerLists));
        }

        /**
         * @param engineVersion Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(@Nullable Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param instanceChargeType Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(@Nullable Output<String> instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param instanceChargeType Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(String instanceChargeType) {
            return instanceChargeType(Output.of(instanceChargeType));
        }

        /**
         * @param kmsEncryptedPassword An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptedPassword(@Nullable Output<String> kmsEncryptedPassword) {
            $.kmsEncryptedPassword = kmsEncryptedPassword;
            return this;
        }

        /**
         * @param kmsEncryptedPassword An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptedPassword(String kmsEncryptedPassword) {
            return kmsEncryptedPassword(Output.of(kmsEncryptedPassword));
        }

        /**
         * @param kmsEncryptionContext An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptionContext(@Nullable Output<Map<String,Object>> kmsEncryptionContext) {
            $.kmsEncryptionContext = kmsEncryptionContext;
            return this;
        }

        /**
         * @param kmsEncryptionContext An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptionContext(Map<String,Object> kmsEncryptionContext) {
            return kmsEncryptionContext(Output.of(kmsEncryptionContext));
        }

        /**
         * @param mongoLists The mongo-node count can be purchased is in range of [2, 32].
         * 
         * @return builder
         * 
         */
        public Builder mongoLists(@Nullable Output<List<ShardingInstanceMongoListArgs>> mongoLists) {
            $.mongoLists = mongoLists;
            return this;
        }

        /**
         * @param mongoLists The mongo-node count can be purchased is in range of [2, 32].
         * 
         * @return builder
         * 
         */
        public Builder mongoLists(List<ShardingInstanceMongoListArgs> mongoLists) {
            return mongoLists(Output.of(mongoLists));
        }

        /**
         * @param mongoLists The mongo-node count can be purchased is in range of [2, 32].
         * 
         * @return builder
         * 
         */
        public Builder mongoLists(ShardingInstanceMongoListArgs... mongoLists) {
            return mongoLists(List.of(mongoLists));
        }

        /**
         * @param name The name of DB instance. It a string of 2 to 256 characters.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of DB instance. It a string of 2 to 256 characters.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkType The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable Output<String> networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param networkType The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(String networkType) {
            return networkType(Output.of(networkType));
        }

        /**
         * @param orderType The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
         * * UPGRADE: The specifications are upgraded.
         * * DOWNGRADE: The specifications are downgraded.
         *   Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder orderType(@Nullable Output<String> orderType) {
            $.orderType = orderType;
            return this;
        }

        /**
         * @param orderType The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
         * * UPGRADE: The specifications are upgraded.
         * * DOWNGRADE: The specifications are downgraded.
         *   Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder orderType(String orderType) {
            return orderType(Output.of(orderType));
        }

        /**
         * @param period The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param protocolType The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
         * 
         * @return builder
         * 
         */
        public Builder protocolType(@Nullable Output<String> protocolType) {
            $.protocolType = protocolType;
            return this;
        }

        /**
         * @param protocolType The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
         * 
         * @return builder
         * 
         */
        public Builder protocolType(String protocolType) {
            return protocolType(Output.of(protocolType));
        }

        /**
         * @param resourceGroupId The ID of the Resource Group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the Resource Group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param retentionPeriod Instance log backup retention days. **NOTE:** Available in 1.42.0+.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriod(@Nullable Output<Integer> retentionPeriod) {
            $.retentionPeriod = retentionPeriod;
            return this;
        }

        /**
         * @param retentionPeriod Instance log backup retention days. **NOTE:** Available in 1.42.0+.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriod(Integer retentionPeriod) {
            return retentionPeriod(Output.of(retentionPeriod));
        }

        /**
         * @param securityGroupId The Security Group ID of ECS.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(@Nullable Output<String> securityGroupId) {
            $.securityGroupId = securityGroupId;
            return this;
        }

        /**
         * @param securityGroupId The Security Group ID of ECS.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(String securityGroupId) {
            return securityGroupId(Output.of(securityGroupId));
        }

        /**
         * @param securityIpLists List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
         * 
         * @return builder
         * 
         */
        public Builder securityIpLists(@Nullable Output<List<String>> securityIpLists) {
            $.securityIpLists = securityIpLists;
            return this;
        }

        /**
         * @param securityIpLists List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
         * 
         * @return builder
         * 
         */
        public Builder securityIpLists(List<String> securityIpLists) {
            return securityIpLists(Output.of(securityIpLists));
        }

        /**
         * @param securityIpLists List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `[&#34;127.0.0.1&#34;]`.
         * 
         * @return builder
         * 
         */
        public Builder securityIpLists(String... securityIpLists) {
            return securityIpLists(List.of(securityIpLists));
        }

        /**
         * @param shardLists the shard-node count can be purchased is in range of [2, 32].
         * 
         * @return builder
         * 
         */
        public Builder shardLists(@Nullable Output<List<ShardingInstanceShardListArgs>> shardLists) {
            $.shardLists = shardLists;
            return this;
        }

        /**
         * @param shardLists the shard-node count can be purchased is in range of [2, 32].
         * 
         * @return builder
         * 
         */
        public Builder shardLists(List<ShardingInstanceShardListArgs> shardLists) {
            return shardLists(Output.of(shardLists));
        }

        /**
         * @param shardLists the shard-node count can be purchased is in range of [2, 32].
         * 
         * @return builder
         * 
         */
        public Builder shardLists(ShardingInstanceShardListArgs... shardLists) {
            return shardLists(List.of(shardLists));
        }

        /**
         * @param storageEngine Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
         * 
         * @return builder
         * 
         */
        public Builder storageEngine(@Nullable Output<String> storageEngine) {
            $.storageEngine = storageEngine;
            return this;
        }

        /**
         * @param storageEngine Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
         * 
         * @return builder
         * 
         */
        public Builder storageEngine(String storageEngine) {
            return storageEngine(Output.of(storageEngine));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tdeStatus The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
         * 
         * @return builder
         * 
         */
        public Builder tdeStatus(@Nullable Output<String> tdeStatus) {
            $.tdeStatus = tdeStatus;
            return this;
        }

        /**
         * @param tdeStatus The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
         * 
         * @return builder
         * 
         */
        public Builder tdeStatus(String tdeStatus) {
            return tdeStatus(Output.of(tdeStatus));
        }

        /**
         * @param vpcId The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The virtual switch ID to launch DB instances in one VPC.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The virtual switch ID to launch DB instances in one VPC.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
         * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
         * If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public ShardingInstanceState build() {
            return $;
        }
    }

}
