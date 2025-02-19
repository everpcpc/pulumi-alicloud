// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.inputs;

import com.pulumi.alicloud.arms.inputs.DispatchRuleNotifyRuleNotifyObjectArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class DispatchRuleNotifyRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final DispatchRuleNotifyRuleArgs Empty = new DispatchRuleNotifyRuleArgs();

    /**
     * The notification method. Valid values: dingTalk, sms, webhook, email, and wechat.
     * 
     */
    @Import(name="notifyChannels", required=true)
    private Output<List<String>> notifyChannels;

    /**
     * @return The notification method. Valid values: dingTalk, sms, webhook, email, and wechat.
     * 
     */
    public Output<List<String>> notifyChannels() {
        return this.notifyChannels;
    }

    /**
     * Sets the notification object. See the following `Block notify_objects`.
     * 
     */
    @Import(name="notifyObjects", required=true)
    private Output<List<DispatchRuleNotifyRuleNotifyObjectArgs>> notifyObjects;

    /**
     * @return Sets the notification object. See the following `Block notify_objects`.
     * 
     */
    public Output<List<DispatchRuleNotifyRuleNotifyObjectArgs>> notifyObjects() {
        return this.notifyObjects;
    }

    private DispatchRuleNotifyRuleArgs() {}

    private DispatchRuleNotifyRuleArgs(DispatchRuleNotifyRuleArgs $) {
        this.notifyChannels = $.notifyChannels;
        this.notifyObjects = $.notifyObjects;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DispatchRuleNotifyRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DispatchRuleNotifyRuleArgs $;

        public Builder() {
            $ = new DispatchRuleNotifyRuleArgs();
        }

        public Builder(DispatchRuleNotifyRuleArgs defaults) {
            $ = new DispatchRuleNotifyRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param notifyChannels The notification method. Valid values: dingTalk, sms, webhook, email, and wechat.
         * 
         * @return builder
         * 
         */
        public Builder notifyChannels(Output<List<String>> notifyChannels) {
            $.notifyChannels = notifyChannels;
            return this;
        }

        /**
         * @param notifyChannels The notification method. Valid values: dingTalk, sms, webhook, email, and wechat.
         * 
         * @return builder
         * 
         */
        public Builder notifyChannels(List<String> notifyChannels) {
            return notifyChannels(Output.of(notifyChannels));
        }

        /**
         * @param notifyChannels The notification method. Valid values: dingTalk, sms, webhook, email, and wechat.
         * 
         * @return builder
         * 
         */
        public Builder notifyChannels(String... notifyChannels) {
            return notifyChannels(List.of(notifyChannels));
        }

        /**
         * @param notifyObjects Sets the notification object. See the following `Block notify_objects`.
         * 
         * @return builder
         * 
         */
        public Builder notifyObjects(Output<List<DispatchRuleNotifyRuleNotifyObjectArgs>> notifyObjects) {
            $.notifyObjects = notifyObjects;
            return this;
        }

        /**
         * @param notifyObjects Sets the notification object. See the following `Block notify_objects`.
         * 
         * @return builder
         * 
         */
        public Builder notifyObjects(List<DispatchRuleNotifyRuleNotifyObjectArgs> notifyObjects) {
            return notifyObjects(Output.of(notifyObjects));
        }

        /**
         * @param notifyObjects Sets the notification object. See the following `Block notify_objects`.
         * 
         * @return builder
         * 
         */
        public Builder notifyObjects(DispatchRuleNotifyRuleNotifyObjectArgs... notifyObjects) {
            return notifyObjects(List.of(notifyObjects));
        }

        public DispatchRuleNotifyRuleArgs build() {
            $.notifyChannels = Objects.requireNonNull($.notifyChannels, "expected parameter 'notifyChannels' to be non-null");
            $.notifyObjects = Objects.requireNonNull($.notifyObjects, "expected parameter 'notifyObjects' to be non-null");
            return $;
        }
    }

}
