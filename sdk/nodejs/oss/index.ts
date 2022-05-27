// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./bucket";
export * from "./bucketObject";
export * from "./bucketReplication";
export * from "./getBucketObjects";
export * from "./getBuckets";
export * from "./getInstanceAttachments";
export * from "./getInstances";
export * from "./getService";
export * from "./getTables";

// Import resources to register:
import { Bucket } from "./bucket";
import { BucketObject } from "./bucketObject";
import { BucketReplication } from "./bucketReplication";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:oss/bucket:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "alicloud:oss/bucketObject:BucketObject":
                return new BucketObject(name, <any>undefined, { urn })
            case "alicloud:oss/bucketReplication:BucketReplication":
                return new BucketReplication(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "oss/bucket", _module)
pulumi.runtime.registerResourceModule("alicloud", "oss/bucketObject", _module)
pulumi.runtime.registerResourceModule("alicloud", "oss/bucketReplication", _module)
