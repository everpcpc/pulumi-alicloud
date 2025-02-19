// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceState Empty = new InstanceState();

    /**
     * Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
     * 
     */
    @Import(name="bandWidth")
    private @Nullable Output<Integer> bandWidth;

    /**
     * @return Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
     * 
     */
    public Optional<Output<Integer>> bandWidth() {
        return Optional.ofNullable(this.bandWidth);
    }

    /**
     * Whether to use log audit. Valid values: `true`, `false`.
     * 
     */
    @Import(name="cfwLog")
    private @Nullable Output<Boolean> cfwLog;

    /**
     * @return Whether to use log audit. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> cfwLog() {
        return Optional.ofNullable(this.cfwLog);
    }

    /**
     * The log storage capacity.
     * 
     */
    @Import(name="cfwLogStorage")
    private @Nullable Output<Integer> cfwLogStorage;

    /**
     * @return The log storage capacity.
     * 
     */
    public Optional<Output<Integer>> cfwLogStorage() {
        return Optional.ofNullable(this.cfwLogStorage);
    }

    /**
     * Whether to use expert service. Valid values: `true`, `false`.
     * 
     */
    @Import(name="cfwService")
    private @Nullable Output<Boolean> cfwService;

    /**
     * @return Whether to use expert service. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> cfwService() {
        return Optional.ofNullable(this.cfwService);
    }

    /**
     * The creation time.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The end time.
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<String> endTime;

    /**
     * @return The end time.
     * 
     */
    public Optional<Output<String>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * The number of protected VPCs. Valid values between 2 and 500.
     * 
     */
    @Import(name="fwVpcNumber")
    private @Nullable Output<Integer> fwVpcNumber;

    /**
     * @return The number of protected VPCs. Valid values between 2 and 500.
     * 
     */
    public Optional<Output<Integer>> fwVpcNumber() {
        return Optional.ofNullable(this.fwVpcNumber);
    }

    /**
     * The number of assets.
     * 
     */
    @Import(name="instanceCount")
    private @Nullable Output<Integer> instanceCount;

    /**
     * @return The number of assets.
     * 
     */
    public Optional<Output<Integer>> instanceCount() {
        return Optional.ofNullable(this.instanceCount);
    }

    /**
     * The number of public IPs that can be protected. Valid values: 20 to 4000.
     * 
     */
    @Import(name="ipNumber")
    private @Nullable Output<Integer> ipNumber;

    /**
     * @return The number of public IPs that can be protected. Valid values: 20 to 4000.
     * 
     */
    public Optional<Output<Integer>> ipNumber() {
        return Optional.ofNullable(this.ipNumber);
    }

    /**
     * The logistics.
     * 
     */
    @Import(name="logistics")
    private @Nullable Output<String> logistics;

    /**
     * @return The logistics.
     * 
     */
    public Optional<Output<String>> logistics() {
        return Optional.ofNullable(this.logistics);
    }

    /**
     * The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
     * 
     */
    @Import(name="modifyType")
    private @Nullable Output<String> modifyType;

    /**
     * @return The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
     * 
     */
    public Optional<Output<String>> modifyType() {
        return Optional.ofNullable(this.modifyType);
    }

    /**
     * The payment type of the resource. Valid values: `Subscription`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `Subscription`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available in 1.204.1+.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available in 1.204.1+.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The release time.
     * 
     */
    @Import(name="releaseTime")
    private @Nullable Output<String> releaseTime;

    /**
     * @return The release time.
     * 
     */
    public Optional<Output<String>> releaseTime() {
        return Optional.ofNullable(this.releaseTime);
    }

    /**
     * Automatic renewal period. **NOTE:** The `renew_period` is required under the condition that renewal_status is `AutoRenewal`.
     * 
     */
    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    /**
     * @return Automatic renewal period. **NOTE:** The `renew_period` is required under the condition that renewal_status is `AutoRenewal`.
     * 
     */
    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    /**
     * Automatic renewal period unit. Valid values: `Month`,`Year`.
     * 
     */
    @Import(name="renewalDurationUnit")
    private @Nullable Output<String> renewalDurationUnit;

    /**
     * @return Automatic renewal period unit. Valid values: `Month`,`Year`.
     * 
     */
    public Optional<Output<String>> renewalDurationUnit() {
        return Optional.ofNullable(this.renewalDurationUnit);
    }

    /**
     * Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    /**
     * Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
     * 
     */
    @Import(name="spec")
    private @Nullable Output<String> spec;

    /**
     * @return Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
     * 
     */
    public Optional<Output<String>> spec() {
        return Optional.ofNullable(this.spec);
    }

    /**
     * The status of Instance.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of Instance.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private InstanceState() {}

    private InstanceState(InstanceState $) {
        this.bandWidth = $.bandWidth;
        this.cfwLog = $.cfwLog;
        this.cfwLogStorage = $.cfwLogStorage;
        this.cfwService = $.cfwService;
        this.createTime = $.createTime;
        this.endTime = $.endTime;
        this.fwVpcNumber = $.fwVpcNumber;
        this.instanceCount = $.instanceCount;
        this.ipNumber = $.ipNumber;
        this.logistics = $.logistics;
        this.modifyType = $.modifyType;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.releaseTime = $.releaseTime;
        this.renewPeriod = $.renewPeriod;
        this.renewalDurationUnit = $.renewalDurationUnit;
        this.renewalStatus = $.renewalStatus;
        this.spec = $.spec;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceState $;

        public Builder() {
            $ = new InstanceState();
        }

        public Builder(InstanceState defaults) {
            $ = new InstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bandWidth Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
         * 
         * @return builder
         * 
         */
        public Builder bandWidth(@Nullable Output<Integer> bandWidth) {
            $.bandWidth = bandWidth;
            return this;
        }

        /**
         * @param bandWidth Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
         * 
         * @return builder
         * 
         */
        public Builder bandWidth(Integer bandWidth) {
            return bandWidth(Output.of(bandWidth));
        }

        /**
         * @param cfwLog Whether to use log audit. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder cfwLog(@Nullable Output<Boolean> cfwLog) {
            $.cfwLog = cfwLog;
            return this;
        }

        /**
         * @param cfwLog Whether to use log audit. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder cfwLog(Boolean cfwLog) {
            return cfwLog(Output.of(cfwLog));
        }

        /**
         * @param cfwLogStorage The log storage capacity.
         * 
         * @return builder
         * 
         */
        public Builder cfwLogStorage(@Nullable Output<Integer> cfwLogStorage) {
            $.cfwLogStorage = cfwLogStorage;
            return this;
        }

        /**
         * @param cfwLogStorage The log storage capacity.
         * 
         * @return builder
         * 
         */
        public Builder cfwLogStorage(Integer cfwLogStorage) {
            return cfwLogStorage(Output.of(cfwLogStorage));
        }

        /**
         * @param cfwService Whether to use expert service. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder cfwService(@Nullable Output<Boolean> cfwService) {
            $.cfwService = cfwService;
            return this;
        }

        /**
         * @param cfwService Whether to use expert service. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder cfwService(Boolean cfwService) {
            return cfwService(Output.of(cfwService));
        }

        /**
         * @param createTime The creation time.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param endTime The end time.
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<String> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime The end time.
         * 
         * @return builder
         * 
         */
        public Builder endTime(String endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param fwVpcNumber The number of protected VPCs. Valid values between 2 and 500.
         * 
         * @return builder
         * 
         */
        public Builder fwVpcNumber(@Nullable Output<Integer> fwVpcNumber) {
            $.fwVpcNumber = fwVpcNumber;
            return this;
        }

        /**
         * @param fwVpcNumber The number of protected VPCs. Valid values between 2 and 500.
         * 
         * @return builder
         * 
         */
        public Builder fwVpcNumber(Integer fwVpcNumber) {
            return fwVpcNumber(Output.of(fwVpcNumber));
        }

        /**
         * @param instanceCount The number of assets.
         * 
         * @return builder
         * 
         */
        public Builder instanceCount(@Nullable Output<Integer> instanceCount) {
            $.instanceCount = instanceCount;
            return this;
        }

        /**
         * @param instanceCount The number of assets.
         * 
         * @return builder
         * 
         */
        public Builder instanceCount(Integer instanceCount) {
            return instanceCount(Output.of(instanceCount));
        }

        /**
         * @param ipNumber The number of public IPs that can be protected. Valid values: 20 to 4000.
         * 
         * @return builder
         * 
         */
        public Builder ipNumber(@Nullable Output<Integer> ipNumber) {
            $.ipNumber = ipNumber;
            return this;
        }

        /**
         * @param ipNumber The number of public IPs that can be protected. Valid values: 20 to 4000.
         * 
         * @return builder
         * 
         */
        public Builder ipNumber(Integer ipNumber) {
            return ipNumber(Output.of(ipNumber));
        }

        /**
         * @param logistics The logistics.
         * 
         * @return builder
         * 
         */
        public Builder logistics(@Nullable Output<String> logistics) {
            $.logistics = logistics;
            return this;
        }

        /**
         * @param logistics The logistics.
         * 
         * @return builder
         * 
         */
        public Builder logistics(String logistics) {
            return logistics(Output.of(logistics));
        }

        /**
         * @param modifyType The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
         * 
         * @return builder
         * 
         */
        public Builder modifyType(@Nullable Output<String> modifyType) {
            $.modifyType = modifyType;
            return this;
        }

        /**
         * @param modifyType The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
         * 
         * @return builder
         * 
         */
        public Builder modifyType(String modifyType) {
            return modifyType(Output.of(modifyType));
        }

        /**
         * @param paymentType The payment type of the resource. Valid values: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The payment type of the resource. Valid values: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available in 1.204.1+.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available in 1.204.1+.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param releaseTime The release time.
         * 
         * @return builder
         * 
         */
        public Builder releaseTime(@Nullable Output<String> releaseTime) {
            $.releaseTime = releaseTime;
            return this;
        }

        /**
         * @param releaseTime The release time.
         * 
         * @return builder
         * 
         */
        public Builder releaseTime(String releaseTime) {
            return releaseTime(Output.of(releaseTime));
        }

        /**
         * @param renewPeriod Automatic renewal period. **NOTE:** The `renew_period` is required under the condition that renewal_status is `AutoRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        /**
         * @param renewPeriod Automatic renewal period. **NOTE:** The `renew_period` is required under the condition that renewal_status is `AutoRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        /**
         * @param renewalDurationUnit Automatic renewal period unit. Valid values: `Month`,`Year`.
         * 
         * @return builder
         * 
         */
        public Builder renewalDurationUnit(@Nullable Output<String> renewalDurationUnit) {
            $.renewalDurationUnit = renewalDurationUnit;
            return this;
        }

        /**
         * @param renewalDurationUnit Automatic renewal period unit. Valid values: `Month`,`Year`.
         * 
         * @return builder
         * 
         */
        public Builder renewalDurationUnit(String renewalDurationUnit) {
            return renewalDurationUnit(Output.of(renewalDurationUnit));
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        /**
         * @param spec Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
         * 
         * @return builder
         * 
         */
        public Builder spec(@Nullable Output<String> spec) {
            $.spec = spec;
            return this;
        }

        /**
         * @param spec Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
         * 
         * @return builder
         * 
         */
        public Builder spec(String spec) {
            return spec(Output.of(spec));
        }

        /**
         * @param status The status of Instance.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of Instance.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public InstanceState build() {
            return $;
        }
    }

}
