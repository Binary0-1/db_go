package src

import (
	"encoding/binary"
)

func assert(cond bool) {
	if !cond {
		panic("assertion Failed")
	}
}

const (
	BNODE_NODE = 1 // internal nodes without values
	BNODE_LEAF = 2 // leaf nodes with values
)

type Bnode struct {
	data []byte
}

type Btree struct {
	root uint64

	get func(uint64) Bnode // dereference a pointer
	new func(Bnode) uint64 // allocate a new page
	del func(uint64)       // deallocate a page
}

const HEADER = 4
const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VALUE_SIZE = 3000

func init() {
	node1Max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VALUE_SIZE
	assert(node1Max <= BTREE_PAGE_SIZE)
}


//Header
func (node Bnode) bType() uint16 {
	return binary.LittleEndian.Uint16(node.data[:2])
}

func (node Bnode) nKeys(){
	return binary.LittleEndian.Uint16(node.data[2:4])
}

func (node Bnode) setHeader(btype uint16, nkeys uint16){
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}


//Pointer 
func (node Bnode) getPtr(idx unint16) uint64{
	assert(idx < node.nkeys())
	pos := HEADER + 8 * idx
	return binary.LittleEndian.uint64(node.data[pos:])
}

func (node Bnode) setPtr(inx uint16, val unit64){
	assert(idx < node.nkeys()){
		pos := HEADER + 8 * idx
		binary.LittleEndian.PutUint16(node.data[pos:], val)
	}
}

