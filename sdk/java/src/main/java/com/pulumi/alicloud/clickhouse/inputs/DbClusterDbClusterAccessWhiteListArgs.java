// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DbClusterDbClusterAccessWhiteListArgs extends com.pulumi.resources.ResourceArgs {

    public static final DbClusterDbClusterAccessWhiteListArgs Empty = new DbClusterDbClusterAccessWhiteListArgs();

    /**
     * Field `db_cluster_ip_array_attribute` has been removed from provider.
     * 
     */
    @Import(name="dbClusterIpArrayAttribute")
    private @Nullable Output<String> dbClusterIpArrayAttribute;

    /**
     * @return Field `db_cluster_ip_array_attribute` has been removed from provider.
     * 
     */
    public Optional<Output<String>> dbClusterIpArrayAttribute() {
        return Optional.ofNullable(this.dbClusterIpArrayAttribute);
    }

    /**
     * Whitelist group name.
     * 
     */
    @Import(name="dbClusterIpArrayName")
    private @Nullable Output<String> dbClusterIpArrayName;

    /**
     * @return Whitelist group name.
     * 
     */
    public Optional<Output<String>> dbClusterIpArrayName() {
        return Optional.ofNullable(this.dbClusterIpArrayName);
    }

    /**
     * The IP address list under the whitelist group.
     * 
     */
    @Import(name="securityIpList")
    private @Nullable Output<String> securityIpList;

    /**
     * @return The IP address list under the whitelist group.
     * 
     */
    public Optional<Output<String>> securityIpList() {
        return Optional.ofNullable(this.securityIpList);
    }

    private DbClusterDbClusterAccessWhiteListArgs() {}

    private DbClusterDbClusterAccessWhiteListArgs(DbClusterDbClusterAccessWhiteListArgs $) {
        this.dbClusterIpArrayAttribute = $.dbClusterIpArrayAttribute;
        this.dbClusterIpArrayName = $.dbClusterIpArrayName;
        this.securityIpList = $.securityIpList;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DbClusterDbClusterAccessWhiteListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DbClusterDbClusterAccessWhiteListArgs $;

        public Builder() {
            $ = new DbClusterDbClusterAccessWhiteListArgs();
        }

        public Builder(DbClusterDbClusterAccessWhiteListArgs defaults) {
            $ = new DbClusterDbClusterAccessWhiteListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterIpArrayAttribute Field `db_cluster_ip_array_attribute` has been removed from provider.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterIpArrayAttribute(@Nullable Output<String> dbClusterIpArrayAttribute) {
            $.dbClusterIpArrayAttribute = dbClusterIpArrayAttribute;
            return this;
        }

        /**
         * @param dbClusterIpArrayAttribute Field `db_cluster_ip_array_attribute` has been removed from provider.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterIpArrayAttribute(String dbClusterIpArrayAttribute) {
            return dbClusterIpArrayAttribute(Output.of(dbClusterIpArrayAttribute));
        }

        /**
         * @param dbClusterIpArrayName Whitelist group name.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterIpArrayName(@Nullable Output<String> dbClusterIpArrayName) {
            $.dbClusterIpArrayName = dbClusterIpArrayName;
            return this;
        }

        /**
         * @param dbClusterIpArrayName Whitelist group name.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterIpArrayName(String dbClusterIpArrayName) {
            return dbClusterIpArrayName(Output.of(dbClusterIpArrayName));
        }

        /**
         * @param securityIpList The IP address list under the whitelist group.
         * 
         * @return builder
         * 
         */
        public Builder securityIpList(@Nullable Output<String> securityIpList) {
            $.securityIpList = securityIpList;
            return this;
        }

        /**
         * @param securityIpList The IP address list under the whitelist group.
         * 
         * @return builder
         * 
         */
        public Builder securityIpList(String securityIpList) {
            return securityIpList(Output.of(securityIpList));
        }

        public DbClusterDbClusterAccessWhiteListArgs build() {
            return $;
        }
    }

}