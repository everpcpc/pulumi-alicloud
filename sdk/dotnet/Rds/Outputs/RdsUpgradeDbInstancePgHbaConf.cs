// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds.Outputs
{

    [OutputType]
    public sealed class RdsUpgradeDbInstancePgHbaConf
    {
        /// <summary>
        /// The IP addresses from which the specified users can access the specified databases. If you set this parameter to 0.0.0.0/0, the specified users are allowed to access the specified databases from all IP addresses.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The name of the database that the specified users are allowed to access. If you set this parameter to all, the specified users are allowed to access all databases in the instance. If you specify multiple databases, separate the database names with commas (,).
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// The mask of the instance. If the value of the `Address` parameter is an IP address, you can use this parameter to specify the mask of the IP address.
        /// </summary>
        public readonly string? Mask;
        /// <summary>
        /// The authentication method of Lightweight Directory Access Protocol (LDAP). Valid values: `trust`, `reject`, `scram-sha-256`, `md5`, `password`, `gss`, `sspi`, `ldap`, `radius`, `cert`, `pam`.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// Optional. The value of this parameter is based on the value of the HbaItem.N.Method parameter. In this topic, LDAP is used as an example. You must configure this parameter. For more information, see [Authentication Methods](https://www.postgresql.org/docs/11/auth-methods.html).
        /// </summary>
        public readonly string? Option;
        /// <summary>
        /// The priority of an AD domain. If you set this parameter to 0, the AD domain has the highest priority. Valid values: 0 to 10000. This parameter is used to identify each AD domain. When you add an AD domain, the value of the PriorityId parameter of the new AD domain cannot be the same as the value of the PriorityId parameter for any existing AD domain. When you modify or delete an AD domain, you must also modify or delete the value of the PriorityId parameter for this AD domain.
        /// </summary>
        public readonly int PriorityId;
        /// <summary>
        /// The type of connection to the instance. Valid values:
        /// * **host**: specifies to verify TCP/IP connections, including SSL connections and non-SSL connections.
        /// * **hostssl**: specifies to verify only TCP/IP connections that are established over SSL connections.
        /// * **hostnossl**: specifies to verify only TCP/IP connections that are established over non-SSL connections.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user that is allowed to access the instance. If you specify multiple users, separate the usernames with commas (,).
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private RdsUpgradeDbInstancePgHbaConf(
            string address,

            string database,

            string? mask,

            string method,

            string? option,

            int priorityId,

            string type,

            string user)
        {
            Address = address;
            Database = database;
            Mask = mask;
            Method = method;
            Option = option;
            PriorityId = priorityId;
            Type = type;
            User = user;
        }
    }
}