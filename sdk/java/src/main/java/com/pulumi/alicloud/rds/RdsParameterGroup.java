// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.RdsParameterGroupArgs;
import com.pulumi.alicloud.rds.inputs.RdsParameterGroupState;
import com.pulumi.alicloud.rds.outputs.RdsParameterGroupParamDetail;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RDS Parameter Group resource.
 * 
 * For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/doc-detail/144839.htm).
 * 
 * &gt; **NOTE:** Available in v1.119.0+.
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
 * import com.pulumi.alicloud.rds.RdsParameterGroup;
 * import com.pulumi.alicloud.rds.RdsParameterGroupArgs;
 * import com.pulumi.alicloud.rds.inputs.RdsParameterGroupParamDetailArgs;
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
 *         var default_ = new RdsParameterGroup(&#34;default&#34;, RdsParameterGroupArgs.builder()        
 *             .engine(&#34;mysql&#34;)
 *             .engineVersion(&#34;5.7&#34;)
 *             .paramDetails(            
 *                 RdsParameterGroupParamDetailArgs.builder()
 *                     .paramName(&#34;back_log&#34;)
 *                     .paramValue(&#34;4000&#34;)
 *                     .build(),
 *                 RdsParameterGroupParamDetailArgs.builder()
 *                     .paramName(&#34;wait_timeout&#34;)
 *                     .paramValue(&#34;86460&#34;)
 *                     .build())
 *             .parameterGroupDesc(&#34;test&#34;)
 *             .parameterGroupName(&#34;test1234&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RDS Parameter Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rds/rdsParameterGroup:RdsParameterGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/rdsParameterGroup:RdsParameterGroup")
public class RdsParameterGroup extends com.pulumi.resources.CustomResource {
    /**
     * The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
     * 
     */
    @Export(name="engine", type=String.class, parameters={})
    private Output<String> engine;

    /**
     * @return The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
     * 
     */
    @Export(name="engineVersion", type=String.class, parameters={})
    private Output<String> engineVersion;

    /**
     * @return The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * Parameter list.
     * 
     */
    @Export(name="paramDetails", type=List.class, parameters={RdsParameterGroupParamDetail.class})
    private Output<List<RdsParameterGroupParamDetail>> paramDetails;

    /**
     * @return Parameter list.
     * 
     */
    public Output<List<RdsParameterGroupParamDetail>> paramDetails() {
        return this.paramDetails;
    }
    /**
     * The description of the parameter template.
     * 
     */
    @Export(name="parameterGroupDesc", type=String.class, parameters={})
    private Output</* @Nullable */ String> parameterGroupDesc;

    /**
     * @return The description of the parameter template.
     * 
     */
    public Output<Optional<String>> parameterGroupDesc() {
        return Codegen.optional(this.parameterGroupDesc);
    }
    /**
     * The name of the parameter template.
     * 
     */
    @Export(name="parameterGroupName", type=String.class, parameters={})
    private Output<String> parameterGroupName;

    /**
     * @return The name of the parameter template.
     * 
     */
    public Output<String> parameterGroupName() {
        return this.parameterGroupName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RdsParameterGroup(String name) {
        this(name, RdsParameterGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RdsParameterGroup(String name, RdsParameterGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RdsParameterGroup(String name, RdsParameterGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsParameterGroup:RdsParameterGroup", name, args == null ? RdsParameterGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RdsParameterGroup(String name, Output<String> id, @Nullable RdsParameterGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsParameterGroup:RdsParameterGroup", name, state, makeResourceOptions(options, id));
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
    public static RdsParameterGroup get(String name, Output<String> id, @Nullable RdsParameterGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RdsParameterGroup(name, id, state, options);
    }
}
