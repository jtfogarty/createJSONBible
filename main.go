package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//	Young's Literal Translation	American King James Version	Weymouth New Testament
type Verse struct {
	BCV        string `json:"bcv"`
	KJVText    string `json:"kvj_text"`
	ASVText    string `json:"asv_text"`
	DRText     string `json:"dr_text"`
	DarbText   string `json:"darb_text"`
	EngRText   string `json:"engr_text"`
	WebText    string `json:"web_text"`
	WldEngText string `json:"wld_eng_text"`
	YngLitText string `json:"young_lit_text"`
	AKJVText   string `json:"amer_kjv_text"`
	WeyText    string `json:"wey_text"`
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func formatVerse(v string, b bool) string {
	if !b {
		return v
	}
	rs1 := strings.Replace(v, "<i>", "", -1)
	rs2 := strings.Replace(rs1, "</i>", "", -1)
	return rs2
}

func main() {
	csvFile, err := os.Open("/Users/jeff/go/src/github.com/jtfogarty/DocuRe/bible-data/bibles.txt")
	checkErr(err)
	var doit = true
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'
	reader.Comment = '@'
	var verses []Verse
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
			break
		}
		verses = append(verses, Verse{
			BCV:        line[0],
			KJVText:    formatVerse(line[1], doit),
			ASVText:    formatVerse(line[2], doit),
			DRText:     formatVerse(line[2], doit),
			DarbText:   formatVerse(line[3], doit),
			EngRText:   formatVerse(line[4], doit),
			WebText:    formatVerse(line[5], doit),
			WldEngText: formatVerse(line[6], doit),
			YngLitText: formatVerse(line[7], doit),
			AKJVText:   formatVerse(line[8], doit),
			WeyText:    formatVerse(line[9], doit),
		})
	}
	verseJSON, err := json.Marshal(verses)
	checkErr(err)
	//fmt.Println(string(verseJson))
	err = ioutil.WriteFile("os-bibles-plain.json", verseJSON, 0644)
	checkErr(err)
}
