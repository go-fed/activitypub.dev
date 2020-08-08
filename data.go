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
	kTechnicalTag   = "technical"
	kCulturalTag    = "cultural"
	kDatashardsTag  = "datashards"
	kHeyFYITag      = "hey-fyi"
	kC2STag         = "c2s"
	kGovernanceTag  = "governance"
	kCooperationTag = "cooperation"
	kBehaviorsTag   = "behaviors"
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
	kAuthorGargron = author{
		DisplayName: "@Gargron@mastodon.social",
		ContactURL:  "https://mastodon.social/@Gargron",
	}
	kAuthorHypolite = author{
		DisplayName: "@hypolite@friendica.mrpetovan.com",
		ContactURL:  "https://friendica.mrpetovan.com/profile/hypolite",
	}
	kAuthorEmacsen = author{
		DisplayName: "@emacsen@emacsen.net",
		ContactURL:  "https://emacsen.net/@emacsen",
	}
	kAuthorCwebber = author{
		DisplayName: "@cwebber@octodon.social",
		ContactURL:  "https://octodon.social/@cwebber",
	}
	kAuthorMariusor = author{
		DisplayName: "@mariusor@metalhead.club",
		ContactURL:  "https://metalhead.club/@mariusor",
	}
	kAuthorLain = author{
		DisplayName: "@lain@lain.com",
		ContactURL:  "https://lain.com/users/lain",
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
		if d.HasAllCategories(c) && d.HasAllTags(t) {
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

func (d data) HasAllCategories(c []string) bool {
	all := true
	for _, oc := range c {
		found := false
		for _, cat := range d.Categories {
			if string(cat) == oc {
				found = true
				break
			}
		}
		all = all && found
	}
	return all
}

func (d data) HasAllTags(t []string) bool {
	all := true
	for _, ot := range t {
		found := false
		for _, tag := range d.Tags {
			if string(tag) == ot {
				found = true
				break
			}
		}
		all = all && found
	}
	return all
}

var allData = tableData{
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-18",
		SourceLink:           "https://blog.soykaf.com/post/encryption/",
		Title:                "Lain Thought on End-To-End Encryption with AP Characteristics for a New Era",
		Author:               kAuthorLain,
		Date:                 "28 May 2020",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kExperimentCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-17",
		SourceLink:           "https://blog.soykaf.com/post/privacy-and-tracking-on-the-fediverse/",
		Title:                "Privacy and Tracking on the Fediverse",
		Author:               kAuthorLain,
		Date:                 "25 March 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag, kBehaviorsTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-16",
		SourceLink:           "https://blog.soykaf.com/post/activity-pub-in-pleroma/",
		Title:                "ActivityPub in Pleroma",
		Author:               kAuthorLain,
		Date:                 "4 March 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kExperimentCategory},
		Tags:                 []tag{kTechnicalTag, kHeyFYITag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-15",
		SourceLink:           "https://blog.soykaf.com/post/pleroma-encyclical-activity-pub/",
		Title:                "Pleroma Encyclical: ActivityPub",
		Author:               kAuthorLain,
		Date:                 "10 February 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kProblemCategory},
		Tags:                 []tag{kTechnicalTag, kCooperationTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "GHUB-1",
		SourceLink:           "https://github.com/go-ap/fedbox/blob/master/doc/c2s.md",
		Title:                "Fed::BOX as an ActivityPub server supporting C2S interactions",
		Author:               kAuthorMariusor,
		Date:                 "20 June 2020",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kExperimentCategory},
		Tags:                 []tag{kTechnicalTag, kC2STag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-14",
		SourceLink:           "http://dustycloud.org/blog/content-addressed-vocabulary/",
		Title:                "Content Addressed Vocabulary",
		Author:               kAuthorCwebber,
		Date:                 "26 February 2020",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-13",
		SourceLink:           "http://dustycloud.org/blog/state-of-spritely-2020-02/",
		Title:                "State of Spritely for February 2020",
		Author:               kAuthorCwebber,
		Date:                 "10 February 2020",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kTechnicalTag, kDatashardsTag, kHeyFYITag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-12",
		SourceLink:           "http://dustycloud.org/blog/2019-10-01-updates/",
		Title:                "Updates: ActivityPub Conference, and more",
		Author:               kAuthorCwebber,
		Date:                 "1 October 2019",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag, kDatashardsTag, kHeyFYITag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-11",
		SourceLink:           "http://dustycloud.org/blog/on-standards-divisions-collaboration/",
		Title:                "On standards divisions and collaboration (or: Why can't the decentralized social web people just get along?)",
		Author:               kAuthorCwebber,
		Date:                 "25 January 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag, kCooperationTag},
		ResponseToShortNames: []shortName{"WEB-10"},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-10",
		SourceLink:           "http://dustycloud.org/blog/activitypub-is-a-w3c-recommendation/",
		Title:                "ActivityPub is a W3C Recommendation",
		Author:               kAuthorCwebber,
		Date:                 "23 January 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kTechnicalTag, kHeyFYITag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{"WEB-11"},
		RelatedShortNames:    []shortName{"WEB-9"},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-9",
		SourceLink:           "http://dustycloud.org/blog/an-even-more-distributed-activitypub/",
		Title:                "An even more distributed ActivityPub",
		Author:               kAuthorCwebber,
		Date:                 "6 October 2016",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{"WEB-10"},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-8",
		SourceLink:           "https://write.emacsen.net/datashards-update-october-2019",
		Title:                "Datashards Update October 2019",
		Author:               kAuthorEmacsen,
		Date:                 "19 October 2019",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kExperimentCategory},
		Tags:                 []tag{kTechnicalTag, kDatashardsTag, kHeyFYITag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-7",
		SourceLink:           "https://write.emacsen.net/thoughts-on-canonical-s-expressions",
		Title:                "Thoughts on Canonical S-Expressions",
		Author:               kAuthorEmacsen,
		Date:                 "2 October 2019",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kTechnicalTag, kDatashardsTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "WEB-6",
		SourceLink:           "https://write.emacsen.net/on-long-form-blogging",
		Title:                "On Long-Form Blogging",
		Author:               kAuthorEmacsen,
		Date:                 "23 September 2019",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag, kBehaviorsTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-5",
		SourceLink:           "https://blog.mrpetovan.com/uncategorized/what-makes-the-fediverse-better-than-twitter/",
		Title:                "What makes the Fediverse better than Twitter",
		Author:               kAuthorHypolite,
		Date:                 "18 May 2020",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag, kBehaviorsTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-4",
		SourceLink:           "https://blog.joinmastodon.org/2018/07/how-to-make-friends-and-verify-requests/",
		Title:                "How to make friends and verify requests",
		Author:               kAuthorGargron,
		Date:                 "3 July 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kExperimentCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{"WEB-3"},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-3",
		SourceLink:           "https://blog.joinmastodon.org/2018/06/how-to-implement-a-basic-activitypub-server/",
		Title:                "How to implement a basic ActivityPub server",
		Author:               kAuthorGargron,
		Date:                 "23 June 2018",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kExperimentCategory},
		Tags:                 []tag{kTechnicalTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{"WEB-4"},
	},
	data{
		License:              kCCBYNCSALicense,
		ShortName:            "WEB-2",
		SourceLink:           "https://ilu.servus.at/getting-our-hands-dirty.html",
		Title:                "Getting our hands dirty",
		Author:               kAuthorAgateBlue,
		Date:                 "Unknown",
		SubmissionDate:       "8 August 2020",
		Categories:           []category{kAnalysisCategory},
		Tags:                 []tag{kCulturalTag, kGovernanceTag, kCooperationTag},
		ResponseToShortNames: []shortName{},
		ResponsesShortNames:  []shortName{},
		RelatedShortNames:    []shortName{},
	},
	data{
		License:              kCCBYSALicense,
		ShortName:            "DS-1",
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
		License:              kW3C2015CopyrightLicense,
		ShortName:            "WEB-1",
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
