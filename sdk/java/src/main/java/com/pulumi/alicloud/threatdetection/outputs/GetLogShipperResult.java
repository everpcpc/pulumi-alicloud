// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetLogShipperResult {
    /**
     * @return Log Analysis Service authorization status.
     * 
     */
    private String authStatus;
    /**
     * @return Cloud Security Center purchase status.
     * 
     */
    private String buyStatus;
    private @Nullable String enable;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Log analysis shipping activation status.
     * 
     */
    private String openStatus;
    /**
     * @return Log analysis project status.
     * 
     */
    private String slsProjectStatus;
    /**
     * @return Log Analysis Service is activated.
     * 
     */
    private String slsServiceStatus;
    /**
     * @return The current service enable status.
     * 
     */
    private String status;

    private GetLogShipperResult() {}
    /**
     * @return Log Analysis Service authorization status.
     * 
     */
    public String authStatus() {
        return this.authStatus;
    }
    /**
     * @return Cloud Security Center purchase status.
     * 
     */
    public String buyStatus() {
        return this.buyStatus;
    }
    public Optional<String> enable() {
        return Optional.ofNullable(this.enable);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Log analysis shipping activation status.
     * 
     */
    public String openStatus() {
        return this.openStatus;
    }
    /**
     * @return Log analysis project status.
     * 
     */
    public String slsProjectStatus() {
        return this.slsProjectStatus;
    }
    /**
     * @return Log Analysis Service is activated.
     * 
     */
    public String slsServiceStatus() {
        return this.slsServiceStatus;
    }
    /**
     * @return The current service enable status.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogShipperResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String authStatus;
        private String buyStatus;
        private @Nullable String enable;
        private String id;
        private String openStatus;
        private String slsProjectStatus;
        private String slsServiceStatus;
        private String status;
        public Builder() {}
        public Builder(GetLogShipperResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authStatus = defaults.authStatus;
    	      this.buyStatus = defaults.buyStatus;
    	      this.enable = defaults.enable;
    	      this.id = defaults.id;
    	      this.openStatus = defaults.openStatus;
    	      this.slsProjectStatus = defaults.slsProjectStatus;
    	      this.slsServiceStatus = defaults.slsServiceStatus;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder authStatus(String authStatus) {
            this.authStatus = Objects.requireNonNull(authStatus);
            return this;
        }
        @CustomType.Setter
        public Builder buyStatus(String buyStatus) {
            this.buyStatus = Objects.requireNonNull(buyStatus);
            return this;
        }
        @CustomType.Setter
        public Builder enable(@Nullable String enable) {
            this.enable = enable;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder openStatus(String openStatus) {
            this.openStatus = Objects.requireNonNull(openStatus);
            return this;
        }
        @CustomType.Setter
        public Builder slsProjectStatus(String slsProjectStatus) {
            this.slsProjectStatus = Objects.requireNonNull(slsProjectStatus);
            return this;
        }
        @CustomType.Setter
        public Builder slsServiceStatus(String slsServiceStatus) {
            this.slsServiceStatus = Objects.requireNonNull(slsServiceStatus);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetLogShipperResult build() {
            final var o = new GetLogShipperResult();
            o.authStatus = authStatus;
            o.buyStatus = buyStatus;
            o.enable = enable;
            o.id = id;
            o.openStatus = openStatus;
            o.slsProjectStatus = slsProjectStatus;
            o.slsServiceStatus = slsServiceStatus;
            o.status = status;
            return o;
        }
    }
}