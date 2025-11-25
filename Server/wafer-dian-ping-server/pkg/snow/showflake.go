package snow

import (
	"github.com/bwmarrin/snowflake"
)

func NewId(nodeVal int64) int64 {

	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(nodeVal)

	// Generate a snowflake ID.
	id := node.Generate()

	return id.Int64()
}
