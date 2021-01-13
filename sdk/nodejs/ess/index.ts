// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alarm";
export * from "./attachment";
export * from "./getAlarms";
export * from "./getLifecycleHooks";
export * from "./getNotifications";
export * from "./getScalingConfigurations";
export * from "./getScalingGroups";
export * from "./getScalingRules";
export * from "./getScheduledTasks";
export * from "./lifecycleHook";
export * from "./notification";
export * from "./scalingConfiguration";
export * from "./scalingGroup";
export * from "./scalingGroupVServerGroups";
export * from "./scalingRule";
export * from "./schedule";
export * from "./scheduledTask";

// Import resources to register:
import { Alarm } from "./alarm";
import { Attachment } from "./attachment";
import { LifecycleHook } from "./lifecycleHook";
import { Notification } from "./notification";
import { ScalingConfiguration } from "./scalingConfiguration";
import { ScalingGroup } from "./scalingGroup";
import { ScalingGroupVServerGroups } from "./scalingGroupVServerGroups";
import { ScalingRule } from "./scalingRule";
import { Schedule } from "./schedule";
import { ScheduledTask } from "./scheduledTask";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:ess/alarm:Alarm":
                return new Alarm(name, <any>undefined, { urn })
            case "alicloud:ess/attachment:Attachment":
                return new Attachment(name, <any>undefined, { urn })
            case "alicloud:ess/lifecycleHook:LifecycleHook":
                return new LifecycleHook(name, <any>undefined, { urn })
            case "alicloud:ess/notification:Notification":
                return new Notification(name, <any>undefined, { urn })
            case "alicloud:ess/scalingConfiguration:ScalingConfiguration":
                return new ScalingConfiguration(name, <any>undefined, { urn })
            case "alicloud:ess/scalingGroup:ScalingGroup":
                return new ScalingGroup(name, <any>undefined, { urn })
            case "alicloud:ess/scalingGroupVServerGroups:ScalingGroupVServerGroups":
                return new ScalingGroupVServerGroups(name, <any>undefined, { urn })
            case "alicloud:ess/scalingRule:ScalingRule":
                return new ScalingRule(name, <any>undefined, { urn })
            case "alicloud:ess/schedule:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            case "alicloud:ess/scheduledTask:ScheduledTask":
                return new ScheduledTask(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "ess/alarm", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/attachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/lifecycleHook", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/notification", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/scalingConfiguration", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/scalingGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/scalingGroupVServerGroups", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/scalingRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/schedule", _module)
pulumi.runtime.registerResourceModule("alicloud", "ess/scheduledTask", _module)
