// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUsersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUsersArgs Empty = new GetUsersArgs();

    /**
     * Specify the New Created the User&#39;s Display Name. Supports up to 128 Characters.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Specify the New Created the User&#39;s Display Name. Supports up to 128 Characters.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * A list of User IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of User IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * You Want to Query the User the Bastion Host ID of.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return You Want to Query the User the Bastion Host ID of.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * Specify the New of the User That Created a Different Mobile Phone Number from Your.
     * 
     */
    @Import(name="mobile")
    private @Nullable Output<String> mobile;

    /**
     * @return Specify the New of the User That Created a Different Mobile Phone Number from Your.
     * 
     */
    public Optional<Output<String>> mobile() {
        return Optional.ofNullable(this.mobile);
    }

    /**
     * A regex string to filter results by User name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by User name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Specify the New of the User That Created the Source. Valid Values: Local: Local User RAM: Ram User.
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return Specify the New of the User That Created the Source. Valid Values: Local: Local User RAM: Ram User.
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    /**
     * Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User&#39;s Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
     * 
     */
    @Import(name="sourceUserId")
    private @Nullable Output<String> sourceUserId;

    /**
     * @return Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User&#39;s Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
     * 
     */
    public Optional<Output<String>> sourceUserId() {
        return Optional.ofNullable(this.sourceUserId);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private GetUsersArgs() {}

    private GetUsersArgs(GetUsersArgs $) {
        this.displayName = $.displayName;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.mobile = $.mobile;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.source = $.source;
        this.sourceUserId = $.sourceUserId;
        this.status = $.status;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUsersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUsersArgs $;

        public Builder() {
            $ = new GetUsersArgs();
        }

        public Builder(GetUsersArgs defaults) {
            $ = new GetUsersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName Specify the New Created the User&#39;s Display Name. Supports up to 128 Characters.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Specify the New Created the User&#39;s Display Name. Supports up to 128 Characters.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param ids A list of User IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of User IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of User IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId You Want to Query the User the Bastion Host ID of.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId You Want to Query the User the Bastion Host ID of.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param mobile Specify the New of the User That Created a Different Mobile Phone Number from Your.
         * 
         * @return builder
         * 
         */
        public Builder mobile(@Nullable Output<String> mobile) {
            $.mobile = mobile;
            return this;
        }

        /**
         * @param mobile Specify the New of the User That Created a Different Mobile Phone Number from Your.
         * 
         * @return builder
         * 
         */
        public Builder mobile(String mobile) {
            return mobile(Output.of(mobile));
        }

        /**
         * @param nameRegex A regex string to filter results by User name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by User name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param source Specify the New of the User That Created the Source. Valid Values: Local: Local User RAM: Ram User.
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source Specify the New of the User That Created the Source. Valid Values: Local: Local User RAM: Ram User.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        /**
         * @param sourceUserId Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User&#39;s Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
         * 
         * @return builder
         * 
         */
        public Builder sourceUserId(@Nullable Output<String> sourceUserId) {
            $.sourceUserId = sourceUserId;
            return this;
        }

        /**
         * @param sourceUserId Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User&#39;s Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
         * 
         * @return builder
         * 
         */
        public Builder sourceUserId(String sourceUserId) {
            return sourceUserId(Output.of(sourceUserId));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param userName Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public GetUsersArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
