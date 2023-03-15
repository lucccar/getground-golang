package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func GenerateID() (int64, error) {

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Generate a snowflake ID.
	id := node.Generate().Int64()
	return id, nil
}
