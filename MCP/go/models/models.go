package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// BatchUpdateDealsResponse represents the BatchUpdateDealsResponse schema from the OpenAPI specification
type BatchUpdateDealsResponse struct {
	Deals []Deal `json:"deals,omitempty"` // Deals updated.
}

// PublisherProfileMobileApplication represents the PublisherProfileMobileApplication schema from the OpenAPI specification
type PublisherProfileMobileApplication struct {
	Appstore string `json:"appStore,omitempty"` // The app store the app belongs to. Can be used to filter the response of the publisherProfiles.list method.
	Externalappid string `json:"externalAppId,omitempty"` // The external ID for the app from its app store. Can be used to filter the response of the publisherProfiles.list method.
	Name string `json:"name,omitempty"` // The name of the app.
}

// ActivateClientUserRequest represents the ActivateClientUserRequest schema from the OpenAPI specification
type ActivateClientUserRequest struct {
}

// FinalizedDeal represents the FinalizedDeal schema from the OpenAPI specification
type FinalizedDeal struct {
	Rtbmetrics RtbMetrics `json:"rtbMetrics,omitempty"` // Real-time bidding metrics. For what each metric means refer to [Report metrics](https://support.google.com/adxbuyer/answer/6115195#report-metrics)
	Deal Deal `json:"deal,omitempty"` // A deal represents a segment of inventory for displaying ads that contains the terms and targeting information that is used for serving as well as the deal stats and status. Note: A proposal may contain multiple deals.
	Dealpausinginfo DealPausingInfo `json:"dealPausingInfo,omitempty"` // Information related to deal pausing.
	Dealservingstatus string `json:"dealServingStatus,omitempty"` // Serving status of the deal.
	Name string `json:"name,omitempty"` // The resource name of the finalized deal. Format: `buyers/{accountId}/finalizedDeals/{finalizedDealId}`
	Readytoserve bool `json:"readyToServe,omitempty"` // Whether the Programmatic Guaranteed deal is ready for serving.
}

// ListClientUsersResponse represents the ListClientUsersResponse schema from the OpenAPI specification
type ListClientUsersResponse struct {
	Clientusers []ClientUser `json:"clientUsers,omitempty"` // The returned list of client users.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token to retrieve the next page of results. Pass this value in the ListClientUsersRequest.pageToken field in the subsequent call to the list method to retrieve the next page of results.
}

// PrivateAuctionTerms represents the PrivateAuctionTerms schema from the OpenAPI specification
type PrivateAuctionTerms struct {
	Floorprice Price `json:"floorPrice,omitempty"` // Represents a price and a pricing type for a deal.
	Openauctionallowed bool `json:"openAuctionAllowed,omitempty"` // Output only. True if open auction buyers are allowed to compete with invited buyers in this private auction.
}

// UriTargeting represents the UriTargeting schema from the OpenAPI specification
type UriTargeting struct {
	Excludeduris []string `json:"excludedUris,omitempty"` // A list of URLs to be excluded.
	Targeteduris []string `json:"targetedUris,omitempty"` // A list of URLs to be included.
}

// SubscribeClientsRequest represents the SubscribeClientsRequest schema from the OpenAPI specification
type SubscribeClientsRequest struct {
	Clients []string `json:"clients,omitempty"` // Optional. A list of client buyers to subscribe to the auction package, with client buyer in the format `buyers/{accountId}/clients/{clientAccountId}`. The current buyer will be subscribed to the auction package regardless of the list contents if not already.
}

// DeliveryControl represents the DeliveryControl schema from the OpenAPI specification
type DeliveryControl struct {
	Roadblockingtype string `json:"roadblockingType,omitempty"` // Output only. Specifies the roadblocking type in display creatives.
	Companiondeliverytype string `json:"companionDeliveryType,omitempty"` // Output only. Specifies roadblocking in a main companion lineitem.
	Creativerotationtype string `json:"creativeRotationType,omitempty"` // Output only. Specifies strategy to use for selecting a creative when multiple creatives of the same size are available.
	Deliveryratetype string `json:"deliveryRateType,omitempty"` // Output only. Specifies how the impression delivery will be paced.
	Frequencycap []FrequencyCap `json:"frequencyCap,omitempty"` // Output only. Specifies any frequency caps. Cannot be filtered within ListDealsRequest.
}

// Deal represents the Deal schema from the OpenAPI specification
type Deal struct {
	Targeting MarketplaceTargeting `json:"targeting,omitempty"` // Targeting represents different criteria that can be used to target inventory. For example, they can choose to target inventory only if the user is in the US. Multiple types of targeting are always applied as a logical AND, unless noted otherwise.
	Privateauctionterms PrivateAuctionTerms `json:"privateAuctionTerms,omitempty"` // Pricing terms for Private Auctions.
	Deliverycontrol DeliveryControl `json:"deliveryControl,omitempty"` // Message contains details about how the deal will be paced.
	Flightstarttime string `json:"flightStartTime,omitempty"` // Proposed flight start time of the deal. This will generally be stored in the granularity of one second since deal serving starts at seconds boundary. Any time specified with more granularity (for example, in milliseconds) will be truncated towards the start of time in seconds.
	Publisherprofile string `json:"publisherProfile,omitempty"` // Immutable. Reference to the seller on the deal. Format: `buyers/{buyerAccountId}/publisherProfiles/{publisherProfileId}`
	Buyer string `json:"buyer,omitempty"` // Output only. Refers to a buyer in The Realtime-bidding API. Format: `buyers/{buyerAccountId}`
	Flightendtime string `json:"flightEndTime,omitempty"` // Proposed flight end time of the deal. This will generally be stored in a granularity of a second. A value is not necessary for Private Auction deals.
	Billedbuyer string `json:"billedBuyer,omitempty"` // Output only. When the client field is populated, this field refers to the buyer who creates and manages the client buyer and gets billed on behalf of the client buyer; when the buyer field is populated, this field is the same value as buyer. Format : `buyers/{buyerAccountId}`
	Client string `json:"client,omitempty"` // Output only. Refers to a Client. Format: `buyers/{buyerAccountId}/clients/{clientAccountid}`
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the deal was last updated.
	Createtime string `json:"createTime,omitempty"` // Output only. The time of the deal creation.
	Description string `json:"description,omitempty"` // Output only. Free text description for the deal terms.
	Programmaticguaranteedterms ProgrammaticGuaranteedTerms `json:"programmaticGuaranteedTerms,omitempty"` // Pricing terms for Programmatic Guaranteed Deals.
	Preferreddealterms PreferredDealTerms `json:"preferredDealTerms,omitempty"` // Pricing terms for Preferred Deals.
	Creativerequirements CreativeRequirements `json:"creativeRequirements,omitempty"` // Message captures data about the creatives in the deal.
	Name string `json:"name,omitempty"` // Immutable. The unique identifier of the deal. Auto-generated by the server when a deal is created. Format: buyers/{accountId}/proposals/{proposalId}/deals/{dealId}
	Proposalrevision string `json:"proposalRevision,omitempty"` // Output only. The revision number for the proposal and is the same value as proposal.proposal_revision. Each update to deal causes the proposal revision number to auto-increment. The buyer keeps track of the last revision number they know of and pass it in when making an update. If the head revision number on the server has since incremented, then an ABORTED error is returned during the update operation to let the buyer know that a subsequent update was made.
	Sellertimezone TimeZone `json:"sellerTimeZone,omitempty"` // Represents a time zone from the [IANA Time Zone Database](https://www.iana.org/time-zones).
	Displayname string `json:"displayName,omitempty"` // Output only. The name of the deal. Maximum length of 255 unicode characters is allowed. Control characters are not allowed. Buyers cannot update this field. Note: Not to be confused with name, which is a unique identifier of the deal.
	Estimatedgrossspend Money `json:"estimatedGrossSpend,omitempty"` // Represents an amount of money with its currency type.
	Dealtype string `json:"dealType,omitempty"` // Output only. Type of deal.
}

// PreferredDealTerms represents the PreferredDealTerms schema from the OpenAPI specification
type PreferredDealTerms struct {
	Fixedprice Price `json:"fixedPrice,omitempty"` // Represents a price and a pricing type for a deal.
}

// Client represents the Client schema from the OpenAPI specification
type Client struct {
	Role string `json:"role,omitempty"` // Required. The role assigned to the client. Each role implies a set of permissions granted to the client.
	Sellervisible bool `json:"sellerVisible,omitempty"` // Whether the client will be visible to sellers.
	State string `json:"state,omitempty"` // Output only. The state of the client.
	Displayname string `json:"displayName,omitempty"` // Required. Display name shown to publishers. Must be unique for clients without partnerClientId specified. Maximum length of 255 characters is allowed.
	Name string `json:"name,omitempty"` // Output only. The resource name of the client. Format: `buyers/{accountId}/clients/{clientAccountId}`
	Partnerclientid string `json:"partnerClientId,omitempty"` // Arbitrary unique identifier provided by the buyer. This field can be used to associate a client with an identifier in the namespace of the buyer, lookup clients by that identifier and verify whether an Authorized Buyers account of the client already exists. If present, must be unique across all the clients.
}

// OperatingSystemTargeting represents the OperatingSystemTargeting schema from the OpenAPI specification
type OperatingSystemTargeting struct {
	Operatingsystemcriteria CriteriaTargeting `json:"operatingSystemCriteria,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
	Operatingsystemversioncriteria CriteriaTargeting `json:"operatingSystemVersionCriteria,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
}

// ListAuctionPackagesResponse represents the ListAuctionPackagesResponse schema from the OpenAPI specification
type ListAuctionPackagesResponse struct {
	Auctionpackages []AuctionPackage `json:"auctionPackages,omitempty"` // The list of auction packages.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Continuation token for fetching the next page of results. Pass this value in the ListAuctionPackagesRequest.pageToken field in the subsequent call to the `ListAuctionPackages` method to retrieve the next page of results.
}

// FrequencyCap represents the FrequencyCap schema from the OpenAPI specification
type FrequencyCap struct {
	Maximpressions int `json:"maxImpressions,omitempty"` // The maximum number of impressions that can be served to a user within the specified time period.
	Timeunittype string `json:"timeUnitType,omitempty"` // The time unit. Along with num_time_units defines the amount of time over which impressions per user are counted and capped.
	Timeunitscount int `json:"timeUnitsCount,omitempty"` // The amount of time, in the units specified by time_unit_type. Defines the amount of time over which impressions per user are counted and capped.
}

// Note represents the Note schema from the OpenAPI specification
type Note struct {
	Createtime string `json:"createTime,omitempty"` // Output only. When this note was created.
	Creatorrole string `json:"creatorRole,omitempty"` // Output only. The role who created the note.
	Note string `json:"note,omitempty"` // The text of the note. Maximum length is 1024 characters.
}

// ClientUser represents the ClientUser schema from the OpenAPI specification
type ClientUser struct {
	Name string `json:"name,omitempty"` // Output only. The resource name of the client user. Format: `buyers/{accountId}/clients/{clientAccountId}/users/{userId}`
	State string `json:"state,omitempty"` // Output only. The state of the client user.
	Email string `json:"email,omitempty"` // Required. The client user's email address that has to be unique across all users for the same client.
}

// DeactivateClientUserRequest represents the DeactivateClientUserRequest schema from the OpenAPI specification
type DeactivateClientUserRequest struct {
}

// ActivateClientRequest represents the ActivateClientRequest schema from the OpenAPI specification
type ActivateClientRequest struct {
}

// DayPartTargeting represents the DayPartTargeting schema from the OpenAPI specification
type DayPartTargeting struct {
	Dayparts []DayPart `json:"dayParts,omitempty"` // The targeted weekdays and times
	Timezonetype string `json:"timeZoneType,omitempty"` // The time zone type of the day parts
}

// MarketplaceTargeting represents the MarketplaceTargeting schema from the OpenAPI specification
type MarketplaceTargeting struct {
	Placementtargeting PlacementTargeting `json:"placementTargeting,omitempty"` // Represents targeting about where the ads can appear, for example, certain sites or mobile applications. Different placement targeting types will be logically OR'ed.
	Technologytargeting TechnologyTargeting `json:"technologyTargeting,omitempty"` // Represents targeting about various types of technology.
	Userlisttargeting CriteriaTargeting `json:"userListTargeting,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
	Videotargeting VideoTargeting `json:"videoTargeting,omitempty"` // Represents targeting information about video.
	Dayparttargeting DayPartTargeting `json:"daypartTargeting,omitempty"` // Represents Daypart targeting.
	Geotargeting CriteriaTargeting `json:"geoTargeting,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
	Inventorysizetargeting InventorySizeTargeting `json:"inventorySizeTargeting,omitempty"` // Represents the size of an ad unit that can be targeted on a bid request.
	Inventorytypetargeting InventoryTypeTargeting `json:"inventoryTypeTargeting,omitempty"` // Targeting of the inventory types a bid request can originate from.
}

// CriteriaTargeting represents the CriteriaTargeting schema from the OpenAPI specification
type CriteriaTargeting struct {
	Targetedcriteriaids []string `json:"targetedCriteriaIds,omitempty"` // A list of numeric IDs to be included.
	Excludedcriteriaids []string `json:"excludedCriteriaIds,omitempty"` // A list of numeric IDs to be excluded.
}

// RtbMetrics represents the RtbMetrics schema from the OpenAPI specification
type RtbMetrics struct {
	Adimpressions7days string `json:"adImpressions7Days,omitempty"` // Ad impressions in last 7 days.
	Bidrate7days float64 `json:"bidRate7Days,omitempty"` // Bid rate in last 7 days, calculated by (bids / bid requests).
	Bidrequests7days string `json:"bidRequests7Days,omitempty"` // Bid requests in last 7 days.
	Bids7days string `json:"bids7Days,omitempty"` // Bids in last 7 days.
	Filteredbidrate7days float64 `json:"filteredBidRate7Days,omitempty"` // Filtered bid rate in last 7 days, calculated by (filtered bids / bids).
	Mustbidratecurrentmonth float64 `json:"mustBidRateCurrentMonth,omitempty"` // Must bid rate for current month.
}

// TimeOfDay represents the TimeOfDay schema from the OpenAPI specification
type TimeOfDay struct {
	Minutes int `json:"minutes,omitempty"` // Minutes of hour of day. Must be from 0 to 59.
	Nanos int `json:"nanos,omitempty"` // Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Seconds int `json:"seconds,omitempty"` // Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	Hours int `json:"hours,omitempty"` // Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
}

// ProgrammaticGuaranteedTerms represents the ProgrammaticGuaranteedTerms schema from the OpenAPI specification
type ProgrammaticGuaranteedTerms struct {
	Guaranteedlooks string `json:"guaranteedLooks,omitempty"` // Count of guaranteed looks. For CPD deals, buyer changes to guaranteed_looks will be ignored.
	Impressioncap string `json:"impressionCap,omitempty"` // The lifetime impression cap for CPM Sponsorship deals. Deal will stop serving when cap is reached.
	Minimumdailylooks string `json:"minimumDailyLooks,omitempty"` // Daily minimum looks for CPD deal types. For CPD deals, buyer should negotiate on this field instead of guaranteed_looks.
	Percentshareofvoice string `json:"percentShareOfVoice,omitempty"` // For sponsorship deals, this is the percentage of the seller's eligible impressions that the deal will serve until the cap is reached. Valid value is within range 0~100.
	Reservationtype string `json:"reservationType,omitempty"` // The reservation type for a Programmatic Guaranteed deal. This indicates whether the number of impressions is fixed, or a percent of available impressions. If not specified, the default reservation type is STANDARD.
	Fixedprice Price `json:"fixedPrice,omitempty"` // Represents a price and a pricing type for a deal.
}

// UnsubscribeClientsRequest represents the UnsubscribeClientsRequest schema from the OpenAPI specification
type UnsubscribeClientsRequest struct {
	Clients []string `json:"clients,omitempty"` // Optional. A list of client buyers to unsubscribe from the auction package, with client buyer in the format `buyers/{accountId}/clients/{clientAccountId}`.
}

// CreativeRequirements represents the CreativeRequirements schema from the OpenAPI specification
type CreativeRequirements struct {
	Creativesafeframecompatibility string `json:"creativeSafeFrameCompatibility,omitempty"` // Output only. Specifies whether the creative is safeFrame compatible.
	Maxaddurationms string `json:"maxAdDurationMs,omitempty"` // Output only. The max duration of the video creative in milliseconds. only applicable for deals with video creatives.
	Programmaticcreativesource string `json:"programmaticCreativeSource,omitempty"` // Output only. Specifies the creative source for programmatic deals. PUBLISHER means creative is provided by seller and ADVERTISER means creative is provided by the buyer.
	Skippableadtype string `json:"skippableAdType,omitempty"` // Output only. Skippable video ads allow viewers to skip ads after 5 seconds. Only applicable for deals with video creatives.
	Creativeformat string `json:"creativeFormat,omitempty"` // Output only. The format of the creative, only applicable for programmatic guaranteed and preferred deals.
	Creativepreapprovalpolicy string `json:"creativePreApprovalPolicy,omitempty"` // Output only. Specifies the creative pre-approval policy.
}

// PublisherProfile represents the PublisherProfile schema from the OpenAPI specification
type PublisherProfile struct {
	Mobileapps []PublisherProfileMobileApplication `json:"mobileApps,omitempty"` // The list of apps represented in this publisher profile. Empty if this is a parent profile.
	Pitchstatement string `json:"pitchStatement,omitempty"` // Statement explaining what's unique about publisher's business, and why buyers should partner with the publisher.
	Logourl string `json:"logoUrl,omitempty"` // A Google public URL to the logo for this publisher profile. The logo is stored as a PNG, JPG, or GIF image.
	Overview string `json:"overview,omitempty"` // Overview of the publisher.
	Audiencedescription string `json:"audienceDescription,omitempty"` // Description on the publisher's audience.
	Domains []string `json:"domains,omitempty"` // The list of domains represented in this publisher profile. Empty if this is a parent profile. These are top private domains, meaning that these will not contain a string like "photos.google.co.uk/123", but will instead contain "google.co.uk". Can be used to filter the response of the publisherProfiles.list method.
	Mediakiturl string `json:"mediaKitUrl,omitempty"` // URL to additional marketing and sales materials.
	Topheadlines []string `json:"topHeadlines,omitempty"` // Up to three key metrics and rankings. For example, "#1 Mobile News Site for 20 Straight Months".
	Displayname string `json:"displayName,omitempty"` // Display name of the publisher profile. Can be used to filter the response of the publisherProfiles.list method.
	Publishercode string `json:"publisherCode,omitempty"` // A unique identifying code for the seller. This value is the same for all of the seller's parent and child publisher profiles. Can be used to filter the response of the publisherProfiles.list method.
	Name string `json:"name,omitempty"` // Name of the publisher profile. Format: `buyers/{buyer}/publisherProfiles/{publisher_profile}`
	Directdealscontact string `json:"directDealsContact,omitempty"` // Contact information for direct reservation deals. This is free text entered by the publisher and may include information like names, phone numbers and email addresses.
	Isparent bool `json:"isParent,omitempty"` // Indicates if this profile is the parent profile of the seller. A parent profile represents all the inventory from the seller, as opposed to child profile that is created to brand a portion of inventory. One seller has only one parent publisher profile, and can have multiple child profiles. See https://support.google.com/admanager/answer/6035806 for details. Can be used to filter the response of the publisherProfiles.list method by setting the filter to "is_parent: true".
	Programmaticdealscontact string `json:"programmaticDealsContact,omitempty"` // Contact information for programmatic deals. This is free text entered by the publisher and may include information like names, phone numbers and email addresses.
	Samplepageurl string `json:"samplePageUrl,omitempty"` // URL to a sample content page.
}

// InventorySizeTargeting represents the InventorySizeTargeting schema from the OpenAPI specification
type InventorySizeTargeting struct {
	Targetedinventorysizes []AdSize `json:"targetedInventorySizes,omitempty"` // A list of inventory sizes to be included.
	Excludedinventorysizes []AdSize `json:"excludedInventorySizes,omitempty"` // A list of inventory sizes to be excluded.
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Displayname string `json:"displayName,omitempty"` // The display_name of the contact.
	Email string `json:"email,omitempty"` // Email address for the contact.
}

// ResumeFinalizedDealRequest represents the ResumeFinalizedDealRequest schema from the OpenAPI specification
type ResumeFinalizedDealRequest struct {
}

// ListProposalsResponse represents the ListProposalsResponse schema from the OpenAPI specification
type ListProposalsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Continuation token for fetching the next page of results.
	Proposals []Proposal `json:"proposals,omitempty"` // The list of proposals.
}

// AddCreativeRequest represents the AddCreativeRequest schema from the OpenAPI specification
type AddCreativeRequest struct {
	Creative string `json:"creative,omitempty"` // Name of the creative to add to the finalized deal, in the format `buyers/{buyerAccountId}/creatives/{creativeId}`. See creative.name.
}

// ListFinalizedDealsResponse represents the ListFinalizedDealsResponse schema from the OpenAPI specification
type ListFinalizedDealsResponse struct {
	Finalizeddeals []FinalizedDeal `json:"finalizedDeals,omitempty"` // The list of finalized deals.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to fetch the next page of results.
}

// AddNoteRequest represents the AddNoteRequest schema from the OpenAPI specification
type AddNoteRequest struct {
	Note Note `json:"note,omitempty"` // A text note attached to the proposal to facilitate the communication between buyers and sellers.
}

// DayPart represents the DayPart schema from the OpenAPI specification
type DayPart struct {
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
	Dayofweek string `json:"dayOfWeek,omitempty"` // Day of week for the period.
	Endtime TimeOfDay `json:"endTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
}

// Price represents the Price schema from the OpenAPI specification
type Price struct {
	Amount Money `json:"amount,omitempty"` // Represents an amount of money with its currency type.
	TypeField string `json:"type,omitempty"` // The pricing type for the deal.
}

// SetReadyToServeRequest represents the SetReadyToServeRequest schema from the OpenAPI specification
type SetReadyToServeRequest struct {
}

// ListPublisherProfilesResponse represents the ListPublisherProfilesResponse schema from the OpenAPI specification
type ListPublisherProfilesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to fetch the next page of results.
	Publisherprofiles []PublisherProfile `json:"publisherProfiles,omitempty"` // The list of matching publisher profiles.
}

// DealPausingInfo represents the DealPausingInfo schema from the OpenAPI specification
type DealPausingInfo struct {
	Pausereason string `json:"pauseReason,omitempty"` // The reason for the pausing of the deal; empty for active deals.
	Pauserole string `json:"pauseRole,omitempty"` // The party that first paused the deal; unspecified for active deals.
	Pausingconsented bool `json:"pausingConsented,omitempty"` // Whether pausing is consented between buyer and seller for the deal.
}

// PlacementTargeting represents the PlacementTargeting schema from the OpenAPI specification
type PlacementTargeting struct {
	Mobileapplicationtargeting MobileApplicationTargeting `json:"mobileApplicationTargeting,omitempty"` // Mobile application targeting settings.
	Uritargeting UriTargeting `json:"uriTargeting,omitempty"` // Represents a list of targeted and excluded URLs (for example, google.com). For Private Auction Deals, URLs are either included or excluded. For Programmatic Guaranteed and Preferred Deals, this doesn't apply.
}

// FirstPartyMobileApplicationTargeting represents the FirstPartyMobileApplicationTargeting schema from the OpenAPI specification
type FirstPartyMobileApplicationTargeting struct {
	Targetedappids []string `json:"targetedAppIds,omitempty"` // A list of application IDs to be included.
	Excludedappids []string `json:"excludedAppIds,omitempty"` // A list of application IDs to be excluded.
}

// MobileApplicationTargeting represents the MobileApplicationTargeting schema from the OpenAPI specification
type MobileApplicationTargeting struct {
	Firstpartytargeting FirstPartyMobileApplicationTargeting `json:"firstPartyTargeting,omitempty"` // Represents a list of targeted and excluded mobile application IDs that publishers own. Android App ID, for example, com.google.android.apps.maps, can be found in Google Play Store URL. iOS App ID (which is a number) can be found at the end of iTunes store URL. First party mobile applications is either included or excluded.
}

// Money represents the Money schema from the OpenAPI specification
type Money struct {
	Currencycode string `json:"currencyCode,omitempty"` // The three-letter currency code defined in ISO 4217.
	Nanos int `json:"nanos,omitempty"` // Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Units string `json:"units,omitempty"` // The whole units of the amount. For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
}

// UnsubscribeAuctionPackageRequest represents the UnsubscribeAuctionPackageRequest schema from the OpenAPI specification
type UnsubscribeAuctionPackageRequest struct {
}

// ListClientsResponse represents the ListClientsResponse schema from the OpenAPI specification
type ListClientsResponse struct {
	Clients []Client `json:"clients,omitempty"` // The returned list of clients.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token to retrieve the next page of results. Pass this value in the ListClientsRequest.pageToken field in the subsequent call to the list method to retrieve the next page of results.
}

// BatchUpdateDealsRequest represents the BatchUpdateDealsRequest schema from the OpenAPI specification
type BatchUpdateDealsRequest struct {
	Requests []UpdateDealRequest `json:"requests,omitempty"` // Required. List of request messages to update deals.
}

// ListDealsResponse represents the ListDealsResponse schema from the OpenAPI specification
type ListDealsResponse struct {
	Deals []Deal `json:"deals,omitempty"` // The list of deals.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to fetch the next page of results.
}

// VideoTargeting represents the VideoTargeting schema from the OpenAPI specification
type VideoTargeting struct {
	Excludedpositiontypes []string `json:"excludedPositionTypes,omitempty"` // A list of video positions to be excluded. When this field is populated, the targeted_position_types field must be empty.
	Targetedpositiontypes []string `json:"targetedPositionTypes,omitempty"` // A list of video positions to be included. When this field is populated, the excluded_position_types field must be empty.
}

// CancelNegotiationRequest represents the CancelNegotiationRequest schema from the OpenAPI specification
type CancelNegotiationRequest struct {
}

// AuctionPackage represents the AuctionPackage schema from the OpenAPI specification
type AuctionPackage struct {
	Updatetime string `json:"updateTime,omitempty"` // Output only. Time the auction package was last updated. This value is only increased when this auction package is updated but never when a buyer subscribed.
	Createtime string `json:"createTime,omitempty"` // Output only. Time the auction package was created.
	Creator string `json:"creator,omitempty"` // Output only. The buyer that created this auction package. Format: `buyers/{buyerAccountId}`
	Description string `json:"description,omitempty"` // Output only. A description of the auction package.
	Displayname string `json:"displayName,omitempty"` // The display_name assigned to the auction package.
	Name string `json:"name,omitempty"` // Immutable. The unique identifier for the auction package. Format: `buyers/{accountId}/auctionPackages/{auctionPackageId}` The auction_package_id part of name is sent in the BidRequest to all RTB bidders and is returned as deal_id by the bidder in the BidResponse.
	Subscribedclients []string `json:"subscribedClients,omitempty"` // Output only. The list of clients of the current buyer that are subscribed to the AuctionPackage. Format: `buyers/{buyerAccountId}/clients/{clientAccountId}`
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// Proposal represents the Proposal schema from the OpenAPI specification
type Proposal struct {
	Notes []Note `json:"notes,omitempty"` // A list of notes from the buyer and the seller attached to this proposal.
	Originatorrole string `json:"originatorRole,omitempty"` // Output only. Indicates whether the buyer/seller created the proposal.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the proposal was last revised.
	Termsandconditions string `json:"termsAndConditions,omitempty"` // Output only. The terms and conditions associated with this proposal. Accepting a proposal implies acceptance of this field. This is created by the seller, the buyer can only view it.
	Buyer string `json:"buyer,omitempty"` // Output only. Refers to a buyer in The Realtime-bidding API. Format: `buyers/{buyerAccountId}`
	Buyerprivatedata PrivateData `json:"buyerPrivateData,omitempty"` // Buyers are allowed to store certain types of private data in a proposal or deal.
	Buyercontacts []Contact `json:"buyerContacts,omitempty"` // Contact information for the buyer.
	Publisherprofile string `json:"publisherProfile,omitempty"` // Immutable. Reference to the seller on the proposal. Format: `buyers/{buyerAccountId}/publisherProfiles/{publisherProfileId}` Note: This field may be set only when creating the resource. Modifying this field while updating the resource will result in an error.
	Sellercontacts []Contact `json:"sellerContacts,omitempty"` // Output only. Contact information for the seller.
	State string `json:"state,omitempty"` // Output only. Indicates the state of the proposal.
	Billedbuyer string `json:"billedBuyer,omitempty"` // Output only. When the client field is populated, this field refers to the buyer who creates and manages the client buyer and gets billed on behalf of the client buyer; when the buyer field is populated, this field is the same value as buyer. Format : `buyers/{buyerAccountId}`
	Client string `json:"client,omitempty"` // Output only. Refers to a Client. Format: `buyers/{buyerAccountId}/clients/{clientAccountid}`
	Proposalrevision string `json:"proposalRevision,omitempty"` // Output only. The revision number for the proposal. Each update to the proposal or deal causes the proposal revision number to auto-increment. The buyer keeps track of the last revision number they know of and pass it in when making an update. If the head revision number on the server has since incremented, then an ABORTED error is returned during the update operation to let the buyer know that a subsequent update was made.
	Displayname string `json:"displayName,omitempty"` // Output only. The descriptive name for the proposal. Maximum length of 255 unicode characters is allowed. Control characters are not allowed. Buyers cannot update this field. Note: Not to be confused with name, which is a unique identifier of the proposal.
	Pausingconsented bool `json:"pausingConsented,omitempty"` // Whether pausing is allowed for the proposal. This is a negotiable term between buyers and publishers.
	Dealtype string `json:"dealType,omitempty"` // Output only. Type of deal the proposal contains.
	Isrenegotiating bool `json:"isRenegotiating,omitempty"` // Output only. True if the proposal was previously finalized and is now being renegotiated.
	Lastupdaterorcommentorrole string `json:"lastUpdaterOrCommentorRole,omitempty"` // Output only. The role of the last user that either updated the proposal or left a comment.
	Name string `json:"name,omitempty"` // Immutable. The name of the proposal serving as a unique identifier. Format: buyers/{accountId}/proposals/{proposalId}
}

// UpdateDealRequest represents the UpdateDealRequest schema from the OpenAPI specification
type UpdateDealRequest struct {
	Deal Deal `json:"deal,omitempty"` // A deal represents a segment of inventory for displaying ads that contains the terms and targeting information that is used for serving as well as the deal stats and status. Note: A proposal may contain multiple deals.
	Updatemask string `json:"updateMask,omitempty"` // List of fields to be updated. If empty or unspecified, the service will update all fields populated in the update request excluding the output only fields and primitive fields with default value. Note that explicit field mask is required in order to reset a primitive field back to its default value, for example, false for boolean fields, 0 for integer fields. A special field mask consisting of a single path "*" can be used to indicate full replacement(the equivalent of PUT method), updatable fields unset or unspecified in the input will be cleared or set to default value. Output only fields will be ignored regardless of the value of updateMask.
}

// AdSize represents the AdSize schema from the OpenAPI specification
type AdSize struct {
	Height string `json:"height,omitempty"` // The height of the ad slot in pixels. This field will be present only when size type is `PIXEL`.
	TypeField string `json:"type,omitempty"` // The type of the ad slot size.
	Width string `json:"width,omitempty"` // The width of the ad slot in pixels. This field will be present only when size type is `PIXEL`.
}

// TimeZone represents the TimeZone schema from the OpenAPI specification
type TimeZone struct {
	Id string `json:"id,omitempty"` // IANA Time Zone Database time zone, e.g. "America/New_York".
	Version string `json:"version,omitempty"` // Optional. IANA Time Zone Database version number, e.g. "2019a".
}

// AcceptProposalRequest represents the AcceptProposalRequest schema from the OpenAPI specification
type AcceptProposalRequest struct {
	Proposalrevision string `json:"proposalRevision,omitempty"` // The last known client revision number of the proposal.
}

// SubscribeAuctionPackageRequest represents the SubscribeAuctionPackageRequest schema from the OpenAPI specification
type SubscribeAuctionPackageRequest struct {
}

// InventoryTypeTargeting represents the InventoryTypeTargeting schema from the OpenAPI specification
type InventoryTypeTargeting struct {
	Inventorytypes []string `json:"inventoryTypes,omitempty"` // The list of targeted inventory types for the bid request.
}

// SendRfpRequest represents the SendRfpRequest schema from the OpenAPI specification
type SendRfpRequest struct {
	Flightstarttime string `json:"flightStartTime,omitempty"` // Required. Proposed flight start time of the RFP. A timestamp in RFC3339 UTC "Zulu" format. Note that the specified value will be truncated to a granularity of one second.
	Inventorysizetargeting InventorySizeTargeting `json:"inventorySizeTargeting,omitempty"` // Represents the size of an ad unit that can be targeted on a bid request.
	Preferreddealterms PreferredDealTerms `json:"preferredDealTerms,omitempty"` // Pricing terms for Preferred Deals.
	Programmaticguaranteedterms ProgrammaticGuaranteedTerms `json:"programmaticGuaranteedTerms,omitempty"` // Pricing terms for Programmatic Guaranteed Deals.
	Flightendtime string `json:"flightEndTime,omitempty"` // Required. Proposed flight end time of the RFP. A timestamp in RFC3339 UTC "Zulu" format. Note that the specified value will be truncated to a granularity of one second.
	Estimatedgrossspend Money `json:"estimatedGrossSpend,omitempty"` // Represents an amount of money with its currency type.
	Geotargeting CriteriaTargeting `json:"geoTargeting,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
	Buyercontacts []Contact `json:"buyerContacts,omitempty"` // Contact information for the buyer.
	Note string `json:"note,omitempty"` // A message that is sent to the publisher. Maximum length is 1024 characters.
	Publisherprofile string `json:"publisherProfile,omitempty"` // Required. The profile of the publisher who will receive this RFP in the format: `buyers/{accountId}/publisherProfiles/{publisherProfileId}`.
	Client string `json:"client,omitempty"` // If the current buyer is sending the RFP on behalf of its client, use this field to specify the name of the client in the format: `buyers/{accountId}/clients/{clientAccountid}`.
	Displayname string `json:"displayName,omitempty"` // Required. The display name of the proposal being created by this RFP.
}

// PrivateData represents the PrivateData schema from the OpenAPI specification
type PrivateData struct {
	Referenceid string `json:"referenceId,omitempty"` // A buyer specified reference ID. This can be queried in the list operations (max-length: 1024 unicode code units).
}

// TechnologyTargeting represents the TechnologyTargeting schema from the OpenAPI specification
type TechnologyTargeting struct {
	Operatingsystemtargeting OperatingSystemTargeting `json:"operatingSystemTargeting,omitempty"` // Represents targeting information for operating systems.
	Devicecapabilitytargeting CriteriaTargeting `json:"deviceCapabilityTargeting,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
	Devicecategorytargeting CriteriaTargeting `json:"deviceCategoryTargeting,omitempty"` // Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.
}

// DeactivateClientRequest represents the DeactivateClientRequest schema from the OpenAPI specification
type DeactivateClientRequest struct {
}

// PauseFinalizedDealRequest represents the PauseFinalizedDealRequest schema from the OpenAPI specification
type PauseFinalizedDealRequest struct {
	Reason string `json:"reason,omitempty"` // The reason to pause the finalized deal, will be displayed to the seller. Maximum length is 1000 characters.
}
