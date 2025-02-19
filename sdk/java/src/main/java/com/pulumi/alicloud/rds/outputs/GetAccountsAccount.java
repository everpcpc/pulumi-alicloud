// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.outputs;

import com.pulumi.alicloud.rds.outputs.GetAccountsAccountDatabasePrivilege;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAccountsAccount {
    /**
     * @return Database description.
     * 
     */
    private String accountDescription;
    /**
     * @return Name of database account.
     * 
     */
    private String accountName;
    /**
     * @return Privilege type of account.
     * 
     */
    private String accountType;
    /**
     * @return A list of database permissions the account has.
     * 
     */
    private List<GetAccountsAccountDatabasePrivilege> databasePrivileges;
    /**
     * @return The ID of the Account.
     * 
     */
    private String id;
    /**
     * @return Whether the maximum number of databases managed by the account is exceeded.
     * 
     */
    private String privExceeded;
    /**
     * @return The status of the resource.
     * 
     */
    private String status;

    private GetAccountsAccount() {}
    /**
     * @return Database description.
     * 
     */
    public String accountDescription() {
        return this.accountDescription;
    }
    /**
     * @return Name of database account.
     * 
     */
    public String accountName() {
        return this.accountName;
    }
    /**
     * @return Privilege type of account.
     * 
     */
    public String accountType() {
        return this.accountType;
    }
    /**
     * @return A list of database permissions the account has.
     * 
     */
    public List<GetAccountsAccountDatabasePrivilege> databasePrivileges() {
        return this.databasePrivileges;
    }
    /**
     * @return The ID of the Account.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether the maximum number of databases managed by the account is exceeded.
     * 
     */
    public String privExceeded() {
        return this.privExceeded;
    }
    /**
     * @return The status of the resource.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccountsAccount defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountDescription;
        private String accountName;
        private String accountType;
        private List<GetAccountsAccountDatabasePrivilege> databasePrivileges;
        private String id;
        private String privExceeded;
        private String status;
        public Builder() {}
        public Builder(GetAccountsAccount defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountDescription = defaults.accountDescription;
    	      this.accountName = defaults.accountName;
    	      this.accountType = defaults.accountType;
    	      this.databasePrivileges = defaults.databasePrivileges;
    	      this.id = defaults.id;
    	      this.privExceeded = defaults.privExceeded;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder accountDescription(String accountDescription) {
            this.accountDescription = Objects.requireNonNull(accountDescription);
            return this;
        }
        @CustomType.Setter
        public Builder accountName(String accountName) {
            this.accountName = Objects.requireNonNull(accountName);
            return this;
        }
        @CustomType.Setter
        public Builder accountType(String accountType) {
            this.accountType = Objects.requireNonNull(accountType);
            return this;
        }
        @CustomType.Setter
        public Builder databasePrivileges(List<GetAccountsAccountDatabasePrivilege> databasePrivileges) {
            this.databasePrivileges = Objects.requireNonNull(databasePrivileges);
            return this;
        }
        public Builder databasePrivileges(GetAccountsAccountDatabasePrivilege... databasePrivileges) {
            return databasePrivileges(List.of(databasePrivileges));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder privExceeded(String privExceeded) {
            this.privExceeded = Objects.requireNonNull(privExceeded);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetAccountsAccount build() {
            final var o = new GetAccountsAccount();
            o.accountDescription = accountDescription;
            o.accountName = accountName;
            o.accountType = accountType;
            o.databasePrivileges = databasePrivileges;
            o.id = id;
            o.privExceeded = privExceeded;
            o.status = status;
            return o;
        }
    }
}
