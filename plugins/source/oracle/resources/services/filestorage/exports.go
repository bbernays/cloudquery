package filestorage

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/filestorage"
)

func Exports() *schema.Table {
	return &schema.Table{
		Name:      "oracle_filestorage_exports",
		Resolver:  fetchExports,
		Multiplex: client.AvailibilityDomainCompartmentMultiplex,
		Transform: client.TransformWithStruct(&filestorage.ExportSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
