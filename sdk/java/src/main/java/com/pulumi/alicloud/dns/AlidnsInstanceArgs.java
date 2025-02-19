// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlidnsInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlidnsInstanceArgs Empty = new AlidnsInstanceArgs();

    /**
     * Alidns security level. Valid values: `no`, `basic`, `advanced`.
     * 
     */
    @Import(name="dnsSecurity", required=true)
    private Output<String> dnsSecurity;

    /**
     * @return Alidns security level. Valid values: `no`, `basic`, `advanced`.
     * 
     */
    public Output<String> dnsSecurity() {
        return this.dnsSecurity;
    }

    /**
     * Number of domain names bound.
     * 
     */
    @Import(name="domainNumbers", required=true)
    private Output<String> domainNumbers;

    /**
     * @return Number of domain names bound.
     * 
     */
    public Output<String> domainNumbers() {
        return this.domainNumbers;
    }

    /**
     * The billing method of the Alidns instance. Valid values: `Subscription`. Default to `Subscription`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The billing method of the Alidns instance. Valid values: `Subscription`. Default to `Subscription`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
     * 
     */
    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    /**
     * @return Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
     * 
     */
    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    /**
     * Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
     * 
     */
    @Import(name="versionCode", required=true)
    private Output<String> versionCode;

    /**
     * @return Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
     * 
     */
    public Output<String> versionCode() {
        return this.versionCode;
    }

    private AlidnsInstanceArgs() {}

    private AlidnsInstanceArgs(AlidnsInstanceArgs $) {
        this.dnsSecurity = $.dnsSecurity;
        this.domainNumbers = $.domainNumbers;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.renewPeriod = $.renewPeriod;
        this.renewalStatus = $.renewalStatus;
        this.versionCode = $.versionCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlidnsInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlidnsInstanceArgs $;

        public Builder() {
            $ = new AlidnsInstanceArgs();
        }

        public Builder(AlidnsInstanceArgs defaults) {
            $ = new AlidnsInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dnsSecurity Alidns security level. Valid values: `no`, `basic`, `advanced`.
         * 
         * @return builder
         * 
         */
        public Builder dnsSecurity(Output<String> dnsSecurity) {
            $.dnsSecurity = dnsSecurity;
            return this;
        }

        /**
         * @param dnsSecurity Alidns security level. Valid values: `no`, `basic`, `advanced`.
         * 
         * @return builder
         * 
         */
        public Builder dnsSecurity(String dnsSecurity) {
            return dnsSecurity(Output.of(dnsSecurity));
        }

        /**
         * @param domainNumbers Number of domain names bound.
         * 
         * @return builder
         * 
         */
        public Builder domainNumbers(Output<String> domainNumbers) {
            $.domainNumbers = domainNumbers;
            return this;
        }

        /**
         * @param domainNumbers Number of domain names bound.
         * 
         * @return builder
         * 
         */
        public Builder domainNumbers(String domainNumbers) {
            return domainNumbers(Output.of(domainNumbers));
        }

        /**
         * @param paymentType The billing method of the Alidns instance. Valid values: `Subscription`. Default to `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The billing method of the Alidns instance. Valid values: `Subscription`. Default to `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param renewPeriod Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        /**
         * @param renewPeriod Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        /**
         * @param versionCode Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
         * 
         * @return builder
         * 
         */
        public Builder versionCode(Output<String> versionCode) {
            $.versionCode = versionCode;
            return this;
        }

        /**
         * @param versionCode Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
         * 
         * @return builder
         * 
         */
        public Builder versionCode(String versionCode) {
            return versionCode(Output.of(versionCode));
        }

        public AlidnsInstanceArgs build() {
            $.dnsSecurity = Objects.requireNonNull($.dnsSecurity, "expected parameter 'dnsSecurity' to be non-null");
            $.domainNumbers = Objects.requireNonNull($.domainNumbers, "expected parameter 'domainNumbers' to be non-null");
            $.versionCode = Objects.requireNonNull($.versionCode, "expected parameter 'versionCode' to be non-null");
            return $;
        }
    }

}
