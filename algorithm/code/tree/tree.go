package tree

import "fmt"

// 二叉树的表示方式
// 一、基于指针的链表
// 二、基于数组的顺序:根节点为i， 左子树为2*i, 右子树为2*i+1
// 完全二叉树浪费了数组的首个位置，非完全二叉树浪费更多。所以完全二叉树最节省内存
// 完全二叉树定义：叶子节点都在最后两层，且除最后一层，其他层的节点都达到最大。最后一层的都靠左

// 二叉树遍历,每个节点最多会访问两次，时间复杂度O(2n) = O(n)
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

func printTree(t *Node) {
	if t == nil {
		return
	}

	fmt.Println(t.Val)
	printTree(t.Left)
	printTree(t.Right)
}

// 层序遍历
func inOrder(t *Node) {
	var q []*Node

	q = append(q, t)
	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		fmt.Println(n.Val)
		if n.Left != nil {
			q = append(q, n.Left)
		}

		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}

// 二叉查找树(二叉搜索树)、支持快快速的增删查
// 特点：左 < 根 < 右
// 散列的增删查时间复杂度O(1)
// 其他操作：最大（右子树的最右）、最小（左子树的最左）、前驱节点[小于该节点的最大值]、后继节点[大于该节点的最小值]

// 利用对象的某个字段构建二叉树。其他字段称之为卫星字段

// 重复二叉查找树：
// 1、节点存储链表或者支持动态扩容的数组，把相同的数据存在一个节点上
// 2、 每个节点仍然存储一个数据，对于相同的则统一存到右子树。查找时相对比较麻烦，查到后需要继续查到底

// 二叉查找树时间复杂度
// 极度不平衡的退化成链表O(n)
// 理想的完全二叉查找树O(h) = O(L - 1)
// L (log (n+1), logn + 1)
// 所以完全二叉查找树的时间复杂度 <= logn
// 需要稳定的 -> 平衡二叉查找树 高度 log(n) 所以查找插入的时间为O(logn)
// 散列 vs 二叉查找树
// 1. 散列无序，二叉查找树中序遍历之后就有序了
// 2. 散列扩容时比较耗时，平衡二叉查找树可以做到O(logn)
// 3. 散列存在hash冲突，如果hash冲突时间可能更多
// 4. 散列的构造比较麻烦，需要考虑hash函数、hash冲突，扩容、缩容。平衡二叉查找树只需要考虑平衡
// 5. 为了避免过多的散列冲突，散列表装载因子不能太大，特别是基于开放寻址法的散列，不然会浪费内存
//
//
// 平衡二叉查找树:二叉树中任意一个节点的左右子树的高度差不能大于1
// 完全二叉树或者满二叉树都是平衡二叉树，非完全二叉树也有可能是平衡二叉树
//
// AVL树是一种平衡二叉树，查找效率高。但是为了维护平衡付出更多代价，插入，删除都要做调整比较大
// 红黑树只做到近似平衡，并非严格平衡，维护成本低于AVL
//
// 平衡二叉树的初衷是解决普通二叉树在频繁插入、删除等动态更新的情况下，出现时间复杂度退化的情况
// 红黑树是平衡二叉树的一种，但不是严格的平衡二叉树
// 满足一下要求：
// 1. 根节点是黑色
// 2. 每个叶子节点是黑色的空节点，即叶子节点不存储数据
// 3. 任何相邻的节点都不能同时为红色，也就是说红色节点是被黑色节点隔开的
// 4. 每个节点，从该节点到达其可以到达的叶子节点的所有路径，都包含相同数目的黑色节点
//
//
// 平衡 = 性能不退化
// 近似平衡 = 性能不严重退化

// 求二叉查找树树的高度（层）
func findL(t *Node) int {
	if t.Left == nil && t.Right == nil {
		return 1
	}

	hl := findL(t.Left)
	hr := findL(t.Right)

	if hl > hr {
		return hl + 1
	} else {
		return hr + 1
	}
}

// 查找：类似二分
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

	p := n
	for p != nil {
		if p.Val == v {
			return n
		} else if p.Val > v {
			p = p.Left
		} else {
			p = p.Right
		}
	}

	return nil
}

// 如果v小于根节点，若跟节点的左子树为空则，改节点插入到其左子树上，若果左子树不为空，则重新遍历左子树
func insert(n *Node, v int) *Node {
	if n == nil {
		return &Node{Val: v}
	}

	p := n
	for p != nil {
		if p.Val > v {
			if p.Left == nil {
				p.Left = &Node{}
				break
			} else {
				p = p.Left
			}
		} else {
			if p.Right == nil {
				p.Right = &Node{}
				break
			} else {
				p = p.Right
			}
		}
	}

	return n
}

// 删除
// case1：没有子节点
// case2: 只有一个子节点
// case3: 有两个子节点
func delete(n *Node, v int) *Node {

	p := n
	for p != nil {
		if p.Val != v {
			if p.Val > v {
				p = p.Left
			} else {
				p = p.Right
			}

			continue
		}

		// case1: 没有子节点直接删除
		if p.Left == nil && p.Right == nil {
			p = nil
			break
			// case2: 有左节点或者右节点 把左右节点提上去，左右节点置为空
		} else if p.Left == nil && p.Right != nil {
			p = p.Right
			p.Right = nil
			break
		} else if p.Left != nil && p.Right == nil {
			p = p.Left
			p.Left = nil
			break
		} else { // 从右子树找到最小的
			// case3:左右节点都有则找到右子树的最左侧节点赋值给当前节点，其他的
			r := p.Right
			for r.Left != nil {
				r = r.Left
			}

			p.Right = r
			r = nil
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
