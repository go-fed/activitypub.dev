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
	kProblemsPage           = "/problems.html"
	kProblemsPath           = "/problems/"
	kAnalysesPage           = "/analyses.html"
	kAnalysesPath           = "/analyses/"
	kExperimentsPage        = "/experiments.html"
	kExperimentsPath        = "/experiments/"
	kObjectionsPage         = "/objections.html"
	kObjectionsPath         = "/objections/"
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
	p, hf = kProblemsPage, handleTemplate(kSiteTableTemplate, problemData)
	http.HandleFunc(p, hf)
	p, hf = kProblemsPath, handleTemplate(kSiteSubmissionTemplate, problemData)
	http.HandleFunc(p, hf)
	// Analyses
	p, hf = kAnalysesPage, handleTemplate(kSiteTableTemplate, analysisData)
	http.HandleFunc(p, hf)
	p, hf = kAnalysesPath, handleTemplate(kSiteSubmissionTemplate, analysisData)
	http.HandleFunc(p, hf)
	// Experiments
	p, hf = kExperimentsPage, handleTemplate(kSiteTableTemplate, experimentData)
	http.HandleFunc(p, hf)
	p, hf = kExperimentsPath, handleTemplate(kSiteSubmissionTemplate, experimentData)
	http.HandleFunc(p, hf)
	// Objections
	p, hf = kObjectionsPage, handleTemplate(kSiteTableTemplate, objectionData)
	http.HandleFunc(p, hf)
	p, hf = kObjectionsPath, handleTemplate(kSiteSubmissionTemplate, objectionData)
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

func handleTemplate(siteTemplate string, data interface{}) http.HandlerFunc {
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

		tmpl, err := template.ParseFiles(cp, lp, fp)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		templateData := struct {
			Data interface{}
			Path string
		}{
			Data: data,
			Path: req.URL.Path,
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
