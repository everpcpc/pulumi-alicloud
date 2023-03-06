// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.databasefilesystem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceAttachmentState Empty = new InstanceAttachmentState();

    /**
     * The ID of the ECS instance.
     * 
     */
    @Import(name="ecsId")
    private @Nullable Output<String> ecsId;

    /**
     * @return The ID of the ECS instance.
     * 
     */
    public Optional<Output<String>> ecsId() {
        return Optional.ofNullable(this.ecsId);
    }

    /**
     * The ID of the database file system.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The ID of the database file system.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private InstanceAttachmentState() {}

    private InstanceAttachmentState(InstanceAttachmentState $) {
        this.ecsId = $.ecsId;
        this.instanceId = $.instanceId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceAttachmentState $;

        public Builder() {
            $ = new InstanceAttachmentState();
        }

        public Builder(InstanceAttachmentState defaults) {
            $ = new InstanceAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ecsId The ID of the ECS instance.
         * 
         * @return builder
         * 
         */
        public Builder ecsId(@Nullable Output<String> ecsId) {
            $.ecsId = ecsId;
            return this;
        }

        /**
         * @param ecsId The ID of the ECS instance.
         * 
         * @return builder
         * 
         */
        public Builder ecsId(String ecsId) {
            return ecsId(Output.of(ecsId));
        }

        /**
         * @param instanceId The ID of the database file system.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the database file system.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param status The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public InstanceAttachmentState build() {
            return $;
        }
    }

}