// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.imp.outputs;

import com.pulumi.alicloud.imp.outputs.GetAppTemplatesTemplateConfigList;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAppTemplatesTemplate {
    /**
     * @return Apply template creator.
     * 
     */
    private String appTemplateCreator;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String appTemplateId;
    /**
     * @return The name of the resource.
     * 
     */
    private String appTemplateName;
    /**
     * @return List of components.
     * 
     */
    private List<String> componentLists;
    /**
     * @return List of config.
     * 
     */
    private List<GetAppTemplatesTemplateConfigList> configLists;
    /**
     * @return Creation time.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the App Template.
     * 
     */
    private String id;
    /**
     * @return Integration mode (Integrated SDK:paasSDK, Model Room: standardRoom).
     * 
     */
    private String integrationMode;
    /**
     * @return Application Template scenario, e-commerce business, classroom classroom.
     * 
     */
    private String scene;
    /**
     * @return SDK information.
     * 
     */
    private String sdkInfo;
    /**
     * @return Model room information.
     * 
     */
    private String standardRoomInfo;
    /**
     * @return Application template usage status.
     * 
     */
    private String status;

    private GetAppTemplatesTemplate() {}
    /**
     * @return Apply template creator.
     * 
     */
    public String appTemplateCreator() {
        return this.appTemplateCreator;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String appTemplateId() {
        return this.appTemplateId;
    }
    /**
     * @return The name of the resource.
     * 
     */
    public String appTemplateName() {
        return this.appTemplateName;
    }
    /**
     * @return List of components.
     * 
     */
    public List<String> componentLists() {
        return this.componentLists;
    }
    /**
     * @return List of config.
     * 
     */
    public List<GetAppTemplatesTemplateConfigList> configLists() {
        return this.configLists;
    }
    /**
     * @return Creation time.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the App Template.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Integration mode (Integrated SDK:paasSDK, Model Room: standardRoom).
     * 
     */
    public String integrationMode() {
        return this.integrationMode;
    }
    /**
     * @return Application Template scenario, e-commerce business, classroom classroom.
     * 
     */
    public String scene() {
        return this.scene;
    }
    /**
     * @return SDK information.
     * 
     */
    public String sdkInfo() {
        return this.sdkInfo;
    }
    /**
     * @return Model room information.
     * 
     */
    public String standardRoomInfo() {
        return this.standardRoomInfo;
    }
    /**
     * @return Application template usage status.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppTemplatesTemplate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String appTemplateCreator;
        private String appTemplateId;
        private String appTemplateName;
        private List<String> componentLists;
        private List<GetAppTemplatesTemplateConfigList> configLists;
        private String createTime;
        private String id;
        private String integrationMode;
        private String scene;
        private String sdkInfo;
        private String standardRoomInfo;
        private String status;
        public Builder() {}
        public Builder(GetAppTemplatesTemplate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appTemplateCreator = defaults.appTemplateCreator;
    	      this.appTemplateId = defaults.appTemplateId;
    	      this.appTemplateName = defaults.appTemplateName;
    	      this.componentLists = defaults.componentLists;
    	      this.configLists = defaults.configLists;
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.integrationMode = defaults.integrationMode;
    	      this.scene = defaults.scene;
    	      this.sdkInfo = defaults.sdkInfo;
    	      this.standardRoomInfo = defaults.standardRoomInfo;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder appTemplateCreator(String appTemplateCreator) {
            this.appTemplateCreator = Objects.requireNonNull(appTemplateCreator);
            return this;
        }
        @CustomType.Setter
        public Builder appTemplateId(String appTemplateId) {
            this.appTemplateId = Objects.requireNonNull(appTemplateId);
            return this;
        }
        @CustomType.Setter
        public Builder appTemplateName(String appTemplateName) {
            this.appTemplateName = Objects.requireNonNull(appTemplateName);
            return this;
        }
        @CustomType.Setter
        public Builder componentLists(List<String> componentLists) {
            this.componentLists = Objects.requireNonNull(componentLists);
            return this;
        }
        public Builder componentLists(String... componentLists) {
            return componentLists(List.of(componentLists));
        }
        @CustomType.Setter
        public Builder configLists(List<GetAppTemplatesTemplateConfigList> configLists) {
            this.configLists = Objects.requireNonNull(configLists);
            return this;
        }
        public Builder configLists(GetAppTemplatesTemplateConfigList... configLists) {
            return configLists(List.of(configLists));
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder integrationMode(String integrationMode) {
            this.integrationMode = Objects.requireNonNull(integrationMode);
            return this;
        }
        @CustomType.Setter
        public Builder scene(String scene) {
            this.scene = Objects.requireNonNull(scene);
            return this;
        }
        @CustomType.Setter
        public Builder sdkInfo(String sdkInfo) {
            this.sdkInfo = Objects.requireNonNull(sdkInfo);
            return this;
        }
        @CustomType.Setter
        public Builder standardRoomInfo(String standardRoomInfo) {
            this.standardRoomInfo = Objects.requireNonNull(standardRoomInfo);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetAppTemplatesTemplate build() {
            final var o = new GetAppTemplatesTemplate();
            o.appTemplateCreator = appTemplateCreator;
            o.appTemplateId = appTemplateId;
            o.appTemplateName = appTemplateName;
            o.componentLists = componentLists;
            o.configLists = configLists;
            o.createTime = createTime;
            o.id = id;
            o.integrationMode = integrationMode;
            o.scene = scene;
            o.sdkInfo = sdkInfo;
            o.standardRoomInfo = standardRoomInfo;
            o.status = status;
            return o;
        }
    }
}