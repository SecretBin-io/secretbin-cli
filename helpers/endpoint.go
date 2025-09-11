package helpers

import (
	"github.com/spf13/viper"

	secretbin "github.com/Nihility-io/SecretBin-Go/v2"
)

// SetEndpoint sets the SecretBin endpoint in the configuration.
// It validates the endpoint by attempting to create a new SecretBin client with it.
// If the endpoint is valid, it saves it to the configuration file.
func SetEndpoint(endpoint string) error {
	if _, err := secretbin.New(endpoint); err != nil {
		return err
	}
	viper.Set("endpoint", endpoint)
	viper.SafeWriteConfig()
	viper.WriteConfig()

	return nil
}
