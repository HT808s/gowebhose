package webhose

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Webhose struct {
	Token      string
	Parameters map[string]string
}

type Response struct {
	Posts []struct {
		Thread struct {
			Uuid              string      `json:"uuid"`
			Url               string      `json:"url"`
			SiteFull          string      `json:"site_full"`
			Site              string      `json:"site"`
			SiteSection       string      `json:"site_section"`
			SectionTitle      string      `json:"section_title"`
			Title             string      `json:"title"`
			TitleFull         string      `json:"title_full"`
			Published         string      `json:"published,omitempty"`
			RepliesCount      int         `json:"replies_count,omitempty"`
			ParticipantsCount int         `json:"participants_count,omitempty"`
			SiteType          string      `json:"site_type,omitempty"`
			Country           string      `json:"country"`
			SpamScore         float32     `json:"spam_score,omitempty"`
			MainImage         string      `json:"main_image"`
			PerfScore         int         `json:"preformance_score,omitempty"`
			Social            interface{} `json:"social,omitempty"`
			DomainRank        int         `json:"domain_rank,omitempty"`
		} `json:"thread"`
		Id             int      `json:"id"`
		Uuid           string   `json:"uuid"`
		Url            string   `json:"url"`
		OrdInThread    int      `json:"ord_in_thread"`
		Author         string   `json:"author"`
		Published      string   `json:"published"`
		Title          string   `json:"title"`
		Text           string   `json:"text"`
		HighLightText  string   `json:"highlightText"`
		HighLightTitle string   `json:"highlightTitle"`
		Language       string   `json:"language"`
		ExternalLinks  []string `json:"external_links"`
		Entities       struct {
			Persons []struct {
				Name      string `json:"name, omitempty"`
				Sentiment string `json:"sentiment, omitempty"`
			} `json:"persons, omitempty"`
			Locations []struct {
				Name      string `json:"name, omitempty"`
				Sentiment string `json:"sentiment, omitempty"`
			} `json:"locations, omitempty"`
			Organizations []struct {
				Name      string `json:"name, omitempty"`
				Sentiment string `json:"sentiment, omitempty"`
			} `json:"organizations, omitempty"`
		} `json:"entities, omitempty"`
		Crawled string `json:"crawled"`
	} `json:"posts"`
}

const (
	Language = "language"
	SiteType = "site_type"
	Site     = "site"
)

func Search(input string, wb Webhose) (*Response, error) {
	var baseUrl = url.URL{
		Host:   "webhose.io",
		Scheme: "https",
		Path:   "search",
	}
	q := baseUrl.Query()
	q.Set("token", wb.Token)
	for key, param := range wb.Parameters {
		q.Set(key, param)
	}
	q.Set("q", input)
	baseUrl.RawQuery = q.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= len(response.Posts); i++ {
		response.Posts[i-1].Id = i
	}
	return &response, nil
}

func (r *Response) EmberID() {
	var i int = 1
	for _, post := range r.Posts {
		post.Id = i
		i++
	}
}

func (r *Response) String() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}
