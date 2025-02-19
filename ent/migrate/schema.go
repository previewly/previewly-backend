// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ErrorResultsColumns holds the columns for the "error_results" table.
	ErrorResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "message", Type: field.TypeString, Nullable: true},
		{Name: "url_errorresult", Type: field.TypeInt, Nullable: true},
	}
	// ErrorResultsTable holds the schema information for the "error_results" table.
	ErrorResultsTable = &schema.Table{
		Name:       "error_results",
		Columns:    ErrorResultsColumns,
		PrimaryKey: []*schema.Column{ErrorResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "error_results_urls_errorresult",
				Columns:    []*schema.Column{ErrorResultsColumns[3]},
				RefColumns: []*schema.Column{UrlsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "filename", Type: field.TypeString},
		{Name: "destination_path", Type: field.TypeString},
		{Name: "original_filename", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "extra_value", Type: field.TypeString, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
	}
	// ImageProcessesColumns holds the columns for the "image_processes" table.
	ImageProcessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"success", "error", "pending"}},
		{Name: "process_hash", Type: field.TypeString},
		{Name: "processes", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "path_prefix", Type: field.TypeString, Nullable: true},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "image_imageprocess", Type: field.TypeInt, Nullable: true},
	}
	// ImageProcessesTable holds the schema information for the "image_processes" table.
	ImageProcessesTable = &schema.Table{
		Name:       "image_processes",
		Columns:    ImageProcessesColumns,
		PrimaryKey: []*schema.Column{ImageProcessesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "image_processes_images_imageprocess",
				Columns:    []*schema.Column{ImageProcessesColumns[8]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "imageprocess_process_hash_image_imageprocess",
				Unique:  true,
				Columns: []*schema.Column{ImageProcessesColumns[2], ImageProcessesColumns[8]},
			},
		},
	}
	// StatsColumns holds the columns for the "stats" table.
	StatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "stat_image", Type: field.TypeInt},
		{Name: "url_stat", Type: field.TypeInt, Nullable: true},
	}
	// StatsTable holds the schema information for the "stats" table.
	StatsTable = &schema.Table{
		Name:       "stats",
		Columns:    StatsColumns,
		PrimaryKey: []*schema.Column{StatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stats_images_image",
				Columns:    []*schema.Column{StatsColumns[3]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "stats_urls_stat",
				Columns:    []*schema.Column{StatsColumns[4]},
				RefColumns: []*schema.Column{UrlsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString, Unique: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// UrlsColumns holds the columns for the "urls" table.
	UrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"success", "error", "pending"}},
		{Name: "relative_path", Type: field.TypeString, Nullable: true},
	}
	// UrlsTable holds the schema information for the "urls" table.
	UrlsTable = &schema.Table{
		Name:       "urls",
		Columns:    UrlsColumns,
		PrimaryKey: []*schema.Column{UrlsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ErrorResultsTable,
		ImagesTable,
		ImageProcessesTable,
		StatsTable,
		TokensTable,
		UrlsTable,
	}
)

func init() {
	ErrorResultsTable.ForeignKeys[0].RefTable = UrlsTable
	ImageProcessesTable.ForeignKeys[0].RefTable = ImagesTable
	StatsTable.ForeignKeys[0].RefTable = ImagesTable
	StatsTable.ForeignKeys[1].RefTable = UrlsTable
}
