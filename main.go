package main

import (
	similigo_api "github.com/Ojelaidi/similigo-api/cmd/similigo-api"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{Use: "similigo"}
	rootCmd.AddCommand(similigo_api.Start())
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
