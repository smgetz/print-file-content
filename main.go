package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type PatientRecord struct {
	PatientId, PatientRecord string
	PatientNumber            int
}

func main() {
	//GLOB returns a slice of data including what files are in this directory, wildcard .txt
	allTxtFiles, err := filepath.Glob(`\\qumulo\BLIS\ScottGetzFiles\*.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}
	patients := []PatientRecord{}
	// v=variable for each file name
	//range-maps through directory
	// _ ...
	for _, v := range allTxtFiles {
		fmt.Printf("Full file path is %s\n", v)
		//fmt.Printf("Filename is %s\n", filepath.Base(v)
		patient, err := grabPatientInfo(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		patients = append(patients, patient)
	}

	fmt.Printf("%+v\n", patients)
}

//Helper function that ensures the defer(40) method to run, that will close each file.
//Opens each file and will scan txt to print line by line

func grabPatientInfo(filePath string) (PatientRecord, error) {
	pr := PatientRecord{}

	file, err := os.Open(filePath)
	if err != nil {
		return pr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var textLine string

	for scanner.Scan() {
		// if multi lined txt file, concat/append
		// textLine = fmt.Sprintf("%s%s", textLine, scanner.Text())
		textLine = scanner.Text()
	}

	contentSplit := strings.Split(textLine, "|")

	if len(contentSplit) != 3 {
		return pr, fmt.Errorf("File %s Does not have 3", filePath)
	}

	indexAsInt, err := strconv.Atoi(contentSplit[2])
	if err != nil {
		return pr, err
	}

	pr.PatientId = contentSplit[0]
	pr.PatientRecord = contentSplit[1]
	pr.PatientNumber = indexAsInt

	return pr, nil
}

// func lastValue(content string) (string, error) {
// 	contentSplit := strings.Split(content, "|")

// 	if len(contentSplit) != 3 {
// 		return "", fmt.Errorf("Does not have 3")
// 	}

// 	return contentSplit[2], nil
// 	// lastNum := strings.SplitN(content, "|", 1)
// 	// for v := range lastNum {
// 	// 	fmt.Println(lastNum[v])
// 	// }
// }
