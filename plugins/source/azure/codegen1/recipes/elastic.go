// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"

func Armelastic() []*Table {
	tables := []*Table{
		{
			NewFunc:        armelastic.NewMonitorsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Elastic/monitors",
			Namespace:      "Microsoft.Elastic",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Elastic)`,
			Pager:          `NewListPager`,
			ResponseStruct: "MonitorsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armelastic())
}
