package cmd

import (
	"github.com/August-Brandt/repet/config"
	"github.com/August-Brandt/repet/path"
	petSync "github.com/August-Brandt/repet/sync"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync snippets",
	Long:  `Sync snippets with gist/gitlab`,
	RunE:  sync,
}

func sync(cmd *cobra.Command, args []string) (err error) {
	filePath, err := path.NewAbsolutePath(config.Conf.General.SnippetFile)
	return petSync.AutoSync(filePath)
}

func init() {
	RootCmd.AddCommand(syncCmd)
}
