// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudsso;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudsso.AccessManagementArgs;
import com.pulumi.alicloud.cloudsso.inputs.AccessManagementState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud SSO Access Assignment resource.
 * 
 * For information about Cloud SSO Access Assignment and how to use it, see [What is Access Assignment](https://www.alibabacloud.com/help/en/doc-detail/265996.htm).
 * 
 * &gt; **NOTE:** When you configure access assignment for the first time, access configuration will be automatically deployed.
 * 
 * &gt; **NOTE:** Available in v1.145.0+.
 * 
 * &gt; **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
 * 
 * ## Import
 * 
 * Cloud SSO Access Assignment can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cloudsso/accessManagement:AccessManagement example &lt;directory_id&gt;:&lt;access_configuration_id&gt;:&lt;target_type&gt;:&lt;target_id&gt;:&lt;principal_type&gt;:&lt;principal_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudsso/accessManagement:AccessManagement")
public class AccessManagement extends com.pulumi.resources.CustomResource {
    /**
     * The Access configuration ID.
     * 
     */
    @Export(name="accessConfigurationId", type=String.class, parameters={})
    private Output<String> accessConfigurationId;

    /**
     * @return The Access configuration ID.
     * 
     */
    public Output<String> accessConfigurationId() {
        return this.accessConfigurationId;
    }
    /**
     * The deprovision strategy. Valid values: `DeprovisionForLastAccessAssignmentOnAccount` and `None`. Default Value: `DeprovisionForLastAccessAssignmentOnAccount`. **NOTE:** When `deprovision_strategy` is `DeprovisionForLastAccessAssignmentOnAccount`, and the access assignment to be deleted is the last access assignment for the same account and the same AC, this option is used for the undeployment operation。
     * 
     */
    @Export(name="deprovisionStrategy", type=String.class, parameters={})
    private Output</* @Nullable */ String> deprovisionStrategy;

    /**
     * @return The deprovision strategy. Valid values: `DeprovisionForLastAccessAssignmentOnAccount` and `None`. Default Value: `DeprovisionForLastAccessAssignmentOnAccount`. **NOTE:** When `deprovision_strategy` is `DeprovisionForLastAccessAssignmentOnAccount`, and the access assignment to be deleted is the last access assignment for the same account and the same AC, this option is used for the undeployment operation。
     * 
     */
    public Output<Optional<String>> deprovisionStrategy() {
        return Codegen.optional(this.deprovisionStrategy);
    }
    /**
     * The ID of the Directory.
     * 
     */
    @Export(name="directoryId", type=String.class, parameters={})
    private Output<String> directoryId;

    /**
     * @return The ID of the Directory.
     * 
     */
    public Output<String> directoryId() {
        return this.directoryId;
    }
    /**
     * The ID of the access assignment.
     * 
     */
    @Export(name="principalId", type=String.class, parameters={})
    private Output<String> principalId;

    /**
     * @return The ID of the access assignment.
     * 
     */
    public Output<String> principalId() {
        return this.principalId;
    }
    /**
     * The identity type of the access assignment, which can be a user or a user group. Valid values: `Group`, `User`.
     * 
     */
    @Export(name="principalType", type=String.class, parameters={})
    private Output<String> principalType;

    /**
     * @return The identity type of the access assignment, which can be a user or a user group. Valid values: `Group`, `User`.
     * 
     */
    public Output<String> principalType() {
        return this.principalType;
    }
    /**
     * The ID of the target to create the resource range.
     * 
     */
    @Export(name="targetId", type=String.class, parameters={})
    private Output<String> targetId;

    /**
     * @return The ID of the target to create the resource range.
     * 
     */
    public Output<String> targetId() {
        return this.targetId;
    }
    /**
     * The type of the resource range target to be accessed. Valid values: `RD-Account`.
     * 
     */
    @Export(name="targetType", type=String.class, parameters={})
    private Output<String> targetType;

    /**
     * @return The type of the resource range target to be accessed. Valid values: `RD-Account`.
     * 
     */
    public Output<String> targetType() {
        return this.targetType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessManagement(String name) {
        this(name, AccessManagementArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessManagement(String name, AccessManagementArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessManagement(String name, AccessManagementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudsso/accessManagement:AccessManagement", name, args == null ? AccessManagementArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessManagement(String name, Output<String> id, @Nullable AccessManagementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudsso/accessManagement:AccessManagement", name, state, makeResourceOptions(options, id));
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
    public static AccessManagement get(String name, Output<String> id, @Nullable AccessManagementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessManagement(name, id, state, options);
    }
}
