package cmd

import (
	"json-convert/types"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func NewRead2Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "read2 [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			file := filename + ".parquet"

			fr, err := local.NewLocalFileReader(file)

			if err != nil {
				log.Println("Can't open file", err)
				return
			}
			defer fr.Close()

			pr, err := reader.NewParquetReader(fr, new(types.ParquetLine), 4)
			if err != nil {
				log.Println("Can't create parquet reader", err)
				return
			}
			num := int(pr.GetNumRows())

			//bar := progressbar.Default(int64(num))
			for i := 0; i < num; i++ {
				lines := make([]types.ParquetLine, 1)
				if err = pr.Read(&lines); err != nil {
					log.Fatalln("Read error", err)
				}
				spew.Dump(lines[0])
			}
			pr.ReadStop()

			//bar.Finish()
		},
	}

	return cmd
}
