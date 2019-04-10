package emojipedia

type node interface{}

type trie interface{}

type Node struct {
	End       bool
	Namespace *Namespace
}

type Trie struct {
	Root *Node
}
