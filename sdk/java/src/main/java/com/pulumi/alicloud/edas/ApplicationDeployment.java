// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.edas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.edas.ApplicationDeploymentArgs;
import com.pulumi.alicloud.edas.inputs.ApplicationDeploymentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Deploys applications on EDAS.
 * 
 * &gt; **NOTE:** Available in 1.82.0+
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
 * import com.pulumi.alicloud.edas.ApplicationDeployment;
 * import com.pulumi.alicloud.edas.ApplicationDeploymentArgs;
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
 *         var default_ = new ApplicationDeployment(&#34;default&#34;, ApplicationDeploymentArgs.builder()        
 *             .appId(var_.app_id())
 *             .groupId(var_.group_id())
 *             .packageVersion(var_.package_version())
 *             .warUrl(var_.war_url())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="alicloud:edas/applicationDeployment:ApplicationDeployment")
public class ApplicationDeployment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the application that you want to deploy.
     * 
     */
    @Export(name="appId", type=String.class, parameters={})
    private Output<String> appId;

    /**
     * @return The ID of the application that you want to deploy.
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
     * 
     */
    @Export(name="groupId", type=String.class, parameters={})
    private Output<String> groupId;

    /**
     * @return The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * Last package version deployed.
     * 
     */
    @Export(name="lastPackageVersion", type=String.class, parameters={})
    private Output<String> lastPackageVersion;

    /**
     * @return Last package version deployed.
     * 
     */
    public Output<String> lastPackageVersion() {
        return this.lastPackageVersion;
    }
    /**
     * The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
     * 
     */
    @Export(name="packageVersion", type=String.class, parameters={})
    private Output</* @Nullable */ String> packageVersion;

    /**
     * @return The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
     * 
     */
    public Output<Optional<String>> packageVersion() {
        return Codegen.optional(this.packageVersion);
    }
    /**
     * The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
     * 
     */
    @Export(name="warUrl", type=String.class, parameters={})
    private Output<String> warUrl;

    /**
     * @return The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
     * 
     */
    public Output<String> warUrl() {
        return this.warUrl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationDeployment(String name) {
        this(name, ApplicationDeploymentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationDeployment(String name, ApplicationDeploymentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationDeployment(String name, ApplicationDeploymentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:edas/applicationDeployment:ApplicationDeployment", name, args == null ? ApplicationDeploymentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationDeployment(String name, Output<String> id, @Nullable ApplicationDeploymentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:edas/applicationDeployment:ApplicationDeployment", name, state, makeResourceOptions(options, id));
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
    public static ApplicationDeployment get(String name, Output<String> id, @Nullable ApplicationDeploymentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationDeployment(name, id, state, options);
    }
}
