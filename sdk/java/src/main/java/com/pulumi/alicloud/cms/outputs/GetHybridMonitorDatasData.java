// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GetHybridMonitorDatasDataLabel;
import com.pulumi.alicloud.cms.outputs.GetHybridMonitorDatasDataValue;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetHybridMonitorDatasData {
    /**
     * @return The label of the time dimension.
     * 
     */
    private List<GetHybridMonitorDatasDataLabel> labels;
    /**
     * @return The name of the monitoring indicator.
     * 
     */
    private String metricName;
    /**
     * @return The metric values that are collected at different timestamps.
     * 
     */
    private List<GetHybridMonitorDatasDataValue> values;

    private GetHybridMonitorDatasData() {}
    /**
     * @return The label of the time dimension.
     * 
     */
    public List<GetHybridMonitorDatasDataLabel> labels() {
        return this.labels;
    }
    /**
     * @return The name of the monitoring indicator.
     * 
     */
    public String metricName() {
        return this.metricName;
    }
    /**
     * @return The metric values that are collected at different timestamps.
     * 
     */
    public List<GetHybridMonitorDatasDataValue> values() {
        return this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHybridMonitorDatasData defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetHybridMonitorDatasDataLabel> labels;
        private String metricName;
        private List<GetHybridMonitorDatasDataValue> values;
        public Builder() {}
        public Builder(GetHybridMonitorDatasData defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.labels = defaults.labels;
    	      this.metricName = defaults.metricName;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder labels(List<GetHybridMonitorDatasDataLabel> labels) {
            this.labels = Objects.requireNonNull(labels);
            return this;
        }
        public Builder labels(GetHybridMonitorDatasDataLabel... labels) {
            return labels(List.of(labels));
        }
        @CustomType.Setter
        public Builder metricName(String metricName) {
            this.metricName = Objects.requireNonNull(metricName);
            return this;
        }
        @CustomType.Setter
        public Builder values(List<GetHybridMonitorDatasDataValue> values) {
            this.values = Objects.requireNonNull(values);
            return this;
        }
        public Builder values(GetHybridMonitorDatasDataValue... values) {
            return values(List.of(values));
        }
        public GetHybridMonitorDatasData build() {
            final var o = new GetHybridMonitorDatasData();
            o.labels = labels;
            o.metricName = metricName;
            o.values = values;
            return o;
        }
    }
}
