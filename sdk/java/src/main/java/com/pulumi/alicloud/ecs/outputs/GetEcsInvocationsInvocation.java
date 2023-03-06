// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetEcsInvocationsInvocationInvokeInstance;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetEcsInvocationsInvocation {
    /**
     * @return The Base64-encoded command content.
     * 
     */
    private String commandContent;
    /**
     * @return The ID of the command.
     * 
     */
    private String commandId;
    /**
     * @return The name of the command.
     * 
     */
    private String commandName;
    /**
     * @return The type of the command.
     * 
     */
    private String commandType;
    /**
     * @return The creation time of the resource.
     * 
     */
    private String createTime;
    /**
     * @return The schedule on which the recurring execution of the command takes place. For information about the value specifications, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
     * 
     */
    private String frequency;
    /**
     * @return The ID of the Invocation.
     * 
     */
    private String id;
    /**
     * @return The ID of the Invocation.
     * 
     */
    private String invocationId;
    /**
     * @return The execution state on a single instance. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopping`, `Stopped`, `PartialFailed`.
     * 
     */
    private String invocationStatus;
    /**
     * @return Execute target instance set type.
     * 
     */
    private List<GetEcsInvocationsInvocationInvokeInstance> invokeInstances;
    /**
     * @return The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
     * 
     */
    private String invokeStatus;
    /**
     * @return The custom parameters in the command.
     * 
     */
    private String parameters;
    /**
     * @return Indicates the execution mode of the command.
     * 
     */
    private String repeatMode;
    /**
     * @return Indicates whether the commands are to be automatically run.
     * * `error_code	` - The code that indicates why the command failed to be sent or run.
     * * `instance_invoke_status	` - **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
     * 
     */
    private Boolean timed;
    /**
     * @return The username that was used to run the command on the instance.
     * 
     */
    private String username;

    private GetEcsInvocationsInvocation() {}
    /**
     * @return The Base64-encoded command content.
     * 
     */
    public String commandContent() {
        return this.commandContent;
    }
    /**
     * @return The ID of the command.
     * 
     */
    public String commandId() {
        return this.commandId;
    }
    /**
     * @return The name of the command.
     * 
     */
    public String commandName() {
        return this.commandName;
    }
    /**
     * @return The type of the command.
     * 
     */
    public String commandType() {
        return this.commandType;
    }
    /**
     * @return The creation time of the resource.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The schedule on which the recurring execution of the command takes place. For information about the value specifications, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
     * 
     */
    public String frequency() {
        return this.frequency;
    }
    /**
     * @return The ID of the Invocation.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the Invocation.
     * 
     */
    public String invocationId() {
        return this.invocationId;
    }
    /**
     * @return The execution state on a single instance. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopping`, `Stopped`, `PartialFailed`.
     * 
     */
    public String invocationStatus() {
        return this.invocationStatus;
    }
    /**
     * @return Execute target instance set type.
     * 
     */
    public List<GetEcsInvocationsInvocationInvokeInstance> invokeInstances() {
        return this.invokeInstances;
    }
    /**
     * @return The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
     * 
     */
    public String invokeStatus() {
        return this.invokeStatus;
    }
    /**
     * @return The custom parameters in the command.
     * 
     */
    public String parameters() {
        return this.parameters;
    }
    /**
     * @return Indicates the execution mode of the command.
     * 
     */
    public String repeatMode() {
        return this.repeatMode;
    }
    /**
     * @return Indicates whether the commands are to be automatically run.
     * * `error_code	` - The code that indicates why the command failed to be sent or run.
     * * `instance_invoke_status	` - **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
     * 
     */
    public Boolean timed() {
        return this.timed;
    }
    /**
     * @return The username that was used to run the command on the instance.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEcsInvocationsInvocation defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String commandContent;
        private String commandId;
        private String commandName;
        private String commandType;
        private String createTime;
        private String frequency;
        private String id;
        private String invocationId;
        private String invocationStatus;
        private List<GetEcsInvocationsInvocationInvokeInstance> invokeInstances;
        private String invokeStatus;
        private String parameters;
        private String repeatMode;
        private Boolean timed;
        private String username;
        public Builder() {}
        public Builder(GetEcsInvocationsInvocation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commandContent = defaults.commandContent;
    	      this.commandId = defaults.commandId;
    	      this.commandName = defaults.commandName;
    	      this.commandType = defaults.commandType;
    	      this.createTime = defaults.createTime;
    	      this.frequency = defaults.frequency;
    	      this.id = defaults.id;
    	      this.invocationId = defaults.invocationId;
    	      this.invocationStatus = defaults.invocationStatus;
    	      this.invokeInstances = defaults.invokeInstances;
    	      this.invokeStatus = defaults.invokeStatus;
    	      this.parameters = defaults.parameters;
    	      this.repeatMode = defaults.repeatMode;
    	      this.timed = defaults.timed;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder commandContent(String commandContent) {
            this.commandContent = Objects.requireNonNull(commandContent);
            return this;
        }
        @CustomType.Setter
        public Builder commandId(String commandId) {
            this.commandId = Objects.requireNonNull(commandId);
            return this;
        }
        @CustomType.Setter
        public Builder commandName(String commandName) {
            this.commandName = Objects.requireNonNull(commandName);
            return this;
        }
        @CustomType.Setter
        public Builder commandType(String commandType) {
            this.commandType = Objects.requireNonNull(commandType);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder frequency(String frequency) {
            this.frequency = Objects.requireNonNull(frequency);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder invocationId(String invocationId) {
            this.invocationId = Objects.requireNonNull(invocationId);
            return this;
        }
        @CustomType.Setter
        public Builder invocationStatus(String invocationStatus) {
            this.invocationStatus = Objects.requireNonNull(invocationStatus);
            return this;
        }
        @CustomType.Setter
        public Builder invokeInstances(List<GetEcsInvocationsInvocationInvokeInstance> invokeInstances) {
            this.invokeInstances = Objects.requireNonNull(invokeInstances);
            return this;
        }
        public Builder invokeInstances(GetEcsInvocationsInvocationInvokeInstance... invokeInstances) {
            return invokeInstances(List.of(invokeInstances));
        }
        @CustomType.Setter
        public Builder invokeStatus(String invokeStatus) {
            this.invokeStatus = Objects.requireNonNull(invokeStatus);
            return this;
        }
        @CustomType.Setter
        public Builder parameters(String parameters) {
            this.parameters = Objects.requireNonNull(parameters);
            return this;
        }
        @CustomType.Setter
        public Builder repeatMode(String repeatMode) {
            this.repeatMode = Objects.requireNonNull(repeatMode);
            return this;
        }
        @CustomType.Setter
        public Builder timed(Boolean timed) {
            this.timed = Objects.requireNonNull(timed);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public GetEcsInvocationsInvocation build() {
            final var o = new GetEcsInvocationsInvocation();
            o.commandContent = commandContent;
            o.commandId = commandId;
            o.commandName = commandName;
            o.commandType = commandType;
            o.createTime = createTime;
            o.frequency = frequency;
            o.id = id;
            o.invocationId = invocationId;
            o.invocationStatus = invocationStatus;
            o.invokeInstances = invokeInstances;
            o.invokeStatus = invokeStatus;
            o.parameters = parameters;
            o.repeatMode = repeatMode;
            o.timed = timed;
            o.username = username;
            return o;
        }
    }
}