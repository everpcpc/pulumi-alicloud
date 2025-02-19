// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AlertContactArgs, AlertContactState } from "./alertContact";
export type AlertContact = import("./alertContact").AlertContact;
export const AlertContact: typeof import("./alertContact").AlertContact = null as any;
utilities.lazyLoad(exports, ["AlertContact"], () => require("./alertContact"));

export { AlertContactGroupArgs, AlertContactGroupState } from "./alertContactGroup";
export type AlertContactGroup = import("./alertContactGroup").AlertContactGroup;
export const AlertContactGroup: typeof import("./alertContactGroup").AlertContactGroup = null as any;
utilities.lazyLoad(exports, ["AlertContactGroup"], () => require("./alertContactGroup"));

export { DispatchRuleArgs, DispatchRuleState } from "./dispatchRule";
export type DispatchRule = import("./dispatchRule").DispatchRule;
export const DispatchRule: typeof import("./dispatchRule").DispatchRule = null as any;
utilities.lazyLoad(exports, ["DispatchRule"], () => require("./dispatchRule"));

export { GetAlertContactGroupsArgs, GetAlertContactGroupsResult, GetAlertContactGroupsOutputArgs } from "./getAlertContactGroups";
export const getAlertContactGroups: typeof import("./getAlertContactGroups").getAlertContactGroups = null as any;
export const getAlertContactGroupsOutput: typeof import("./getAlertContactGroups").getAlertContactGroupsOutput = null as any;
utilities.lazyLoad(exports, ["getAlertContactGroups","getAlertContactGroupsOutput"], () => require("./getAlertContactGroups"));

export { GetAlertContactsArgs, GetAlertContactsResult, GetAlertContactsOutputArgs } from "./getAlertContacts";
export const getAlertContacts: typeof import("./getAlertContacts").getAlertContacts = null as any;
export const getAlertContactsOutput: typeof import("./getAlertContacts").getAlertContactsOutput = null as any;
utilities.lazyLoad(exports, ["getAlertContacts","getAlertContactsOutput"], () => require("./getAlertContacts"));

export { GetDispatchRulesArgs, GetDispatchRulesResult, GetDispatchRulesOutputArgs } from "./getDispatchRules";
export const getDispatchRules: typeof import("./getDispatchRules").getDispatchRules = null as any;
export const getDispatchRulesOutput: typeof import("./getDispatchRules").getDispatchRulesOutput = null as any;
utilities.lazyLoad(exports, ["getDispatchRules","getDispatchRulesOutput"], () => require("./getDispatchRules"));

export { GetIntegrationExportersArgs, GetIntegrationExportersResult, GetIntegrationExportersOutputArgs } from "./getIntegrationExporters";
export const getIntegrationExporters: typeof import("./getIntegrationExporters").getIntegrationExporters = null as any;
export const getIntegrationExportersOutput: typeof import("./getIntegrationExporters").getIntegrationExportersOutput = null as any;
utilities.lazyLoad(exports, ["getIntegrationExporters","getIntegrationExportersOutput"], () => require("./getIntegrationExporters"));

export { GetPrometheisArgs, GetPrometheisResult, GetPrometheisOutputArgs } from "./getPrometheis";
export const getPrometheis: typeof import("./getPrometheis").getPrometheis = null as any;
export const getPrometheisOutput: typeof import("./getPrometheis").getPrometheisOutput = null as any;
utilities.lazyLoad(exports, ["getPrometheis","getPrometheisOutput"], () => require("./getPrometheis"));

export { GetPrometheusAlertRulesArgs, GetPrometheusAlertRulesResult, GetPrometheusAlertRulesOutputArgs } from "./getPrometheusAlertRules";
export const getPrometheusAlertRules: typeof import("./getPrometheusAlertRules").getPrometheusAlertRules = null as any;
export const getPrometheusAlertRulesOutput: typeof import("./getPrometheusAlertRules").getPrometheusAlertRulesOutput = null as any;
utilities.lazyLoad(exports, ["getPrometheusAlertRules","getPrometheusAlertRulesOutput"], () => require("./getPrometheusAlertRules"));

export { GetRemoteWritesArgs, GetRemoteWritesResult, GetRemoteWritesOutputArgs } from "./getRemoteWrites";
export const getRemoteWrites: typeof import("./getRemoteWrites").getRemoteWrites = null as any;
export const getRemoteWritesOutput: typeof import("./getRemoteWrites").getRemoteWritesOutput = null as any;
utilities.lazyLoad(exports, ["getRemoteWrites","getRemoteWritesOutput"], () => require("./getRemoteWrites"));

export { IntegrationExporterArgs, IntegrationExporterState } from "./integrationExporter";
export type IntegrationExporter = import("./integrationExporter").IntegrationExporter;
export const IntegrationExporter: typeof import("./integrationExporter").IntegrationExporter = null as any;
utilities.lazyLoad(exports, ["IntegrationExporter"], () => require("./integrationExporter"));

export { PrometheusArgs, PrometheusState } from "./prometheus";
export type Prometheus = import("./prometheus").Prometheus;
export const Prometheus: typeof import("./prometheus").Prometheus = null as any;
utilities.lazyLoad(exports, ["Prometheus"], () => require("./prometheus"));

export { PrometheusAlertRuleArgs, PrometheusAlertRuleState } from "./prometheusAlertRule";
export type PrometheusAlertRule = import("./prometheusAlertRule").PrometheusAlertRule;
export const PrometheusAlertRule: typeof import("./prometheusAlertRule").PrometheusAlertRule = null as any;
utilities.lazyLoad(exports, ["PrometheusAlertRule"], () => require("./prometheusAlertRule"));

export { RemoteWriteArgs, RemoteWriteState } from "./remoteWrite";
export type RemoteWrite = import("./remoteWrite").RemoteWrite;
export const RemoteWrite: typeof import("./remoteWrite").RemoteWrite = null as any;
utilities.lazyLoad(exports, ["RemoteWrite"], () => require("./remoteWrite"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:arms/alertContact:AlertContact":
                return new AlertContact(name, <any>undefined, { urn })
            case "alicloud:arms/alertContactGroup:AlertContactGroup":
                return new AlertContactGroup(name, <any>undefined, { urn })
            case "alicloud:arms/dispatchRule:DispatchRule":
                return new DispatchRule(name, <any>undefined, { urn })
            case "alicloud:arms/integrationExporter:IntegrationExporter":
                return new IntegrationExporter(name, <any>undefined, { urn })
            case "alicloud:arms/prometheus:Prometheus":
                return new Prometheus(name, <any>undefined, { urn })
            case "alicloud:arms/prometheusAlertRule:PrometheusAlertRule":
                return new PrometheusAlertRule(name, <any>undefined, { urn })
            case "alicloud:arms/remoteWrite:RemoteWrite":
                return new RemoteWrite(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "arms/alertContact", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/alertContactGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/dispatchRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/integrationExporter", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/prometheus", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/prometheusAlertRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/remoteWrite", _module)
