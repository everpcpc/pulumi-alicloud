// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ram.GroupMembershipArgs;
import com.pulumi.alicloud.ram.inputs.GroupMembershipState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a RAM Group membership resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ram.Group;
 * import com.pulumi.alicloud.ram.GroupArgs;
 * import com.pulumi.alicloud.ram.User;
 * import com.pulumi.alicloud.ram.UserArgs;
 * import com.pulumi.alicloud.ram.GroupMembership;
 * import com.pulumi.alicloud.ram.GroupMembershipArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var group = new Group(&#34;group&#34;, GroupArgs.builder()        
 *             .comments(&#34;this is a group comments.&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .displayName(&#34;user_display_name&#34;)
 *             .mobile(&#34;86-18688888888&#34;)
 *             .email(&#34;hello.uuu@aaa.com&#34;)
 *             .comments(&#34;yoyoyo&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var user1 = new User(&#34;user1&#34;, UserArgs.builder()        
 *             .displayName(&#34;user_display_name1&#34;)
 *             .mobile(&#34;86-18688888889&#34;)
 *             .email(&#34;hello.uuu@aaa.com&#34;)
 *             .comments(&#34;yoyoyo&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var membership = new GroupMembership(&#34;membership&#34;, GroupMembershipArgs.builder()        
 *             .groupName(group.name())
 *             .userNames(            
 *                 user.name(),
 *                 user1.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RAM Group membership can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ram/groupMembership:GroupMembership example my-group
 * ```
 * 
 */
@ResourceType(type="alicloud:ram/groupMembership:GroupMembership")
public class GroupMembership extends com.pulumi.resources.CustomResource {
    /**
     * Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen &#34;-&#34;, and must not begin with a hyphen.
     * 
     */
    @Export(name="groupName", type=String.class, parameters={})
    private Output<String> groupName;

    /**
     * @return Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen &#34;-&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }
    /**
     * Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    @Export(name="userNames", type=List.class, parameters={String.class})
    private Output<List<String>> userNames;

    /**
     * @return Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<List<String>> userNames() {
        return this.userNames;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupMembership(String name) {
        this(name, GroupMembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMembership(String name, GroupMembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMembership(String name, GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/groupMembership:GroupMembership", name, args == null ? GroupMembershipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupMembership(String name, Output<String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/groupMembership:GroupMembership", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static GroupMembership get(String name, Output<String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMembership(name, id, state, options);
    }
}
