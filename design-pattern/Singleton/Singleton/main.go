package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(city string) int {
	return db.capitals[city]
}

// Thread safe
// Using sync.Once and init()
// Laziness (only construct when needed) --> sync.Once

var once sync.Once
var instance *singletonDatabase

func readDataFromFile(path string) (map[string]int, error) {
	// Not needed in Mac
	//executable, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//
	//symLinks, _ := filepath.EvalSymlinks(executable)
	//directory := filepath.Dir(symLinks)
	//
	//fmt.Println(directory)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		capitals := scanner.Text()
		scanner.Scan()
		pops, _ := strconv.Atoi(scanner.Text())
		result[capitals] = pops
	}

	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readDataFromFile("capitals.txt")
		if err != nil {
			panic(err)
		}
		db := singletonDatabase{caps}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()
	fmt.Println(db.GetPopulation("Seoul"))
}
