package main

import "fmt"

func (f *Folder) print(indentation string) {
	fmt.Println("|" + indentation + f.name)
	for _, i := range f.children {
		i.print("-" + indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	var children []INode

	cloneFolder := &Folder{name: f.name}

	children = append(children, f.children...)

	cloneFolder.children = children
	return cloneFolder
}
