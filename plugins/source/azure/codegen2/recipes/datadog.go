// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"

func init() {
	tables := []Table{
		{
			Service:        "armdatadog",
			Name:           "marketplace_agreements",
			Struct:         &armdatadog.AgreementResource{},
			ResponseStruct: &armdatadog.MarketplaceAgreementsClientListResponse{},
			Client:         &armdatadog.MarketplaceAgreementsClient{},
			ListFunc:       (&armdatadog.MarketplaceAgreementsClient{}).NewListPager,
			NewFunc:        armdatadog.NewMarketplaceAgreementsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/agreements",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Datadog)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armdatadog",
			Name:           "monitors",
			Struct:         &armdatadog.MonitorResource{},
			ResponseStruct: &armdatadog.MonitorsClientListResponse{},
			Client:         &armdatadog.MonitorsClient{},
			ListFunc:       (&armdatadog.MonitorsClient{}).NewListPager,
			NewFunc:        armdatadog.NewMonitorsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/monitors",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Datadog)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
