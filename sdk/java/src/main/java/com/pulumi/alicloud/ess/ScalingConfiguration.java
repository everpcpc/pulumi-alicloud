// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ess.ScalingConfigurationArgs;
import com.pulumi.alicloud.ess.inputs.ScalingConfigurationState;
import com.pulumi.alicloud.ess.outputs.ScalingConfigurationDataDisk;
import com.pulumi.alicloud.ess.outputs.ScalingConfigurationInstancePatternInfo;
import com.pulumi.alicloud.ess.outputs.ScalingConfigurationSpotPriceLimit;
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
 * ESS scaling configuration can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ess/scalingConfiguration:ScalingConfiguration example asg-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:ess/scalingConfiguration:ScalingConfiguration")
public class ScalingConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * Whether active current scaling configuration in the specified scaling group. Default to `false`.
     * 
     */
    @Export(name="active", type=Boolean.class, parameters={})
    private Output<Boolean> active;

    /**
     * @return Whether active current scaling configuration in the specified scaling group. Default to `false`.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * Performance mode of the t5 burstable instance. Valid values: &#39;Standard&#39;, &#39;Unlimited&#39;.
     * 
     */
    @Export(name="creditSpecification", type=String.class, parameters={})
    private Output</* @Nullable */ String> creditSpecification;

    /**
     * @return Performance mode of the t5 burstable instance. Valid values: &#39;Standard&#39;, &#39;Unlimited&#39;.
     * 
     */
    public Output<Optional<String>> creditSpecification() {
        return Codegen.optional(this.creditSpecification);
    }
    /**
     * DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
     * 
     */
    @Export(name="dataDisks", type=List.class, parameters={ScalingConfigurationDataDisk.class})
    private Output</* @Nullable */ List<ScalingConfigurationDataDisk>> dataDisks;

    /**
     * @return DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
     * 
     */
    public Output<Optional<List<ScalingConfigurationDataDisk>>> dataDisks() {
        return Codegen.optional(this.dataDisks);
    }
    /**
     * Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
     * 
     */
    @Export(name="enable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enable;

    /**
     * @return Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
     * 
     */
    public Output<Optional<Boolean>> enable() {
        return Codegen.optional(this.enable);
    }
    /**
     * The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
     * 
     */
    @Export(name="forceDelete", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceDelete;

    /**
     * @return The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
     * 
     */
    public Output<Optional<Boolean>> forceDelete() {
        return Codegen.optional(this.forceDelete);
    }
    /**
     * Hostname of an ECS instance.
     * 
     */
    @Export(name="hostName", type=String.class, parameters={})
    private Output</* @Nullable */ String> hostName;

    /**
     * @return Hostname of an ECS instance.
     * 
     */
    public Output<Optional<String>> hostName() {
        return Codegen.optional(this.hostName);
    }
    /**
     * ID of an image file, indicating the image resource selected when an instance is enabled.
     * 
     */
    @Export(name="imageId", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageId;

    /**
     * @return ID of an image file, indicating the image resource selected when an instance is enabled.
     * 
     */
    public Output<Optional<String>> imageId() {
        return Codegen.optional(this.imageId);
    }
    /**
     * Name of an image file, indicating the image resource selected when an instance is enabled.
     * 
     */
    @Export(name="imageName", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageName;

    /**
     * @return Name of an image file, indicating the image resource selected when an instance is enabled.
     * 
     */
    public Output<Optional<String>> imageName() {
        return Codegen.optional(this.imageName);
    }
    /**
     * It has been deprecated from version 1.6.0. New resource `alicloud.ess.Attachment` replaces it.
     * 
     * @deprecated
     * Field &#39;instance_ids&#39; has been deprecated from provider version 1.6.0. New resource &#39;alicloud_ess_attachment&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it. */
    @Export(name="instanceIds", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> instanceIds;

    /**
     * @return It has been deprecated from version 1.6.0. New resource `alicloud.ess.Attachment` replaces it.
     * 
     */
    public Output<Optional<List<String>>> instanceIds() {
        return Codegen.optional(this.instanceIds);
    }
    /**
     * Name of an ECS instance. Default to &#34;ESS-Instance&#34;. It is valid from version 1.7.1.
     * 
     */
    @Export(name="instanceName", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return Name of an ECS instance. Default to &#34;ESS-Instance&#34;. It is valid from version 1.7.1.
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * intelligent configuration mode. In this mode, you only need to specify the number of vCPUs, memory size, instance family, and maximum price. The system selects an instance type that is provided at the lowest price based on your configurations to create ECS instances. This mode is available only for scaling groups that reside in virtual private clouds (VPCs). This mode helps reduce the failures of scale-out activities caused by insufficient inventory of instance types
     * 
     */
    @Export(name="instancePatternInfos", type=List.class, parameters={ScalingConfigurationInstancePatternInfo.class})
    private Output</* @Nullable */ List<ScalingConfigurationInstancePatternInfo>> instancePatternInfos;

    /**
     * @return intelligent configuration mode. In this mode, you only need to specify the number of vCPUs, memory size, instance family, and maximum price. The system selects an instance type that is provided at the lowest price based on your configurations to create ECS instances. This mode is available only for scaling groups that reside in virtual private clouds (VPCs). This mode helps reduce the failures of scale-out activities caused by insufficient inventory of instance types
     * 
     */
    public Output<Optional<List<ScalingConfigurationInstancePatternInfo>>> instancePatternInfos() {
        return Codegen.optional(this.instancePatternInfos);
    }
    /**
     * Resource type of an ECS instance.
     * 
     */
    @Export(name="instanceType", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceType;

    /**
     * @return Resource type of an ECS instance.
     * 
     */
    public Output<Optional<String>> instanceType() {
        return Codegen.optional(this.instanceType);
    }
    /**
     * Resource types of an ECS instance.
     * 
     */
    @Export(name="instanceTypes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> instanceTypes;

    /**
     * @return Resource types of an ECS instance.
     * 
     */
    public Output<Optional<List<String>>> instanceTypes() {
        return Codegen.optional(this.instanceTypes);
    }
    /**
     * Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
     * 
     */
    @Export(name="internetChargeType", type=String.class, parameters={})
    private Output</* @Nullable */ String> internetChargeType;

    /**
     * @return Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
     * 
     */
    public Output<Optional<String>> internetChargeType() {
        return Codegen.optional(this.internetChargeType);
    }
    /**
     * Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
     * 
     */
    @Export(name="internetMaxBandwidthIn", type=Integer.class, parameters={})
    private Output<Integer> internetMaxBandwidthIn;

    /**
     * @return Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
     * 
     */
    public Output<Integer> internetMaxBandwidthIn() {
        return this.internetMaxBandwidthIn;
    }
    /**
     * Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
     * 
     */
    @Export(name="internetMaxBandwidthOut", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> internetMaxBandwidthOut;

    /**
     * @return Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
     * 
     */
    public Output<Optional<Integer>> internetMaxBandwidthOut() {
        return Codegen.optional(this.internetMaxBandwidthOut);
    }
    /**
     * It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
     * 
     * @deprecated
     * Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
     * 
     */
    @Deprecated /* Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template. */
    @Export(name="ioOptimized", type=String.class, parameters={})
    private Output</* @Nullable */ String> ioOptimized;

    /**
     * @return It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
     * 
     */
    public Output<Optional<String>> ioOptimized() {
        return Codegen.optional(this.ioOptimized);
    }
    /**
     * Whether to use outdated instance type. Default to false.
     * 
     */
    @Export(name="isOutdated", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isOutdated;

    /**
     * @return Whether to use outdated instance type. Default to false.
     * 
     */
    public Output<Optional<Boolean>> isOutdated() {
        return Codegen.optional(this.isOutdated);
    }
    /**
     * The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
     * 
     */
    @Export(name="keyName", type=String.class, parameters={})
    private Output</* @Nullable */ String> keyName;

    /**
     * @return The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
     * 
     */
    public Output<Optional<String>> keyName() {
        return Codegen.optional(this.keyName);
    }
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedPassword() {
        return Codegen.optional(this.kmsEncryptedPassword);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Export(name="kmsEncryptionContext", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }
    /**
     * Indicates whether to overwrite the existing data. Default to false.
     * 
     */
    @Export(name="override", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> override;

    /**
     * @return Indicates whether to overwrite the existing data. Default to false.
     * 
     */
    public Output<Optional<Boolean>> override() {
        return Codegen.optional(this.override);
    }
    /**
     * The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&amp;*-_+=\|{}[]:;&#39;&lt;&gt;,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&amp;*-_+=\|{}[]:;&#39;&lt;&gt;,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kms_encrypted_password` will be ignored. You must ensure that the selected image has a password configured.
     * 
     */
    @Export(name="passwordInherit", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> passwordInherit;

    /**
     * @return Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kms_encrypted_password` will be ignored. You must ensure that the selected image has a password configured.
     * 
     */
    public Output<Optional<Boolean>> passwordInherit() {
        return Codegen.optional(this.passwordInherit);
    }
    /**
     * ID of resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return ID of resource group.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * Instance RAM role name. The name is provided and maintained by RAM. You can use `alicloud.ram.Role` to create a new one.
     * 
     */
    @Export(name="roleName", type=String.class, parameters={})
    private Output</* @Nullable */ String> roleName;

    /**
     * @return Instance RAM role name. The name is provided and maintained by RAM. You can use `alicloud.ram.Role` to create a new one.
     * 
     */
    public Output<Optional<String>> roleName() {
        return Codegen.optional(this.roleName);
    }
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
     * 
     */
    @Export(name="scalingConfigurationName", type=String.class, parameters={})
    private Output<String> scalingConfigurationName;

    /**
     * @return Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
     * 
     */
    public Output<String> scalingConfigurationName() {
        return this.scalingConfigurationName;
    }
    /**
     * ID of the scaling group of a scaling configuration.
     * 
     */
    @Export(name="scalingGroupId", type=String.class, parameters={})
    private Output<String> scalingGroupId;

    /**
     * @return ID of the scaling group of a scaling configuration.
     * 
     */
    public Output<String> scalingGroupId() {
        return this.scalingGroupId;
    }
    /**
     * ID of the security group used to create new instance. It is conflict with `security_group_ids`.
     * 
     */
    @Export(name="securityGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityGroupId;

    /**
     * @return ID of the security group used to create new instance. It is conflict with `security_group_ids`.
     * 
     */
    public Output<Optional<String>> securityGroupId() {
        return Codegen.optional(this.securityGroupId);
    }
    /**
     * List IDs of the security group used to create new instances. It is conflict with `security_group_id`.
     * 
     */
    @Export(name="securityGroupIds", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> securityGroupIds;

    /**
     * @return List IDs of the security group used to create new instances. It is conflict with `security_group_id`.
     * 
     */
    public Output<Optional<List<String>>> securityGroupIds() {
        return Codegen.optional(this.securityGroupIds);
    }
    /**
     * Sets the maximum price hourly for instance types. See Block spotPriceLimit below for details.
     * 
     */
    @Export(name="spotPriceLimits", type=List.class, parameters={ScalingConfigurationSpotPriceLimit.class})
    private Output</* @Nullable */ List<ScalingConfigurationSpotPriceLimit>> spotPriceLimits;

    /**
     * @return Sets the maximum price hourly for instance types. See Block spotPriceLimit below for details.
     * 
     */
    public Output<Optional<List<ScalingConfigurationSpotPriceLimit>>> spotPriceLimits() {
        return Codegen.optional(this.spotPriceLimits);
    }
    /**
     * The spot strategy for a Pay-As-You-Go instance. Valid values: `NoSpot`, `SpotAsPriceGo`, `SpotWithPriceLimit`.
     * 
     */
    @Export(name="spotStrategy", type=String.class, parameters={})
    private Output</* @Nullable */ String> spotStrategy;

    /**
     * @return The spot strategy for a Pay-As-You-Go instance. Valid values: `NoSpot`, `SpotAsPriceGo`, `SpotWithPriceLimit`.
     * 
     */
    public Output<Optional<String>> spotStrategy() {
        return Codegen.optional(this.spotStrategy);
    }
    /**
     * The another scaling configuration which will be active automatically and replace current configuration when setting `active` to &#39;false&#39;. It is invalid when `active` is &#39;true&#39;.
     * 
     */
    @Export(name="substitute", type=String.class, parameters={})
    private Output<String> substitute;

    /**
     * @return The another scaling configuration which will be active automatically and replace current configuration when setting `active` to &#39;false&#39;. It is invalid when `active` is &#39;true&#39;.
     * 
     */
    public Output<String> substitute() {
        return this.substitute;
    }
    /**
     * The id of auto snapshot policy for system disk.
     * 
     */
    @Export(name="systemDiskAutoSnapshotPolicyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskAutoSnapshotPolicyId;

    /**
     * @return The id of auto snapshot policy for system disk.
     * 
     */
    public Output<Optional<String>> systemDiskAutoSnapshotPolicyId() {
        return Codegen.optional(this.systemDiskAutoSnapshotPolicyId);
    }
    /**
     * Category of the system disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloud_efficiency`.
     * 
     */
    @Export(name="systemDiskCategory", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskCategory;

    /**
     * @return Category of the system disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloud_efficiency`.
     * 
     */
    public Output<Optional<String>> systemDiskCategory() {
        return Codegen.optional(this.systemDiskCategory);
    }
    /**
     * The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    @Export(name="systemDiskDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskDescription;

    /**
     * @return The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    public Output<Optional<String>> systemDiskDescription() {
        return Codegen.optional(this.systemDiskDescription);
    }
    /**
     * Whether to encrypt the system disk.
     * 
     */
    @Export(name="systemDiskEncrypted", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> systemDiskEncrypted;

    /**
     * @return Whether to encrypt the system disk.
     * 
     */
    public Output<Optional<Boolean>> systemDiskEncrypted() {
        return Codegen.optional(this.systemDiskEncrypted);
    }
    /**
     * The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
     * 
     */
    @Export(name="systemDiskName", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskName;

    /**
     * @return The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
     * 
     */
    public Output<Optional<String>> systemDiskName() {
        return Codegen.optional(this.systemDiskName);
    }
    /**
     * The performance level of the ESSD used as the system disk.
     * 
     */
    @Export(name="systemDiskPerformanceLevel", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskPerformanceLevel;

    /**
     * @return The performance level of the ESSD used as the system disk.
     * 
     */
    public Output<Optional<String>> systemDiskPerformanceLevel() {
        return Codegen.optional(this.systemDiskPerformanceLevel);
    }
    /**
     * Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
     * 
     */
    @Export(name="systemDiskSize", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> systemDiskSize;

    /**
     * @return Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
     * 
     */
    public Output<Optional<Integer>> systemDiskSize() {
        return Codegen.optional(this.systemDiskSize);
    }
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;http://&#34;, or &#34;https://&#34; It can be a null string.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;http://&#34;, or &#34;https://&#34; It can be a null string.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
     * 
     */
    @Export(name="userData", type=String.class, parameters={})
    private Output</* @Nullable */ String> userData;

    /**
     * @return User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
     * 
     */
    public Output<Optional<String>> userData() {
        return Codegen.optional(this.userData);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScalingConfiguration(String name) {
        this(name, ScalingConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScalingConfiguration(String name, ScalingConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScalingConfiguration(String name, ScalingConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/scalingConfiguration:ScalingConfiguration", name, args == null ? ScalingConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ScalingConfiguration(String name, Output<String> id, @Nullable ScalingConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/scalingConfiguration:ScalingConfiguration", name, state, makeResourceOptions(options, id));
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
    public static ScalingConfiguration get(String name, Output<String> id, @Nullable ScalingConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScalingConfiguration(name, id, state, options);
    }
}
