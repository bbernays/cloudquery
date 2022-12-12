// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"

func Armdesktopvirtualization() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdesktopvirtualization.NewHostPoolsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/hostPools",
			Namespace:      "Microsoft.DesktopVirtualization",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DesktopVirtualization)`,
			Pager:          `NewListPager`,
			ResponseStruct: "HostPoolsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdesktopvirtualization())
}
