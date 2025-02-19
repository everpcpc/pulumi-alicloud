// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.amqp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBindingsBinding {
    /**
     * @return X-match Attributes. Valid Values: All: Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match. Any: at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match. This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     * 
     */
    private String argument;
    /**
     * @return The Binding Key. The Source of the Binding Exchange Non-Topic Type: Can Only Contain Letters, Lowercase Letters, Numbers, and the Dash (-), the Underscore Character (_), English Periods (.) and the at Sign (@). Length from 1 to 255 Characters. The Source of the Binding Exchange Topic Type: Can Contain Letters, Lowercase Letters, Numbers, and the Dash (-), the Underscore Character (_), English Periods (.) and the at Sign (@). If You Include the Hash (.
     * 
     */
    private String bindingKey;
    /**
     * @return The Target Binding Types.
     * 
     */
    private String bindingType;
    /**
     * @return The Target Queue Or Exchange of the Name.
     * 
     */
    private String destinationName;
    /**
     * @return The ID of the Binding. The value formats as `&lt;instance_id&gt;:&lt;virtual_host_name&gt;:&lt;source_exchange&gt;:&lt;destination_name&gt;`.
     * 
     */
    private String id;
    /**
     * @return Instance Id.
     * 
     */
    private String instanceId;
    /**
     * @return The Source Exchange Name.
     * 
     */
    private String sourceExchange;
    /**
     * @return Virtualhost Name.
     * 
     */
    private String virtualHostName;

    private GetBindingsBinding() {}
    /**
     * @return X-match Attributes. Valid Values: All: Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match. Any: at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match. This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     * 
     */
    public String argument() {
        return this.argument;
    }
    /**
     * @return The Binding Key. The Source of the Binding Exchange Non-Topic Type: Can Only Contain Letters, Lowercase Letters, Numbers, and the Dash (-), the Underscore Character (_), English Periods (.) and the at Sign (@). Length from 1 to 255 Characters. The Source of the Binding Exchange Topic Type: Can Contain Letters, Lowercase Letters, Numbers, and the Dash (-), the Underscore Character (_), English Periods (.) and the at Sign (@). If You Include the Hash (.
     * 
     */
    public String bindingKey() {
        return this.bindingKey;
    }
    /**
     * @return The Target Binding Types.
     * 
     */
    public String bindingType() {
        return this.bindingType;
    }
    /**
     * @return The Target Queue Or Exchange of the Name.
     * 
     */
    public String destinationName() {
        return this.destinationName;
    }
    /**
     * @return The ID of the Binding. The value formats as `&lt;instance_id&gt;:&lt;virtual_host_name&gt;:&lt;source_exchange&gt;:&lt;destination_name&gt;`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Instance Id.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The Source Exchange Name.
     * 
     */
    public String sourceExchange() {
        return this.sourceExchange;
    }
    /**
     * @return Virtualhost Name.
     * 
     */
    public String virtualHostName() {
        return this.virtualHostName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBindingsBinding defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String argument;
        private String bindingKey;
        private String bindingType;
        private String destinationName;
        private String id;
        private String instanceId;
        private String sourceExchange;
        private String virtualHostName;
        public Builder() {}
        public Builder(GetBindingsBinding defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.argument = defaults.argument;
    	      this.bindingKey = defaults.bindingKey;
    	      this.bindingType = defaults.bindingType;
    	      this.destinationName = defaults.destinationName;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.sourceExchange = defaults.sourceExchange;
    	      this.virtualHostName = defaults.virtualHostName;
        }

        @CustomType.Setter
        public Builder argument(String argument) {
            this.argument = Objects.requireNonNull(argument);
            return this;
        }
        @CustomType.Setter
        public Builder bindingKey(String bindingKey) {
            this.bindingKey = Objects.requireNonNull(bindingKey);
            return this;
        }
        @CustomType.Setter
        public Builder bindingType(String bindingType) {
            this.bindingType = Objects.requireNonNull(bindingType);
            return this;
        }
        @CustomType.Setter
        public Builder destinationName(String destinationName) {
            this.destinationName = Objects.requireNonNull(destinationName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceExchange(String sourceExchange) {
            this.sourceExchange = Objects.requireNonNull(sourceExchange);
            return this;
        }
        @CustomType.Setter
        public Builder virtualHostName(String virtualHostName) {
            this.virtualHostName = Objects.requireNonNull(virtualHostName);
            return this;
        }
        public GetBindingsBinding build() {
            final var o = new GetBindingsBinding();
            o.argument = argument;
            o.bindingKey = bindingKey;
            o.bindingType = bindingType;
            o.destinationName = destinationName;
            o.id = id;
            o.instanceId = instanceId;
            o.sourceExchange = sourceExchange;
            o.virtualHostName = virtualHostName;
            return o;
        }
    }
}
