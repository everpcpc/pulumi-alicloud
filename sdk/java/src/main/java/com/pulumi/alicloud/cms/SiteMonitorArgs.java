// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.cms.inputs.SiteMonitorIspCityArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SiteMonitorArgs extends com.pulumi.resources.ResourceArgs {

    public static final SiteMonitorArgs Empty = new SiteMonitorArgs();

    /**
     * The URL or IP address monitored by the site monitoring task.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return The URL or IP address monitored by the site monitoring task.
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * The IDs of existing alert rules to be associated with the site monitoring task.
     * 
     */
    @Import(name="alertIds")
    private @Nullable Output<List<String>> alertIds;

    /**
     * @return The IDs of existing alert rules to be associated with the site monitoring task.
     * 
     */
    public Optional<Output<List<String>>> alertIds() {
        return Optional.ofNullable(this.alertIds);
    }

    /**
     * The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
     * 
     */
    @Import(name="interval")
    private @Nullable Output<Integer> interval;

    /**
     * @return The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
     * 
     */
    public Optional<Output<Integer>> interval() {
        return Optional.ofNullable(this.interval);
    }

    /**
     * The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
     * 
     */
    @Import(name="ispCities")
    private @Nullable Output<List<SiteMonitorIspCityArgs>> ispCities;

    /**
     * @return The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
     * 
     */
    public Optional<Output<List<SiteMonitorIspCityArgs>>> ispCities() {
        return Optional.ofNullable(this.ispCities);
    }

    /**
     * The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
     * 
     */
    @Import(name="optionsJson")
    private @Nullable Output<String> optionsJson;

    /**
     * @return The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
     * 
     */
    public Optional<Output<String>> optionsJson() {
        return Optional.ofNullable(this.optionsJson);
    }

    /**
     * The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
     * 
     */
    @Import(name="taskName", required=true)
    private Output<String> taskName;

    /**
     * @return The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
     * 
     */
    public Output<String> taskName() {
        return this.taskName;
    }

    /**
     * The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
     * 
     */
    @Import(name="taskType", required=true)
    private Output<String> taskType;

    /**
     * @return The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
     * 
     */
    public Output<String> taskType() {
        return this.taskType;
    }

    private SiteMonitorArgs() {}

    private SiteMonitorArgs(SiteMonitorArgs $) {
        this.address = $.address;
        this.alertIds = $.alertIds;
        this.interval = $.interval;
        this.ispCities = $.ispCities;
        this.optionsJson = $.optionsJson;
        this.taskName = $.taskName;
        this.taskType = $.taskType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SiteMonitorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SiteMonitorArgs $;

        public Builder() {
            $ = new SiteMonitorArgs();
        }

        public Builder(SiteMonitorArgs defaults) {
            $ = new SiteMonitorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The URL or IP address monitored by the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The URL or IP address monitored by the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param alertIds The IDs of existing alert rules to be associated with the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder alertIds(@Nullable Output<List<String>> alertIds) {
            $.alertIds = alertIds;
            return this;
        }

        /**
         * @param alertIds The IDs of existing alert rules to be associated with the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder alertIds(List<String> alertIds) {
            return alertIds(Output.of(alertIds));
        }

        /**
         * @param alertIds The IDs of existing alert rules to be associated with the site monitoring task.
         * 
         * @return builder
         * 
         */
        public Builder alertIds(String... alertIds) {
            return alertIds(List.of(alertIds));
        }

        /**
         * @param interval The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param ispCities The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
         * 
         * @return builder
         * 
         */
        public Builder ispCities(@Nullable Output<List<SiteMonitorIspCityArgs>> ispCities) {
            $.ispCities = ispCities;
            return this;
        }

        /**
         * @param ispCities The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
         * 
         * @return builder
         * 
         */
        public Builder ispCities(List<SiteMonitorIspCityArgs> ispCities) {
            return ispCities(Output.of(ispCities));
        }

        /**
         * @param ispCities The detection points in a JSON array. For example, `[{&#34;city&#34;:&#34;546&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;572&#34;,&#34;isp&#34;:&#34;465&#34;},{&#34;city&#34;:&#34;738&#34;,&#34;isp&#34;:&#34;465&#34;}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
         * 
         * @return builder
         * 
         */
        public Builder ispCities(SiteMonitorIspCityArgs... ispCities) {
            return ispCities(List.of(ispCities));
        }

        /**
         * @param optionsJson The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
         * 
         * @return builder
         * 
         */
        public Builder optionsJson(@Nullable Output<String> optionsJson) {
            $.optionsJson = optionsJson;
            return this;
        }

        /**
         * @param optionsJson The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
         * 
         * @return builder
         * 
         */
        public Builder optionsJson(String optionsJson) {
            return optionsJson(Output.of(optionsJson));
        }

        /**
         * @param taskName The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
         * 
         * @return builder
         * 
         */
        public Builder taskName(Output<String> taskName) {
            $.taskName = taskName;
            return this;
        }

        /**
         * @param taskName The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
         * 
         * @return builder
         * 
         */
        public Builder taskName(String taskName) {
            return taskName(Output.of(taskName));
        }

        /**
         * @param taskType The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
         * 
         * @return builder
         * 
         */
        public Builder taskType(Output<String> taskType) {
            $.taskType = taskType;
            return this;
        }

        /**
         * @param taskType The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
         * 
         * @return builder
         * 
         */
        public Builder taskType(String taskType) {
            return taskType(Output.of(taskType));
        }

        public SiteMonitorArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.taskName = Objects.requireNonNull($.taskName, "expected parameter 'taskName' to be non-null");
            $.taskType = Objects.requireNonNull($.taskType, "expected parameter 'taskType' to be non-null");
            return $;
        }
    }

}
