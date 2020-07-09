package cmd

import (
	"github.com/mmatur/dockerhub/list"
	"github.com/spf13/cobra"
)

func NewListAll() *cobra.Command {
	listAll := &cobra.Command{
		Use:   "list-all",
		Short: "List all",
		Long:  "List all",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return list.All()
		},
	}

	return listAll
}
