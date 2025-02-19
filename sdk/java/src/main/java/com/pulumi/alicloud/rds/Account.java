// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.AccountArgs;
import com.pulumi.alicloud.rds.inputs.AccountState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an RDS account resource and used to manage databases.
 * 
 * &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.120.0`. Please use new resource alicloud_rds_account.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.Account;
 * import com.pulumi.alicloud.rds.AccountArgs;
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
 *         final var config = ctx.config();
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;Rds&#34;);
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;dbaccountmysql&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .instanceType(&#34;rds.mysql.s1.small&#34;)
 *             .instanceStorage(&#34;10&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .instanceName(name)
 *             .build());
 * 
 *         var account = new Account(&#34;account&#34;, AccountArgs.builder()        
 *             .instanceId(instance.id())
 *             .password(&#34;Test12345&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RDS account can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rds/account:Account example &#34;rm-12345:tf_account&#34;
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/account:Account")
public class Account extends com.pulumi.resources.CustomResource {
    @Export(name="accountDescription", type=String.class, parameters={})
    private Output<String> accountDescription;

    public Output<String> accountDescription() {
        return this.accountDescription;
    }
    @Export(name="accountName", type=String.class, parameters={})
    private Output<String> accountName;

    public Output<String> accountName() {
        return this.accountName;
    }
    @Export(name="accountPassword", type=String.class, parameters={})
    private Output<String> accountPassword;

    public Output<String> accountPassword() {
        return this.accountPassword;
    }
    @Export(name="accountType", type=String.class, parameters={})
    private Output<String> accountType;

    public Output<String> accountType() {
        return this.accountType;
    }
    @Export(name="dbInstanceId", type=String.class, parameters={})
    private Output<String> dbInstanceId;

    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     * @deprecated
     * Field &#39;description&#39; has been deprecated from provider version 1.120.0. New field &#39;account_description&#39; instead.
     * 
     */
    @Deprecated /* Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead. */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The Id of instance in which account belongs.
     * 
     * @deprecated
     * Field &#39;instance_id&#39; has been deprecated from provider version 1.120.0. New field &#39;db_instance_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead. */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The Id of instance in which account belongs.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     * 
     */
    @Export(name="kmsEncryptedPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     * 
     */
    public Output<Optional<String>> kmsEncryptedPassword() {
        return Codegen.optional(this.kmsEncryptedPassword);
    }
    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Export(name="kmsEncryptionContext", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Output<Optional<Map<String,Object>>> kmsEncryptionContext() {
        return Codegen.optional(this.kmsEncryptionContext);
    }
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;account_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     * @deprecated
     * Field &#39;password&#39; has been deprecated from provider version 1.120.0. New field &#39;account_password&#39; instead.
     * 
     */
    @Deprecated /* Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead. */
    @Export(name="password", type=String.class, parameters={})
    private Output<String> password;

    /**
     * @return Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    @Export(name="resetPermissionFlag", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> resetPermissionFlag;

    public Output<Optional<Boolean>> resetPermissionFlag() {
        return Codegen.optional(this.resetPermissionFlag);
    }
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     * 
     * Default to Normal.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.120.0. New field &#39;account_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead. */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     * 
     * Default to Normal.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Account(String name) {
        this(name, AccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Account(String name, @Nullable AccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Account(String name, @Nullable AccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/account:Account", name, args == null ? AccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Account(String name, Output<String> id, @Nullable AccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/account:Account", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountPassword",
                "password"
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
    public static Account get(String name, Output<String> id, @Nullable AccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Account(name, id, state, options);
    }
}
