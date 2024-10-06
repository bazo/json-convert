package cmd

import (
	"bufio"
	"encoding/json"
	"json-convert/parquet"
	"json-convert/utils"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	"github.com/xitongsys/parquet-go-source/local"
)

func NewConvert3Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "convert3 [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			infile := filename + ".jsonl"

			outfile := filename + ".parquet"

			f, err := os.Open(infile)
			if err != nil {
				log.Fatalln("Cannot open file", infile)
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)

			stats, err := f.Stat()

			if err != nil {
				log.Fatalln("Cannot stat file", infile)
			}

			// Create a new Parquet writer
			fw, err := local.NewLocalFileWriter(outfile)
			if err != nil {
				log.Println("Can't create local file", err)
				return
			}
			defer fw.Close()

			pw, err := parquet.NewParquetWriter(fw, parquet.MaxPageSize(10000), parquet.Snappy)
			if err != nil {
				log.Fatal(err)
			}

			i := 0
			bar := progressbar.DefaultBytes(stats.Size())
			for scanner.Scan() {
				b := scanner.Bytes()
				line := &parquet.Line{}
				json.Unmarshal(b, line)

				pw.Add(*line.ToParquet())

				bar.Add(len(b))
				bar.Describe(utils.GetMemUsage())
				//b = nil
				//line = nil
				i++

				if i%100 == 0 {
					err = pw.Write()
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
			bar.Finish()
			err = pw.Close()
			if err != nil {
				log.Fatal("Error stopping writer:", err)
			}

			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

		},
	}

	return cmd
}
