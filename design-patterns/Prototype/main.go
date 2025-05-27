package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1", children: []INode{file1, file2}}
	folder2 := &Folder{name: "Folder2", children: []INode{file1, file2, file3}}

	clone1 := folder1.clone()
	clone1.print(" ")

	clone2 := folder2.clone()
	clone2.print(" ")
}
