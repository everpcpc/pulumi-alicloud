// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final BgpGroupArgs Empty = new BgpGroupArgs();

    /**
     * The authentication key of the BGP group.
     * 
     */
    @Import(name="authKey")
    private @Nullable Output<String> authKey;

    /**
     * @return The authentication key of the BGP group.
     * 
     */
    public Optional<Output<String>> authKey() {
        return Optional.ofNullable(this.authKey);
    }

    /**
     * The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="bgpGroupName")
    private @Nullable Output<String> bgpGroupName;

    /**
     * @return The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> bgpGroupName() {
        return Optional.ofNullable(this.bgpGroupName);
    }

    /**
     * The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
     * 
     */
    @Import(name="isFakeAsn")
    private @Nullable Output<Boolean> isFakeAsn;

    /**
     * @return The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
     * 
     */
    public Optional<Output<Boolean>> isFakeAsn() {
        return Optional.ofNullable(this.isFakeAsn);
    }

    /**
     * The AS number on the Alibaba Cloud side.
     * 
     */
    @Import(name="localAsn")
    private @Nullable Output<Integer> localAsn;

    /**
     * @return The AS number on the Alibaba Cloud side.
     * 
     */
    public Optional<Output<Integer>> localAsn() {
        return Optional.ofNullable(this.localAsn);
    }

    /**
     * The AS number of the BGP peer.
     * 
     */
    @Import(name="peerAsn", required=true)
    private Output<Integer> peerAsn;

    /**
     * @return The AS number of the BGP peer.
     * 
     */
    public Output<Integer> peerAsn() {
        return this.peerAsn;
    }

    /**
     * The ID of the VBR.
     * 
     */
    @Import(name="routerId", required=true)
    private Output<String> routerId;

    /**
     * @return The ID of the VBR.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }

    private BgpGroupArgs() {}

    private BgpGroupArgs(BgpGroupArgs $) {
        this.authKey = $.authKey;
        this.bgpGroupName = $.bgpGroupName;
        this.description = $.description;
        this.isFakeAsn = $.isFakeAsn;
        this.localAsn = $.localAsn;
        this.peerAsn = $.peerAsn;
        this.routerId = $.routerId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpGroupArgs $;

        public Builder() {
            $ = new BgpGroupArgs();
        }

        public Builder(BgpGroupArgs defaults) {
            $ = new BgpGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authKey The authentication key of the BGP group.
         * 
         * @return builder
         * 
         */
        public Builder authKey(@Nullable Output<String> authKey) {
            $.authKey = authKey;
            return this;
        }

        /**
         * @param authKey The authentication key of the BGP group.
         * 
         * @return builder
         * 
         */
        public Builder authKey(String authKey) {
            return authKey(Output.of(authKey));
        }

        /**
         * @param bgpGroupName The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder bgpGroupName(@Nullable Output<String> bgpGroupName) {
            $.bgpGroupName = bgpGroupName;
            return this;
        }

        /**
         * @param bgpGroupName The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder bgpGroupName(String bgpGroupName) {
            return bgpGroupName(Output.of(bgpGroupName));
        }

        /**
         * @param description The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param isFakeAsn The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
         * 
         * @return builder
         * 
         */
        public Builder isFakeAsn(@Nullable Output<Boolean> isFakeAsn) {
            $.isFakeAsn = isFakeAsn;
            return this;
        }

        /**
         * @param isFakeAsn The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
         * 
         * @return builder
         * 
         */
        public Builder isFakeAsn(Boolean isFakeAsn) {
            return isFakeAsn(Output.of(isFakeAsn));
        }

        /**
         * @param localAsn The AS number on the Alibaba Cloud side.
         * 
         * @return builder
         * 
         */
        public Builder localAsn(@Nullable Output<Integer> localAsn) {
            $.localAsn = localAsn;
            return this;
        }

        /**
         * @param localAsn The AS number on the Alibaba Cloud side.
         * 
         * @return builder
         * 
         */
        public Builder localAsn(Integer localAsn) {
            return localAsn(Output.of(localAsn));
        }

        /**
         * @param peerAsn The AS number of the BGP peer.
         * 
         * @return builder
         * 
         */
        public Builder peerAsn(Output<Integer> peerAsn) {
            $.peerAsn = peerAsn;
            return this;
        }

        /**
         * @param peerAsn The AS number of the BGP peer.
         * 
         * @return builder
         * 
         */
        public Builder peerAsn(Integer peerAsn) {
            return peerAsn(Output.of(peerAsn));
        }

        /**
         * @param routerId The ID of the VBR.
         * 
         * @return builder
         * 
         */
        public Builder routerId(Output<String> routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerId The ID of the VBR.
         * 
         * @return builder
         * 
         */
        public Builder routerId(String routerId) {
            return routerId(Output.of(routerId));
        }

        public BgpGroupArgs build() {
            $.peerAsn = Objects.requireNonNull($.peerAsn, "expected parameter 'peerAsn' to be non-null");
            $.routerId = Objects.requireNonNull($.routerId, "expected parameter 'routerId' to be non-null");
            return $;
        }
    }

}