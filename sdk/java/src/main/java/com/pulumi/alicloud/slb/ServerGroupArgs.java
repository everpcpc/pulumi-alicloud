// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.slb.inputs.ServerGroupServerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerGroupArgs Empty = new ServerGroupArgs();

    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     * 
     */
    @Import(name="deleteProtectionValidation")
    private @Nullable Output<Boolean> deleteProtectionValidation;

    /**
     * @return Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     * 
     */
    public Optional<Output<Boolean>> deleteProtectionValidation() {
        return Optional.ofNullable(this.deleteProtectionValidation);
    }

    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     * 
     */
    @Import(name="loadBalancerId", required=true)
    private Output<String> loadBalancerId;

    /**
     * @return The Load Balancer ID which is used to launch a new virtual server group.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }

    /**
     * Name of the virtual server group. Our plugin provides a default name: &#34;tf-server-group&#34;.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the virtual server group. Our plugin provides a default name: &#34;tf-server-group&#34;.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of ECS instances to be added. **NOTE:** Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     * 
     * @deprecated
     * Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;.
     * 
     */
    @Deprecated /* Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. */
    @Import(name="servers")
    private @Nullable Output<List<ServerGroupServerArgs>> servers;

    /**
     * @return A list of ECS instances to be added. **NOTE:** Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     * 
     * @deprecated
     * Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;.
     * 
     */
    @Deprecated /* Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. */
    public Optional<Output<List<ServerGroupServerArgs>>> servers() {
        return Optional.ofNullable(this.servers);
    }

    private ServerGroupArgs() {}

    private ServerGroupArgs(ServerGroupArgs $) {
        this.deleteProtectionValidation = $.deleteProtectionValidation;
        this.loadBalancerId = $.loadBalancerId;
        this.name = $.name;
        this.servers = $.servers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerGroupArgs $;

        public Builder() {
            $ = new ServerGroupArgs();
        }

        public Builder(ServerGroupArgs defaults) {
            $ = new ServerGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deleteProtectionValidation Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
         * 
         * @return builder
         * 
         */
        public Builder deleteProtectionValidation(@Nullable Output<Boolean> deleteProtectionValidation) {
            $.deleteProtectionValidation = deleteProtectionValidation;
            return this;
        }

        /**
         * @param deleteProtectionValidation Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
         * 
         * @return builder
         * 
         */
        public Builder deleteProtectionValidation(Boolean deleteProtectionValidation) {
            return deleteProtectionValidation(Output.of(deleteProtectionValidation));
        }

        /**
         * @param loadBalancerId The Load Balancer ID which is used to launch a new virtual server group.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(Output<String> loadBalancerId) {
            $.loadBalancerId = loadBalancerId;
            return this;
        }

        /**
         * @param loadBalancerId The Load Balancer ID which is used to launch a new virtual server group.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(String loadBalancerId) {
            return loadBalancerId(Output.of(loadBalancerId));
        }

        /**
         * @param name Name of the virtual server group. Our plugin provides a default name: &#34;tf-server-group&#34;.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the virtual server group. Our plugin provides a default name: &#34;tf-server-group&#34;.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param servers A list of ECS instances to be added. **NOTE:** Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;.
         * 
         */
        @Deprecated /* Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. */
        public Builder servers(@Nullable Output<List<ServerGroupServerArgs>> servers) {
            $.servers = servers;
            return this;
        }

        /**
         * @param servers A list of ECS instances to be added. **NOTE:** Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;.
         * 
         */
        @Deprecated /* Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. */
        public Builder servers(List<ServerGroupServerArgs> servers) {
            return servers(Output.of(servers));
        }

        /**
         * @param servers A list of ECS instances to be added. **NOTE:** Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;servers&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_server_group_server_attachment&#39;.
         * 
         */
        @Deprecated /* Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. */
        public Builder servers(ServerGroupServerArgs... servers) {
            return servers(List.of(servers));
        }

        public ServerGroupArgs build() {
            $.loadBalancerId = Objects.requireNonNull($.loadBalancerId, "expected parameter 'loadBalancerId' to be non-null");
            return $;
        }
    }

}