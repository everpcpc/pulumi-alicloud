// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHostAccountsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHostAccountsPlainArgs Empty = new GetHostAccountsPlainArgs();

    /**
     * Specify the new hosting account&#39;s name, support the longest 128 characters.
     * 
     */
    @Import(name="hostAccountName")
    private @Nullable String hostAccountName;

    /**
     * @return Specify the new hosting account&#39;s name, support the longest 128 characters.
     * 
     */
    public Optional<String> hostAccountName() {
        return Optional.ofNullable(this.hostAccountName);
    }

    /**
     * Specifies the database where you want to create your hosting account&#39;s host ID.
     * 
     */
    @Import(name="hostId", required=true)
    private String hostId;

    /**
     * @return Specifies the database where you want to create your hosting account&#39;s host ID.
     * 
     */
    public String hostId() {
        return this.hostId;
    }

    /**
     * A list of Host Account IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Host Account IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Specifies the database where you want to create your hosting account&#39;s host bastion host ID of.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return Specifies the database where you want to create your hosting account&#39;s host bastion host ID of.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }

    /**
     * A regex string to filter results by Host Account name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Host Account name.
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
     * Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
     * 
     */
    @Import(name="protocolName")
    private @Nullable String protocolName;

    /**
     * @return Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
     * 
     */
    public Optional<String> protocolName() {
        return Optional.ofNullable(this.protocolName);
    }

    private GetHostAccountsPlainArgs() {}

    private GetHostAccountsPlainArgs(GetHostAccountsPlainArgs $) {
        this.hostAccountName = $.hostAccountName;
        this.hostId = $.hostId;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.protocolName = $.protocolName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHostAccountsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHostAccountsPlainArgs $;

        public Builder() {
            $ = new GetHostAccountsPlainArgs();
        }

        public Builder(GetHostAccountsPlainArgs defaults) {
            $ = new GetHostAccountsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hostAccountName Specify the new hosting account&#39;s name, support the longest 128 characters.
         * 
         * @return builder
         * 
         */
        public Builder hostAccountName(@Nullable String hostAccountName) {
            $.hostAccountName = hostAccountName;
            return this;
        }

        /**
         * @param hostId Specifies the database where you want to create your hosting account&#39;s host ID.
         * 
         * @return builder
         * 
         */
        public Builder hostId(String hostId) {
            $.hostId = hostId;
            return this;
        }

        /**
         * @param ids A list of Host Account IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Host Account IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId Specifies the database where you want to create your hosting account&#39;s host bastion host ID of.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Host Account name.
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
         * @param protocolName Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
         * 
         * @return builder
         * 
         */
        public Builder protocolName(@Nullable String protocolName) {
            $.protocolName = protocolName;
            return this;
        }

        public GetHostAccountsPlainArgs build() {
            $.hostId = Objects.requireNonNull($.hostId, "expected parameter 'hostId' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}