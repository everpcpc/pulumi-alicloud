// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vod.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EditingProjectState extends com.pulumi.resources.ResourceArgs {

    public static final EditingProjectState Empty = new EditingProjectState();

    /**
     * The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
     * 
     */
    @Import(name="coverUrl")
    private @Nullable Output<String> coverUrl;

    /**
     * @return The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
     * 
     */
    public Optional<Output<String>> coverUrl() {
        return Optional.ofNullable(this.coverUrl);
    }

    /**
     * The region where you want to create the online editing project.
     * 
     */
    @Import(name="division")
    private @Nullable Output<String> division;

    /**
     * @return The region where you want to create the online editing project.
     * 
     */
    public Optional<Output<String>> division() {
        return Optional.ofNullable(this.division);
    }

    /**
     * The description of the online editing project.
     * 
     */
    @Import(name="editingProjectName")
    private @Nullable Output<String> editingProjectName;

    /**
     * @return The description of the online editing project.
     * 
     */
    public Optional<Output<String>> editingProjectName() {
        return Optional.ofNullable(this.editingProjectName);
    }

    /**
     * The Status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The Status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
     * 
     */
    @Import(name="timeline")
    private @Nullable Output<String> timeline;

    /**
     * @return The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
     * 
     */
    public Optional<Output<String>> timeline() {
        return Optional.ofNullable(this.timeline);
    }

    /**
     * The title of the online editing project.
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return The title of the online editing project.
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    private EditingProjectState() {}

    private EditingProjectState(EditingProjectState $) {
        this.coverUrl = $.coverUrl;
        this.division = $.division;
        this.editingProjectName = $.editingProjectName;
        this.status = $.status;
        this.timeline = $.timeline;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EditingProjectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EditingProjectState $;

        public Builder() {
            $ = new EditingProjectState();
        }

        public Builder(EditingProjectState defaults) {
            $ = new EditingProjectState(Objects.requireNonNull(defaults));
        }

        /**
         * @param coverUrl The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
         * 
         * @return builder
         * 
         */
        public Builder coverUrl(@Nullable Output<String> coverUrl) {
            $.coverUrl = coverUrl;
            return this;
        }

        /**
         * @param coverUrl The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
         * 
         * @return builder
         * 
         */
        public Builder coverUrl(String coverUrl) {
            return coverUrl(Output.of(coverUrl));
        }

        /**
         * @param division The region where you want to create the online editing project.
         * 
         * @return builder
         * 
         */
        public Builder division(@Nullable Output<String> division) {
            $.division = division;
            return this;
        }

        /**
         * @param division The region where you want to create the online editing project.
         * 
         * @return builder
         * 
         */
        public Builder division(String division) {
            return division(Output.of(division));
        }

        /**
         * @param editingProjectName The description of the online editing project.
         * 
         * @return builder
         * 
         */
        public Builder editingProjectName(@Nullable Output<String> editingProjectName) {
            $.editingProjectName = editingProjectName;
            return this;
        }

        /**
         * @param editingProjectName The description of the online editing project.
         * 
         * @return builder
         * 
         */
        public Builder editingProjectName(String editingProjectName) {
            return editingProjectName(Output.of(editingProjectName));
        }

        /**
         * @param status The Status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The Status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param timeline The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
         * 
         * @return builder
         * 
         */
        public Builder timeline(@Nullable Output<String> timeline) {
            $.timeline = timeline;
            return this;
        }

        /**
         * @param timeline The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
         * 
         * @return builder
         * 
         */
        public Builder timeline(String timeline) {
            return timeline(Output.of(timeline));
        }

        /**
         * @param title The title of the online editing project.
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title The title of the online editing project.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public EditingProjectState build() {
            return $;
        }
    }

}