package cmd

import (
	"bufio"
	"encoding/json"
	"json-convert/types"
	"json-convert/utils"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

func NewConvert2Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "convert2 [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			infile := filename + ".jsonl"

			outfile := filename + ".parquet"

			f, err := os.Open(utils.DockerPath(infile))
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
			fw, err := local.NewLocalFileWriter(utils.DockerPath(outfile))
			if err != nil {
				log.Println("Can't create local file", err)
				return
			}
			defer fw.Close()

			pw, err := writer.NewParquetWriter(fw, new(types.ParquetLine), 4)
			//pw, err := writer.NewJSONWriter(md, fw, 4)
			if err != nil {
				log.Fatal("Can't create parquet writer", err)
			}
			defer pw.WriteStop() // Flushes the writer and closes

			pw.RowGroupSize = 128 * 1024 * 1024 //128M
			pw.PageSize = 8 * 1024              //8K
			pw.CompressionType = parquet.CompressionCodec_SNAPPY

			bar := progressbar.DefaultBytes(stats.Size())
			for scanner.Scan() {
				b := scanner.Bytes()
				line := &types.Line{}
				json.Unmarshal(b, line)

				err = pw.Write(line.ToParquet())
				if err != nil {
					log.Fatalln(err)
				}

				bar.Add(len(b))
				bar.Describe(utils.GetMemUsage())
				//b = nil
				//line = nil
			}

			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

			err = pw.WriteStop()
			if err != nil {
				log.Fatal("Error stopping writer:", err)
			}

			bar.Finish()
		},
	}

	return cmd
}
