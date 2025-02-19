// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.ConnectionArgs;
import com.pulumi.alicloud.rds.inputs.ConnectionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an RDS connection resource to allocate an Internet connection string for RDS instance.
 * 
 * &gt; **NOTE:** Each RDS instance will allocate a intranet connnection string automatically and its prifix is RDS instance ID.
 *  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
 * 
 * ## Example Usage
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
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.Connection;
 * import com.pulumi.alicloud.rds.ConnectionArgs;
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
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;Rds&#34;);
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;dbconnectionbasic&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.6&#34;)
 *             .instanceType(&#34;rds.mysql.t1.small&#34;)
 *             .instanceStorage(&#34;10&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .instanceName(name)
 *             .build());
 * 
 *         var foo = new Connection(&#34;foo&#34;, ConnectionArgs.builder()        
 *             .instanceId(instance.id())
 *             .connectionPrefix(&#34;testabc&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RDS connection can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rds/connection:Connection example abc12345678
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
     * 
     * &gt; **NOTE:** This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/babelfish-for-pg).
     * 
     */
    @Export(name="babelfishPort", type=String.class, parameters={})
    private Output<String> babelfishPort;

    /**
     * @return The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
     * 
     * &gt; **NOTE:** This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/babelfish-for-pg).
     * 
     */
    public Output<String> babelfishPort() {
        return this.babelfishPort;
    }
    /**
     * Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 40 characters. Default to &lt;instance_id&gt; + &#39;tf&#39;.
     * 
     */
    @Export(name="connectionPrefix", type=String.class, parameters={})
    private Output<String> connectionPrefix;

    /**
     * @return Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 40 characters. Default to &lt;instance_id&gt; + &#39;tf&#39;.
     * 
     */
    public Output<String> connectionPrefix() {
        return this.connectionPrefix;
    }
    /**
     * Connection instance string.
     * 
     */
    @Export(name="connectionString", type=String.class, parameters={})
    private Output<String> connectionString;

    /**
     * @return Connection instance string.
     * 
     */
    public Output<String> connectionString() {
        return this.connectionString;
    }
    /**
     * The Id of instance that can run database.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The Id of instance that can run database.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The ip address of connection string.
     * 
     */
    @Export(name="ipAddress", type=String.class, parameters={})
    private Output<String> ipAddress;

    /**
     * @return The ip address of connection string.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * Internet connection port. Valid value: [1000-5999]. Default to 3306.
     * 
     */
    @Export(name="port", type=String.class, parameters={})
    private Output</* @Nullable */ String> port;

    /**
     * @return Internet connection port. Valid value: [1000-5999]. Default to 3306.
     * 
     */
    public Output<Optional<String>> port() {
        return Codegen.optional(this.port);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/connection:Connection", name, state, makeResourceOptions(options, id));
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
