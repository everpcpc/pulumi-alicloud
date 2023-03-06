// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * Number of servers.
     * 
     */
    @Import(name="buyNumber")
    private @Nullable Output<String> buyNumber;

    /**
     * @return Number of servers.
     * 
     */
    public Optional<Output<String>> buyNumber() {
        return Optional.ofNullable(this.buyNumber);
    }

    /**
     * Container Image security scan.
     * 
     */
    @Import(name="containerImageScan")
    private @Nullable Output<String> containerImageScan;

    /**
     * @return Container Image security scan.
     * 
     */
    public Optional<Output<String>> containerImageScan() {
        return Optional.ofNullable(this.containerImageScan);
    }

    /**
     * Cloud honeypot authorization number.
     * 
     */
    @Import(name="honeypot")
    private @Nullable Output<String> honeypot;

    /**
     * @return Cloud honeypot authorization number.
     * 
     */
    public Optional<Output<String>> honeypot() {
        return Optional.ofNullable(this.honeypot);
    }

    /**
     * Cloud honeypot. Valid values: `1`, `2`.
     * 
     */
    @Import(name="honeypotSwitch")
    private @Nullable Output<String> honeypotSwitch;

    /**
     * @return Cloud honeypot. Valid values: `1`, `2`.
     * 
     */
    public Optional<Output<String>> honeypotSwitch() {
        return Optional.ofNullable(this.honeypotSwitch);
    }

    /**
     * The first ID of the resource
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Change configuration type, value
     * - Upgrade: Upgrade.
     * - Downgrade: Downgrade.
     * 
     */
    @Import(name="modifyType")
    private @Nullable Output<String> modifyType;

    /**
     * @return Change configuration type, value
     * - Upgrade: Upgrade.
     * - Downgrade: Downgrade.
     * 
     */
    public Optional<Output<String>> modifyType() {
        return Optional.ofNullable(this.modifyType);
    }

    /**
     * The payment type of the resource.
     * 
     */
    @Import(name="paymentType", required=true)
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
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return Prepaid cycle. The unit is Monthly, please enter an integer multiple of 12 for annual paid products. **NOTE:** must be set when creating a prepaid instance.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * Automatic renewal cycle, in months. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`.
     * 
     */
    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    /**
     * @return Automatic renewal cycle, in months. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`.
     * 
     */
    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    /**
     * The unit of the auto-renewal period. **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`. Valid values:
     * 
     */
    @Import(name="renewalPeriodUnit")
    private @Nullable Output<String> renewalPeriodUnit;

    /**
     * @return The unit of the auto-renewal period. **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`. Valid values:
     * 
     */
    public Optional<Output<String>> renewalPeriodUnit() {
        return Optional.ofNullable(this.renewalPeriodUnit);
    }

    /**
     * Automatic renewal status, Default ManualRenewal. value:
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Automatic renewal status, Default ManualRenewal. value:
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    /**
     * Anti-extortion.
     * 
     */
    @Import(name="sasAntiRansomware")
    private @Nullable Output<String> sasAntiRansomware;

    /**
     * @return Anti-extortion.
     * 
     */
    public Optional<Output<String>> sasAntiRansomware() {
        return Optional.ofNullable(this.sasAntiRansomware);
    }

    /**
     * Large security screen.
     * 
     */
    @Import(name="sasSc")
    private @Nullable Output<Boolean> sasSc;

    /**
     * @return Large security screen.
     * 
     */
    public Optional<Output<Boolean>> sasSc() {
        return Optional.ofNullable(this.sasSc);
    }

    /**
     * Number of malicious file detections.
     * 
     */
    @Import(name="sasSdk")
    private @Nullable Output<String> sasSdk;

    /**
     * @return Number of malicious file detections.
     * 
     */
    public Optional<Output<String>> sasSdk() {
        return Optional.ofNullable(this.sasSdk);
    }

    /**
     * Malicious file detection SDK. Valid values: `0`, `1`.
     * 
     */
    @Import(name="sasSdkSwitch")
    private @Nullable Output<String> sasSdkSwitch;

    /**
     * @return Malicious file detection SDK. Valid values: `0`, `1`.
     * 
     */
    public Optional<Output<String>> sasSdkSwitch() {
        return Optional.ofNullable(this.sasSdkSwitch);
    }

    /**
     * Log analysis.
     * 
     */
    @Import(name="sasSlsStorage")
    private @Nullable Output<String> sasSlsStorage;

    /**
     * @return Log analysis.
     * 
     */
    public Optional<Output<String>> sasSlsStorage() {
        return Optional.ofNullable(this.sasSlsStorage);
    }

    /**
     * Web page tamper-proof.  Valid values: `0`, `1`.
     * 
     */
    @Import(name="sasWebguardBoolean")
    private @Nullable Output<String> sasWebguardBoolean;

    /**
     * @return Web page tamper-proof.  Valid values: `0`, `1`.
     * 
     */
    public Optional<Output<String>> sasWebguardBoolean() {
        return Optional.ofNullable(this.sasWebguardBoolean);
    }

    /**
     * Number of tamper-proof authorizations.
     * 
     */
    @Import(name="sasWebguardOrderNum")
    private @Nullable Output<String> sasWebguardOrderNum;

    /**
     * @return Number of tamper-proof authorizations.
     * 
     */
    public Optional<Output<String>> sasWebguardOrderNum() {
        return Optional.ofNullable(this.sasWebguardOrderNum);
    }

    /**
     * The amount of threat analysis log storage.
     * 
     */
    @Import(name="threatAnalysis")
    private @Nullable Output<String> threatAnalysis;

    /**
     * @return The amount of threat analysis log storage.
     * 
     */
    public Optional<Output<String>> threatAnalysis() {
        return Optional.ofNullable(this.threatAnalysis);
    }

    /**
     * Threat analysis.  Valid values: `0`, `1`.
     * 
     */
    @Import(name="threatAnalysisSwitch")
    private @Nullable Output<String> threatAnalysisSwitch;

    /**
     * @return Threat analysis.  Valid values: `0`, `1`.
     * 
     */
    public Optional<Output<String>> threatAnalysisSwitch() {
        return Optional.ofNullable(this.threatAnalysisSwitch);
    }

    /**
     * Number of cores.
     * 
     */
    @Import(name="vCore")
    private @Nullable Output<String> vCore;

    /**
     * @return Number of cores.
     * 
     */
    public Optional<Output<String>> vCore() {
        return Optional.ofNullable(this.vCore);
    }

    /**
     * Version selection. Valid values: `level10`, `level2`, `level3`, `level7`, `level8`.
     * 
     */
    @Import(name="versionCode", required=true)
    private Output<String> versionCode;

    /**
     * @return Version selection. Valid values: `level10`, `level2`, `level3`, `level7`, `level8`.
     * 
     */
    public Output<String> versionCode() {
        return this.versionCode;
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.buyNumber = $.buyNumber;
        this.containerImageScan = $.containerImageScan;
        this.honeypot = $.honeypot;
        this.honeypotSwitch = $.honeypotSwitch;
        this.instanceId = $.instanceId;
        this.modifyType = $.modifyType;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.renewPeriod = $.renewPeriod;
        this.renewalPeriodUnit = $.renewalPeriodUnit;
        this.renewalStatus = $.renewalStatus;
        this.sasAntiRansomware = $.sasAntiRansomware;
        this.sasSc = $.sasSc;
        this.sasSdk = $.sasSdk;
        this.sasSdkSwitch = $.sasSdkSwitch;
        this.sasSlsStorage = $.sasSlsStorage;
        this.sasWebguardBoolean = $.sasWebguardBoolean;
        this.sasWebguardOrderNum = $.sasWebguardOrderNum;
        this.threatAnalysis = $.threatAnalysis;
        this.threatAnalysisSwitch = $.threatAnalysisSwitch;
        this.vCore = $.vCore;
        this.versionCode = $.versionCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param buyNumber Number of servers.
         * 
         * @return builder
         * 
         */
        public Builder buyNumber(@Nullable Output<String> buyNumber) {
            $.buyNumber = buyNumber;
            return this;
        }

        /**
         * @param buyNumber Number of servers.
         * 
         * @return builder
         * 
         */
        public Builder buyNumber(String buyNumber) {
            return buyNumber(Output.of(buyNumber));
        }

        /**
         * @param containerImageScan Container Image security scan.
         * 
         * @return builder
         * 
         */
        public Builder containerImageScan(@Nullable Output<String> containerImageScan) {
            $.containerImageScan = containerImageScan;
            return this;
        }

        /**
         * @param containerImageScan Container Image security scan.
         * 
         * @return builder
         * 
         */
        public Builder containerImageScan(String containerImageScan) {
            return containerImageScan(Output.of(containerImageScan));
        }

        /**
         * @param honeypot Cloud honeypot authorization number.
         * 
         * @return builder
         * 
         */
        public Builder honeypot(@Nullable Output<String> honeypot) {
            $.honeypot = honeypot;
            return this;
        }

        /**
         * @param honeypot Cloud honeypot authorization number.
         * 
         * @return builder
         * 
         */
        public Builder honeypot(String honeypot) {
            return honeypot(Output.of(honeypot));
        }

        /**
         * @param honeypotSwitch Cloud honeypot. Valid values: `1`, `2`.
         * 
         * @return builder
         * 
         */
        public Builder honeypotSwitch(@Nullable Output<String> honeypotSwitch) {
            $.honeypotSwitch = honeypotSwitch;
            return this;
        }

        /**
         * @param honeypotSwitch Cloud honeypot. Valid values: `1`, `2`.
         * 
         * @return builder
         * 
         */
        public Builder honeypotSwitch(String honeypotSwitch) {
            return honeypotSwitch(Output.of(honeypotSwitch));
        }

        /**
         * @param instanceId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param modifyType Change configuration type, value
         * - Upgrade: Upgrade.
         * - Downgrade: Downgrade.
         * 
         * @return builder
         * 
         */
        public Builder modifyType(@Nullable Output<String> modifyType) {
            $.modifyType = modifyType;
            return this;
        }

        /**
         * @param modifyType Change configuration type, value
         * - Upgrade: Upgrade.
         * - Downgrade: Downgrade.
         * 
         * @return builder
         * 
         */
        public Builder modifyType(String modifyType) {
            return modifyType(Output.of(modifyType));
        }

        /**
         * @param paymentType The payment type of the resource.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The payment type of the resource.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period Prepaid cycle. The unit is Monthly, please enter an integer multiple of 12 for annual paid products. **NOTE:** must be set when creating a prepaid instance.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period Prepaid cycle. The unit is Monthly, please enter an integer multiple of 12 for annual paid products. **NOTE:** must be set when creating a prepaid instance.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param renewPeriod Automatic renewal cycle, in months. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        /**
         * @param renewPeriod Automatic renewal cycle, in months. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        /**
         * @param renewalPeriodUnit The unit of the auto-renewal period. **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder renewalPeriodUnit(@Nullable Output<String> renewalPeriodUnit) {
            $.renewalPeriodUnit = renewalPeriodUnit;
            return this;
        }

        /**
         * @param renewalPeriodUnit The unit of the auto-renewal period. **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder renewalPeriodUnit(String renewalPeriodUnit) {
            return renewalPeriodUnit(Output.of(renewalPeriodUnit));
        }

        /**
         * @param renewalStatus Automatic renewal status, Default ManualRenewal. value:
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Automatic renewal status, Default ManualRenewal. value:
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        /**
         * @param sasAntiRansomware Anti-extortion.
         * 
         * @return builder
         * 
         */
        public Builder sasAntiRansomware(@Nullable Output<String> sasAntiRansomware) {
            $.sasAntiRansomware = sasAntiRansomware;
            return this;
        }

        /**
         * @param sasAntiRansomware Anti-extortion.
         * 
         * @return builder
         * 
         */
        public Builder sasAntiRansomware(String sasAntiRansomware) {
            return sasAntiRansomware(Output.of(sasAntiRansomware));
        }

        /**
         * @param sasSc Large security screen.
         * 
         * @return builder
         * 
         */
        public Builder sasSc(@Nullable Output<Boolean> sasSc) {
            $.sasSc = sasSc;
            return this;
        }

        /**
         * @param sasSc Large security screen.
         * 
         * @return builder
         * 
         */
        public Builder sasSc(Boolean sasSc) {
            return sasSc(Output.of(sasSc));
        }

        /**
         * @param sasSdk Number of malicious file detections.
         * 
         * @return builder
         * 
         */
        public Builder sasSdk(@Nullable Output<String> sasSdk) {
            $.sasSdk = sasSdk;
            return this;
        }

        /**
         * @param sasSdk Number of malicious file detections.
         * 
         * @return builder
         * 
         */
        public Builder sasSdk(String sasSdk) {
            return sasSdk(Output.of(sasSdk));
        }

        /**
         * @param sasSdkSwitch Malicious file detection SDK. Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder sasSdkSwitch(@Nullable Output<String> sasSdkSwitch) {
            $.sasSdkSwitch = sasSdkSwitch;
            return this;
        }

        /**
         * @param sasSdkSwitch Malicious file detection SDK. Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder sasSdkSwitch(String sasSdkSwitch) {
            return sasSdkSwitch(Output.of(sasSdkSwitch));
        }

        /**
         * @param sasSlsStorage Log analysis.
         * 
         * @return builder
         * 
         */
        public Builder sasSlsStorage(@Nullable Output<String> sasSlsStorage) {
            $.sasSlsStorage = sasSlsStorage;
            return this;
        }

        /**
         * @param sasSlsStorage Log analysis.
         * 
         * @return builder
         * 
         */
        public Builder sasSlsStorage(String sasSlsStorage) {
            return sasSlsStorage(Output.of(sasSlsStorage));
        }

        /**
         * @param sasWebguardBoolean Web page tamper-proof.  Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder sasWebguardBoolean(@Nullable Output<String> sasWebguardBoolean) {
            $.sasWebguardBoolean = sasWebguardBoolean;
            return this;
        }

        /**
         * @param sasWebguardBoolean Web page tamper-proof.  Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder sasWebguardBoolean(String sasWebguardBoolean) {
            return sasWebguardBoolean(Output.of(sasWebguardBoolean));
        }

        /**
         * @param sasWebguardOrderNum Number of tamper-proof authorizations.
         * 
         * @return builder
         * 
         */
        public Builder sasWebguardOrderNum(@Nullable Output<String> sasWebguardOrderNum) {
            $.sasWebguardOrderNum = sasWebguardOrderNum;
            return this;
        }

        /**
         * @param sasWebguardOrderNum Number of tamper-proof authorizations.
         * 
         * @return builder
         * 
         */
        public Builder sasWebguardOrderNum(String sasWebguardOrderNum) {
            return sasWebguardOrderNum(Output.of(sasWebguardOrderNum));
        }

        /**
         * @param threatAnalysis The amount of threat analysis log storage.
         * 
         * @return builder
         * 
         */
        public Builder threatAnalysis(@Nullable Output<String> threatAnalysis) {
            $.threatAnalysis = threatAnalysis;
            return this;
        }

        /**
         * @param threatAnalysis The amount of threat analysis log storage.
         * 
         * @return builder
         * 
         */
        public Builder threatAnalysis(String threatAnalysis) {
            return threatAnalysis(Output.of(threatAnalysis));
        }

        /**
         * @param threatAnalysisSwitch Threat analysis.  Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder threatAnalysisSwitch(@Nullable Output<String> threatAnalysisSwitch) {
            $.threatAnalysisSwitch = threatAnalysisSwitch;
            return this;
        }

        /**
         * @param threatAnalysisSwitch Threat analysis.  Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder threatAnalysisSwitch(String threatAnalysisSwitch) {
            return threatAnalysisSwitch(Output.of(threatAnalysisSwitch));
        }

        /**
         * @param vCore Number of cores.
         * 
         * @return builder
         * 
         */
        public Builder vCore(@Nullable Output<String> vCore) {
            $.vCore = vCore;
            return this;
        }

        /**
         * @param vCore Number of cores.
         * 
         * @return builder
         * 
         */
        public Builder vCore(String vCore) {
            return vCore(Output.of(vCore));
        }

        /**
         * @param versionCode Version selection. Valid values: `level10`, `level2`, `level3`, `level7`, `level8`.
         * 
         * @return builder
         * 
         */
        public Builder versionCode(Output<String> versionCode) {
            $.versionCode = versionCode;
            return this;
        }

        /**
         * @param versionCode Version selection. Valid values: `level10`, `level2`, `level3`, `level7`, `level8`.
         * 
         * @return builder
         * 
         */
        public Builder versionCode(String versionCode) {
            return versionCode(Output.of(versionCode));
        }

        public InstanceArgs build() {
            $.paymentType = Objects.requireNonNull($.paymentType, "expected parameter 'paymentType' to be non-null");
            $.versionCode = Objects.requireNonNull($.versionCode, "expected parameter 'versionCode' to be non-null");
            return $;
        }
    }

}