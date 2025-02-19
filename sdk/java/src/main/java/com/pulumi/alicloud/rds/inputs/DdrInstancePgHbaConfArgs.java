// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DdrInstancePgHbaConfArgs extends com.pulumi.resources.ResourceArgs {

    public static final DdrInstancePgHbaConfArgs Empty = new DdrInstancePgHbaConfArgs();

    /**
     * The IP addresses from which the specified users can access the specified databases. If you set this parameter to 0.0.0.0/0, the specified users are allowed to access the specified databases from all IP addresses.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return The IP addresses from which the specified users can access the specified databases. If you set this parameter to 0.0.0.0/0, the specified users are allowed to access the specified databases from all IP addresses.
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * The name of the database that the specified users are allowed to access. If you set this parameter to all, the specified users are allowed to access all databases in the instance. If you specify multiple databases, separate the database names with commas (,).
     * 
     */
    @Import(name="database", required=true)
    private Output<String> database;

    /**
     * @return The name of the database that the specified users are allowed to access. If you set this parameter to all, the specified users are allowed to access all databases in the instance. If you specify multiple databases, separate the database names with commas (,).
     * 
     */
    public Output<String> database() {
        return this.database;
    }

    /**
     * The mask of the instance. If the value of the `Address` parameter is an IP address, you can use this parameter to specify the mask of the IP address.
     * 
     */
    @Import(name="mask")
    private @Nullable Output<String> mask;

    /**
     * @return The mask of the instance. If the value of the `Address` parameter is an IP address, you can use this parameter to specify the mask of the IP address.
     * 
     */
    public Optional<Output<String>> mask() {
        return Optional.ofNullable(this.mask);
    }

    /**
     * The authentication method of Lightweight Directory Access Protocol (LDAP). Valid values: `trust`, `reject`, `scram-sha-256`, `md5`, `password`, `gss`, `sspi`, `ldap`, `radius`, `cert`, `pam`.
     * 
     */
    @Import(name="method", required=true)
    private Output<String> method;

    /**
     * @return The authentication method of Lightweight Directory Access Protocol (LDAP). Valid values: `trust`, `reject`, `scram-sha-256`, `md5`, `password`, `gss`, `sspi`, `ldap`, `radius`, `cert`, `pam`.
     * 
     */
    public Output<String> method() {
        return this.method;
    }

    /**
     * Optional. The value of this parameter is based on the value of the HbaItem.N.Method parameter. In this topic, LDAP is used as an example. You must configure this parameter. For more information, see [Authentication Methods](https://www.postgresql.org/docs/11/auth-methods.html).
     * 
     */
    @Import(name="option")
    private @Nullable Output<String> option;

    /**
     * @return Optional. The value of this parameter is based on the value of the HbaItem.N.Method parameter. In this topic, LDAP is used as an example. You must configure this parameter. For more information, see [Authentication Methods](https://www.postgresql.org/docs/11/auth-methods.html).
     * 
     */
    public Optional<Output<String>> option() {
        return Optional.ofNullable(this.option);
    }

    /**
     * The priority of an AD domain. If you set this parameter to 0, the AD domain has the highest priority. Valid values: 0 to 10000. This parameter is used to identify each AD domain. When you add an AD domain, the value of the PriorityId parameter of the new AD domain cannot be the same as the value of the PriorityId parameter for any existing AD domain. When you modify or delete an AD domain, you must also modify or delete the value of the PriorityId parameter for this AD domain.
     * 
     */
    @Import(name="priorityId", required=true)
    private Output<Integer> priorityId;

    /**
     * @return The priority of an AD domain. If you set this parameter to 0, the AD domain has the highest priority. Valid values: 0 to 10000. This parameter is used to identify each AD domain. When you add an AD domain, the value of the PriorityId parameter of the new AD domain cannot be the same as the value of the PriorityId parameter for any existing AD domain. When you modify or delete an AD domain, you must also modify or delete the value of the PriorityId parameter for this AD domain.
     * 
     */
    public Output<Integer> priorityId() {
        return this.priorityId;
    }

    /**
     * The type of connection to the instance. Valid values:
     * * **host**: specifies to verify TCP/IP connections, including SSL connections and non-SSL connections.
     * * **hostssl**: specifies to verify only TCP/IP connections that are established over SSL connections.
     * * **hostnossl**: specifies to verify only TCP/IP connections that are established over non-SSL connections.
     * 
     * &gt; **NOTE:** You can set this parameter to hostssl only when SSL encryption is enabled for the instance. For more information, see [Configure SSL encryption for an ApsaraDB RDS for PostgreSQL instance](https://www.alibabacloud.com/help/en/doc-detail/229518.htm).
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of connection to the instance. Valid values:
     * * **host**: specifies to verify TCP/IP connections, including SSL connections and non-SSL connections.
     * * **hostssl**: specifies to verify only TCP/IP connections that are established over SSL connections.
     * * **hostnossl**: specifies to verify only TCP/IP connections that are established over non-SSL connections.
     * 
     * &gt; **NOTE:** You can set this parameter to hostssl only when SSL encryption is enabled for the instance. For more information, see [Configure SSL encryption for an ApsaraDB RDS for PostgreSQL instance](https://www.alibabacloud.com/help/en/doc-detail/229518.htm).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * The user that is allowed to access the instance. If you specify multiple users, separate the usernames with commas (,).
     * 
     */
    @Import(name="user", required=true)
    private Output<String> user;

    /**
     * @return The user that is allowed to access the instance. If you specify multiple users, separate the usernames with commas (,).
     * 
     */
    public Output<String> user() {
        return this.user;
    }

    private DdrInstancePgHbaConfArgs() {}

    private DdrInstancePgHbaConfArgs(DdrInstancePgHbaConfArgs $) {
        this.address = $.address;
        this.database = $.database;
        this.mask = $.mask;
        this.method = $.method;
        this.option = $.option;
        this.priorityId = $.priorityId;
        this.type = $.type;
        this.user = $.user;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DdrInstancePgHbaConfArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DdrInstancePgHbaConfArgs $;

        public Builder() {
            $ = new DdrInstancePgHbaConfArgs();
        }

        public Builder(DdrInstancePgHbaConfArgs defaults) {
            $ = new DdrInstancePgHbaConfArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The IP addresses from which the specified users can access the specified databases. If you set this parameter to 0.0.0.0/0, the specified users are allowed to access the specified databases from all IP addresses.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The IP addresses from which the specified users can access the specified databases. If you set this parameter to 0.0.0.0/0, the specified users are allowed to access the specified databases from all IP addresses.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param database The name of the database that the specified users are allowed to access. If you set this parameter to all, the specified users are allowed to access all databases in the instance. If you specify multiple databases, separate the database names with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The name of the database that the specified users are allowed to access. If you set this parameter to all, the specified users are allowed to access all databases in the instance. If you specify multiple databases, separate the database names with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param mask The mask of the instance. If the value of the `Address` parameter is an IP address, you can use this parameter to specify the mask of the IP address.
         * 
         * @return builder
         * 
         */
        public Builder mask(@Nullable Output<String> mask) {
            $.mask = mask;
            return this;
        }

        /**
         * @param mask The mask of the instance. If the value of the `Address` parameter is an IP address, you can use this parameter to specify the mask of the IP address.
         * 
         * @return builder
         * 
         */
        public Builder mask(String mask) {
            return mask(Output.of(mask));
        }

        /**
         * @param method The authentication method of Lightweight Directory Access Protocol (LDAP). Valid values: `trust`, `reject`, `scram-sha-256`, `md5`, `password`, `gss`, `sspi`, `ldap`, `radius`, `cert`, `pam`.
         * 
         * @return builder
         * 
         */
        public Builder method(Output<String> method) {
            $.method = method;
            return this;
        }

        /**
         * @param method The authentication method of Lightweight Directory Access Protocol (LDAP). Valid values: `trust`, `reject`, `scram-sha-256`, `md5`, `password`, `gss`, `sspi`, `ldap`, `radius`, `cert`, `pam`.
         * 
         * @return builder
         * 
         */
        public Builder method(String method) {
            return method(Output.of(method));
        }

        /**
         * @param option Optional. The value of this parameter is based on the value of the HbaItem.N.Method parameter. In this topic, LDAP is used as an example. You must configure this parameter. For more information, see [Authentication Methods](https://www.postgresql.org/docs/11/auth-methods.html).
         * 
         * @return builder
         * 
         */
        public Builder option(@Nullable Output<String> option) {
            $.option = option;
            return this;
        }

        /**
         * @param option Optional. The value of this parameter is based on the value of the HbaItem.N.Method parameter. In this topic, LDAP is used as an example. You must configure this parameter. For more information, see [Authentication Methods](https://www.postgresql.org/docs/11/auth-methods.html).
         * 
         * @return builder
         * 
         */
        public Builder option(String option) {
            return option(Output.of(option));
        }

        /**
         * @param priorityId The priority of an AD domain. If you set this parameter to 0, the AD domain has the highest priority. Valid values: 0 to 10000. This parameter is used to identify each AD domain. When you add an AD domain, the value of the PriorityId parameter of the new AD domain cannot be the same as the value of the PriorityId parameter for any existing AD domain. When you modify or delete an AD domain, you must also modify or delete the value of the PriorityId parameter for this AD domain.
         * 
         * @return builder
         * 
         */
        public Builder priorityId(Output<Integer> priorityId) {
            $.priorityId = priorityId;
            return this;
        }

        /**
         * @param priorityId The priority of an AD domain. If you set this parameter to 0, the AD domain has the highest priority. Valid values: 0 to 10000. This parameter is used to identify each AD domain. When you add an AD domain, the value of the PriorityId parameter of the new AD domain cannot be the same as the value of the PriorityId parameter for any existing AD domain. When you modify or delete an AD domain, you must also modify or delete the value of the PriorityId parameter for this AD domain.
         * 
         * @return builder
         * 
         */
        public Builder priorityId(Integer priorityId) {
            return priorityId(Output.of(priorityId));
        }

        /**
         * @param type The type of connection to the instance. Valid values:
         * * **host**: specifies to verify TCP/IP connections, including SSL connections and non-SSL connections.
         * * **hostssl**: specifies to verify only TCP/IP connections that are established over SSL connections.
         * * **hostnossl**: specifies to verify only TCP/IP connections that are established over non-SSL connections.
         * 
         * &gt; **NOTE:** You can set this parameter to hostssl only when SSL encryption is enabled for the instance. For more information, see [Configure SSL encryption for an ApsaraDB RDS for PostgreSQL instance](https://www.alibabacloud.com/help/en/doc-detail/229518.htm).
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of connection to the instance. Valid values:
         * * **host**: specifies to verify TCP/IP connections, including SSL connections and non-SSL connections.
         * * **hostssl**: specifies to verify only TCP/IP connections that are established over SSL connections.
         * * **hostnossl**: specifies to verify only TCP/IP connections that are established over non-SSL connections.
         * 
         * &gt; **NOTE:** You can set this parameter to hostssl only when SSL encryption is enabled for the instance. For more information, see [Configure SSL encryption for an ApsaraDB RDS for PostgreSQL instance](https://www.alibabacloud.com/help/en/doc-detail/229518.htm).
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param user The user that is allowed to access the instance. If you specify multiple users, separate the usernames with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder user(Output<String> user) {
            $.user = user;
            return this;
        }

        /**
         * @param user The user that is allowed to access the instance. If you specify multiple users, separate the usernames with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder user(String user) {
            return user(Output.of(user));
        }

        public DdrInstancePgHbaConfArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            $.method = Objects.requireNonNull($.method, "expected parameter 'method' to be non-null");
            $.priorityId = Objects.requireNonNull($.priorityId, "expected parameter 'priorityId' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.user = Objects.requireNonNull($.user, "expected parameter 'user' to be non-null");
            return $;
        }
    }

}
