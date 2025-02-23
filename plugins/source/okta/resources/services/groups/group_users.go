package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupUsers() *schema.Table {
	return &schema.Table{
		Name:      "okta_group_users",
		Resolver:  fetchGroupUsers,
		Transform: client.TransformWithStruct(&models.GroupUser{}),
		Columns: []schema.Column{
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
