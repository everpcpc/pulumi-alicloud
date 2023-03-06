// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupsPlainArgs Empty = new GetGroupsPlainArgs();

    /**
     * A regex string to filter the returned groups by their names.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter the returned groups by their names.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Filter the results by a specific policy name. If you set this parameter without setting `policy_type`, it will be automatically set to `System`.
     * 
     */
    @Import(name="policyName")
    private @Nullable String policyName;

    /**
     * @return Filter the results by a specific policy name. If you set this parameter without setting `policy_type`, it will be automatically set to `System`.
     * 
     */
    public Optional<String> policyName() {
        return Optional.ofNullable(this.policyName);
    }

    /**
     * Filter the results by a specific policy type. Valid items are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
     * 
     */
    @Import(name="policyType")
    private @Nullable String policyType;

    /**
     * @return Filter the results by a specific policy type. Valid items are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
     * 
     */
    public Optional<String> policyType() {
        return Optional.ofNullable(this.policyType);
    }

    /**
     * Filter the results by a specific the user name.
     * 
     */
    @Import(name="userName")
    private @Nullable String userName;

    /**
     * @return Filter the results by a specific the user name.
     * 
     */
    public Optional<String> userName() {
        return Optional.ofNullable(this.userName);
    }

    private GetGroupsPlainArgs() {}

    private GetGroupsPlainArgs(GetGroupsPlainArgs $) {
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.policyName = $.policyName;
        this.policyType = $.policyType;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupsPlainArgs $;

        public Builder() {
            $ = new GetGroupsPlainArgs();
        }

        public Builder(GetGroupsPlainArgs defaults) {
            $ = new GetGroupsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param nameRegex A regex string to filter the returned groups by their names.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param policyName Filter the results by a specific policy name. If you set this parameter without setting `policy_type`, it will be automatically set to `System`.
         * 
         * @return builder
         * 
         */
        public Builder policyName(@Nullable String policyName) {
            $.policyName = policyName;
            return this;
        }

        /**
         * @param policyType Filter the results by a specific policy type. Valid items are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
         * 
         * @return builder
         * 
         */
        public Builder policyType(@Nullable String policyType) {
            $.policyType = policyType;
            return this;
        }

        /**
         * @param userName Filter the results by a specific the user name.
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable String userName) {
            $.userName = userName;
            return this;
        }

        public GetGroupsPlainArgs build() {
            return $;
        }
    }

}