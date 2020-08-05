package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	kFaviconFile            = "./static/favicon.ico"
	kFaviconPath            = "/favicon.ico"
	kStaticDir              = "./static"
	kStaticPath             = "/static/"
	kRootPath               = "/"
	kTemplatesDir           = "./templates"
	kCommonTemplate         = "common.tmpl"
	kSiteTemplate           = "site.tmpl"
	kSiteTableTemplate      = "site_table.tmpl"
	kSiteSubmissionTemplate = "site_submission.tmpl"
	kSiteTemplateName       = "site"
	kDirIndex               = "index.html"
	kShelvesPage            = "/shelves.html"
	kShelvesPath            = "/shelves/"
	kCategoryQuery          = "category"
	kTagQuery               = "tag"
)

var (
	certFileFlag  = flag.String("cert", "tls.crt", "Path to certificate public file")
	keyFileFlag   = flag.String("key", "tls.key", "Path to certificate private key file")
	httpPortFlag  = flag.String("http_port", ":http", "Server http port")
	httpsPortFlag = flag.String("https_port", ":https", "Server https port")
	httpsFlag     = flag.Bool("https", false, "Whether to use https")
)

func main() {
	flag.Parse()
	if len(*httpPortFlag) == 0 {
		log.Println("must provide --http_port")
		return
	}
	if *httpsFlag {
		if len(*certFileFlag) == 0 {
			log.Println("must provide --cert with --https flag")
			return
		} else if len(*keyFileFlag) == 0 {
			log.Println("must provide --key with --https flag")
			return
		} else if len(*httpsPortFlag) == 0 {
			log.Println("must provide --https_port with --https flag")
			return
		}
	}

	p, hf := handleFavicon()
	http.HandleFunc(p, hf)
	p, h := handleStatic()
	http.Handle(p, h)
	p, hf = kRootPath, handleTemplate(kSiteTemplate, nil)
	http.HandleFunc(p, hf)
	// Problems
	p, hf = kShelvesPage, handleTemplate(kSiteTableTemplate, allData)
	http.HandleFunc(p, hf)
	p, hf = kShelvesPath, handleTemplate(kSiteSubmissionTemplate, allData)
	http.HandleFunc(p, hf)

	if *httpsFlag {
		log.Printf("Https & http-redirect: Listening on %s and %s\n", *httpPortFlag, *httpsPortFlag)
		go func() {
			redir := http.NewServeMux()
			redir.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Connection", "close")
				http.Redirect(w, req, fmt.Sprintf("https://%s%s", req.Host, req.URL), http.StatusFound)
			})
			err := http.ListenAndServe(*httpPortFlag, redir)
			if err != nil {
				log.Fatal(err)
			}
		}()
		err := http.ListenAndServeTLS(*httpsPortFlag, *certFileFlag, *keyFileFlag, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Http-only: Listening on %s\n", *httpPortFlag)
		err := http.ListenAndServe(*httpPortFlag, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handleFavicon() (string, http.HandlerFunc) {
	return kFaviconPath, func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, kFaviconFile)
	}
}

func handleStatic() (string, http.Handler) {
	fs := http.FileServer(http.Dir(kStaticDir))
	return kStaticPath, http.StripPrefix(kStaticPath, fs)
}

func handleTemplate(siteTemplate string, data tableData) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cp := filepath.Join(kTemplatesDir, kCommonTemplate)
		lp := filepath.Join(kTemplatesDir, siteTemplate)
		fp := filepath.Join(kTemplatesDir, filepath.Clean(req.URL.Path))

		info, err := os.Stat(fp)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, req)
				return
			}
		} else if info.IsDir() {
			fp = filepath.Join(fp, kDirIndex)
			info, err = os.Stat(fp)
			if err != nil {
				if os.IsNotExist(err) {
					http.NotFound(w, req)
					return
				}
			} else if info.IsDir() {
				http.NotFound(w, req)
				return
			}
		}

		// Process any query filtering
		query := req.URL.Query()
		categories, _ := query[kCategoryQuery]
		tags, _ := query[kTagQuery]
		u := req.URL
		addQueryToU := func(k, v string) string {
			c := *u
			q := c.Query()
			q.Add(k, v)
			c.RawQuery = q.Encode()
			return (&c).RequestURI()
		}
		removeQueryFromU := func(k, v string) string {
			c := *u
			q := c.Query()
			vals, ok := q[k]
			if ok {
				di := -1
				for i := 0; i < len(vals); i++ {
					if vals[i] == v {
						di = i
						break
					}
				}
				if di >= 0 {
					copy(vals[di:], vals[di+1:])
					vals[len(vals)-1] = ""
					vals = vals[:len(vals)-1]
					q[k] = vals
				}
			}
			c.RawQuery = q.Encode()
			return (&c).RequestURI()
		}

		fm := template.FuncMap{
			"RequestURIWithQueryCategory": func(v category) string {
				return addQueryToU(kCategoryQuery, string(v))
			},
			"RequestURIWithQueryTag": func(t tag) string {
				return addQueryToU(kTagQuery, string(t))
			},
			"RequestURIWithoutQueryCategory": func(v string) string {
				return removeQueryFromU(kCategoryQuery, v)
			},
			"RequestURIWithoutQueryTag": func(t string) string {
				return removeQueryFromU(kTagQuery, t)
			},
		}
		tmpl := template.New("")
		tmpl = tmpl.Funcs(fm)
		tmpl, err = tmpl.ParseFiles(cp, lp, fp)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		templateData := struct {
			Data           tableData
			Path           string
			SiteLicense    licenseType
			CategoryFilter []string
			TagFilter      []string
		}{
			Data:           data.FilterByCategoriesAndTags(categories, tags),
			Path:           req.URL.Path,
			SiteLicense:    kSiteLicense,
			CategoryFilter: categories,
			TagFilter:      tags,
		}
		err = tmpl.ExecuteTemplate(w, kSiteTemplateName, templateData)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}
}
