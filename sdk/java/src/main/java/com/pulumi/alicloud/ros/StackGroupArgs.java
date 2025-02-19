// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros;

import com.pulumi.alicloud.ros.inputs.StackGroupParameterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StackGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final StackGroupArgs Empty = new StackGroupArgs();

    /**
     * The list of target account IDs, in JSON format. A maximum of 20 accounts can be specified.
     * 
     */
    @Import(name="accountIds")
    private @Nullable Output<String> accountIds;

    /**
     * @return The list of target account IDs, in JSON format. A maximum of 20 accounts can be specified.
     * 
     */
    public Optional<Output<String>> accountIds() {
        return Optional.ofNullable(this.accountIds);
    }

    /**
     * The name of the RAM administrator role assumed by ROS. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
     * 
     */
    @Import(name="administrationRoleName")
    private @Nullable Output<String> administrationRoleName;

    /**
     * @return The name of the RAM administrator role assumed by ROS. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
     * 
     */
    public Optional<Output<String>> administrationRoleName() {
        return Optional.ofNullable(this.administrationRoleName);
    }

    /**
     * The description of the stack group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the stack group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the RAM execution role assumed by the administrator role. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
     * 
     */
    @Import(name="executionRoleName")
    private @Nullable Output<String> executionRoleName;

    /**
     * @return The name of the RAM execution role assumed by the administrator role. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
     * 
     */
    public Optional<Output<String>> executionRoleName() {
        return Optional.ofNullable(this.executionRoleName);
    }

    /**
     * The description of the operation.
     * 
     */
    @Import(name="operationDescription")
    private @Nullable Output<String> operationDescription;

    /**
     * @return The description of the operation.
     * 
     */
    public Optional<Output<String>> operationDescription() {
        return Optional.ofNullable(this.operationDescription);
    }

    /**
     * The operation settings, in JSON format.
     * 
     */
    @Import(name="operationPreferences")
    private @Nullable Output<String> operationPreferences;

    /**
     * @return The operation settings, in JSON format.
     * 
     */
    public Optional<Output<String>> operationPreferences() {
        return Optional.ofNullable(this.operationPreferences);
    }

    /**
     * The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<List<StackGroupParameterArgs>> parameters;

    /**
     * @return The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
     * 
     */
    public Optional<Output<List<StackGroupParameterArgs>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * The list of target regions, in JSON format. A maximum of 20 accounts can be specified.
     * 
     */
    @Import(name="regionIds")
    private @Nullable Output<String> regionIds;

    /**
     * @return The list of target regions, in JSON format. A maximum of 20 accounts can be specified.
     * 
     */
    public Optional<Output<String>> regionIds() {
        return Optional.ofNullable(this.regionIds);
    }

    /**
     * The name of the stack group. The name must be unique in a region.
     * 
     */
    @Import(name="stackGroupName", required=true)
    private Output<String> stackGroupName;

    /**
     * @return The name of the stack group. The name must be unique in a region.
     * 
     */
    public Output<String> stackGroupName() {
        return this.stackGroupName;
    }

    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
     * 
     */
    @Import(name="templateBody")
    private @Nullable Output<String> templateBody;

    /**
     * @return The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
     * 
     */
    public Optional<Output<String>> templateBody() {
        return Optional.ofNullable(this.templateBody);
    }

    /**
     * The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    @Import(name="templateUrl")
    private @Nullable Output<String> templateUrl;

    /**
     * @return The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
     * 
     */
    public Optional<Output<String>> templateUrl() {
        return Optional.ofNullable(this.templateUrl);
    }

    /**
     * The version of the template.
     * 
     */
    @Import(name="templateVersion")
    private @Nullable Output<String> templateVersion;

    /**
     * @return The version of the template.
     * 
     */
    public Optional<Output<String>> templateVersion() {
        return Optional.ofNullable(this.templateVersion);
    }

    private StackGroupArgs() {}

    private StackGroupArgs(StackGroupArgs $) {
        this.accountIds = $.accountIds;
        this.administrationRoleName = $.administrationRoleName;
        this.description = $.description;
        this.executionRoleName = $.executionRoleName;
        this.operationDescription = $.operationDescription;
        this.operationPreferences = $.operationPreferences;
        this.parameters = $.parameters;
        this.regionIds = $.regionIds;
        this.stackGroupName = $.stackGroupName;
        this.templateBody = $.templateBody;
        this.templateUrl = $.templateUrl;
        this.templateVersion = $.templateVersion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StackGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StackGroupArgs $;

        public Builder() {
            $ = new StackGroupArgs();
        }

        public Builder(StackGroupArgs defaults) {
            $ = new StackGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountIds The list of target account IDs, in JSON format. A maximum of 20 accounts can be specified.
         * 
         * @return builder
         * 
         */
        public Builder accountIds(@Nullable Output<String> accountIds) {
            $.accountIds = accountIds;
            return this;
        }

        /**
         * @param accountIds The list of target account IDs, in JSON format. A maximum of 20 accounts can be specified.
         * 
         * @return builder
         * 
         */
        public Builder accountIds(String accountIds) {
            return accountIds(Output.of(accountIds));
        }

        /**
         * @param administrationRoleName The name of the RAM administrator role assumed by ROS. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
         * 
         * @return builder
         * 
         */
        public Builder administrationRoleName(@Nullable Output<String> administrationRoleName) {
            $.administrationRoleName = administrationRoleName;
            return this;
        }

        /**
         * @param administrationRoleName The name of the RAM administrator role assumed by ROS. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
         * 
         * @return builder
         * 
         */
        public Builder administrationRoleName(String administrationRoleName) {
            return administrationRoleName(Output.of(administrationRoleName));
        }

        /**
         * @param description The description of the stack group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the stack group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param executionRoleName The name of the RAM execution role assumed by the administrator role. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
         * 
         * @return builder
         * 
         */
        public Builder executionRoleName(@Nullable Output<String> executionRoleName) {
            $.executionRoleName = executionRoleName;
            return this;
        }

        /**
         * @param executionRoleName The name of the RAM execution role assumed by the administrator role. ROS assumes this role to perform operations on the stack corresponding to the stack instance in the stack group.
         * 
         * @return builder
         * 
         */
        public Builder executionRoleName(String executionRoleName) {
            return executionRoleName(Output.of(executionRoleName));
        }

        /**
         * @param operationDescription The description of the operation.
         * 
         * @return builder
         * 
         */
        public Builder operationDescription(@Nullable Output<String> operationDescription) {
            $.operationDescription = operationDescription;
            return this;
        }

        /**
         * @param operationDescription The description of the operation.
         * 
         * @return builder
         * 
         */
        public Builder operationDescription(String operationDescription) {
            return operationDescription(Output.of(operationDescription));
        }

        /**
         * @param operationPreferences The operation settings, in JSON format.
         * 
         * @return builder
         * 
         */
        public Builder operationPreferences(@Nullable Output<String> operationPreferences) {
            $.operationPreferences = operationPreferences;
            return this;
        }

        /**
         * @param operationPreferences The operation settings, in JSON format.
         * 
         * @return builder
         * 
         */
        public Builder operationPreferences(String operationPreferences) {
            return operationPreferences(Output.of(operationPreferences));
        }

        /**
         * @param parameters The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<List<StackGroupParameterArgs>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
         * 
         * @return builder
         * 
         */
        public Builder parameters(List<StackGroupParameterArgs> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param parameters The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
         * 
         * @return builder
         * 
         */
        public Builder parameters(StackGroupParameterArgs... parameters) {
            return parameters(List.of(parameters));
        }

        /**
         * @param regionIds The list of target regions, in JSON format. A maximum of 20 accounts can be specified.
         * 
         * @return builder
         * 
         */
        public Builder regionIds(@Nullable Output<String> regionIds) {
            $.regionIds = regionIds;
            return this;
        }

        /**
         * @param regionIds The list of target regions, in JSON format. A maximum of 20 accounts can be specified.
         * 
         * @return builder
         * 
         */
        public Builder regionIds(String regionIds) {
            return regionIds(Output.of(regionIds));
        }

        /**
         * @param stackGroupName The name of the stack group. The name must be unique in a region.
         * 
         * @return builder
         * 
         */
        public Builder stackGroupName(Output<String> stackGroupName) {
            $.stackGroupName = stackGroupName;
            return this;
        }

        /**
         * @param stackGroupName The name of the stack group. The name must be unique in a region.
         * 
         * @return builder
         * 
         */
        public Builder stackGroupName(String stackGroupName) {
            return stackGroupName(Output.of(stackGroupName));
        }

        /**
         * @param templateBody The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
         * 
         * @return builder
         * 
         */
        public Builder templateBody(@Nullable Output<String> templateBody) {
            $.templateBody = templateBody;
            return this;
        }

        /**
         * @param templateBody The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
         * 
         * @return builder
         * 
         */
        public Builder templateBody(String templateBody) {
            return templateBody(Output.of(templateBody));
        }

        /**
         * @param templateUrl The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder templateUrl(@Nullable Output<String> templateUrl) {
            $.templateUrl = templateUrl;
            return this;
        }

        /**
         * @param templateUrl The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
         * 
         * @return builder
         * 
         */
        public Builder templateUrl(String templateUrl) {
            return templateUrl(Output.of(templateUrl));
        }

        /**
         * @param templateVersion The version of the template.
         * 
         * @return builder
         * 
         */
        public Builder templateVersion(@Nullable Output<String> templateVersion) {
            $.templateVersion = templateVersion;
            return this;
        }

        /**
         * @param templateVersion The version of the template.
         * 
         * @return builder
         * 
         */
        public Builder templateVersion(String templateVersion) {
            return templateVersion(Output.of(templateVersion));
        }

        public StackGroupArgs build() {
            $.stackGroupName = Objects.requireNonNull($.stackGroupName, "expected parameter 'stackGroupName' to be non-null");
            return $;
        }
    }

}
