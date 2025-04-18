package dbmlgengomodel

import (
	"fmt"

	"github.com/duythinht/dbml-go/internal/gen-go-model/gen"
	"github.com/spf13/cobra"
)

var (
	from             string
	out              string
	gopackage        string
	fieldtags        []string
	shouldGenTblName bool
	rememberAlias    bool
	recursive        bool
	exclude          string
)

var DBMLGenGoModelCmd = &cobra.Command{
	Use:   "dbml-gen-go-model",
	Short: "Generate Go models from DBML files",
	Long: `A CLI tool to generate Go model structs from DBML (Database Markup Language) files.
Supports generating models from single files or directories of DBML files.`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.Generate(gen.Opts{
			From:             from,
			Out:              out,
			Package:          gopackage,
			FieldTags:        fieldtags,
			ShouldGenTblName: shouldGenTblName,
			RememberAlias:    rememberAlias,
			Recursive:        recursive,
			Exclude:          exclude,
		})
	},
}

func init() {
	DBMLGenGoModelCmd.Flags().
		StringVarP(&from, "from", "f", "database.dbml", "source of dbml, can be https://dbdiagram.io/... | fire_name.dbml")
	DBMLGenGoModelCmd.Flags().StringVarP(&out, "out", "o", "model", "output folder")
	DBMLGenGoModelCmd.Flags().
		StringVarP(&gopackage, "package", "p", "model", "package name for generated files")
	DBMLGenGoModelCmd.Flags().
		StringArrayVarP(&fieldtags, "fieldtags", "t", []string{"db", "json", "mapstructure"}, "go field tags to generate")
	DBMLGenGoModelCmd.Flags().
		BoolVar(&shouldGenTblName, "gen-table-name", false, "generate \"TableName\" function for models")
	DBMLGenGoModelCmd.Flags().
		BoolVar(&rememberAlias, "remember-alias", false, "remember table alias (only when 'from' is a directory)")
	DBMLGenGoModelCmd.Flags().
		BoolVar(&recursive, "recursive", false, "recursively search directories (only when 'from' is a directory)")
	DBMLGenGoModelCmd.Flags().
		StringVarP(&exclude, "exclude", "E", "", "regex pattern to exclude files (only when 'from' is a directory)")

	DBMLGenGoModelCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dbml-gen-go-model v1.0.0")
		},
	})
}
