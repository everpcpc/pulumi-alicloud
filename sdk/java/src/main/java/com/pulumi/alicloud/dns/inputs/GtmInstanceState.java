// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.inputs;

import com.pulumi.alicloud.dns.inputs.GtmInstanceAlertConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GtmInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final GtmInstanceState Empty = new GtmInstanceState();

    /**
     * The alert notification methods. See the following `Block alert_config`.
     * 
     */
    @Import(name="alertConfigs")
    private @Nullable Output<List<GtmInstanceAlertConfigArgs>> alertConfigs;

    /**
     * @return The alert notification methods. See the following `Block alert_config`.
     * 
     */
    public Optional<Output<List<GtmInstanceAlertConfigArgs>>> alertConfigs() {
        return Optional.ofNullable(this.alertConfigs);
    }

    /**
     * The alert group.
     * 
     */
    @Import(name="alertGroups")
    private @Nullable Output<List<String>> alertGroups;

    /**
     * @return The alert group.
     * 
     */
    public Optional<Output<List<String>>> alertGroups() {
        return Optional.ofNullable(this.alertGroups);
    }

    /**
     * The access type of the CNAME domain name. Valid value: `PUBLIC`.
     * 
     */
    @Import(name="cnameType")
    private @Nullable Output<String> cnameType;

    /**
     * @return The access type of the CNAME domain name. Valid value: `PUBLIC`.
     * 
     */
    public Optional<Output<String>> cnameType() {
        return Optional.ofNullable(this.cnameType);
    }

    /**
     * The force update.
     * 
     */
    @Import(name="forceUpdate")
    private @Nullable Output<Boolean> forceUpdate;

    /**
     * @return The force update.
     * 
     */
    public Optional<Output<Boolean>> forceUpdate() {
        return Optional.ofNullable(this.forceUpdate);
    }

    /**
     * The quota of detection tasks.
     * 
     */
    @Import(name="healthCheckTaskCount")
    private @Nullable Output<Integer> healthCheckTaskCount;

    /**
     * @return The quota of detection tasks.
     * 
     */
    public Optional<Output<Integer>> healthCheckTaskCount() {
        return Optional.ofNullable(this.healthCheckTaskCount);
    }

    /**
     * The name of the instance.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return The name of the instance.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * The lang.
     * 
     */
    @Import(name="lang")
    private @Nullable Output<String> lang;

    /**
     * @return The lang.
     * 
     */
    public Optional<Output<String>> lang() {
        return Optional.ofNullable(this.lang);
    }

    /**
     * Paid package version. Valid values: `ultimate`, `standard`.
     * 
     */
    @Import(name="packageEdition")
    private @Nullable Output<String> packageEdition;

    /**
     * @return Paid package version. Valid values: `ultimate`, `standard`.
     * 
     */
    public Optional<Output<String>> packageEdition() {
        return Optional.ofNullable(this.packageEdition);
    }

    /**
     * The Payment Type of the resource. Valid value: `Subscription`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The Payment Type of the resource. Valid value: `Subscription`.
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
     * The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
     * 
     */
    @Import(name="publicCnameMode")
    private @Nullable Output<String> publicCnameMode;

    /**
     * @return The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
     * 
     */
    public Optional<Output<String>> publicCnameMode() {
        return Optional.ofNullable(this.publicCnameMode);
    }

    /**
     * The CNAME access domain name.
     * 
     */
    @Import(name="publicRr")
    private @Nullable Output<String> publicRr;

    /**
     * @return The CNAME access domain name.
     * 
     */
    public Optional<Output<String>> publicRr() {
        return Optional.ofNullable(this.publicRr);
    }

    /**
     * The website domain name that the user uses on the Internet.
     * 
     */
    @Import(name="publicUserDomainName")
    private @Nullable Output<String> publicUserDomainName;

    /**
     * @return The website domain name that the user uses on the Internet.
     * 
     */
    public Optional<Output<String>> publicUserDomainName() {
        return Optional.ofNullable(this.publicUserDomainName);
    }

    /**
     * The domain name that is used to access GTM over the Internet.
     * 
     */
    @Import(name="publicZoneName")
    private @Nullable Output<String> publicZoneName;

    /**
     * @return The domain name that is used to access GTM over the Internet.
     * 
     */
    public Optional<Output<String>> publicZoneName() {
        return Optional.ofNullable(this.publicZoneName);
    }

    /**
     * Automatic renewal period, the unit is month. When setting `renewal_status` to AutoRenewal, it must be set.
     * 
     */
    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    /**
     * @return Automatic renewal period, the unit is month. When setting `renewal_status` to AutoRenewal, it must be set.
     * 
     */
    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The quota of SMS notifications.
     * 
     */
    @Import(name="smsNotificationCount")
    private @Nullable Output<Integer> smsNotificationCount;

    /**
     * @return The quota of SMS notifications.
     * 
     */
    public Optional<Output<Integer>> smsNotificationCount() {
        return Optional.ofNullable(this.smsNotificationCount);
    }

    /**
     * The type of the access policy. Valid values: `GEO`, `LATENCY`.
     * 
     */
    @Import(name="strategyMode")
    private @Nullable Output<String> strategyMode;

    /**
     * @return The type of the access policy. Valid values: `GEO`, `LATENCY`.
     * 
     */
    public Optional<Output<String>> strategyMode() {
        return Optional.ofNullable(this.strategyMode);
    }

    /**
     * The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private GtmInstanceState() {}

    private GtmInstanceState(GtmInstanceState $) {
        this.alertConfigs = $.alertConfigs;
        this.alertGroups = $.alertGroups;
        this.cnameType = $.cnameType;
        this.forceUpdate = $.forceUpdate;
        this.healthCheckTaskCount = $.healthCheckTaskCount;
        this.instanceName = $.instanceName;
        this.lang = $.lang;
        this.packageEdition = $.packageEdition;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.publicCnameMode = $.publicCnameMode;
        this.publicRr = $.publicRr;
        this.publicUserDomainName = $.publicUserDomainName;
        this.publicZoneName = $.publicZoneName;
        this.renewPeriod = $.renewPeriod;
        this.renewalStatus = $.renewalStatus;
        this.resourceGroupId = $.resourceGroupId;
        this.smsNotificationCount = $.smsNotificationCount;
        this.strategyMode = $.strategyMode;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmInstanceState $;

        public Builder() {
            $ = new GtmInstanceState();
        }

        public Builder(GtmInstanceState defaults) {
            $ = new GtmInstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertConfigs The alert notification methods. See the following `Block alert_config`.
         * 
         * @return builder
         * 
         */
        public Builder alertConfigs(@Nullable Output<List<GtmInstanceAlertConfigArgs>> alertConfigs) {
            $.alertConfigs = alertConfigs;
            return this;
        }

        /**
         * @param alertConfigs The alert notification methods. See the following `Block alert_config`.
         * 
         * @return builder
         * 
         */
        public Builder alertConfigs(List<GtmInstanceAlertConfigArgs> alertConfigs) {
            return alertConfigs(Output.of(alertConfigs));
        }

        /**
         * @param alertConfigs The alert notification methods. See the following `Block alert_config`.
         * 
         * @return builder
         * 
         */
        public Builder alertConfigs(GtmInstanceAlertConfigArgs... alertConfigs) {
            return alertConfigs(List.of(alertConfigs));
        }

        /**
         * @param alertGroups The alert group.
         * 
         * @return builder
         * 
         */
        public Builder alertGroups(@Nullable Output<List<String>> alertGroups) {
            $.alertGroups = alertGroups;
            return this;
        }

        /**
         * @param alertGroups The alert group.
         * 
         * @return builder
         * 
         */
        public Builder alertGroups(List<String> alertGroups) {
            return alertGroups(Output.of(alertGroups));
        }

        /**
         * @param alertGroups The alert group.
         * 
         * @return builder
         * 
         */
        public Builder alertGroups(String... alertGroups) {
            return alertGroups(List.of(alertGroups));
        }

        /**
         * @param cnameType The access type of the CNAME domain name. Valid value: `PUBLIC`.
         * 
         * @return builder
         * 
         */
        public Builder cnameType(@Nullable Output<String> cnameType) {
            $.cnameType = cnameType;
            return this;
        }

        /**
         * @param cnameType The access type of the CNAME domain name. Valid value: `PUBLIC`.
         * 
         * @return builder
         * 
         */
        public Builder cnameType(String cnameType) {
            return cnameType(Output.of(cnameType));
        }

        /**
         * @param forceUpdate The force update.
         * 
         * @return builder
         * 
         */
        public Builder forceUpdate(@Nullable Output<Boolean> forceUpdate) {
            $.forceUpdate = forceUpdate;
            return this;
        }

        /**
         * @param forceUpdate The force update.
         * 
         * @return builder
         * 
         */
        public Builder forceUpdate(Boolean forceUpdate) {
            return forceUpdate(Output.of(forceUpdate));
        }

        /**
         * @param healthCheckTaskCount The quota of detection tasks.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTaskCount(@Nullable Output<Integer> healthCheckTaskCount) {
            $.healthCheckTaskCount = healthCheckTaskCount;
            return this;
        }

        /**
         * @param healthCheckTaskCount The quota of detection tasks.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTaskCount(Integer healthCheckTaskCount) {
            return healthCheckTaskCount(Output.of(healthCheckTaskCount));
        }

        /**
         * @param instanceName The name of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The name of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param lang The lang.
         * 
         * @return builder
         * 
         */
        public Builder lang(@Nullable Output<String> lang) {
            $.lang = lang;
            return this;
        }

        /**
         * @param lang The lang.
         * 
         * @return builder
         * 
         */
        public Builder lang(String lang) {
            return lang(Output.of(lang));
        }

        /**
         * @param packageEdition Paid package version. Valid values: `ultimate`, `standard`.
         * 
         * @return builder
         * 
         */
        public Builder packageEdition(@Nullable Output<String> packageEdition) {
            $.packageEdition = packageEdition;
            return this;
        }

        /**
         * @param packageEdition Paid package version. Valid values: `ultimate`, `standard`.
         * 
         * @return builder
         * 
         */
        public Builder packageEdition(String packageEdition) {
            return packageEdition(Output.of(packageEdition));
        }

        /**
         * @param paymentType The Payment Type of the resource. Valid value: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The Payment Type of the resource. Valid value: `Subscription`.
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
         * @param publicCnameMode The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
         * 
         * @return builder
         * 
         */
        public Builder publicCnameMode(@Nullable Output<String> publicCnameMode) {
            $.publicCnameMode = publicCnameMode;
            return this;
        }

        /**
         * @param publicCnameMode The Public Network domain name access method. Valid values: `CUSTOM`, `SYSTEM_ASSIGN`.
         * 
         * @return builder
         * 
         */
        public Builder publicCnameMode(String publicCnameMode) {
            return publicCnameMode(Output.of(publicCnameMode));
        }

        /**
         * @param publicRr The CNAME access domain name.
         * 
         * @return builder
         * 
         */
        public Builder publicRr(@Nullable Output<String> publicRr) {
            $.publicRr = publicRr;
            return this;
        }

        /**
         * @param publicRr The CNAME access domain name.
         * 
         * @return builder
         * 
         */
        public Builder publicRr(String publicRr) {
            return publicRr(Output.of(publicRr));
        }

        /**
         * @param publicUserDomainName The website domain name that the user uses on the Internet.
         * 
         * @return builder
         * 
         */
        public Builder publicUserDomainName(@Nullable Output<String> publicUserDomainName) {
            $.publicUserDomainName = publicUserDomainName;
            return this;
        }

        /**
         * @param publicUserDomainName The website domain name that the user uses on the Internet.
         * 
         * @return builder
         * 
         */
        public Builder publicUserDomainName(String publicUserDomainName) {
            return publicUserDomainName(Output.of(publicUserDomainName));
        }

        /**
         * @param publicZoneName The domain name that is used to access GTM over the Internet.
         * 
         * @return builder
         * 
         */
        public Builder publicZoneName(@Nullable Output<String> publicZoneName) {
            $.publicZoneName = publicZoneName;
            return this;
        }

        /**
         * @param publicZoneName The domain name that is used to access GTM over the Internet.
         * 
         * @return builder
         * 
         */
        public Builder publicZoneName(String publicZoneName) {
            return publicZoneName(Output.of(publicZoneName));
        }

        /**
         * @param renewPeriod Automatic renewal period, the unit is month. When setting `renewal_status` to AutoRenewal, it must be set.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        /**
         * @param renewPeriod Automatic renewal period, the unit is month. When setting `renewal_status` to AutoRenewal, it must be set.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param smsNotificationCount The quota of SMS notifications.
         * 
         * @return builder
         * 
         */
        public Builder smsNotificationCount(@Nullable Output<Integer> smsNotificationCount) {
            $.smsNotificationCount = smsNotificationCount;
            return this;
        }

        /**
         * @param smsNotificationCount The quota of SMS notifications.
         * 
         * @return builder
         * 
         */
        public Builder smsNotificationCount(Integer smsNotificationCount) {
            return smsNotificationCount(Output.of(smsNotificationCount));
        }

        /**
         * @param strategyMode The type of the access policy. Valid values: `GEO`, `LATENCY`.
         * 
         * @return builder
         * 
         */
        public Builder strategyMode(@Nullable Output<String> strategyMode) {
            $.strategyMode = strategyMode;
            return this;
        }

        /**
         * @param strategyMode The type of the access policy. Valid values: `GEO`, `LATENCY`.
         * 
         * @return builder
         * 
         */
        public Builder strategyMode(String strategyMode) {
            return strategyMode(Output.of(strategyMode));
        }

        /**
         * @param ttl The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The global time to live. Valid values: `60`, `120`, `300`, `600`. Unit: second.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        public GtmInstanceState build() {
            return $;
        }
    }

}
