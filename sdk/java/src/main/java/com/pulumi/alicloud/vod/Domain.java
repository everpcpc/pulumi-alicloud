// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vod;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vod.DomainArgs;
import com.pulumi.alicloud.vod.inputs.DomainState;
import com.pulumi.alicloud.vod.outputs.DomainSource;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VOD Domain resource.
 * 
 * For information about VOD Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/product/29932.html).
 * 
 * &gt; **NOTE:** Available in v1.136.0+.
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
 * import com.pulumi.alicloud.vod.Domain;
 * import com.pulumi.alicloud.vod.DomainArgs;
 * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
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
 *         var default_ = new Domain(&#34;default&#34;, DomainArgs.builder()        
 *             .domainName(&#34;your_domain_name&#34;)
 *             .scope(&#34;domestic&#34;)
 *             .sources(DomainSourceArgs.builder()
 *                 .sourceContent(&#34;your_source_content&#34;)
 *                 .sourcePort(&#34;80&#34;)
 *                 .sourceType(&#34;domain&#34;)
 *                 .build())
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
 *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VOD Domain can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vod/domain:Domain example &lt;domain_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vod/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    @Export(name="certName", type=String.class, parameters={})
    private Output<String> certName;

    /**
     * @return The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    public Output<String> certName() {
        return this.certName;
    }
    /**
     * The URL that is used for health checks.
     * 
     */
    @Export(name="checkUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> checkUrl;

    /**
     * @return The URL that is used for health checks.
     * 
     */
    public Output<Optional<String>> checkUrl() {
        return Codegen.optional(this.checkUrl);
    }
    /**
     * The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
     * 
     */
    @Export(name="cname", type=String.class, parameters={})
    private Output<String> cname;

    /**
     * @return The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
     * 
     */
    public Output<String> cname() {
        return this.cname;
    }
    /**
     * The description of the domain name for CDN.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return The description of the domain name for CDN.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     * 
     */
    @Export(name="domainName", type=String.class, parameters={})
    private Output<String> domainName;

    /**
     * @return The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="gmtCreated", type=String.class, parameters={})
    private Output<String> gmtCreated;

    /**
     * @return The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> gmtCreated() {
        return this.gmtCreated;
    }
    /**
     * The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="gmtModified", type=String.class, parameters={})
    private Output<String> gmtModified;

    /**
     * @return The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> gmtModified() {
        return this.gmtModified;
    }
    /**
     * This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     * 
     */
    @Export(name="scope", type=String.class, parameters={})
    private Output</* @Nullable */ String> scope;

    /**
     * @return This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
    }
    /**
     * The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     * 
     */
    @Export(name="sources", type=List.class, parameters={DomainSource.class})
    private Output<List<DomainSource>> sources;

    /**
     * @return The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     * 
     */
    public Output<List<DomainSource>> sources() {
        return this.sources;
    }
    /**
     * Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
     * 
     */
    @Export(name="sslProtocol", type=String.class, parameters={})
    private Output<String> sslProtocol;

    /**
     * @return Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
     * 
     */
    public Output<String> sslProtocol() {
        return this.sslProtocol;
    }
    /**
     * The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    @Export(name="sslPub", type=String.class, parameters={})
    private Output<String> sslPub;

    /**
     * @return The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    public Output<String> sslPub() {
        return this.sslPub;
    }
    /**
     * The status of the domain name for CDN. Valid values:
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the domain name for CDN. Valid values:
     * 
     */
    public Output<String> status() {
        return this.status;
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
     * The top-level domain name.
     * 
     */
    @Export(name="topLevelDomain", type=String.class, parameters={})
    private Output</* @Nullable */ String> topLevelDomain;

    /**
     * @return The top-level domain name.
     * 
     */
    public Output<Optional<String>> topLevelDomain() {
        return Codegen.optional(this.topLevelDomain);
    }
    /**
     * The weight of the origin server.
     * 
     */
    @Export(name="weight", type=String.class, parameters={})
    private Output<String> weight;

    /**
     * @return The weight of the origin server.
     * 
     */
    public Output<String> weight() {
        return this.weight;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vod/domain:Domain", name, args == null ? DomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Domain(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vod/domain:Domain", name, state, makeResourceOptions(options, id));
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
    public static Domain get(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}