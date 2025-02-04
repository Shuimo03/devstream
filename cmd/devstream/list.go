package main

import (
	"github.com/spf13/cobra"

	"github.com/devstream-io/devstream/cmd/devstream/list"
)

var (
	pluginFilter string
)

var listCMD = &cobra.Command{
	Use:   "list",
	Short: "This command only supports listing plugins now",
}

var listPluginsCMD = &cobra.Command{
	Use:   "plugins",
	Short: "List all plugins",
	Long: `This command lists all of the plugins.
Examples:
  dtm list plugins
  dtm list plugins --filter=argo.*
  dtm list plugins -r ^argo
`,
	Run: listPluginsCMDFunc,
}

func listPluginsCMDFunc(_ *cobra.Command, _ []string) {
	list.List(pluginFilter)
}

// TODO Use `--group=somegroup` to filter the specified groups on feature
func init() {
	listCMD.AddCommand(listPluginsCMD)

	listPluginsCMD.PersistentFlags().StringVarP(&pluginFilter, "filter", "r", "", "filter plugin by regex")
}
