package main

import "fmt"

type Mahasiswa struct {
	Nama_Mahasiswa string
	Usia_Mahasiswa int
}

type Node struct {
	Data *Mahasiswa
	Next *Node
}

func (n *Node) addNode(data *Mahasiswa) {
	for n.Next != nil {
		n = n.Next
	}

	newNode := &Node{
		Data: data,
		Next: nil,
	}

	n.Next = newNode
}

func (n *Node) deleteNode() {
	if n.Next == nil {
		fmt.Println("True")
		n = &Node{}
		return
	}

	for n.Next.Next != nil {
		n = n.Next
	}

	n.Next = nil
}

func (n *Node) showAll() {
	for n != nil {
		fmt.Println("Data Mahasiswa :", n)
		n = n.Next
	}
}

func main() {
	joshua := &Mahasiswa{
		Nama_Mahasiswa: "Joshua Ryandafres Pangaribuan",
		Usia_Mahasiswa: 21,
	}

	chesya := &Mahasiswa{
		Nama_Mahasiswa: "Chesya Sitorus",
		Usia_Mahasiswa: 17,
	}

	irwan := &Mahasiswa{
		Nama_Mahasiswa: "Irwan Siagian",
		Usia_Mahasiswa: 20,
	}

	mainNode := &Node{
		Data: joshua,
		Next: nil,
	}

	mainNode.addNode(chesya)
	mainNode.addNode(irwan)

	mainNode.deleteNode()
	mainNode.deleteNode()
	mainNode.deleteNode()

	mainNode.showAll()
}
