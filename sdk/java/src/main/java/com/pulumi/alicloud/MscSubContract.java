// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud;

import com.pulumi.alicloud.MscSubContractArgs;
import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.inputs.MscSubContractState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Msc Sub Contact resource.
 * 
 * &gt; **NOTE:** Available in v1.132.0+.
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
 * import com.pulumi.alicloud.MscSubContract;
 * import com.pulumi.alicloud.MscSubContractArgs;
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
 *         var default_ = new MscSubContract(&#34;default&#34;, MscSubContractArgs.builder()        
 *             .contactName(example_value)
 *             .position(&#34;CEO&#34;)
 *             .email(&#34;123@163.com&#34;)
 *             .mobile(&#34;153xxxxx906&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Msc Sub Contact can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:index/mscSubContract:MscSubContract example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:index/mscSubContract:MscSubContract")
public class MscSubContract extends com.pulumi.resources.CustomResource {
    /**
     * The User&#39;s Contact Name. **Note:** The name must be 2 to 12 characters in length.
     * 
     */
    @Export(name="contactName", type=String.class, parameters={})
    private Output<String> contactName;

    /**
     * @return The User&#39;s Contact Name. **Note:** The name must be 2 to 12 characters in length.
     * 
     */
    public Output<String> contactName() {
        return this.contactName;
    }
    /**
     * The User&#39;s Contact Email Address.
     * 
     */
    @Export(name="email", type=String.class, parameters={})
    private Output<String> email;

    /**
     * @return The User&#39;s Contact Email Address.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * The User&#39;s Telephone.
     * 
     */
    @Export(name="mobile", type=String.class, parameters={})
    private Output<String> mobile;

    /**
     * @return The User&#39;s Telephone.
     * 
     */
    public Output<String> mobile() {
        return this.mobile;
    }
    /**
     * The User&#39;s Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
     * 
     * &gt; **NOTE:** When the user creates a contact, the user should use `alicloud.getMscSubContactVerificationMessage` to receive the verification message and confirm it.
     * 
     */
    @Export(name="position", type=String.class, parameters={})
    private Output<String> position;

    /**
     * @return The User&#39;s Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
     * 
     * &gt; **NOTE:** When the user creates a contact, the user should use `alicloud.getMscSubContactVerificationMessage` to receive the verification message and confirm it.
     * 
     */
    public Output<String> position() {
        return this.position;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MscSubContract(String name) {
        this(name, MscSubContractArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MscSubContract(String name, MscSubContractArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MscSubContract(String name, MscSubContractArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:index/mscSubContract:MscSubContract", name, args == null ? MscSubContractArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MscSubContract(String name, Output<String> id, @Nullable MscSubContractState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:index/mscSubContract:MscSubContract", name, state, makeResourceOptions(options, id));
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
    public static MscSubContract get(String name, Output<String> id, @Nullable MscSubContractState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MscSubContract(name, id, state, options);
    }
}
