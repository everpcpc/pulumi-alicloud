// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.InstanceArgs;
import com.pulumi.alicloud.threatdetection.inputs.InstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Instance resource.
 * 
 * For information about Threat Detection Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/security-center/latest/what-is-security-center).
 * 
 * &gt; **NOTE:** Available in v1.199.0+.
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
 * import com.pulumi.alicloud.threatdetection.Instance;
 * import com.pulumi.alicloud.threatdetection.InstanceArgs;
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
 *         var default_ = new Instance(&#34;default&#34;, InstanceArgs.builder()        
 *             .buyNumber(&#34;30&#34;)
 *             .containerImageScan(&#34;100&#34;)
 *             .honeypot(&#34;32&#34;)
 *             .honeypotSwitch(&#34;1&#34;)
 *             .paymentType(&#34;Subscription&#34;)
 *             .period(12)
 *             .renewalStatus(&#34;ManualRenewal&#34;)
 *             .sasAntiRansomware(&#34;100&#34;)
 *             .sasSc(&#34;true&#34;)
 *             .sasSdk(&#34;1000&#34;)
 *             .sasSdkSwitch(&#34;1&#34;)
 *             .sasSlsStorage(&#34;100&#34;)
 *             .sasWebguardOrderNum(&#34;100&#34;)
 *             .vCore(&#34;100&#34;)
 *             .versionCode(&#34;level2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Threat Detection Instance do not support import.
 * 
 */
@ResourceType(type="alicloud:threatdetection/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Number of servers.
     * 
     */
    @Export(name="buyNumber", type=String.class, parameters={})
    private Output</* @Nullable */ String> buyNumber;

    /**
     * @return Number of servers.
     * 
     */
    public Output<Optional<String>> buyNumber() {
        return Codegen.optional(this.buyNumber);
    }
    /**
     * Container Image security scan.
     * 
     */
    @Export(name="containerImageScan", type=String.class, parameters={})
    private Output</* @Nullable */ String> containerImageScan;

    /**
     * @return Container Image security scan.
     * 
     */
    public Output<Optional<String>> containerImageScan() {
        return Codegen.optional(this.containerImageScan);
    }
    /**
     * The creation time of the resource
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The creation time of the resource
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Cloud honeypot authorization number.
     * 
     */
    @Export(name="honeypot", type=String.class, parameters={})
    private Output</* @Nullable */ String> honeypot;

    /**
     * @return Cloud honeypot authorization number.
     * 
     */
    public Output<Optional<String>> honeypot() {
        return Codegen.optional(this.honeypot);
    }
    /**
     * Cloud honeypot. Valid values: `1`, `2`.
     * 
     */
    @Export(name="honeypotSwitch", type=String.class, parameters={})
    private Output</* @Nullable */ String> honeypotSwitch;

    /**
     * @return Cloud honeypot. Valid values: `1`, `2`.
     * 
     */
    public Output<Optional<String>> honeypotSwitch() {
        return Codegen.optional(this.honeypotSwitch);
    }
    /**
     * The first ID of the resource
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Change configuration type, value
     * - Upgrade: Upgrade.
     * - Downgrade: Downgrade.
     * 
     */
    @Export(name="modifyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> modifyType;

    /**
     * @return Change configuration type, value
     * - Upgrade: Upgrade.
     * - Downgrade: Downgrade.
     * 
     */
    public Output<Optional<String>> modifyType() {
        return Codegen.optional(this.modifyType);
    }
    /**
     * The payment type of the resource.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * Prepaid cycle. The unit is Monthly, please enter an integer multiple of 12 for annual paid products. **NOTE:** must be set when creating a prepaid instance.
     * 
     */
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    /**
     * @return Prepaid cycle. The unit is Monthly, please enter an integer multiple of 12 for annual paid products. **NOTE:** must be set when creating a prepaid instance.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Automatic renewal cycle, in months. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`.
     * 
     */
    @Export(name="renewPeriod", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> renewPeriod;

    /**
     * @return Automatic renewal cycle, in months. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`.
     * 
     */
    public Output<Optional<Integer>> renewPeriod() {
        return Codegen.optional(this.renewPeriod);
    }
    /**
     * The unit of the auto-renewal period. **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`. Valid values:
     * 
     */
    @Export(name="renewalPeriodUnit", type=String.class, parameters={})
    private Output<String> renewalPeriodUnit;

    /**
     * @return The unit of the auto-renewal period. **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`. Valid values:
     * 
     */
    public Output<String> renewalPeriodUnit() {
        return this.renewalPeriodUnit;
    }
    /**
     * Automatic renewal status, Default ManualRenewal. value:
     * 
     */
    @Export(name="renewalStatus", type=String.class, parameters={})
    private Output<String> renewalStatus;

    /**
     * @return Automatic renewal status, Default ManualRenewal. value:
     * 
     */
    public Output<String> renewalStatus() {
        return this.renewalStatus;
    }
    /**
     * Anti-extortion.
     * 
     */
    @Export(name="sasAntiRansomware", type=String.class, parameters={})
    private Output</* @Nullable */ String> sasAntiRansomware;

    /**
     * @return Anti-extortion.
     * 
     */
    public Output<Optional<String>> sasAntiRansomware() {
        return Codegen.optional(this.sasAntiRansomware);
    }
    /**
     * Large security screen.
     * 
     */
    @Export(name="sasSc", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> sasSc;

    /**
     * @return Large security screen.
     * 
     */
    public Output<Optional<Boolean>> sasSc() {
        return Codegen.optional(this.sasSc);
    }
    /**
     * Number of malicious file detections.
     * 
     */
    @Export(name="sasSdk", type=String.class, parameters={})
    private Output</* @Nullable */ String> sasSdk;

    /**
     * @return Number of malicious file detections.
     * 
     */
    public Output<Optional<String>> sasSdk() {
        return Codegen.optional(this.sasSdk);
    }
    /**
     * Malicious file detection SDK. Valid values: `0`, `1`.
     * 
     */
    @Export(name="sasSdkSwitch", type=String.class, parameters={})
    private Output</* @Nullable */ String> sasSdkSwitch;

    /**
     * @return Malicious file detection SDK. Valid values: `0`, `1`.
     * 
     */
    public Output<Optional<String>> sasSdkSwitch() {
        return Codegen.optional(this.sasSdkSwitch);
    }
    /**
     * Log analysis.
     * 
     */
    @Export(name="sasSlsStorage", type=String.class, parameters={})
    private Output</* @Nullable */ String> sasSlsStorage;

    /**
     * @return Log analysis.
     * 
     */
    public Output<Optional<String>> sasSlsStorage() {
        return Codegen.optional(this.sasSlsStorage);
    }
    /**
     * Web page tamper-proof.  Valid values: `0`, `1`.
     * 
     */
    @Export(name="sasWebguardBoolean", type=String.class, parameters={})
    private Output</* @Nullable */ String> sasWebguardBoolean;

    /**
     * @return Web page tamper-proof.  Valid values: `0`, `1`.
     * 
     */
    public Output<Optional<String>> sasWebguardBoolean() {
        return Codegen.optional(this.sasWebguardBoolean);
    }
    /**
     * Number of tamper-proof authorizations.
     * 
     */
    @Export(name="sasWebguardOrderNum", type=String.class, parameters={})
    private Output</* @Nullable */ String> sasWebguardOrderNum;

    /**
     * @return Number of tamper-proof authorizations.
     * 
     */
    public Output<Optional<String>> sasWebguardOrderNum() {
        return Codegen.optional(this.sasWebguardOrderNum);
    }
    /**
     * The status of the resource
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The amount of threat analysis log storage.
     * 
     */
    @Export(name="threatAnalysis", type=String.class, parameters={})
    private Output</* @Nullable */ String> threatAnalysis;

    /**
     * @return The amount of threat analysis log storage.
     * 
     */
    public Output<Optional<String>> threatAnalysis() {
        return Codegen.optional(this.threatAnalysis);
    }
    /**
     * Threat analysis.  Valid values: `0`, `1`.
     * 
     */
    @Export(name="threatAnalysisSwitch", type=String.class, parameters={})
    private Output</* @Nullable */ String> threatAnalysisSwitch;

    /**
     * @return Threat analysis.  Valid values: `0`, `1`.
     * 
     */
    public Output<Optional<String>> threatAnalysisSwitch() {
        return Codegen.optional(this.threatAnalysisSwitch);
    }
    /**
     * Number of cores.
     * 
     */
    @Export(name="vCore", type=String.class, parameters={})
    private Output</* @Nullable */ String> vCore;

    /**
     * @return Number of cores.
     * 
     */
    public Output<Optional<String>> vCore() {
        return Codegen.optional(this.vCore);
    }
    /**
     * Version selection. Valid values: `level10`, `level2`, `level3`, `level7`, `level8`.
     * 
     */
    @Export(name="versionCode", type=String.class, parameters={})
    private Output<String> versionCode;

    /**
     * @return Version selection. Valid values: `level10`, `level2`, `level3`, `level7`, `level8`.
     * 
     */
    public Output<String> versionCode() {
        return this.versionCode;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
