// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleRuleConditionCookieConfigValueArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleRuleConditionCookieConfigValueArgs Empty = new RuleRuleConditionCookieConfigValueArgs();

    /**
     * The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The value must be 1 to 128 characters in length, and can contain lowercase letters, printable characters, asterisks (*), and question marks (?). The value cannot contain spaces or the following special characters: # [ ] { } \ | &lt; &gt; &amp;.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value must be 1 to 128 characters in length, and can contain lowercase letters, printable characters, asterisks (*), and question marks (?). The value cannot contain spaces or the following special characters: # [ ] { } \ | &lt; &gt; &amp;.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private RuleRuleConditionCookieConfigValueArgs() {}

    private RuleRuleConditionCookieConfigValueArgs(RuleRuleConditionCookieConfigValueArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleRuleConditionCookieConfigValueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleRuleConditionCookieConfigValueArgs $;

        public Builder() {
            $ = new RuleRuleConditionCookieConfigValueArgs();
        }

        public Builder(RuleRuleConditionCookieConfigValueArgs defaults) {
            $ = new RuleRuleConditionCookieConfigValueArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The value must be 1 to 128 characters in length, and can contain lowercase letters, printable characters, asterisks (*), and question marks (?). The value cannot contain spaces or the following special characters: # [ ] { } \ | &lt; &gt; &amp;.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value must be 1 to 128 characters in length, and can contain lowercase letters, printable characters, asterisks (*), and question marks (?). The value cannot contain spaces or the following special characters: # [ ] { } \ | &lt; &gt; &amp;.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public RuleRuleConditionCookieConfigValueArgs build() {
            return $;
        }
    }

}
