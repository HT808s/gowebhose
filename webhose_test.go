package webhose

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestSearch(t *testing.T) {
	resp, err := Search("tesla", Webhose{
		Token: "4cf8c9c1-0ccf-41d3-b530-806f4cab2d02",
		Parameters: map[string]string{
			"language": "english",
		},
	})
	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
}
