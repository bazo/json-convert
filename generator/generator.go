package generator

import (
	"fmt"
	"json-convert/types"

	"github.com/go-faker/faker/v4"
)

func Generate() *types.Line {
	line := &types.Line{}
	err := faker.FakeData(line)

	if err != nil {
		fmt.Println(err)
	}

	return line
}
