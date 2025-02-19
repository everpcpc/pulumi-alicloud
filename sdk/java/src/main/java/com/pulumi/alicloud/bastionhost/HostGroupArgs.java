// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HostGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final HostGroupArgs Empty = new HostGroupArgs();

    /**
     * Specify the New Host Group of Notes, Supports up to 500 Characters.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return Specify the New Host Group of Notes, Supports up to 500 Characters.
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * Specify the New Host Group Name, Supports up to 128 Characters.
     * 
     */
    @Import(name="hostGroupName", required=true)
    private Output<String> hostGroupName;

    /**
     * @return Specify the New Host Group Name, Supports up to 128 Characters.
     * 
     */
    public Output<String> hostGroupName() {
        return this.hostGroupName;
    }

    /**
     * Specify the New Host Group Where the Bastion Host ID of.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return Specify the New Host Group Where the Bastion Host ID of.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    private HostGroupArgs() {}

    private HostGroupArgs(HostGroupArgs $) {
        this.comment = $.comment;
        this.hostGroupName = $.hostGroupName;
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HostGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HostGroupArgs $;

        public Builder() {
            $ = new HostGroupArgs();
        }

        public Builder(HostGroupArgs defaults) {
            $ = new HostGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param comment Specify the New Host Group of Notes, Supports up to 500 Characters.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment Specify the New Host Group of Notes, Supports up to 500 Characters.
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param hostGroupName Specify the New Host Group Name, Supports up to 128 Characters.
         * 
         * @return builder
         * 
         */
        public Builder hostGroupName(Output<String> hostGroupName) {
            $.hostGroupName = hostGroupName;
            return this;
        }

        /**
         * @param hostGroupName Specify the New Host Group Name, Supports up to 128 Characters.
         * 
         * @return builder
         * 
         */
        public Builder hostGroupName(String hostGroupName) {
            return hostGroupName(Output.of(hostGroupName));
        }

        /**
         * @param instanceId Specify the New Host Group Where the Bastion Host ID of.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Specify the New Host Group Where the Bastion Host ID of.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        public HostGroupArgs build() {
            $.hostGroupName = Objects.requireNonNull($.hostGroupName, "expected parameter 'hostGroupName' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
