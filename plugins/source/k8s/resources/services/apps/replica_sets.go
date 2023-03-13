package apps

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ReplicaSets() *schema.Table {
	return &schema.Table{
		Name:      "k8s_apps_replica_sets",
		Resolver:  fetchReplicaSets,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.ReplicaSet{}),
		Columns: []schema.Column{
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: client.ResolveContext,
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchReplicaSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().AppsV1().ReplicaSets("")
	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
