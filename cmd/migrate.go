package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jurevic/facegrinder/pkg/datastore/migrate"
	"github.com/jurevic/facegrinder/pkg/datastore"
)

var reset bool

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "uses Facegrinder migration tool to migrate Postgres db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		argsMig := args[:0]
		for _, arg := range args {
			switch arg {
			case "migrate", "--db_debug", "--reset":
			default:
				argsMig = append(argsMig, arg)
			}
		}

		if reset {
			migrate.Reset()
		}
		migrate.Migrate(argsMig)
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().BoolVar(
		&reset,
		"reset",
		false,
		"migrate down to version 0 then up to latest. WARNING: all data will be lost!")

	datastore.Init()
}