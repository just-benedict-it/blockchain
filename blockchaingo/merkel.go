package blockchain

import "crypto/sha256"

type MerkelTree struct {
	RootNode *MerkelNode
}

type MerkelNode struct {
	Left  *MerkelNode
	Right *MerkelNode
	Data  []byte
}

func NewMerkelNode(left, right *MerkelNode, data []byte) *MerkelNode {
	node := MerkelNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkelTree(data [][]byte) *MerkelTree {
	var nodes []MerkelNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, dat := range data {
		node := NewMerkelNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var level []MerkelNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkelNode(&nodes[j], &nodes[j+1], nil)
			level = append(level, *node)
		}
		nodes = level
	}

	tree := MerkelTree{&nodes[0]}

	return &tree

}
