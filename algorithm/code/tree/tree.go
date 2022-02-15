/*
 * @Author : wb
 * @Date : 2022/2/12 10:57 上午
 */

package tree

// 遍历二叉树，层序遍历

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(n *Node, v int) *Node {
	if n == nil {
		return &Node{
			Val: v,
		}
	}

	head := n
	for n != nil {
		if v > n.Val {
			if n.Right == nil {
				n.Right = &Node{Val: v}
				return head
			}

			n = n.Right
		} else {
			if n.Left == nil {
				n.Left = &Node{Val: v}
				return head
			}

			n = n.Left
		}
	}

	return head
}
func find(n *Node, v int) *Node {
	if n == nil {
		return nil
	}

	for n != nil {
		if n.Val == v {
			return n
		} else if n.Val > v {
			n = n.Left
		} else {
			n = n.Right
		}
	}

	return nil
}
