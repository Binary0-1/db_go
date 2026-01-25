package src;

import (
	"fmt"
)

func assert(cond bool){
	if(!cond){
		panic("assertion Failed")
	}
}

const (
	BNODE_NODE = 1 // internal nodes without values
	BNODE_LEAF = 2 // leaf nodes with values
 )

type Bnode struct{
	data []byte
}

type Btree struct {
	root uinit64

	get func(uinit64) Bnode // dereference a pointer
	new func(Bnode) uinit64 // allocate a new page
	del func(uinit64) // deallocate a page
}

const HEADER = 4
const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VALUE_SIZE = 3000


func init(){
	node1Max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VALUE_SIZE
	assert(node1Max <= BTREE_PAGE_SIZE)
}
