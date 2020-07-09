package list

import (
	"context"
	"os"
	"sort"
	"strconv"

	"github.com/fatih/color"
	"github.com/mmatur/dockerhub/docker"
	"github.com/olekukonko/tablewriter"
)

func All() error {
	ctx := context.Background()
	client, err := docker.NewClient()
	if err != nil {
		return err
	}

	results, err := client.ListAllRepo(ctx)
	if err != nil {
		return err
	}

	sort.Sort(results)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Position", "Name", "Pull", "Stars"})

	for i, result := range results {
		table.Append([]string{strconv.Itoa(i + 1), color.BlueString(result.Source + "/" + result.Slug), color.YellowString(strconv.FormatInt(result.Popularity, 10)), color.YellowString(strconv.Itoa(result.StarCount))})
	}
	table.Render()

	return nil
}
