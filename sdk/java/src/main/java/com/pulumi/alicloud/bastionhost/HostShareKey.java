// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.bastionhost.HostShareKeyArgs;
import com.pulumi.alicloud.bastionhost.inputs.HostShareKeyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Bastion Host Share Key resource.
 * 
 * For information about Bastion Host Host Share Key and how to use it, see [What is Host Share Key](https://www.alibabacloud.com/help/en/bastion-host/latest/createhostsharekey).
 * 
 * &gt; **NOTE:** Available in v1.165.0+.
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
 * import com.pulumi.alicloud.bastionhost.BastionhostFunctions;
 * import com.pulumi.alicloud.bastionhost.inputs.GetInstancesArgs;
 * import com.pulumi.alicloud.bastionhost.HostShareKey;
 * import com.pulumi.alicloud.bastionhost.HostShareKeyArgs;
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
 *         final var defaultInstances = BastionhostFunctions.getInstances();
 * 
 *         var defaultHostShareKey = new HostShareKey(&#34;defaultHostShareKey&#34;, HostShareKeyArgs.builder()        
 *             .hostShareKeyName(&#34;example_name&#34;)
 *             .instanceId(defaultInstances.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()))
 *             .passPhrase(&#34;example_value&#34;)
 *             .privateKey(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Bastion Host Share Key can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:bastionhost/hostShareKey:HostShareKey example &lt;instance_id&gt;:&lt;host_share_key_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:bastionhost/hostShareKey:HostShareKey")
public class HostShareKey extends com.pulumi.resources.CustomResource {
    /**
     * The first ID of the resource.
     * 
     */
    @Export(name="hostShareKeyId", type=String.class, parameters={})
    private Output<String> hostShareKeyId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Output<String> hostShareKeyId() {
        return this.hostShareKeyId;
    }
    /**
     * The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
     * 
     */
    @Export(name="hostShareKeyName", type=String.class, parameters={})
    private Output<String> hostShareKeyName;

    /**
     * @return The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
     * 
     */
    public Output<String> hostShareKeyName() {
        return this.hostShareKeyName;
    }
    /**
     * The ID of the Bastion instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The ID of the Bastion instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The password of the private key. The value is a Base64-encoded string.
     * 
     */
    @Export(name="passPhrase", type=String.class, parameters={})
    private Output</* @Nullable */ String> passPhrase;

    /**
     * @return The password of the private key. The value is a Base64-encoded string.
     * 
     */
    public Output<Optional<String>> passPhrase() {
        return Codegen.optional(this.passPhrase);
    }
    /**
     * The private key. The value is a Base64-encoded string.
     * 
     */
    @Export(name="privateKey", type=String.class, parameters={})
    private Output<String> privateKey;

    /**
     * @return The private key. The value is a Base64-encoded string.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * The fingerprint of the private key.
     * 
     */
    @Export(name="privateKeyFingerPrint", type=String.class, parameters={})
    private Output<String> privateKeyFingerPrint;

    /**
     * @return The fingerprint of the private key.
     * 
     */
    public Output<String> privateKeyFingerPrint() {
        return this.privateKeyFingerPrint;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostShareKey(String name) {
        this(name, HostShareKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostShareKey(String name, HostShareKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostShareKey(String name, HostShareKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/hostShareKey:HostShareKey", name, args == null ? HostShareKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HostShareKey(String name, Output<String> id, @Nullable HostShareKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/hostShareKey:HostShareKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "passPhrase",
                "privateKey"
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
    public static HostShareKey get(String name, Output<String> id, @Nullable HostShareKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostShareKey(name, id, state, options);
    }
}
