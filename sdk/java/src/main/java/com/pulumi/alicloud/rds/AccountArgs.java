// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccountArgs Empty = new AccountArgs();

    @Import(name="accountDescription")
    private @Nullable Output<String> accountDescription;

    public Optional<Output<String>> accountDescription() {
        return Optional.ofNullable(this.accountDescription);
    }

    @Import(name="accountName")
    private @Nullable Output<String> accountName;

    public Optional<Output<String>> accountName() {
        return Optional.ofNullable(this.accountName);
    }

    @Import(name="accountPassword")
    private @Nullable Output<String> accountPassword;

    public Optional<Output<String>> accountPassword() {
        return Optional.ofNullable(this.accountPassword);
    }

    @Import(name="accountType")
    private @Nullable Output<String> accountType;

    public Optional<Output<String>> accountType() {
        return Optional.ofNullable(this.accountType);
    }

    @Import(name="dbInstanceId")
    private @Nullable Output<String> dbInstanceId;

    public Optional<Output<String>> dbInstanceId() {
        return Optional.ofNullable(this.dbInstanceId);
    }

    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     * @deprecated
     * Field &#39;description&#39; has been deprecated from provider version 1.120.0. New field &#39;account_description&#39; instead.
     * 
     */
    @Deprecated /* Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead. */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     * 
     * @deprecated
     * Field &#39;description&#39; has been deprecated from provider version 1.120.0. New field &#39;account_description&#39; instead.
     * 
     */
    @Deprecated /* Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead. */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Id of instance in which account belongs.
     * 
     * @deprecated
     * Field &#39;instance_id&#39; has been deprecated from provider version 1.120.0. New field &#39;db_instance_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead. */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The Id of instance in which account belongs.
     * 
     * @deprecated
     * Field &#39;instance_id&#39; has been deprecated from provider version 1.120.0. New field &#39;db_instance_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead. */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     * 
     */
    @Import(name="kmsEncryptedPassword")
    private @Nullable Output<String> kmsEncryptedPassword;

    /**
     * @return An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     * 
     */
    public Optional<Output<String>> kmsEncryptedPassword() {
        return Optional.ofNullable(this.kmsEncryptedPassword);
    }

    /**
     * An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    @Import(name="kmsEncryptionContext")
    private @Nullable Output<Map<String,Object>> kmsEncryptionContext;

    /**
     * @return An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
     * 
     */
    public Optional<Output<Map<String,Object>>> kmsEncryptionContext() {
        return Optional.ofNullable(this.kmsEncryptionContext);
    }

    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;account_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;account_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     * @deprecated
     * Field &#39;password&#39; has been deprecated from provider version 1.120.0. New field &#39;account_password&#39; instead.
     * 
     */
    @Deprecated /* Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead. */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
     * 
     * @deprecated
     * Field &#39;password&#39; has been deprecated from provider version 1.120.0. New field &#39;account_password&#39; instead.
     * 
     */
    @Deprecated /* Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead. */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="resetPermissionFlag")
    private @Nullable Output<Boolean> resetPermissionFlag;

    public Optional<Output<Boolean>> resetPermissionFlag() {
        return Optional.ofNullable(this.resetPermissionFlag);
    }

    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.120.0. New field &#39;account_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead. */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     * 
     * @deprecated
     * Field &#39;type&#39; has been deprecated from provider version 1.120.0. New field &#39;account_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead. */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private AccountArgs() {}

    private AccountArgs(AccountArgs $) {
        this.accountDescription = $.accountDescription;
        this.accountName = $.accountName;
        this.accountPassword = $.accountPassword;
        this.accountType = $.accountType;
        this.dbInstanceId = $.dbInstanceId;
        this.description = $.description;
        this.instanceId = $.instanceId;
        this.kmsEncryptedPassword = $.kmsEncryptedPassword;
        this.kmsEncryptionContext = $.kmsEncryptionContext;
        this.name = $.name;
        this.password = $.password;
        this.resetPermissionFlag = $.resetPermissionFlag;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountArgs $;

        public Builder() {
            $ = new AccountArgs();
        }

        public Builder(AccountArgs defaults) {
            $ = new AccountArgs(Objects.requireNonNull(defaults));
        }

        public Builder accountDescription(@Nullable Output<String> accountDescription) {
            $.accountDescription = accountDescription;
            return this;
        }

        public Builder accountDescription(String accountDescription) {
            return accountDescription(Output.of(accountDescription));
        }

        public Builder accountName(@Nullable Output<String> accountName) {
            $.accountName = accountName;
            return this;
        }

        public Builder accountName(String accountName) {
            return accountName(Output.of(accountName));
        }

        public Builder accountPassword(@Nullable Output<String> accountPassword) {
            $.accountPassword = accountPassword;
            return this;
        }

        public Builder accountPassword(String accountPassword) {
            return accountPassword(Output.of(accountPassword));
        }

        public Builder accountType(@Nullable Output<String> accountType) {
            $.accountType = accountType;
            return this;
        }

        public Builder accountType(String accountType) {
            return accountType(Output.of(accountType));
        }

        public Builder dbInstanceId(@Nullable Output<String> dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        public Builder dbInstanceId(String dbInstanceId) {
            return dbInstanceId(Output.of(dbInstanceId));
        }

        /**
         * @param description Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;description&#39; has been deprecated from provider version 1.120.0. New field &#39;account_description&#39; instead.
         * 
         */
        @Deprecated /* Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead. */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;description&#39; has been deprecated from provider version 1.120.0. New field &#39;account_description&#39; instead.
         * 
         */
        @Deprecated /* Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead. */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param instanceId The Id of instance in which account belongs.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;instance_id&#39; has been deprecated from provider version 1.120.0. New field &#39;db_instance_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead. */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The Id of instance in which account belongs.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;instance_id&#39; has been deprecated from provider version 1.120.0. New field &#39;db_instance_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead. */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param kmsEncryptedPassword An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptedPassword(@Nullable Output<String> kmsEncryptedPassword) {
            $.kmsEncryptedPassword = kmsEncryptedPassword;
            return this;
        }

        /**
         * @param kmsEncryptedPassword An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptedPassword(String kmsEncryptedPassword) {
            return kmsEncryptedPassword(Output.of(kmsEncryptedPassword));
        }

        /**
         * @param kmsEncryptionContext An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptionContext(@Nullable Output<Map<String,Object>> kmsEncryptionContext) {
            $.kmsEncryptionContext = kmsEncryptionContext;
            return this;
        }

        /**
         * @param kmsEncryptionContext An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
         * 
         * @return builder
         * 
         */
        public Builder kmsEncryptionContext(Map<String,Object> kmsEncryptionContext) {
            return kmsEncryptionContext(Output.of(kmsEncryptionContext));
        }

        /**
         * @param name Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;account_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;account_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;password&#39; has been deprecated from provider version 1.120.0. New field &#39;account_password&#39; instead.
         * 
         */
        @Deprecated /* Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead. */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;password&#39; has been deprecated from provider version 1.120.0. New field &#39;account_password&#39; instead.
         * 
         */
        @Deprecated /* Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead. */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder resetPermissionFlag(@Nullable Output<Boolean> resetPermissionFlag) {
            $.resetPermissionFlag = resetPermissionFlag;
            return this;
        }

        public Builder resetPermissionFlag(Boolean resetPermissionFlag) {
            return resetPermissionFlag(Output.of(resetPermissionFlag));
        }

        /**
         * @param type Privilege type of account.
         * - Normal: Common privilege.
         * - Super: High privilege.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;type&#39; has been deprecated from provider version 1.120.0. New field &#39;account_type&#39; instead.
         * 
         */
        @Deprecated /* Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead. */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Privilege type of account.
         * - Normal: Common privilege.
         * - Super: High privilege.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;type&#39; has been deprecated from provider version 1.120.0. New field &#39;account_type&#39; instead.
         * 
         */
        @Deprecated /* Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead. */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AccountArgs build() {
            return $;
        }
    }

}
