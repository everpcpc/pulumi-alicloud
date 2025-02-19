// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alikafka;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alikafka.SaslAclArgs;
import com.pulumi.alicloud.alikafka.inputs.SaslAclState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an ALIKAFKA sasl acl resource.
 * 
 * &gt; **NOTE:** Available in 1.66.0+
 * 
 * &gt; **NOTE:**  Only the following regions support create alikafka sasl user.
 * [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.alikafka.Instance;
 * import com.pulumi.alicloud.alikafka.InstanceArgs;
 * import com.pulumi.alicloud.alikafka.Topic;
 * import com.pulumi.alicloud.alikafka.TopicArgs;
 * import com.pulumi.alicloud.alikafka.SaslUser;
 * import com.pulumi.alicloud.alikafka.SaslUserArgs;
 * import com.pulumi.alicloud.alikafka.SaslAcl;
 * import com.pulumi.alicloud.alikafka.SaslAclArgs;
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
 *         final var config = ctx.config();
 *         final var username = config.get(&#34;username&#34;).orElse(&#34;testusername&#34;);
 *         final var password = config.get(&#34;password&#34;).orElse(&#34;testpassword&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .partitionNum(&#34;50&#34;)
 *             .diskType(&#34;1&#34;)
 *             .diskSize(&#34;500&#34;)
 *             .deployType(&#34;5&#34;)
 *             .ioMax(&#34;20&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .build());
 * 
 *         var defaultTopic = new Topic(&#34;defaultTopic&#34;, TopicArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .topic(&#34;test-topic&#34;)
 *             .remark(&#34;topic-remark&#34;)
 *             .build());
 * 
 *         var defaultSaslUser = new SaslUser(&#34;defaultSaslUser&#34;, SaslUserArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .username(username)
 *             .password(password)
 *             .build());
 * 
 *         var defaultSaslAcl = new SaslAcl(&#34;defaultSaslAcl&#34;, SaslAclArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .username(defaultSaslUser.username())
 *             .aclResourceType(&#34;Topic&#34;)
 *             .aclResourceName(defaultTopic.topic())
 *             .aclResourcePatternType(&#34;LITERAL&#34;)
 *             .aclOperationType(&#34;Write&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ALIKAFKA GROUP can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:alikafka/saslAcl:SaslAcl acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write
 * ```
 * 
 */
@ResourceType(type="alicloud:alikafka/saslAcl:SaslAcl")
public class SaslAcl extends com.pulumi.resources.CustomResource {
    /**
     * Operation type for this acl. The operation type can only be &#34;Write&#34; and &#34;Read&#34;.
     * 
     */
    @Export(name="aclOperationType", type=String.class, parameters={})
    private Output<String> aclOperationType;

    /**
     * @return Operation type for this acl. The operation type can only be &#34;Write&#34; and &#34;Read&#34;.
     * 
     */
    public Output<String> aclOperationType() {
        return this.aclOperationType;
    }
    /**
     * Resource name for this acl. The resource name should be a topic or consumer group name.
     * 
     */
    @Export(name="aclResourceName", type=String.class, parameters={})
    private Output<String> aclResourceName;

    /**
     * @return Resource name for this acl. The resource name should be a topic or consumer group name.
     * 
     */
    public Output<String> aclResourceName() {
        return this.aclResourceName;
    }
    /**
     * Resource pattern type for this acl. The resource pattern support two types &#34;LITERAL&#34; and &#34;PREFIXED&#34;. &#34;LITERAL&#34;: A literal name defines the full name of a resource. The special wildcard character &#34;*&#34; can be used to represent a resource with any name. &#34;PREFIXED&#34;: A prefixed name defines a prefix for a resource.
     * 
     */
    @Export(name="aclResourcePatternType", type=String.class, parameters={})
    private Output<String> aclResourcePatternType;

    /**
     * @return Resource pattern type for this acl. The resource pattern support two types &#34;LITERAL&#34; and &#34;PREFIXED&#34;. &#34;LITERAL&#34;: A literal name defines the full name of a resource. The special wildcard character &#34;*&#34; can be used to represent a resource with any name. &#34;PREFIXED&#34;: A prefixed name defines a prefix for a resource.
     * 
     */
    public Output<String> aclResourcePatternType() {
        return this.aclResourcePatternType;
    }
    /**
     * Resource type for this acl. The resource type can only be &#34;Topic&#34; and &#34;Group&#34;.
     * 
     */
    @Export(name="aclResourceType", type=String.class, parameters={})
    private Output<String> aclResourceType;

    /**
     * @return Resource type for this acl. The resource type can only be &#34;Topic&#34; and &#34;Group&#34;.
     * 
     */
    public Output<String> aclResourceType() {
        return this.aclResourceType;
    }
    /**
     * The host of the acl.
     * 
     */
    @Export(name="host", type=String.class, parameters={})
    private Output<String> host;

    /**
     * @return The host of the acl.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * ID of the ALIKAFKA Instance that owns the groups.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return ID of the ALIKAFKA Instance that owns the groups.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SaslAcl(String name) {
        this(name, SaslAclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SaslAcl(String name, SaslAclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SaslAcl(String name, SaslAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/saslAcl:SaslAcl", name, args == null ? SaslAclArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SaslAcl(String name, Output<String> id, @Nullable SaslAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alikafka/saslAcl:SaslAcl", name, state, makeResourceOptions(options, id));
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
    public static SaslAcl get(String name, Output<String> id, @Nullable SaslAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SaslAcl(name, id, state, options);
    }
}
