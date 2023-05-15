// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.securitycenter.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupsGroup {
    /**
     * @return GroupFlag, &#39;0&#39; mean default group(created by system), &#39;1&#39; means customer defined group.
     * 
     */
    private Integer groupFlag;
    /**
     * @return The ID of Group.
     * 
     */
    private String groupId;
    /**
     * @return The name of Group.
     * 
     */
    private String groupName;
    /**
     * @return The ID of the Group(same as the group_id).
     * 
     */
    private String id;

    private GetGroupsGroup() {}
    /**
     * @return GroupFlag, &#39;0&#39; mean default group(created by system), &#39;1&#39; means customer defined group.
     * 
     */
    public Integer groupFlag() {
        return this.groupFlag;
    }
    /**
     * @return The ID of Group.
     * 
     */
    public String groupId() {
        return this.groupId;
    }
    /**
     * @return The name of Group.
     * 
     */
    public String groupName() {
        return this.groupName;
    }
    /**
     * @return The ID of the Group(same as the group_id).
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer groupFlag;
        private String groupId;
        private String groupName;
        private String id;
        public Builder() {}
        public Builder(GetGroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupFlag = defaults.groupFlag;
    	      this.groupId = defaults.groupId;
    	      this.groupName = defaults.groupName;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder groupFlag(Integer groupFlag) {
            this.groupFlag = Objects.requireNonNull(groupFlag);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(String groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder groupName(String groupName) {
            this.groupName = Objects.requireNonNull(groupName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetGroupsGroup build() {
            final var o = new GetGroupsGroup();
            o.groupFlag = groupFlag;
            o.groupId = groupId;
            o.groupName = groupName;
            o.id = id;
            return o;
        }
    }
}
