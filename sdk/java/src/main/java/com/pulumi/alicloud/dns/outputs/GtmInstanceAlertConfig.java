// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GtmInstanceAlertConfig {
    /**
     * @return Whether to configure DingTalk notifications. Valid values: `true`, `false`.
     * 
     */
    private @Nullable Boolean dingtalkNotice;
    /**
     * @return Whether to configure mail notification. Valid values: `true`, `false`.
     * 
     */
    private @Nullable Boolean emailNotice;
    /**
     * @return The Alarm Event Type.
     * 
     */
    private @Nullable String noticeType;
    /**
     * @return Whether to configure SMS notification. Valid values: `true`, `false`.
     * 
     */
    private @Nullable Boolean smsNotice;

    private GtmInstanceAlertConfig() {}
    /**
     * @return Whether to configure DingTalk notifications. Valid values: `true`, `false`.
     * 
     */
    public Optional<Boolean> dingtalkNotice() {
        return Optional.ofNullable(this.dingtalkNotice);
    }
    /**
     * @return Whether to configure mail notification. Valid values: `true`, `false`.
     * 
     */
    public Optional<Boolean> emailNotice() {
        return Optional.ofNullable(this.emailNotice);
    }
    /**
     * @return The Alarm Event Type.
     * 
     */
    public Optional<String> noticeType() {
        return Optional.ofNullable(this.noticeType);
    }
    /**
     * @return Whether to configure SMS notification. Valid values: `true`, `false`.
     * 
     */
    public Optional<Boolean> smsNotice() {
        return Optional.ofNullable(this.smsNotice);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GtmInstanceAlertConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean dingtalkNotice;
        private @Nullable Boolean emailNotice;
        private @Nullable String noticeType;
        private @Nullable Boolean smsNotice;
        public Builder() {}
        public Builder(GtmInstanceAlertConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dingtalkNotice = defaults.dingtalkNotice;
    	      this.emailNotice = defaults.emailNotice;
    	      this.noticeType = defaults.noticeType;
    	      this.smsNotice = defaults.smsNotice;
        }

        @CustomType.Setter
        public Builder dingtalkNotice(@Nullable Boolean dingtalkNotice) {
            this.dingtalkNotice = dingtalkNotice;
            return this;
        }
        @CustomType.Setter
        public Builder emailNotice(@Nullable Boolean emailNotice) {
            this.emailNotice = emailNotice;
            return this;
        }
        @CustomType.Setter
        public Builder noticeType(@Nullable String noticeType) {
            this.noticeType = noticeType;
            return this;
        }
        @CustomType.Setter
        public Builder smsNotice(@Nullable Boolean smsNotice) {
            this.smsNotice = smsNotice;
            return this;
        }
        public GtmInstanceAlertConfig build() {
            final var o = new GtmInstanceAlertConfig();
            o.dingtalkNotice = dingtalkNotice;
            o.emailNotice = emailNotice;
            o.noticeType = noticeType;
            o.smsNotice = smsNotice;
            return o;
        }
    }
}
