// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cfg.AggregateCompliancePackArgs;
import com.pulumi.alicloud.cfg.inputs.AggregateCompliancePackState;
import com.pulumi.alicloud.cfg.outputs.AggregateCompliancePackConfigRule;
import com.pulumi.alicloud.cfg.outputs.AggregateCompliancePackConfigRuleId;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Config Aggregate Compliance Pack resource.
 * 
 * For information about Cloud Config Aggregate Compliance Pack and how to use it, see [What is Aggregate Compliance Pack](https://www.alibabacloud.com/help/en/doc-detail/194753.html).
 * 
 * &gt; **NOTE:** Available in v1.124.0+.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstancesArgs;
 * import com.pulumi.alicloud.cfg.Aggregator;
 * import com.pulumi.alicloud.cfg.AggregatorArgs;
 * import com.pulumi.alicloud.cfg.inputs.AggregatorAggregatorAccountArgs;
 * import com.pulumi.alicloud.cfg.AggregateConfigRule;
 * import com.pulumi.alicloud.cfg.AggregateConfigRuleArgs;
 * import com.pulumi.alicloud.cfg.AggregateCompliancePack;
 * import com.pulumi.alicloud.cfg.AggregateCompliancePackArgs;
 * import com.pulumi.alicloud.cfg.inputs.AggregateCompliancePackConfigRuleIdArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;example_name&#34;);
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status(&#34;OK&#34;)
 *             .build());
 * 
 *         final var defaultInstances = EcsFunctions.getInstances();
 * 
 *         var defaultAggregator = new Aggregator(&#34;defaultAggregator&#34;, AggregatorArgs.builder()        
 *             .aggregatorAccounts(AggregatorAggregatorAccountArgs.builder()
 *                 .accountId(&#34;140278452670****&#34;)
 *                 .accountName(&#34;test-2&#34;)
 *                 .accountType(&#34;ResourceDirectory&#34;)
 *                 .build())
 *             .aggregatorName(&#34;tf-testaccaggregator&#34;)
 *             .description(&#34;tf-testaccaggregator&#34;)
 *             .build());
 * 
 *         var defaultAggregateConfigRule = new AggregateConfigRule(&#34;defaultAggregateConfigRule&#34;, AggregateConfigRuleArgs.builder()        
 *             .aggregatorId(defaultAggregator.id())
 *             .aggregateConfigRuleName(name)
 *             .sourceOwner(&#34;ALIYUN&#34;)
 *             .sourceIdentifier(&#34;ecs-cpu-min-count-limit&#34;)
 *             .configRuleTriggerTypes(&#34;ConfigurationItemChangeNotification&#34;)
 *             .resourceTypesScopes(&#34;ACS::ECS::Instance&#34;)
 *             .riskLevel(1)
 *             .description(name)
 *             .excludeResourceIdsScope(defaultInstances.applyValue(getInstancesResult -&gt; getInstancesResult.ids()[0]))
 *             .inputParameters(Map.of(&#34;cpuCount&#34;, &#34;4&#34;))
 *             .regionIdsScope(&#34;cn-hangzhou&#34;)
 *             .resourceGroupIdsScope(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .tagKeyScope(&#34;tFTest&#34;)
 *             .tagValueScope(&#34;forTF 123&#34;)
 *             .build());
 * 
 *         var defaultAggregateCompliancePack = new AggregateCompliancePack(&#34;defaultAggregateCompliancePack&#34;, AggregateCompliancePackArgs.builder()        
 *             .aggregateCompliancePackName(&#34;tf-testaccConfig1234&#34;)
 *             .aggregatorId(defaultAggregator.id())
 *             .description(&#34;tf-testaccConfig1234&#34;)
 *             .riskLevel(1)
 *             .configRuleIds(AggregateCompliancePackConfigRuleIdArgs.builder()
 *                 .configRuleId(defaultAggregateConfigRule.configRuleId())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Config Aggregate Compliance Pack can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack example &lt;aggregator_id&gt;:&lt;aggregator_compliance_pack_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack")
public class AggregateCompliancePack extends com.pulumi.resources.CustomResource {
    /**
     * The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
     * 
     */
    @Export(name="aggregateCompliancePackName", type=String.class, parameters={})
    private Output<String> aggregateCompliancePackName;

    /**
     * @return The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
     * 
     */
    public Output<String> aggregateCompliancePackName() {
        return this.aggregateCompliancePackName;
    }
    /**
     * The ID of aggregator.
     * 
     */
    @Export(name="aggregatorId", type=String.class, parameters={})
    private Output<String> aggregatorId;

    /**
     * @return The ID of aggregator.
     * 
     */
    public Output<String> aggregatorId() {
        return this.aggregatorId;
    }
    /**
     * The Template ID of compliance package.
     * 
     */
    @Export(name="compliancePackTemplateId", type=String.class, parameters={})
    private Output</* @Nullable */ String> compliancePackTemplateId;

    /**
     * @return The Template ID of compliance package.
     * 
     */
    public Output<Optional<String>> compliancePackTemplateId() {
        return Codegen.optional(this.compliancePackTemplateId);
    }
    /**
     * A list of Config Rule IDs.
     * 
     */
    @Export(name="configRuleIds", type=List.class, parameters={AggregateCompliancePackConfigRuleId.class})
    private Output</* @Nullable */ List<AggregateCompliancePackConfigRuleId>> configRuleIds;

    /**
     * @return A list of Config Rule IDs.
     * 
     */
    public Output<Optional<List<AggregateCompliancePackConfigRuleId>>> configRuleIds() {
        return Codegen.optional(this.configRuleIds);
    }
    /**
     * A list of Config Rules.
     * 
     * @deprecated
     * Field &#39;config_rules&#39; has been deprecated from provider version 1.141.0. New field &#39;config_rule_ids&#39; instead.
     * 
     */
    @Deprecated /* Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead. */
    @Export(name="configRules", type=List.class, parameters={AggregateCompliancePackConfigRule.class})
    private Output</* @Nullable */ List<AggregateCompliancePackConfigRule>> configRules;

    /**
     * @return A list of Config Rules.
     * 
     */
    public Output<Optional<List<AggregateCompliancePackConfigRule>>> configRules() {
        return Codegen.optional(this.configRules);
    }
    /**
     * The description of compliance package.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return The description of compliance package.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The Risk Level. Valid values: `1`: critical `2`: warning `3`: info.
     * 
     */
    @Export(name="riskLevel", type=Integer.class, parameters={})
    private Output<Integer> riskLevel;

    /**
     * @return The Risk Level. Valid values: `1`: critical `2`: warning `3`: info.
     * 
     */
    public Output<Integer> riskLevel() {
        return this.riskLevel;
    }
    /**
     * The status of the resource. The valid values: `CREATING`, `ACTIVE`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource. The valid values: `CREATING`, `ACTIVE`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AggregateCompliancePack(String name) {
        this(name, AggregateCompliancePackArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AggregateCompliancePack(String name, AggregateCompliancePackArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AggregateCompliancePack(String name, AggregateCompliancePackArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack", name, args == null ? AggregateCompliancePackArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AggregateCompliancePack(String name, Output<String> id, @Nullable AggregateCompliancePackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack", name, state, makeResourceOptions(options, id));
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
    public static AggregateCompliancePack get(String name, Output<String> id, @Nullable AggregateCompliancePackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AggregateCompliancePack(name, id, state, options);
    }
}
