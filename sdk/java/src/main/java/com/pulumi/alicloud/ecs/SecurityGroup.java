// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.SecurityGroupArgs;
import com.pulumi.alicloud.ecs.inputs.SecurityGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Security Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecs/securityGroup:SecurityGroup example sg-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/securityGroup:SecurityGroup")
public class SecurityGroup extends com.pulumi.resources.CustomResource {
    /**
     * The security group description. Defaults to null.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The security group description. Defaults to null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Field &#39;inner_access&#39; has been deprecated from provider version 1.55.3. Use &#39;inner_access_policy&#39; replaces it.
     * 
     * @deprecated
     * Field &#39;inner_access&#39; has been deprecated from provider version 1.55.3. Use &#39;inner_access_policy&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it. */
    @Export(name="innerAccess", type=Boolean.class, parameters={})
    private Output<Boolean> innerAccess;

    /**
     * @return Field &#39;inner_access&#39; has been deprecated from provider version 1.55.3. Use &#39;inner_access_policy&#39; replaces it.
     * 
     */
    public Output<Boolean> innerAccess() {
        return this.innerAccess;
    }
    /**
     * Whether to allow both machines to access each other on all ports in the same security group. Valid values: [&#34;Accept&#34;, &#34;Drop&#34;]
     * 
     */
    @Export(name="innerAccessPolicy", type=String.class, parameters={})
    private Output<String> innerAccessPolicy;

    /**
     * @return Whether to allow both machines to access each other on all ports in the same security group. Valid values: [&#34;Accept&#34;, &#34;Drop&#34;]
     * 
     */
    public Output<String> innerAccessPolicy() {
        return this.innerAccessPolicy;
    }
    /**
     * The name of the security group. Defaults to null.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the security group. Defaults to null.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Id of resource group which the security_group belongs.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The Id of resource group which the security_group belongs.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The type of the security group. Valid values:
     * `normal`: basic security group.
     * `enterprise`: advanced security group For more information.
     * 
     */
    @Export(name="securityGroupType", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityGroupType;

    /**
     * @return The type of the security group. Valid values:
     * `normal`: basic security group.
     * `enterprise`: advanced security group For more information.
     * 
     */
    public Output<Optional<String>> securityGroupType() {
        return Codegen.optional(this.securityGroupType);
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
     * The VPC ID.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The VPC ID.
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityGroup(String name) {
        this(name, SecurityGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityGroup(String name, @Nullable SecurityGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityGroup(String name, @Nullable SecurityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/securityGroup:SecurityGroup", name, args == null ? SecurityGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecurityGroup(String name, Output<String> id, @Nullable SecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/securityGroup:SecurityGroup", name, state, makeResourceOptions(options, id));
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
    public static SecurityGroup get(String name, Output<String> id, @Nullable SecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityGroup(name, id, state, options);
    }
}