package ginSample

type Node struct {
	part         string  // 当前part字符串
	children     []*Node // child
	nodeType     int     // 0 正常，1 :，2 *
	dynamicIndex int     // child动态个数
	pattern      string  // 完整part
}

func getNodeType(part string) int {
	nodeType := 0

	if part[0:1] == ":" {
		nodeType = 1
	}

	if part[0:1] == "*" {
		nodeType = 2
	}

	return nodeType
}

func (n *Node) findChild(part string, exact bool) *Node {
	if n.children == nil {
		return nil
	}

	nodeType := getNodeType(part)

	nodes := make([]*Node, 0)

	for _, node := range n.children {
		if node.part == part {
			return node
		}

		if !exact && node.nodeType > 0 {
			nodes = append(nodes, node)
			continue
		}

		if exact && node.nodeType > 0 && node.nodeType == nodeType {
			nodes = append(nodes, node)
			continue
		}
	}

	if len(nodes) == 0 {
		return nil
	} else {
		return nodes[0]
	}
}
