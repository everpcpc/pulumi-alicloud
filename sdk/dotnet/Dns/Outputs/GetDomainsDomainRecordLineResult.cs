// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainRecordLineResult
    {
        public readonly string FatherCode;
        public readonly string LineCode;
        public readonly string LineDisplayName;
        public readonly string LineName;

        [OutputConstructor]
        private GetDomainsDomainRecordLineResult(
            string fatherCode,

            string lineCode,

            string lineDisplayName,

            string lineName)
        {
            FatherCode = fatherCode;
            LineCode = lineCode;
            LineDisplayName = lineDisplayName;
            LineName = lineName;
        }
    }
}
