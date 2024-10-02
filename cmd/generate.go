package cmd

import (
	"encoding/json"
	"json-convert/generator"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var newLine = []byte("\n")

func NewGenerateCommand() *cobra.Command {
	var lines int
	cmd := &cobra.Command{
		Use:  "generate [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			f, err := os.Create(filename + ".jsonl")
			if err != nil {
				log.Fatalln("Cannot create output file")
			}
			defer f.Close()

			bar := progressbar.Default(int64(lines))
			for i := 0; i < lines; i++ {
				line := generator.Generate()

				b, err := json.Marshal(line)

				if i < lines-1 {
					b = append(b, newLine...)
				}
				if err != nil {
					log.Fatalln("Couldn't marshal line")
				}

				_, err = f.Write(b)

				if err != nil {
					log.Fatalln("Couldn't write line")
				}
				bar.Add(1)
			}
		},
	}

	cmd.Flags().IntVarP(&lines, "lines", "l", 100, "How many lines to generate")
	return cmd
}
