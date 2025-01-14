// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServerGroupHealthCheck {
    /**
     * @return The backend port that is used for health checks. Valid values: 0 to 65535. Default value: 0. If you set the value to 0, the port of a backend server is used for health checks.
     * 
     */
    private @Nullable Integer healthCheckConnectPort;
    /**
     * @return The maximum timeout period of a health check response. Unit: seconds. Valid values: 1 to 300. Default value: 5.
     * 
     */
    private @Nullable Integer healthCheckConnectTimeout;
    /**
     * @return The domain name that is used for health checks. Valid values:
     * - `$SERVER_IP`: the private IP address of a backend server.
     * 
     */
    private @Nullable String healthCheckDomain;
    /**
     * @return Specifies whether to enable health checks.
     * 
     */
    private @Nullable Boolean healthCheckEnabled;
    /**
     * @return The HTTP status codes to return to health checks. Separate multiple HTTP status codes with commas (,). Valid values: http_2xx (default), http_3xx, http_4xx, and http_5xx. **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
     * 
     */
    private @Nullable List<String> healthCheckHttpCodes;
    /**
     * @return The interval between two consecutive health checks. Unit: seconds. Valid values: 5 to 5000. Default value: 10.
     * 
     */
    private @Nullable Integer healthCheckInterval;
    /**
     * @return The protocol that is used for health checks. Valid values: `TCP` (default) and `HTTP`.
     * 
     */
    private @Nullable String healthCheckType;
    /**
     * @return The path to which health check requests are sent. The path must be 1 to 80 characters in length, and can contain only letters, digits, and the following special characters: `- / . % ? # &amp; =`. It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : &#39; , +`. The path must start with a forward slash (/). **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
     * 
     */
    private @Nullable String healthCheckUrl;
    /**
     * @return The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from fail to success. Valid values: 2 to 10. Default value: 2.
     * 
     */
    private @Nullable Integer healthyThreshold;
    /**
     * @return The HTTP method that is used for health checks. Valid values: `GET` and `HEAD`. **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
     * 
     */
    private @Nullable String httpCheckMethod;
    /**
     * @return The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from success to fail. Valid values: 2 to 10. Default value: 2.
     * 
     */
    private @Nullable Integer unhealthyThreshold;

    private ServerGroupHealthCheck() {}
    /**
     * @return The backend port that is used for health checks. Valid values: 0 to 65535. Default value: 0. If you set the value to 0, the port of a backend server is used for health checks.
     * 
     */
    public Optional<Integer> healthCheckConnectPort() {
        return Optional.ofNullable(this.healthCheckConnectPort);
    }
    /**
     * @return The maximum timeout period of a health check response. Unit: seconds. Valid values: 1 to 300. Default value: 5.
     * 
     */
    public Optional<Integer> healthCheckConnectTimeout() {
        return Optional.ofNullable(this.healthCheckConnectTimeout);
    }
    /**
     * @return The domain name that is used for health checks. Valid values:
     * - `$SERVER_IP`: the private IP address of a backend server.
     * 
     */
    public Optional<String> healthCheckDomain() {
        return Optional.ofNullable(this.healthCheckDomain);
    }
    /**
     * @return Specifies whether to enable health checks.
     * 
     */
    public Optional<Boolean> healthCheckEnabled() {
        return Optional.ofNullable(this.healthCheckEnabled);
    }
    /**
     * @return The HTTP status codes to return to health checks. Separate multiple HTTP status codes with commas (,). Valid values: http_2xx (default), http_3xx, http_4xx, and http_5xx. **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
     * 
     */
    public List<String> healthCheckHttpCodes() {
        return this.healthCheckHttpCodes == null ? List.of() : this.healthCheckHttpCodes;
    }
    /**
     * @return The interval between two consecutive health checks. Unit: seconds. Valid values: 5 to 5000. Default value: 10.
     * 
     */
    public Optional<Integer> healthCheckInterval() {
        return Optional.ofNullable(this.healthCheckInterval);
    }
    /**
     * @return The protocol that is used for health checks. Valid values: `TCP` (default) and `HTTP`.
     * 
     */
    public Optional<String> healthCheckType() {
        return Optional.ofNullable(this.healthCheckType);
    }
    /**
     * @return The path to which health check requests are sent. The path must be 1 to 80 characters in length, and can contain only letters, digits, and the following special characters: `- / . % ? # &amp; =`. It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : &#39; , +`. The path must start with a forward slash (/). **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
     * 
     */
    public Optional<String> healthCheckUrl() {
        return Optional.ofNullable(this.healthCheckUrl);
    }
    /**
     * @return The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from fail to success. Valid values: 2 to 10. Default value: 2.
     * 
     */
    public Optional<Integer> healthyThreshold() {
        return Optional.ofNullable(this.healthyThreshold);
    }
    /**
     * @return The HTTP method that is used for health checks. Valid values: `GET` and `HEAD`. **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
     * 
     */
    public Optional<String> httpCheckMethod() {
        return Optional.ofNullable(this.httpCheckMethod);
    }
    /**
     * @return The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from success to fail. Valid values: 2 to 10. Default value: 2.
     * 
     */
    public Optional<Integer> unhealthyThreshold() {
        return Optional.ofNullable(this.unhealthyThreshold);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerGroupHealthCheck defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer healthCheckConnectPort;
        private @Nullable Integer healthCheckConnectTimeout;
        private @Nullable String healthCheckDomain;
        private @Nullable Boolean healthCheckEnabled;
        private @Nullable List<String> healthCheckHttpCodes;
        private @Nullable Integer healthCheckInterval;
        private @Nullable String healthCheckType;
        private @Nullable String healthCheckUrl;
        private @Nullable Integer healthyThreshold;
        private @Nullable String httpCheckMethod;
        private @Nullable Integer unhealthyThreshold;
        public Builder() {}
        public Builder(ServerGroupHealthCheck defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.healthCheckConnectPort = defaults.healthCheckConnectPort;
    	      this.healthCheckConnectTimeout = defaults.healthCheckConnectTimeout;
    	      this.healthCheckDomain = defaults.healthCheckDomain;
    	      this.healthCheckEnabled = defaults.healthCheckEnabled;
    	      this.healthCheckHttpCodes = defaults.healthCheckHttpCodes;
    	      this.healthCheckInterval = defaults.healthCheckInterval;
    	      this.healthCheckType = defaults.healthCheckType;
    	      this.healthCheckUrl = defaults.healthCheckUrl;
    	      this.healthyThreshold = defaults.healthyThreshold;
    	      this.httpCheckMethod = defaults.httpCheckMethod;
    	      this.unhealthyThreshold = defaults.unhealthyThreshold;
        }

        @CustomType.Setter
        public Builder healthCheckConnectPort(@Nullable Integer healthCheckConnectPort) {
            this.healthCheckConnectPort = healthCheckConnectPort;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckConnectTimeout(@Nullable Integer healthCheckConnectTimeout) {
            this.healthCheckConnectTimeout = healthCheckConnectTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckDomain(@Nullable String healthCheckDomain) {
            this.healthCheckDomain = healthCheckDomain;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckEnabled(@Nullable Boolean healthCheckEnabled) {
            this.healthCheckEnabled = healthCheckEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckHttpCodes(@Nullable List<String> healthCheckHttpCodes) {
            this.healthCheckHttpCodes = healthCheckHttpCodes;
            return this;
        }
        public Builder healthCheckHttpCodes(String... healthCheckHttpCodes) {
            return healthCheckHttpCodes(List.of(healthCheckHttpCodes));
        }
        @CustomType.Setter
        public Builder healthCheckInterval(@Nullable Integer healthCheckInterval) {
            this.healthCheckInterval = healthCheckInterval;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckType(@Nullable String healthCheckType) {
            this.healthCheckType = healthCheckType;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckUrl(@Nullable String healthCheckUrl) {
            this.healthCheckUrl = healthCheckUrl;
            return this;
        }
        @CustomType.Setter
        public Builder healthyThreshold(@Nullable Integer healthyThreshold) {
            this.healthyThreshold = healthyThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder httpCheckMethod(@Nullable String httpCheckMethod) {
            this.httpCheckMethod = httpCheckMethod;
            return this;
        }
        @CustomType.Setter
        public Builder unhealthyThreshold(@Nullable Integer unhealthyThreshold) {
            this.unhealthyThreshold = unhealthyThreshold;
            return this;
        }
        public ServerGroupHealthCheck build() {
            final var o = new ServerGroupHealthCheck();
            o.healthCheckConnectPort = healthCheckConnectPort;
            o.healthCheckConnectTimeout = healthCheckConnectTimeout;
            o.healthCheckDomain = healthCheckDomain;
            o.healthCheckEnabled = healthCheckEnabled;
            o.healthCheckHttpCodes = healthCheckHttpCodes;
            o.healthCheckInterval = healthCheckInterval;
            o.healthCheckType = healthCheckType;
            o.healthCheckUrl = healthCheckUrl;
            o.healthyThreshold = healthyThreshold;
            o.httpCheckMethod = httpCheckMethod;
            o.unhealthyThreshold = unhealthyThreshold;
            return o;
        }
    }
}
