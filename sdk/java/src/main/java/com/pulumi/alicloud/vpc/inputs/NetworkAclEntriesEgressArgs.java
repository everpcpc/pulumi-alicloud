// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkAclEntriesEgressArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkAclEntriesEgressArgs Empty = new NetworkAclEntriesEgressArgs();

    /**
     * The description of the ingress entry.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the ingress entry.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The destination ip of the egress entry.
     * 
     */
    @Import(name="destinationCidrIp")
    private @Nullable Output<String> destinationCidrIp;

    /**
     * @return The destination ip of the egress entry.
     * 
     */
    public Optional<Output<String>> destinationCidrIp() {
        return Optional.ofNullable(this.destinationCidrIp);
    }

    /**
     * The entry type of the ingress entry. It must be `custom` or `system`. Default value is `custom`.
     * 
     */
    @Import(name="entryType")
    private @Nullable Output<String> entryType;

    /**
     * @return The entry type of the ingress entry. It must be `custom` or `system`. Default value is `custom`.
     * 
     */
    public Optional<Output<String>> entryType() {
        return Optional.ofNullable(this.entryType);
    }

    /**
     * The name of the ingress entry.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the ingress entry.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The policy of the ingress entry. It must be `accept` or `drop`.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return The policy of the ingress entry. It must be `accept` or `drop`.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    /**
     * The port of the ingress entry.
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return The port of the ingress entry.
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The protocol of the ingress entry.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The protocol of the ingress entry.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    private NetworkAclEntriesEgressArgs() {}

    private NetworkAclEntriesEgressArgs(NetworkAclEntriesEgressArgs $) {
        this.description = $.description;
        this.destinationCidrIp = $.destinationCidrIp;
        this.entryType = $.entryType;
        this.name = $.name;
        this.policy = $.policy;
        this.port = $.port;
        this.protocol = $.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkAclEntriesEgressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkAclEntriesEgressArgs $;

        public Builder() {
            $ = new NetworkAclEntriesEgressArgs();
        }

        public Builder(NetworkAclEntriesEgressArgs defaults) {
            $ = new NetworkAclEntriesEgressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param destinationCidrIp The destination ip of the egress entry.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidrIp(@Nullable Output<String> destinationCidrIp) {
            $.destinationCidrIp = destinationCidrIp;
            return this;
        }

        /**
         * @param destinationCidrIp The destination ip of the egress entry.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidrIp(String destinationCidrIp) {
            return destinationCidrIp(Output.of(destinationCidrIp));
        }

        /**
         * @param entryType The entry type of the ingress entry. It must be `custom` or `system`. Default value is `custom`.
         * 
         * @return builder
         * 
         */
        public Builder entryType(@Nullable Output<String> entryType) {
            $.entryType = entryType;
            return this;
        }

        /**
         * @param entryType The entry type of the ingress entry. It must be `custom` or `system`. Default value is `custom`.
         * 
         * @return builder
         * 
         */
        public Builder entryType(String entryType) {
            return entryType(Output.of(entryType));
        }

        /**
         * @param name The name of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policy The policy of the ingress entry. It must be `accept` or `drop`.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The policy of the ingress entry. It must be `accept` or `drop`.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param port The port of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The protocol of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol of the ingress entry.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        public NetworkAclEntriesEgressArgs build() {
            return $;
        }
    }

}
