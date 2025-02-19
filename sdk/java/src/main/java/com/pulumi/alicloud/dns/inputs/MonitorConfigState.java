// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.inputs;

import com.pulumi.alicloud.dns.inputs.MonitorConfigIspCityNodeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MonitorConfigState extends com.pulumi.resources.ResourceArgs {

    public static final MonitorConfigState Empty = new MonitorConfigState();

    /**
     * The ID of the address pool.
     * 
     */
    @Import(name="addrPoolId")
    private @Nullable Output<String> addrPoolId;

    /**
     * @return The ID of the address pool.
     * 
     */
    public Optional<Output<String>> addrPoolId() {
        return Optional.ofNullable(this.addrPoolId);
    }

    /**
     * The number of consecutive times of failed health check attempts. Valid values: `1`, `2`, `3`.
     * 
     */
    @Import(name="evaluationCount")
    private @Nullable Output<Integer> evaluationCount;

    /**
     * @return The number of consecutive times of failed health check attempts. Valid values: `1`, `2`, `3`.
     * 
     */
    public Optional<Output<Integer>> evaluationCount() {
        return Optional.ofNullable(this.evaluationCount);
    }

    /**
     * The health check interval. Unit: seconds. Valid values: `60`.
     * 
     */
    @Import(name="interval")
    private @Nullable Output<Integer> interval;

    /**
     * @return The health check interval. Unit: seconds. Valid values: `60`.
     * 
     */
    public Optional<Output<Integer>> interval() {
        return Optional.ofNullable(this.interval);
    }

    /**
     * The Monitoring node. See the following `Block isp_city_node`.
     * 
     */
    @Import(name="ispCityNodes")
    private @Nullable Output<List<MonitorConfigIspCityNodeArgs>> ispCityNodes;

    /**
     * @return The Monitoring node. See the following `Block isp_city_node`.
     * 
     */
    public Optional<Output<List<MonitorConfigIspCityNodeArgs>>> ispCityNodes() {
        return Optional.ofNullable(this.ispCityNodes);
    }

    /**
     * The lang.
     * 
     */
    @Import(name="lang")
    private @Nullable Output<String> lang;

    /**
     * @return The lang.
     * 
     */
    public Optional<Output<String>> lang() {
        return Optional.ofNullable(this.lang);
    }

    /**
     * The extended information. This value follows the json format. For more details, see the [description of MonitorExtendInfo in the Request parameters table for details](https://www.alibabacloud.com/help/en/doc-detail/198064.html).
     * 
     */
    @Import(name="monitorExtendInfo")
    private @Nullable Output<String> monitorExtendInfo;

    /**
     * @return The extended information. This value follows the json format. For more details, see the [description of MonitorExtendInfo in the Request parameters table for details](https://www.alibabacloud.com/help/en/doc-detail/198064.html).
     * 
     */
    public Optional<Output<String>> monitorExtendInfo() {
        return Optional.ofNullable(this.monitorExtendInfo);
    }

    /**
     * The health check protocol. Valid values: `HTTP`, `HTTPS`, `PING`, `TCP`.
     * 
     */
    @Import(name="protocolType")
    private @Nullable Output<String> protocolType;

    /**
     * @return The health check protocol. Valid values: `HTTP`, `HTTPS`, `PING`, `TCP`.
     * 
     */
    public Optional<Output<String>> protocolType() {
        return Optional.ofNullable(this.protocolType);
    }

    /**
     * The timeout period. Unit: milliseconds. Valid values: `2000`, `3000`, `5000`, `10000`.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return The timeout period. Unit: milliseconds. Valid values: `2000`, `3000`, `5000`, `10000`.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private MonitorConfigState() {}

    private MonitorConfigState(MonitorConfigState $) {
        this.addrPoolId = $.addrPoolId;
        this.evaluationCount = $.evaluationCount;
        this.interval = $.interval;
        this.ispCityNodes = $.ispCityNodes;
        this.lang = $.lang;
        this.monitorExtendInfo = $.monitorExtendInfo;
        this.protocolType = $.protocolType;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MonitorConfigState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MonitorConfigState $;

        public Builder() {
            $ = new MonitorConfigState();
        }

        public Builder(MonitorConfigState defaults) {
            $ = new MonitorConfigState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addrPoolId The ID of the address pool.
         * 
         * @return builder
         * 
         */
        public Builder addrPoolId(@Nullable Output<String> addrPoolId) {
            $.addrPoolId = addrPoolId;
            return this;
        }

        /**
         * @param addrPoolId The ID of the address pool.
         * 
         * @return builder
         * 
         */
        public Builder addrPoolId(String addrPoolId) {
            return addrPoolId(Output.of(addrPoolId));
        }

        /**
         * @param evaluationCount The number of consecutive times of failed health check attempts. Valid values: `1`, `2`, `3`.
         * 
         * @return builder
         * 
         */
        public Builder evaluationCount(@Nullable Output<Integer> evaluationCount) {
            $.evaluationCount = evaluationCount;
            return this;
        }

        /**
         * @param evaluationCount The number of consecutive times of failed health check attempts. Valid values: `1`, `2`, `3`.
         * 
         * @return builder
         * 
         */
        public Builder evaluationCount(Integer evaluationCount) {
            return evaluationCount(Output.of(evaluationCount));
        }

        /**
         * @param interval The health check interval. Unit: seconds. Valid values: `60`.
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval The health check interval. Unit: seconds. Valid values: `60`.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param ispCityNodes The Monitoring node. See the following `Block isp_city_node`.
         * 
         * @return builder
         * 
         */
        public Builder ispCityNodes(@Nullable Output<List<MonitorConfigIspCityNodeArgs>> ispCityNodes) {
            $.ispCityNodes = ispCityNodes;
            return this;
        }

        /**
         * @param ispCityNodes The Monitoring node. See the following `Block isp_city_node`.
         * 
         * @return builder
         * 
         */
        public Builder ispCityNodes(List<MonitorConfigIspCityNodeArgs> ispCityNodes) {
            return ispCityNodes(Output.of(ispCityNodes));
        }

        /**
         * @param ispCityNodes The Monitoring node. See the following `Block isp_city_node`.
         * 
         * @return builder
         * 
         */
        public Builder ispCityNodes(MonitorConfigIspCityNodeArgs... ispCityNodes) {
            return ispCityNodes(List.of(ispCityNodes));
        }

        /**
         * @param lang The lang.
         * 
         * @return builder
         * 
         */
        public Builder lang(@Nullable Output<String> lang) {
            $.lang = lang;
            return this;
        }

        /**
         * @param lang The lang.
         * 
         * @return builder
         * 
         */
        public Builder lang(String lang) {
            return lang(Output.of(lang));
        }

        /**
         * @param monitorExtendInfo The extended information. This value follows the json format. For more details, see the [description of MonitorExtendInfo in the Request parameters table for details](https://www.alibabacloud.com/help/en/doc-detail/198064.html).
         * 
         * @return builder
         * 
         */
        public Builder monitorExtendInfo(@Nullable Output<String> monitorExtendInfo) {
            $.monitorExtendInfo = monitorExtendInfo;
            return this;
        }

        /**
         * @param monitorExtendInfo The extended information. This value follows the json format. For more details, see the [description of MonitorExtendInfo in the Request parameters table for details](https://www.alibabacloud.com/help/en/doc-detail/198064.html).
         * 
         * @return builder
         * 
         */
        public Builder monitorExtendInfo(String monitorExtendInfo) {
            return monitorExtendInfo(Output.of(monitorExtendInfo));
        }

        /**
         * @param protocolType The health check protocol. Valid values: `HTTP`, `HTTPS`, `PING`, `TCP`.
         * 
         * @return builder
         * 
         */
        public Builder protocolType(@Nullable Output<String> protocolType) {
            $.protocolType = protocolType;
            return this;
        }

        /**
         * @param protocolType The health check protocol. Valid values: `HTTP`, `HTTPS`, `PING`, `TCP`.
         * 
         * @return builder
         * 
         */
        public Builder protocolType(String protocolType) {
            return protocolType(Output.of(protocolType));
        }

        /**
         * @param timeout The timeout period. Unit: milliseconds. Valid values: `2000`, `3000`, `5000`, `10000`.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout The timeout period. Unit: milliseconds. Valid values: `2000`, `3000`, `5000`, `10000`.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public MonitorConfigState build() {
            return $;
        }
    }

}
