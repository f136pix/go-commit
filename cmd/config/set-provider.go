package config

import (
	"fmt"
	"go-commit/internal/configstore"

	"go-commit/cmd/env"

	"github.com/spf13/cobra"
)

func isValidProvider(provider string) bool {
	for _, p := range env.ValidProviders {
		if p == provider {
			return true
		}
	}
}

var setProviderCmd = &cobra.Command{
	Use:   "set-provider <provider-name> --api-key <your-api-key>",
	Short: "Set the API key for a provider",
	Long:  `Set the API key for a provider and store it in the configuration file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		providerName := args[0]

		if !isValidProvider(providerName) {
			fmt.Printf("Invalid provider: %s\n", providerName)
			return
		}

		apiKey, _ := cmd.Flags().GetString("api-key")

		if apiKey == "" {
			fmt.Println("API key is required")
			return
		}

		config := configstore.LoadConfig()
		config.Providers[providerName] = apiKey
		configstore.SaveConfig(config)

		fmt.Printf("API key for provider '%s' set successfully\n", providerName)
	},
}

func init() {

	// -k flag
	setProviderCmd.Flags().StringP("api-key", "k", "", "API key for the provider")
	setProviderCmd.MarkFlagRequired("api-key")

	ConfigCmd.AddCommand(setProviderCmd)
}
