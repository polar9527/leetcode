package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if root.Val < p.Val && root.Val < q.Val {
	// 	return lowestCommonAncestor(root.Right, p, q)
	// }
	// if root.Val > p.Val && root.Val > q.Val {
	// 	return lowestCommonAncestor(root.Left, p, q)
	// }
	//
	// return root

	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}
