// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudstoragegateway.GatewayBlockVolumeArgs;
import com.pulumi.alicloud.cloudstoragegateway.inputs.GatewayBlockVolumeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Storage Gateway Gateway Block Volume resource.
 * 
 * For information about Cloud Storage Gateway Gateway Block Volume and how to use it, see [What is Gateway Block Volume](https://www.alibabacloud.com/help/en/doc-detail/53972.htm).
 * 
 * &gt; **NOTE:** Available in v1.144.0+.
 * 
 * ## Import
 * 
 * Cloud Storage Gateway Gateway Block Volume can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume example &lt;gateway_id&gt;:&lt;index_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume")
public class GatewayBlockVolume extends com.pulumi.resources.CustomResource {
    /**
     * The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
     * 
     */
    @Export(name="cacheMode", type=String.class, parameters={})
    private Output<String> cacheMode;

    /**
     * @return The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
     * 
     */
    public Output<String> cacheMode() {
        return this.cacheMode;
    }
    /**
     * Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
     * 
     */
    @Export(name="chapEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> chapEnabled;

    /**
     * @return Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
     * 
     */
    public Output<Boolean> chapEnabled() {
        return this.chapEnabled;
    }
    /**
     * The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
     * 
     */
    @Export(name="chapInPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> chapInPassword;

    /**
     * @return The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
     * 
     */
    public Output<Optional<String>> chapInPassword() {
        return Codegen.optional(this.chapInPassword);
    }
    /**
     * The Inbound CHAP user. The `chap_in_user` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
     * 
     */
    @Export(name="chapInUser", type=String.class, parameters={})
    private Output</* @Nullable */ String> chapInUser;

    /**
     * @return The Inbound CHAP user. The `chap_in_user` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chap_enabled` is  `true` is,The `chap_in_password` is valid.
     * 
     */
    public Output<Optional<String>> chapInUser() {
        return Codegen.optional(this.chapInUser);
    }
    /**
     * The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
     * 
     */
    @Export(name="chunkSize", type=Integer.class, parameters={})
    private Output<Integer> chunkSize;

    /**
     * @return The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
     * 
     */
    public Output<Integer> chunkSize() {
        return this.chunkSize;
    }
    /**
     * The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
     * 
     */
    @Export(name="gatewayBlockVolumeName", type=String.class, parameters={})
    private Output<String> gatewayBlockVolumeName;

    /**
     * @return The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
     * 
     */
    public Output<String> gatewayBlockVolumeName() {
        return this.gatewayBlockVolumeName;
    }
    /**
     * The Gateway ID.
     * 
     */
    @Export(name="gatewayId", type=String.class, parameters={})
    private Output<String> gatewayId;

    /**
     * @return The Gateway ID.
     * 
     */
    public Output<String> gatewayId() {
        return this.gatewayId;
    }
    /**
     * The ID of the index.
     * 
     */
    @Export(name="indexId", type=String.class, parameters={})
    private Output<String> indexId;

    /**
     * @return The ID of the index.
     * 
     */
    public Output<String> indexId() {
        return this.indexId;
    }
    /**
     * Whether to delete the source data. Default value `true`. **NOTE:** When `is_source_deletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
     * 
     */
    @Export(name="isSourceDeletion", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isSourceDeletion;

    /**
     * @return Whether to delete the source data. Default value `true`. **NOTE:** When `is_source_deletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
     * 
     */
    public Output<Optional<Boolean>> isSourceDeletion() {
        return Codegen.optional(this.isSourceDeletion);
    }
    /**
     * The Cache disk to local path. **NOTE:**  When the `cache_mode` is  `Cache` is,The `chap_in_password` is valid.
     * 
     */
    @Export(name="localPath", type=String.class, parameters={})
    private Output</* @Nullable */ String> localPath;

    /**
     * @return The Cache disk to local path. **NOTE:**  When the `cache_mode` is  `Cache` is,The `chap_in_password` is valid.
     * 
     */
    public Output<Optional<String>> localPath() {
        return Codegen.optional(this.localPath);
    }
    /**
     * The name of the OSS Bucket.
     * 
     */
    @Export(name="ossBucketName", type=String.class, parameters={})
    private Output<String> ossBucketName;

    /**
     * @return The name of the OSS Bucket.
     * 
     */
    public Output<String> ossBucketName() {
        return this.ossBucketName;
    }
    /**
     * Whether to enable SSL access your OSS Buckets. Default value: `true`.
     * 
     */
    @Export(name="ossBucketSsl", type=Boolean.class, parameters={})
    private Output<Boolean> ossBucketSsl;

    /**
     * @return Whether to enable SSL access your OSS Buckets. Default value: `true`.
     * 
     */
    public Output<Boolean> ossBucketSsl() {
        return this.ossBucketSsl;
    }
    /**
     * The endpoint of the OSS Bucket.
     * 
     */
    @Export(name="ossEndpoint", type=String.class, parameters={})
    private Output<String> ossEndpoint;

    /**
     * @return The endpoint of the OSS Bucket.
     * 
     */
    public Output<String> ossEndpoint() {
        return this.ossEndpoint;
    }
    /**
     * The Protocol. Valid values: `iSCSI`.
     * 
     */
    @Export(name="protocol", type=String.class, parameters={})
    private Output<String> protocol;

    /**
     * @return The Protocol. Valid values: `iSCSI`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The recovery.
     * 
     */
    @Export(name="recovery", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> recovery;

    /**
     * @return The recovery.
     * 
     */
    public Output<Optional<Boolean>> recovery() {
        return Codegen.optional(this.recovery);
    }
    /**
     * The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
     * 
     */
    @Export(name="size", type=Integer.class, parameters={})
    private Output<Integer> size;

    /**
     * @return The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The status of volume. Valid values:
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of volume. Valid values:
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GatewayBlockVolume(String name) {
        this(name, GatewayBlockVolumeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GatewayBlockVolume(String name, GatewayBlockVolumeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GatewayBlockVolume(String name, GatewayBlockVolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume", name, args == null ? GatewayBlockVolumeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GatewayBlockVolume(String name, Output<String> id, @Nullable GatewayBlockVolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume", name, state, makeResourceOptions(options, id));
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
    public static GatewayBlockVolume get(String name, Output<String> id, @Nullable GatewayBlockVolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GatewayBlockVolume(name, id, state, options);
    }
}
