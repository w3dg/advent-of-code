package main

const (
	NODE_TYPE_DIR  string = "dir"  // To indicate the node is a directory
	NODE_TYPE_FILE string = "file" // To indicate the node is a directory
)

type Node struct {
	// AbsolutePath string

	Parent   *Node // to make `cd ..` possible
	Children *map[string]*Node
	Name     string
	Type     string
	Size     int // dirs have size -1 until recursively calculated
}

var Root *Node = nil

// name will be either a child dir name or `..`
// In the problem, cd to / is only ever done once. (assumption)
func (n *Node) MoveIntoDir(name string) *Node {
	if name == ".." {
		// root's parent is itself (similar to linux file systems)
		if n.Parent == nil {
			return n
		}
		return n.Parent
	}

	ch := *n.Children
	v, ok := ch[name]
	if !ok {
		panic("No such directory " + name)
	}
	return v
}

func MakeNewDir(parent *Node, name string) *Node {
	m := make(map[string]*Node)
	return &Node{
		Parent:   parent,
		Children: &m,
		Name:     name,
		Type:     NODE_TYPE_DIR,
		Size:     -1,
	}
}

func MakeNewFile(parent *Node, name string, size int) *Node {
	return &Node{
		Parent:   parent,
		Children: nil,
		Name:     name,
		Type:     NODE_TYPE_FILE,
		Size:     size,
	}
}

func MakeRootDir() *Node {
	Root = MakeNewDir(nil, "/")
	return Root
}

func (n *Node) MakeChildDir(name string) {
	if n.Type != NODE_TYPE_DIR {
		panic("Attempt to make a child dir under a node which is not a dir")
	}

	ch := *n.Children

	if _, ok := ch[name]; ok {
		panic("dir " + name + " already exists under dir " + n.Name)
	}

	d := MakeNewDir(n, name)
	ch[name] = d
}

func (n *Node) MakeChildFile(name string, size int) {
	if n.Type != NODE_TYPE_DIR {
		panic("Attempt to make a child file under a node which is not a dir")
	}
	ch := *n.Children

	if _, ok := ch[name]; ok {
		panic("file " + name + " already exists under dir " + n.Name)
	}

	f := MakeNewFile(n, name, size)
	ch[name] = f
}
