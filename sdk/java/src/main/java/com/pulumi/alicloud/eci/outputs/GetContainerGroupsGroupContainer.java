// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.alicloud.eci.outputs.GetContainerGroupsGroupContainerEnvironmentVar;
import com.pulumi.alicloud.eci.outputs.GetContainerGroupsGroupContainerPort;
import com.pulumi.alicloud.eci.outputs.GetContainerGroupsGroupContainerVolumeMount;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetContainerGroupsGroupContainer {
    /**
     * @return The arguments passed to the commands.
     * 
     */
    private List<String> args;
    /**
     * @return The commands run by the container.
     * 
     */
    private List<String> commands;
    /**
     * @return The amount of CPU resources allocated to the container.
     * 
     */
    private Double cpu;
    /**
     * @return The environment variables.
     * 
     */
    private List<GetContainerGroupsGroupContainerEnvironmentVar> environmentVars;
    /**
     * @return The amount of GPU resources allocated to the container.
     * 
     */
    private Integer gpu;
    /**
     * @return The image of the container.
     * 
     */
    private String image;
    /**
     * @return The policy for pulling an image.
     * 
     */
    private String imagePullPolicy;
    /**
     * @return The amount of memory resources allocated to the container group.
     * 
     */
    private Double memory;
    /**
     * @return The name of the volume.
     * 
     */
    private String name;
    /**
     * @return The exposed ports and protocols. Maximum: `100`.
     * 
     */
    private List<GetContainerGroupsGroupContainerPort> ports;
    /**
     * @return Indicates whether the container is ready.
     * 
     */
    private Boolean ready;
    /**
     * @return The number of times that the container has restarted.
     * 
     */
    private Integer restartCount;
    /**
     * @return The list of volumes mounted to the container.
     * 
     */
    private List<GetContainerGroupsGroupContainerVolumeMount> volumeMounts;
    /**
     * @return The working directory of the container.
     * 
     */
    private String workingDir;

    private GetContainerGroupsGroupContainer() {}
    /**
     * @return The arguments passed to the commands.
     * 
     */
    public List<String> args() {
        return this.args;
    }
    /**
     * @return The commands run by the container.
     * 
     */
    public List<String> commands() {
        return this.commands;
    }
    /**
     * @return The amount of CPU resources allocated to the container.
     * 
     */
    public Double cpu() {
        return this.cpu;
    }
    /**
     * @return The environment variables.
     * 
     */
    public List<GetContainerGroupsGroupContainerEnvironmentVar> environmentVars() {
        return this.environmentVars;
    }
    /**
     * @return The amount of GPU resources allocated to the container.
     * 
     */
    public Integer gpu() {
        return this.gpu;
    }
    /**
     * @return The image of the container.
     * 
     */
    public String image() {
        return this.image;
    }
    /**
     * @return The policy for pulling an image.
     * 
     */
    public String imagePullPolicy() {
        return this.imagePullPolicy;
    }
    /**
     * @return The amount of memory resources allocated to the container group.
     * 
     */
    public Double memory() {
        return this.memory;
    }
    /**
     * @return The name of the volume.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The exposed ports and protocols. Maximum: `100`.
     * 
     */
    public List<GetContainerGroupsGroupContainerPort> ports() {
        return this.ports;
    }
    /**
     * @return Indicates whether the container is ready.
     * 
     */
    public Boolean ready() {
        return this.ready;
    }
    /**
     * @return The number of times that the container has restarted.
     * 
     */
    public Integer restartCount() {
        return this.restartCount;
    }
    /**
     * @return The list of volumes mounted to the container.
     * 
     */
    public List<GetContainerGroupsGroupContainerVolumeMount> volumeMounts() {
        return this.volumeMounts;
    }
    /**
     * @return The working directory of the container.
     * 
     */
    public String workingDir() {
        return this.workingDir;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerGroupsGroupContainer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> args;
        private List<String> commands;
        private Double cpu;
        private List<GetContainerGroupsGroupContainerEnvironmentVar> environmentVars;
        private Integer gpu;
        private String image;
        private String imagePullPolicy;
        private Double memory;
        private String name;
        private List<GetContainerGroupsGroupContainerPort> ports;
        private Boolean ready;
        private Integer restartCount;
        private List<GetContainerGroupsGroupContainerVolumeMount> volumeMounts;
        private String workingDir;
        public Builder() {}
        public Builder(GetContainerGroupsGroupContainer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.args = defaults.args;
    	      this.commands = defaults.commands;
    	      this.cpu = defaults.cpu;
    	      this.environmentVars = defaults.environmentVars;
    	      this.gpu = defaults.gpu;
    	      this.image = defaults.image;
    	      this.imagePullPolicy = defaults.imagePullPolicy;
    	      this.memory = defaults.memory;
    	      this.name = defaults.name;
    	      this.ports = defaults.ports;
    	      this.ready = defaults.ready;
    	      this.restartCount = defaults.restartCount;
    	      this.volumeMounts = defaults.volumeMounts;
    	      this.workingDir = defaults.workingDir;
        }

        @CustomType.Setter
        public Builder args(List<String> args) {
            this.args = Objects.requireNonNull(args);
            return this;
        }
        public Builder args(String... args) {
            return args(List.of(args));
        }
        @CustomType.Setter
        public Builder commands(List<String> commands) {
            this.commands = Objects.requireNonNull(commands);
            return this;
        }
        public Builder commands(String... commands) {
            return commands(List.of(commands));
        }
        @CustomType.Setter
        public Builder cpu(Double cpu) {
            this.cpu = Objects.requireNonNull(cpu);
            return this;
        }
        @CustomType.Setter
        public Builder environmentVars(List<GetContainerGroupsGroupContainerEnvironmentVar> environmentVars) {
            this.environmentVars = Objects.requireNonNull(environmentVars);
            return this;
        }
        public Builder environmentVars(GetContainerGroupsGroupContainerEnvironmentVar... environmentVars) {
            return environmentVars(List.of(environmentVars));
        }
        @CustomType.Setter
        public Builder gpu(Integer gpu) {
            this.gpu = Objects.requireNonNull(gpu);
            return this;
        }
        @CustomType.Setter
        public Builder image(String image) {
            this.image = Objects.requireNonNull(image);
            return this;
        }
        @CustomType.Setter
        public Builder imagePullPolicy(String imagePullPolicy) {
            this.imagePullPolicy = Objects.requireNonNull(imagePullPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder memory(Double memory) {
            this.memory = Objects.requireNonNull(memory);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder ports(List<GetContainerGroupsGroupContainerPort> ports) {
            this.ports = Objects.requireNonNull(ports);
            return this;
        }
        public Builder ports(GetContainerGroupsGroupContainerPort... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder ready(Boolean ready) {
            this.ready = Objects.requireNonNull(ready);
            return this;
        }
        @CustomType.Setter
        public Builder restartCount(Integer restartCount) {
            this.restartCount = Objects.requireNonNull(restartCount);
            return this;
        }
        @CustomType.Setter
        public Builder volumeMounts(List<GetContainerGroupsGroupContainerVolumeMount> volumeMounts) {
            this.volumeMounts = Objects.requireNonNull(volumeMounts);
            return this;
        }
        public Builder volumeMounts(GetContainerGroupsGroupContainerVolumeMount... volumeMounts) {
            return volumeMounts(List.of(volumeMounts));
        }
        @CustomType.Setter
        public Builder workingDir(String workingDir) {
            this.workingDir = Objects.requireNonNull(workingDir);
            return this;
        }
        public GetContainerGroupsGroupContainer build() {
            final var o = new GetContainerGroupsGroupContainer();
            o.args = args;
            o.commands = commands;
            o.cpu = cpu;
            o.environmentVars = environmentVars;
            o.gpu = gpu;
            o.image = image;
            o.imagePullPolicy = imagePullPolicy;
            o.memory = memory;
            o.name = name;
            o.ports = ports;
            o.ready = ready;
            o.restartCount = restartCount;
            o.volumeMounts = volumeMounts;
            o.workingDir = workingDir;
            return o;
        }
    }
}
