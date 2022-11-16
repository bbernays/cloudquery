// Code generated by codegen; DO NOT EDIT.

package ecs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterTasks() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_cluster_tasks",
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html`,
		Resolver:    fetchEcsClusterTasks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "task_protection",
				Type:     schema.TypeJSON,
				Resolver: getEcsTaskProtection,
			},
			{
				Name:     "attachments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attachments"),
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "capacity_provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CapacityProviderName"),
			},
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
			},
			{
				Name:     "connectivity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Connectivity"),
			},
			{
				Name:     "connectivity_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ConnectivityAt"),
			},
			{
				Name:     "container_instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerInstanceArn"),
			},
			{
				Name:     "containers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Containers"),
			},
			{
				Name:     "cpu",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Cpu"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "desired_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DesiredStatus"),
			},
			{
				Name:     "enable_execute_command",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableExecuteCommand"),
			},
			{
				Name:     "ephemeral_storage",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EphemeralStorage"),
			},
			{
				Name:     "execution_stopped_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ExecutionStoppedAt"),
			},
			{
				Name:     "group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Group"),
			},
			{
				Name:     "health_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthStatus"),
			},
			{
				Name:     "inference_accelerators",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InferenceAccelerators"),
			},
			{
				Name:     "last_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastStatus"),
			},
			{
				Name:     "launch_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchType"),
			},
			{
				Name:     "memory",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Memory"),
			},
			{
				Name:     "overrides",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Overrides"),
			},
			{
				Name:     "platform_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformFamily"),
			},
			{
				Name:     "platform_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformVersion"),
			},
			{
				Name:     "pull_started_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PullStartedAt"),
			},
			{
				Name:     "pull_stopped_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PullStoppedAt"),
			},
			{
				Name:     "started_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartedAt"),
			},
			{
				Name:     "started_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StartedBy"),
			},
			{
				Name:     "stop_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StopCode"),
			},
			{
				Name:     "stopped_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StoppedAt"),
			},
			{
				Name:     "stopped_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StoppedReason"),
			},
			{
				Name:     "stopping_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StoppingAt"),
			},
			{
				Name:     "task_definition_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskDefinitionArn"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}
