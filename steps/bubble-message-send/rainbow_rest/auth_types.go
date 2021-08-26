package rainbow_rest

import "time"

type LoginResponse struct {
	Token        string `json:"token"`
	LoggedInUser struct {
		Developer struct {
			Account struct {
				Status         string    `json:"status"`
				LastUpdateDate time.Time `json:"lastUpdateDate"`
			} `json:"account"`
			Sandbox struct {
				Status         string    `json:"status"`
				LastUpdateDate time.Time `json:"lastUpdateDate"`
			} `json:"sandbox"`
			BsAccountID         string    `json:"bsAccountId"`
			ZuoraAccountID      string    `json:"zuoraAccountId"`
			Currency            string    `json:"currency"`
			AccountCreationDate time.Time `json:"accountCreationDate"`
		} `json:"developer"`
		FirstName                          string        `json:"firstName"`
		LastName                           string        `json:"lastName"`
		DisplayName                        string        `json:"displayName"`
		Tags                               []interface{} `json:"tags"`
		IsActive                           bool          `json:"isActive"`
		Roles                              []string      `json:"roles"`
		AdminType                          string        `json:"adminType"`
		AccountType                        string        `json:"accountType"`
		OrganisationID                     string        `json:"organisationId"`
		SiteID                             interface{}   `json:"siteId"`
		SystemID                           interface{}   `json:"systemId"`
		IsInitialized                      bool          `json:"isInitialized"`
		InitializationDate                 time.Time     `json:"initializationDate"`
		LastUpdateDate                     time.Time     `json:"lastUpdateDate"`
		LastAvatarUpdateDate               time.Time     `json:"lastAvatarUpdateDate"`
		CreatedBySelfRegister              bool          `json:"createdBySelfRegister"`
		CreatedByAppID                     interface{}   `json:"createdByAppId"`
		FirstLoginDate                     time.Time     `json:"firstLoginDate"`
		LastLoginDate                      time.Time     `json:"lastLoginDate"`
		LoggedSince                        time.Time     `json:"loggedSince"`
		FailedLoginAttempts                int           `json:"failedLoginAttempts"`
		LastLoginFailureDate               interface{}   `json:"lastLoginFailureDate"`
		LastExpiredTokenRenewedDate        interface{}   `json:"lastExpiredTokenRenewedDate"`
		LastPasswordUpdateDate             interface{}   `json:"lastPasswordUpdateDate"`
		TimeToLive                         int           `json:"timeToLive"`
		TimeToLiveDate                     interface{}   `json:"timeToLiveDate"`
		TerminatedDate                     interface{}   `json:"terminatedDate"`
		IsTerminated                       bool          `json:"isTerminated"`
		GuestMode                          bool          `json:"guestMode"`
		FileSharingCustomisation           string        `json:"fileSharingCustomisation"`
		UserTitleNameCustomisation         string        `json:"userTitleNameCustomisation"`
		SoftphoneOnlyCustomisation         string        `json:"softphoneOnlyCustomisation"`
		UseRoomCustomisation               string        `json:"useRoomCustomisation"`
		PhoneMeetingCustomisation          string        `json:"phoneMeetingCustomisation"`
		UseChannelCustomisation            string        `json:"useChannelCustomisation"`
		UseWebRTCVideoCustomisation        string        `json:"useWebRTCVideoCustomisation"`
		UseWebRTCAudioCustomisation        string        `json:"useWebRTCAudioCustomisation"`
		InstantMessagesCustomisation       string        `json:"instantMessagesCustomisation"`
		UserProfileCustomisation           string        `json:"userProfileCustomisation"`
		FileStorageCustomisation           string        `json:"fileStorageCustomisation"`
		OverridePresenceCustomisation      string        `json:"overridePresenceCustomisation"`
		ChangeTelephonyCustomisation       string        `json:"changeTelephonyCustomisation"`
		ChangeSettingsCustomisation        string        `json:"changeSettingsCustomisation"`
		RecordingConversationCustomisation string        `json:"recordingConversationCustomisation"`
		UseGifCustomisation                string        `json:"useGifCustomisation"`
		UseDialOutCustomisation            string        `json:"useDialOutCustomisation"`
		FileCopyCustomisation              string        `json:"fileCopyCustomisation"`
		FileTransferCustomisation          string        `json:"fileTransferCustomisation"`
		ForbidFileOwnerChangeCustomisation string        `json:"forbidFileOwnerChangeCustomisation"`
		SelectedAppCustomisationTemplate   interface{}   `json:"selectedAppCustomisationTemplate"`
		AlertNotificationReception         string        `json:"alertNotificationReception"`
		SelectedDeviceFirmware             string        `json:"selectedDeviceFirmware"`
		Visibility                         string        `json:"visibility"`
		LoginEmail                         string        `json:"loginEmail"`
		CompanyID                          string        `json:"companyId"`
		JidIm                              string        `json:"jid_im"`
		JidTel                             string        `json:"jid_tel"`
		JidPassword                        string        `json:"jid_password"`
		CreationDate                       time.Time     `json:"creationDate"`
		Emails                             []struct {
			Email string `json:"email"`
			Type  string `json:"type"`
		} `json:"emails"`
		PhoneNumbers []interface{} `json:"phoneNumbers"`
		Profiles     []struct {
			IsDefault                 bool          `json:"isDefault"`
			Status                    string        `json:"status"`
			CanBeSold                 bool          `json:"canBeSold"`
			BusinessModel             string        `json:"businessModel"`
			BusinessSpecific          []string      `json:"businessSpecific"`
			IsExclusive               bool          `json:"isExclusive"`
			IsPrepaid                 bool          `json:"isPrepaid"`
			HasConference             bool          `json:"hasConference"`
			IsBundle                  bool          `json:"isBundle"`
			AssignationDate           time.Time     `json:"assignationDate"`
			ProvisioningNeeded        []interface{} `json:"provisioningNeeded"`
			SubscriptionID            string        `json:"subscriptionId"`
			OfferID                   string        `json:"offerId"`
			OfferName                 string        `json:"offerName"`
			ProfileID                 string        `json:"profileId"`
			ProfileName               string        `json:"profileName"`
			OfferDescription          string        `json:"offerDescription"`
			OfferTechnicalDescription string        `json:"offerTechnicalDescription"`
			OfferReference            string        `json:"offerReference"`
			IsDemo                    bool          `json:"isDemo,omitempty"`
			GroupName                 string        `json:"groupName,omitempty"`
		} `json:"profiles"`
		ActivationDate              time.Time     `json:"activationDate"`
		LastOfflineMailReceivedDate []interface{} `json:"lastOfflineMailReceivedDate"`
		Country                     string        `json:"country"`
		CompanyName                 string        `json:"companyName"`
		Language                    string        `json:"language"`
		State                       interface{}   `json:"state"`
		JobTitle                    string        `json:"jobTitle"`
		NickName                    string        `json:"nickName"`
		Timezone                    string        `json:"timezone"`
		Title                       string        `json:"title"`
		ID                          string        `json:"id"`
		IsInDefaultCompany          bool          `json:"isInDefaultCompany"`
		IsADSearchAvailable         bool          `json:"isADSearchAvailable"`
	} `json:"loggedInUser"`
	LoggedInApplication struct {
		OauthRedirectUris               []interface{} `json:"oauthRedirectUris"`
		EnableOAuthImplicitGrant        bool          `json:"enableOAuthImplicitGrant"`
		State                           string        `json:"state"`
		Env                             string        `json:"env"`
		Kpi                             string        `json:"kpi"`
		DateOfLastTokenWhenRenew        interface{}   `json:"dateOfLastTokenWhenRenew"`
		DateOfDeploymentRequest         time.Time     `json:"dateOfDeploymentRequest"`
		DateOfDeployment                time.Time     `json:"dateOfDeployment"`
		IsRainbowClientDefault          bool          `json:"isRainbowClientDefault"`
		Platform                        string        `json:"platform"`
		Name                            string        `json:"name"`
		Type                            string        `json:"type"`
		Activity                        string        `json:"activity"`
		OwnerID                         string        `json:"ownerId"`
		ProfileID                       string        `json:"profileId"`
		DateOfCreation                  time.Time     `json:"dateOfCreation"`
		Subscriptions                   []interface{} `json:"subscriptions"`
		TurnCredentials                 []interface{} `json:"turnCredentials"`
		TurnPassword                    string        `json:"turnPassword"`
		TurnHosts                       []interface{} `json:"turnHosts"`
		EnableOAuthRefreshTokenRotation bool          `json:"enableOAuthRefreshTokenRotation"`
		ExpirationNotified              bool          `json:"expirationNotified"`
		ExpiringNotified                bool          `json:"expiringNotified"`
		NumberOfRenewal                 int           `json:"numberOfRenewal"`
		Origin                          string        `json:"origin"`
		ID                              string        `json:"id"`
	} `json:"loggedInApplication"`
}
