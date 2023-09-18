package parser

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

const ruleSndType = "Individual"
const elementName = "sdnEntry"

type sndList struct {
	SndEntry []SndEntry `xml:"sdnEntry"`
}

type SndEntry struct {
	Uid       int    `xml:"uid"`
	FirstName string `xml:"firstName"`
	LastName  string `xml:"lastName"`
	SdnType   string `xml:"sdnType"`
}

var (
	list  sndList
	entry SndEntry
)

func FetchData(url string) string {
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("ACCEPT", "application/xhtml+xml,application/xml")
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal("File is unacceptable")
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func Parse(data string) sndList {
	xmlData := bytes.NewBufferString(data)
	decoder := xml.NewDecoder(xmlData)

	for i, _ := decoder.Token(); i != nil; i, _ = decoder.Token() {
		switch item := i.(type) {
		case xml.StartElement:
			if item.Name.Local == elementName {
				decoder.DecodeElement(&entry, &item)
				if entry.SdnType == ruleSndType {
					list.SndEntry = append(list.SndEntry, entry)
				}
			}
		}
	}
	return list
}
