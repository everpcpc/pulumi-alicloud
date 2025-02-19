// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emrv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterNodeAttributeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterNodeAttributeArgs Empty = new ClusterNodeAttributeArgs();

    /**
     * Whether to enable data disk encryption.
     * 
     */
    @Import(name="dataDiskEncrypted")
    private @Nullable Output<Boolean> dataDiskEncrypted;

    /**
     * @return Whether to enable data disk encryption.
     * 
     */
    public Optional<Output<Boolean>> dataDiskEncrypted() {
        return Optional.ofNullable(this.dataDiskEncrypted);
    }

    /**
     * The kms key id used to encrypt the data disk. It takes effect when data_disk_encrypted is true.
     * 
     */
    @Import(name="dataDiskKmsKeyId")
    private @Nullable Output<String> dataDiskKmsKeyId;

    /**
     * @return The kms key id used to encrypt the data disk. It takes effect when data_disk_encrypted is true.
     * 
     */
    public Optional<Output<String>> dataDiskKmsKeyId() {
        return Optional.ofNullable(this.dataDiskKmsKeyId);
    }

    /**
     * The name of the key pair.
     * 
     */
    @Import(name="keyPairName", required=true)
    private Output<String> keyPairName;

    /**
     * @return The name of the key pair.
     * 
     */
    public Output<String> keyPairName() {
        return this.keyPairName;
    }

    /**
     * Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * 
     */
    @Import(name="ramRole", required=true)
    private Output<String> ramRole;

    /**
     * @return Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * 
     */
    public Output<String> ramRole() {
        return this.ramRole;
    }

    /**
     * Security Group ID for Cluster.
     * 
     */
    @Import(name="securityGroupId", required=true)
    private Output<String> securityGroupId;

    /**
     * @return Security Group ID for Cluster.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }

    /**
     * Used to retrieve instances belong to specified VPC.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return Used to retrieve instances belong to specified VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     * Zone ID, e.g. cn-hangzhou-i
     * 
     */
    @Import(name="zoneId", required=true)
    private Output<String> zoneId;

    /**
     * @return Zone ID, e.g. cn-hangzhou-i
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    private ClusterNodeAttributeArgs() {}

    private ClusterNodeAttributeArgs(ClusterNodeAttributeArgs $) {
        this.dataDiskEncrypted = $.dataDiskEncrypted;
        this.dataDiskKmsKeyId = $.dataDiskKmsKeyId;
        this.keyPairName = $.keyPairName;
        this.ramRole = $.ramRole;
        this.securityGroupId = $.securityGroupId;
        this.vpcId = $.vpcId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterNodeAttributeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterNodeAttributeArgs $;

        public Builder() {
            $ = new ClusterNodeAttributeArgs();
        }

        public Builder(ClusterNodeAttributeArgs defaults) {
            $ = new ClusterNodeAttributeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dataDiskEncrypted Whether to enable data disk encryption.
         * 
         * @return builder
         * 
         */
        public Builder dataDiskEncrypted(@Nullable Output<Boolean> dataDiskEncrypted) {
            $.dataDiskEncrypted = dataDiskEncrypted;
            return this;
        }

        /**
         * @param dataDiskEncrypted Whether to enable data disk encryption.
         * 
         * @return builder
         * 
         */
        public Builder dataDiskEncrypted(Boolean dataDiskEncrypted) {
            return dataDiskEncrypted(Output.of(dataDiskEncrypted));
        }

        /**
         * @param dataDiskKmsKeyId The kms key id used to encrypt the data disk. It takes effect when data_disk_encrypted is true.
         * 
         * @return builder
         * 
         */
        public Builder dataDiskKmsKeyId(@Nullable Output<String> dataDiskKmsKeyId) {
            $.dataDiskKmsKeyId = dataDiskKmsKeyId;
            return this;
        }

        /**
         * @param dataDiskKmsKeyId The kms key id used to encrypt the data disk. It takes effect when data_disk_encrypted is true.
         * 
         * @return builder
         * 
         */
        public Builder dataDiskKmsKeyId(String dataDiskKmsKeyId) {
            return dataDiskKmsKeyId(Output.of(dataDiskKmsKeyId));
        }

        /**
         * @param keyPairName The name of the key pair.
         * 
         * @return builder
         * 
         */
        public Builder keyPairName(Output<String> keyPairName) {
            $.keyPairName = keyPairName;
            return this;
        }

        /**
         * @param keyPairName The name of the key pair.
         * 
         * @return builder
         * 
         */
        public Builder keyPairName(String keyPairName) {
            return keyPairName(Output.of(keyPairName));
        }

        /**
         * @param ramRole Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
         * 
         * @return builder
         * 
         */
        public Builder ramRole(Output<String> ramRole) {
            $.ramRole = ramRole;
            return this;
        }

        /**
         * @param ramRole Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
         * 
         * @return builder
         * 
         */
        public Builder ramRole(String ramRole) {
            return ramRole(Output.of(ramRole));
        }

        /**
         * @param securityGroupId Security Group ID for Cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(Output<String> securityGroupId) {
            $.securityGroupId = securityGroupId;
            return this;
        }

        /**
         * @param securityGroupId Security Group ID for Cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(String securityGroupId) {
            return securityGroupId(Output.of(securityGroupId));
        }

        /**
         * @param vpcId Used to retrieve instances belong to specified VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId Used to retrieve instances belong to specified VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param zoneId Zone ID, e.g. cn-hangzhou-i
         * 
         * @return builder
         * 
         */
        public Builder zoneId(Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId Zone ID, e.g. cn-hangzhou-i
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public ClusterNodeAttributeArgs build() {
            $.keyPairName = Objects.requireNonNull($.keyPairName, "expected parameter 'keyPairName' to be non-null");
            $.ramRole = Objects.requireNonNull($.ramRole, "expected parameter 'ramRole' to be non-null");
            $.securityGroupId = Objects.requireNonNull($.securityGroupId, "expected parameter 'securityGroupId' to be non-null");
            $.vpcId = Objects.requireNonNull($.vpcId, "expected parameter 'vpcId' to be non-null");
            $.zoneId = Objects.requireNonNull($.zoneId, "expected parameter 'zoneId' to be non-null");
            return $;
        }
    }

}
