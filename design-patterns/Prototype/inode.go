package main

type INode interface {
	print(string)
	clone() INode
}

type File struct {
	name string
}

type Folder struct {
	children []INode
	name     string
}
