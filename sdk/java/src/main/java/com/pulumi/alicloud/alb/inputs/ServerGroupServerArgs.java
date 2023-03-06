// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerGroupServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerGroupServerArgs Empty = new ServerGroupServerArgs();

    /**
     * The description of the server.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the server.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The port that is used by the server. Valid values: `1` to `65535`. **Note:** This parameter is required if the `server_type` parameter is set to `Ecs`, `Eni`, `Eci`, or `Ip`. You do not need to configure this parameter if you set `server_type` to `Fc`.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port that is used by the server. Valid values: `1` to `65535`. **Note:** This parameter is required if the `server_type` parameter is set to `Ecs`, `Eni`, `Eci`, or `Ip`. You do not need to configure this parameter if you set `server_type` to `Fc`.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Specifies whether to enable the remote IP address feature. You can specify up to 40 servers in each call. **Note:** If `server_type` is set to `Ip`, this parameter is available.
     * 
     */
    @Import(name="remoteIpEnabled")
    private @Nullable Output<Boolean> remoteIpEnabled;

    /**
     * @return Specifies whether to enable the remote IP address feature. You can specify up to 40 servers in each call. **Note:** If `server_type` is set to `Ip`, this parameter is available.
     * 
     */
    public Optional<Output<Boolean>> remoteIpEnabled() {
        return Optional.ofNullable(this.remoteIpEnabled);
    }

    /**
     * The ID of the backend server.
     * - If `server_group_type` is set to `Instance`, set the parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by Ecs, Eni, or Eci.
     * - If `server_group_type` is set to `Ip`, set the parameter to an IP address specified in the server group.
     * - If `server_group_type` is set to `Fc`, set the parameter to the Alibaba Cloud Resource Name (ARN) of a function specified in the server group.
     * 
     */
    @Import(name="serverId", required=true)
    private Output<String> serverId;

    /**
     * @return The ID of the backend server.
     * - If `server_group_type` is set to `Instance`, set the parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by Ecs, Eni, or Eci.
     * - If `server_group_type` is set to `Ip`, set the parameter to an IP address specified in the server group.
     * - If `server_group_type` is set to `Fc`, set the parameter to the Alibaba Cloud Resource Name (ARN) of a function specified in the server group.
     * 
     */
    public Output<String> serverId() {
        return this.serverId;
    }

    /**
     * The IP address of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. **Note:** If `server_group_type` is set to `Fc`, you do not need to configure parameters, otherwise this attribute is required. If `server_group_type` is set to `Ip`, the value of this property is the same as the `server_id` value.
     * 
     */
    @Import(name="serverIp")
    private @Nullable Output<String> serverIp;

    /**
     * @return The IP address of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. **Note:** If `server_group_type` is set to `Fc`, you do not need to configure parameters, otherwise this attribute is required. If `server_group_type` is set to `Ip`, the value of this property is the same as the `server_id` value.
     * 
     */
    public Optional<Output<String>> serverIp() {
        return Optional.ofNullable(this.serverIp);
    }

    /**
     * The type of the server. The type of the server. Valid values:
     * - Ecs: an ECS instance.
     * - Eni: an ENI.
     * - Eci: an elastic container instance.
     * - Ip(Available in v1.194.0+): an IP address.
     * - fc(Available in v1.194.0+): a function.
     * 
     */
    @Import(name="serverType", required=true)
    private Output<String> serverType;

    /**
     * @return The type of the server. The type of the server. Valid values:
     * - Ecs: an ECS instance.
     * - Eni: an ENI.
     * - Eci: an elastic container instance.
     * - Ip(Available in v1.194.0+): an IP address.
     * - fc(Available in v1.194.0+): a function.
     * 
     */
    public Output<String> serverType() {
        return this.serverType;
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The weight of the server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no
     * requests are forwarded to the server. **Note:** You do not need to set this parameter if you set `server_type` to `Fc`.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return The weight of the server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no
     * requests are forwarded to the server. **Note:** You do not need to set this parameter if you set `server_type` to `Fc`.
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private ServerGroupServerArgs() {}

    private ServerGroupServerArgs(ServerGroupServerArgs $) {
        this.description = $.description;
        this.port = $.port;
        this.remoteIpEnabled = $.remoteIpEnabled;
        this.serverId = $.serverId;
        this.serverIp = $.serverIp;
        this.serverType = $.serverType;
        this.status = $.status;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerGroupServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerGroupServerArgs $;

        public Builder() {
            $ = new ServerGroupServerArgs();
        }

        public Builder(ServerGroupServerArgs defaults) {
            $ = new ServerGroupServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the server.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the server.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param port The port that is used by the server. Valid values: `1` to `65535`. **Note:** This parameter is required if the `server_type` parameter is set to `Ecs`, `Eni`, `Eci`, or `Ip`. You do not need to configure this parameter if you set `server_type` to `Fc`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port that is used by the server. Valid values: `1` to `65535`. **Note:** This parameter is required if the `server_type` parameter is set to `Ecs`, `Eni`, `Eci`, or `Ip`. You do not need to configure this parameter if you set `server_type` to `Fc`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param remoteIpEnabled Specifies whether to enable the remote IP address feature. You can specify up to 40 servers in each call. **Note:** If `server_type` is set to `Ip`, this parameter is available.
         * 
         * @return builder
         * 
         */
        public Builder remoteIpEnabled(@Nullable Output<Boolean> remoteIpEnabled) {
            $.remoteIpEnabled = remoteIpEnabled;
            return this;
        }

        /**
         * @param remoteIpEnabled Specifies whether to enable the remote IP address feature. You can specify up to 40 servers in each call. **Note:** If `server_type` is set to `Ip`, this parameter is available.
         * 
         * @return builder
         * 
         */
        public Builder remoteIpEnabled(Boolean remoteIpEnabled) {
            return remoteIpEnabled(Output.of(remoteIpEnabled));
        }

        /**
         * @param serverId The ID of the backend server.
         * - If `server_group_type` is set to `Instance`, set the parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by Ecs, Eni, or Eci.
         * - If `server_group_type` is set to `Ip`, set the parameter to an IP address specified in the server group.
         * - If `server_group_type` is set to `Fc`, set the parameter to the Alibaba Cloud Resource Name (ARN) of a function specified in the server group.
         * 
         * @return builder
         * 
         */
        public Builder serverId(Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        /**
         * @param serverId The ID of the backend server.
         * - If `server_group_type` is set to `Instance`, set the parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by Ecs, Eni, or Eci.
         * - If `server_group_type` is set to `Ip`, set the parameter to an IP address specified in the server group.
         * - If `server_group_type` is set to `Fc`, set the parameter to the Alibaba Cloud Resource Name (ARN) of a function specified in the server group.
         * 
         * @return builder
         * 
         */
        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        /**
         * @param serverIp The IP address of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. **Note:** If `server_group_type` is set to `Fc`, you do not need to configure parameters, otherwise this attribute is required. If `server_group_type` is set to `Ip`, the value of this property is the same as the `server_id` value.
         * 
         * @return builder
         * 
         */
        public Builder serverIp(@Nullable Output<String> serverIp) {
            $.serverIp = serverIp;
            return this;
        }

        /**
         * @param serverIp The IP address of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. **Note:** If `server_group_type` is set to `Fc`, you do not need to configure parameters, otherwise this attribute is required. If `server_group_type` is set to `Ip`, the value of this property is the same as the `server_id` value.
         * 
         * @return builder
         * 
         */
        public Builder serverIp(String serverIp) {
            return serverIp(Output.of(serverIp));
        }

        /**
         * @param serverType The type of the server. The type of the server. Valid values:
         * - Ecs: an ECS instance.
         * - Eni: an ENI.
         * - Eci: an elastic container instance.
         * - Ip(Available in v1.194.0+): an IP address.
         * - fc(Available in v1.194.0+): a function.
         * 
         * @return builder
         * 
         */
        public Builder serverType(Output<String> serverType) {
            $.serverType = serverType;
            return this;
        }

        /**
         * @param serverType The type of the server. The type of the server. Valid values:
         * - Ecs: an ECS instance.
         * - Eni: an ENI.
         * - Eci: an elastic container instance.
         * - Ip(Available in v1.194.0+): an IP address.
         * - fc(Available in v1.194.0+): a function.
         * 
         * @return builder
         * 
         */
        public Builder serverType(String serverType) {
            return serverType(Output.of(serverType));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param weight The weight of the server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no
         * requests are forwarded to the server. **Note:** You do not need to set this parameter if you set `server_type` to `Fc`.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight The weight of the server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no
         * requests are forwarded to the server. **Note:** You do not need to set this parameter if you set `server_type` to `Fc`.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public ServerGroupServerArgs build() {
            $.serverId = Objects.requireNonNull($.serverId, "expected parameter 'serverId' to be non-null");
            $.serverType = Objects.requireNonNull($.serverType, "expected parameter 'serverType' to be non-null");
            return $;
        }
    }

}