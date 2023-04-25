// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.tag;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.tag.PolicyArgs;
import com.pulumi.alicloud.tag.inputs.PolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Tag Policy resource.
 * 
 * For information about Tag Policy and how to use it,
 * see [What is Policy](https://www.alibabacloud.com/help/en/resource-management/latest/create-policy).
 * 
 * &gt; **NOTE:** Available in v1.203.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.tag.Policy;
 * import com.pulumi.alicloud.tag.PolicyArgs;
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
 *         var example = new Policy(&#34;example&#34;, PolicyArgs.builder()        
 *             .policyDesc(&#34;testDesc&#34;)
 *             .policyDocument(&#34;&#34;&#34;
 * 		{&#34;tags&#34;:{&#34;CostCenter&#34;:{&#34;tag_value&#34;:{&#34;@@assign&#34;:[&#34;Beijing&#34;,&#34;Shanghai&#34;]},&#34;tag_key&#34;:{&#34;@@assign&#34;:&#34;CostCenter&#34;}}}}
 *     
 *             &#34;&#34;&#34;)
 *             .policyName(&#34;testName&#34;)
 *             .userType(&#34;USER&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Tag Policy can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:tag/policy:Policy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:tag/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * The content of the policy.
     * 
     */
    @Export(name="policyContent", type=String.class, parameters={})
    private Output<String> policyContent;

    /**
     * @return The content of the policy.
     * 
     */
    public Output<String> policyContent() {
        return this.policyContent;
    }
    /**
     * The description of the policy. The description must be 1 to 512 characters in length.
     * 
     */
    @Export(name="policyDesc", type=String.class, parameters={})
    private Output</* @Nullable */ String> policyDesc;

    /**
     * @return The description of the policy. The description must be 1 to 512 characters in length.
     * 
     */
    public Output<Optional<String>> policyDesc() {
        return Codegen.optional(this.policyDesc);
    }
    /**
     * The name of the policy. name must be 1 to 128 characters in length and can contain letters,
     * digits, and hyphens (-).
     * 
     */
    @Export(name="policyName", type=String.class, parameters={})
    private Output<String> policyName;

    /**
     * @return The name of the policy. name must be 1 to 128 characters in length and can contain letters,
     * digits, and hyphens (-).
     * 
     */
    public Output<String> policyName() {
        return this.policyName;
    }
    /**
     * The type of the tag policy. Valid values: `USER`, `RD`.
     * 
     */
    @Export(name="userType", type=String.class, parameters={})
    private Output</* @Nullable */ String> userType;

    /**
     * @return The type of the tag policy. Valid values: `USER`, `RD`.
     * 
     */
    public Output<Optional<String>> userType() {
        return Codegen.optional(this.userType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:tag/policy:Policy", name, args == null ? PolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Policy(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:tag/policy:Policy", name, state, makeResourceOptions(options, id));
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
    public static Policy get(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
