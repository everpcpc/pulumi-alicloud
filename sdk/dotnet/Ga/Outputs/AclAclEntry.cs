// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Outputs
{

    [OutputType]
    public sealed class AclAclEntry
    {
        /// <summary>
        /// The IP entry that you want to add to the ACL.
        /// </summary>
        public readonly string? Entry;
        /// <summary>
        /// The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
        /// </summary>
        public readonly string? EntryDescription;

        [OutputConstructor]
        private AclAclEntry(
            string? entry,

            string? entryDescription)
        {
            Entry = entry;
            EntryDescription = entryDescription;
        }
    }
}