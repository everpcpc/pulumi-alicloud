// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BucketLifecycleRuleTransitionArgs extends com.pulumi.resources.ResourceArgs {

    public static final BucketLifecycleRuleTransitionArgs Empty = new BucketLifecycleRuleTransitionArgs();

    /**
     * Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
     * 
     */
    @Import(name="createdBeforeDate")
    private @Nullable Output<String> createdBeforeDate;

    /**
     * @return Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
     * 
     */
    public Optional<Output<String>> createdBeforeDate() {
        return Optional.ofNullable(this.createdBeforeDate);
    }

    /**
     * Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     */
    @Import(name="days")
    private @Nullable Output<Integer> days;

    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     */
    public Optional<Output<Integer>> days() {
        return Optional.ofNullable(this.days);
    }

    /**
     * The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    @Import(name="storageClass")
    private @Nullable Output<String> storageClass;

    /**
     * @return The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    public Optional<Output<String>> storageClass() {
        return Optional.ofNullable(this.storageClass);
    }

    private BucketLifecycleRuleTransitionArgs() {}

    private BucketLifecycleRuleTransitionArgs(BucketLifecycleRuleTransitionArgs $) {
        this.createdBeforeDate = $.createdBeforeDate;
        this.days = $.days;
        this.storageClass = $.storageClass;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BucketLifecycleRuleTransitionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BucketLifecycleRuleTransitionArgs $;

        public Builder() {
            $ = new BucketLifecycleRuleTransitionArgs();
        }

        public Builder(BucketLifecycleRuleTransitionArgs defaults) {
            $ = new BucketLifecycleRuleTransitionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdBeforeDate Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
         * 
         * @return builder
         * 
         */
        public Builder createdBeforeDate(@Nullable Output<String> createdBeforeDate) {
            $.createdBeforeDate = createdBeforeDate;
            return this;
        }

        /**
         * @param createdBeforeDate Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
         * 
         * @return builder
         * 
         */
        public Builder createdBeforeDate(String createdBeforeDate) {
            return createdBeforeDate(Output.of(createdBeforeDate));
        }

        /**
         * @param days Specifies the number of days after object creation when the specific rule action takes effect.
         * 
         * @return builder
         * 
         */
        public Builder days(@Nullable Output<Integer> days) {
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
        public Builder storageClass(@Nullable Output<String> storageClass) {
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

        public BucketLifecycleRuleTransitionArgs build() {
            return $;
        }
    }

}
