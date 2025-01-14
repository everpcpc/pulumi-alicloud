// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oss.BucketArgs;
import com.pulumi.alicloud.oss.inputs.BucketState;
import com.pulumi.alicloud.oss.outputs.BucketCorsRule;
import com.pulumi.alicloud.oss.outputs.BucketLifecycleRule;
import com.pulumi.alicloud.oss.outputs.BucketLogging;
import com.pulumi.alicloud.oss.outputs.BucketRefererConfig;
import com.pulumi.alicloud.oss.outputs.BucketServerSideEncryptionRule;
import com.pulumi.alicloud.oss.outputs.BucketTransferAcceleration;
import com.pulumi.alicloud.oss.outputs.BucketVersioning;
import com.pulumi.alicloud.oss.outputs.BucketWebsite;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to create a oss bucket and set its attribution.
 * 
 * &gt; **NOTE:** The bucket namespace is shared by all users of the OSS system. Please set bucket name as unique as possible.
 * 
 * ## Example Usage
 * 
 * Private Bucket
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
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
 *         var bucket_acl = new Bucket(&#34;bucket-acl&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-acl&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Static Website
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketWebsiteArgs;
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
 *         var bucket_website = new Bucket(&#34;bucket-website&#34;, BucketArgs.builder()        
 *             .bucket(&#34;bucket-170309-website&#34;)
 *             .website(BucketWebsiteArgs.builder()
 *                 .errorDocument(&#34;error.html&#34;)
 *                 .indexDocument(&#34;index.html&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Enable Logging
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketLoggingArgs;
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
 *         var bucket_target = new Bucket(&#34;bucket-target&#34;, BucketArgs.builder()        
 *             .bucket(&#34;bucket-170309-acl&#34;)
 *             .acl(&#34;public-read&#34;)
 *             .build());
 * 
 *         var bucket_logging = new Bucket(&#34;bucket-logging&#34;, BucketArgs.builder()        
 *             .bucket(&#34;bucket-170309-logging&#34;)
 *             .logging(BucketLoggingArgs.builder()
 *                 .targetBucket(bucket_target.id())
 *                 .targetPrefix(&#34;log/&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Referer configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketRefererConfigArgs;
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
 *         var bucket_referer = new Bucket(&#34;bucket-referer&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-referer&#34;)
 *             .refererConfig(BucketRefererConfigArgs.builder()
 *                 .allowEmpty(false)
 *                 .referers(                
 *                     &#34;http://www.aliyun.com&#34;,
 *                     &#34;https://www.aliyun.com&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set lifecycle rule
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketLifecycleRuleArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketVersioningArgs;
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
 *         var bucket_lifecycle = new Bucket(&#34;bucket-lifecycle&#34;, BucketArgs.builder()        
 *             .acl(&#34;public-read&#34;)
 *             .bucket(&#34;bucket-170309-lifecycle&#34;)
 *             .lifecycleRules(BucketLifecycleRuleArgs.builder()
 *                 .abortMultipartUploads(BucketLifecycleRuleAbortMultipartUploadArgs.builder()
 *                     .days(128)
 *                     .build())
 *                 .enabled(true)
 *                 .id(&#34;rule-abort-multipart-upload&#34;)
 *                 .prefix(&#34;path3/&#34;)
 *                 .build())
 *             .build());
 * 
 *         var bucket_versioning_lifecycle = new Bucket(&#34;bucket-versioning-lifecycle&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-lifecycle&#34;)
 *             .lifecycleRules(BucketLifecycleRuleArgs.builder()
 *                 .enabled(true)
 *                 .expirations(BucketLifecycleRuleExpirationArgs.builder()
 *                     .expiredObjectDeleteMarker(true)
 *                     .build())
 *                 .id(&#34;rule-versioning&#34;)
 *                 .noncurrentVersionExpirations(BucketLifecycleRuleNoncurrentVersionExpirationArgs.builder()
 *                     .days(240)
 *                     .build())
 *                 .noncurrentVersionTransitions(                
 *                     BucketLifecycleRuleNoncurrentVersionTransitionArgs.builder()
 *                         .days(180)
 *                         .storageClass(&#34;Archive&#34;)
 *                         .build(),
 *                     BucketLifecycleRuleNoncurrentVersionTransitionArgs.builder()
 *                         .days(60)
 *                         .storageClass(&#34;IA&#34;)
 *                         .build())
 *                 .prefix(&#34;path1/&#34;)
 *                 .build())
 *             .versioning(BucketVersioningArgs.builder()
 *                 .status(&#34;Enabled&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set bucket policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
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
 *         var bucket_policy = new Bucket(&#34;bucket-policy&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-policy&#34;)
 *             .policy(&#34;&#34;&#34;
 *   {&#34;Statement&#34;:
 *       [{&#34;Action&#34;:
 *           [&#34;oss:PutObject&#34;, &#34;oss:GetObject&#34;, &#34;oss:DeleteBucket&#34;],
 *         &#34;Effect&#34;:&#34;Allow&#34;,
 *         &#34;Resource&#34;:
 *             [&#34;acs:oss:*:*:*&#34;]}],
 *    &#34;Version&#34;:&#34;1&#34;}
 *   
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * IA Bucket
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
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
 *         var bucket_storageclass = new Bucket(&#34;bucket-storageclass&#34;, BucketArgs.builder()        
 *             .bucket(&#34;bucket-170309-storageclass&#34;)
 *             .storageClass(&#34;IA&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set bucket server-side encryption rule
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketServerSideEncryptionRuleArgs;
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
 *         var bucket_sserule = new Bucket(&#34;bucket-sserule&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-sserule&#34;)
 *             .serverSideEncryptionRule(BucketServerSideEncryptionRuleArgs.builder()
 *                 .kmsMasterKeyId(&#34;your kms key id&#34;)
 *                 .sseAlgorithm(&#34;KMS&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set bucket tags
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
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
 *         var bucket_tags = new Bucket(&#34;bucket-tags&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-tags&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
 *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Enable bucket versioning
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketVersioningArgs;
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
 *         var bucket_versioning = new Bucket(&#34;bucket-versioning&#34;, BucketArgs.builder()        
 *             .acl(&#34;private&#34;)
 *             .bucket(&#34;bucket-170309-versioning&#34;)
 *             .versioning(BucketVersioningArgs.builder()
 *                 .status(&#34;Enabled&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set bucket redundancy type
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
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
 *         var bucket_redundancytype = new Bucket(&#34;bucket-redundancytype&#34;, BucketArgs.builder()        
 *             .bucket(&#34;bucket_name&#34;)
 *             .redundancyType(&#34;ZRS&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set bucket accelerate configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketTransferAccelerationArgs;
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
 *         var bucket_accelerate = new Bucket(&#34;bucket-accelerate&#34;, BucketArgs.builder()        
 *             .bucket(&#34;bucket_name&#34;)
 *             .transferAcceleration(BucketTransferAccelerationArgs.builder()
 *                 .enabled(false)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OSS bucket can be imported using the bucket name, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:oss/bucket:Bucket bucket bucket-12345678
 * ```
 * 
 */
@ResourceType(type="alicloud:oss/bucket:Bucket")
public class Bucket extends com.pulumi.resources.CustomResource {
    /**
     * The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be &#34;private&#34;, &#34;public-read&#34; and &#34;public-read-write&#34;. Defaults to &#34;private&#34;.
     * 
     */
    @Export(name="acl", type=String.class, parameters={})
    private Output</* @Nullable */ String> acl;

    /**
     * @return The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be &#34;private&#34;, &#34;public-read&#34; and &#34;public-read-write&#34;. Defaults to &#34;private&#34;.
     * 
     */
    public Output<Optional<String>> acl() {
        return Codegen.optional(this.acl);
    }
    @Export(name="bucket", type=String.class, parameters={})
    private Output</* @Nullable */ String> bucket;

    public Output<Optional<String>> bucket() {
        return Codegen.optional(this.bucket);
    }
    /**
     * A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
     * 
     */
    @Export(name="corsRules", type=List.class, parameters={BucketCorsRule.class})
    private Output</* @Nullable */ List<BucketCorsRule>> corsRules;

    /**
     * @return A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
     * 
     */
    public Output<Optional<List<BucketCorsRule>>> corsRules() {
        return Codegen.optional(this.corsRules);
    }
    /**
     * The creation date of the bucket.
     * 
     */
    @Export(name="creationDate", type=String.class, parameters={})
    private Output<String> creationDate;

    /**
     * @return The creation date of the bucket.
     * 
     */
    public Output<String> creationDate() {
        return this.creationDate;
    }
    /**
     * The extranet access endpoint of the bucket.
     * 
     */
    @Export(name="extranetEndpoint", type=String.class, parameters={})
    private Output<String> extranetEndpoint;

    /**
     * @return The extranet access endpoint of the bucket.
     * 
     */
    public Output<String> extranetEndpoint() {
        return this.extranetEndpoint;
    }
    /**
     * A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to &#34;false&#34;.
     * 
     */
    @Export(name="forceDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to &#34;false&#34;.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * The intranet access endpoint of the bucket.
     * 
     */
    @Export(name="intranetEndpoint", type=String.class, parameters={})
    private Output<String> intranetEndpoint;

    /**
     * @return The intranet access endpoint of the bucket.
     * 
     */
    public Output<String> intranetEndpoint() {
        return this.intranetEndpoint;
    }
    /**
     * A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
     * 
     */
    @Export(name="lifecycleRules", type=List.class, parameters={BucketLifecycleRule.class})
    private Output</* @Nullable */ List<BucketLifecycleRule>> lifecycleRules;

    /**
     * @return A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
     * 
     */
    public Output<Optional<List<BucketLifecycleRule>>> lifecycleRules() {
        return Codegen.optional(this.lifecycleRules);
    }
    /**
     * The location of the bucket.
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The location of the bucket.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
     * 
     */
    @Export(name="logging", type=BucketLogging.class, parameters={})
    private Output</* @Nullable */ BucketLogging> logging;

    /**
     * @return A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
     * 
     */
    public Output<Optional<BucketLogging>> logging() {
        return Codegen.optional(this.logging);
    }
    /**
     * The flag of using logging enable container. Defaults true.
     * 
     * @deprecated
     * Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
     * 
     */
    @Deprecated /* Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able. */
    @Export(name="loggingIsenable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> loggingIsenable;

    /**
     * @return The flag of using logging enable container. Defaults true.
     * 
     */
    public Output<Optional<Boolean>> loggingIsenable() {
        return Codegen.optional(this.loggingIsenable);
    }
    /**
     * The bucket owner.
     * 
     */
    @Export(name="owner", type=String.class, parameters={})
    private Output<String> owner;

    /**
     * @return The bucket owner.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
     * 
     */
    @Export(name="policy", type=String.class, parameters={})
    private Output</* @Nullable */ String> policy;

    /**
     * @return Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
     * 
     */
    public Output<Optional<String>> policy() {
        return Codegen.optional(this.policy);
    }
    /**
     * The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be &#34;LRS&#34;, and &#34;ZRS&#34;. Defaults to &#34;LRS&#34;.
     * 
     */
    @Export(name="redundancyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> redundancyType;

    /**
     * @return The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be &#34;LRS&#34;, and &#34;ZRS&#34;. Defaults to &#34;LRS&#34;.
     * 
     */
    public Output<Optional<String>> redundancyType() {
        return Codegen.optional(this.redundancyType);
    }
    /**
     * The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
     * 
     */
    @Export(name="refererConfig", type=BucketRefererConfig.class, parameters={})
    private Output</* @Nullable */ BucketRefererConfig> refererConfig;

    /**
     * @return The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
     * 
     */
    public Output<Optional<BucketRefererConfig>> refererConfig() {
        return Codegen.optional(this.refererConfig);
    }
    /**
     * A configuration of server-side encryption (documented below).
     * 
     */
    @Export(name="serverSideEncryptionRule", type=BucketServerSideEncryptionRule.class, parameters={})
    private Output</* @Nullable */ BucketServerSideEncryptionRule> serverSideEncryptionRule;

    /**
     * @return A configuration of server-side encryption (documented below).
     * 
     */
    public Output<Optional<BucketServerSideEncryptionRule>> serverSideEncryptionRule() {
        return Codegen.optional(this.serverSideEncryptionRule);
    }
    /**
     * The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    @Export(name="storageClass", type=String.class, parameters={})
    private Output</* @Nullable */ String> storageClass;

    /**
     * @return The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    public Output<Optional<String>> storageClass() {
        return Codegen.optional(this.storageClass);
    }
    /**
     * A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A transfer acceleration status of a bucket (documented below).
     * 
     */
    @Export(name="transferAcceleration", type=BucketTransferAcceleration.class, parameters={})
    private Output</* @Nullable */ BucketTransferAcceleration> transferAcceleration;

    /**
     * @return A transfer acceleration status of a bucket (documented below).
     * 
     */
    public Output<Optional<BucketTransferAcceleration>> transferAcceleration() {
        return Codegen.optional(this.transferAcceleration);
    }
    /**
     * A state of versioning (documented below).
     * 
     */
    @Export(name="versioning", type=BucketVersioning.class, parameters={})
    private Output</* @Nullable */ BucketVersioning> versioning;

    /**
     * @return A state of versioning (documented below).
     * 
     */
    public Output<Optional<BucketVersioning>> versioning() {
        return Codegen.optional(this.versioning);
    }
    /**
     * A website object(documented below).
     * 
     */
    @Export(name="website", type=BucketWebsite.class, parameters={})
    private Output</* @Nullable */ BucketWebsite> website;

    /**
     * @return A website object(documented below).
     * 
     */
    public Output<Optional<BucketWebsite>> website() {
        return Codegen.optional(this.website);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Bucket(String name) {
        this(name, BucketArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Bucket(String name, @Nullable BucketArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Bucket(String name, @Nullable BucketArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucket:Bucket", name, args == null ? BucketArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Bucket(String name, Output<String> id, @Nullable BucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucket:Bucket", name, state, makeResourceOptions(options, id));
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
    public static Bucket get(String name, Output<String> id, @Nullable BucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Bucket(name, id, state, options);
    }
}
