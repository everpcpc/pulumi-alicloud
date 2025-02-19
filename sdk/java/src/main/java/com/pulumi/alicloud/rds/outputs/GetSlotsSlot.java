// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSlotsSlot {
    /**
     * @return The name of the database where Replication Slot is located.
     * 
     */
    private String database;
    /**
     * @return The plugin used by Replication Slot.
     * 
     */
    private String plugin;
    /**
     * @return The Replication Slot name.
     * 
     */
    private String slotName;
    /**
     * @return The Replication Slot status.
     * 
     */
    private String slotStatus;
    /**
     * @return The Replication Slot type.
     * 
     */
    private String slotType;
    /**
     * @return Is the Replication Slot temporary.
     * 
     */
    private String temporary;
    /**
     * @return The amount of logs accumulated by Replication Slot.
     * 
     */
    private String walDelay;

    private GetSlotsSlot() {}
    /**
     * @return The name of the database where Replication Slot is located.
     * 
     */
    public String database() {
        return this.database;
    }
    /**
     * @return The plugin used by Replication Slot.
     * 
     */
    public String plugin() {
        return this.plugin;
    }
    /**
     * @return The Replication Slot name.
     * 
     */
    public String slotName() {
        return this.slotName;
    }
    /**
     * @return The Replication Slot status.
     * 
     */
    public String slotStatus() {
        return this.slotStatus;
    }
    /**
     * @return The Replication Slot type.
     * 
     */
    public String slotType() {
        return this.slotType;
    }
    /**
     * @return Is the Replication Slot temporary.
     * 
     */
    public String temporary() {
        return this.temporary;
    }
    /**
     * @return The amount of logs accumulated by Replication Slot.
     * 
     */
    public String walDelay() {
        return this.walDelay;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSlotsSlot defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String database;
        private String plugin;
        private String slotName;
        private String slotStatus;
        private String slotType;
        private String temporary;
        private String walDelay;
        public Builder() {}
        public Builder(GetSlotsSlot defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.database = defaults.database;
    	      this.plugin = defaults.plugin;
    	      this.slotName = defaults.slotName;
    	      this.slotStatus = defaults.slotStatus;
    	      this.slotType = defaults.slotType;
    	      this.temporary = defaults.temporary;
    	      this.walDelay = defaults.walDelay;
        }

        @CustomType.Setter
        public Builder database(String database) {
            this.database = Objects.requireNonNull(database);
            return this;
        }
        @CustomType.Setter
        public Builder plugin(String plugin) {
            this.plugin = Objects.requireNonNull(plugin);
            return this;
        }
        @CustomType.Setter
        public Builder slotName(String slotName) {
            this.slotName = Objects.requireNonNull(slotName);
            return this;
        }
        @CustomType.Setter
        public Builder slotStatus(String slotStatus) {
            this.slotStatus = Objects.requireNonNull(slotStatus);
            return this;
        }
        @CustomType.Setter
        public Builder slotType(String slotType) {
            this.slotType = Objects.requireNonNull(slotType);
            return this;
        }
        @CustomType.Setter
        public Builder temporary(String temporary) {
            this.temporary = Objects.requireNonNull(temporary);
            return this;
        }
        @CustomType.Setter
        public Builder walDelay(String walDelay) {
            this.walDelay = Objects.requireNonNull(walDelay);
            return this;
        }
        public GetSlotsSlot build() {
            final var o = new GetSlotsSlot();
            o.database = database;
            o.plugin = plugin;
            o.slotName = slotName;
            o.slotStatus = slotStatus;
            o.slotType = slotType;
            o.temporary = temporary;
            o.walDelay = walDelay;
            return o;
        }
    }
}
