// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kms.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPlaintextPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPlaintextPlainArgs Empty = new GetPlaintextPlainArgs();

    /**
     * The ciphertext to be decrypted.
     * 
     */
    @Import(name="ciphertextBlob", required=true)
    private String ciphertextBlob;

    /**
     * @return The ciphertext to be decrypted.
     * 
     */
    public String ciphertextBlob() {
        return this.ciphertextBlob;
    }

    /**
     * (Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
     * 
     */
    @Import(name="encryptionContext")
    private @Nullable Map<String,String> encryptionContext;

    /**
     * @return (Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
     * 
     */
    public Optional<Map<String,String>> encryptionContext() {
        return Optional.ofNullable(this.encryptionContext);
    }

    private GetPlaintextPlainArgs() {}

    private GetPlaintextPlainArgs(GetPlaintextPlainArgs $) {
        this.ciphertextBlob = $.ciphertextBlob;
        this.encryptionContext = $.encryptionContext;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPlaintextPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPlaintextPlainArgs $;

        public Builder() {
            $ = new GetPlaintextPlainArgs();
        }

        public Builder(GetPlaintextPlainArgs defaults) {
            $ = new GetPlaintextPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ciphertextBlob The ciphertext to be decrypted.
         * 
         * @return builder
         * 
         */
        public Builder ciphertextBlob(String ciphertextBlob) {
            $.ciphertextBlob = ciphertextBlob;
            return this;
        }

        /**
         * @param encryptionContext (Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
         * 
         * @return builder
         * 
         */
        public Builder encryptionContext(@Nullable Map<String,String> encryptionContext) {
            $.encryptionContext = encryptionContext;
            return this;
        }

        public GetPlaintextPlainArgs build() {
            $.ciphertextBlob = Objects.requireNonNull($.ciphertextBlob, "expected parameter 'ciphertextBlob' to be non-null");
            return $;
        }
    }

}
