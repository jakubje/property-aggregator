package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"github.com/jakubje/property-aggregator/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/url"
	"strings"
)

func (s *PropertyScraperProcessor) zooplaInit() {
	req, err := http.NewRequest("GET", "https://www.zoopla.co.uk/", nil)
	if err != nil {
		log.Fatal().Err(err)
	}
	req.Header.Set("authority", "www.zoopla.co.uk")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "en-GB,en;q=0.9")
	req.Header.Set("referer", "https://www.google.com/")
	req.Header.Set("sec-ch-ua", `"Opera";v="103", "Not;A=Brand";v="8", "Chromium";v="117"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 OPR/103.0.0.0")
	resp, err := s.client.Do(req)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal().Msg("failed zoopla init")
		return
	}
	//fmt.Println(string(bodyText))
	buildId := strings.Split(strings.Split(string(bodyText), `"buildId":"`)[1], `",`)[0]
	s.getZooplaData(buildId)

}

func (s *PropertyScraperProcessor) getZooplaData(buildId string) {
	log.Info().Msg("getting zoopla data")
	requestUrl := fmt.Sprintf("https://www.zoopla.co.uk/_next/data/%s/to-rent/property/reading.json", buildId)
	params := url.Values{}
	params.Add("search_source", "home")
	params.Add("beds_min", "1")
	params.Add("beds_max", "2")

	params.Add("price_frequency", "per_month")
	params.Add("available_from", "3months")
	params.Add("results_sort", "newest_listings")
	params.Add("is_retirement_home", "false")
	params.Add("added", "24_hours")
	//params.Add("added", "3_days")
	//params.Add("furnished_state", "unfurnished")
	params.Add("is_shared_accommodation", "false")
	params.Add("q", "Reading, Berkshire")
	params.Add("search-path", "property")
	params.Add("search-path", "reading")
	constructedURL := requestUrl + "?" + params.Encode()
	req, err := http.NewRequest("GET", constructedURL, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	req.Header.Set("Host", "www.zoopla.co.uk")
	req.Header.Set("sec-ch-ua", `"Opera";v="103", "Not;A=Brand";v="8", "Chromium";v="117"`)
	req.Header.Set("x-nextjs-data", "1")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 OPR/103.0.0.0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://www.zoopla.co.uk/to-rent/property")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	resp, err := s.client.Do(req)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err)
	}

	if resp.StatusCode == 200 {
		var responseData models.ZooplaResponse
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			log.Fatal().Err(err).Msg("error decoding json")
		}
		var docs []interface{}
		for _, property := range responseData.PageProps.RegularListingsFormatted {

			exists, err := s.ExistsInDB("listingid", property.ListingID)
			if err != nil {
				log.Fatal().Err(err)
			}
			if !exists {
				var images []any
				for _, image := range property.Gallery {
					images = append(images, fmt.Sprintf("https://lid.zoocdn.com/645/430%s", image[1]))
				}
				dbDoc := DBStructure{
					Tag:           "Zoopla",
					AvailableFrom: property.AvailableFrom,
					ListingID:     property.ListingID,
					ListingUrl:    fmt.Sprintf("https://www.zoopla.co.uk%s", property.ListingUris.Detail),
					Price:         property.Price,
					PropertyType:  property.PropertyType,
					Title:         property.Title,
					Summary:       property.SummaryDescription,
					PublishedOn:   property.PublishedOn,
					Images:        images,
				}
				fmt.Println(images)
				discordData := models.NewPropertyNotification{
					Title:    property.Title,
					Image:    fmt.Sprintf("%s", images[0]),
					Price:    property.Price,
					Url:      fmt.Sprintf("https://www.zoopla.co.uk%s", property.ListingUris.Detail),
					Website:  "Zoopla",
					Location: "",
				}
				docs = append(docs, dbDoc)

				sendWebhook(s.config, &discordData)

			}
		}
		if len(docs) == 0 {
			return
		}

		s.mongoCollection.InsertMany(context.Background(), docs)
		log.Info().Msg("Saved to db")
	} else {
		log.Fatal().Msg(fmt.Sprintf("Request failed: %d", resp.StatusCode))
	}

}

func (s *PropertyScraperProcessor) ExistsInDB(fieldId string, fieldValue string) (bool, error) {
	filter := bson.M{fieldId: fieldValue}
	count, err := s.mongoCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
