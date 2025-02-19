// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class RoleAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final RoleAttachmentArgs Empty = new RoleAttachmentArgs();

    /**
     * The list of ECS instance&#39;s IDs.
     * 
     */
    @Import(name="instanceIds", required=true)
    private Output<List<String>> instanceIds;

    /**
     * @return The list of ECS instance&#39;s IDs.
     * 
     */
    public Output<List<String>> instanceIds() {
        return this.instanceIds;
    }

    /**
     * The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    @Import(name="roleName", required=true)
    private Output<String> roleName;

    /**
     * @return The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }

    private RoleAttachmentArgs() {}

    private RoleAttachmentArgs(RoleAttachmentArgs $) {
        this.instanceIds = $.instanceIds;
        this.roleName = $.roleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RoleAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RoleAttachmentArgs $;

        public Builder() {
            $ = new RoleAttachmentArgs();
        }

        public Builder(RoleAttachmentArgs defaults) {
            $ = new RoleAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceIds The list of ECS instance&#39;s IDs.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(Output<List<String>> instanceIds) {
            $.instanceIds = instanceIds;
            return this;
        }

        /**
         * @param instanceIds The list of ECS instance&#39;s IDs.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(List<String> instanceIds) {
            return instanceIds(Output.of(instanceIds));
        }

        /**
         * @param instanceIds The list of ECS instance&#39;s IDs.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(String... instanceIds) {
            return instanceIds(List.of(instanceIds));
        }

        /**
         * @param roleName The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;_&#34;, and must not begin with a hyphen.
         * 
         * @return builder
         * 
         */
        public Builder roleName(Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;_&#34;, and must not begin with a hyphen.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        public RoleAttachmentArgs build() {
            $.instanceIds = Objects.requireNonNull($.instanceIds, "expected parameter 'instanceIds' to be non-null");
            $.roleName = Objects.requireNonNull($.roleName, "expected parameter 'roleName' to be non-null");
            return $;
        }
    }

}
