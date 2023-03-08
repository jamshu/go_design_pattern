package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


var entryCount = 0
type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries,entry)
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries,"\n")
}

func SavetoFile(filename string,j *Journal) error {
	err := ioutil.WriteFile(filename,[]byte(j.String()),0644)
	return err
}

func main() {
	journal := Journal{}
	journal.AddEntry("i am crired today")
	journal.AddEntry("i studied")
	journal.AddEntry("I read Quran")
	data := journal.String()
	fmt.Println(data)

	err := SavetoFile("journal.txt",&journal)
	if err != nil {
		fmt.Println(err)
	}


}