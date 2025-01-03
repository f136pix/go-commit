package commit

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	files []string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a commit message based on the current Git repository context",
	Long:  `This command is going to analyse all the tracked files and generate a commit message for the changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		files = append(files, args...)
		for _, file := range files {
			fmt.Println(file)
		}

	},
}

func init() {
	// -f flag
	generateCmd.Flags().StringSliceVarP(&files, "file", "f", []string{}, "Files to analyze")

	// -o flag

	CommitCmd.AddCommand(generateCmd)
}
