// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Outputs
{

    [OutputType]
    public sealed class ProviderEndpoint
    {
        public readonly string? Actiontrail;
        public readonly string? Adb;
        public readonly string? Alb;
        public readonly string? Alidns;
        public readonly string? Alikafka;
        public readonly string? Apigateway;
        public readonly string? Arms;
        public readonly string? BrainIndustrial;
        public readonly string? Bssopenapi;
        public readonly string? Cas;
        public readonly string? Cassandra;
        public readonly string? Cbn;
        public readonly string? Cdn;
        public readonly string? Cds;
        public readonly string? Cloudphone;
        public readonly string? Cms;
        public readonly string? Config;
        public readonly string? Cr;
        public readonly string? Cs;
        public readonly string? Datahub;
        public readonly string? Dcdn;
        public readonly string? Ddosbgp;
        public readonly string? Ddoscoo;
        public readonly string? Dds;
        public readonly string? Dm;
        public readonly string? DmsEnterprise;
        public readonly string? Dns;
        public readonly string? Drds;
        public readonly string? Eci;
        public readonly string? Ecs;
        public readonly string? Eipanycast;
        public readonly string? Elasticsearch;
        public readonly string? Emr;
        public readonly string? Ess;
        public readonly string? Eventbridge;
        public readonly string? Fc;
        public readonly string? Fnf;
        public readonly string? Ga;
        public readonly string? Gpdb;
        public readonly string? Gwsecd;
        public readonly string? Hbr;
        public readonly string? Hitsdb;
        public readonly string? Ims;
        public readonly string? Kms;
        public readonly string? Kvstore;
        public readonly string? Location;
        public readonly string? Log;
        public readonly string? Market;
        public readonly string? Maxcompute;
        public readonly string? Mns;
        public readonly string? Mse;
        public readonly string? Nas;
        public readonly string? Ons;
        public readonly string? Onsproxy;
        public readonly string? Oos;
        public readonly string? Oss;
        public readonly string? Ots;
        public readonly string? Polardb;
        public readonly string? Privatelink;
        public readonly string? Pvtz;
        public readonly string? Quotas;
        public readonly string? RKvstore;
        public readonly string? Ram;
        public readonly string? Rds;
        public readonly string? Redisa;
        public readonly string? Resourcemanager;
        public readonly string? Resourcesharing;
        public readonly string? Ros;
        public readonly string? Serverless;
        public readonly string? Sgw;
        public readonly string? Slb;
        public readonly string? Sts;
        public readonly string? Vpc;
        public readonly string? WafOpenapi;

        [OutputConstructor]
        private ProviderEndpoint(
            string? actiontrail,

            string? adb,

            string? alb,

            string? alidns,

            string? alikafka,

            string? apigateway,

            string? arms,

            string? brainIndustrial,

            string? bssopenapi,

            string? cas,

            string? cassandra,

            string? cbn,

            string? cdn,

            string? cds,

            string? cloudphone,

            string? cms,

            string? config,

            string? cr,

            string? cs,

            string? datahub,

            string? dcdn,

            string? ddosbgp,

            string? ddoscoo,

            string? dds,

            string? dm,

            string? dmsEnterprise,

            string? dns,

            string? drds,

            string? eci,

            string? ecs,

            string? eipanycast,

            string? elasticsearch,

            string? emr,

            string? ess,

            string? eventbridge,

            string? fc,

            string? fnf,

            string? ga,

            string? gpdb,

            string? gwsecd,

            string? hbr,

            string? hitsdb,

            string? ims,

            string? kms,

            string? kvstore,

            string? location,

            string? log,

            string? market,

            string? maxcompute,

            string? mns,

            string? mse,

            string? nas,

            string? ons,

            string? onsproxy,

            string? oos,

            string? oss,

            string? ots,

            string? polardb,

            string? privatelink,

            string? pvtz,

            string? quotas,

            string? rKvstore,

            string? ram,

            string? rds,

            string? redisa,

            string? resourcemanager,

            string? resourcesharing,

            string? ros,

            string? serverless,

            string? sgw,

            string? slb,

            string? sts,

            string? vpc,

            string? wafOpenapi)
        {
            Actiontrail = actiontrail;
            Adb = adb;
            Alb = alb;
            Alidns = alidns;
            Alikafka = alikafka;
            Apigateway = apigateway;
            Arms = arms;
            BrainIndustrial = brainIndustrial;
            Bssopenapi = bssopenapi;
            Cas = cas;
            Cassandra = cassandra;
            Cbn = cbn;
            Cdn = cdn;
            Cds = cds;
            Cloudphone = cloudphone;
            Cms = cms;
            Config = config;
            Cr = cr;
            Cs = cs;
            Datahub = datahub;
            Dcdn = dcdn;
            Ddosbgp = ddosbgp;
            Ddoscoo = ddoscoo;
            Dds = dds;
            Dm = dm;
            DmsEnterprise = dmsEnterprise;
            Dns = dns;
            Drds = drds;
            Eci = eci;
            Ecs = ecs;
            Eipanycast = eipanycast;
            Elasticsearch = elasticsearch;
            Emr = emr;
            Ess = ess;
            Eventbridge = eventbridge;
            Fc = fc;
            Fnf = fnf;
            Ga = ga;
            Gpdb = gpdb;
            Gwsecd = gwsecd;
            Hbr = hbr;
            Hitsdb = hitsdb;
            Ims = ims;
            Kms = kms;
            Kvstore = kvstore;
            Location = location;
            Log = log;
            Market = market;
            Maxcompute = maxcompute;
            Mns = mns;
            Mse = mse;
            Nas = nas;
            Ons = ons;
            Onsproxy = onsproxy;
            Oos = oos;
            Oss = oss;
            Ots = ots;
            Polardb = polardb;
            Privatelink = privatelink;
            Pvtz = pvtz;
            Quotas = quotas;
            RKvstore = rKvstore;
            Ram = ram;
            Rds = rds;
            Redisa = redisa;
            Resourcemanager = resourcemanager;
            Resourcesharing = resourcesharing;
            Ros = ros;
            Serverless = serverless;
            Sgw = sgw;
            Slb = slb;
            Sts = sts;
            Vpc = vpc;
            WafOpenapi = wafOpenapi;
        }
    }
}
