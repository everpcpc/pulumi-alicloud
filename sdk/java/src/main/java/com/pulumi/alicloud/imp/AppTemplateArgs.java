// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.imp;

import com.pulumi.alicloud.imp.inputs.AppTemplateConfigListArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppTemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppTemplateArgs Empty = new AppTemplateArgs();

    /**
     * The name of the resource.
     * 
     */
    @Import(name="appTemplateName", required=true)
    private Output<String> appTemplateName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> appTemplateName() {
        return this.appTemplateName;
    }

    /**
     * List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
     * 
     */
    @Import(name="componentLists", required=true)
    private Output<List<String>> componentLists;

    /**
     * @return List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
     * 
     */
    public Output<List<String>> componentLists() {
        return this.componentLists;
    }

    /**
     * Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
     * 
     */
    @Import(name="configLists")
    private @Nullable Output<List<AppTemplateConfigListArgs>> configLists;

    /**
     * @return Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
     * 
     */
    public Optional<Output<List<AppTemplateConfigListArgs>>> configLists() {
        return Optional.ofNullable(this.configLists);
    }

    /**
     * Integration mode. Valid values:
     * * paasSDK: Integrated SDK.
     * * standardRoom: Model Room.
     * 
     */
    @Import(name="integrationMode")
    private @Nullable Output<String> integrationMode;

    /**
     * @return Integration mode. Valid values:
     * * paasSDK: Integrated SDK.
     * * standardRoom: Model Room.
     * 
     */
    public Optional<Output<String>> integrationMode() {
        return Optional.ofNullable(this.integrationMode);
    }

    /**
     * Application Template scenario. Valid values: [&#34;business&#34;, &#34;classroom&#34;].
     * 
     */
    @Import(name="scene")
    private @Nullable Output<String> scene;

    /**
     * @return Application Template scenario. Valid values: [&#34;business&#34;, &#34;classroom&#34;].
     * 
     */
    public Optional<Output<String>> scene() {
        return Optional.ofNullable(this.scene);
    }

    private AppTemplateArgs() {}

    private AppTemplateArgs(AppTemplateArgs $) {
        this.appTemplateName = $.appTemplateName;
        this.componentLists = $.componentLists;
        this.configLists = $.configLists;
        this.integrationMode = $.integrationMode;
        this.scene = $.scene;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppTemplateArgs $;

        public Builder() {
            $ = new AppTemplateArgs();
        }

        public Builder(AppTemplateArgs defaults) {
            $ = new AppTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appTemplateName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder appTemplateName(Output<String> appTemplateName) {
            $.appTemplateName = appTemplateName;
            return this;
        }

        /**
         * @param appTemplateName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder appTemplateName(String appTemplateName) {
            return appTemplateName(Output.of(appTemplateName));
        }

        /**
         * @param componentLists List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
         * 
         * @return builder
         * 
         */
        public Builder componentLists(Output<List<String>> componentLists) {
            $.componentLists = componentLists;
            return this;
        }

        /**
         * @param componentLists List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
         * 
         * @return builder
         * 
         */
        public Builder componentLists(List<String> componentLists) {
            return componentLists(Output.of(componentLists));
        }

        /**
         * @param componentLists List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
         * 
         * @return builder
         * 
         */
        public Builder componentLists(String... componentLists) {
            return componentLists(List.of(componentLists));
        }

        /**
         * @param configLists Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
         * 
         * @return builder
         * 
         */
        public Builder configLists(@Nullable Output<List<AppTemplateConfigListArgs>> configLists) {
            $.configLists = configLists;
            return this;
        }

        /**
         * @param configLists Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
         * 
         * @return builder
         * 
         */
        public Builder configLists(List<AppTemplateConfigListArgs> configLists) {
            return configLists(Output.of(configLists));
        }

        /**
         * @param configLists Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
         * 
         * @return builder
         * 
         */
        public Builder configLists(AppTemplateConfigListArgs... configLists) {
            return configLists(List.of(configLists));
        }

        /**
         * @param integrationMode Integration mode. Valid values:
         * * paasSDK: Integrated SDK.
         * * standardRoom: Model Room.
         * 
         * @return builder
         * 
         */
        public Builder integrationMode(@Nullable Output<String> integrationMode) {
            $.integrationMode = integrationMode;
            return this;
        }

        /**
         * @param integrationMode Integration mode. Valid values:
         * * paasSDK: Integrated SDK.
         * * standardRoom: Model Room.
         * 
         * @return builder
         * 
         */
        public Builder integrationMode(String integrationMode) {
            return integrationMode(Output.of(integrationMode));
        }

        /**
         * @param scene Application Template scenario. Valid values: [&#34;business&#34;, &#34;classroom&#34;].
         * 
         * @return builder
         * 
         */
        public Builder scene(@Nullable Output<String> scene) {
            $.scene = scene;
            return this;
        }

        /**
         * @param scene Application Template scenario. Valid values: [&#34;business&#34;, &#34;classroom&#34;].
         * 
         * @return builder
         * 
         */
        public Builder scene(String scene) {
            return scene(Output.of(scene));
        }

        public AppTemplateArgs build() {
            $.appTemplateName = Objects.requireNonNull($.appTemplateName, "expected parameter 'appTemplateName' to be non-null");
            $.componentLists = Objects.requireNonNull($.componentLists, "expected parameter 'componentLists' to be non-null");
            return $;
        }
    }

}
