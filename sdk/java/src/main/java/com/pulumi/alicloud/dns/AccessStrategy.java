// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.AccessStrategyArgs;
import com.pulumi.alicloud.dns.inputs.AccessStrategyState;
import com.pulumi.alicloud.dns.outputs.AccessStrategyDefaultAddrPool;
import com.pulumi.alicloud.dns.outputs.AccessStrategyFailoverAddrPool;
import com.pulumi.alicloud.dns.outputs.AccessStrategyLine;
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
 * Provides a DNS Access Strategy resource.
 * 
 * For information about DNS Access Strategy and how to use it, see [What is Access Strategy](https://www.alibabacloud.com/help/doc-detail/189620.html).
 * 
 * &gt; **NOTE:** Available in v1.152.0+.
 * 
 * ## Import
 * 
 * DNS Access Strategy can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dns/accessStrategy:AccessStrategy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/accessStrategy:AccessStrategy")
public class AccessStrategy extends com.pulumi.resources.CustomResource {
    /**
     * The primary/secondary switchover policy for address pool groups. Valid values: `AUTO`, `DEFAULT`, `FAILOVER`.
     * 
     */
    @Export(name="accessMode", type=String.class, parameters={})
    private Output<String> accessMode;

    /**
     * @return The primary/secondary switchover policy for address pool groups. Valid values: `AUTO`, `DEFAULT`, `FAILOVER`.
     * 
     */
    public Output<String> accessMode() {
        return this.accessMode;
    }
    /**
     * The type of the primary address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    @Export(name="defaultAddrPoolType", type=String.class, parameters={})
    private Output<String> defaultAddrPoolType;

    /**
     * @return The type of the primary address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    public Output<String> defaultAddrPoolType() {
        return this.defaultAddrPoolType;
    }
    /**
     * List of primary address pool collections. See the following `Block default_addr_pools`.
     * 
     */
    @Export(name="defaultAddrPools", type=List.class, parameters={AccessStrategyDefaultAddrPool.class})
    private Output<List<AccessStrategyDefaultAddrPool>> defaultAddrPools;

    /**
     * @return List of primary address pool collections. See the following `Block default_addr_pools`.
     * 
     */
    public Output<List<AccessStrategyDefaultAddrPool>> defaultAddrPools() {
        return this.defaultAddrPools;
    }
    /**
     * Specifies whether to enable scheduling optimization for latency resolution for the primary address pool group. Valid values: `OPEN`, `CLOSE`.
     * 
     */
    @Export(name="defaultLatencyOptimization", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultLatencyOptimization;

    /**
     * @return Specifies whether to enable scheduling optimization for latency resolution for the primary address pool group. Valid values: `OPEN`, `CLOSE`.
     * 
     */
    public Output<Optional<String>> defaultLatencyOptimization() {
        return Codegen.optional(this.defaultLatencyOptimization);
    }
    /**
     * The load balancing policy of the primary address pool group. Valid values: `ALL_RR`, `RATIO`. **NOTE:** The `default_lba_strategy` is required under the condition that `strategy_mode` is `GEO`.
     * 
     */
    @Export(name="defaultLbaStrategy", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultLbaStrategy;

    /**
     * @return The load balancing policy of the primary address pool group. Valid values: `ALL_RR`, `RATIO`. **NOTE:** The `default_lba_strategy` is required under the condition that `strategy_mode` is `GEO`.
     * 
     */
    public Output<Optional<String>> defaultLbaStrategy() {
        return Codegen.optional(this.defaultLbaStrategy);
    }
    /**
     * The maximum number of addresses returned by the primary address pool set. **NOTE:** The `default_max_return_addr_num` is required under the condition that `strategy_mode` is `LATENCY`.
     * 
     */
    @Export(name="defaultMaxReturnAddrNum", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> defaultMaxReturnAddrNum;

    /**
     * @return The maximum number of addresses returned by the primary address pool set. **NOTE:** The `default_max_return_addr_num` is required under the condition that `strategy_mode` is `LATENCY`.
     * 
     */
    public Output<Optional<Integer>> defaultMaxReturnAddrNum() {
        return Codegen.optional(this.defaultMaxReturnAddrNum);
    }
    /**
     * The minimum number of available addresses for the primary address pool set.
     * 
     */
    @Export(name="defaultMinAvailableAddrNum", type=Integer.class, parameters={})
    private Output<Integer> defaultMinAvailableAddrNum;

    /**
     * @return The minimum number of available addresses for the primary address pool set.
     * 
     */
    public Output<Integer> defaultMinAvailableAddrNum() {
        return this.defaultMinAvailableAddrNum;
    }
    /**
     * The type of the secondary address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    @Export(name="failoverAddrPoolType", type=String.class, parameters={})
    private Output</* @Nullable */ String> failoverAddrPoolType;

    /**
     * @return The type of the secondary address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    public Output<Optional<String>> failoverAddrPoolType() {
        return Codegen.optional(this.failoverAddrPoolType);
    }
    /**
     * List of backup address pool sets. See the following `Block failover_addr_pools`.
     * 
     */
    @Export(name="failoverAddrPools", type=List.class, parameters={AccessStrategyFailoverAddrPool.class})
    private Output</* @Nullable */ List<AccessStrategyFailoverAddrPool>> failoverAddrPools;

    /**
     * @return List of backup address pool sets. See the following `Block failover_addr_pools`.
     * 
     */
    public Output<Optional<List<AccessStrategyFailoverAddrPool>>> failoverAddrPools() {
        return Codegen.optional(this.failoverAddrPools);
    }
    /**
     * Specifies whether to enable scheduling optimization for latency resolution for the secondary address pool group. Valid values: `OPEN`, `CLOSE`.
     * 
     */
    @Export(name="failoverLatencyOptimization", type=String.class, parameters={})
    private Output</* @Nullable */ String> failoverLatencyOptimization;

    /**
     * @return Specifies whether to enable scheduling optimization for latency resolution for the secondary address pool group. Valid values: `OPEN`, `CLOSE`.
     * 
     */
    public Output<Optional<String>> failoverLatencyOptimization() {
        return Codegen.optional(this.failoverLatencyOptimization);
    }
    /**
     * The load balancing policy of the secondary address pool group. Valid values: `ALL_RR`, `RATIO`.
     * 
     */
    @Export(name="failoverLbaStrategy", type=String.class, parameters={})
    private Output</* @Nullable */ String> failoverLbaStrategy;

    /**
     * @return The load balancing policy of the secondary address pool group. Valid values: `ALL_RR`, `RATIO`.
     * 
     */
    public Output<Optional<String>> failoverLbaStrategy() {
        return Codegen.optional(this.failoverLbaStrategy);
    }
    /**
     * The maximum number of returned addresses in the standby address pool.
     * 
     */
    @Export(name="failoverMaxReturnAddrNum", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> failoverMaxReturnAddrNum;

    /**
     * @return The maximum number of returned addresses in the standby address pool.
     * 
     */
    public Output<Optional<Integer>> failoverMaxReturnAddrNum() {
        return Codegen.optional(this.failoverMaxReturnAddrNum);
    }
    /**
     * The minimum number of available addresses in the standby address pool.
     * 
     */
    @Export(name="failoverMinAvailableAddrNum", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> failoverMinAvailableAddrNum;

    /**
     * @return The minimum number of available addresses in the standby address pool.
     * 
     */
    public Output<Optional<Integer>> failoverMinAvailableAddrNum() {
        return Codegen.optional(this.failoverMinAvailableAddrNum);
    }
    /**
     * The Id of the associated instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The Id of the associated instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The lang.
     * 
     */
    @Export(name="lang", type=String.class, parameters={})
    private Output</* @Nullable */ String> lang;

    /**
     * @return The lang.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The source regions. See the following `Block lines`. **NOTE:** The `lines` is required under the condition that `strategy_mode` is `GEO`.
     * 
     */
    @Export(name="lines", type=List.class, parameters={AccessStrategyLine.class})
    private Output</* @Nullable */ List<AccessStrategyLine>> lines;

    /**
     * @return The source regions. See the following `Block lines`. **NOTE:** The `lines` is required under the condition that `strategy_mode` is `GEO`.
     * 
     */
    public Output<Optional<List<AccessStrategyLine>>> lines() {
        return Codegen.optional(this.lines);
    }
    /**
     * The type of the access policy. Valid values: `GEO` or `LATENCY`. `GEO`: based on geographic location. `LATENCY`: Based on delay.
     * 
     */
    @Export(name="strategyMode", type=String.class, parameters={})
    private Output<String> strategyMode;

    /**
     * @return The type of the access policy. Valid values: `GEO` or `LATENCY`. `GEO`: based on geographic location. `LATENCY`: Based on delay.
     * 
     */
    public Output<String> strategyMode() {
        return this.strategyMode;
    }
    /**
     * The name of the access policy.
     * 
     */
    @Export(name="strategyName", type=String.class, parameters={})
    private Output<String> strategyName;

    /**
     * @return The name of the access policy.
     * 
     */
    public Output<String> strategyName() {
        return this.strategyName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessStrategy(String name) {
        this(name, AccessStrategyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessStrategy(String name, AccessStrategyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessStrategy(String name, AccessStrategyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/accessStrategy:AccessStrategy", name, args == null ? AccessStrategyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessStrategy(String name, Output<String> id, @Nullable AccessStrategyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/accessStrategy:AccessStrategy", name, state, makeResourceOptions(options, id));
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
    public static AccessStrategy get(String name, Output<String> id, @Nullable AccessStrategyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessStrategy(name, id, state, options);
    }
}
