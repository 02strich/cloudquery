package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Customconversions() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_customconversions",
		Resolver:  fetchCustomconversions,
		Transform: client.TransformWithStruct(&rest.Customconversion{}, transformers.WithPrimaryKeys("Id")),
	}
}
