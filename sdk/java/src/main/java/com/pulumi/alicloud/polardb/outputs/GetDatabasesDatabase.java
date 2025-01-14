// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.outputs;

import com.pulumi.alicloud.polardb.outputs.GetDatabasesDatabaseAccount;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDatabasesDatabase {
    /**
     * @return A list of accounts of database. Each element contains the following attributes.
     * 
     */
    private List<GetDatabasesDatabaseAccount> accounts;
    /**
     * @return The character set name of database.
     * 
     */
    private String characterSetName;
    /**
     * @return Database description.
     * 
     */
    private String dbDescription;
    /**
     * @return Database name.
     * 
     */
    private String dbName;
    /**
     * @return The status of database.
     * 
     */
    private String dbStatus;
    /**
     * @return The engine of database.
     * 
     */
    private String engine;

    private GetDatabasesDatabase() {}
    /**
     * @return A list of accounts of database. Each element contains the following attributes.
     * 
     */
    public List<GetDatabasesDatabaseAccount> accounts() {
        return this.accounts;
    }
    /**
     * @return The character set name of database.
     * 
     */
    public String characterSetName() {
        return this.characterSetName;
    }
    /**
     * @return Database description.
     * 
     */
    public String dbDescription() {
        return this.dbDescription;
    }
    /**
     * @return Database name.
     * 
     */
    public String dbName() {
        return this.dbName;
    }
    /**
     * @return The status of database.
     * 
     */
    public String dbStatus() {
        return this.dbStatus;
    }
    /**
     * @return The engine of database.
     * 
     */
    public String engine() {
        return this.engine;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabasesDatabase defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDatabasesDatabaseAccount> accounts;
        private String characterSetName;
        private String dbDescription;
        private String dbName;
        private String dbStatus;
        private String engine;
        public Builder() {}
        public Builder(GetDatabasesDatabase defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accounts = defaults.accounts;
    	      this.characterSetName = defaults.characterSetName;
    	      this.dbDescription = defaults.dbDescription;
    	      this.dbName = defaults.dbName;
    	      this.dbStatus = defaults.dbStatus;
    	      this.engine = defaults.engine;
        }

        @CustomType.Setter
        public Builder accounts(List<GetDatabasesDatabaseAccount> accounts) {
            this.accounts = Objects.requireNonNull(accounts);
            return this;
        }
        public Builder accounts(GetDatabasesDatabaseAccount... accounts) {
            return accounts(List.of(accounts));
        }
        @CustomType.Setter
        public Builder characterSetName(String characterSetName) {
            this.characterSetName = Objects.requireNonNull(characterSetName);
            return this;
        }
        @CustomType.Setter
        public Builder dbDescription(String dbDescription) {
            this.dbDescription = Objects.requireNonNull(dbDescription);
            return this;
        }
        @CustomType.Setter
        public Builder dbName(String dbName) {
            this.dbName = Objects.requireNonNull(dbName);
            return this;
        }
        @CustomType.Setter
        public Builder dbStatus(String dbStatus) {
            this.dbStatus = Objects.requireNonNull(dbStatus);
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            this.engine = Objects.requireNonNull(engine);
            return this;
        }
        public GetDatabasesDatabase build() {
            final var o = new GetDatabasesDatabase();
            o.accounts = accounts;
            o.characterSetName = characterSetName;
            o.dbDescription = dbDescription;
            o.dbName = dbName;
            o.dbStatus = dbStatus;
            o.engine = engine;
            return o;
        }
    }
}
