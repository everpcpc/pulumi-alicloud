// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dcdn.DomainArgs;
import com.pulumi.alicloud.dcdn.inputs.DomainState;
import com.pulumi.alicloud.dcdn.outputs.DomainSource;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * You can use DCDN to improve the overall performance of your website and accelerate content delivery to improve user experience. For information about Alicloud DCDN Domain and how to use it, see [What is Resource Alicloud DCDN Domain](https://www.alibabacloud.com/help/en/doc-detail/130628.htm).
 * 
 * &gt; **NOTE:** Available in v1.94.0+.
 * 
 * &gt; **NOTE:** You must activate the Dynamic Route for CDN (DCDN) service before you create an accelerated domain.
 * 
 * &gt; **NOTE:** Make sure that you have obtained an Internet content provider (ICP) filling for the accelerated domain.
 * 
 * &gt; **NOTE:** If the origin content is not saved on Alibaba Cloud, the content must be reviewed by Alibaba Cloud. The review will be completed by the next working day after you submit the application.
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
 * import com.pulumi.alicloud.dcdn.Domain;
 * import com.pulumi.alicloud.dcdn.DomainArgs;
 * import com.pulumi.alicloud.dcdn.inputs.DomainSourceArgs;
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
 *         var example = new Domain(&#34;example&#34;, DomainArgs.builder()        
 *             .domainName(&#34;example.com&#34;)
 *             .scope(&#34;overseas&#34;)
 *             .sources(DomainSourceArgs.builder()
 *                 .content(&#34;1.1.1.1&#34;)
 *                 .port(&#34;80&#34;)
 *                 .priority(&#34;20&#34;)
 *                 .type(&#34;ipaddr&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DCDN Domain can be imported using the id or DCDN Domain name, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dcdn/domain:Domain example example.com
 * ```
 * 
 */
@ResourceType(type="alicloud:dcdn/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * Indicates the name of the certificate if the HTTPS protocol is enabled.
     * 
     */
    @Export(name="certName", type=String.class, parameters={})
    private Output<String> certName;

    /**
     * @return Indicates the name of the certificate if the HTTPS protocol is enabled.
     * 
     */
    public Output<String> certName() {
        return this.certName;
    }
    /**
     * The type of the certificate. Valid values:
     * `free`: a free certificate.
     * `cas`: a certificate purchased from Alibaba Cloud SSL Certificates Service.
     * `upload`: a user uploaded certificate.
     * 
     */
    @Export(name="certType", type=String.class, parameters={})
    private Output</* @Nullable */ String> certType;

    /**
     * @return The type of the certificate. Valid values:
     * `free`: a free certificate.
     * `cas`: a certificate purchased from Alibaba Cloud SSL Certificates Service.
     * `upload`: a user uploaded certificate.
     * 
     */
    public Output<Optional<String>> certType() {
        return Codegen.optional(this.certType);
    }
    /**
     * The URL that is used to test the accessibility of the origin.
     * 
     */
    @Export(name="checkUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> checkUrl;

    /**
     * @return The URL that is used to test the accessibility of the origin.
     * 
     */
    public Output<Optional<String>> checkUrl() {
        return Codegen.optional(this.checkUrl);
    }
    /**
     * (Available in 1.198.0+)- The canonical name (CNAME) of the accelerated domain.
     * 
     */
    @Export(name="cname", type=String.class, parameters={})
    private Output<String> cname;

    /**
     * @return (Available in 1.198.0+)- The canonical name (CNAME) of the accelerated domain.
     * 
     */
    public Output<String> cname() {
        return this.cname;
    }
    /**
     * The name of the accelerated domain.
     * 
     */
    @Export(name="domainName", type=String.class, parameters={})
    private Output<String> domainName;

    /**
     * @return The name of the accelerated domain.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * Specifies whether to check the certificate name for duplicates. If you set the value to 1, the system does not perform the check and overwrites the information of the existing certificate with the same name.
     * 
     */
    @Export(name="forceSet", type=String.class, parameters={})
    private Output</* @Nullable */ String> forceSet;

    /**
     * @return Specifies whether to check the certificate name for duplicates. If you set the value to 1, the system does not perform the check and overwrites the information of the existing certificate with the same name.
     * 
     */
    public Output<Optional<String>> forceSet() {
        return Codegen.optional(this.forceSet);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The acceleration region.
     * 
     */
    @Export(name="scope", type=String.class, parameters={})
    private Output</* @Nullable */ String> scope;

    /**
     * @return The acceleration region.
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
    }
    /**
     * The top-level domain name.
     * 
     */
    @Export(name="securityToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityToken;

    /**
     * @return The top-level domain name.
     * 
     */
    public Output<Optional<String>> securityToken() {
        return Codegen.optional(this.securityToken);
    }
    /**
     * The origin information.
     * 
     */
    @Export(name="sources", type=List.class, parameters={DomainSource.class})
    private Output<List<DomainSource>> sources;

    /**
     * @return The origin information.
     * 
     */
    public Output<List<DomainSource>> sources() {
        return this.sources;
    }
    /**
     * The private key. Specify this parameter only if you enable the SSL certificate.
     * 
     */
    @Export(name="sslPri", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslPri;

    /**
     * @return The private key. Specify this parameter only if you enable the SSL certificate.
     * 
     */
    public Output<Optional<String>> sslPri() {
        return Codegen.optional(this.sslPri);
    }
    /**
     * Indicates whether the SSL certificate is enabled. Valid values: `on` enabled, `off` disabled.
     * 
     */
    @Export(name="sslProtocol", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslProtocol;

    /**
     * @return Indicates whether the SSL certificate is enabled. Valid values: `on` enabled, `off` disabled.
     * 
     */
    public Output<Optional<String>> sslProtocol() {
        return Codegen.optional(this.sslProtocol);
    }
    /**
     * Indicates the public key of the certificate if the HTTPS protocol is enabled.
     * 
     */
    @Export(name="sslPub", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslPub;

    /**
     * @return Indicates the public key of the certificate if the HTTPS protocol is enabled.
     * 
     */
    public Output<Optional<String>> sslPub() {
        return Codegen.optional(this.sslPub);
    }
    /**
     * The status of DCDN Domain. Valid values: `online`, `offline`. Default to `online`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output</* @Nullable */ String> status;

    /**
     * @return The status of DCDN Domain. Valid values: `online`, `offline`. Default to `online`.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
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
        super("alicloud:dcdn/domain:Domain", name, args == null ? DomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Domain(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/domain:Domain", name, state, makeResourceOptions(options, id));
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
