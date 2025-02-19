// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.securitycenter;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.securitycenter.ServiceLinkedRoleArgs;
import com.pulumi.alicloud.securitycenter.inputs.ServiceLinkedRoleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import javax.annotation.Nullable;

/**
 * Using this resource can create SecurityCenter service-linked role : `AliyunServiceRolePolicyForSas`.  This Role is a Resource Access Management (RAM) role, which to obtain permissions to access another Alibaba Cloud service.
 * 
 * For information about Security Center Service Role and how to use it, see [What is Security Center](https://www.alibabacloud.com/help/en/doc-detail/42302.htm).
 * 
 * &gt; **NOTE:** Available in v1.142.0+.
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
 * import com.pulumi.alicloud.securitycenter.ServiceLinkedRole;
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
 *         var serviceLinkedRole = new ServiceLinkedRole(&#34;serviceLinkedRole&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SecurityCenter service-linked roles(SLR) can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole example &lt;product_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole")
public class ServiceLinkedRole extends com.pulumi.resources.CustomResource {
    /**
     * The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
     * 
     */
    @Export(name="status", type=Boolean.class, parameters={})
    private Output<Boolean> status;

    /**
     * @return The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
     * 
     */
    public Output<Boolean> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceLinkedRole(String name) {
        this(name, ServiceLinkedRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceLinkedRole(String name, @Nullable ServiceLinkedRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceLinkedRole(String name, @Nullable ServiceLinkedRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole", name, args == null ? ServiceLinkedRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceLinkedRole(String name, Output<String> id, @Nullable ServiceLinkedRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole", name, state, makeResourceOptions(options, id));
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
    public static ServiceLinkedRole get(String name, Output<String> id, @Nullable ServiceLinkedRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceLinkedRole(name, id, state, options);
    }
}
