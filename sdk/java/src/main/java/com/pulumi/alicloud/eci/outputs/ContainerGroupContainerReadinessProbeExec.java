// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ContainerGroupContainerReadinessProbeExec {
    /**
     * @return The commands run by the init container.
     * 
     */
    private @Nullable List<String> commands;

    private ContainerGroupContainerReadinessProbeExec() {}
    /**
     * @return The commands run by the init container.
     * 
     */
    public List<String> commands() {
        return this.commands == null ? List.of() : this.commands;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerGroupContainerReadinessProbeExec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> commands;
        public Builder() {}
        public Builder(ContainerGroupContainerReadinessProbeExec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commands = defaults.commands;
        }

        @CustomType.Setter
        public Builder commands(@Nullable List<String> commands) {
            this.commands = commands;
            return this;
        }
        public Builder commands(String... commands) {
            return commands(List.of(commands));
        }
        public ContainerGroupContainerReadinessProbeExec build() {
            final var o = new ContainerGroupContainerReadinessProbeExec();
            o.commands = commands;
            return o;
        }
    }
}
