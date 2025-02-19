// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEndpointAclServiceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEndpointAclServiceArgs Empty = new GetEndpointAclServiceArgs();

    /**
     * Whether to enable Acl Service.  Valid values: `true` and `false`.
     * 
     */
    @Import(name="enable", required=true)
    private Output<Boolean> enable;

    /**
     * @return Whether to enable Acl Service.  Valid values: `true` and `false`.
     * 
     */
    public Output<Boolean> enable() {
        return this.enable;
    }

    /**
     * The type of endpoint. Valid values: `internet`.
     * 
     */
    @Import(name="endpointType", required=true)
    private Output<String> endpointType;

    /**
     * @return The type of endpoint. Valid values: `internet`.
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }

    /**
     * The ID of the CR Instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the CR Instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The ModuleName. Valid values: `Registry`.
     * 
     * &gt; **NOTE:** After You enable access over the Internet, the Classless Inter-Domain Routing (CIDR) block `127.0.0.1/32` is automatically added to the whitelist.
     * 
     * &gt; **NOTE:** You may want to allow all ECS instances to access the Container Registry Enterprise Edition instance over the Internet. To achieve this purpose, you can enable access over the Internet and delete all IP addresses from the whitelist for Internet access. After you perform the preceding operation, the Container Registry Enterprise Edition instance is completely exposed to the Internet and may be attacked.
     * 
     */
    @Import(name="moduleName")
    private @Nullable Output<String> moduleName;

    /**
     * @return The ModuleName. Valid values: `Registry`.
     * 
     * &gt; **NOTE:** After You enable access over the Internet, the Classless Inter-Domain Routing (CIDR) block `127.0.0.1/32` is automatically added to the whitelist.
     * 
     * &gt; **NOTE:** You may want to allow all ECS instances to access the Container Registry Enterprise Edition instance over the Internet. To achieve this purpose, you can enable access over the Internet and delete all IP addresses from the whitelist for Internet access. After you perform the preceding operation, the Container Registry Enterprise Edition instance is completely exposed to the Internet and may be attacked.
     * 
     */
    public Optional<Output<String>> moduleName() {
        return Optional.ofNullable(this.moduleName);
    }

    private GetEndpointAclServiceArgs() {}

    private GetEndpointAclServiceArgs(GetEndpointAclServiceArgs $) {
        this.enable = $.enable;
        this.endpointType = $.endpointType;
        this.instanceId = $.instanceId;
        this.moduleName = $.moduleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEndpointAclServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEndpointAclServiceArgs $;

        public Builder() {
            $ = new GetEndpointAclServiceArgs();
        }

        public Builder(GetEndpointAclServiceArgs defaults) {
            $ = new GetEndpointAclServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enable Whether to enable Acl Service.  Valid values: `true` and `false`.
         * 
         * @return builder
         * 
         */
        public Builder enable(Output<Boolean> enable) {
            $.enable = enable;
            return this;
        }

        /**
         * @param enable Whether to enable Acl Service.  Valid values: `true` and `false`.
         * 
         * @return builder
         * 
         */
        public Builder enable(Boolean enable) {
            return enable(Output.of(enable));
        }

        /**
         * @param endpointType The type of endpoint. Valid values: `internet`.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(Output<String> endpointType) {
            $.endpointType = endpointType;
            return this;
        }

        /**
         * @param endpointType The type of endpoint. Valid values: `internet`.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(String endpointType) {
            return endpointType(Output.of(endpointType));
        }

        /**
         * @param instanceId The ID of the CR Instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the CR Instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param moduleName The ModuleName. Valid values: `Registry`.
         * 
         * &gt; **NOTE:** After You enable access over the Internet, the Classless Inter-Domain Routing (CIDR) block `127.0.0.1/32` is automatically added to the whitelist.
         * 
         * &gt; **NOTE:** You may want to allow all ECS instances to access the Container Registry Enterprise Edition instance over the Internet. To achieve this purpose, you can enable access over the Internet and delete all IP addresses from the whitelist for Internet access. After you perform the preceding operation, the Container Registry Enterprise Edition instance is completely exposed to the Internet and may be attacked.
         * 
         * @return builder
         * 
         */
        public Builder moduleName(@Nullable Output<String> moduleName) {
            $.moduleName = moduleName;
            return this;
        }

        /**
         * @param moduleName The ModuleName. Valid values: `Registry`.
         * 
         * &gt; **NOTE:** After You enable access over the Internet, the Classless Inter-Domain Routing (CIDR) block `127.0.0.1/32` is automatically added to the whitelist.
         * 
         * &gt; **NOTE:** You may want to allow all ECS instances to access the Container Registry Enterprise Edition instance over the Internet. To achieve this purpose, you can enable access over the Internet and delete all IP addresses from the whitelist for Internet access. After you perform the preceding operation, the Container Registry Enterprise Edition instance is completely exposed to the Internet and may be attacked.
         * 
         * @return builder
         * 
         */
        public Builder moduleName(String moduleName) {
            return moduleName(Output.of(moduleName));
        }

        public GetEndpointAclServiceArgs build() {
            $.enable = Objects.requireNonNull($.enable, "expected parameter 'enable' to be non-null");
            $.endpointType = Objects.requireNonNull($.endpointType, "expected parameter 'endpointType' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
