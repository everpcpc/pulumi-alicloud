// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.DiskArgs;
import com.pulumi.alicloud.ecs.inputs.DiskState;
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
 * Provides a ECS disk resource.
 * 
 * &gt; **DEPRECATED:** This resource has been renamed to alicloud.ecs.EcsDisk from version 1.122.0.
 * 
 * &gt; **NOTE:** One of `size` or `snapshot_id` is required when specifying an ECS disk. If all of them be specified, `size` must more than the size of snapshot which `snapshot_id` represents. Currently, `alicloud.ecs.Disk` doesn&#39;t resize disk.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ecs.Disk;
 * import com.pulumi.alicloud.ecs.DiskArgs;
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
 *         var ecsDisk = new Disk(&#34;ecsDisk&#34;, DiskArgs.builder()        
 *             .availabilityZone(&#34;cn-beijing-b&#34;)
 *             .category(&#34;cloud_efficiency&#34;)
 *             .description(&#34;Hello ecs disk.&#34;)
 *             .encrypted(true)
 *             .kmsKeyId(&#34;2a6767f0-a16c-4679-a60f-13bf*****&#34;)
 *             .size(&#34;30&#34;)
 *             .tags(Map.of(&#34;Name&#34;, &#34;TerraformTest&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud disk can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecs/disk:Disk example d-abc12345678
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/disk:Disk")
public class Disk extends com.pulumi.resources.CustomResource {
    @Export(name="advancedFeatures", type=String.class, parameters={})
    private Output</* @Nullable */ String> advancedFeatures;

    public Output<Optional<String>> advancedFeatures() {
        return Codegen.optional(this.advancedFeatures);
    }
    /**
     * The Zone to create the disk in.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
    @Export(name="availabilityZone", type=String.class, parameters={})
    private Output<String> availabilityZone;

    /**
     * @return The Zone to create the disk in.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`. Default is `cloud_efficiency`.
     * 
     */
    @Export(name="category", type=String.class, parameters={})
    private Output</* @Nullable */ String> category;

    /**
     * @return Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`. Default is `cloud_efficiency`.
     * 
     */
    public Output<Optional<String>> category() {
        return Codegen.optional(this.category);
    }
    @Export(name="dedicatedBlockStorageClusterId", type=String.class, parameters={})
    private Output</* @Nullable */ String> dedicatedBlockStorageClusterId;

    public Output<Optional<String>> dedicatedBlockStorageClusterId() {
        return Codegen.optional(this.dedicatedBlockStorageClusterId);
    }
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     * 
     */
    @Export(name="deleteAutoSnapshot", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deleteAutoSnapshot;

    /**
     * @return Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     * 
     */
    public Output<Optional<Boolean>> deleteAutoSnapshot() {
        return Codegen.optional(this.deleteAutoSnapshot);
    }
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     * 
     */
    @Export(name="deleteWithInstance", type=Boolean.class, parameters={})
    private Output<Boolean> deleteWithInstance;

    /**
     * @return Indicates whether the disk is released together with the instance: Default value: false.
     * 
     */
    public Output<Boolean> deleteWithInstance() {
        return this.deleteWithInstance;
    }
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="diskName", type=String.class, parameters={})
    private Output<String> diskName;

    public Output<String> diskName() {
        return this.diskName;
    }
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     * 
     */
    @Export(name="enableAutoSnapshot", type=Boolean.class, parameters={})
    private Output<Boolean> enableAutoSnapshot;

    /**
     * @return Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     * 
     */
    public Output<Boolean> enableAutoSnapshot() {
        return this.enableAutoSnapshot;
    }
    @Export(name="encryptAlgorithm", type=String.class, parameters={})
    private Output</* @Nullable */ String> encryptAlgorithm;

    public Output<Optional<String>> encryptAlgorithm() {
        return Codegen.optional(this.encryptAlgorithm);
    }
    /**
     * If true, the disk will be encrypted, conflict with `snapshot_id`.
     * 
     */
    @Export(name="encrypted", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> encrypted;

    /**
     * @return If true, the disk will be encrypted, conflict with `snapshot_id`.
     * 
     */
    public Output<Optional<Boolean>> encrypted() {
        return Codegen.optional(this.encrypted);
    }
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     * 
     */
    @Export(name="kmsKeyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
     * 
     */
    @Export(name="performanceLevel", type=String.class, parameters={})
    private Output<String> performanceLevel;

    /**
     * @return Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
     * 
     */
    public Output<String> performanceLevel() {
        return this.performanceLevel;
    }
    /**
     * The Id of resource group which the disk belongs.
     * &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The Id of resource group which the disk belongs.
     * &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     * 
     */
    @Export(name="size", type=Integer.class, parameters={})
    private Output<Integer> size;

    /**
     * @return The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     * 
     */
    @Export(name="snapshotId", type=String.class, parameters={})
    private Output</* @Nullable */ String> snapshotId;

    /**
     * @return A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     * 
     */
    public Output<Optional<String>> snapshotId() {
        return Codegen.optional(this.snapshotId);
    }
    /**
     * The disk status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The disk status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    @Export(name="storageSetId", type=String.class, parameters={})
    private Output</* @Nullable */ String> storageSetId;

    public Output<Optional<String>> storageSetId() {
        return Codegen.optional(this.storageSetId);
    }
    @Export(name="storageSetPartitionNumber", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> storageSetPartitionNumber;

    public Output<Optional<Integer>> storageSetPartitionNumber() {
        return Codegen.optional(this.storageSetPartitionNumber);
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
    @Export(name="type", type=String.class, parameters={})
    private Output</* @Nullable */ String> type;

    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    @Export(name="zoneId", type=String.class, parameters={})
    private Output<String> zoneId;

    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Disk(String name) {
        this(name, DiskArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Disk(String name, @Nullable DiskArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Disk(String name, @Nullable DiskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/disk:Disk", name, args == null ? DiskArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Disk(String name, Output<String> id, @Nullable DiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/disk:Disk", name, state, makeResourceOptions(options, id));
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
    public static Disk get(String name, Output<String> id, @Nullable DiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Disk(name, id, state, options);
    }
}
