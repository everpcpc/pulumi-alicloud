// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AclAclEntryArgs extends com.pulumi.resources.ResourceArgs {

    public static final AclAclEntryArgs Empty = new AclAclEntryArgs();

    /**
     * The IP entry that you want to add to the ACL.
     * 
     */
    @Import(name="entry")
    private @Nullable Output<String> entry;

    /**
     * @return The IP entry that you want to add to the ACL.
     * 
     */
    public Optional<Output<String>> entry() {
        return Optional.ofNullable(this.entry);
    }

    /**
     * The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
     * 
     */
    @Import(name="entryDescription")
    private @Nullable Output<String> entryDescription;

    /**
     * @return The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
     * 
     */
    public Optional<Output<String>> entryDescription() {
        return Optional.ofNullable(this.entryDescription);
    }

    private AclAclEntryArgs() {}

    private AclAclEntryArgs(AclAclEntryArgs $) {
        this.entry = $.entry;
        this.entryDescription = $.entryDescription;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AclAclEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AclAclEntryArgs $;

        public Builder() {
            $ = new AclAclEntryArgs();
        }

        public Builder(AclAclEntryArgs defaults) {
            $ = new AclAclEntryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param entry The IP entry that you want to add to the ACL.
         * 
         * @return builder
         * 
         */
        public Builder entry(@Nullable Output<String> entry) {
            $.entry = entry;
            return this;
        }

        /**
         * @param entry The IP entry that you want to add to the ACL.
         * 
         * @return builder
         * 
         */
        public Builder entry(String entry) {
            return entry(Output.of(entry));
        }

        /**
         * @param entryDescription The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
         * 
         * @return builder
         * 
         */
        public Builder entryDescription(@Nullable Output<String> entryDescription) {
            $.entryDescription = entryDescription;
            return this;
        }

        /**
         * @param entryDescription The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
         * 
         * @return builder
         * 
         */
        public Builder entryDescription(String entryDescription) {
            return entryDescription(Output.of(entryDescription));
        }

        public AclAclEntryArgs build() {
            return $;
        }
    }

}
