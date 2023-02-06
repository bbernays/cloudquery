package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEcsAccountSettingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcsClient(ctrl)
	services := client.Services{
		Ecs: m,
	}
	setting := types.Setting{}
	err := faker.FakeObject(&setting)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAccountSettings(gomock.Any(), gomock.Any(), gomock.Any()).Return(ecs.ListAccountSettingsOutput{
		Settings: []types.Setting{setting},
	}, nil)

	return services
}

func TestEcsAccountSettings(t *testing.T) {
	client.AwsMockTestHelper(t, AccountSettings(), buildEcsAccountSettingsMock, client.TestOptions{})
}
