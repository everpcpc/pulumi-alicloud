// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.OssShipperArgs;
import com.pulumi.alicloud.log.inputs.OssShipperState;
import com.pulumi.alicloud.log.outputs.OssShipperParquetConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Log service data delivery management, this service provides the function of delivering data in logstore to oss product storage.
 * [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/43724.htm).
 * 
 * &gt; **NOTE:** Available in 1.121.0+
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
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.log.OssShipper;
 * import com.pulumi.alicloud.log.OssShipperArgs;
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;created by terraform&#34;)
 *             .tags(Map.of(&#34;test&#34;, &#34;test&#34;))
 *             .build());
 * 
 *         var exampleStore = new Store(&#34;exampleStore&#34;, StoreArgs.builder()        
 *             .project(exampleProject.name())
 *             .retentionPeriod(3650)
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         var exampleOssShipper = new OssShipper(&#34;exampleOssShipper&#34;, OssShipperArgs.builder()        
 *             .projectName(exampleProject.name())
 *             .logstoreName(exampleStore.name())
 *             .shipperName(&#34;oss_shipper_name&#34;)
 *             .ossBucket(&#34;test_bucket&#34;)
 *             .ossPrefix(&#34;root&#34;)
 *             .bufferInterval(300)
 *             .bufferSize(250)
 *             .compressType(&#34;none&#34;)
 *             .pathFormat(&#34;%Y/%m/%d/%H/%M&#34;)
 *             .format(&#34;json&#34;)
 *             .jsonEnableTag(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Log oss shipper can be imported using the id or name, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:log/ossShipper:OssShipper example tf-log-project:tf-log-logstore:tf-log-shipper
 * ```
 * 
 */
@ResourceType(type="alicloud:log/ossShipper:OssShipper")
public class OssShipper extends com.pulumi.resources.CustomResource {
    /**
     * How often is it delivered every interval.
     * 
     */
    @Export(name="bufferInterval", type=Integer.class, parameters={})
    private Output<Integer> bufferInterval;

    /**
     * @return How often is it delivered every interval.
     * 
     */
    public Output<Integer> bufferInterval() {
        return this.bufferInterval;
    }
    /**
     * Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     * 
     */
    @Export(name="bufferSize", type=Integer.class, parameters={})
    private Output<Integer> bufferSize;

    /**
     * @return Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
     * 
     */
    public Output<Integer> bufferSize() {
        return this.bufferSize;
    }
    /**
     * OSS data storage compression method, support: none, snappy. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     * 
     */
    @Export(name="compressType", type=String.class, parameters={})
    private Output</* @Nullable */ String> compressType;

    /**
     * @return OSS data storage compression method, support: none, snappy. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
     * 
     */
    public Output<Optional<String>> compressType() {
        return Codegen.optional(this.compressType);
    }
    @Export(name="csvConfigColumns", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> csvConfigColumns;

    public Output<Optional<List<String>>> csvConfigColumns() {
        return Codegen.optional(this.csvConfigColumns);
    }
    @Export(name="csvConfigDelimiter", type=String.class, parameters={})
    private Output</* @Nullable */ String> csvConfigDelimiter;

    public Output<Optional<String>> csvConfigDelimiter() {
        return Codegen.optional(this.csvConfigDelimiter);
    }
    @Export(name="csvConfigHeader", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> csvConfigHeader;

    public Output<Optional<Boolean>> csvConfigHeader() {
        return Codegen.optional(this.csvConfigHeader);
    }
    @Export(name="csvConfigLinefeed", type=String.class, parameters={})
    private Output</* @Nullable */ String> csvConfigLinefeed;

    public Output<Optional<String>> csvConfigLinefeed() {
        return Codegen.optional(this.csvConfigLinefeed);
    }
    @Export(name="csvConfigNullidentifier", type=String.class, parameters={})
    private Output</* @Nullable */ String> csvConfigNullidentifier;

    public Output<Optional<String>> csvConfigNullidentifier() {
        return Codegen.optional(this.csvConfigNullidentifier);
    }
    @Export(name="csvConfigQuote", type=String.class, parameters={})
    private Output</* @Nullable */ String> csvConfigQuote;

    public Output<Optional<String>> csvConfigQuote() {
        return Codegen.optional(this.csvConfigQuote);
    }
    /**
     * Storage format, only supports three types: `json`, `parquet`, `csv`.
     * **According to the different format, please select the following parameters**
     * - format = `json`
     *   `json_enable_tag` - (Optional) Whether to deliver the label.
     * - format = `csv`
     *   `csv_config_delimiter` - (Optional) Separator configuration in csv configuration format.
     *   `csv_config_columns` - (Optional) Field configuration in csv configuration format.
     *   `csv_config_nullidentifier` - (Optional) Invalid field content.
     *   `csv_config_quote` - (Optional) Escape character under csv configuration.
     *   `csv_config_header` - (Optional) Indicates whether to write the field name to the CSV file, the default value is `false`.
     *   `csv_config_linefeed` - (Optional) lineFeed in csv configuration.
     * - format = `parquet`
     *   `parquet_config` - (Optional) Configure to use parquet storage format.
     *   `name` - (Required) The name of the key.
     *   `type` - (Required) Type of configuration name.
     * 
     */
    @Export(name="format", type=String.class, parameters={})
    private Output<String> format;

    /**
     * @return Storage format, only supports three types: `json`, `parquet`, `csv`.
     * **According to the different format, please select the following parameters**
     * - format = `json`
     *   `json_enable_tag` - (Optional) Whether to deliver the label.
     * - format = `csv`
     *   `csv_config_delimiter` - (Optional) Separator configuration in csv configuration format.
     *   `csv_config_columns` - (Optional) Field configuration in csv configuration format.
     *   `csv_config_nullidentifier` - (Optional) Invalid field content.
     *   `csv_config_quote` - (Optional) Escape character under csv configuration.
     *   `csv_config_header` - (Optional) Indicates whether to write the field name to the CSV file, the default value is `false`.
     *   `csv_config_linefeed` - (Optional) lineFeed in csv configuration.
     * - format = `parquet`
     *   `parquet_config` - (Optional) Configure to use parquet storage format.
     *   `name` - (Required) The name of the key.
     *   `type` - (Required) Type of configuration name.
     * 
     */
    public Output<String> format() {
        return this.format;
    }
    @Export(name="jsonEnableTag", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> jsonEnableTag;

    public Output<Optional<Boolean>> jsonEnableTag() {
        return Codegen.optional(this.jsonEnableTag);
    }
    /**
     * The name of the log logstore.
     * 
     */
    @Export(name="logstoreName", type=String.class, parameters={})
    private Output<String> logstoreName;

    /**
     * @return The name of the log logstore.
     * 
     */
    public Output<String> logstoreName() {
        return this.logstoreName;
    }
    /**
     * The name of the oss bucket.
     * 
     */
    @Export(name="ossBucket", type=String.class, parameters={})
    private Output<String> ossBucket;

    /**
     * @return The name of the oss bucket.
     * 
     */
    public Output<String> ossBucket() {
        return this.ossBucket;
    }
    /**
     * The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     * 
     */
    @Export(name="ossPrefix", type=String.class, parameters={})
    private Output</* @Nullable */ String> ossPrefix;

    /**
     * @return The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
     * 
     */
    public Output<Optional<String>> ossPrefix() {
        return Codegen.optional(this.ossPrefix);
    }
    @Export(name="parquetConfigs", type=List.class, parameters={OssShipperParquetConfig.class})
    private Output</* @Nullable */ List<OssShipperParquetConfig>> parquetConfigs;

    public Output<Optional<List<OssShipperParquetConfig>>> parquetConfigs() {
        return Codegen.optional(this.parquetConfigs);
    }
    /**
     * The OSS Bucket directory is dynamically generated according to the creation time of the shipper task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
     * 
     */
    @Export(name="pathFormat", type=String.class, parameters={})
    private Output<String> pathFormat;

    /**
     * @return The OSS Bucket directory is dynamically generated according to the creation time of the shipper task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
     * 
     */
    public Output<String> pathFormat() {
        return this.pathFormat;
    }
    /**
     * The name of the log project. It is the only in one Alicloud account.
     * 
     */
    @Export(name="projectName", type=String.class, parameters={})
    private Output<String> projectName;

    /**
     * @return The name of the log project. It is the only in one Alicloud account.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }
    /**
     * Used for access control, the OSS Bucket owner creates the role mark, such as `acs:ram::13234:role/logrole`
     * 
     */
    @Export(name="roleArn", type=String.class, parameters={})
    private Output</* @Nullable */ String> roleArn;

    /**
     * @return Used for access control, the OSS Bucket owner creates the role mark, such as `acs:ram::13234:role/logrole`
     * 
     */
    public Output<Optional<String>> roleArn() {
        return Codegen.optional(this.roleArn);
    }
    /**
     * Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     * 
     */
    @Export(name="shipperName", type=String.class, parameters={})
    private Output<String> shipperName;

    /**
     * @return Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
     * 
     */
    public Output<String> shipperName() {
        return this.shipperName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OssShipper(String name) {
        this(name, OssShipperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OssShipper(String name, OssShipperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OssShipper(String name, OssShipperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/ossShipper:OssShipper", name, args == null ? OssShipperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OssShipper(String name, Output<String> id, @Nullable OssShipperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/ossShipper:OssShipper", name, state, makeResourceOptions(options, id));
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
    public static OssShipper get(String name, Output<String> id, @Nullable OssShipperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OssShipper(name, id, state, options);
    }
}