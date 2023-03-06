// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpn.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionIkeConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionIkeConfigArgs Empty = new ConnectionIkeConfigArgs();

    /**
     * The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 . Default value: md5
     * 
     */
    @Import(name="ikeAuthAlg")
    private @Nullable Output<String> ikeAuthAlg;

    /**
     * @return The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 . Default value: md5
     * 
     */
    public Optional<Output<String>> ikeAuthAlg() {
        return Optional.ofNullable(this.ikeAuthAlg);
    }

    /**
     * The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
     * 
     */
    @Import(name="ikeEncAlg")
    private @Nullable Output<String> ikeEncAlg;

    /**
     * @return The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
     * 
     */
    public Optional<Output<String>> ikeEncAlg() {
        return Optional.ofNullable(this.ikeEncAlg);
    }

    /**
     * The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
     * 
     */
    @Import(name="ikeLifetime")
    private @Nullable Output<Integer> ikeLifetime;

    /**
     * @return The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
     * 
     */
    public Optional<Output<Integer>> ikeLifetime() {
        return Optional.ofNullable(this.ikeLifetime);
    }

    /**
     * The identification of the VPN gateway.
     * 
     */
    @Import(name="ikeLocalId")
    private @Nullable Output<String> ikeLocalId;

    /**
     * @return The identification of the VPN gateway.
     * 
     */
    public Optional<Output<String>> ikeLocalId() {
        return Optional.ofNullable(this.ikeLocalId);
    }

    /**
     * The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
     * 
     */
    @Import(name="ikeMode")
    private @Nullable Output<String> ikeMode;

    /**
     * @return The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
     * 
     */
    public Optional<Output<String>> ikeMode() {
        return Optional.ofNullable(this.ikeMode);
    }

    /**
     * The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
     * 
     */
    @Import(name="ikePfs")
    private @Nullable Output<String> ikePfs;

    /**
     * @return The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
     * 
     */
    public Optional<Output<String>> ikePfs() {
        return Optional.ofNullable(this.ikePfs);
    }

    /**
     * The identification of the customer gateway.
     * 
     */
    @Import(name="ikeRemoteId")
    private @Nullable Output<String> ikeRemoteId;

    /**
     * @return The identification of the customer gateway.
     * 
     */
    public Optional<Output<String>> ikeRemoteId() {
        return Optional.ofNullable(this.ikeRemoteId);
    }

    /**
     * The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
     * 
     */
    @Import(name="ikeVersion")
    private @Nullable Output<String> ikeVersion;

    /**
     * @return The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
     * 
     */
    public Optional<Output<String>> ikeVersion() {
        return Optional.ofNullable(this.ikeVersion);
    }

    /**
     * Used for authentication between the IPsec VPN gateway and the customer gateway.
     * 
     */
    @Import(name="psk")
    private @Nullable Output<String> psk;

    /**
     * @return Used for authentication between the IPsec VPN gateway and the customer gateway.
     * 
     */
    public Optional<Output<String>> psk() {
        return Optional.ofNullable(this.psk);
    }

    private ConnectionIkeConfigArgs() {}

    private ConnectionIkeConfigArgs(ConnectionIkeConfigArgs $) {
        this.ikeAuthAlg = $.ikeAuthAlg;
        this.ikeEncAlg = $.ikeEncAlg;
        this.ikeLifetime = $.ikeLifetime;
        this.ikeLocalId = $.ikeLocalId;
        this.ikeMode = $.ikeMode;
        this.ikePfs = $.ikePfs;
        this.ikeRemoteId = $.ikeRemoteId;
        this.ikeVersion = $.ikeVersion;
        this.psk = $.psk;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionIkeConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionIkeConfigArgs $;

        public Builder() {
            $ = new ConnectionIkeConfigArgs();
        }

        public Builder(ConnectionIkeConfigArgs defaults) {
            $ = new ConnectionIkeConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ikeAuthAlg The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 . Default value: md5
         * 
         * @return builder
         * 
         */
        public Builder ikeAuthAlg(@Nullable Output<String> ikeAuthAlg) {
            $.ikeAuthAlg = ikeAuthAlg;
            return this;
        }

        /**
         * @param ikeAuthAlg The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 . Default value: md5
         * 
         * @return builder
         * 
         */
        public Builder ikeAuthAlg(String ikeAuthAlg) {
            return ikeAuthAlg(Output.of(ikeAuthAlg));
        }

        /**
         * @param ikeEncAlg The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
         * 
         * @return builder
         * 
         */
        public Builder ikeEncAlg(@Nullable Output<String> ikeEncAlg) {
            $.ikeEncAlg = ikeEncAlg;
            return this;
        }

        /**
         * @param ikeEncAlg The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
         * 
         * @return builder
         * 
         */
        public Builder ikeEncAlg(String ikeEncAlg) {
            return ikeEncAlg(Output.of(ikeEncAlg));
        }

        /**
         * @param ikeLifetime The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
         * 
         * @return builder
         * 
         */
        public Builder ikeLifetime(@Nullable Output<Integer> ikeLifetime) {
            $.ikeLifetime = ikeLifetime;
            return this;
        }

        /**
         * @param ikeLifetime The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
         * 
         * @return builder
         * 
         */
        public Builder ikeLifetime(Integer ikeLifetime) {
            return ikeLifetime(Output.of(ikeLifetime));
        }

        /**
         * @param ikeLocalId The identification of the VPN gateway.
         * 
         * @return builder
         * 
         */
        public Builder ikeLocalId(@Nullable Output<String> ikeLocalId) {
            $.ikeLocalId = ikeLocalId;
            return this;
        }

        /**
         * @param ikeLocalId The identification of the VPN gateway.
         * 
         * @return builder
         * 
         */
        public Builder ikeLocalId(String ikeLocalId) {
            return ikeLocalId(Output.of(ikeLocalId));
        }

        /**
         * @param ikeMode The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
         * 
         * @return builder
         * 
         */
        public Builder ikeMode(@Nullable Output<String> ikeMode) {
            $.ikeMode = ikeMode;
            return this;
        }

        /**
         * @param ikeMode The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
         * 
         * @return builder
         * 
         */
        public Builder ikeMode(String ikeMode) {
            return ikeMode(Output.of(ikeMode));
        }

        /**
         * @param ikePfs The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
         * 
         * @return builder
         * 
         */
        public Builder ikePfs(@Nullable Output<String> ikePfs) {
            $.ikePfs = ikePfs;
            return this;
        }

        /**
         * @param ikePfs The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
         * 
         * @return builder
         * 
         */
        public Builder ikePfs(String ikePfs) {
            return ikePfs(Output.of(ikePfs));
        }

        /**
         * @param ikeRemoteId The identification of the customer gateway.
         * 
         * @return builder
         * 
         */
        public Builder ikeRemoteId(@Nullable Output<String> ikeRemoteId) {
            $.ikeRemoteId = ikeRemoteId;
            return this;
        }

        /**
         * @param ikeRemoteId The identification of the customer gateway.
         * 
         * @return builder
         * 
         */
        public Builder ikeRemoteId(String ikeRemoteId) {
            return ikeRemoteId(Output.of(ikeRemoteId));
        }

        /**
         * @param ikeVersion The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
         * 
         * @return builder
         * 
         */
        public Builder ikeVersion(@Nullable Output<String> ikeVersion) {
            $.ikeVersion = ikeVersion;
            return this;
        }

        /**
         * @param ikeVersion The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
         * 
         * @return builder
         * 
         */
        public Builder ikeVersion(String ikeVersion) {
            return ikeVersion(Output.of(ikeVersion));
        }

        /**
         * @param psk Used for authentication between the IPsec VPN gateway and the customer gateway.
         * 
         * @return builder
         * 
         */
        public Builder psk(@Nullable Output<String> psk) {
            $.psk = psk;
            return this;
        }

        /**
         * @param psk Used for authentication between the IPsec VPN gateway and the customer gateway.
         * 
         * @return builder
         * 
         */
        public Builder psk(String psk) {
            return psk(Output.of(psk));
        }

        public ConnectionIkeConfigArgs build() {
            return $;
        }
    }

}