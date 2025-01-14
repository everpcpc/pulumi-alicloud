// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMscSubContactsContact {
    /**
     * @return UID.
     * 
     */
    private String accountUid;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String contactId;
    /**
     * @return The User&#39;s Contact Name. **Note:** The name must be 2 to 12 characters in length, and can contain uppercase and lowercase letters.
     * 
     */
    private String contactName;
    /**
     * @return The User&#39;s Contact Email Address.
     * 
     */
    private String email;
    /**
     * @return The ID of the Contact.
     * 
     */
    private String id;
    /**
     * @return Indicates Whether the BGP Group Is the Account Itself.
     * 
     */
    private Boolean isAccount;
    /**
     * @return Whether They Have Expired Or Not.
     * 
     */
    private Boolean isObsolete;
    /**
     * @return Email Validation for.
     * 
     */
    private Boolean isVerifiedEmail;
    /**
     * @return If the Phone Verification.
     * 
     */
    private Boolean isVerifiedMobile;
    /**
     * @return Last Verification Email Transmission Time.
     * 
     */
    private String lastEmailVerificationTimeStamp;
    /**
     * @return The Pieces of Authentication SMS Sending Time.
     * 
     */
    private String lastMobileVerificationTimeStamp;
    /**
     * @return The User&#39;s Telephone.
     * 
     */
    private String mobile;
    /**
     * @return The User&#39;s Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
     * 
     */
    private String position;

    private GetMscSubContactsContact() {}
    /**
     * @return UID.
     * 
     */
    public String accountUid() {
        return this.accountUid;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String contactId() {
        return this.contactId;
    }
    /**
     * @return The User&#39;s Contact Name. **Note:** The name must be 2 to 12 characters in length, and can contain uppercase and lowercase letters.
     * 
     */
    public String contactName() {
        return this.contactName;
    }
    /**
     * @return The User&#39;s Contact Email Address.
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return The ID of the Contact.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Indicates Whether the BGP Group Is the Account Itself.
     * 
     */
    public Boolean isAccount() {
        return this.isAccount;
    }
    /**
     * @return Whether They Have Expired Or Not.
     * 
     */
    public Boolean isObsolete() {
        return this.isObsolete;
    }
    /**
     * @return Email Validation for.
     * 
     */
    public Boolean isVerifiedEmail() {
        return this.isVerifiedEmail;
    }
    /**
     * @return If the Phone Verification.
     * 
     */
    public Boolean isVerifiedMobile() {
        return this.isVerifiedMobile;
    }
    /**
     * @return Last Verification Email Transmission Time.
     * 
     */
    public String lastEmailVerificationTimeStamp() {
        return this.lastEmailVerificationTimeStamp;
    }
    /**
     * @return The Pieces of Authentication SMS Sending Time.
     * 
     */
    public String lastMobileVerificationTimeStamp() {
        return this.lastMobileVerificationTimeStamp;
    }
    /**
     * @return The User&#39;s Telephone.
     * 
     */
    public String mobile() {
        return this.mobile;
    }
    /**
     * @return The User&#39;s Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
     * 
     */
    public String position() {
        return this.position;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMscSubContactsContact defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountUid;
        private String contactId;
        private String contactName;
        private String email;
        private String id;
        private Boolean isAccount;
        private Boolean isObsolete;
        private Boolean isVerifiedEmail;
        private Boolean isVerifiedMobile;
        private String lastEmailVerificationTimeStamp;
        private String lastMobileVerificationTimeStamp;
        private String mobile;
        private String position;
        public Builder() {}
        public Builder(GetMscSubContactsContact defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountUid = defaults.accountUid;
    	      this.contactId = defaults.contactId;
    	      this.contactName = defaults.contactName;
    	      this.email = defaults.email;
    	      this.id = defaults.id;
    	      this.isAccount = defaults.isAccount;
    	      this.isObsolete = defaults.isObsolete;
    	      this.isVerifiedEmail = defaults.isVerifiedEmail;
    	      this.isVerifiedMobile = defaults.isVerifiedMobile;
    	      this.lastEmailVerificationTimeStamp = defaults.lastEmailVerificationTimeStamp;
    	      this.lastMobileVerificationTimeStamp = defaults.lastMobileVerificationTimeStamp;
    	      this.mobile = defaults.mobile;
    	      this.position = defaults.position;
        }

        @CustomType.Setter
        public Builder accountUid(String accountUid) {
            this.accountUid = Objects.requireNonNull(accountUid);
            return this;
        }
        @CustomType.Setter
        public Builder contactId(String contactId) {
            this.contactId = Objects.requireNonNull(contactId);
            return this;
        }
        @CustomType.Setter
        public Builder contactName(String contactName) {
            this.contactName = Objects.requireNonNull(contactName);
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isAccount(Boolean isAccount) {
            this.isAccount = Objects.requireNonNull(isAccount);
            return this;
        }
        @CustomType.Setter
        public Builder isObsolete(Boolean isObsolete) {
            this.isObsolete = Objects.requireNonNull(isObsolete);
            return this;
        }
        @CustomType.Setter
        public Builder isVerifiedEmail(Boolean isVerifiedEmail) {
            this.isVerifiedEmail = Objects.requireNonNull(isVerifiedEmail);
            return this;
        }
        @CustomType.Setter
        public Builder isVerifiedMobile(Boolean isVerifiedMobile) {
            this.isVerifiedMobile = Objects.requireNonNull(isVerifiedMobile);
            return this;
        }
        @CustomType.Setter
        public Builder lastEmailVerificationTimeStamp(String lastEmailVerificationTimeStamp) {
            this.lastEmailVerificationTimeStamp = Objects.requireNonNull(lastEmailVerificationTimeStamp);
            return this;
        }
        @CustomType.Setter
        public Builder lastMobileVerificationTimeStamp(String lastMobileVerificationTimeStamp) {
            this.lastMobileVerificationTimeStamp = Objects.requireNonNull(lastMobileVerificationTimeStamp);
            return this;
        }
        @CustomType.Setter
        public Builder mobile(String mobile) {
            this.mobile = Objects.requireNonNull(mobile);
            return this;
        }
        @CustomType.Setter
        public Builder position(String position) {
            this.position = Objects.requireNonNull(position);
            return this;
        }
        public GetMscSubContactsContact build() {
            final var o = new GetMscSubContactsContact();
            o.accountUid = accountUid;
            o.contactId = contactId;
            o.contactName = contactName;
            o.email = email;
            o.id = id;
            o.isAccount = isAccount;
            o.isObsolete = isObsolete;
            o.isVerifiedEmail = isVerifiedEmail;
            o.isVerifiedMobile = isVerifiedMobile;
            o.lastEmailVerificationTimeStamp = lastEmailVerificationTimeStamp;
            o.lastMobileVerificationTimeStamp = lastMobileVerificationTimeStamp;
            o.mobile = mobile;
            o.position = position;
            return o;
        }
    }
}
