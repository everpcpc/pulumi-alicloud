// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh;

import com.pulumi.alicloud.servicemesh.inputs.UserPermissionPermissionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserPermissionArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserPermissionArgs Empty = new UserPermissionArgs();

    /**
     * List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<List<UserPermissionPermissionArgs>> permissions;

    /**
     * @return List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
     * 
     */
    public Optional<Output<List<UserPermissionPermissionArgs>>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * The configuration of the Load Balancer. See the following `Block load_balancer`.
     * 
     */
    @Import(name="subAccountUserId", required=true)
    private Output<String> subAccountUserId;

    /**
     * @return The configuration of the Load Balancer. See the following `Block load_balancer`.
     * 
     */
    public Output<String> subAccountUserId() {
        return this.subAccountUserId;
    }

    private UserPermissionArgs() {}

    private UserPermissionArgs(UserPermissionArgs $) {
        this.permissions = $.permissions;
        this.subAccountUserId = $.subAccountUserId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserPermissionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserPermissionArgs $;

        public Builder() {
            $ = new UserPermissionArgs();
        }

        public Builder(UserPermissionArgs defaults) {
            $ = new UserPermissionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param permissions List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<List<UserPermissionPermissionArgs>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
         * 
         * @return builder
         * 
         */
        public Builder permissions(List<UserPermissionPermissionArgs> permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param permissions List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
         * 
         * @return builder
         * 
         */
        public Builder permissions(UserPermissionPermissionArgs... permissions) {
            return permissions(List.of(permissions));
        }

        /**
         * @param subAccountUserId The configuration of the Load Balancer. See the following `Block load_balancer`.
         * 
         * @return builder
         * 
         */
        public Builder subAccountUserId(Output<String> subAccountUserId) {
            $.subAccountUserId = subAccountUserId;
            return this;
        }

        /**
         * @param subAccountUserId The configuration of the Load Balancer. See the following `Block load_balancer`.
         * 
         * @return builder
         * 
         */
        public Builder subAccountUserId(String subAccountUserId) {
            return subAccountUserId(Output.of(subAccountUserId));
        }

        public UserPermissionArgs build() {
            $.subAccountUserId = Objects.requireNonNull($.subAccountUserId, "expected parameter 'subAccountUserId' to be non-null");
            return $;
        }
    }

}
