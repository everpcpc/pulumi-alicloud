// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AclAclEntry {
    /**
     * @return The description of the ACL entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_). It can also contain Chinese characters.
     * 
     */
    private @Nullable String description;
    /**
     * @return The IP address for the ACL entry.
     * 
     */
    private @Nullable String entry;
    /**
     * @return The state of the ACL. Valid values:`Provisioning`, `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     * 
     */
    private @Nullable String status;

    private AclAclEntry() {}
    /**
     * @return The description of the ACL entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_). It can also contain Chinese characters.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The IP address for the ACL entry.
     * 
     */
    public Optional<String> entry() {
        return Optional.ofNullable(this.entry);
    }
    /**
     * @return The state of the ACL. Valid values:`Provisioning`, `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AclAclEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String entry;
        private @Nullable String status;
        public Builder() {}
        public Builder(AclAclEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.entry = defaults.entry;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder entry(@Nullable String entry) {
            this.entry = entry;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public AclAclEntry build() {
            final var o = new AclAclEntry();
            o.description = description;
            o.entry = entry;
            o.status = status;
            return o;
        }
    }
}
