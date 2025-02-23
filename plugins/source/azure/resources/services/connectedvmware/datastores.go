package connectedvmware

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Datastores() *schema.Table {
	return &schema.Table{
		Name:        "azure_connectedvmware_datastores",
		Resolver:    fetchDatastores,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware@v0.1.0#Datastore",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_connectedvmware_datastores", client.Namespacemicrosoft_connectedvmwarevsphere),
		Transform:   transformers.TransformWithStruct(&armconnectedvmware.Datastore{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDatastores(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconnectedvmware.NewDatastoresClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
