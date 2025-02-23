package apikeys

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	apikeys "cloud.google.com/go/apikeys/apiv2"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_apikeys_keys",
		Description: `https://cloud.google.com/api-keys/docs/reference/rest/v2/projects.locations.keys#Key`,
		Resolver:    fetchKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("apikeys.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Key{}, transformers.WithPrimaryKeys("Uid")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListKeysRequest{
		Parent: "projects/" + c.ProjectId + "/locations/global",
	}
	gcpClient, err := apikeys.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListKeys(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
