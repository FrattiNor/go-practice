package ginSample

import (
	"path"
	"strings"
)

type API struct {
	pattern string
	method  string
}

type Router struct {
	root       *Node
	handlerMap map[string][]*HandleFunc
	apis       []API
}

func newRouter() *Router {
	return &Router{
		root:       &Node{part: "root"},
		apis:       make([]API, 0),
		handlerMap: make(map[string][]*HandleFunc, 0),
	}
}

func splitPattern(method string, pattern string) []string {
	parts := make([]string, 1)
	parts[0] = method

	for _, item := range strings.Split(pattern, "/") {
		if item != "" {
			parts = append(parts, item)
		}
	}

	return parts
}

func recordDynamicIndex(node *Node, nodeType int) {
	if nodeType > 0 {
		if node.dynamicIndex > 0 {
			panic("同级只能存在一个动态路由")
		} else {
			node.dynamicIndex = 1
		}
	}
}

func (r *Router) addRoute(method string, pattern string, handler HandleFunc, middleware []*HandleFunc) {
	parts := splitPattern(method, pattern)
	currentNode := r.root

	for _, part := range parts {
		node := currentNode.findChild(part, true)

		if node == nil {
			nodeType := getNodeType(part)
			recordDynamicIndex(currentNode, nodeType)
			nextNode := &Node{part: part, nodeType: nodeType}
			currentNode.children = append(currentNode.children, nextNode)
			currentNode = nextNode
		} else {
			currentNode = node
		}
	}

	r.apis = append(r.apis, API{path.Join(parts[1:]...), method})

	currentNode.pattern = path.Join(parts[:]...)

	if _, ok := r.handlerMap[currentNode.pattern]; ok {
		panic("已存在接口")
	} else {
		handlers := make([]*HandleFunc, 0)
		handlers = append(handlers, middleware...)
		handlers = append(handlers, &handler)
		r.handlerMap[currentNode.pattern] = handlers
	}
}

func (r *Router) searchRoute(method string, pattern string) (*Node, map[string]string) {
	parts := splitPattern(method, pattern)
	currentNode := r.root
	params := make(map[string]string)

	for _, part := range parts {

		node := currentNode.findChild(part, false)
		if node == nil {
			return nil, nil
		} else {
			if node.nodeType > 0 {
				params[node.part[1:]] = part
			}
			currentNode = node
		}
	}

	return currentNode, params
}

func (r *Router) findHandlers(method string, pattern string) ([]*HandleFunc, map[string]string) {
	node, params := r.searchRoute(method, pattern)

	if node == nil || node.pattern == "" {
		return nil, nil
	}

	return r.handlerMap[node.pattern], params
}

func (r *Router) logApis() {
	for _, value := range r.apis {
		colorPrint(value.method, value.pattern)
	}
}
