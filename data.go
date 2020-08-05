package main

import (
	"fmt"
	"strings"
)

type licenseType string

const (
	kW3C2015CopyrightLicense = "w3c-copyright-software-and-document"
	kCCBYNCSALicense         = "CC BY-NC-SA 4.0"
	kCCBYSALicense           = "CC BY-SA 4.0"
	kCCBYLicense             = "CC BY 4.0"
)

const (
	kSiteLicense = kCCBYNCSALicense
)

type category string

const (
	kProblemCategory    = "Problem"
	kAnalysisCategory   = "Analysis"
	kExperimentCategory = "Experimentation"
	kObjectionCategory  = "Objection"
)

type tag string

const (
	kTechnicalTag = "technical"
	kCulturalTag  = "cultural"
)

type shortName string

func (s shortName) URLPath() string {
	return fmt.Sprintf("%s%s.html", kShelvesPath, strings.ToLower(string(s)))
}

type author struct {
	DisplayName string
	ContactURL  string
}

var (
	kAuthorCJMastodonTechnology = author{
		DisplayName: "@cj@mastodon.technology",
		ContactURL:  "https://mastodon.technology/@cj",
	}
	kAuthorW3CActivityPub = author{
		DisplayName: "W3C Social Web Working Group: Christopher Lemmer Webber, Jessica Tallon, Erin Shepherd, Amy Guy, Evan Prodromou",
		ContactURL:  "https://www.w3.org/TR/activitypub/",
	}
	kAuthorAgateBlue = author{
		DisplayName: "@agateblue@mastodon.technology",
		ContactURL:  "https://mastodon.technology/@agateblue",
	}
)

// TableData lists submissions.
type tableData []data

func (td tableData) FilterByCategoriesAndTags(c, t []string) tableData {
	if len(c) == 0 && len(t) == 0 {
		return td
	}
	var out []data
	for _, d := range td {
		if d.HasAnyCategory(c) || d.HasAnyTag(t) {
			out = append(out, d)
		}
	}
	return out
}

type data struct {
	License              licenseType
	ShortName            shortName
	SourceLink           string
	Title                string
	Author               author
	Date                 string
	SubmissionDate       string
	Categories           []category
	Tags                 []tag
	ResponseToShortNames []shortName
	ResponsesShortNames  []shortName
	RelatedShortNames    []shortName
}

func (d data) HasAnyCategory(c []string) bool {
	for _, cat := range d.Categories {
		for _, oc := range c {
			if string(cat) == oc {
				return true
			}
		}
	}
	return false
}

func (d data) HasAnyTag(t []string) bool {
	for _, tag := range d.Tags {
		for _, ot := range t {
			if string(tag) == ot {
				return true
			}
		}
	}
	return false
}

var allData = tableData{
	data{
		License:              kCCBYSALicense,
		ShortName:            "COR-2",
		SourceLink:           "",
		Title:                "Problems With Relying On The Domain Name Service (DNS)",
		Author:               kAuthorCJMastodonTechnology,
		Date:                 "23 June 2020",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kProblemCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "COR-3",
		SourceLink:           "https://ilu.servus.at/getting-our-hands-dirty.html",
		Title:                "Getting our hands dirty",
		Author:               kAuthorAgateBlue,
		Date:                 "Unknown",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kW3C2015CopyrightLicense,
		ShortName:            "COR-1",
		SourceLink:           "https://www.w3.org/TR/activitypub/",
		Title:                "ActivityPub Authentication & Authorization",
		Author:               kAuthorW3CActivityPub,
		Date:                 "23 January 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kProblemCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
}
