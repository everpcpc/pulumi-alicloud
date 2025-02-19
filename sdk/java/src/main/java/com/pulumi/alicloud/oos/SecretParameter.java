// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oos.SecretParameterArgs;
import com.pulumi.alicloud.oos.inputs.SecretParameterState;
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
 * Provides a OOS Secret Parameter resource.
 * 
 * For information about OOS Secret Parameter and how to use it, see [What is Secret Parameter](https://www.alibabacloud.com/help/en/doc-detail/183418.html).
 * 
 * &gt; **NOTE:** Available in v1.147.0+.
 * 
 * ## Import
 * 
 * OOS Secret Parameter can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:oos/secretParameter:SecretParameter example &lt;secret_parameter_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:oos/secretParameter:SecretParameter")
public class SecretParameter extends com.pulumi.resources.CustomResource {
    /**
     * The constraints of the encryption parameter. By default, this parameter is null. Valid values:
     * 
     */
    @Export(name="constraints", type=String.class, parameters={})
    private Output</* @Nullable */ String> constraints;

    /**
     * @return The constraints of the encryption parameter. By default, this parameter is null. Valid values:
     * 
     */
    public Output<Optional<String>> constraints() {
        return Codegen.optional(this.constraints);
    }
    /**
     * The description of the encryption parameter. The description must be `1` to `200` characters in length.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the encryption parameter. The description must be `1` to `200` characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
     * 
     */
    @Export(name="keyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> keyId;

    /**
     * @return The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
     * 
     */
    public Output<Optional<String>> keyId() {
        return Codegen.optional(this.keyId);
    }
    /**
     * The ID of the Resource Group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the Resource Group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     * 
     */
    @Export(name="secretParameterName", type=String.class, parameters={})
    private Output<String> secretParameterName;

    /**
     * @return The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     * 
     */
    public Output<String> secretParameterName() {
        return this.secretParameterName;
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
     * The data type of the encryption parameter. Valid values: `Secret`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The data type of the encryption parameter. Valid values: `Secret`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The value of the encryption parameter. The value must be `1` to `4096` characters in length.
     * 
     */
    @Export(name="value", type=String.class, parameters={})
    private Output<String> value;

    /**
     * @return The value of the encryption parameter. The value must be `1` to `4096` characters in length.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretParameter(String name) {
        this(name, SecretParameterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretParameter(String name, SecretParameterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretParameter(String name, SecretParameterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oos/secretParameter:SecretParameter", name, args == null ? SecretParameterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretParameter(String name, Output<String> id, @Nullable SecretParameterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oos/secretParameter:SecretParameter", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "value"
            ))
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
    public static SecretParameter get(String name, Output<String> id, @Nullable SecretParameterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretParameter(name, id, state, options);
    }
}
