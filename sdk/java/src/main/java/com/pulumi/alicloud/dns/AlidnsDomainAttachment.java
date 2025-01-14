// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.AlidnsDomainAttachmentArgs;
import com.pulumi.alicloud.dns.inputs.AlidnsDomainAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides bind the domain name to the Alidns instance resource.
 * 
 * &gt; **NOTE:** Available in v1.99.0+.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.dns.AlidnsDomainAttachment;
 * import com.pulumi.alicloud.dns.AlidnsDomainAttachmentArgs;
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
 *         var dns = new AlidnsDomainAttachment(&#34;dns&#34;, AlidnsDomainAttachmentArgs.builder()        
 *             .domainNames(            
 *                 &#34;test111.abc&#34;,
 *                 &#34;test222.abc&#34;)
 *             .instanceId(&#34;dns-cn-mp91lyq9xxxx&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DNS domain attachment can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment example dns-cn-v0h1ldjhxxx
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment")
public class AlidnsDomainAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The domain names bound to the DNS instance.
     * 
     */
    @Export(name="domainNames", type=List.class, parameters={String.class})
    private Output<List<String>> domainNames;

    /**
     * @return The domain names bound to the DNS instance.
     * 
     */
    public Output<List<String>> domainNames() {
        return this.domainNames;
    }
    /**
     * The id of the DNS instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The id of the DNS instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlidnsDomainAttachment(String name) {
        this(name, AlidnsDomainAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlidnsDomainAttachment(String name, AlidnsDomainAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlidnsDomainAttachment(String name, AlidnsDomainAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment", name, args == null ? AlidnsDomainAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AlidnsDomainAttachment(String name, Output<String> id, @Nullable AlidnsDomainAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment", name, state, makeResourceOptions(options, id));
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
    public static AlidnsDomainAttachment get(String name, Output<String> id, @Nullable AlidnsDomainAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlidnsDomainAttachment(name, id, state, options);
    }
}
