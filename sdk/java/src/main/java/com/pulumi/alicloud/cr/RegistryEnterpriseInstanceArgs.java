// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegistryEnterpriseInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegistryEnterpriseInstanceArgs Empty = new RegistryEnterpriseInstanceArgs();

    /**
     * Name of your customized oss bucket. Use this bucket as instance storage if set.
     * 
     */
    @Import(name="customOssBucket")
    private @Nullable Output<String> customOssBucket;

    /**
     * @return Name of your customized oss bucket. Use this bucket as instance storage if set.
     * 
     */
    public Optional<Output<String>> customOssBucket() {
        return Optional.ofNullable(this.customOssBucket);
    }

    /**
     * Name of Container Registry Enterprise Edition instance.
     * 
     */
    @Import(name="instanceName", required=true)
    private Output<String> instanceName;

    /**
     * @return Name of Container Registry Enterprise Edition instance.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }

    /**
     * Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn&#39;t supports `Standard`.
     * 
     */
    @Import(name="instanceType", required=true)
    private Output<String> instanceType;

    /**
     * @return Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn&#39;t supports `Standard`.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }

    /**
     * An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
     * 
     */
    @Import(name="kmsEncryptedPassword")
    private @Nullable Output<String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
     * 
     */
    public Optional<Output<String>> kmsEncryptedPassword() {
        return Optional.ofNullable(this.kmsEncryptedPassword);
    }

    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Import(name="kmsEncryptionContext")
    private @Nullable Output<Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Optional<Output<Map<String,Object>>> kmsEncryptionContext() {
        return Optional.ofNullable(this.kmsEncryptionContext);
    }

    /**
     * The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
     * 
     */
    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    /**
     * @return Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
     * 
     */
    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    /**
     * Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    private RegistryEnterpriseInstanceArgs() {}

    private RegistryEnterpriseInstanceArgs(RegistryEnterpriseInstanceArgs $) {
        this.customOssBucket = $.customOssBucket;
        this.instanceName = $.instanceName;
        this.instanceType = $.instanceType;
        this.kmsEncryptedPassword = $.kmsEncryptedPassword;
        this.kmsEncryptionContext = $.kmsEncryptionContext;
        this.password = $.password;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.renewPeriod = $.renewPeriod;
        this.renewalStatus = $.renewalStatus;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegistryEnterpriseInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegistryEnterpriseInstanceArgs $;

        public Builder() {
            $ = new RegistryEnterpriseInstanceArgs();
        }

        public Builder(RegistryEnterpriseInstanceArgs defaults) {
            $ = new RegistryEnterpriseInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customOssBucket Name of your customized oss bucket. Use this bucket as instance storage if set.
         * 
         * @return builder
         * 
         */
        public Builder customOssBucket(@Nullable Output<String> customOssBucket) {
            $.customOssBucket = customOssBucket;
            return this;
        }

        /**
         * @param customOssBucket Name of your customized oss bucket. Use this bucket as instance storage if set.
         * 
         * @return builder
         * 
         */
        public Builder customOssBucket(String customOssBucket) {
            return customOssBucket(Output.of(customOssBucket));
        }

        /**
         * @param instanceName Name of Container Registry Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName Name of Container Registry Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param instanceType Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn&#39;t supports `Standard`.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn&#39;t supports `Standard`.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param kmsEncryptedPassword An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptedPassword(@Nullable Output<String> kmsEncryptedPassword) {
            $.kmsEncryptedPassword = kmsEncryptedPassword;
            return this;
        }

        /**
         * @param kmsEncryptedPassword An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptedPassword(String kmsEncryptedPassword) {
            return kmsEncryptedPassword(Output.of(kmsEncryptedPassword));
        }

        /**
         * @param kmsEncryptionContext An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptionContext(@Nullable Output<Map<String,Object>> kmsEncryptionContext) {
            $.kmsEncryptionContext = kmsEncryptionContext;
            return this;
        }

        /**
         * @param kmsEncryptionContext An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptionContext(Map<String,Object> kmsEncryptionContext) {
            return kmsEncryptionContext(Output.of(kmsEncryptionContext));
        }

        /**
         * @param password The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param paymentType Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param renewPeriod Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        /**
         * @param renewPeriod Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        /**
         * @param renewalStatus Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        public RegistryEnterpriseInstanceArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            $.instanceType = Objects.requireNonNull($.instanceType, "expected parameter 'instanceType' to be non-null");
            return $;
        }
    }

}
