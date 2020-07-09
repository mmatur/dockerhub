package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmatur/dockerhub/choose"
	"github.com/mmatur/dockerhub/cmd"
	"github.com/mmatur/dockerhub/list"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// rootCmd represents the base command when called without any subcommands.
	rootCmd := &cobra.Command{
		Use:   "dockerhub",
		Short: "dockerhub",
		Long:  "dockerhub",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return rootRun()
		},
	}

	rootCmd.AddCommand(cmd.NewVersion())
	rootCmd.AddCommand(cmd.NewDoc())
	rootCmd.AddCommand(cmd.NewListAll())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootRun() error {
	action, err := choose.Action()
	if err != nil {
		return err
	}

	if action == choose.ActionListAll {
		return list.All()
	}

	return nil
}
