package models

import "time"

type RightMoveResponse struct {
	Properties []struct {
		ID                   int    `json:"id"`
		Bedrooms             int    `json:"bedrooms"`
		Bathrooms            int    `json:"bathrooms"`
		NumberOfImages       int    `json:"numberOfImages"`
		NumberOfFloorplans   int    `json:"numberOfFloorplans"`
		NumberOfVirtualTours int    `json:"numberOfVirtualTours"`
		Summary              string `json:"summary"`
		DisplayAddress       string `json:"displayAddress"`
		CountryCode          string `json:"countryCode"`
		Location             struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		PropertyImages struct {
			Images []struct {
				SrcURL  string `json:"srcUrl"`
				URL     string `json:"url"`
				Caption string `json:"caption"`
			} `json:"images"`
			MainImageSrc    string `json:"mainImageSrc"`
			MainMapImageSrc string `json:"mainMapImageSrc"`
		} `json:"propertyImages"`
		PropertySubType string `json:"propertySubType"`
		ListingUpdate   struct {
			ListingUpdateReason string    `json:"listingUpdateReason"`
			ListingUpdateDate   time.Time `json:"listingUpdateDate"`
		} `json:"listingUpdate"`
		PremiumListing   bool `json:"premiumListing"`
		FeaturedProperty bool `json:"featuredProperty"`
		Price            struct {
			Amount        int    `json:"amount"`
			Frequency     string `json:"frequency"`
			CurrencyCode  string `json:"currencyCode"`
			DisplayPrices []struct {
				DisplayPrice          string `json:"displayPrice"`
				DisplayPriceQualifier string `json:"displayPriceQualifier"`
			} `json:"displayPrices"`
		} `json:"price"`
		Customer struct {
			BranchID              int    `json:"branchId"`
			BrandPlusLogoURI      string `json:"brandPlusLogoURI"`
			ContactTelephone      string `json:"contactTelephone"`
			BranchDisplayName     string `json:"branchDisplayName"`
			BranchName            string `json:"branchName"`
			BrandTradingName      string `json:"brandTradingName"`
			BranchLandingPageURL  string `json:"branchLandingPageUrl"`
			Development           bool   `json:"development"`
			ShowReducedProperties bool   `json:"showReducedProperties"`
			Commercial            bool   `json:"commercial"`
			ShowOnMap             bool   `json:"showOnMap"`
			EnhancedListing       bool   `json:"enhancedListing"`
			DevelopmentContent    any    `json:"developmentContent"`
			BuildToRent           bool   `json:"buildToRent"`
			BuildToRentBenefits   []any  `json:"buildToRentBenefits"`
			BrandPlusLogoURL      string `json:"brandPlusLogoUrl"`
		} `json:"customer"`
		Distance        any    `json:"distance"`
		TransactionType string `json:"transactionType"`
		ProductLabel    struct {
			ProductLabelText string `json:"productLabelText"`
			SpotlightLabel   bool   `json:"spotlightLabel"`
		} `json:"productLabel"`
		Commercial              bool      `json:"commercial"`
		Development             bool      `json:"development"`
		Residential             bool      `json:"residential"`
		Students                bool      `json:"students"`
		Auction                 bool      `json:"auction"`
		FeesApply               bool      `json:"feesApply"`
		FeesApplyText           string    `json:"feesApplyText"`
		DisplaySize             string    `json:"displaySize"`
		ShowOnMap               bool      `json:"showOnMap"`
		PropertyURL             string    `json:"propertyUrl"`
		ContactURL              string    `json:"contactUrl"`
		StaticMapURL            any       `json:"staticMapUrl"`
		Channel                 string    `json:"channel"`
		FirstVisibleDate        time.Time `json:"firstVisibleDate"`
		Keywords                []any     `json:"keywords"`
		KeywordMatchType        string    `json:"keywordMatchType"`
		Saved                   bool      `json:"saved"`
		Hidden                  bool      `json:"hidden"`
		OnlineViewingsAvailable bool      `json:"onlineViewingsAvailable"`
		LozengeModel            struct {
			MatchingLozenges []any `json:"matchingLozenges"`
		} `json:"lozengeModel"`
		HasBrandPlus                bool   `json:"hasBrandPlus"`
		DisplayStatus               string `json:"displayStatus"`
		EnquiredTimestamp           any    `json:"enquiredTimestamp"`
		PropertyTypeFullDescription string `json:"propertyTypeFullDescription"`
		Heading                     string `json:"heading"`
		FormattedBranchName         string `json:"formattedBranchName"`
		AddedOrReduced              string `json:"addedOrReduced"`
		IsRecent                    bool   `json:"isRecent"`
		FormattedDistance           string `json:"formattedDistance"`
		EnhancedListing             bool   `json:"enhancedListing"`
	} `json:"properties"`
	ResultCount                 string `json:"resultCount"`
	SearchParametersDescription string `json:"searchParametersDescription"`
	RadiusOptions               []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"radiusOptions"`
	PriceOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"priceOptions"`
	BedroomOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"bedroomOptions"`
	AddedToSiteOptions []struct {
		Value       any    `json:"value"`
		Description string `json:"description"`
	} `json:"addedToSiteOptions"`
	MustHaveOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"mustHaveOptions"`
	DontShowOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"dontShowOptions"`
	FurnishOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"furnishOptions"`
	LetTypeOptions []struct {
		Value       any    `json:"value"`
		Description string `json:"description"`
	} `json:"letTypeOptions"`
	SortOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"sortOptions"`
	StaticMapURL             string `json:"staticMapUrl"`
	ShortLocationDescription string `json:"shortLocationDescription"`
	Timestamp                int64  `json:"timestamp"`
	Bot                      bool   `json:"bot"`
	DeviceType               string `json:"deviceType"`
	PropertySchema           struct {
		ID                   int `json:"id"`
		Bedrooms             int `json:"bedrooms"`
		Bathrooms            int `json:"bathrooms"`
		NumberOfImages       int `json:"numberOfImages"`
		NumberOfFloorplans   int `json:"numberOfFloorplans"`
		NumberOfVirtualTours int `json:"numberOfVirtualTours"`
		Summary              any `json:"summary"`
		DisplayAddress       any `json:"displayAddress"`
		CountryCode          any `json:"countryCode"`
		Location             struct {
			Latitude  any `json:"latitude"`
			Longitude any `json:"longitude"`
		} `json:"location"`
		PropertyImages struct {
			Images          []any  `json:"images"`
			MainImageSrc    string `json:"mainImageSrc"`
			MainMapImageSrc string `json:"mainMapImageSrc"`
		} `json:"propertyImages"`
		PropertySubType any `json:"propertySubType"`
		ListingUpdate   struct {
			ListingUpdateReason any       `json:"listingUpdateReason"`
			ListingUpdateDate   time.Time `json:"listingUpdateDate"`
		} `json:"listingUpdate"`
		PremiumListing   bool `json:"premiumListing"`
		FeaturedProperty bool `json:"featuredProperty"`
		Price            struct {
			Amount        int    `json:"amount"`
			Frequency     string `json:"frequency"`
			CurrencyCode  string `json:"currencyCode"`
			DisplayPrices []struct {
				DisplayPrice          string `json:"displayPrice"`
				DisplayPriceQualifier string `json:"displayPriceQualifier"`
			} `json:"displayPrices"`
		} `json:"price"`
		Customer struct {
			BranchID              any    `json:"branchId"`
			BrandPlusLogoURI      any    `json:"brandPlusLogoURI"`
			ContactTelephone      any    `json:"contactTelephone"`
			BranchDisplayName     any    `json:"branchDisplayName"`
			BranchName            any    `json:"branchName"`
			BrandTradingName      any    `json:"brandTradingName"`
			BranchLandingPageURL  any    `json:"branchLandingPageUrl"`
			Development           bool   `json:"development"`
			ShowReducedProperties bool   `json:"showReducedProperties"`
			Commercial            bool   `json:"commercial"`
			ShowOnMap             bool   `json:"showOnMap"`
			EnhancedListing       bool   `json:"enhancedListing"`
			DevelopmentContent    any    `json:"developmentContent"`
			BuildToRent           bool   `json:"buildToRent"`
			BuildToRentBenefits   any    `json:"buildToRentBenefits"`
			BrandPlusLogoURL      string `json:"brandPlusLogoUrl"`
		} `json:"customer"`
		Distance        any `json:"distance"`
		TransactionType any `json:"transactionType"`
		ProductLabel    struct {
			ProductLabelText any  `json:"productLabelText"`
			SpotlightLabel   bool `json:"spotlightLabel"`
		} `json:"productLabel"`
		Commercial              bool      `json:"commercial"`
		Development             bool      `json:"development"`
		Residential             bool      `json:"residential"`
		Students                bool      `json:"students"`
		Auction                 bool      `json:"auction"`
		FeesApply               bool      `json:"feesApply"`
		FeesApplyText           string    `json:"feesApplyText"`
		DisplaySize             string    `json:"displaySize"`
		ShowOnMap               bool      `json:"showOnMap"`
		PropertyURL             string    `json:"propertyUrl"`
		ContactURL              string    `json:"contactUrl"`
		StaticMapURL            string    `json:"staticMapUrl"`
		Channel                 string    `json:"channel"`
		FirstVisibleDate        time.Time `json:"firstVisibleDate"`
		Keywords                []any     `json:"keywords"`
		KeywordMatchType        string    `json:"keywordMatchType"`
		Saved                   bool      `json:"saved"`
		Hidden                  bool      `json:"hidden"`
		OnlineViewingsAvailable bool      `json:"onlineViewingsAvailable"`
		LozengeModel            struct {
			MatchingLozenges []any `json:"matchingLozenges"`
		} `json:"lozengeModel"`
		HasBrandPlus                bool   `json:"hasBrandPlus"`
		DisplayStatus               string `json:"displayStatus"`
		EnquiredTimestamp           any    `json:"enquiredTimestamp"`
		PropertyTypeFullDescription string `json:"propertyTypeFullDescription"`
		Heading                     string `json:"heading"`
		FormattedBranchName         string `json:"formattedBranchName"`
		AddedOrReduced              string `json:"addedOrReduced"`
		IsRecent                    bool   `json:"isRecent"`
		FormattedDistance           string `json:"formattedDistance"`
		EnhancedListing             bool   `json:"enhancedListing"`
	} `json:"propertySchema"`
	SidebarModel struct {
		SoldHousePricesLinks struct {
			Heading    string `json:"heading"`
			SubHeading string `json:"subHeading"`
			Model      []struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"model"`
			HeadingLink any `json:"headingLink"`
		} `json:"soldHousePricesLinks"`
		RelatedHouseSearches struct {
			Heading    string `json:"heading"`
			SubHeading any    `json:"subHeading"`
			Model      []struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"model"`
			HeadingLink any `json:"headingLink"`
		} `json:"relatedHouseSearches"`
		RelatedFlatSearches struct {
			Heading    string `json:"heading"`
			SubHeading any    `json:"subHeading"`
			Model      []struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"model"`
			HeadingLink struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"headingLink"`
		} `json:"relatedFlatSearches"`
		RelatedPopularSearches struct {
			Heading    string `json:"heading"`
			SubHeading any    `json:"subHeading"`
			Model      []struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"model"`
			HeadingLink any `json:"headingLink"`
		} `json:"relatedPopularSearches"`
		RelatedRegionsSearches any `json:"relatedRegionsSearches"`
		ChannelSwitchLink      struct {
			Heading    string `json:"heading"`
			SubHeading any    `json:"subHeading"`
			Model      []struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"model"`
			HeadingLink any `json:"headingLink"`
		} `json:"channelSwitchLink"`
		RelatedStudentLinks any `json:"relatedStudentLinks"`
		BranchMPU           any `json:"branchMPU"`
		CountryGuideMPU     any `json:"countryGuideMPU"`
		SuggestedLinks      struct {
			Heading    string `json:"heading"`
			SubHeading any    `json:"subHeading"`
			Model      []struct {
				Text     string `json:"text"`
				URL      string `json:"url"`
				NoFollow bool   `json:"noFollow"`
			} `json:"model"`
			HeadingLink any `json:"headingLink"`
		} `json:"suggestedLinks"`
	} `json:"sidebarModel"`
	SeoModel struct {
		CanonicalURL string `json:"canonicalUrl"`
		MetaRobots   string `json:"metaRobots"`
	} `json:"seoModel"`
	MapViewURL        string `json:"mapViewUrl"`
	LegacyURL         string `json:"legacyUrl"`
	ListViewURL       string `json:"listViewUrl"`
	PageTitle         string `json:"pageTitle"`
	MetaDescription   string `json:"metaDescription"`
	RecentSearchModel struct {
		LinkDisplayText                 string `json:"linkDisplayText"`
		TitleDisplayText                string `json:"titleDisplayText"`
		SearchCriteriaMobile            string `json:"searchCriteriaMobile"`
		CreateDate                      int64  `json:"createDate"`
		LocationIdentifierAndSearchType string `json:"locationIdentifierAndSearchType"`
	} `json:"recentSearchModel"`
	MaxCardsPerPage     int    `json:"maxCardsPerPage"`
	CountryCode         string `json:"countryCode"`
	CountryID           int    `json:"countryId"`
	CurrencyCodeOptions []struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"currencyCodeOptions"`
	AreaSizeUnitOptions []struct {
		Value        string `json:"value"`
		Description  string `json:"description"`
		Abbreviation string `json:"abbreviation"`
	} `json:"areaSizeUnitOptions"`
	SizeOptions []struct {
		Value        string `json:"value"`
		Description  string `json:"description"`
		Abbreviation string `json:"abbreviation"`
	} `json:"sizeOptions"`
	PriceTypeOptions []struct {
		Value        string `json:"value"`
		Description  string `json:"description"`
		Abbreviation string `json:"abbreviation"`
	} `json:"priceTypeOptions"`
	ShowFeaturedAgent      bool   `json:"showFeaturedAgent"`
	ShowNewDrawASearch     bool   `json:"showNewDrawASearch"`
	CommercialChannel      bool   `json:"commercialChannel"`
	DisambiguationPagePath string `json:"disambiguationPagePath"`
	DfpModel               struct {
		SidebarSlots []struct {
			ID         string  `json:"id"`
			AdUnitPath string  `json:"adUnitPath"`
			Sizes      [][]int `json:"sizes"`
			Mappings   []any   `json:"mappings"`
		} `json:"sidebarSlots"`
		Targeting []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"targeting"`
	} `json:"dfpModel"`
	NoResultsModel struct {
		SuggestionPods        []any `json:"suggestionPods"`
		IntelligentSuggestion any   `json:"intelligentSuggestion"`
	} `json:"noResultsModel"`
	URLPath        any    `json:"urlPath"`
	TileGeometry   any    `json:"tileGeometry"`
	Comscore       string `json:"comscore"`
	CookiePolicies struct {
		Functional        bool `json:"functional"`
		Targeting         bool `json:"targeting"`
		StrictlyNecessary bool `json:"strictly necessary"`
	} `json:"cookiePolicies"`
	FormattedExchangeRateDate string `json:"formattedExchangeRateDate"`
	Authenticated             bool   `json:"authenticated"`
	ApplicationProperties     struct {
		DfpInterstitial2AdUnitPath                          string `json:"dfp.interstitial2.adUnitPath"`
		GtmAuth                                             string `json:"gtm.auth"`
		SidebarMpuAdUnitTimeout                             string `json:"sidebar.mpu.adUnitTimeout"`
		DfpInterstitial1AdUnitPath                          string `json:"dfp.interstitial1.adUnitPath"`
		SeoGlobalConsumerFooterURL                          string `json:"seo.global.consumerFooter.url"`
		DfpSoldByMeInterstitialAdUnitPath                   string `json:"dfp.sold-by-me.interstitial.adUnitPath"`
		SidebarMpuAdUnitPath                                string `json:"sidebar.mpu.adUnitPath"`
		DfpSidebarUkCreditCheckSponsorshipAdUnitPath        string `json:"dfp.sidebar.uk-credit-check-sponsorship.adUnitPath"`
		LocationProductWebServerPort                        string `json:"location.product.web.server.port"`
		RightmovePlusLandingPageURL                         string `json:"rightmovePlusLandingPageUrl"`
		LabsHostname                                        string `json:"labsHostname"`
		OptimizeMapPins                                     bool   `json:"optimize.map.pins"`
		DfpSoldByMeInterstitialCommercialLettingsAdUnitPath string `json:"dfp.sold-by-me.interstitial.commercial-lettings.adUnitPath"`
		DfpSoldByMeInterstitialCommercialSalesAdUnitPath    string `json:"dfp.sold-by-me.interstitial.commercial-sales.adUnitPath"`
		MetadataServiceURL                                  string `json:"metadata.service.url"`
		AnalyticsTypeformURL                                string `json:"analytics.typeform.url"`
		PublicsiteServerPort                                string `json:"publicsite.server.port"`
		LocationProductWebServerHost                        string `json:"location.product.web.server.host"`
		GtmPreview                                          string `json:"gtm.preview"`
		MyRightmoveWebServerPort                            string `json:"my.rightmove.web.server.port"`
		InfoBuildVersion                                    string `json:"info.build.version"`
		LocationSearchServerPort                            string `json:"location.search.server.port"`
		RecaptchaSiteKey                                    string `json:"recaptchaSiteKey"`
		PublicsiteServerHost                                string `json:"publicsite.server.host"`
		ClickstreamEnabled                                  bool   `json:"clickstream.enabled"`
		MyRightmoveWebServerHost                            string `json:"my.rightmove.web.server.host"`
		LocationSearchServerHost                            string `json:"location.search.server.host"`
		GtmEnabled                                          bool   `json:"gtm.enabled"`
		RecaptchaEnabled                                    bool   `json:"recaptchaEnabled"`
		DfpOverseasInterstitial3AdUnitPath                  string `json:"dfp.overseas.interstitial3.adUnitPath"`
		SeoGlobalConsumerHeaderURL                          string `json:"seo.global.consumerHeader.url"`
		MediaServerHost                                     string `json:"media.server.host"`
		DfpOverseasInterstitial2AdUnitPath                  string `json:"dfp.overseas.interstitial2.adUnitPath"`
		GtmID                                               string `json:"gtm.id"`
		ComscoreEnabled                                     bool   `json:"comscore.enabled"`
		DfpOverseasInterstitial1AdUnitPath                  string `json:"dfp.overseas.interstitial1.adUnitPath"`
		DfpInterstitial3AdUnitPath                          string `json:"dfp.interstitial3.adUnitPath"`
	} `json:"applicationProperties"`
	TermsOfUse string `json:"termsOfUse"`
	Location   struct {
		ID               int    `json:"id"`
		DisplayName      string `json:"displayName"`
		ShortDisplayName string `json:"shortDisplayName"`
		LocationType     string `json:"locationType"`
		ListingCurrency  string `json:"listingCurrency"`
	} `json:"location"`
	LocationBranchSearchPath any `json:"locationBranchSearchPath"`
	SearchParameters         struct {
		LocationIdentifier        string   `json:"locationIdentifier"`
		MaxBedrooms               string   `json:"maxBedrooms"`
		MinBedrooms               string   `json:"minBedrooms"`
		NumberOfPropertiesPerPage string   `json:"numberOfPropertiesPerPage"`
		Radius                    string   `json:"radius"`
		SortType                  string   `json:"sortType"`
		Index                     string   `json:"index"`
		PropertyTypes             []any    `json:"propertyTypes"`
		MaxDaysSinceAdded         string   `json:"maxDaysSinceAdded"`
		IncludeLetAgreed          bool     `json:"includeLetAgreed"`
		ViewType                  string   `json:"viewType"`
		MustHave                  []any    `json:"mustHave"`
		DontShow                  []string `json:"dontShow"`
		FurnishTypes              []any    `json:"furnishTypes"`
		Channel                   string   `json:"channel"`
		AreaSizeUnit              string   `json:"areaSizeUnit"`
		CurrencyCode              string   `json:"currencyCode"`
		Keywords                  []any    `json:"keywords"`
	} `json:"searchParameters"`
	Pagination struct {
		Total   int `json:"total"`
		Options []struct {
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"options"`
		First string `json:"first"`
		Last  string `json:"last"`
		Page  string `json:"page"`
	} `json:"pagination"`
	FeatureSwitchStateForUser struct {
		IndividualFeatureSwitchStates []struct {
			Label     string `json:"label"`
			State     string `json:"state"`
			ShouldLog bool   `json:"shouldLog"`
		} `json:"individualFeatureSwitchStates"`
		FeatureUser struct {
			UniqueIdentifier string `json:"uniqueIdentifier"`
		} `json:"featureUser"`
	} `json:"featureSwitchStateForUser"`
}

type ZooplaResponse struct {
	PageProps struct {
		InitialState struct {
			Favourites struct {
				Listings []any `json:"listings"`
			} `json:"favourites"`
			DfpAdTargeting struct {
				Activity           string   `json:"activity"`
				AreaName           string   `json:"area_name"`
				BedsMax            string   `json:"beds_max"`
				BedsMin            string   `json:"beds_min"`
				Brand              string   `json:"brand"`
				BrandPrimary       string   `json:"brand_primary"`
				CountryCode        string   `json:"country_code"`
				CountyAreaName     string   `json:"county_area_name"`
				CurrencyCode       string   `json:"currency_code"`
				ListingsCategory   string   `json:"listings_category"`
				Outcode            string   `json:"outcode"`
				Outcodes           []string `json:"outcodes"`
				Page               string   `json:"page"`
				PostalArea         string   `json:"postal_area"`
				PriceMax           string   `json:"price_max"`
				PriceMin           string   `json:"price_min"`
				RegionName         string   `json:"region_name"`
				SearchLocation     string   `json:"search_location"`
				SearchResultsCount string   `json:"search_results_count"`
				Section            string   `json:"section"`
				TotalResults       string   `json:"total_results"`
				URL                string   `json:"url"`
				ViewType           string   `json:"view_type"`
			} `json:"dfpAdTargeting"`
			SeoAccordions []struct {
				Category     string `json:"category"`
				Geo          string `json:"geo"`
				PropertyType string `json:"propertyType"`
				Rows         []struct {
					Links []struct {
						Category        string `json:"category"`
						Geo             string `json:"geo"`
						PropertyType    string `json:"propertyType"`
						Section         string `json:"section"`
						TransactionType string `json:"transactionType"`
						URI             string `json:"uri"`
					} `json:"links"`
				} `json:"rows"`
				Section         string `json:"section"`
				TransactionType string `json:"transactionType"`
				Links           any    `json:"links"`
				Name            string `json:"name"`
				Expanded        bool   `json:"expanded"`
			} `json:"seoAccordions"`
			Links struct {
			} `json:"links"`
			CurrentSearchSource            string `json:"currentSearchSource"`
			SearchType                     string `json:"searchType"`
			SearchCategory                 string `json:"searchCategory"`
			TransactionType                string `json:"transactionType"`
			PageTitle                      string `json:"pageTitle"`
			ShowLocationValidation         bool   `json:"showLocationValidation"`
			UserAlertID                    string `json:"userAlertId"`
			SavedSearchAndAlerts           any    `json:"savedSearchAndAlerts"`
			PreviousSavedSearchURI         any    `json:"previousSavedSearchUri"`
			IsEditSavedSearchBannerVisible bool   `json:"isEditSavedSearchBannerVisible"`
			IsEditSavedSearchSuccess       any    `json:"isEditSavedSearchSuccess"`
			ActionAfterAuth                any    `json:"actionAfterAuth"`
			TotalResults                   int    `json:"totalResults"`
			PageNumber                     int    `json:"pageNumber"`
			PageNumberMax                  int    `json:"pageNumberMax"`
			TotalResultsWasLimited         bool   `json:"totalResultsWasLimited"`
			Polyenc                        any    `json:"polyenc"`
			LocationByIdentifier           string `json:"locationByIdentifier"`
		} `json:"initialState"`
		RegularListingsFormatted []struct {
			Address                       string `json:"address"`
			AlternativeRentFrequencyLabel string `json:"alternativeRentFrequencyLabel"`
			AvailableFrom                 string `json:"availableFrom"`
			Branch                        struct {
				BranchID         int    `json:"branchId"`
				BranchDetailsURI string `json:"branchDetailsUri"`
				LogoURL          string `json:"logoUrl"`
				Name             string `json:"name"`
				Phone            string `json:"phone"`
			} `json:"branch"`
			DisplayType  string `json:"displayType"`
			FeaturedType any    `json:"featuredType"`
			Features     []struct {
				IconID  string `json:"iconId"`
				Content int    `json:"content"`
			} `json:"features"`
			Flag       string `json:"flag"`
			Highlights []any  `json:"highlights"`
			Image      struct {
				Src               string `json:"src"`
				ResponsiveImgList []struct {
					Width int    `json:"width"`
					Src   string `json:"src"`
				} `json:"responsiveImgList"`
				Caption string `json:"caption"`
			} `json:"image"`
			IsPremium         bool   `json:"isPremium"`
			LastPublishedDate string `json:"lastPublishedDate"`
			ListingID         string `json:"listingId"`
			ListingUris       struct {
				Contact string `json:"contact"`
				Detail  string `json:"detail"`
				Success string `json:"success"`
			} `json:"listingUris"`
			Location struct {
				Coordinates struct {
					IsApproximate bool    `json:"isApproximate"`
					Latitude      float64 `json:"latitude"`
					Longitude     float64 `json:"longitude"`
				} `json:"coordinates"`
			} `json:"location"`
			NumberOfFloorPlans int     `json:"numberOfFloorPlans"`
			NumberOfImages     int     `json:"numberOfImages"`
			NumberOfVideos     int     `json:"numberOfVideos"`
			Price              string  `json:"price"`
			PriceDrop          any     `json:"priceDrop"`
			PriceTitle         string  `json:"priceTitle"`
			PropertyType       string  `json:"propertyType"`
			PublishedOn        string  `json:"publishedOn"`
			PublishedOnLabel   string  `json:"publishedOnLabel"`
			ShortPriceTitle    string  `json:"shortPriceTitle"`
			SummaryDescription string  `json:"summaryDescription"`
			Tags               []any   `json:"tags"`
			Title              string  `json:"title"`
			Transports         []any   `json:"transports"`
			UnderOffer         bool    `json:"underOffer"`
			AvailableFromLabel string  `json:"availableFromLabel"`
			IsFavourite        bool    `json:"isFavourite"`
			Gallery            [][]any `json:"gallery"`
		} `json:"regularListingsFormatted"`
		ExtendedListingsFormatted []any `json:"extendedListingsFormatted"`
		FeaturedProperties        []struct {
			Address                       string `json:"address"`
			AlternativeRentFrequencyLabel string `json:"alternativeRentFrequencyLabel"`
			AvailableFrom                 string `json:"availableFrom"`
			Branch                        struct {
				BranchID         int    `json:"branchId"`
				BranchDetailsURI string `json:"branchDetailsUri"`
				LogoURL          string `json:"logoUrl"`
				Name             string `json:"name"`
				Phone            string `json:"phone"`
			} `json:"branch"`
			DisplayType  string `json:"displayType"`
			FeaturedType string `json:"featuredType"`
			Features     []struct {
				IconID  string `json:"iconId"`
				Content int    `json:"content"`
			} `json:"features"`
			Flag       string `json:"flag"`
			Highlights []any  `json:"highlights"`
			Image      struct {
				Src               string `json:"src"`
				ResponsiveImgList []struct {
					Width int    `json:"width"`
					Src   string `json:"src"`
				} `json:"responsiveImgList"`
				Caption any `json:"caption"`
			} `json:"image"`
			IsPremium         bool   `json:"isPremium"`
			LastPublishedDate string `json:"lastPublishedDate"`
			ListingID         string `json:"listingId"`
			ListingUris       struct {
				Contact string `json:"contact"`
				Detail  string `json:"detail"`
				Success string `json:"success"`
			} `json:"listingUris"`
			Location struct {
				Coordinates struct {
					IsApproximate bool    `json:"isApproximate"`
					Latitude      float64 `json:"latitude"`
					Longitude     float64 `json:"longitude"`
				} `json:"coordinates"`
			} `json:"location"`
			NumberOfFloorPlans int    `json:"numberOfFloorPlans"`
			NumberOfImages     int    `json:"numberOfImages"`
			NumberOfVideos     int    `json:"numberOfVideos"`
			Price              string `json:"price"`
			PriceDrop          struct {
				FirstPriceDate        string `json:"firstPriceDate"`
				LastPriceChangeDate   string `json:"lastPriceChangeDate"`
				PercentageChangeLabel any    `json:"percentageChangeLabel"`
			} `json:"priceDrop"`
			PriceTitle         string `json:"priceTitle"`
			PropertyType       string `json:"propertyType"`
			PublishedOn        string `json:"publishedOn"`
			PublishedOnLabel   string `json:"publishedOnLabel"`
			ShortPriceTitle    string `json:"shortPriceTitle"`
			SummaryDescription string `json:"summaryDescription"`
			Tags               []any  `json:"tags"`
			Title              string `json:"title"`
			Transports         []any  `json:"transports"`
		} `json:"featuredProperties"`
		TransactionType   any `json:"transactionType"`
		SearchResultsMeta struct {
			AnalyticsTaxonomy struct {
				Activity             string   `json:"activity"`
				AreaName             string   `json:"areaName"`
				BedsMax              string   `json:"bedsMax"`
				BedsMin              string   `json:"bedsMin"`
				Brand                string   `json:"brand"`
				CountryCode          string   `json:"countryCode"`
				CountyAreaName       string   `json:"countyAreaName"`
				CurrencyCode         string   `json:"currencyCode"`
				ExpandedResultsCount int      `json:"expandedResultsCount"`
				GeoIdentifier        string   `json:"geoIdentifier"`
				ListingsCategory     string   `json:"listingsCategory"`
				Outcode              string   `json:"outcode"`
				Outcodes             []string `json:"outcodes"`
				Page                 string   `json:"page"`
				PostalArea           string   `json:"postalArea"`
				PriceMax             any      `json:"priceMax"`
				PriceMin             any      `json:"priceMin"`
				Radius               any      `json:"radius"`
				RegionName           string   `json:"regionName"`
				ResultsSort          string   `json:"resultsSort"`
				SearchGUID           string   `json:"searchGuid"`
				SearchIdentifier     string   `json:"searchIdentifier"`
				SearchLocation       string   `json:"searchLocation"`
				SearchResultsCount   int      `json:"searchResultsCount"`
				Section              string   `json:"section"`
				TotalResults         int      `json:"totalResults"`
				URL                  string   `json:"url"`
				ViewType             string   `json:"viewType"`
			} `json:"analyticsTaxonomy"`
			AnalyticsEcommerce struct {
				CurrencyCode string `json:"currencyCode"`
				Impressions  []struct {
					ID       int    `json:"id"`
					List     string `json:"list"`
					Position int    `json:"position"`
					Variant  string `json:"variant"`
				} `json:"impressions"`
			} `json:"analyticsEcommerce"`
			MetaInfo struct {
				Title        string `json:"title"`
				CanonicalURI string `json:"canonicalUri"`
				Description  string `json:"description"`
			} `json:"metaInfo"`
			SeoBlurb string `json:"seoBlurb"`
			Title    string `json:"title"`
		} `json:"searchResultsMeta"`
		Breadcrumbs []struct {
			Label string `json:"label"`
			URI   string `json:"uri"`
		} `json:"breadcrumbs"`
		Links struct {
			CreateAlert any `json:"createAlert"`
			SaveSearch  any `json:"saveSearch"`
		} `json:"links"`
		SimilarGeoIds []struct {
			SimilarGeoIds []struct {
				Identifier string `json:"identifier"`
				Label      string `json:"label"`
			} `json:"similarGeoIds"`
			GeoID string `json:"geoId"`
		} `json:"similarGeoIds"`
		OptimizelyState []string `json:"__OPTIMIZELY_STATE__"`
	} `json:"pageProps"`
	NSsp bool `json:"__N_SSP"`
}

type NewPropertyNotification struct {
	Title    string
	Image    string
	Price    string
	Url      string
	Website  string
	Location string
}
