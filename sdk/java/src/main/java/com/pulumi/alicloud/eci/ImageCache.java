// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eci.ImageCacheArgs;
import com.pulumi.alicloud.eci.inputs.ImageCacheState;
import com.pulumi.alicloud.eci.outputs.ImageCacheImageRegistryCredential;
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
 * An ECI Image Cache can help user to solve the time-consuming problem of image pull. For information about Alicloud ECI Image Cache and how to use it, see [What is Resource Alicloud ECI Image Cache](https://www.alibabacloud.com/help/doc-detail/146891.htm).
 * 
 * &gt; **NOTE:** Available in v1.89.0+.
 * 
 * &gt; **NOTE:** Each image cache corresponds to a snapshot, and the user does not delete the snapshot directly, otherwise the cache will fail.
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
 * import com.pulumi.alicloud.eci.ImageCache;
 * import com.pulumi.alicloud.eci.ImageCacheArgs;
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
 *         var example = new ImageCache(&#34;example&#34;, ImageCacheArgs.builder()        
 *             .eipInstanceId(&#34;eip-uf60c7cqb2pcrkgxhxxxx&#34;)
 *             .imageCacheName(&#34;tf-test&#34;)
 *             .images(&#34;registry.cn-beijing.aliyuncs.com/sceneplatform/sae-image-xxxx:latest&#34;)
 *             .securityGroupId(&#34;sg-2zeef68b66fxxxx&#34;)
 *             .vswitchId(&#34;vsw-2zef9k7ng82xxxx&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ECI Image Cache can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:eci/imageCache:ImageCache example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:eci/imageCache:ImageCache")
public class ImageCache extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the container group job that is used to create the image cache.
     * 
     */
    @Export(name="containerGroupId", type=String.class, parameters={})
    private Output<String> containerGroupId;

    /**
     * @return The ID of the container group job that is used to create the image cache.
     * 
     */
    public Output<String> containerGroupId() {
        return this.containerGroupId;
    }
    /**
     * The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
     * 
     */
    @Export(name="eipInstanceId", type=String.class, parameters={})
    private Output</* @Nullable */ String> eipInstanceId;

    /**
     * @return The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
     * 
     */
    public Output<Optional<String>> eipInstanceId() {
        return Codegen.optional(this.eipInstanceId);
    }
    /**
     * The name of the image cache.
     * 
     */
    @Export(name="imageCacheName", type=String.class, parameters={})
    private Output<String> imageCacheName;

    /**
     * @return The name of the image cache.
     * 
     */
    public Output<String> imageCacheName() {
        return this.imageCacheName;
    }
    /**
     * The size of the image cache. Default to `20`. Unit: GiB.
     * 
     */
    @Export(name="imageCacheSize", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> imageCacheSize;

    /**
     * @return The size of the image cache. Default to `20`. Unit: GiB.
     * 
     */
    public Output<Optional<Integer>> imageCacheSize() {
        return Codegen.optional(this.imageCacheSize);
    }
    /**
     * The Image Registry parameters about the image to be cached.
     * 
     */
    @Export(name="imageRegistryCredentials", type=List.class, parameters={ImageCacheImageRegistryCredential.class})
    private Output</* @Nullable */ List<ImageCacheImageRegistryCredential>> imageRegistryCredentials;

    /**
     * @return The Image Registry parameters about the image to be cached.
     * 
     */
    public Output<Optional<List<ImageCacheImageRegistryCredential>>> imageRegistryCredentials() {
        return Codegen.optional(this.imageRegistryCredentials);
    }
    /**
     * The images to be cached. The image name must be versioned.
     * 
     */
    @Export(name="images", type=List.class, parameters={String.class})
    private Output<List<String>> images;

    /**
     * @return The images to be cached. The image name must be versioned.
     * 
     */
    public Output<List<String>> images() {
        return this.images;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
     * 
     */
    @Export(name="retentionDays", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> retentionDays;

    /**
     * @return The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
     * 
     */
    public Output<Optional<Integer>> retentionDays() {
        return Codegen.optional(this.retentionDays);
    }
    /**
     * The ID of the security group. You do not need to specify the same security group as the container group.
     * 
     */
    @Export(name="securityGroupId", type=String.class, parameters={})
    private Output<String> securityGroupId;

    /**
     * @return The ID of the security group. You do not need to specify the same security group as the container group.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * The status of the image cache.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the image cache.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The zone id to cache image.
     * 
     */
    @Export(name="zoneId", type=String.class, parameters={})
    private Output</* @Nullable */ String> zoneId;

    /**
     * @return The zone id to cache image.
     * 
     */
    public Output<Optional<String>> zoneId() {
        return Codegen.optional(this.zoneId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageCache(String name) {
        this(name, ImageCacheArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageCache(String name, ImageCacheArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageCache(String name, ImageCacheArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eci/imageCache:ImageCache", name, args == null ? ImageCacheArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ImageCache(String name, Output<String> id, @Nullable ImageCacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eci/imageCache:ImageCache", name, state, makeResourceOptions(options, id));
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
    public static ImageCache get(String name, Output<String> id, @Nullable ImageCacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageCache(name, id, state, options);
    }
}
