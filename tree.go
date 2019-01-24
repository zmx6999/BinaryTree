package main

import (
	"fmt"
	"reflect"
)

type BTreeNode struct {
	Data interface{}
	LeftChild *BTreeNode
	RightChild *BTreeNode
}

func CreateBTree(data interface{}) *BTreeNode {
	node:=new(BTreeNode)
	node.Data=data
	return node
}

func (t *BTreeNode) PreTravel()  {
	if t==nil {
		return
	}
	fmt.Print(t.Data," ")
	t.LeftChild.PreTravel()
	t.RightChild.PreTravel()
}

func (t *BTreeNode) MidTravel()  {
	if t==nil {
		return
	}
	t.LeftChild.MidTravel()
	fmt.Print(t.Data," ")
	t.RightChild.MidTravel()
}

func (t *BTreeNode) RearTravel()  {
	if t==nil {
		return
	}
	t.LeftChild.RearTravel()
	t.RightChild.RearTravel()
	fmt.Print(t.Data," ")
}

func (t *BTreeNode) Height() int {
	if t==nil {
		return 0
	}
	lh:=t.LeftChild.Height()
	rh:=t.RightChild.Height()
	if lh>rh {
		return lh+1
	} else {
		return rh+1
	}
}

func (t *BTreeNode) LeafCount(n *int)  {
	if t==nil {
		return
	}
	if t.LeftChild==nil && t.RightChild==nil {
		*n++
	}
	t.LeftChild.LeafCount(n)
	t.RightChild.LeafCount(n)
}

func (t *BTreeNode) Search(data interface{})  {
	if t==nil {
		return
	}
	if reflect.TypeOf(data)==reflect.TypeOf(t.Data) && data==t.Data {
		fmt.Print(t.Data)
		return
	}
	t.LeftChild.Search(data)
	t.RightChild.Search(data)
}

func (t *BTreeNode) Reverse()  {
	if t==nil {
		return
	}
	t.LeftChild.Reverse()
	t.RightChild.Reverse()
	t.LeftChild,t.RightChild=t.RightChild,t.LeftChild
}

func (t *BTreeNode) Copy() *BTreeNode {
	if t==nil {
		return nil
	}
	leftChild:=t.LeftChild.Copy()
	rightChild:=t.RightChild.Copy()
	copiedNode:=new(BTreeNode)
	copiedNode.Data=t.Data
	copiedNode.LeftChild=leftChild
	copiedNode.RightChild=rightChild
	return copiedNode
}

func (t *BTreeNode) Destroy()  {
	if t==nil {
		return
	}
	t.LeftChild.Destroy()
	t.RightChild.Destroy()
	t.Data=nil
	t.LeftChild=nil
	t.RightChild=nil
	t=nil
}

func main()  {
	tree:=CreateBTree(0)

	node1:=CreateBTree(1)
	tree.LeftChild=node1
	node2:=CreateBTree(2)
	tree.RightChild=node2

	node3:=CreateBTree(3)
	node1.LeftChild=node3
	node4:=CreateBTree(4)
	node1.RightChild=node4

	node5:=CreateBTree(5)
	node2.LeftChild=node5
	node6:=CreateBTree(6)
	node2.RightChild=node6

	node7:=CreateBTree(7)
	node3.LeftChild=node7

	tree.PreTravel()
	fmt.Println()
	tree.MidTravel()
	fmt.Println()
	tree.RearTravel()
	fmt.Println()
	fmt.Println(tree.Height())
	n:=0
	tree.LeafCount(&n)
	fmt.Println(n)
	for i:=0; i<=8; i++ {
		tree.Search(i)
		fmt.Println()
	}

	tree.Reverse()
	tree.PreTravel()
	fmt.Println()
	tree.MidTravel()
	fmt.Println()
	tree.RearTravel()
	fmt.Println()
	fmt.Printf("%p\n",tree)
	copiedTree:=tree.Copy()
	fmt.Printf("%p\n",copiedTree)
	tree.PreTravel()
	fmt.Println()
	tree.MidTravel()
	fmt.Println()
	tree.RearTravel()
	tree.Destroy()
	fmt.Print(tree)
}
