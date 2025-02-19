// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.databasegateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGatewaysGatewayGatewayInstance {
    /**
     * @return The connection type of Gateway instance.
     * 
     */
    private String connectEndpointType;
    /**
     * @return The process of version number of Gateway instance.
     * 
     */
    private String currentDaemonVersion;
    /**
     * @return The version of Gateway instance.
     * 
     */
    private String currentVersion;
    /**
     * @return The endpoint address of Gateway instance.
     * 
     */
    private String endPoint;
    /**
     * @return The id of Gateway instance.
     * 
     */
    private String gatewayInstanceId;
    /**
     * @return The status of Gateway instance. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
     * 
     */
    private String gatewayInstanceStatus;
    /**
     * @return The last Updated time stamp of Gateway instance.
     * 
     */
    private String lastUpdateTime;
    /**
     * @return The Local IP ADDRESS of Gateway instance.
     * 
     */
    private String localIp;
    /**
     * @return The prompt information of Gateway instance.
     * 
     */
    private String message;
    /**
     * @return The host of Gateway instance.
     * 
     */
    private String outputIp;

    private GetGatewaysGatewayGatewayInstance() {}
    /**
     * @return The connection type of Gateway instance.
     * 
     */
    public String connectEndpointType() {
        return this.connectEndpointType;
    }
    /**
     * @return The process of version number of Gateway instance.
     * 
     */
    public String currentDaemonVersion() {
        return this.currentDaemonVersion;
    }
    /**
     * @return The version of Gateway instance.
     * 
     */
    public String currentVersion() {
        return this.currentVersion;
    }
    /**
     * @return The endpoint address of Gateway instance.
     * 
     */
    public String endPoint() {
        return this.endPoint;
    }
    /**
     * @return The id of Gateway instance.
     * 
     */
    public String gatewayInstanceId() {
        return this.gatewayInstanceId;
    }
    /**
     * @return The status of Gateway instance. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
     * 
     */
    public String gatewayInstanceStatus() {
        return this.gatewayInstanceStatus;
    }
    /**
     * @return The last Updated time stamp of Gateway instance.
     * 
     */
    public String lastUpdateTime() {
        return this.lastUpdateTime;
    }
    /**
     * @return The Local IP ADDRESS of Gateway instance.
     * 
     */
    public String localIp() {
        return this.localIp;
    }
    /**
     * @return The prompt information of Gateway instance.
     * 
     */
    public String message() {
        return this.message;
    }
    /**
     * @return The host of Gateway instance.
     * 
     */
    public String outputIp() {
        return this.outputIp;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGatewaysGatewayGatewayInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String connectEndpointType;
        private String currentDaemonVersion;
        private String currentVersion;
        private String endPoint;
        private String gatewayInstanceId;
        private String gatewayInstanceStatus;
        private String lastUpdateTime;
        private String localIp;
        private String message;
        private String outputIp;
        public Builder() {}
        public Builder(GetGatewaysGatewayGatewayInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectEndpointType = defaults.connectEndpointType;
    	      this.currentDaemonVersion = defaults.currentDaemonVersion;
    	      this.currentVersion = defaults.currentVersion;
    	      this.endPoint = defaults.endPoint;
    	      this.gatewayInstanceId = defaults.gatewayInstanceId;
    	      this.gatewayInstanceStatus = defaults.gatewayInstanceStatus;
    	      this.lastUpdateTime = defaults.lastUpdateTime;
    	      this.localIp = defaults.localIp;
    	      this.message = defaults.message;
    	      this.outputIp = defaults.outputIp;
        }

        @CustomType.Setter
        public Builder connectEndpointType(String connectEndpointType) {
            this.connectEndpointType = Objects.requireNonNull(connectEndpointType);
            return this;
        }
        @CustomType.Setter
        public Builder currentDaemonVersion(String currentDaemonVersion) {
            this.currentDaemonVersion = Objects.requireNonNull(currentDaemonVersion);
            return this;
        }
        @CustomType.Setter
        public Builder currentVersion(String currentVersion) {
            this.currentVersion = Objects.requireNonNull(currentVersion);
            return this;
        }
        @CustomType.Setter
        public Builder endPoint(String endPoint) {
            this.endPoint = Objects.requireNonNull(endPoint);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayInstanceId(String gatewayInstanceId) {
            this.gatewayInstanceId = Objects.requireNonNull(gatewayInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayInstanceStatus(String gatewayInstanceStatus) {
            this.gatewayInstanceStatus = Objects.requireNonNull(gatewayInstanceStatus);
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdateTime(String lastUpdateTime) {
            this.lastUpdateTime = Objects.requireNonNull(lastUpdateTime);
            return this;
        }
        @CustomType.Setter
        public Builder localIp(String localIp) {
            this.localIp = Objects.requireNonNull(localIp);
            return this;
        }
        @CustomType.Setter
        public Builder message(String message) {
            this.message = Objects.requireNonNull(message);
            return this;
        }
        @CustomType.Setter
        public Builder outputIp(String outputIp) {
            this.outputIp = Objects.requireNonNull(outputIp);
            return this;
        }
        public GetGatewaysGatewayGatewayInstance build() {
            final var o = new GetGatewaysGatewayGatewayInstance();
            o.connectEndpointType = connectEndpointType;
            o.currentDaemonVersion = currentDaemonVersion;
            o.currentVersion = currentVersion;
            o.endPoint = endPoint;
            o.gatewayInstanceId = gatewayInstanceId;
            o.gatewayInstanceStatus = gatewayInstanceStatus;
            o.lastUpdateTime = lastUpdateTime;
            o.localIp = localIp;
            o.message = message;
            o.outputIp = outputIp;
            return o;
        }
    }
}
