// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VbrHaState extends com.pulumi.resources.ResourceArgs {

    public static final VbrHaState Empty = new VbrHaState();

    /**
     * The description of the VBR switching group. It must be `2` to `256` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the VBR switching group. It must be `2` to `256` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The ID of the other VBR in the VBR failover group.
     * 
     */
    @Import(name="peerVbrId")
    private @Nullable Output<String> peerVbrId;

    /**
     * @return The ID of the other VBR in the VBR failover group.
     * 
     */
    public Optional<Output<String>> peerVbrId() {
        return Optional.ofNullable(this.peerVbrId);
    }

    /**
     * The state of the VBR failover group.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The state of the VBR failover group.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The name of the VBR failover group.
     * 
     */
    @Import(name="vbrHaName")
    private @Nullable Output<String> vbrHaName;

    /**
     * @return The name of the VBR failover group.
     * 
     */
    public Optional<Output<String>> vbrHaName() {
        return Optional.ofNullable(this.vbrHaName);
    }

    /**
     * The ID of the VBR instance.
     * 
     */
    @Import(name="vbrId")
    private @Nullable Output<String> vbrId;

    /**
     * @return The ID of the VBR instance.
     * 
     */
    public Optional<Output<String>> vbrId() {
        return Optional.ofNullable(this.vbrId);
    }

    private VbrHaState() {}

    private VbrHaState(VbrHaState $) {
        this.description = $.description;
        this.dryRun = $.dryRun;
        this.peerVbrId = $.peerVbrId;
        this.status = $.status;
        this.vbrHaName = $.vbrHaName;
        this.vbrId = $.vbrId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VbrHaState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VbrHaState $;

        public Builder() {
            $ = new VbrHaState();
        }

        public Builder(VbrHaState defaults) {
            $ = new VbrHaState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the VBR switching group. It must be `2` to `256` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the VBR switching group. It must be `2` to `256` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param peerVbrId The ID of the other VBR in the VBR failover group.
         * 
         * @return builder
         * 
         */
        public Builder peerVbrId(@Nullable Output<String> peerVbrId) {
            $.peerVbrId = peerVbrId;
            return this;
        }

        /**
         * @param peerVbrId The ID of the other VBR in the VBR failover group.
         * 
         * @return builder
         * 
         */
        public Builder peerVbrId(String peerVbrId) {
            return peerVbrId(Output.of(peerVbrId));
        }

        /**
         * @param status The state of the VBR failover group.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The state of the VBR failover group.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param vbrHaName The name of the VBR failover group.
         * 
         * @return builder
         * 
         */
        public Builder vbrHaName(@Nullable Output<String> vbrHaName) {
            $.vbrHaName = vbrHaName;
            return this;
        }

        /**
         * @param vbrHaName The name of the VBR failover group.
         * 
         * @return builder
         * 
         */
        public Builder vbrHaName(String vbrHaName) {
            return vbrHaName(Output.of(vbrHaName));
        }

        /**
         * @param vbrId The ID of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder vbrId(@Nullable Output<String> vbrId) {
            $.vbrId = vbrId;
            return this;
        }

        /**
         * @param vbrId The ID of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder vbrId(String vbrId) {
            return vbrId(Output.of(vbrId));
        }

        public VbrHaState build() {
            return $;
        }
    }

}
