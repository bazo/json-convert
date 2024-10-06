//go:generate parquetgen -input parquet/types.go -type ParquetLine -package parquet -output parquet/parquet.go
package main

import (
	"json-convert/cmd"
)

func main() {
	cli := cmd.NewCmd()
	cli.Execute()

}
