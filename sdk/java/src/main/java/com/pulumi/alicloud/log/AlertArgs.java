// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.log.inputs.AlertAnnotationArgs;
import com.pulumi.alicloud.log.inputs.AlertGroupConfigurationArgs;
import com.pulumi.alicloud.log.inputs.AlertJoinConfigurationArgs;
import com.pulumi.alicloud.log.inputs.AlertLabelArgs;
import com.pulumi.alicloud.log.inputs.AlertNotificationListArgs;
import com.pulumi.alicloud.log.inputs.AlertPolicyConfigurationArgs;
import com.pulumi.alicloud.log.inputs.AlertQueryListArgs;
import com.pulumi.alicloud.log.inputs.AlertScheduleArgs;
import com.pulumi.alicloud.log.inputs.AlertSeverityConfigurationArgs;
import com.pulumi.alicloud.log.inputs.AlertTemplateConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlertArgs Empty = new AlertArgs();

    /**
     * Alert description.
     * 
     */
    @Import(name="alertDescription")
    private @Nullable Output<String> alertDescription;

    /**
     * @return Alert description.
     * 
     */
    public Optional<Output<String>> alertDescription() {
        return Optional.ofNullable(this.alertDescription);
    }

    /**
     * Alert displayname.
     * 
     */
    @Import(name="alertDisplayname", required=true)
    private Output<String> alertDisplayname;

    /**
     * @return Alert displayname.
     * 
     */
    public Output<String> alertDisplayname() {
        return this.alertDisplayname;
    }

    /**
     * Name of logstore for configuring alarm service.
     * 
     */
    @Import(name="alertName", required=true)
    private Output<String> alertName;

    /**
     * @return Name of logstore for configuring alarm service.
     * 
     */
    public Output<String> alertName() {
        return this.alertName;
    }

    /**
     * Alert template annotations.
     * 
     */
    @Import(name="annotations")
    private @Nullable Output<List<AlertAnnotationArgs>> annotations;

    /**
     * @return Alert template annotations.
     * 
     */
    public Optional<Output<List<AlertAnnotationArgs>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    /**
     * whether to add automatic annotation, default is false.
     * 
     */
    @Import(name="autoAnnotation")
    private @Nullable Output<Boolean> autoAnnotation;

    /**
     * @return whether to add automatic annotation, default is false.
     * 
     */
    public Optional<Output<Boolean>> autoAnnotation() {
        return Optional.ofNullable(this.autoAnnotation);
    }

    /**
     * Join condition.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use eval_condition in severity_configurations
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use eval_condition in severity_configurations */
    @Import(name="condition")
    private @Nullable Output<String> condition;

    /**
     * @return Join condition.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use eval_condition in severity_configurations
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use eval_condition in severity_configurations */
    public Optional<Output<String>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * @deprecated
     * Deprecated from 1.161.0+, use dashboardId in query_list
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use dashboardId in query_list */
    @Import(name="dashboard")
    private @Nullable Output<String> dashboard;

    /**
     * @deprecated
     * Deprecated from 1.161.0+, use dashboardId in query_list
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use dashboardId in query_list */
    public Optional<Output<String>> dashboard() {
        return Optional.ofNullable(this.dashboard);
    }

    /**
     * Group configuration for new alert.
     * 
     */
    @Import(name="groupConfiguration")
    private @Nullable Output<AlertGroupConfigurationArgs> groupConfiguration;

    /**
     * @return Group configuration for new alert.
     * 
     */
    public Optional<Output<AlertGroupConfigurationArgs>> groupConfiguration() {
        return Optional.ofNullable(this.groupConfiguration);
    }

    /**
     * Join configuration for different queries.
     * 
     */
    @Import(name="joinConfigurations")
    private @Nullable Output<List<AlertJoinConfigurationArgs>> joinConfigurations;

    /**
     * @return Join configuration for different queries.
     * 
     */
    public Optional<Output<List<AlertJoinConfigurationArgs>>> joinConfigurations() {
        return Optional.ofNullable(this.joinConfigurations);
    }

    /**
     * Labels for new alert.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<AlertLabelArgs>> labels;

    /**
     * @return Labels for new alert.
     * 
     */
    public Optional<Output<List<AlertLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Timestamp, notifications before closing again.
     * 
     */
    @Import(name="muteUntil")
    private @Nullable Output<Integer> muteUntil;

    /**
     * @return Timestamp, notifications before closing again.
     * 
     */
    public Optional<Output<Integer>> muteUntil() {
        return Optional.ofNullable(this.muteUntil);
    }

    /**
     * Switch for whether new alert fires when no data happens, default is false.
     * 
     */
    @Import(name="noDataFire")
    private @Nullable Output<Boolean> noDataFire;

    /**
     * @return Switch for whether new alert fires when no data happens, default is false.
     * 
     */
    public Optional<Output<Boolean>> noDataFire() {
        return Optional.ofNullable(this.noDataFire);
    }

    /**
     * when no data happens, the severity of new alert.
     * 
     */
    @Import(name="noDataSeverity")
    private @Nullable Output<Integer> noDataSeverity;

    /**
     * @return when no data happens, the severity of new alert.
     * 
     */
    public Optional<Output<Integer>> noDataSeverity() {
        return Optional.ofNullable(this.noDataSeverity);
    }

    /**
     * Alarm information notification list, Deprecated from 1.161.0+.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use policy_configuration for notification
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use policy_configuration for notification */
    @Import(name="notificationLists")
    private @Nullable Output<List<AlertNotificationListArgs>> notificationLists;

    /**
     * @return Alarm information notification list, Deprecated from 1.161.0+.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use policy_configuration for notification
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use policy_configuration for notification */
    public Optional<Output<List<AlertNotificationListArgs>>> notificationLists() {
        return Optional.ofNullable(this.notificationLists);
    }

    /**
     * Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use threshold
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use threshold */
    @Import(name="notifyThreshold")
    private @Nullable Output<Integer> notifyThreshold;

    /**
     * @return Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use threshold
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use threshold */
    public Optional<Output<Integer>> notifyThreshold() {
        return Optional.ofNullable(this.notifyThreshold);
    }

    /**
     * Policy configuration for new alert.
     * 
     */
    @Import(name="policyConfiguration")
    private @Nullable Output<AlertPolicyConfigurationArgs> policyConfiguration;

    /**
     * @return Policy configuration for new alert.
     * 
     */
    public Optional<Output<AlertPolicyConfigurationArgs>> policyConfiguration() {
        return Optional.ofNullable(this.policyConfiguration);
    }

    /**
     * The project name.
     * 
     */
    @Import(name="projectName", required=true)
    private Output<String> projectName;

    /**
     * @return The project name.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }

    /**
     * Multiple conditions for configured alarm query.
     * 
     */
    @Import(name="queryLists")
    private @Nullable Output<List<AlertQueryListArgs>> queryLists;

    /**
     * @return Multiple conditions for configured alarm query.
     * 
     */
    public Optional<Output<List<AlertQueryListArgs>>> queryLists() {
        return Optional.ofNullable(this.queryLists);
    }

    /**
     * schedule for alert.
     * 
     */
    @Import(name="schedule")
    private @Nullable Output<AlertScheduleArgs> schedule;

    /**
     * @return schedule for alert.
     * 
     */
    public Optional<Output<AlertScheduleArgs>> schedule() {
        return Optional.ofNullable(this.schedule);
    }

    /**
     * Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
     * 
     * @deprecated
     * Field &#39;schedule_interval&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
     * 
     */
    @Deprecated /* Field 'schedule_interval' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
    @Import(name="scheduleInterval")
    private @Nullable Output<String> scheduleInterval;

    /**
     * @return Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
     * 
     * @deprecated
     * Field &#39;schedule_interval&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
     * 
     */
    @Deprecated /* Field 'schedule_interval' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
    public Optional<Output<String>> scheduleInterval() {
        return Optional.ofNullable(this.scheduleInterval);
    }

    /**
     * Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
     * 
     * @deprecated
     * Field &#39;schedule_type&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
     * 
     */
    @Deprecated /* Field 'schedule_type' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
    @Import(name="scheduleType")
    private @Nullable Output<String> scheduleType;

    /**
     * @return Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
     * 
     * @deprecated
     * Field &#39;schedule_type&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
     * 
     */
    @Deprecated /* Field 'schedule_type' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
    public Optional<Output<String>> scheduleType() {
        return Optional.ofNullable(this.scheduleType);
    }

    /**
     * when new alert is resolved, whether to notify, default is false.
     * 
     */
    @Import(name="sendResolved")
    private @Nullable Output<Boolean> sendResolved;

    /**
     * @return when new alert is resolved, whether to notify, default is false.
     * 
     */
    public Optional<Output<Boolean>> sendResolved() {
        return Optional.ofNullable(this.sendResolved);
    }

    /**
     * Severity configuration for new alert.
     * 
     */
    @Import(name="severityConfigurations")
    private @Nullable Output<List<AlertSeverityConfigurationArgs>> severityConfigurations;

    /**
     * @return Severity configuration for new alert.
     * 
     */
    public Optional<Output<List<AlertSeverityConfigurationArgs>>> severityConfigurations() {
        return Optional.ofNullable(this.severityConfigurations);
    }

    /**
     * Template configuration for alert, when `type` is `tpl`.
     * 
     */
    @Import(name="templateConfiguration")
    private @Nullable Output<AlertTemplateConfigurationArgs> templateConfiguration;

    /**
     * @return Template configuration for alert, when `type` is `tpl`.
     * 
     */
    public Optional<Output<AlertTemplateConfigurationArgs>> templateConfiguration() {
        return Optional.ofNullable(this.templateConfiguration);
    }

    /**
     * Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
     * 
     */
    @Import(name="threshold")
    private @Nullable Output<Integer> threshold;

    /**
     * @return Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
     * 
     */
    public Optional<Output<Integer>> threshold() {
        return Optional.ofNullable(this.threshold);
    }

    /**
     * Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use repeat_interval in policy_configuration
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use repeat_interval in policy_configuration */
    @Import(name="throttling")
    private @Nullable Output<String> throttling;

    /**
     * @return Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
     * 
     * @deprecated
     * Deprecated from 1.161.0+, use repeat_interval in policy_configuration
     * 
     */
    @Deprecated /* Deprecated from 1.161.0+, use repeat_interval in policy_configuration */
    public Optional<Output<String>> throttling() {
        return Optional.ofNullable(this.throttling);
    }

    /**
     * including FixedRate,Hourly,Daily,Weekly,Cron.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return including FixedRate,Hourly,Daily,Weekly,Cron.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The version of alert, new alert is 2.0.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version of alert, new alert is 2.0.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private AlertArgs() {}

    private AlertArgs(AlertArgs $) {
        this.alertDescription = $.alertDescription;
        this.alertDisplayname = $.alertDisplayname;
        this.alertName = $.alertName;
        this.annotations = $.annotations;
        this.autoAnnotation = $.autoAnnotation;
        this.condition = $.condition;
        this.dashboard = $.dashboard;
        this.groupConfiguration = $.groupConfiguration;
        this.joinConfigurations = $.joinConfigurations;
        this.labels = $.labels;
        this.muteUntil = $.muteUntil;
        this.noDataFire = $.noDataFire;
        this.noDataSeverity = $.noDataSeverity;
        this.notificationLists = $.notificationLists;
        this.notifyThreshold = $.notifyThreshold;
        this.policyConfiguration = $.policyConfiguration;
        this.projectName = $.projectName;
        this.queryLists = $.queryLists;
        this.schedule = $.schedule;
        this.scheduleInterval = $.scheduleInterval;
        this.scheduleType = $.scheduleType;
        this.sendResolved = $.sendResolved;
        this.severityConfigurations = $.severityConfigurations;
        this.templateConfiguration = $.templateConfiguration;
        this.threshold = $.threshold;
        this.throttling = $.throttling;
        this.type = $.type;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertArgs $;

        public Builder() {
            $ = new AlertArgs();
        }

        public Builder(AlertArgs defaults) {
            $ = new AlertArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertDescription Alert description.
         * 
         * @return builder
         * 
         */
        public Builder alertDescription(@Nullable Output<String> alertDescription) {
            $.alertDescription = alertDescription;
            return this;
        }

        /**
         * @param alertDescription Alert description.
         * 
         * @return builder
         * 
         */
        public Builder alertDescription(String alertDescription) {
            return alertDescription(Output.of(alertDescription));
        }

        /**
         * @param alertDisplayname Alert displayname.
         * 
         * @return builder
         * 
         */
        public Builder alertDisplayname(Output<String> alertDisplayname) {
            $.alertDisplayname = alertDisplayname;
            return this;
        }

        /**
         * @param alertDisplayname Alert displayname.
         * 
         * @return builder
         * 
         */
        public Builder alertDisplayname(String alertDisplayname) {
            return alertDisplayname(Output.of(alertDisplayname));
        }

        /**
         * @param alertName Name of logstore for configuring alarm service.
         * 
         * @return builder
         * 
         */
        public Builder alertName(Output<String> alertName) {
            $.alertName = alertName;
            return this;
        }

        /**
         * @param alertName Name of logstore for configuring alarm service.
         * 
         * @return builder
         * 
         */
        public Builder alertName(String alertName) {
            return alertName(Output.of(alertName));
        }

        /**
         * @param annotations Alert template annotations.
         * 
         * @return builder
         * 
         */
        public Builder annotations(@Nullable Output<List<AlertAnnotationArgs>> annotations) {
            $.annotations = annotations;
            return this;
        }

        /**
         * @param annotations Alert template annotations.
         * 
         * @return builder
         * 
         */
        public Builder annotations(List<AlertAnnotationArgs> annotations) {
            return annotations(Output.of(annotations));
        }

        /**
         * @param annotations Alert template annotations.
         * 
         * @return builder
         * 
         */
        public Builder annotations(AlertAnnotationArgs... annotations) {
            return annotations(List.of(annotations));
        }

        /**
         * @param autoAnnotation whether to add automatic annotation, default is false.
         * 
         * @return builder
         * 
         */
        public Builder autoAnnotation(@Nullable Output<Boolean> autoAnnotation) {
            $.autoAnnotation = autoAnnotation;
            return this;
        }

        /**
         * @param autoAnnotation whether to add automatic annotation, default is false.
         * 
         * @return builder
         * 
         */
        public Builder autoAnnotation(Boolean autoAnnotation) {
            return autoAnnotation(Output.of(autoAnnotation));
        }

        /**
         * @param condition Join condition.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use eval_condition in severity_configurations
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use eval_condition in severity_configurations */
        public Builder condition(@Nullable Output<String> condition) {
            $.condition = condition;
            return this;
        }

        /**
         * @param condition Join condition.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use eval_condition in severity_configurations
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use eval_condition in severity_configurations */
        public Builder condition(String condition) {
            return condition(Output.of(condition));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use dashboardId in query_list
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use dashboardId in query_list */
        public Builder dashboard(@Nullable Output<String> dashboard) {
            $.dashboard = dashboard;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use dashboardId in query_list
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use dashboardId in query_list */
        public Builder dashboard(String dashboard) {
            return dashboard(Output.of(dashboard));
        }

        /**
         * @param groupConfiguration Group configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder groupConfiguration(@Nullable Output<AlertGroupConfigurationArgs> groupConfiguration) {
            $.groupConfiguration = groupConfiguration;
            return this;
        }

        /**
         * @param groupConfiguration Group configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder groupConfiguration(AlertGroupConfigurationArgs groupConfiguration) {
            return groupConfiguration(Output.of(groupConfiguration));
        }

        /**
         * @param joinConfigurations Join configuration for different queries.
         * 
         * @return builder
         * 
         */
        public Builder joinConfigurations(@Nullable Output<List<AlertJoinConfigurationArgs>> joinConfigurations) {
            $.joinConfigurations = joinConfigurations;
            return this;
        }

        /**
         * @param joinConfigurations Join configuration for different queries.
         * 
         * @return builder
         * 
         */
        public Builder joinConfigurations(List<AlertJoinConfigurationArgs> joinConfigurations) {
            return joinConfigurations(Output.of(joinConfigurations));
        }

        /**
         * @param joinConfigurations Join configuration for different queries.
         * 
         * @return builder
         * 
         */
        public Builder joinConfigurations(AlertJoinConfigurationArgs... joinConfigurations) {
            return joinConfigurations(List.of(joinConfigurations));
        }

        /**
         * @param labels Labels for new alert.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<AlertLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Labels for new alert.
         * 
         * @return builder
         * 
         */
        public Builder labels(List<AlertLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels Labels for new alert.
         * 
         * @return builder
         * 
         */
        public Builder labels(AlertLabelArgs... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param muteUntil Timestamp, notifications before closing again.
         * 
         * @return builder
         * 
         */
        public Builder muteUntil(@Nullable Output<Integer> muteUntil) {
            $.muteUntil = muteUntil;
            return this;
        }

        /**
         * @param muteUntil Timestamp, notifications before closing again.
         * 
         * @return builder
         * 
         */
        public Builder muteUntil(Integer muteUntil) {
            return muteUntil(Output.of(muteUntil));
        }

        /**
         * @param noDataFire Switch for whether new alert fires when no data happens, default is false.
         * 
         * @return builder
         * 
         */
        public Builder noDataFire(@Nullable Output<Boolean> noDataFire) {
            $.noDataFire = noDataFire;
            return this;
        }

        /**
         * @param noDataFire Switch for whether new alert fires when no data happens, default is false.
         * 
         * @return builder
         * 
         */
        public Builder noDataFire(Boolean noDataFire) {
            return noDataFire(Output.of(noDataFire));
        }

        /**
         * @param noDataSeverity when no data happens, the severity of new alert.
         * 
         * @return builder
         * 
         */
        public Builder noDataSeverity(@Nullable Output<Integer> noDataSeverity) {
            $.noDataSeverity = noDataSeverity;
            return this;
        }

        /**
         * @param noDataSeverity when no data happens, the severity of new alert.
         * 
         * @return builder
         * 
         */
        public Builder noDataSeverity(Integer noDataSeverity) {
            return noDataSeverity(Output.of(noDataSeverity));
        }

        /**
         * @param notificationLists Alarm information notification list, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use policy_configuration for notification
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use policy_configuration for notification */
        public Builder notificationLists(@Nullable Output<List<AlertNotificationListArgs>> notificationLists) {
            $.notificationLists = notificationLists;
            return this;
        }

        /**
         * @param notificationLists Alarm information notification list, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use policy_configuration for notification
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use policy_configuration for notification */
        public Builder notificationLists(List<AlertNotificationListArgs> notificationLists) {
            return notificationLists(Output.of(notificationLists));
        }

        /**
         * @param notificationLists Alarm information notification list, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use policy_configuration for notification
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use policy_configuration for notification */
        public Builder notificationLists(AlertNotificationListArgs... notificationLists) {
            return notificationLists(List.of(notificationLists));
        }

        /**
         * @param notifyThreshold Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use threshold
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use threshold */
        public Builder notifyThreshold(@Nullable Output<Integer> notifyThreshold) {
            $.notifyThreshold = notifyThreshold;
            return this;
        }

        /**
         * @param notifyThreshold Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use threshold
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use threshold */
        public Builder notifyThreshold(Integer notifyThreshold) {
            return notifyThreshold(Output.of(notifyThreshold));
        }

        /**
         * @param policyConfiguration Policy configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder policyConfiguration(@Nullable Output<AlertPolicyConfigurationArgs> policyConfiguration) {
            $.policyConfiguration = policyConfiguration;
            return this;
        }

        /**
         * @param policyConfiguration Policy configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder policyConfiguration(AlertPolicyConfigurationArgs policyConfiguration) {
            return policyConfiguration(Output.of(policyConfiguration));
        }

        /**
         * @param projectName The project name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName The project name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param queryLists Multiple conditions for configured alarm query.
         * 
         * @return builder
         * 
         */
        public Builder queryLists(@Nullable Output<List<AlertQueryListArgs>> queryLists) {
            $.queryLists = queryLists;
            return this;
        }

        /**
         * @param queryLists Multiple conditions for configured alarm query.
         * 
         * @return builder
         * 
         */
        public Builder queryLists(List<AlertQueryListArgs> queryLists) {
            return queryLists(Output.of(queryLists));
        }

        /**
         * @param queryLists Multiple conditions for configured alarm query.
         * 
         * @return builder
         * 
         */
        public Builder queryLists(AlertQueryListArgs... queryLists) {
            return queryLists(List.of(queryLists));
        }

        /**
         * @param schedule schedule for alert.
         * 
         * @return builder
         * 
         */
        public Builder schedule(@Nullable Output<AlertScheduleArgs> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule schedule for alert.
         * 
         * @return builder
         * 
         */
        public Builder schedule(AlertScheduleArgs schedule) {
            return schedule(Output.of(schedule));
        }

        /**
         * @param scheduleInterval Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;schedule_interval&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
         * 
         */
        @Deprecated /* Field 'schedule_interval' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
        public Builder scheduleInterval(@Nullable Output<String> scheduleInterval) {
            $.scheduleInterval = scheduleInterval;
            return this;
        }

        /**
         * @param scheduleInterval Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;schedule_interval&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
         * 
         */
        @Deprecated /* Field 'schedule_interval' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
        public Builder scheduleInterval(String scheduleInterval) {
            return scheduleInterval(Output.of(scheduleInterval));
        }

        /**
         * @param scheduleType Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;schedule_type&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
         * 
         */
        @Deprecated /* Field 'schedule_type' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
        public Builder scheduleType(@Nullable Output<String> scheduleType) {
            $.scheduleType = scheduleType;
            return this;
        }

        /**
         * @param scheduleType Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;schedule_type&#39; has been deprecated from provider version 1.176.0. New field &#39;schedule&#39; instead.
         * 
         */
        @Deprecated /* Field 'schedule_type' has been deprecated from provider version 1.176.0. New field 'schedule' instead. */
        public Builder scheduleType(String scheduleType) {
            return scheduleType(Output.of(scheduleType));
        }

        /**
         * @param sendResolved when new alert is resolved, whether to notify, default is false.
         * 
         * @return builder
         * 
         */
        public Builder sendResolved(@Nullable Output<Boolean> sendResolved) {
            $.sendResolved = sendResolved;
            return this;
        }

        /**
         * @param sendResolved when new alert is resolved, whether to notify, default is false.
         * 
         * @return builder
         * 
         */
        public Builder sendResolved(Boolean sendResolved) {
            return sendResolved(Output.of(sendResolved));
        }

        /**
         * @param severityConfigurations Severity configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder severityConfigurations(@Nullable Output<List<AlertSeverityConfigurationArgs>> severityConfigurations) {
            $.severityConfigurations = severityConfigurations;
            return this;
        }

        /**
         * @param severityConfigurations Severity configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder severityConfigurations(List<AlertSeverityConfigurationArgs> severityConfigurations) {
            return severityConfigurations(Output.of(severityConfigurations));
        }

        /**
         * @param severityConfigurations Severity configuration for new alert.
         * 
         * @return builder
         * 
         */
        public Builder severityConfigurations(AlertSeverityConfigurationArgs... severityConfigurations) {
            return severityConfigurations(List.of(severityConfigurations));
        }

        /**
         * @param templateConfiguration Template configuration for alert, when `type` is `tpl`.
         * 
         * @return builder
         * 
         */
        public Builder templateConfiguration(@Nullable Output<AlertTemplateConfigurationArgs> templateConfiguration) {
            $.templateConfiguration = templateConfiguration;
            return this;
        }

        /**
         * @param templateConfiguration Template configuration for alert, when `type` is `tpl`.
         * 
         * @return builder
         * 
         */
        public Builder templateConfiguration(AlertTemplateConfigurationArgs templateConfiguration) {
            return templateConfiguration(Output.of(templateConfiguration));
        }

        /**
         * @param threshold Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
         * 
         * @return builder
         * 
         */
        public Builder threshold(@Nullable Output<Integer> threshold) {
            $.threshold = threshold;
            return this;
        }

        /**
         * @param threshold Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
         * 
         * @return builder
         * 
         */
        public Builder threshold(Integer threshold) {
            return threshold(Output.of(threshold));
        }

        /**
         * @param throttling Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use repeat_interval in policy_configuration
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use repeat_interval in policy_configuration */
        public Builder throttling(@Nullable Output<String> throttling) {
            $.throttling = throttling;
            return this;
        }

        /**
         * @param throttling Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated from 1.161.0+, use repeat_interval in policy_configuration
         * 
         */
        @Deprecated /* Deprecated from 1.161.0+, use repeat_interval in policy_configuration */
        public Builder throttling(String throttling) {
            return throttling(Output.of(throttling));
        }

        /**
         * @param type including FixedRate,Hourly,Daily,Weekly,Cron.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type including FixedRate,Hourly,Daily,Weekly,Cron.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param version The version of alert, new alert is 2.0.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of alert, new alert is 2.0.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public AlertArgs build() {
            $.alertDisplayname = Objects.requireNonNull($.alertDisplayname, "expected parameter 'alertDisplayname' to be non-null");
            $.alertName = Objects.requireNonNull($.alertName, "expected parameter 'alertName' to be non-null");
            $.projectName = Objects.requireNonNull($.projectName, "expected parameter 'projectName' to be non-null");
            return $;
        }
    }

}
