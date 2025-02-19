// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.hbr.HanaBackupPlanArgs;
import com.pulumi.alicloud.hbr.inputs.HanaBackupPlanState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Hybrid Backup Recovery (HBR) Hana Backup Plan resource.
 * 
 * For information about Hybrid Backup Recovery (HBR) Hana Backup Plan and how to use it, see [What is Hana Backup Plan](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-doc-hbr-2017-09-08-api-doc-createhanabackupplan).
 * 
 * &gt; **NOTE:** Available in v1.179.0+.
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
 * import com.pulumi.alicloud.hbr.Vault;
 * import com.pulumi.alicloud.hbr.VaultArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.hbr.HanaInstance;
 * import com.pulumi.alicloud.hbr.HanaInstanceArgs;
 * import com.pulumi.alicloud.hbr.HanaBackupPlan;
 * import com.pulumi.alicloud.hbr.HanaBackupPlanArgs;
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
 *         var defaultVault = new Vault(&#34;defaultVault&#34;, VaultArgs.builder()        
 *             .vaultName(var_.name())
 *             .build());
 * 
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status(&#34;OK&#34;)
 *             .build());
 * 
 *         var defaultHanaInstance = new HanaInstance(&#34;defaultHanaInstance&#34;, HanaInstanceArgs.builder()        
 *             .alertSetting(&#34;INHERITED&#34;)
 *             .hanaName(var_.name())
 *             .host(&#34;1.1.1.1&#34;)
 *             .instanceNumber(&#34;1&#34;)
 *             .password(&#34;YouPassword123&#34;)
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.groups()[0].id()))
 *             .sid(&#34;HXE&#34;)
 *             .useSsl(false)
 *             .userName(&#34;admin&#34;)
 *             .validateCertificate(false)
 *             .vaultId(defaultVault.id())
 *             .build());
 * 
 *         var defaultHanaBackupPlan = new HanaBackupPlan(&#34;defaultHanaBackupPlan&#34;, HanaBackupPlanArgs.builder()        
 *             .backupPrefix(&#34;DIFF_DATA_BACKUP&#34;)
 *             .backupType(&#34;COMPLETE&#34;)
 *             .clusterId(defaultHanaInstance.hanaInstanceId())
 *             .databaseName(&#34;SYSTEMDB&#34;)
 *             .planName(var_.name())
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.groups()[0].id()))
 *             .schedule(&#34;I|1602673264|P1D&#34;)
 *             .vaultId(defaultHanaInstance.vaultId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Hybrid Backup Recovery (HBR) Hana Backup Plan can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:hbr/hanaBackupPlan:HanaBackupPlan example &lt;plan_id&gt;:&lt;vault_id&gt;:&lt;cluster_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:hbr/hanaBackupPlan:HanaBackupPlan")
public class HanaBackupPlan extends com.pulumi.resources.CustomResource {
    /**
     * The backup prefix.
     * 
     */
    @Export(name="backupPrefix", type=String.class, parameters={})
    private Output</* @Nullable */ String> backupPrefix;

    /**
     * @return The backup prefix.
     * 
     */
    public Output<Optional<String>> backupPrefix() {
        return Codegen.optional(this.backupPrefix);
    }
    /**
     * The backup type. Valid values:
     * 
     */
    @Export(name="backupType", type=String.class, parameters={})
    private Output<String> backupType;

    /**
     * @return The backup type. Valid values:
     * 
     */
    public Output<String> backupType() {
        return this.backupType;
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
     * The name of the database.
     * 
     */
    @Export(name="databaseName", type=String.class, parameters={})
    private Output<String> databaseName;

    /**
     * @return The name of the database.
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }
    /**
     * The id of the plan.
     * 
     */
    @Export(name="planId", type=String.class, parameters={})
    private Output<String> planId;

    /**
     * @return The id of the plan.
     * 
     */
    public Output<String> planId() {
        return this.planId;
    }
    /**
     * The name of the backup plan.
     * 
     */
    @Export(name="planName", type=String.class, parameters={})
    private Output<String> planName;

    /**
     * @return The name of the backup plan.
     * 
     */
    public Output<String> planName() {
        return this.planName;
    }
    /**
     * The resource attribute field that represents the resource group ID.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The resource attribute field that represents the resource group ID.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
     * 
     */
    @Export(name="schedule", type=String.class, parameters={})
    private Output<String> schedule;

    /**
     * @return The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
     * 
     */
    public Output<String> schedule() {
        return this.schedule;
    }
    /**
     * The status of the resource. Valid values: `Enabled`, `Disabled`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Enabled`, `Disabled`.
     * 
     */
    public Output<String> status() {
        return this.status;
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
    public HanaBackupPlan(String name) {
        this(name, HanaBackupPlanArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HanaBackupPlan(String name, HanaBackupPlanArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HanaBackupPlan(String name, HanaBackupPlanArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/hanaBackupPlan:HanaBackupPlan", name, args == null ? HanaBackupPlanArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HanaBackupPlan(String name, Output<String> id, @Nullable HanaBackupPlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/hanaBackupPlan:HanaBackupPlan", name, state, makeResourceOptions(options, id));
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
    public static HanaBackupPlan get(String name, Output<String> id, @Nullable HanaBackupPlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HanaBackupPlan(name, id, state, options);
    }
}
