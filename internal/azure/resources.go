package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/charmbracelet/lipgloss"
)

func ListResources(ctx context.Context, cred azcore.TokenCredential, subscriptionID string) error {
	client, err := armresources.NewClient(subscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create Azure Resource Manager client: %w", err)
	}

	pager := client.NewListPager(nil)

	// Step 1: Collect data into slices
	var names, types, locations []string

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to retrieve page: %w", err)
		}

		for _, r := range page.Value {
			names = append(names, safeStr(r.Name))
			types = append(types, safeStr(r.Type))
			locations = append(locations, safeStr(r.Location))
		}
	}

	nameW, typeW, locW := maxWidth(names), maxWidth(types), maxWidth(locations)

	nameW += 2
	typeW += 2
	locW += 2

	headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF"))
	separator := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).
		Render("─" + repeatRune('─', nameW+typeW+locW+4))

	plainName := fmt.Sprintf("%-*s", nameW, "NAME")
	plainType := fmt.Sprintf("%-*s", typeW, "TYPE")
	plainLoc := fmt.Sprintf("%-*s", locW, "LOCATION")

	nameHeader := headerStyle.Render(plainName)
	typeHeader := headerStyle.Render(plainType)
	locHeader := headerStyle.Render(plainLoc)

	fmt.Println(separator)
	fmt.Printf("%s%s%s\n", nameHeader, typeHeader, locHeader)
	fmt.Println(separator)

	for i := range names {
		fmt.Printf("%-*s %-*s %-*s\n", nameW, names[i], typeW, types[i], locW, locations[i])
	}

	fmt.Println(separator)
	return nil
}

func safeStr(s *string) string {
	if s == nil {
		return "-"
	}
	return *s
}

func maxWidth(list []string) int {
	max := 0
	for _, s := range list {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

func repeatRune(r rune, count int) string {
	out := make([]rune, count)
	for i := range out {
		out[i] = r
	}
	return string(out)
}
