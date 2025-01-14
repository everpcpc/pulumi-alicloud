// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.databasefilesystem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class AutoSnapShotPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final AutoSnapShotPolicyArgs Empty = new AutoSnapShotPolicyArgs();

    /**
     * Automatic snapshot policy name
     * 
     */
    @Import(name="policyName", required=true)
    private Output<String> policyName;

    /**
     * @return Automatic snapshot policy name
     * 
     */
    public Output<String> policyName() {
        return this.policyName;
    }

    /**
     * A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
     * 
     */
    @Import(name="repeatWeekdays", required=true)
    private Output<List<String>> repeatWeekdays;

    /**
     * @return A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
     * 
     */
    public Output<List<String>> repeatWeekdays() {
        return this.repeatWeekdays;
    }

    /**
     * Automatic snapshot retention days.
     * 
     */
    @Import(name="retentionDays", required=true)
    private Output<Integer> retentionDays;

    /**
     * @return Automatic snapshot retention days.
     * 
     */
    public Output<Integer> retentionDays() {
        return this.retentionDays;
    }

    /**
     * The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
     * 
     */
    @Import(name="timePoints", required=true)
    private Output<List<String>> timePoints;

    /**
     * @return The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
     * 
     */
    public Output<List<String>> timePoints() {
        return this.timePoints;
    }

    private AutoSnapShotPolicyArgs() {}

    private AutoSnapShotPolicyArgs(AutoSnapShotPolicyArgs $) {
        this.policyName = $.policyName;
        this.repeatWeekdays = $.repeatWeekdays;
        this.retentionDays = $.retentionDays;
        this.timePoints = $.timePoints;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AutoSnapShotPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AutoSnapShotPolicyArgs $;

        public Builder() {
            $ = new AutoSnapShotPolicyArgs();
        }

        public Builder(AutoSnapShotPolicyArgs defaults) {
            $ = new AutoSnapShotPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param policyName Automatic snapshot policy name
         * 
         * @return builder
         * 
         */
        public Builder policyName(Output<String> policyName) {
            $.policyName = policyName;
            return this;
        }

        /**
         * @param policyName Automatic snapshot policy name
         * 
         * @return builder
         * 
         */
        public Builder policyName(String policyName) {
            return policyName(Output.of(policyName));
        }

        /**
         * @param repeatWeekdays A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
         * 
         * @return builder
         * 
         */
        public Builder repeatWeekdays(Output<List<String>> repeatWeekdays) {
            $.repeatWeekdays = repeatWeekdays;
            return this;
        }

        /**
         * @param repeatWeekdays A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
         * 
         * @return builder
         * 
         */
        public Builder repeatWeekdays(List<String> repeatWeekdays) {
            return repeatWeekdays(Output.of(repeatWeekdays));
        }

        /**
         * @param repeatWeekdays A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
         * 
         * @return builder
         * 
         */
        public Builder repeatWeekdays(String... repeatWeekdays) {
            return repeatWeekdays(List.of(repeatWeekdays));
        }

        /**
         * @param retentionDays Automatic snapshot retention days.
         * 
         * @return builder
         * 
         */
        public Builder retentionDays(Output<Integer> retentionDays) {
            $.retentionDays = retentionDays;
            return this;
        }

        /**
         * @param retentionDays Automatic snapshot retention days.
         * 
         * @return builder
         * 
         */
        public Builder retentionDays(Integer retentionDays) {
            return retentionDays(Output.of(retentionDays));
        }

        /**
         * @param timePoints The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
         * 
         * @return builder
         * 
         */
        public Builder timePoints(Output<List<String>> timePoints) {
            $.timePoints = timePoints;
            return this;
        }

        /**
         * @param timePoints The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
         * 
         * @return builder
         * 
         */
        public Builder timePoints(List<String> timePoints) {
            return timePoints(Output.of(timePoints));
        }

        /**
         * @param timePoints The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
         * 
         * @return builder
         * 
         */
        public Builder timePoints(String... timePoints) {
            return timePoints(List.of(timePoints));
        }

        public AutoSnapShotPolicyArgs build() {
            $.policyName = Objects.requireNonNull($.policyName, "expected parameter 'policyName' to be non-null");
            $.repeatWeekdays = Objects.requireNonNull($.repeatWeekdays, "expected parameter 'repeatWeekdays' to be non-null");
            $.retentionDays = Objects.requireNonNull($.retentionDays, "expected parameter 'retentionDays' to be non-null");
            $.timePoints = Objects.requireNonNull($.timePoints, "expected parameter 'timePoints' to be non-null");
            return $;
        }
    }

}
