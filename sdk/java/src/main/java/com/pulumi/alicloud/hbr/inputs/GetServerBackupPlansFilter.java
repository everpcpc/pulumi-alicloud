// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServerBackupPlansFilter extends com.pulumi.resources.InvokeArgs {

    public static final GetServerBackupPlansFilter Empty = new GetServerBackupPlansFilter();

    /**
     * The key of the field to filter. Valid values: `planId`, `instanceId`, `planName`.
     * 
     */
    @Import(name="key")
    private @Nullable String key;

    /**
     * @return The key of the field to filter. Valid values: `planId`, `instanceId`, `planName`.
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Set of values that are accepted for the given field.
     * 
     */
    @Import(name="values")
    private @Nullable List<String> values;

    /**
     * @return Set of values that are accepted for the given field.
     * 
     */
    public Optional<List<String>> values() {
        return Optional.ofNullable(this.values);
    }

    private GetServerBackupPlansFilter() {}

    private GetServerBackupPlansFilter(GetServerBackupPlansFilter $) {
        this.key = $.key;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerBackupPlansFilter defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerBackupPlansFilter $;

        public Builder() {
            $ = new GetServerBackupPlansFilter();
        }

        public Builder(GetServerBackupPlansFilter defaults) {
            $ = new GetServerBackupPlansFilter(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The key of the field to filter. Valid values: `planId`, `instanceId`, `planName`.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable String key) {
            $.key = key;
            return this;
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(@Nullable List<String> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetServerBackupPlansFilter build() {
            return $;
        }
    }

}