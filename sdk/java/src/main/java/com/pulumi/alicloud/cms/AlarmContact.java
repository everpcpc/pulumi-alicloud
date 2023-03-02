// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cms.AlarmContactArgs;
import com.pulumi.alicloud.cms.inputs.AlarmContactState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates or modifies an alarm contact. For information about alarm contact and how to use it, see [What is alarm contact](https://www.alibabacloud.com/help/en/doc-detail/114923.htm).
 * 
 * &gt; **NOTE:** Available in v1.99.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.cms.AlarmContact;
 * import com.pulumi.alicloud.cms.AlarmContactArgs;
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
 *         var example = new AlarmContact(&#34;example&#34;, AlarmContactArgs.builder()        
 *             .alarmContactName(&#34;zhangsan&#34;)
 *             .channelsMail(&#34;terraform@test.com&#34;)
 *             .describe(&#34;For Test&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.cms.AlarmContact;
 * import com.pulumi.alicloud.cms.AlarmContactArgs;
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
 *         var example = new AlarmContact(&#34;example&#34;, AlarmContactArgs.builder()        
 *             .alarmContactName(&#34;zhangsan&#34;)
 *             .describe(&#34;For Test&#34;)
 *             .channelsMail(&#34;terraform@test.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Alarm contact can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cms/alarmContact:AlarmContact example abc12345
 * ```
 * 
 */
@ResourceType(type="alicloud:cms/alarmContact:AlarmContact")
public class AlarmContact extends com.pulumi.resources.CustomResource {
    /**
     * The name of the alarm contact. The length should between 2 and 40 characters.
     * 
     */
    @Export(name="alarmContactName", type=String.class, parameters={})
    private Output<String> alarmContactName;

    /**
     * @return The name of the alarm contact. The length should between 2 and 40 characters.
     * 
     */
    public Output<String> alarmContactName() {
        return this.alarmContactName;
    }
    /**
     * The TradeManager ID of the alarm contact.
     * 
     */
    @Export(name="channelsAliim", type=String.class, parameters={})
    private Output</* @Nullable */ String> channelsAliim;

    /**
     * @return The TradeManager ID of the alarm contact.
     * 
     */
    public Output<Optional<String>> channelsAliim() {
        return Codegen.optional(this.channelsAliim);
    }
    /**
     * The webhook URL of the DingTalk chatbot.
     * 
     */
    @Export(name="channelsDingWebHook", type=String.class, parameters={})
    private Output</* @Nullable */ String> channelsDingWebHook;

    /**
     * @return The webhook URL of the DingTalk chatbot.
     * 
     */
    public Output<Optional<String>> channelsDingWebHook() {
        return Codegen.optional(this.channelsDingWebHook);
    }
    /**
     * The email address of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     * 
     */
    @Export(name="channelsMail", type=String.class, parameters={})
    private Output</* @Nullable */ String> channelsMail;

    /**
     * @return The email address of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     * 
     */
    public Output<Optional<String>> channelsMail() {
        return Codegen.optional(this.channelsMail);
    }
    /**
     * The phone number of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     * 
     */
    @Export(name="channelsSms", type=String.class, parameters={})
    private Output</* @Nullable */ String> channelsSms;

    /**
     * @return The phone number of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     * 
     */
    public Output<Optional<String>> channelsSms() {
        return Codegen.optional(this.channelsSms);
    }
    /**
     * The description of the alarm contact.
     * 
     */
    @Export(name="describe", type=String.class, parameters={})
    private Output<String> describe;

    /**
     * @return The description of the alarm contact.
     * 
     */
    public Output<String> describe() {
        return this.describe;
    }
    /**
     * The language type of the alarm. Valid values: `en`, `zh-cn`.
     * 
     */
    @Export(name="lang", type=String.class, parameters={})
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language type of the alarm. Valid values: `en`, `zh-cn`.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlarmContact(String name) {
        this(name, AlarmContactArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlarmContact(String name, AlarmContactArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlarmContact(String name, AlarmContactArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/alarmContact:AlarmContact", name, args == null ? AlarmContactArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AlarmContact(String name, Output<String> id, @Nullable AlarmContactState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/alarmContact:AlarmContact", name, state, makeResourceOptions(options, id));
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
    public static AlarmContact get(String name, Output<String> id, @Nullable AlarmContactState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlarmContact(name, id, state, options);
    }
}
