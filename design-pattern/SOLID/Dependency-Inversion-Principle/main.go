package main

import "fmt"

// DIP
// High Level Module should not depend on Low Level Module
// Both of them should depend on abstraction

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// Another attribute here
}

// Model relationship between people
type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Store the information about relationship between people

// Low Level Module
// Why LLM? -> Because we need this program just storage
// We can change this with API or DB
// In this case we store the relations Info on memory
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(
	parent, child *Person) {

	// Add relations
	r.relations = append(r.relations,
		Info{parent, Parent, child})
	r.relations = append(r.relations,
		Info{child, Child, parent})
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

// Solution DIP
// Make a Abstraction!
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// High Level Module
// Why HLM? -> Because we need this in program to operate and perform
// some research or operation in data from LLM (storage)
type Research struct {

	// Break DIP Case
	// Because this struct depends on LLM
	//relationships Relationships

	// Solution
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	/*
		If we still let the High Level Module depends on
		Low Level Module may be this function no longer work correctly
	*/
	//relations := r.relationships.relations
	//for _, rel := range relations {
	//	if rel.from.name == "Joshua" &&
	//		rel.relationship == Parent {
	//		fmt.Println("John has a child called", rel.to.name)
	//	}
	//}

	// Handle the new change of LLM
	for _, p := range r.browser.FindAllChildrenOf("Joshua") {
		fmt.Println("Joshua has a child called", p.name)
	}
}

func main() {
	parent := &Person{"Joshua"}
	child1 := &Person{"Irwan"}
	child2 := &Person{"Wilson"}

	// Add relationships
	relationships := Relationships{}
	relationships.AddParentAndChild(parent, child1)
	relationships.AddParentAndChild(parent, child2)

	r := Research{&relationships}
	r.Investigate()
}
