// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudconnect;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class NetworkAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkAttachmentArgs Empty = new NetworkAttachmentArgs();

    /**
     * The ID of the CCN instance.
     * 
     */
    @Import(name="ccnId", required=true)
    private Output<String> ccnId;

    /**
     * @return The ID of the CCN instance.
     * 
     */
    public Output<String> ccnId() {
        return this.ccnId;
    }

    /**
     * The ID of the Smart Access Gateway instance.
     * 
     */
    @Import(name="sagId", required=true)
    private Output<String> sagId;

    /**
     * @return The ID of the Smart Access Gateway instance.
     * 
     */
    public Output<String> sagId() {
        return this.sagId;
    }

    private NetworkAttachmentArgs() {}

    private NetworkAttachmentArgs(NetworkAttachmentArgs $) {
        this.ccnId = $.ccnId;
        this.sagId = $.sagId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkAttachmentArgs $;

        public Builder() {
            $ = new NetworkAttachmentArgs();
        }

        public Builder(NetworkAttachmentArgs defaults) {
            $ = new NetworkAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ccnId The ID of the CCN instance.
         * 
         * @return builder
         * 
         */
        public Builder ccnId(Output<String> ccnId) {
            $.ccnId = ccnId;
            return this;
        }

        /**
         * @param ccnId The ID of the CCN instance.
         * 
         * @return builder
         * 
         */
        public Builder ccnId(String ccnId) {
            return ccnId(Output.of(ccnId));
        }

        /**
         * @param sagId The ID of the Smart Access Gateway instance.
         * 
         * @return builder
         * 
         */
        public Builder sagId(Output<String> sagId) {
            $.sagId = sagId;
            return this;
        }

        /**
         * @param sagId The ID of the Smart Access Gateway instance.
         * 
         * @return builder
         * 
         */
        public Builder sagId(String sagId) {
            return sagId(Output.of(sagId));
        }

        public NetworkAttachmentArgs build() {
            $.ccnId = Objects.requireNonNull($.ccnId, "expected parameter 'ccnId' to be non-null");
            $.sagId = Objects.requireNonNull($.sagId, "expected parameter 'sagId' to be non-null");
            return $;
        }
    }

}
