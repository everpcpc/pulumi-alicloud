// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetChainsChainChainConfigNodeNodeConfigDenyPolicy {
    /**
     * @return The action of trigger blocking. Valid values: `BLOCK`, `BLOCK_RETAG`, `BLOCK_DELETE_TAG`. While `Block` means block the delivery chain from continuing to execute, `BLOCK_RETAG` means block overwriting push image tag, `BLOCK_DELETE_TAG` means block deletion of mirror tags.
     * 
     */
    private @Nullable String action;
    /**
     * @return The count of scanning vulnerabilities that triggers blocking.
     * 
     */
    private String issueCount;
    /**
     * @return The level of scanning vulnerability that triggers blocking. Valid values: `LOW`, `MEDIUM`, `HIGH`, `UNKNOWN`.
     * 
     */
    private String issueLevel;
    /**
     * @return The logic of trigger blocking. Valid values: `AND`, `OR`.
     * 
     */
    private String logic;

    private GetChainsChainChainConfigNodeNodeConfigDenyPolicy() {}
    /**
     * @return The action of trigger blocking. Valid values: `BLOCK`, `BLOCK_RETAG`, `BLOCK_DELETE_TAG`. While `Block` means block the delivery chain from continuing to execute, `BLOCK_RETAG` means block overwriting push image tag, `BLOCK_DELETE_TAG` means block deletion of mirror tags.
     * 
     */
    public Optional<String> action() {
        return Optional.ofNullable(this.action);
    }
    /**
     * @return The count of scanning vulnerabilities that triggers blocking.
     * 
     */
    public String issueCount() {
        return this.issueCount;
    }
    /**
     * @return The level of scanning vulnerability that triggers blocking. Valid values: `LOW`, `MEDIUM`, `HIGH`, `UNKNOWN`.
     * 
     */
    public String issueLevel() {
        return this.issueLevel;
    }
    /**
     * @return The logic of trigger blocking. Valid values: `AND`, `OR`.
     * 
     */
    public String logic() {
        return this.logic;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetChainsChainChainConfigNodeNodeConfigDenyPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String action;
        private String issueCount;
        private String issueLevel;
        private String logic;
        public Builder() {}
        public Builder(GetChainsChainChainConfigNodeNodeConfigDenyPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.action = defaults.action;
    	      this.issueCount = defaults.issueCount;
    	      this.issueLevel = defaults.issueLevel;
    	      this.logic = defaults.logic;
        }

        @CustomType.Setter
        public Builder action(@Nullable String action) {
            this.action = action;
            return this;
        }
        @CustomType.Setter
        public Builder issueCount(String issueCount) {
            this.issueCount = Objects.requireNonNull(issueCount);
            return this;
        }
        @CustomType.Setter
        public Builder issueLevel(String issueLevel) {
            this.issueLevel = Objects.requireNonNull(issueLevel);
            return this;
        }
        @CustomType.Setter
        public Builder logic(String logic) {
            this.logic = Objects.requireNonNull(logic);
            return this;
        }
        public GetChainsChainChainConfigNodeNodeConfigDenyPolicy build() {
            final var o = new GetChainsChainChainConfigNodeNodeConfigDenyPolicy();
            o.action = action;
            o.issueCount = issueCount;
            o.issueLevel = issueLevel;
            o.logic = logic;
            return o;
        }
    }
}
