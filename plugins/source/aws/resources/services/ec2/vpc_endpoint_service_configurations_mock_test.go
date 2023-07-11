package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2VpcEndpointServiceConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	sc := types.ServiceConfiguration{}
	require.NoError(t, faker.FakeObject(&sc))

	m.EXPECT().DescribeVpcEndpointServiceConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointServiceConfigurationsOutput{
			ServiceConfigurations: []types.ServiceConfiguration{sc},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcEndpointServiceConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpointServiceConfigurations(), buildEc2VpcEndpointServiceConfigurations, client.TestOptions{})
}
