package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const CourseLearnGoLang = "CourseLearnGoLang"
const SummaryMD = "summary.md"
const SummaryText = "summary.txt"

type topic struct {
	chapter string
	title   string
	summary []topic
}

var summary = make([]topic, 0)

func main() {
	root, _ := os.Getwd()
	log.Println("Init", root)

	PopulateSlice()
	CreateFile(SummaryMD)
	WriteSummaryMD()

	CreateDirectory(CourseLearnGoLang)
	os.Chdir(CourseLearnGoLang)

	CreateFoldersBasedOnSummary()
	fmt.Println()
}

func PopulateSlice() {
	log.Println("Populate summary populated")
	log.Println("")

	file, err := os.Open(SummaryText)

	if err != nil {
		log.Println("Erro on populate", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sTitle := strings.Split(scanner.Text(), "–")

		chapter := sTitle[0]
		title := sTitle[1]
		subtopic := sTitle[2]

		canAppend := true
		if len(summary) >= 1 {
			last := len(summary) - 1
			currentTopic := &summary[last]

			if (*currentTopic).chapter == chapter {
				canAppend = false

				currentTopic.summary = append(currentTopic.summary, topic{chapter, subtopic, make([]topic, 0)})
			}
		}

		if canAppend {

			t := topic{chapter, title, make([]topic, 0)}
			t.summary = append(t.summary, topic{chapter, subtopic, make([]topic, 0)})

			summary = append(summary, t)
		}
	}

	log.Println("Successfully summary populated")

}

func CreateDirectory(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if os.IsExist(err) {
			log.Printf("Error: %v on created directory: %v\n", err, path)
		}
		log.Printf("Created directory :%v\n", path)
	}
}

func WriteSummaryMD() {
	file, err := os.OpenFile(SummaryMD, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	if _, err := file.WriteString("### Summary Course Learn Go Lang\n"); err != nil {
		log.Fatal(err)
	}
	s := ""

	for _, v := range summary {

		sn := fmt.Sprintf("\n### %v\n"+
			"\r* %v  ",
			v.chapter, v.title)

		for _, v := range v.summary {
			sn = fmt.Sprintf("%v\r\t* %v  ", sn, v.title)
		}

		s = fmt.Sprint(s, sn)
		fmt.Println(s)
	}

	if _, err := file.WriteString(s); err != nil {
		log.Fatal(err)
	}
}

func CreateFoldersBasedOnSummary() {
	for _, v := range summary {
		CreateDirectory(v.title)
	}
}

func CreateFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(path)

		if isError(err) {
			return
		}
		defer file.Close()

		log.Println("File Created Successfully", path)
	}
}

func isError(err error) bool {
	if err != nil {
		log.Println(err)

		return true
	}

	return false
}
