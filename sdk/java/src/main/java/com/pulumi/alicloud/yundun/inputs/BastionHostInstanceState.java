// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.yundun.inputs;

import com.pulumi.alicloud.yundun.inputs.BastionHostInstanceAdAuthServerArgs;
import com.pulumi.alicloud.yundun.inputs.BastionHostInstanceLdapAuthServerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BastionHostInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final BastionHostInstanceState Empty = new BastionHostInstanceState();

    @Import(name="adAuthServers")
    private @Nullable Output<List<BastionHostInstanceAdAuthServerArgs>> adAuthServers;

    public Optional<Output<List<BastionHostInstanceAdAuthServerArgs>>> adAuthServers() {
        return Optional.ofNullable(this.adAuthServers);
    }

    @Import(name="bandwidth")
    private @Nullable Output<String> bandwidth;

    public Optional<Output<String>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="enablePublicAccess")
    private @Nullable Output<Boolean> enablePublicAccess;

    public Optional<Output<Boolean>> enablePublicAccess() {
        return Optional.ofNullable(this.enablePublicAccess);
    }

    @Import(name="ldapAuthServers")
    private @Nullable Output<List<BastionHostInstanceLdapAuthServerArgs>> ldapAuthServers;

    public Optional<Output<List<BastionHostInstanceLdapAuthServerArgs>>> ldapAuthServers() {
        return Optional.ofNullable(this.ldapAuthServers);
    }

    @Import(name="licenseCode")
    private @Nullable Output<String> licenseCode;

    public Optional<Output<String>> licenseCode() {
        return Optional.ofNullable(this.licenseCode);
    }

    @Import(name="period")
    private @Nullable Output<Integer> period;

    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    @Import(name="planCode")
    private @Nullable Output<String> planCode;

    public Optional<Output<String>> planCode() {
        return Optional.ofNullable(this.planCode);
    }

    @Import(name="publicWhiteLists")
    private @Nullable Output<List<String>> publicWhiteLists;

    public Optional<Output<List<String>>> publicWhiteLists() {
        return Optional.ofNullable(this.publicWhiteLists);
    }

    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    @Import(name="renewalPeriodUnit")
    private @Nullable Output<String> renewalPeriodUnit;

    public Optional<Output<String>> renewalPeriodUnit() {
        return Optional.ofNullable(this.renewalPeriodUnit);
    }

    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    @Import(name="storage")
    private @Nullable Output<String> storage;

    public Optional<Output<String>> storage() {
        return Optional.ofNullable(this.storage);
    }

    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private BastionHostInstanceState() {}

    private BastionHostInstanceState(BastionHostInstanceState $) {
        this.adAuthServers = $.adAuthServers;
        this.bandwidth = $.bandwidth;
        this.description = $.description;
        this.enablePublicAccess = $.enablePublicAccess;
        this.ldapAuthServers = $.ldapAuthServers;
        this.licenseCode = $.licenseCode;
        this.period = $.period;
        this.planCode = $.planCode;
        this.publicWhiteLists = $.publicWhiteLists;
        this.renewPeriod = $.renewPeriod;
        this.renewalPeriodUnit = $.renewalPeriodUnit;
        this.renewalStatus = $.renewalStatus;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupIds = $.securityGroupIds;
        this.storage = $.storage;
        this.tags = $.tags;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BastionHostInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BastionHostInstanceState $;

        public Builder() {
            $ = new BastionHostInstanceState();
        }

        public Builder(BastionHostInstanceState defaults) {
            $ = new BastionHostInstanceState(Objects.requireNonNull(defaults));
        }

        public Builder adAuthServers(@Nullable Output<List<BastionHostInstanceAdAuthServerArgs>> adAuthServers) {
            $.adAuthServers = adAuthServers;
            return this;
        }

        public Builder adAuthServers(List<BastionHostInstanceAdAuthServerArgs> adAuthServers) {
            return adAuthServers(Output.of(adAuthServers));
        }

        public Builder adAuthServers(BastionHostInstanceAdAuthServerArgs... adAuthServers) {
            return adAuthServers(List.of(adAuthServers));
        }

        public Builder bandwidth(@Nullable Output<String> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        public Builder bandwidth(String bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder enablePublicAccess(@Nullable Output<Boolean> enablePublicAccess) {
            $.enablePublicAccess = enablePublicAccess;
            return this;
        }

        public Builder enablePublicAccess(Boolean enablePublicAccess) {
            return enablePublicAccess(Output.of(enablePublicAccess));
        }

        public Builder ldapAuthServers(@Nullable Output<List<BastionHostInstanceLdapAuthServerArgs>> ldapAuthServers) {
            $.ldapAuthServers = ldapAuthServers;
            return this;
        }

        public Builder ldapAuthServers(List<BastionHostInstanceLdapAuthServerArgs> ldapAuthServers) {
            return ldapAuthServers(Output.of(ldapAuthServers));
        }

        public Builder ldapAuthServers(BastionHostInstanceLdapAuthServerArgs... ldapAuthServers) {
            return ldapAuthServers(List.of(ldapAuthServers));
        }

        public Builder licenseCode(@Nullable Output<String> licenseCode) {
            $.licenseCode = licenseCode;
            return this;
        }

        public Builder licenseCode(String licenseCode) {
            return licenseCode(Output.of(licenseCode));
        }

        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        public Builder planCode(@Nullable Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        public Builder publicWhiteLists(@Nullable Output<List<String>> publicWhiteLists) {
            $.publicWhiteLists = publicWhiteLists;
            return this;
        }

        public Builder publicWhiteLists(List<String> publicWhiteLists) {
            return publicWhiteLists(Output.of(publicWhiteLists));
        }

        public Builder publicWhiteLists(String... publicWhiteLists) {
            return publicWhiteLists(List.of(publicWhiteLists));
        }

        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        public Builder renewalPeriodUnit(@Nullable Output<String> renewalPeriodUnit) {
            $.renewalPeriodUnit = renewalPeriodUnit;
            return this;
        }

        public Builder renewalPeriodUnit(String renewalPeriodUnit) {
            return renewalPeriodUnit(Output.of(renewalPeriodUnit));
        }

        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        public Builder storage(@Nullable Output<String> storage) {
            $.storage = storage;
            return this;
        }

        public Builder storage(String storage) {
            return storage(Output.of(storage));
        }

        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public BastionHostInstanceState build() {
            return $;
        }
    }

}
