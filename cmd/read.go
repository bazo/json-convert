package cmd

import (
	"json-convert/types2"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/parquet-go/parquet-go"
	"github.com/spf13/cobra"
)

func NewReadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "read [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			file := filename + ".parquet"

			rows, err := parquet.ReadFile[types2.ParquetLine](file)
			if err != nil {
				log.Fatalln(err)
			}

			for _, c := range rows {
				spew.Dump(c)
			}
		},
	}

	return cmd
}
