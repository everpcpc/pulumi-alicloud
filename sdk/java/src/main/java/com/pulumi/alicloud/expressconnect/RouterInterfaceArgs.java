// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouterInterfaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouterInterfaceArgs Empty = new RouterInterfaceArgs();

    /**
     * The access point ID to which the VBR belongs.
     * 
     */
    @Import(name="accessPointId")
    private @Nullable Output<String> accessPointId;

    /**
     * @return The access point ID to which the VBR belongs.
     * 
     */
    public Optional<Output<String>> accessPointId() {
        return Optional.ofNullable(this.accessPointId);
    }

    /**
     * Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
     * 
     */
    @Import(name="autoPay")
    private @Nullable Output<Boolean> autoPay;

    /**
     * @return Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
     * 
     */
    public Optional<Output<Boolean>> autoPay() {
        return Optional.ofNullable(this.autoPay);
    }

    /**
     * Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
     * 
     */
    @Import(name="deleteHealthCheckIp")
    private @Nullable Output<Boolean> deleteHealthCheckIp;

    /**
     * @return Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
     * 
     */
    public Optional<Output<Boolean>> deleteHealthCheckIp() {
        return Optional.ofNullable(this.deleteHealthCheckIp);
    }

    /**
     * The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
     * 
     */
    @Import(name="hcRate")
    private @Nullable Output<Integer> hcRate;

    /**
     * @return The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
     * 
     */
    public Optional<Output<Integer>> hcRate() {
        return Optional.ofNullable(this.hcRate);
    }

    /**
     * The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
     * 
     */
    @Import(name="hcThreshold")
    private @Nullable Output<String> hcThreshold;

    /**
     * @return The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
     * 
     */
    public Optional<Output<String>> hcThreshold() {
        return Optional.ofNullable(this.hcThreshold);
    }

    /**
     * The health check source IP address, must be an unused IP within the local VPC.
     * 
     */
    @Import(name="healthCheckSourceIp")
    private @Nullable Output<String> healthCheckSourceIp;

    /**
     * @return The health check source IP address, must be an unused IP within the local VPC.
     * 
     */
    public Optional<Output<String>> healthCheckSourceIp() {
        return Optional.ofNullable(this.healthCheckSourceIp);
    }

    /**
     * The IP address for health screening purposes.
     * 
     */
    @Import(name="healthCheckTargetIp")
    private @Nullable Output<String> healthCheckTargetIp;

    /**
     * @return The IP address for health screening purposes.
     * 
     */
    public Optional<Output<String>> healthCheckTargetIp() {
        return Optional.ofNullable(this.healthCheckTargetIp);
    }

    /**
     * The Access point ID to which the other end belongs.
     * 
     */
    @Import(name="oppositeAccessPointId")
    private @Nullable Output<String> oppositeAccessPointId;

    /**
     * @return The Access point ID to which the other end belongs.
     * 
     */
    public Optional<Output<String>> oppositeAccessPointId() {
        return Optional.ofNullable(this.oppositeAccessPointId);
    }

    /**
     * The Interface ID of the router at the other end.
     * 
     */
    @Import(name="oppositeInterfaceId")
    private @Nullable Output<String> oppositeInterfaceId;

    /**
     * @return The Interface ID of the router at the other end.
     * 
     */
    public Optional<Output<String>> oppositeInterfaceId() {
        return Optional.ofNullable(this.oppositeInterfaceId);
    }

    /**
     * The AliCloud account ID of the owner of the router interface on the other end.
     * 
     */
    @Import(name="oppositeInterfaceOwnerId")
    private @Nullable Output<String> oppositeInterfaceOwnerId;

    /**
     * @return The AliCloud account ID of the owner of the router interface on the other end.
     * 
     */
    public Optional<Output<String>> oppositeInterfaceOwnerId() {
        return Optional.ofNullable(this.oppositeInterfaceOwnerId);
    }

    /**
     * The geographical ID of the location of the receiving end of the connection.
     * 
     */
    @Import(name="oppositeRegionId", required=true)
    private Output<String> oppositeRegionId;

    /**
     * @return The geographical ID of the location of the receiving end of the connection.
     * 
     */
    public Output<String> oppositeRegionId() {
        return this.oppositeRegionId;
    }

    /**
     * The id of the router at the other end.
     * 
     */
    @Import(name="oppositeRouterId")
    private @Nullable Output<String> oppositeRouterId;

    /**
     * @return The id of the router at the other end.
     * 
     */
    public Optional<Output<String>> oppositeRouterId() {
        return Optional.ofNullable(this.oppositeRouterId);
    }

    /**
     * The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
     * 
     */
    @Import(name="oppositeRouterType")
    private @Nullable Output<String> oppositeRouterType;

    /**
     * @return The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
     * 
     */
    public Optional<Output<String>> oppositeRouterType() {
        return Optional.ofNullable(this.oppositeRouterType);
    }

    /**
     * The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
     * 
     */
    @Import(name="pricingCycle")
    private @Nullable Output<String> pricingCycle;

    /**
     * @return The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
     * 
     */
    public Optional<Output<String>> pricingCycle() {
        return Optional.ofNullable(this.pricingCycle);
    }

    /**
     * The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     * The router id associated with the router interface.
     * 
     */
    @Import(name="routerId", required=true)
    private Output<String> routerId;

    /**
     * @return The router id associated with the router interface.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }

    /**
     * The first ID of the resource.
     * 
     */
    @Import(name="routerInterfaceId")
    private @Nullable Output<String> routerInterfaceId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Optional<Output<String>> routerInterfaceId() {
        return Optional.ofNullable(this.routerInterfaceId);
    }

    /**
     * The name of the resource.
     * 
     */
    @Import(name="routerInterfaceName")
    private @Nullable Output<String> routerInterfaceName;

    /**
     * @return The name of the resource.
     * 
     */
    public Optional<Output<String>> routerInterfaceName() {
        return Optional.ofNullable(this.routerInterfaceName);
    }

    /**
     * The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
     * 
     */
    @Import(name="routerType", required=true)
    private Output<String> routerType;

    /**
     * @return The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
     * 
     */
    public Output<String> routerType() {
        return this.routerType;
    }

    /**
     * The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
     * 
     */
    @Import(name="spec", required=true)
    private Output<String> spec;

    /**
     * @return The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
     * 
     */
    public Output<String> spec() {
        return this.spec;
    }

    /**
     * The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private RouterInterfaceArgs() {}

    private RouterInterfaceArgs(RouterInterfaceArgs $) {
        this.accessPointId = $.accessPointId;
        this.autoPay = $.autoPay;
        this.deleteHealthCheckIp = $.deleteHealthCheckIp;
        this.description = $.description;
        this.hcRate = $.hcRate;
        this.hcThreshold = $.hcThreshold;
        this.healthCheckSourceIp = $.healthCheckSourceIp;
        this.healthCheckTargetIp = $.healthCheckTargetIp;
        this.oppositeAccessPointId = $.oppositeAccessPointId;
        this.oppositeInterfaceId = $.oppositeInterfaceId;
        this.oppositeInterfaceOwnerId = $.oppositeInterfaceOwnerId;
        this.oppositeRegionId = $.oppositeRegionId;
        this.oppositeRouterId = $.oppositeRouterId;
        this.oppositeRouterType = $.oppositeRouterType;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.pricingCycle = $.pricingCycle;
        this.role = $.role;
        this.routerId = $.routerId;
        this.routerInterfaceId = $.routerInterfaceId;
        this.routerInterfaceName = $.routerInterfaceName;
        this.routerType = $.routerType;
        this.spec = $.spec;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouterInterfaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouterInterfaceArgs $;

        public Builder() {
            $ = new RouterInterfaceArgs();
        }

        public Builder(RouterInterfaceArgs defaults) {
            $ = new RouterInterfaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessPointId The access point ID to which the VBR belongs.
         * 
         * @return builder
         * 
         */
        public Builder accessPointId(@Nullable Output<String> accessPointId) {
            $.accessPointId = accessPointId;
            return this;
        }

        /**
         * @param accessPointId The access point ID to which the VBR belongs.
         * 
         * @return builder
         * 
         */
        public Builder accessPointId(String accessPointId) {
            return accessPointId(Output.of(accessPointId));
        }

        /**
         * @param autoPay Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder autoPay(@Nullable Output<Boolean> autoPay) {
            $.autoPay = autoPay;
            return this;
        }

        /**
         * @param autoPay Whether to pay automatically, value:-**false** (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.-**true**: Enable automatic payment to automatically pay for orders.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder autoPay(Boolean autoPay) {
            return autoPay(Output.of(autoPay));
        }

        /**
         * @param deleteHealthCheckIp Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
         * 
         * @return builder
         * 
         */
        public Builder deleteHealthCheckIp(@Nullable Output<Boolean> deleteHealthCheckIp) {
            $.deleteHealthCheckIp = deleteHealthCheckIp;
            return this;
        }

        /**
         * @param deleteHealthCheckIp Whether to delete the health check IP address configured on the router interface. Value:-**true**: deletes the health check IP address.-**false** (default): does not delete the health check IP address.
         * 
         * @return builder
         * 
         */
        public Builder deleteHealthCheckIp(Boolean deleteHealthCheckIp) {
            return deleteHealthCheckIp(Output.of(deleteHealthCheckIp));
        }

        /**
         * @param description The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the router interface. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param hcRate The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
         * 
         * @return builder
         * 
         */
        public Builder hcRate(@Nullable Output<Integer> hcRate) {
            $.hcRate = hcRate;
            return this;
        }

        /**
         * @param hcRate The health check rate. Unit: seconds. The recommended value is 2. This indicates the interval between successive probe messages sent during the specified health check.
         * 
         * @return builder
         * 
         */
        public Builder hcRate(Integer hcRate) {
            return hcRate(Output.of(hcRate));
        }

        /**
         * @param hcThreshold The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
         * 
         * @return builder
         * 
         */
        public Builder hcThreshold(@Nullable Output<String> hcThreshold) {
            $.hcThreshold = hcThreshold;
            return this;
        }

        /**
         * @param hcThreshold The health check thresholds. Unit: pcs. The recommended value is 8. This indicates the number of probe messages to be sent during the specified health check.
         * 
         * @return builder
         * 
         */
        public Builder hcThreshold(String hcThreshold) {
            return hcThreshold(Output.of(hcThreshold));
        }

        /**
         * @param healthCheckSourceIp The health check source IP address, must be an unused IP within the local VPC.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckSourceIp(@Nullable Output<String> healthCheckSourceIp) {
            $.healthCheckSourceIp = healthCheckSourceIp;
            return this;
        }

        /**
         * @param healthCheckSourceIp The health check source IP address, must be an unused IP within the local VPC.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckSourceIp(String healthCheckSourceIp) {
            return healthCheckSourceIp(Output.of(healthCheckSourceIp));
        }

        /**
         * @param healthCheckTargetIp The IP address for health screening purposes.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTargetIp(@Nullable Output<String> healthCheckTargetIp) {
            $.healthCheckTargetIp = healthCheckTargetIp;
            return this;
        }

        /**
         * @param healthCheckTargetIp The IP address for health screening purposes.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTargetIp(String healthCheckTargetIp) {
            return healthCheckTargetIp(Output.of(healthCheckTargetIp));
        }

        /**
         * @param oppositeAccessPointId The Access point ID to which the other end belongs.
         * 
         * @return builder
         * 
         */
        public Builder oppositeAccessPointId(@Nullable Output<String> oppositeAccessPointId) {
            $.oppositeAccessPointId = oppositeAccessPointId;
            return this;
        }

        /**
         * @param oppositeAccessPointId The Access point ID to which the other end belongs.
         * 
         * @return builder
         * 
         */
        public Builder oppositeAccessPointId(String oppositeAccessPointId) {
            return oppositeAccessPointId(Output.of(oppositeAccessPointId));
        }

        /**
         * @param oppositeInterfaceId The Interface ID of the router at the other end.
         * 
         * @return builder
         * 
         */
        public Builder oppositeInterfaceId(@Nullable Output<String> oppositeInterfaceId) {
            $.oppositeInterfaceId = oppositeInterfaceId;
            return this;
        }

        /**
         * @param oppositeInterfaceId The Interface ID of the router at the other end.
         * 
         * @return builder
         * 
         */
        public Builder oppositeInterfaceId(String oppositeInterfaceId) {
            return oppositeInterfaceId(Output.of(oppositeInterfaceId));
        }

        /**
         * @param oppositeInterfaceOwnerId The AliCloud account ID of the owner of the router interface on the other end.
         * 
         * @return builder
         * 
         */
        public Builder oppositeInterfaceOwnerId(@Nullable Output<String> oppositeInterfaceOwnerId) {
            $.oppositeInterfaceOwnerId = oppositeInterfaceOwnerId;
            return this;
        }

        /**
         * @param oppositeInterfaceOwnerId The AliCloud account ID of the owner of the router interface on the other end.
         * 
         * @return builder
         * 
         */
        public Builder oppositeInterfaceOwnerId(String oppositeInterfaceOwnerId) {
            return oppositeInterfaceOwnerId(Output.of(oppositeInterfaceOwnerId));
        }

        /**
         * @param oppositeRegionId The geographical ID of the location of the receiving end of the connection.
         * 
         * @return builder
         * 
         */
        public Builder oppositeRegionId(Output<String> oppositeRegionId) {
            $.oppositeRegionId = oppositeRegionId;
            return this;
        }

        /**
         * @param oppositeRegionId The geographical ID of the location of the receiving end of the connection.
         * 
         * @return builder
         * 
         */
        public Builder oppositeRegionId(String oppositeRegionId) {
            return oppositeRegionId(Output.of(oppositeRegionId));
        }

        /**
         * @param oppositeRouterId The id of the router at the other end.
         * 
         * @return builder
         * 
         */
        public Builder oppositeRouterId(@Nullable Output<String> oppositeRouterId) {
            $.oppositeRouterId = oppositeRouterId;
            return this;
        }

        /**
         * @param oppositeRouterId The id of the router at the other end.
         * 
         * @return builder
         * 
         */
        public Builder oppositeRouterId(String oppositeRouterId) {
            return oppositeRouterId(Output.of(oppositeRouterId));
        }

        /**
         * @param oppositeRouterType The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
         * 
         * @return builder
         * 
         */
        public Builder oppositeRouterType(@Nullable Output<String> oppositeRouterType) {
            $.oppositeRouterType = oppositeRouterType;
            return this;
        }

        /**
         * @param oppositeRouterType The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
         * 
         * @return builder
         * 
         */
        public Builder oppositeRouterType(String oppositeRouterType) {
            return oppositeRouterType(Output.of(oppositeRouterType));
        }

        /**
         * @param paymentType The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The payment methods for router interfaces. Valid Values: `PayAsYouGo`, `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period Purchase duration, value:-When you choose to pay on a monthly basis, the value range is **1 to 9 * *.-When you choose to pay per year, the value range is **1 to 3 * *.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param pricingCycle The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder pricingCycle(@Nullable Output<String> pricingCycle) {
            $.pricingCycle = pricingCycle;
            return this;
        }

        /**
         * @param pricingCycle The billing cycle of the prepaid fee. Valid values:-**Month** (default): monthly payment.-**Year**: Pay per Year.&gt; **InstanceChargeType** is required when the value of the parameter is **PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder pricingCycle(String pricingCycle) {
            return pricingCycle(Output.of(pricingCycle));
        }

        /**
         * @param role The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param routerId The router id associated with the router interface.
         * 
         * @return builder
         * 
         */
        public Builder routerId(Output<String> routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerId The router id associated with the router interface.
         * 
         * @return builder
         * 
         */
        public Builder routerId(String routerId) {
            return routerId(Output.of(routerId));
        }

        /**
         * @param routerInterfaceId The first ID of the resource.
         * 
         * @return builder
         * 
         */
        public Builder routerInterfaceId(@Nullable Output<String> routerInterfaceId) {
            $.routerInterfaceId = routerInterfaceId;
            return this;
        }

        /**
         * @param routerInterfaceId The first ID of the resource.
         * 
         * @return builder
         * 
         */
        public Builder routerInterfaceId(String routerInterfaceId) {
            return routerInterfaceId(Output.of(routerInterfaceId));
        }

        /**
         * @param routerInterfaceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder routerInterfaceName(@Nullable Output<String> routerInterfaceName) {
            $.routerInterfaceName = routerInterfaceName;
            return this;
        }

        /**
         * @param routerInterfaceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder routerInterfaceName(String routerInterfaceName) {
            return routerInterfaceName(Output.of(routerInterfaceName));
        }

        /**
         * @param routerType The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
         * 
         * @return builder
         * 
         */
        public Builder routerType(Output<String> routerType) {
            $.routerType = routerType;
            return this;
        }

        /**
         * @param routerType The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
         * 
         * @return builder
         * 
         */
        public Builder routerType(String routerType) {
            return routerType(Output.of(routerType));
        }

        /**
         * @param spec The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
         * 
         * @return builder
         * 
         */
        public Builder spec(Output<String> spec) {
            $.spec = spec;
            return this;
        }

        /**
         * @param spec The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
         * 
         * @return builder
         * 
         */
        public Builder spec(String spec) {
            return spec(Output.of(spec));
        }

        /**
         * @param status The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public RouterInterfaceArgs build() {
            $.oppositeRegionId = Objects.requireNonNull($.oppositeRegionId, "expected parameter 'oppositeRegionId' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            $.routerId = Objects.requireNonNull($.routerId, "expected parameter 'routerId' to be non-null");
            $.routerType = Objects.requireNonNull($.routerType, "expected parameter 'routerType' to be non-null");
            $.spec = Objects.requireNonNull($.spec, "expected parameter 'spec' to be non-null");
            return $;
        }
    }

}
