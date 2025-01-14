// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.inputs;

import com.pulumi.core.Output;
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


public final class GetLoadBalancersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLoadBalancersArgs Empty = new GetLoadBalancersArgs();

    /**
     * Service address of the SLBs.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return Service address of the SLBs.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    @Import(name="addressIpVersion")
    private @Nullable Output<String> addressIpVersion;

    public Optional<Output<String>> addressIpVersion() {
        return Optional.ofNullable(this.addressIpVersion);
    }

    @Import(name="addressType")
    private @Nullable Output<String> addressType;

    public Optional<Output<String>> addressType() {
        return Optional.ofNullable(this.addressType);
    }

    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of SLBs IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of SLBs IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="internetChargeType")
    private @Nullable Output<String> internetChargeType;

    public Optional<Output<String>> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    @Import(name="loadBalancerName")
    private @Nullable Output<String> loadBalancerName;

    public Optional<Output<String>> loadBalancerName() {
        return Optional.ofNullable(this.loadBalancerName);
    }

    @Import(name="masterZoneId")
    private @Nullable Output<String> masterZoneId;

    public Optional<Output<String>> masterZoneId() {
        return Optional.ofNullable(this.masterZoneId);
    }

    /**
     * A regex string to filter results by SLB name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by SLB name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * Network type of the SLBs. Valid values: `vpc` and `classic`.
     * 
     */
    @Import(name="networkType")
    private @Nullable Output<String> networkType;

    /**
     * @return Network type of the SLBs. Valid values: `vpc` and `classic`.
     * 
     */
    public Optional<Output<String>> networkType() {
        return Optional.ofNullable(this.networkType);
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

    @Import(name="pageNumber")
    private @Nullable Output<Integer> pageNumber;

    public Optional<Output<Integer>> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Output<Integer> pageSize;

    public Optional<Output<Integer>> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The Id of resource group which SLB belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which SLB belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="serverId")
    private @Nullable Output<String> serverId;

    public Optional<Output<String>> serverId() {
        return Optional.ofNullable(this.serverId);
    }

    @Import(name="serverIntranetAddress")
    private @Nullable Output<String> serverIntranetAddress;

    public Optional<Output<String>> serverIntranetAddress() {
        return Optional.ofNullable(this.serverIntranetAddress);
    }

    @Import(name="slaveZoneId")
    private @Nullable Output<String> slaveZoneId;

    public Optional<Output<String>> slaveZoneId() {
        return Optional.ofNullable(this.slaveZoneId);
    }

    /**
     * SLB current status. Possible values: `inactive`, `active` and `locked`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return SLB current status. Possible values: `inactive`, `active` and `locked`.
     * 
     */
    public Optional<Output<String>> status() {
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
    private @Nullable Output<Map<String,Object>> tags;

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
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * ID of the VPC linked to the SLBs.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return ID of the VPC linked to the SLBs.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * ID of the VSwitch linked to the SLBs.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return ID of the VSwitch linked to the SLBs.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetLoadBalancersArgs() {}

    private GetLoadBalancersArgs(GetLoadBalancersArgs $) {
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
    public static Builder builder(GetLoadBalancersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLoadBalancersArgs $;

        public Builder() {
            $ = new GetLoadBalancersArgs();
        }

        public Builder(GetLoadBalancersArgs defaults) {
            $ = new GetLoadBalancersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Service address of the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Service address of the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        public Builder addressIpVersion(@Nullable Output<String> addressIpVersion) {
            $.addressIpVersion = addressIpVersion;
            return this;
        }

        public Builder addressIpVersion(String addressIpVersion) {
            return addressIpVersion(Output.of(addressIpVersion));
        }

        public Builder addressType(@Nullable Output<String> addressType) {
            $.addressType = addressType;
            return this;
        }

        public Builder addressType(String addressType) {
            return addressType(Output.of(addressType));
        }

        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param ids A list of SLBs IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of SLBs IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
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

        public Builder internetChargeType(@Nullable Output<String> internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        public Builder internetChargeType(String internetChargeType) {
            return internetChargeType(Output.of(internetChargeType));
        }

        public Builder loadBalancerName(@Nullable Output<String> loadBalancerName) {
            $.loadBalancerName = loadBalancerName;
            return this;
        }

        public Builder loadBalancerName(String loadBalancerName) {
            return loadBalancerName(Output.of(loadBalancerName));
        }

        public Builder masterZoneId(@Nullable Output<String> masterZoneId) {
            $.masterZoneId = masterZoneId;
            return this;
        }

        public Builder masterZoneId(String masterZoneId) {
            return masterZoneId(Output.of(masterZoneId));
        }

        /**
         * @param nameRegex A regex string to filter results by SLB name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by SLB name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param networkType Network type of the SLBs. Valid values: `vpc` and `classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable Output<String> networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param networkType Network type of the SLBs. Valid values: `vpc` and `classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(String networkType) {
            return networkType(Output.of(networkType));
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

        public Builder pageNumber(@Nullable Output<Integer> pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageNumber(Integer pageNumber) {
            return pageNumber(Output.of(pageNumber));
        }

        public Builder pageSize(@Nullable Output<Integer> pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        public Builder pageSize(Integer pageSize) {
            return pageSize(Output.of(pageSize));
        }

        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param resourceGroupId The Id of resource group which SLB belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which SLB belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder serverId(@Nullable Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        public Builder serverIntranetAddress(@Nullable Output<String> serverIntranetAddress) {
            $.serverIntranetAddress = serverIntranetAddress;
            return this;
        }

        public Builder serverIntranetAddress(String serverIntranetAddress) {
            return serverIntranetAddress(Output.of(serverIntranetAddress));
        }

        public Builder slaveZoneId(@Nullable Output<String> slaveZoneId) {
            $.slaveZoneId = slaveZoneId;
            return this;
        }

        public Builder slaveZoneId(String slaveZoneId) {
            return slaveZoneId(Output.of(slaveZoneId));
        }

        /**
         * @param status SLB current status. Possible values: `inactive`, `active` and `locked`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status SLB current status. Possible values: `inactive`, `active` and `locked`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
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
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
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
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId ID of the VPC linked to the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId ID of the VPC linked to the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId ID of the VSwitch linked to the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId ID of the VSwitch linked to the SLBs.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public GetLoadBalancersArgs build() {
            return $;
        }
    }

}
