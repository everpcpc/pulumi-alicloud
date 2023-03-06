// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log.outputs;

import com.pulumi.alicloud.log.outputs.StoreIndexFieldSearchJsonKey;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class StoreIndexFieldSearch {
    /**
     * @return The alias of one field.
     * 
     */
    private @Nullable String alias;
    /**
     * @return Whether the case sensitive for the field. Default to false. It is valid when &#34;type&#34; is &#34;text&#34; or &#34;json&#34;.
     * 
     */
    private @Nullable Boolean caseSensitive;
    /**
     * @return Whether to enable field analytics. Default to true.
     * 
     */
    private @Nullable Boolean enableAnalytics;
    /**
     * @return Whether includes the chinese for the field. Default to false. It is valid when &#34;type&#34; is &#34;text&#34; or &#34;json&#34;.
     * 
     */
    private @Nullable Boolean includeChinese;
    /**
     * @return Use nested index when type is json
     * 
     */
    private @Nullable List<StoreIndexFieldSearchJsonKey> jsonKeys;
    /**
     * @return When using the json_keys field, this field is required.
     * 
     */
    private String name;
    /**
     * @return The string of several split words, like &#34;\r&#34;, &#34;#&#34;. It is valid when &#34;type&#34; is &#34;text&#34; or &#34;json&#34;.
     * 
     */
    private @Nullable String token;
    /**
     * @return The type of one field. Valid values: [&#34;long&#34;, &#34;text&#34;, &#34;double&#34;]. Default to &#34;long&#34;
     * 
     */
    private @Nullable String type;

    private StoreIndexFieldSearch() {}
    /**
     * @return The alias of one field.
     * 
     */
    public Optional<String> alias() {
        return Optional.ofNullable(this.alias);
    }
    /**
     * @return Whether the case sensitive for the field. Default to false. It is valid when &#34;type&#34; is &#34;text&#34; or &#34;json&#34;.
     * 
     */
    public Optional<Boolean> caseSensitive() {
        return Optional.ofNullable(this.caseSensitive);
    }
    /**
     * @return Whether to enable field analytics. Default to true.
     * 
     */
    public Optional<Boolean> enableAnalytics() {
        return Optional.ofNullable(this.enableAnalytics);
    }
    /**
     * @return Whether includes the chinese for the field. Default to false. It is valid when &#34;type&#34; is &#34;text&#34; or &#34;json&#34;.
     * 
     */
    public Optional<Boolean> includeChinese() {
        return Optional.ofNullable(this.includeChinese);
    }
    /**
     * @return Use nested index when type is json
     * 
     */
    public List<StoreIndexFieldSearchJsonKey> jsonKeys() {
        return this.jsonKeys == null ? List.of() : this.jsonKeys;
    }
    /**
     * @return When using the json_keys field, this field is required.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The string of several split words, like &#34;\r&#34;, &#34;#&#34;. It is valid when &#34;type&#34; is &#34;text&#34; or &#34;json&#34;.
     * 
     */
    public Optional<String> token() {
        return Optional.ofNullable(this.token);
    }
    /**
     * @return The type of one field. Valid values: [&#34;long&#34;, &#34;text&#34;, &#34;double&#34;]. Default to &#34;long&#34;
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(StoreIndexFieldSearch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String alias;
        private @Nullable Boolean caseSensitive;
        private @Nullable Boolean enableAnalytics;
        private @Nullable Boolean includeChinese;
        private @Nullable List<StoreIndexFieldSearchJsonKey> jsonKeys;
        private String name;
        private @Nullable String token;
        private @Nullable String type;
        public Builder() {}
        public Builder(StoreIndexFieldSearch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alias = defaults.alias;
    	      this.caseSensitive = defaults.caseSensitive;
    	      this.enableAnalytics = defaults.enableAnalytics;
    	      this.includeChinese = defaults.includeChinese;
    	      this.jsonKeys = defaults.jsonKeys;
    	      this.name = defaults.name;
    	      this.token = defaults.token;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder alias(@Nullable String alias) {
            this.alias = alias;
            return this;
        }
        @CustomType.Setter
        public Builder caseSensitive(@Nullable Boolean caseSensitive) {
            this.caseSensitive = caseSensitive;
            return this;
        }
        @CustomType.Setter
        public Builder enableAnalytics(@Nullable Boolean enableAnalytics) {
            this.enableAnalytics = enableAnalytics;
            return this;
        }
        @CustomType.Setter
        public Builder includeChinese(@Nullable Boolean includeChinese) {
            this.includeChinese = includeChinese;
            return this;
        }
        @CustomType.Setter
        public Builder jsonKeys(@Nullable List<StoreIndexFieldSearchJsonKey> jsonKeys) {
            this.jsonKeys = jsonKeys;
            return this;
        }
        public Builder jsonKeys(StoreIndexFieldSearchJsonKey... jsonKeys) {
            return jsonKeys(List.of(jsonKeys));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder token(@Nullable String token) {
            this.token = token;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public StoreIndexFieldSearch build() {
            final var o = new StoreIndexFieldSearch();
            o.alias = alias;
            o.caseSensitive = caseSensitive;
            o.enableAnalytics = enableAnalytics;
            o.includeChinese = includeChinese;
            o.jsonKeys = jsonKeys;
            o.name = name;
            o.token = token;
            o.type = type;
            return o;
        }
    }
}