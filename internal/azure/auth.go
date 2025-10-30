package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func GetAzureCredential(ctx context.Context) (*azidentity.DefaultAzureCredential, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain Azure credential: %w", err)
	}
	return cred, nil
}
