// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExecutionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExecutionArgs Empty = new ExecutionArgs();

    /**
     * The description of OOS Execution.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of OOS Execution.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The loop mode of OOS Execution.
     * 
     */
    @Import(name="loopMode")
    private @Nullable Output<String> loopMode;

    /**
     * @return The loop mode of OOS Execution.
     * 
     */
    public Optional<Output<String>> loopMode() {
        return Optional.ofNullable(this.loopMode);
    }

    /**
     * The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The parameters required by the template. Default to `{}`.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<String> parameters;

    /**
     * @return The parameters required by the template. Default to `{}`.
     * 
     */
    public Optional<Output<String>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * The id of parent execution.
     * 
     */
    @Import(name="parentExecutionId")
    private @Nullable Output<String> parentExecutionId;

    /**
     * @return The id of parent execution.
     * 
     */
    public Optional<Output<String>> parentExecutionId() {
        return Optional.ofNullable(this.parentExecutionId);
    }

    /**
     * The mode of safety check.
     * 
     */
    @Import(name="safetyCheck")
    private @Nullable Output<String> safetyCheck;

    /**
     * @return The mode of safety check.
     * 
     */
    public Optional<Output<String>> safetyCheck() {
        return Optional.ofNullable(this.safetyCheck);
    }

    /**
     * The content of template. When the user selects an existing template to create and execute a task, it is not necessary to pass in this field.
     * 
     */
    @Import(name="templateContent")
    private @Nullable Output<String> templateContent;

    /**
     * @return The content of template. When the user selects an existing template to create and execute a task, it is not necessary to pass in this field.
     * 
     */
    public Optional<Output<String>> templateContent() {
        return Optional.ofNullable(this.templateContent);
    }

    /**
     * The name of execution template.
     * 
     */
    @Import(name="templateName", required=true)
    private Output<String> templateName;

    /**
     * @return The name of execution template.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    /**
     * The version of execution template.
     * 
     */
    @Import(name="templateVersion")
    private @Nullable Output<String> templateVersion;

    /**
     * @return The version of execution template.
     * 
     */
    public Optional<Output<String>> templateVersion() {
        return Optional.ofNullable(this.templateVersion);
    }

    private ExecutionArgs() {}

    private ExecutionArgs(ExecutionArgs $) {
        this.description = $.description;
        this.loopMode = $.loopMode;
        this.mode = $.mode;
        this.parameters = $.parameters;
        this.parentExecutionId = $.parentExecutionId;
        this.safetyCheck = $.safetyCheck;
        this.templateContent = $.templateContent;
        this.templateName = $.templateName;
        this.templateVersion = $.templateVersion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExecutionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExecutionArgs $;

        public Builder() {
            $ = new ExecutionArgs();
        }

        public Builder(ExecutionArgs defaults) {
            $ = new ExecutionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of OOS Execution.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of OOS Execution.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param loopMode The loop mode of OOS Execution.
         * 
         * @return builder
         * 
         */
        public Builder loopMode(@Nullable Output<String> loopMode) {
            $.loopMode = loopMode;
            return this;
        }

        /**
         * @param loopMode The loop mode of OOS Execution.
         * 
         * @return builder
         * 
         */
        public Builder loopMode(String loopMode) {
            return loopMode(Output.of(loopMode));
        }

        /**
         * @param mode The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param parameters The parameters required by the template. Default to `{}`.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<String> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters The parameters required by the template. Default to `{}`.
         * 
         * @return builder
         * 
         */
        public Builder parameters(String parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param parentExecutionId The id of parent execution.
         * 
         * @return builder
         * 
         */
        public Builder parentExecutionId(@Nullable Output<String> parentExecutionId) {
            $.parentExecutionId = parentExecutionId;
            return this;
        }

        /**
         * @param parentExecutionId The id of parent execution.
         * 
         * @return builder
         * 
         */
        public Builder parentExecutionId(String parentExecutionId) {
            return parentExecutionId(Output.of(parentExecutionId));
        }

        /**
         * @param safetyCheck The mode of safety check.
         * 
         * @return builder
         * 
         */
        public Builder safetyCheck(@Nullable Output<String> safetyCheck) {
            $.safetyCheck = safetyCheck;
            return this;
        }

        /**
         * @param safetyCheck The mode of safety check.
         * 
         * @return builder
         * 
         */
        public Builder safetyCheck(String safetyCheck) {
            return safetyCheck(Output.of(safetyCheck));
        }

        /**
         * @param templateContent The content of template. When the user selects an existing template to create and execute a task, it is not necessary to pass in this field.
         * 
         * @return builder
         * 
         */
        public Builder templateContent(@Nullable Output<String> templateContent) {
            $.templateContent = templateContent;
            return this;
        }

        /**
         * @param templateContent The content of template. When the user selects an existing template to create and execute a task, it is not necessary to pass in this field.
         * 
         * @return builder
         * 
         */
        public Builder templateContent(String templateContent) {
            return templateContent(Output.of(templateContent));
        }

        /**
         * @param templateName The name of execution template.
         * 
         * @return builder
         * 
         */
        public Builder templateName(Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName The name of execution template.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        /**
         * @param templateVersion The version of execution template.
         * 
         * @return builder
         * 
         */
        public Builder templateVersion(@Nullable Output<String> templateVersion) {
            $.templateVersion = templateVersion;
            return this;
        }

        /**
         * @param templateVersion The version of execution template.
         * 
         * @return builder
         * 
         */
        public Builder templateVersion(String templateVersion) {
            return templateVersion(Output.of(templateVersion));
        }

        public ExecutionArgs build() {
            $.templateName = Objects.requireNonNull($.templateName, "expected parameter 'templateName' to be non-null");
            return $;
        }
    }

}
