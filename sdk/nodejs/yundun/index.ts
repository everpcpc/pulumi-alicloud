// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./bastionHostInstance";
export * from "./dbauditInstance";
export * from "./getBastionHostInstances";
export * from "./getDBAuditInstance";

// Import resources to register:
import { BastionHostInstance } from "./bastionHostInstance";
import { DBAuditInstance } from "./dbauditInstance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:yundun/bastionHostInstance:BastionHostInstance":
                return new BastionHostInstance(name, <any>undefined, { urn })
            case "alicloud:yundun/dBAuditInstance:DBAuditInstance":
                return new DBAuditInstance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "yundun/bastionHostInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "yundun/dBAuditInstance", _module)
