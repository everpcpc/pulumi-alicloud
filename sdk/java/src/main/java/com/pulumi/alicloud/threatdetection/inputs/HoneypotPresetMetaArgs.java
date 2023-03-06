// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HoneypotPresetMetaArgs extends com.pulumi.resources.ResourceArgs {

    public static final HoneypotPresetMetaArgs Empty = new HoneypotPresetMetaArgs();

    /**
     * Burp counter.
     * 
     */
    @Import(name="burp", required=true)
    private Output<String> burp;

    /**
     * @return Burp counter.
     * 
     */
    public Output<String> burp() {
        return this.burp;
    }

    /**
     * Social traceability.
     * 
     */
    @Import(name="portraitOption")
    private @Nullable Output<Boolean> portraitOption;

    /**
     * @return Social traceability.
     * 
     */
    public Optional<Output<Boolean>> portraitOption() {
        return Optional.ofNullable(this.portraitOption);
    }

    /**
     * Git countered.
     * 
     */
    @Import(name="trojanGit")
    private @Nullable Output<String> trojanGit;

    /**
     * @return Git countered.
     * 
     */
    public Optional<Output<String>> trojanGit() {
        return Optional.ofNullable(this.trojanGit);
    }

    private HoneypotPresetMetaArgs() {}

    private HoneypotPresetMetaArgs(HoneypotPresetMetaArgs $) {
        this.burp = $.burp;
        this.portraitOption = $.portraitOption;
        this.trojanGit = $.trojanGit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HoneypotPresetMetaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HoneypotPresetMetaArgs $;

        public Builder() {
            $ = new HoneypotPresetMetaArgs();
        }

        public Builder(HoneypotPresetMetaArgs defaults) {
            $ = new HoneypotPresetMetaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param burp Burp counter.
         * 
         * @return builder
         * 
         */
        public Builder burp(Output<String> burp) {
            $.burp = burp;
            return this;
        }

        /**
         * @param burp Burp counter.
         * 
         * @return builder
         * 
         */
        public Builder burp(String burp) {
            return burp(Output.of(burp));
        }

        /**
         * @param portraitOption Social traceability.
         * 
         * @return builder
         * 
         */
        public Builder portraitOption(@Nullable Output<Boolean> portraitOption) {
            $.portraitOption = portraitOption;
            return this;
        }

        /**
         * @param portraitOption Social traceability.
         * 
         * @return builder
         * 
         */
        public Builder portraitOption(Boolean portraitOption) {
            return portraitOption(Output.of(portraitOption));
        }

        /**
         * @param trojanGit Git countered.
         * 
         * @return builder
         * 
         */
        public Builder trojanGit(@Nullable Output<String> trojanGit) {
            $.trojanGit = trojanGit;
            return this;
        }

        /**
         * @param trojanGit Git countered.
         * 
         * @return builder
         * 
         */
        public Builder trojanGit(String trojanGit) {
            return trojanGit(Output.of(trojanGit));
        }

        public HoneypotPresetMetaArgs build() {
            $.burp = Objects.requireNonNull($.burp, "expected parameter 'burp' to be non-null");
            return $;
        }
    }

}