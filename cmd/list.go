package cmd

import (
	"context"
	"log"

	"azure-resource-inspector/internal/azure"

	"github.com/spf13/cobra"
)

var subscriptionID string

var listcmd = &cobra.Command{
	Use:   "list",
	Short: "List all resources in the specified Azure subscription",
	Long:  `This command lists all resources in the specified Azure subscription using the Azure SDK for Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cred, err := azure.GetAzureCredential(ctx)
		if err != nil {
			log.Fatalf("Failed to get Azure credentials: %v", err)
		}

		if err := azure.ListResources(ctx, cred, subscriptionID); err != nil {
			log.Fatalf("Failed to list resources: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(listcmd)

	listcmd.Flags().StringVarP(&subscriptionID, "subscription", "s", "", "Azure Subscription ID (required)")
	listcmd.MarkFlagRequired("subscription")
}
