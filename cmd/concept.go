package cmd

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/concept"
	"github.com/phodal/coca/core/models"
	. "github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var parsedDeps []models.JClassNode

var conceptCmd = &cobra.Command{
	Use:   "concept",
	Short: "concept api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := cmd.Flag("dependence").Value.String()

		if dependence != "" {
			analyser := concept.NewConceptAnalyser()
			file := ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}

			_ = json.Unmarshal(file, &parsedDeps)

			wordCounts := analyser.Analysis(&parsedDeps)

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Words", "Counts"})

			for _, word := range wordCounts {
				if word.Value > 0 {
					table.Append([]string{word.Key, strconv.Itoa(word.Value)})
				}
			}

			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(conceptCmd)

	conceptCmd.PersistentFlags().StringP("dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
