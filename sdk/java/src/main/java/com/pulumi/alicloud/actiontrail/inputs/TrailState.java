// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TrailState extends com.pulumi.resources.ResourceArgs {

    public static final TrailState Empty = new TrailState();

    /**
     * Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     * 
     */
    @Import(name="eventRw")
    private @Nullable Output<String> eventRw;

    /**
     * @return Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     * 
     */
    public Optional<Output<String>> eventRw() {
        return Optional.ofNullable(this.eventRw);
    }

    /**
     * Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     * 
     */
    @Import(name="isOrganizationTrail")
    private @Nullable Output<Boolean> isOrganizationTrail;

    /**
     * @return Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     * 
     */
    public Optional<Output<Boolean>> isOrganizationTrail() {
        return Optional.ofNullable(this.isOrganizationTrail);
    }

    /**
     * Field `mns_topic_arn` has been deprecated from version 1.118.0.
     * 
     * @deprecated
     * Field &#39;mns_topic_arn&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'mns_topic_arn' has been deprecated from version 1.118.0 */
    @Import(name="mnsTopicArn")
    private @Nullable Output<String> mnsTopicArn;

    /**
     * @return Field `mns_topic_arn` has been deprecated from version 1.118.0.
     * 
     * @deprecated
     * Field &#39;mns_topic_arn&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'mns_topic_arn' has been deprecated from version 1.118.0 */
    public Optional<Output<String>> mnsTopicArn() {
        return Optional.ofNullable(this.mnsTopicArn);
    }

    /**
     * Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.95.0. Use &#39;trail_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.95.0. Use &#39;trail_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     * 
     */
    @Import(name="ossBucketName")
    private @Nullable Output<String> ossBucketName;

    /**
     * @return The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     * 
     */
    public Optional<Output<String>> ossBucketName() {
        return Optional.ofNullable(this.ossBucketName);
    }

    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     * 
     */
    @Import(name="ossKeyPrefix")
    private @Nullable Output<String> ossKeyPrefix;

    /**
     * @return The prefix of the specified OSS bucket name. This parameter can be left empty.
     * 
     */
    public Optional<Output<String>> ossKeyPrefix() {
        return Optional.ofNullable(this.ossKeyPrefix);
    }

    /**
     * The unique ARN of the Oss role.
     * 
     */
    @Import(name="ossWriteRoleArn")
    private @Nullable Output<String> ossWriteRoleArn;

    /**
     * @return The unique ARN of the Oss role.
     * 
     */
    public Optional<Output<String>> ossWriteRoleArn() {
        return Optional.ofNullable(this.ossWriteRoleArn);
    }

    /**
     * Field `name` has been deprecated from version 1.118.0.
     * 
     * @deprecated
     * Field &#39;role_name&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'role_name' has been deprecated from version 1.118.0 */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return Field `name` has been deprecated from version 1.118.0.
     * 
     * @deprecated
     * Field &#39;role_name&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'role_name' has been deprecated from version 1.118.0 */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
     * 
     */
    @Import(name="slsProjectArn")
    private @Nullable Output<String> slsProjectArn;

    /**
     * @return The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
     * 
     */
    public Optional<Output<String>> slsProjectArn() {
        return Optional.ofNullable(this.slsProjectArn);
    }

    /**
     * The unique ARN of the Log Service role.
     * 
     */
    @Import(name="slsWriteRoleArn")
    private @Nullable Output<String> slsWriteRoleArn;

    /**
     * @return The unique ARN of the Log Service role.
     * 
     */
    public Optional<Output<String>> slsWriteRoleArn() {
        return Optional.ofNullable(this.slsWriteRoleArn);
    }

    /**
     * The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The name of the trail to be created, which must be unique for an account.
     * 
     */
    @Import(name="trailName")
    private @Nullable Output<String> trailName;

    /**
     * @return The name of the trail to be created, which must be unique for an account.
     * 
     */
    public Optional<Output<String>> trailName() {
        return Optional.ofNullable(this.trailName);
    }

    /**
     * The regions to which the trail is applied. Default to `All`.
     * 
     */
    @Import(name="trailRegion")
    private @Nullable Output<String> trailRegion;

    /**
     * @return The regions to which the trail is applied. Default to `All`.
     * 
     */
    public Optional<Output<String>> trailRegion() {
        return Optional.ofNullable(this.trailRegion);
    }

    private TrailState() {}

    private TrailState(TrailState $) {
        this.eventRw = $.eventRw;
        this.isOrganizationTrail = $.isOrganizationTrail;
        this.mnsTopicArn = $.mnsTopicArn;
        this.name = $.name;
        this.ossBucketName = $.ossBucketName;
        this.ossKeyPrefix = $.ossKeyPrefix;
        this.ossWriteRoleArn = $.ossWriteRoleArn;
        this.roleName = $.roleName;
        this.slsProjectArn = $.slsProjectArn;
        this.slsWriteRoleArn = $.slsWriteRoleArn;
        this.status = $.status;
        this.trailName = $.trailName;
        this.trailRegion = $.trailRegion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrailState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrailState $;

        public Builder() {
            $ = new TrailState();
        }

        public Builder(TrailState defaults) {
            $ = new TrailState(Objects.requireNonNull(defaults));
        }

        /**
         * @param eventRw Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
         * 
         * @return builder
         * 
         */
        public Builder eventRw(@Nullable Output<String> eventRw) {
            $.eventRw = eventRw;
            return this;
        }

        /**
         * @param eventRw Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
         * 
         * @return builder
         * 
         */
        public Builder eventRw(String eventRw) {
            return eventRw(Output.of(eventRw));
        }

        /**
         * @param isOrganizationTrail Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
         * 
         * @return builder
         * 
         */
        public Builder isOrganizationTrail(@Nullable Output<Boolean> isOrganizationTrail) {
            $.isOrganizationTrail = isOrganizationTrail;
            return this;
        }

        /**
         * @param isOrganizationTrail Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
         * 
         * @return builder
         * 
         */
        public Builder isOrganizationTrail(Boolean isOrganizationTrail) {
            return isOrganizationTrail(Output.of(isOrganizationTrail));
        }

        /**
         * @param mnsTopicArn Field `mns_topic_arn` has been deprecated from version 1.118.0.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;mns_topic_arn&#39; has been deprecated from version 1.118.0
         * 
         */
        @Deprecated /* Field 'mns_topic_arn' has been deprecated from version 1.118.0 */
        public Builder mnsTopicArn(@Nullable Output<String> mnsTopicArn) {
            $.mnsTopicArn = mnsTopicArn;
            return this;
        }

        /**
         * @param mnsTopicArn Field `mns_topic_arn` has been deprecated from version 1.118.0.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;mns_topic_arn&#39; has been deprecated from version 1.118.0
         * 
         */
        @Deprecated /* Field 'mns_topic_arn' has been deprecated from version 1.118.0 */
        public Builder mnsTopicArn(String mnsTopicArn) {
            return mnsTopicArn(Output.of(mnsTopicArn));
        }

        /**
         * @param name Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from version 1.95.0. Use &#39;trail_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from version 1.95.0. Use &#39;trail_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param ossBucketName The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
         * 
         * @return builder
         * 
         */
        public Builder ossBucketName(@Nullable Output<String> ossBucketName) {
            $.ossBucketName = ossBucketName;
            return this;
        }

        /**
         * @param ossBucketName The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
         * 
         * @return builder
         * 
         */
        public Builder ossBucketName(String ossBucketName) {
            return ossBucketName(Output.of(ossBucketName));
        }

        /**
         * @param ossKeyPrefix The prefix of the specified OSS bucket name. This parameter can be left empty.
         * 
         * @return builder
         * 
         */
        public Builder ossKeyPrefix(@Nullable Output<String> ossKeyPrefix) {
            $.ossKeyPrefix = ossKeyPrefix;
            return this;
        }

        /**
         * @param ossKeyPrefix The prefix of the specified OSS bucket name. This parameter can be left empty.
         * 
         * @return builder
         * 
         */
        public Builder ossKeyPrefix(String ossKeyPrefix) {
            return ossKeyPrefix(Output.of(ossKeyPrefix));
        }

        /**
         * @param ossWriteRoleArn The unique ARN of the Oss role.
         * 
         * @return builder
         * 
         */
        public Builder ossWriteRoleArn(@Nullable Output<String> ossWriteRoleArn) {
            $.ossWriteRoleArn = ossWriteRoleArn;
            return this;
        }

        /**
         * @param ossWriteRoleArn The unique ARN of the Oss role.
         * 
         * @return builder
         * 
         */
        public Builder ossWriteRoleArn(String ossWriteRoleArn) {
            return ossWriteRoleArn(Output.of(ossWriteRoleArn));
        }

        /**
         * @param roleName Field `name` has been deprecated from version 1.118.0.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;role_name&#39; has been deprecated from version 1.118.0
         * 
         */
        @Deprecated /* Field 'role_name' has been deprecated from version 1.118.0 */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName Field `name` has been deprecated from version 1.118.0.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;role_name&#39; has been deprecated from version 1.118.0
         * 
         */
        @Deprecated /* Field 'role_name' has been deprecated from version 1.118.0 */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param slsProjectArn The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
         * 
         * @return builder
         * 
         */
        public Builder slsProjectArn(@Nullable Output<String> slsProjectArn) {
            $.slsProjectArn = slsProjectArn;
            return this;
        }

        /**
         * @param slsProjectArn The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
         * 
         * @return builder
         * 
         */
        public Builder slsProjectArn(String slsProjectArn) {
            return slsProjectArn(Output.of(slsProjectArn));
        }

        /**
         * @param slsWriteRoleArn The unique ARN of the Log Service role.
         * 
         * @return builder
         * 
         */
        public Builder slsWriteRoleArn(@Nullable Output<String> slsWriteRoleArn) {
            $.slsWriteRoleArn = slsWriteRoleArn;
            return this;
        }

        /**
         * @param slsWriteRoleArn The unique ARN of the Log Service role.
         * 
         * @return builder
         * 
         */
        public Builder slsWriteRoleArn(String slsWriteRoleArn) {
            return slsWriteRoleArn(Output.of(slsWriteRoleArn));
        }

        /**
         * @param status The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param trailName The name of the trail to be created, which must be unique for an account.
         * 
         * @return builder
         * 
         */
        public Builder trailName(@Nullable Output<String> trailName) {
            $.trailName = trailName;
            return this;
        }

        /**
         * @param trailName The name of the trail to be created, which must be unique for an account.
         * 
         * @return builder
         * 
         */
        public Builder trailName(String trailName) {
            return trailName(Output.of(trailName));
        }

        /**
         * @param trailRegion The regions to which the trail is applied. Default to `All`.
         * 
         * @return builder
         * 
         */
        public Builder trailRegion(@Nullable Output<String> trailRegion) {
            $.trailRegion = trailRegion;
            return this;
        }

        /**
         * @param trailRegion The regions to which the trail is applied. Default to `All`.
         * 
         * @return builder
         * 
         */
        public Builder trailRegion(String trailRegion) {
            return trailRegion(Output.of(trailRegion));
        }

        public TrailState build() {
            return $;
        }
    }

}