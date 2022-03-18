// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./getAccount";
export * from "./getCallerIdentity";
export * from "./getFileCrc64Checksum";
export * from "./getMscSubContactVerificationMessage";
export * from "./getMscSubContacts";
export * from "./getMscSubSubscriptions";
export * from "./getMscSubWebhooks";
export * from "./getRegions";
export * from "./getZones";
export * from "./mscSubContract";
export * from "./mscSubSubscription";
export * from "./mscSubWebhook";
export * from "./provider";

// Export sub-modules:
import * as actiontrail from "./actiontrail";
import * as adb from "./adb";
import * as alb from "./alb";
import * as alikafka from "./alikafka";
import * as amqp from "./amqp";
import * as apigateway from "./apigateway";
import * as arms from "./arms";
import * as bastionhost from "./bastionhost";
import * as brain from "./brain";
import * as cas from "./cas";
import * as cassandra from "./cassandra";
import * as cddc from "./cddc";
import * as cdn from "./cdn";
import * as cen from "./cen";
import * as cfg from "./cfg";
import * as clickhouse from "./clickhouse";
import * as cloudauth from "./cloudauth";
import * as cloudconnect from "./cloudconnect";
import * as cloudfirewall from "./cloudfirewall";
import * as cloudsso from "./cloudsso";
import * as cloudstoragegateway from "./cloudstoragegateway";
import * as cms from "./cms";
import * as config from "./config";
import * as cr from "./cr";
import * as cs from "./cs";
import * as databasefilesystem from "./databasefilesystem";
import * as databasegateway from "./databasegateway";
import * as datahub from "./datahub";
import * as dataworks from "./dataworks";
import * as dcdn from "./dcdn";
import * as ddos from "./ddos";
import * as dds from "./dds";
import * as dfs from "./dfs";
import * as directmail from "./directmail";
import * as dms from "./dms";
import * as dns from "./dns";
import * as drds from "./drds";
import * as dts from "./dts";
import * as eais from "./eais";
import * as eci from "./eci";
import * as ecp from "./ecp";
import * as ecs from "./ecs";
import * as edas from "./edas";
import * as eds from "./eds";
import * as ehpc from "./ehpc";
import * as eipanycast from "./eipanycast";
import * as elasticsearch from "./elasticsearch";
import * as emr from "./emr";
import * as ens from "./ens";
import * as ess from "./ess";
import * as eventbridge from "./eventbridge";
import * as expressconnect from "./expressconnect";
import * as fc from "./fc";
import * as fnf from "./fnf";
import * as ga from "./ga";
import * as gpdb from "./gpdb";
import * as graphdatabase from "./graphdatabase";
import * as hbase from "./hbase";
import * as hbr from "./hbr";
import * as imm from "./imm";
import * as imp from "./imp";
import * as iot from "./iot";
import * as kms from "./kms";
import * as kvstore from "./kvstore";
import * as lindorm from "./lindorm";
import * as log from "./log";
import * as marketplace from "./marketplace";
import * as maxcompute from "./maxcompute";
import * as mhub from "./mhub";
import * as mns from "./mns";
import * as mongodb from "./mongodb";
import * as mse from "./mse";
import * as nas from "./nas";
import * as oos from "./oos";
import * as opensearch from "./opensearch";
import * as oss from "./oss";
import * as ots from "./ots";
import * as polardb from "./polardb";
import * as privatelink from "./privatelink";
import * as pvtz from "./pvtz";
import * as quickbi from "./quickbi";
import * as quotas from "./quotas";
import * as ram from "./ram";
import * as rdc from "./rdc";
import * as rds from "./rds";
import * as resourcemanager from "./resourcemanager";
import * as rocketmq from "./rocketmq";
import * as ros from "./ros";
import * as sae from "./sae";
import * as sag from "./sag";
import * as scdn from "./scdn";
import * as sddp from "./sddp";
import * as securitycenter from "./securitycenter";
import * as servicemesh from "./servicemesh";
import * as simpleapplicationserver from "./simpleapplicationserver";
import * as slb from "./slb";
import * as tsdb from "./tsdb";
import * as types from "./types";
import * as videosurveillance from "./videosurveillance";
import * as vod from "./vod";
import * as vpc from "./vpc";
import * as vpn from "./vpn";
import * as waf from "./waf";
import * as yundun from "./yundun";

export {
    actiontrail,
    adb,
    alb,
    alikafka,
    amqp,
    apigateway,
    arms,
    bastionhost,
    brain,
    cas,
    cassandra,
    cddc,
    cdn,
    cen,
    cfg,
    clickhouse,
    cloudauth,
    cloudconnect,
    cloudfirewall,
    cloudsso,
    cloudstoragegateway,
    cms,
    config,
    cr,
    cs,
    databasefilesystem,
    databasegateway,
    datahub,
    dataworks,
    dcdn,
    ddos,
    dds,
    dfs,
    directmail,
    dms,
    dns,
    drds,
    dts,
    eais,
    eci,
    ecp,
    ecs,
    edas,
    eds,
    ehpc,
    eipanycast,
    elasticsearch,
    emr,
    ens,
    ess,
    eventbridge,
    expressconnect,
    fc,
    fnf,
    ga,
    gpdb,
    graphdatabase,
    hbase,
    hbr,
    imm,
    imp,
    iot,
    kms,
    kvstore,
    lindorm,
    log,
    marketplace,
    maxcompute,
    mhub,
    mns,
    mongodb,
    mse,
    nas,
    oos,
    opensearch,
    oss,
    ots,
    polardb,
    privatelink,
    pvtz,
    quickbi,
    quotas,
    ram,
    rdc,
    rds,
    resourcemanager,
    rocketmq,
    ros,
    sae,
    sag,
    scdn,
    sddp,
    securitycenter,
    servicemesh,
    simpleapplicationserver,
    slb,
    tsdb,
    types,
    videosurveillance,
    vod,
    vpc,
    vpn,
    waf,
    yundun,
};

// Import resources to register:
import { MscSubContract } from "./mscSubContract";
import { MscSubSubscription } from "./mscSubSubscription";
import { MscSubWebhook } from "./mscSubWebhook";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:index/mscSubContract:MscSubContract":
                return new MscSubContract(name, <any>undefined, { urn })
            case "alicloud:index/mscSubSubscription:MscSubSubscription":
                return new MscSubSubscription(name, <any>undefined, { urn })
            case "alicloud:index/mscSubWebhook:MscSubWebhook":
                return new MscSubWebhook(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "index/mscSubContract", _module)
pulumi.runtime.registerResourceModule("alicloud", "index/mscSubSubscription", _module)
pulumi.runtime.registerResourceModule("alicloud", "index/mscSubWebhook", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("alicloud", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:alicloud") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
