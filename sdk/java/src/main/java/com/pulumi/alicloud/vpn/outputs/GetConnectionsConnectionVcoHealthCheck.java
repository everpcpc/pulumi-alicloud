// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetConnectionsConnectionVcoHealthCheck {
    /**
     * @return The destination ip address.
     * 
     */
    private @Nullable String dip;
    /**
     * @return The health check on status. Valid values: `true`, `false`.
     * 
     */
    private @Nullable String enable;
    /**
     * @return The time interval between health checks.
     * 
     */
    private @Nullable Integer interval;
    /**
     * @return The number of retries for health checks issued.
     * 
     */
    private @Nullable Integer retry;
    /**
     * @return The source ip address.
     * 
     */
    private @Nullable String sip;
    /**
     * @return The negotiation status of the BGP routing protocol. Valid values: `success`, `false`.
     * 
     */
    private @Nullable String status;

    private GetConnectionsConnectionVcoHealthCheck() {}
    /**
     * @return The destination ip address.
     * 
     */
    public Optional<String> dip() {
        return Optional.ofNullable(this.dip);
    }
    /**
     * @return The health check on status. Valid values: `true`, `false`.
     * 
     */
    public Optional<String> enable() {
        return Optional.ofNullable(this.enable);
    }
    /**
     * @return The time interval between health checks.
     * 
     */
    public Optional<Integer> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return The number of retries for health checks issued.
     * 
     */
    public Optional<Integer> retry() {
        return Optional.ofNullable(this.retry);
    }
    /**
     * @return The source ip address.
     * 
     */
    public Optional<String> sip() {
        return Optional.ofNullable(this.sip);
    }
    /**
     * @return The negotiation status of the BGP routing protocol. Valid values: `success`, `false`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionsConnectionVcoHealthCheck defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String dip;
        private @Nullable String enable;
        private @Nullable Integer interval;
        private @Nullable Integer retry;
        private @Nullable String sip;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetConnectionsConnectionVcoHealthCheck defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dip = defaults.dip;
    	      this.enable = defaults.enable;
    	      this.interval = defaults.interval;
    	      this.retry = defaults.retry;
    	      this.sip = defaults.sip;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder dip(@Nullable String dip) {
            this.dip = dip;
            return this;
        }
        @CustomType.Setter
        public Builder enable(@Nullable String enable) {
            this.enable = enable;
            return this;
        }
        @CustomType.Setter
        public Builder interval(@Nullable Integer interval) {
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder retry(@Nullable Integer retry) {
            this.retry = retry;
            return this;
        }
        @CustomType.Setter
        public Builder sip(@Nullable String sip) {
            this.sip = sip;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetConnectionsConnectionVcoHealthCheck build() {
            final var o = new GetConnectionsConnectionVcoHealthCheck();
            o.dip = dip;
            o.enable = enable;
            o.interval = interval;
            o.retry = retry;
            o.sip = sip;
            o.status = status;
            return o;
        }
    }
}
