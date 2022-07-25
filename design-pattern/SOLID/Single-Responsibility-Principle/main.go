package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	Single Responsibility Principle (SRP)
	- Make sure in your code, every object/class ONLY have one responsibility
	- Don't create a GOD object/class
	- Separate object/class by their responsibility
*/

// Here we add Journal struct
// This Journal struct have many responsibility (AddEntries, GetEntriesCount)
type Journal struct {
	entries    []string
	entryCount int
}

func (j *Journal) AddEntries(text string) int {
	j.entryCount++
	entry := fmt.Sprintf("%d:%s", j.entryCount, text)
	j.entries = append(j.entries, entry)
	return j.entryCount
}

func (j *Journal) GetEntriesCount() int {
	return j.entryCount
}

type Utils struct {
	lineSeparator string
}

func (u *Utils) Stringer(obj interface{}) string {
	switch obj.(type) {
	case Journal:
		v, ok := obj.(Journal)
		if ok {
			return strings.Join(v.entries, u.lineSeparator)
		}
	}
	return ""
}

/*
	But in this section we add some persistence responsibility to Journal struct
	We add (SaveToFile, and LoadFromFile)
	So, the example below in SRP is not allowed.
	Because, the journal struct have more than one responsibility
	And the solution to this is to make a new struct that has persistence responsibility
*/

//func (j *Journal) SaveToFile(filename string) {
//	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
//}

//func (j *Journal) LoadFromFile(filename string) {
//
//}

// We add new struct here to separate the responsibility
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(journal *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(journal.entries, p.lineSeparator)), 0644)
}

func main() {
	// The responsibility of this journal variable is
	// - Adding Entries
	// - Get Total Entries
	journal := Journal{}
	journal.AddEntries("Joshua Ryandafres")
	journal.AddEntries("Pangaribuan")

	// The responsibility of this util variable is
	// - Modify data in interface{} into string
	util := Utils{"\n"}
	fmt.Println(util.Stringer(journal))

	// The responsibility of this persist variable is
	// - Save the journal entries to specific file
	persist := Persistence{"\n"}
	persist.SaveToFile(&journal, "firstJournal.txt")
}
