// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.hbr.HanaBackupClientArgs;
import com.pulumi.alicloud.hbr.inputs.HanaBackupClientState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Hybrid Backup Recovery (HBR) Hana Backup Client resource.
 * 
 * For information about Hybrid Backup Recovery (HBR) Hana Backup Client and how to use it, see [What is Hana Backup Client](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-doc-hbr-2017-09-08-api-doc-createclients).
 * 
 * &gt; **NOTE:** Available in v1.198.0+.
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
 * import com.pulumi.alicloud.hbr.HanaBackupClient;
 * import com.pulumi.alicloud.hbr.HanaBackupClientArgs;
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
 *         var default_ = new HanaBackupClient(&#34;default&#34;, HanaBackupClientArgs.builder()        
 *             .vaultId(data.alicloud_hbr_vaults().default().vaults()[0].id())
 *             .clientInfo(&#34;[ { \&#34;instanceId\&#34;: \&#34;i-bp116lr******te9q2\&#34;, \&#34;clusterId\&#34;: \&#34;cl-000csy09q******9rfz9\&#34;, \&#34;sourceTypes\&#34;: [ \&#34;HANA\&#34; ]  }]&#34;)
 *             .alertSetting(&#34;INHERITED&#34;)
 *             .useHttps(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Hybrid Backup Recovery (HBR) Hana Backup Client can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:hbr/hanaBackupClient:HanaBackupClient example &lt;vault_id&gt;:&lt;client_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:hbr/hanaBackupClient:HanaBackupClient")
public class HanaBackupClient extends com.pulumi.resources.CustomResource {
    /**
     * The alert settings. Valid value: `INHERITED`.
     * 
     */
    @Export(name="alertSetting", type=String.class, parameters={})
    private Output<String> alertSetting;

    /**
     * @return The alert settings. Valid value: `INHERITED`.
     * 
     */
    public Output<String> alertSetting() {
        return this.alertSetting;
    }
    /**
     * The ID of the backup client.
     * 
     */
    @Export(name="clientId", type=String.class, parameters={})
    private Output<String> clientId;

    /**
     * @return The ID of the backup client.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * The installation information of the HBR clients.
     * 
     */
    @Export(name="clientInfo", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientInfo;

    /**
     * @return The installation information of the HBR clients.
     * 
     */
    public Output<Optional<String>> clientInfo() {
        return Codegen.optional(this.clientInfo);
    }
    /**
     * The ID of the SAP HANA instance.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return The ID of the SAP HANA instance.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The status of the Hana Backup Client.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the Hana Backup Client.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
     * 
     */
    @Export(name="useHttps", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> useHttps;

    /**
     * @return Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> useHttps() {
        return Codegen.optional(this.useHttps);
    }
    /**
     * The ID of the backup vault.
     * 
     */
    @Export(name="vaultId", type=String.class, parameters={})
    private Output<String> vaultId;

    /**
     * @return The ID of the backup vault.
     * 
     */
    public Output<String> vaultId() {
        return this.vaultId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HanaBackupClient(String name) {
        this(name, HanaBackupClientArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HanaBackupClient(String name, HanaBackupClientArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HanaBackupClient(String name, HanaBackupClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/hanaBackupClient:HanaBackupClient", name, args == null ? HanaBackupClientArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HanaBackupClient(String name, Output<String> id, @Nullable HanaBackupClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/hanaBackupClient:HanaBackupClient", name, state, makeResourceOptions(options, id));
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
    public static HanaBackupClient get(String name, Output<String> id, @Nullable HanaBackupClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HanaBackupClient(name, id, state, options);
    }
}
