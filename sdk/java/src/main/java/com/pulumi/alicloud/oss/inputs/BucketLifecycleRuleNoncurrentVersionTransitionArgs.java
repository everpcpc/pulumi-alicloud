// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class BucketLifecycleRuleNoncurrentVersionTransitionArgs extends com.pulumi.resources.ResourceArgs {

    public static final BucketLifecycleRuleNoncurrentVersionTransitionArgs Empty = new BucketLifecycleRuleNoncurrentVersionTransitionArgs();

    /**
     * Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     */
    @Import(name="days", required=true)
    private Output<Integer> days;

    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     */
    public Output<Integer> days() {
        return this.days;
    }

    /**
     * The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    @Import(name="storageClass", required=true)
    private Output<String> storageClass;

    /**
     * @return The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    public Output<String> storageClass() {
        return this.storageClass;
    }

    private BucketLifecycleRuleNoncurrentVersionTransitionArgs() {}

    private BucketLifecycleRuleNoncurrentVersionTransitionArgs(BucketLifecycleRuleNoncurrentVersionTransitionArgs $) {
        this.days = $.days;
        this.storageClass = $.storageClass;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BucketLifecycleRuleNoncurrentVersionTransitionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BucketLifecycleRuleNoncurrentVersionTransitionArgs $;

        public Builder() {
            $ = new BucketLifecycleRuleNoncurrentVersionTransitionArgs();
        }

        public Builder(BucketLifecycleRuleNoncurrentVersionTransitionArgs defaults) {
            $ = new BucketLifecycleRuleNoncurrentVersionTransitionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param days Specifies the number of days after object creation when the specific rule action takes effect.
         * 
         * @return builder
         * 
         */
        public Builder days(Output<Integer> days) {
            $.days = days;
            return this;
        }

        /**
         * @param days Specifies the number of days after object creation when the specific rule action takes effect.
         * 
         * @return builder
         * 
         */
        public Builder days(Integer days) {
            return days(Output.of(days));
        }

        /**
         * @param storageClass The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(Output<String> storageClass) {
            $.storageClass = storageClass;
            return this;
        }

        /**
         * @param storageClass The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(String storageClass) {
            return storageClass(Output.of(storageClass));
        }

        public BucketLifecycleRuleNoncurrentVersionTransitionArgs build() {
            $.days = Objects.requireNonNull($.days, "expected parameter 'days' to be non-null");
            $.storageClass = Objects.requireNonNull($.storageClass, "expected parameter 'storageClass' to be non-null");
            return $;
        }
    }

}
