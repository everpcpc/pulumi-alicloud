// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLoadBalancersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLoadBalancersPlainArgs Empty = new GetLoadBalancersPlainArgs();

    /**
     * Service address of the SLBs.
     * 
     */
    @Import(name="address")
    private @Nullable String address;

    /**
     * @return Service address of the SLBs.
     * 
     */
    public Optional<String> address() {
        return Optional.ofNullable(this.address);
    }

    @Import(name="addressIpVersion")
    private @Nullable String addressIpVersion;

    public Optional<String> addressIpVersion() {
        return Optional.ofNullable(this.addressIpVersion);
    }

    @Import(name="addressType")
    private @Nullable String addressType;

    public Optional<String> addressType() {
        return Optional.ofNullable(this.addressType);
    }

    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of SLBs IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of SLBs IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="internetChargeType")
    private @Nullable String internetChargeType;

    public Optional<String> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    @Import(name="loadBalancerName")
    private @Nullable String loadBalancerName;

    public Optional<String> loadBalancerName() {
        return Optional.ofNullable(this.loadBalancerName);
    }

    @Import(name="masterZoneId")
    private @Nullable String masterZoneId;

    public Optional<String> masterZoneId() {
        return Optional.ofNullable(this.masterZoneId);
    }

    /**
     * A regex string to filter results by SLB name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by SLB name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * Network type of the SLBs. Valid values: `vpc` and `classic`.
     * 
     */
    @Import(name="networkType")
    private @Nullable String networkType;

    /**
     * @return Network type of the SLBs. Valid values: `vpc` and `classic`.
     * 
     */
    public Optional<String> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    @Import(name="paymentType")
    private @Nullable String paymentType;

    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The Id of resource group which SLB belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The Id of resource group which SLB belongs.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="serverId")
    private @Nullable String serverId;

    public Optional<String> serverId() {
        return Optional.ofNullable(this.serverId);
    }

    @Import(name="serverIntranetAddress")
    private @Nullable String serverIntranetAddress;

    public Optional<String> serverIntranetAddress() {
        return Optional.ofNullable(this.serverIntranetAddress);
    }

    @Import(name="slaveZoneId")
    private @Nullable String slaveZoneId;

    public Optional<String> slaveZoneId() {
        return Optional.ofNullable(this.slaveZoneId);
    }

    /**
     * SLB current status. Possible values: `inactive`, `active` and `locked`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return SLB current status. Possible values: `inactive`, `active` and `locked`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.slb.SlbFunctions;
     * import com.pulumi.alicloud.slb.inputs.GetLoadBalancersArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var taggedInstances = SlbFunctions.getLoadBalancers(GetLoadBalancersArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
     *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
     *             ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.slb.SlbFunctions;
     * import com.pulumi.alicloud.slb.inputs.GetLoadBalancersArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var taggedInstances = SlbFunctions.getLoadBalancers(GetLoadBalancersArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
     *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
     *             ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * ID of the VPC linked to the SLBs.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return ID of the VPC linked to the SLBs.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * ID of the VSwitch linked to the SLBs.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable String vswitchId;

    /**
     * @return ID of the VSwitch linked to the SLBs.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetLoadBalancersPlainArgs() {}

    private GetLoadBalancersPlainArgs(GetLoadBalancersPlainArgs $) {
        this.address = $.address;
        this.addressIpVersion = $.addressIpVersion;
        this.addressType = $.addressType;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.internetChargeType = $.internetChargeType;
        this.loadBalancerName = $.loadBalancerName;
        this.masterZoneId = $.masterZoneId;
        this.nameRegex = $.nameRegex;
        this.networkType = $.networkType;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.paymentType = $.paymentType;
        this.resourceGroupId = $.resourceGroupId;
        this.serverId = $.serverId;
        this.serverIntranetAddress = $.serverIntranetAddress;
        this.slaveZoneId = $.slaveZoneId;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLoadBalancersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLoadBalancersPlainArgs $;

        public Builder() {
            $ = new GetLoadBalancersPlainArgs();
        }

        public Builder(GetLoadBalancersPlainArgs defaults) {
            $ = new GetLoadBalancersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Service address of the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable String address) {
            $.address = address;
            return this;
        }

        public Builder addressIpVersion(@Nullable String addressIpVersion) {
            $.addressIpVersion = addressIpVersion;
            return this;
        }

        public Builder addressType(@Nullable String addressType) {
            $.addressType = addressType;
            return this;
        }

        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of SLBs IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of SLBs IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder internetChargeType(@Nullable String internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        public Builder loadBalancerName(@Nullable String loadBalancerName) {
            $.loadBalancerName = loadBalancerName;
            return this;
        }

        public Builder masterZoneId(@Nullable String masterZoneId) {
            $.masterZoneId = masterZoneId;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by SLB name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param networkType Network type of the SLBs. Valid values: `vpc` and `classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable String networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        public Builder paymentType(@Nullable String paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which SLB belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        public Builder serverId(@Nullable String serverId) {
            $.serverId = serverId;
            return this;
        }

        public Builder serverIntranetAddress(@Nullable String serverIntranetAddress) {
            $.serverIntranetAddress = serverIntranetAddress;
            return this;
        }

        public Builder slaveZoneId(@Nullable String slaveZoneId) {
            $.slaveZoneId = slaveZoneId;
            return this;
        }

        /**
         * @param status SLB current status. Possible values: `inactive`, `active` and `locked`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import com.pulumi.alicloud.slb.SlbFunctions;
         * import com.pulumi.alicloud.slb.inputs.GetLoadBalancersArgs;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *         final var taggedInstances = SlbFunctions.getLoadBalancers(GetLoadBalancersArgs.builder()
         *             .tags(Map.ofEntries(
         *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
         *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
         *             ))
         *             .build());
         * 
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param vpcId ID of the VPC linked to the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vswitchId ID of the VSwitch linked to the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable String vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        public GetLoadBalancersPlainArgs build() {
            return $;
        }
    }

}
