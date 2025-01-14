// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ManagedKubernetesMaintenanceWindow {
    /**
     * @return The maintenance time, values range from 1 to 24,unit is hour. For example: &#34;3h&#34;.
     * 
     */
    private String duration;
    /**
     * @return Whether to open the maintenance window. The following parameters take effect only `enable = true`.
     * 
     */
    private Boolean enable;
    /**
     * @return Initial maintenance time, For example:&#34;03:00:00Z&#34;.
     * 
     */
    private String maintenanceTime;
    /**
     * @return Maintenance cycle, you can set the values from Monday to Sunday, separated by commas when the values are multiple. The default is Thursday.
     * 
     * for example:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
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
     *     }
     * }
     * ```
     * 
     */
    private String weeklyPeriod;

    private ManagedKubernetesMaintenanceWindow() {}
    /**
     * @return The maintenance time, values range from 1 to 24,unit is hour. For example: &#34;3h&#34;.
     * 
     */
    public String duration() {
        return this.duration;
    }
    /**
     * @return Whether to open the maintenance window. The following parameters take effect only `enable = true`.
     * 
     */
    public Boolean enable() {
        return this.enable;
    }
    /**
     * @return Initial maintenance time, For example:&#34;03:00:00Z&#34;.
     * 
     */
    public String maintenanceTime() {
        return this.maintenanceTime;
    }
    /**
     * @return Maintenance cycle, you can set the values from Monday to Sunday, separated by commas when the values are multiple. The default is Thursday.
     * 
     * for example:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
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
     *     }
     * }
     * ```
     * 
     */
    public String weeklyPeriod() {
        return this.weeklyPeriod;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ManagedKubernetesMaintenanceWindow defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String duration;
        private Boolean enable;
        private String maintenanceTime;
        private String weeklyPeriod;
        public Builder() {}
        public Builder(ManagedKubernetesMaintenanceWindow defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.duration = defaults.duration;
    	      this.enable = defaults.enable;
    	      this.maintenanceTime = defaults.maintenanceTime;
    	      this.weeklyPeriod = defaults.weeklyPeriod;
        }

        @CustomType.Setter
        public Builder duration(String duration) {
            this.duration = Objects.requireNonNull(duration);
            return this;
        }
        @CustomType.Setter
        public Builder enable(Boolean enable) {
            this.enable = Objects.requireNonNull(enable);
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceTime(String maintenanceTime) {
            this.maintenanceTime = Objects.requireNonNull(maintenanceTime);
            return this;
        }
        @CustomType.Setter
        public Builder weeklyPeriod(String weeklyPeriod) {
            this.weeklyPeriod = Objects.requireNonNull(weeklyPeriod);
            return this;
        }
        public ManagedKubernetesMaintenanceWindow build() {
            final var o = new ManagedKubernetesMaintenanceWindow();
            o.duration = duration;
            o.enable = enable;
            o.maintenanceTime = maintenanceTime;
            o.weeklyPeriod = weeklyPeriod;
            return o;
        }
    }
}
