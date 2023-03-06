// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceLinkedRoleState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceLinkedRoleState Empty = new ServiceLinkedRoleState();

    /**
     * The product name for SLR. EventBridge can automatically create the following service-linked roles:
     * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
     * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
     * 
     */
    @Import(name="productName")
    private @Nullable Output<String> productName;

    /**
     * @return The product name for SLR. EventBridge can automatically create the following service-linked roles:
     * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
     * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
     * 
     */
    public Optional<Output<String>> productName() {
        return Optional.ofNullable(this.productName);
    }

    private ServiceLinkedRoleState() {}

    private ServiceLinkedRoleState(ServiceLinkedRoleState $) {
        this.productName = $.productName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceLinkedRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceLinkedRoleState $;

        public Builder() {
            $ = new ServiceLinkedRoleState();
        }

        public Builder(ServiceLinkedRoleState defaults) {
            $ = new ServiceLinkedRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param productName The product name for SLR. EventBridge can automatically create the following service-linked roles:
         * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
         * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
         * 
         * @return builder
         * 
         */
        public Builder productName(@Nullable Output<String> productName) {
            $.productName = productName;
            return this;
        }

        /**
         * @param productName The product name for SLR. EventBridge can automatically create the following service-linked roles:
         * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
         * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
         * 
         * @return builder
         * 
         */
        public Builder productName(String productName) {
            return productName(Output.of(productName));
        }

        public ServiceLinkedRoleState build() {
            return $;
        }
    }

}