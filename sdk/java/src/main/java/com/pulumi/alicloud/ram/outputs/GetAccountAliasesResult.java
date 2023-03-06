// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAccountAliasesResult {
    /**
     * @return Alias of the account.
     * 
     */
    private String accountAlias;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String outputFile;

    private GetAccountAliasesResult() {}
    /**
     * @return Alias of the account.
     * 
     */
    public String accountAlias() {
        return this.accountAlias;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccountAliasesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountAlias;
        private String id;
        private @Nullable String outputFile;
        public Builder() {}
        public Builder(GetAccountAliasesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountAlias = defaults.accountAlias;
    	      this.id = defaults.id;
    	      this.outputFile = defaults.outputFile;
        }

        @CustomType.Setter
        public Builder accountAlias(String accountAlias) {
            this.accountAlias = Objects.requireNonNull(accountAlias);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        public GetAccountAliasesResult build() {
            final var o = new GetAccountAliasesResult();
            o.accountAlias = accountAlias;
            o.id = id;
            o.outputFile = outputFile;
            return o;
        }
    }
}