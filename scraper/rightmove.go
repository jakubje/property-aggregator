package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"github.com/jakubje/property-aggregator/models"
	"github.com/rs/zerolog/log"
	"io"
	"net/url"
	"strconv"
)

const (
	baseUrl = "https://www.rightmove.co.uk/api/_search"
)

func (s *PropertyScraperProcessor) getRightMoveData() {
	log.Info().Msg("getting rightmove data")
	requestUrl := baseUrl
	params := url.Values{}
	params.Add("locationIdentifier", "REGION^1114")
	params.Add("maxBedrooms", "2")
	params.Add("numberOfPropertiesPerPage", "100")
	params.Add("radius", "0.0")
	params.Add("sortType", "6")
	params.Add("index", "0")
	params.Add("maxDaysSinceAdded", "1")
	params.Add("includeLetAgreed", "false")
	params.Add("viewType", "LIST")
	params.Add("dontShow[0]", "houseShare")
	params.Add("dontShow[1]", "retirement")
	params.Add("dontShow[2]", "student")
	params.Add("channel", "RENT")
	params.Add("areaSizeUnit", "sqft")
	params.Add("currencyCode", "GBP")
	params.Add("isFetching", "false")
	constructedURL := requestUrl + "?" + params.Encode()

	req, err := http.NewRequest("GET", constructedURL, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	req.Header.Set("Host", "www.rightmove.co.uk")
	req.Header.Set("sec-ch-ua", `"Opera";v="103", "Not;A=Brand";v="8", "Chromium";v="117"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 OPR/103.0.0.0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.rightmove.co.uk/property-to-rent/find.html?locationIdentifier=REGION%5E1114&maxBedrooms=2&minBedrooms=1&propertyTypes=&maxDaysSinceAdded=1&includeLetAgreed=false&mustHave=&dontShow=houseShare%2Cretirement%2Cstudent&furnishTypes=&keywords=")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	resp, err := s.client.Do(req)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err)
	}

	var responseData models.RightMoveResponse

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatal().Err(err).Msg("error decoding json")
	}
	var docs []interface{}

	for _, property := range responseData.Properties {
		exists, err := s.ExistsInDB("id", string(property.ID))
		if err != nil {
			log.Fatal().Err(err)
		}
		if !exists {
			var images []any
			for _, image := range property.PropertyImages.Images {
				images = append(images, image.SrcURL)
			}

			dbDoc := DBStructure{
				Tag:          "Rightmove",
				ListingID:    strconv.Itoa(property.ID),
				ListingUrl:   fmt.Sprintf("https://www.rightmove.co.uk%s", property.PropertyURL),
				Price:        strconv.Itoa(property.Price.Amount),
				PropertyType: property.PropertySubType,
				Bedrooms:     strconv.Itoa(property.Bedrooms),
				Bathrooms:    strconv.Itoa(property.Bathrooms),
				Title:        property.PropertyTypeFullDescription,
				Summary:      property.Summary,
				PublishedOn:  property.ListingUpdate.ListingUpdateDate.String(),
				Images:       images,
			}
			docs = append(docs, dbDoc)
		}
	}
	if len(docs) == 0 {
		return
	}
	s.mongoCollection.InsertMany(context.Background(), docs)
}

type DBStructure struct {
	Tag           string `json:"tag"`
	AvailableFrom string `json:"availableFrom"`
	ListingID     string `json:"listingId"`
	ListingUrl    string `json:"listingUrl"`
	Price         string `json:"price"`
	PropertyType  string `json:"propertyType"`
	Title         string `json:"title"`
	Bedrooms      string `json:"bedrooms"`
	Bathrooms     string `json:"bathrooms"`
	Summary       string `json:"summary"`
	PublishedOn   string `json:"publishedOn"`
	Images        []any  `json:"images"`
}
