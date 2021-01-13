// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accelerator";
export * from "./bandwidthPackage";
export * from "./endpointGroup";
export * from "./getAccelerators";
export * from "./getBandwidthPackages";
export * from "./getEndpointGroups";
export * from "./getListeners";
export * from "./listener";

// Import resources to register:
import { Accelerator } from "./accelerator";
import { BandwidthPackage } from "./bandwidthPackage";
import { EndpointGroup } from "./endpointGroup";
import { Listener } from "./listener";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:ga/accelerator:Accelerator":
                return new Accelerator(name, <any>undefined, { urn })
            case "alicloud:ga/bandwidthPackage:BandwidthPackage":
                return new BandwidthPackage(name, <any>undefined, { urn })
            case "alicloud:ga/endpointGroup:EndpointGroup":
                return new EndpointGroup(name, <any>undefined, { urn })
            case "alicloud:ga/listener:Listener":
                return new Listener(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "ga/accelerator", _module)
pulumi.runtime.registerResourceModule("alicloud", "ga/bandwidthPackage", _module)
pulumi.runtime.registerResourceModule("alicloud", "ga/endpointGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "ga/listener", _module)
