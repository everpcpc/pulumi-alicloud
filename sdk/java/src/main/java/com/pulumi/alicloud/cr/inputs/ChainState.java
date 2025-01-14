// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr.inputs;

import com.pulumi.alicloud.cr.inputs.ChainChainConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ChainState extends com.pulumi.resources.ResourceArgs {

    public static final ChainState Empty = new ChainState();

    /**
     * The configuration of delivery chain. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    @Import(name="chainConfigs")
    private @Nullable Output<List<ChainChainConfigArgs>> chainConfigs;

    /**
     * @return The configuration of delivery chain. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    public Optional<Output<List<ChainChainConfigArgs>>> chainConfigs() {
        return Optional.ofNullable(this.chainConfigs);
    }

    /**
     * Delivery chain ID.
     * 
     */
    @Import(name="chainId")
    private @Nullable Output<String> chainId;

    /**
     * @return Delivery chain ID.
     * 
     */
    public Optional<Output<String>> chainId() {
        return Optional.ofNullable(this.chainId);
    }

    /**
     * The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators &#34;_&#34;, &#34;-&#34;, &#34;.&#34; can be used, noted that the separator cannot be at the first or last position.
     * 
     */
    @Import(name="chainName")
    private @Nullable Output<String> chainName;

    /**
     * @return The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators &#34;_&#34;, &#34;-&#34;, &#34;.&#34; can be used, noted that the separator cannot be at the first or last position.
     * 
     */
    public Optional<Output<String>> chainName() {
        return Optional.ofNullable(this.chainName);
    }

    /**
     * The description delivery chain.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description delivery chain.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of CR Enterprise Edition instance.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The ID of CR Enterprise Edition instance.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
     * 
     */
    @Import(name="repoName")
    private @Nullable Output<String> repoName;

    /**
     * @return The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
     * 
     */
    public Optional<Output<String>> repoName() {
        return Optional.ofNullable(this.repoName);
    }

    /**
     * The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    @Import(name="repoNamespaceName")
    private @Nullable Output<String> repoNamespaceName;

    /**
     * @return The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    public Optional<Output<String>> repoNamespaceName() {
        return Optional.ofNullable(this.repoNamespaceName);
    }

    private ChainState() {}

    private ChainState(ChainState $) {
        this.chainConfigs = $.chainConfigs;
        this.chainId = $.chainId;
        this.chainName = $.chainName;
        this.description = $.description;
        this.instanceId = $.instanceId;
        this.repoName = $.repoName;
        this.repoNamespaceName = $.repoNamespaceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChainState $;

        public Builder() {
            $ = new ChainState();
        }

        public Builder(ChainState defaults) {
            $ = new ChainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param chainConfigs The configuration of delivery chain. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder chainConfigs(@Nullable Output<List<ChainChainConfigArgs>> chainConfigs) {
            $.chainConfigs = chainConfigs;
            return this;
        }

        /**
         * @param chainConfigs The configuration of delivery chain. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder chainConfigs(List<ChainChainConfigArgs> chainConfigs) {
            return chainConfigs(Output.of(chainConfigs));
        }

        /**
         * @param chainConfigs The configuration of delivery chain. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder chainConfigs(ChainChainConfigArgs... chainConfigs) {
            return chainConfigs(List.of(chainConfigs));
        }

        /**
         * @param chainId Delivery chain ID.
         * 
         * @return builder
         * 
         */
        public Builder chainId(@Nullable Output<String> chainId) {
            $.chainId = chainId;
            return this;
        }

        /**
         * @param chainId Delivery chain ID.
         * 
         * @return builder
         * 
         */
        public Builder chainId(String chainId) {
            return chainId(Output.of(chainId));
        }

        /**
         * @param chainName The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators &#34;_&#34;, &#34;-&#34;, &#34;.&#34; can be used, noted that the separator cannot be at the first or last position.
         * 
         * @return builder
         * 
         */
        public Builder chainName(@Nullable Output<String> chainName) {
            $.chainName = chainName;
            return this;
        }

        /**
         * @param chainName The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators &#34;_&#34;, &#34;-&#34;, &#34;.&#34; can be used, noted that the separator cannot be at the first or last position.
         * 
         * @return builder
         * 
         */
        public Builder chainName(String chainName) {
            return chainName(Output.of(chainName));
        }

        /**
         * @param description The description delivery chain.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description delivery chain.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param instanceId The ID of CR Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of CR Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param repoName The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder repoName(@Nullable Output<String> repoName) {
            $.repoName = repoName;
            return this;
        }

        /**
         * @param repoName The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder repoName(String repoName) {
            return repoName(Output.of(repoName));
        }

        /**
         * @param repoNamespaceName The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder repoNamespaceName(@Nullable Output<String> repoNamespaceName) {
            $.repoNamespaceName = repoNamespaceName;
            return this;
        }

        /**
         * @param repoNamespaceName The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
         * 
         * @return builder
         * 
         */
        public Builder repoNamespaceName(String repoNamespaceName) {
            return repoNamespaceName(Output.of(repoNamespaceName));
        }

        public ChainState build() {
            return $;
        }
    }

}
