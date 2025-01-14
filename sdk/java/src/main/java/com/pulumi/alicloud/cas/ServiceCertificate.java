// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cas.ServiceCertificateArgs;
import com.pulumi.alicloud.cas.inputs.ServiceCertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a SSL Certificates Certificate resource.
 * 
 * For information about SSL Certificates Certificate and how to use it, see [What is Certificate](https://www.alibabacloud.com/help/product/28533.html).
 * 
 * &gt; **NOTE:** Available in v1.129.0+.
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
 * import com.pulumi.alicloud.cas.ServiceCertificate;
 * import com.pulumi.alicloud.cas.ServiceCertificateArgs;
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
 *         var example = new ServiceCertificate(&#34;example&#34;, ServiceCertificateArgs.builder()        
 *             .certificateName(&#34;test&#34;)
 *             .cert(Files.readString(Paths.get(String.format(&#34;%s/test.crt&#34;, path.module()))))
 *             .key(Files.readString(Paths.get(String.format(&#34;%s/test.key&#34;, path.module()))))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SSL Certificates Certificate can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cas/serviceCertificate:ServiceCertificate example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cas/serviceCertificate:ServiceCertificate")
public class ServiceCertificate extends com.pulumi.resources.CustomResource {
    /**
     * Cert of the Certificate in which the Certificate will add.
     * 
     */
    @Export(name="cert", type=String.class, parameters={})
    private Output<String> cert;

    /**
     * @return Cert of the Certificate in which the Certificate will add.
     * 
     */
    public Output<String> cert() {
        return this.cert;
    }
    /**
     * Name of the Certificate.
     * This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;,
     * and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time.
     * Suffix .sh and .tel are not supported.
     * **NOTE:** One of `certificate_name` and `name` must be specified.
     * 
     */
    @Export(name="certificateName", type=String.class, parameters={})
    private Output<String> certificateName;

    /**
     * @return Name of the Certificate.
     * This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;,
     * and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time.
     * Suffix .sh and .tel are not supported.
     * **NOTE:** One of `certificate_name` and `name` must be specified.
     * 
     */
    public Output<String> certificateName() {
        return this.certificateName;
    }
    /**
     * Key of the Certificate in which the Certificate will add.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return Key of the Certificate in which the Certificate will add.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The lang.
     * 
     */
    @Export(name="lang", type=String.class, parameters={})
    private Output</* @Nullable */ String> lang;

    /**
     * @return The lang.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * It has been deprecated from version 1.129.0 and using `certificate_name` instead.
     * 
     * @deprecated
     * attribute &#39;name&#39; has been deprecated from provider version 1.129.0 and it will be remove in the future version. Please use the new attribute &#39;certificate_name&#39; instead.
     * 
     */
    @Deprecated /* attribute 'name' has been deprecated from provider version 1.129.0 and it will be remove in the future version. Please use the new attribute 'certificate_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return It has been deprecated from version 1.129.0 and using `certificate_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceCertificate(String name) {
        this(name, ServiceCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceCertificate(String name, ServiceCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceCertificate(String name, ServiceCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cas/serviceCertificate:ServiceCertificate", name, args == null ? ServiceCertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceCertificate(String name, Output<String> id, @Nullable ServiceCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cas/serviceCertificate:ServiceCertificate", name, state, makeResourceOptions(options, id));
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
    public static ServiceCertificate get(String name, Output<String> id, @Nullable ServiceCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceCertificate(name, id, state, options);
    }
}
