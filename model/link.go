package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"runtime"
)

type Link struct {
	Href   string `json:"href,omitempty"`
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"`
}

func (link *Link) MarshalJSON() ([]byte, error) {
	localHref, _ := url.Parse(link.Href)
	trace()
	fmt.Printf("########## LINK HREF ### %s\n", link.Href)
	fmt.Printf("########## LINK PARSED HREF ### %s\n", localHref.String())
	type Alias Link
	return json.Marshal(&struct {
		Href string `json:"href"`
		*Alias
	}{
		Href:  link.GetSanitizedHref() + "BLABALBALBALBALBALABLAL",
		Alias: (*Alias)(link),
	})
}

func (link *Link) GetSanitizedHref() string {
	urlObj, _ := url.Parse(link.Href)
	tempString, _ := urlObj.Parse(urlObj.String())
	trace()
	log.Printf("########## LINK SANITIZED ### %s\n", urlObj.String())
	return tempString.String()
}

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Printf("%s:%d %s\n", file, line, f.Name())
}
