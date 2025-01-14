// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GetMetricRuleTemplatesTemplateAlertTemplateEscalation;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMetricRuleTemplatesTemplateAlertTemplate {
    /**
     * @return The abbreviation of the service name. Valid values: `ecs`, `rds`, `ads`, `slb`, `vpc`, `apigateway`, `cdn`, `cs`, `dcdn`, `ddos`, `eip`, `elasticsearch`, `emr`, `ess`, `hbase`, `iot_edge`, `kvstore_sharding`, `kvstore_splitrw`, `kvstore_standard`, `memcache`, `mns`, `mongodb`, `mongodb_cluster`, `mongodb_sharding`, `mq_topic`, `ocs`, `opensearch`, `oss`, `polardb`, `petadata`, `scdn`, `sharebandwidthpackages`, `sls`, `vpn`.
     * 
     */
    private String category;
    /**
     * @return The information about the trigger condition based on the alert level.
     * 
     */
    private List<GetMetricRuleTemplatesTemplateAlertTemplateEscalation> escalations;
    /**
     * @return The name of the metric.
     * 
     */
    private String metricName;
    /**
     * @return The namespace of the service.
     * 
     */
    private String namespace;
    /**
     * @return The name of the alert rule.
     * 
     */
    private String ruleName;
    private String selector;
    /**
     * @return The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
     * 
     */
    private String webhook;

    private GetMetricRuleTemplatesTemplateAlertTemplate() {}
    /**
     * @return The abbreviation of the service name. Valid values: `ecs`, `rds`, `ads`, `slb`, `vpc`, `apigateway`, `cdn`, `cs`, `dcdn`, `ddos`, `eip`, `elasticsearch`, `emr`, `ess`, `hbase`, `iot_edge`, `kvstore_sharding`, `kvstore_splitrw`, `kvstore_standard`, `memcache`, `mns`, `mongodb`, `mongodb_cluster`, `mongodb_sharding`, `mq_topic`, `ocs`, `opensearch`, `oss`, `polardb`, `petadata`, `scdn`, `sharebandwidthpackages`, `sls`, `vpn`.
     * 
     */
    public String category() {
        return this.category;
    }
    /**
     * @return The information about the trigger condition based on the alert level.
     * 
     */
    public List<GetMetricRuleTemplatesTemplateAlertTemplateEscalation> escalations() {
        return this.escalations;
    }
    /**
     * @return The name of the metric.
     * 
     */
    public String metricName() {
        return this.metricName;
    }
    /**
     * @return The namespace of the service.
     * 
     */
    public String namespace() {
        return this.namespace;
    }
    /**
     * @return The name of the alert rule.
     * 
     */
    public String ruleName() {
        return this.ruleName;
    }
    public String selector() {
        return this.selector;
    }
    /**
     * @return The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
     * 
     */
    public String webhook() {
        return this.webhook;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMetricRuleTemplatesTemplateAlertTemplate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String category;
        private List<GetMetricRuleTemplatesTemplateAlertTemplateEscalation> escalations;
        private String metricName;
        private String namespace;
        private String ruleName;
        private String selector;
        private String webhook;
        public Builder() {}
        public Builder(GetMetricRuleTemplatesTemplateAlertTemplate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.escalations = defaults.escalations;
    	      this.metricName = defaults.metricName;
    	      this.namespace = defaults.namespace;
    	      this.ruleName = defaults.ruleName;
    	      this.selector = defaults.selector;
    	      this.webhook = defaults.webhook;
        }

        @CustomType.Setter
        public Builder category(String category) {
            this.category = Objects.requireNonNull(category);
            return this;
        }
        @CustomType.Setter
        public Builder escalations(List<GetMetricRuleTemplatesTemplateAlertTemplateEscalation> escalations) {
            this.escalations = Objects.requireNonNull(escalations);
            return this;
        }
        public Builder escalations(GetMetricRuleTemplatesTemplateAlertTemplateEscalation... escalations) {
            return escalations(List.of(escalations));
        }
        @CustomType.Setter
        public Builder metricName(String metricName) {
            this.metricName = Objects.requireNonNull(metricName);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        @CustomType.Setter
        public Builder ruleName(String ruleName) {
            this.ruleName = Objects.requireNonNull(ruleName);
            return this;
        }
        @CustomType.Setter
        public Builder selector(String selector) {
            this.selector = Objects.requireNonNull(selector);
            return this;
        }
        @CustomType.Setter
        public Builder webhook(String webhook) {
            this.webhook = Objects.requireNonNull(webhook);
            return this;
        }
        public GetMetricRuleTemplatesTemplateAlertTemplate build() {
            final var o = new GetMetricRuleTemplatesTemplateAlertTemplate();
            o.category = category;
            o.escalations = escalations;
            o.metricName = metricName;
            o.namespace = namespace;
            o.ruleName = ruleName;
            o.selector = selector;
            o.webhook = webhook;
            return o;
        }
    }
}
