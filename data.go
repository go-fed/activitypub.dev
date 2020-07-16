package main

// TableData lists submissions.
type tableData []rowData

type rowData struct {
	ShortName            string
	UrlPath              string
	Title                string
	Author               string
	AuthorUrl            string
	Date                 string
	ResponseToShortNames []string
	ResponseToURLs       []string
	ResponsesShortNames  []string
	ResponsesURLs        []string
	RelatedShortNames    []string
	RelatedURLs          []string
}

var problemData = tableData{
	rowData{
		"PB-2",
		"/problems/pb-2.html",
		"Domain Name Service (DNS)",
		"@cj@mastodon.technology",
		"https://mastodon.technology/@cj",
		"23 June 2020",
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	},
	rowData{
		"PB-1",
		"/problems/pb-1.html",
		"ActivityPub Authentication & Authorization",
		"W3C",
		"https://www.w3.org/TR/activitypub",
		"23 January 2018",
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	},
}

var analysisData = tableData{}
var experimentData = tableData{}
var objectionData = tableData{}
