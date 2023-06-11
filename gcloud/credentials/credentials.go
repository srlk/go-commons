package credentials

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

// NewCredentials returns credentials as option.ClientOption
// credentials can be provided as a file or json
func NewCredentials(credentials string) (option.ClientOption, error) {
	if _, err := os.Stat(credentials); err != nil {
		credentials = strings.Trim(credentials, " ")
		var tmp map[string]string
		if err := json.Unmarshal([]byte(credentials), &tmp); err != nil {
			return nil, errors.Wrap(err, "credentials is non-json")
		}
		return option.WithCredentialsJSON([]byte(credentials)), nil
	} else {
		return option.WithCredentialsFile(credentials), nil
	}
}
