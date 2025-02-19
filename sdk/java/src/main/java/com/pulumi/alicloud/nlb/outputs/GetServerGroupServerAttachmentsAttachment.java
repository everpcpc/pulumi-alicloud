// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServerGroupServerAttachmentsAttachment {
    /**
     * @return The description of the backend server.
     * 
     */
    private String description;
    /**
     * @return The ID of the server group.
     * 
     */
    private String id;
    /**
     * @return The port used by the backend server.
     * 
     */
    private Integer port;
    /**
     * @return The ID of the server group.
     * 
     */
    private String serverGroupId;
    /**
     * @return The ID of the server.
     * 
     */
    private String serverId;
    /**
     * @return The IP address of the backend server.
     * 
     */
    private String serverIp;
    /**
     * @return The type of the backend server.
     * 
     */
    private String serverType;
    /**
     * @return Indicates the status of the backend server.
     * 
     */
    private String status;
    /**
     * @return The weight of the backend server.
     * 
     */
    private Integer weight;
    /**
     * @return The zone ID of the server.
     * 
     */
    private String zoneId;

    private GetServerGroupServerAttachmentsAttachment() {}
    /**
     * @return The description of the backend server.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the server group.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The port used by the backend server.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return The ID of the server group.
     * 
     */
    public String serverGroupId() {
        return this.serverGroupId;
    }
    /**
     * @return The ID of the server.
     * 
     */
    public String serverId() {
        return this.serverId;
    }
    /**
     * @return The IP address of the backend server.
     * 
     */
    public String serverIp() {
        return this.serverIp;
    }
    /**
     * @return The type of the backend server.
     * 
     */
    public String serverType() {
        return this.serverType;
    }
    /**
     * @return Indicates the status of the backend server.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The weight of the backend server.
     * 
     */
    public Integer weight() {
        return this.weight;
    }
    /**
     * @return The zone ID of the server.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerGroupServerAttachmentsAttachment defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private Integer port;
        private String serverGroupId;
        private String serverId;
        private String serverIp;
        private String serverType;
        private String status;
        private Integer weight;
        private String zoneId;
        public Builder() {}
        public Builder(GetServerGroupServerAttachmentsAttachment defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.port = defaults.port;
    	      this.serverGroupId = defaults.serverGroupId;
    	      this.serverId = defaults.serverId;
    	      this.serverIp = defaults.serverIp;
    	      this.serverType = defaults.serverType;
    	      this.status = defaults.status;
    	      this.weight = defaults.weight;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder serverGroupId(String serverGroupId) {
            this.serverGroupId = Objects.requireNonNull(serverGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder serverId(String serverId) {
            this.serverId = Objects.requireNonNull(serverId);
            return this;
        }
        @CustomType.Setter
        public Builder serverIp(String serverIp) {
            this.serverIp = Objects.requireNonNull(serverIp);
            return this;
        }
        @CustomType.Setter
        public Builder serverType(String serverType) {
            this.serverType = Objects.requireNonNull(serverType);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetServerGroupServerAttachmentsAttachment build() {
            final var o = new GetServerGroupServerAttachmentsAttachment();
            o.description = description;
            o.id = id;
            o.port = port;
            o.serverGroupId = serverGroupId;
            o.serverId = serverId;
            o.serverIp = serverIp;
            o.serverType = serverType;
            o.status = status;
            o.weight = weight;
            o.zoneId = zoneId;
            return o;
        }
    }
}
