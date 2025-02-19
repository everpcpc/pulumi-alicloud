// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseState extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseState Empty = new DatabaseState();

    /**
     * Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is &#34;utf8&#34; \(`utf8mb4` only supports versions 5.5 and 5.6\).
     * 
     */
    @Import(name="characterSetName")
    private @Nullable Output<String> characterSetName;

    /**
     * @return Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is &#34;utf8&#34; \(`utf8mb4` only supports versions 5.5 and 5.6\).
     * 
     */
    public Optional<Output<String>> characterSetName() {
        return Optional.ofNullable(this.characterSetName);
    }

    /**
     * The Id of cluster that can run database.
     * 
     */
    @Import(name="dbClusterId")
    private @Nullable Output<String> dbClusterId;

    /**
     * @return The Id of cluster that can run database.
     * 
     */
    public Optional<Output<String>> dbClusterId() {
        return Optional.ofNullable(this.dbClusterId);
    }

    /**
     * Database description. It must start with a Chinese character or English letter, cannot start with &#34;http://&#34; or &#34;https://&#34;. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length must be 2-256 characters.
     * 
     */
    @Import(name="dbDescription")
    private @Nullable Output<String> dbDescription;

    /**
     * @return Database description. It must start with a Chinese character or English letter, cannot start with &#34;http://&#34; or &#34;https://&#34;. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length must be 2-256 characters.
     * 
     */
    public Optional<Output<String>> dbDescription() {
        return Optional.ofNullable(this.dbDescription);
    }

    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
     * 
     */
    @Import(name="dbName")
    private @Nullable Output<String> dbName;

    /**
     * @return Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
     * 
     */
    public Optional<Output<String>> dbName() {
        return Optional.ofNullable(this.dbName);
    }

    private DatabaseState() {}

    private DatabaseState(DatabaseState $) {
        this.characterSetName = $.characterSetName;
        this.dbClusterId = $.dbClusterId;
        this.dbDescription = $.dbDescription;
        this.dbName = $.dbName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseState $;

        public Builder() {
            $ = new DatabaseState();
        }

        public Builder(DatabaseState defaults) {
            $ = new DatabaseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param characterSetName Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is &#34;utf8&#34; \(`utf8mb4` only supports versions 5.5 and 5.6\).
         * 
         * @return builder
         * 
         */
        public Builder characterSetName(@Nullable Output<String> characterSetName) {
            $.characterSetName = characterSetName;
            return this;
        }

        /**
         * @param characterSetName Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is &#34;utf8&#34; \(`utf8mb4` only supports versions 5.5 and 5.6\).
         * 
         * @return builder
         * 
         */
        public Builder characterSetName(String characterSetName) {
            return characterSetName(Output.of(characterSetName));
        }

        /**
         * @param dbClusterId The Id of cluster that can run database.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(@Nullable Output<String> dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param dbClusterId The Id of cluster that can run database.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(String dbClusterId) {
            return dbClusterId(Output.of(dbClusterId));
        }

        /**
         * @param dbDescription Database description. It must start with a Chinese character or English letter, cannot start with &#34;http://&#34; or &#34;https://&#34;. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length must be 2-256 characters.
         * 
         * @return builder
         * 
         */
        public Builder dbDescription(@Nullable Output<String> dbDescription) {
            $.dbDescription = dbDescription;
            return this;
        }

        /**
         * @param dbDescription Database description. It must start with a Chinese character or English letter, cannot start with &#34;http://&#34; or &#34;https://&#34;. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length must be 2-256 characters.
         * 
         * @return builder
         * 
         */
        public Builder dbDescription(String dbDescription) {
            return dbDescription(Output.of(dbDescription));
        }

        /**
         * @param dbName Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
         * 
         * @return builder
         * 
         */
        public Builder dbName(@Nullable Output<String> dbName) {
            $.dbName = dbName;
            return this;
        }

        /**
         * @param dbName Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
         * 
         * @return builder
         * 
         */
        public Builder dbName(String dbName) {
            return dbName(Output.of(dbName));
        }

        public DatabaseState build() {
            return $;
        }
    }

}
