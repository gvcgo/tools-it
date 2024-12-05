package cloc

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/hhatto/gocloc"
)

/*
	type CmdOptions struct {
		ByFile         bool   `long:"by-file" description:"report results for every encountered source file"`
		SkipDuplicated bool   `long:"skip-duplicated" description:"skip duplicated files"`
		SortTag        string `long:"sort" default:"code" description:"sort based on a certain column" choice:"name" choice:"files" choice:"blank" choice:"comment" choice:"code"`
		ExcludeExt     string `long:"exclude-ext" description:"exclude file name extensions (separated commas)"`
		IncludeLang    string `long:"include-lang" description:"include language name (separated commas)"`
		Match          string `long:"match" description:"include file name (regex)"`
		NotMatch       string `long:"not-match" description:"exclude file name (regex)"`
		MatchDir       string `long:"match-d" description:"include dir name (regex)"`
		NotMatchDir    string `long:"not-match-d" description:"exclude dir name (regex)"`
		Debug          bool   `long:"debug" description:"dump debug log for developer"`
		OutputType     string `long:"output-type" default:"default" description:"output type [values: default,cloc-xml,sloccount,json]"`
		ShowLang       bool   `long:"show-lang" description:"print about all languages and extensions"`
		ShowVersion    bool   `long:"version" description:"print version info"`
	}
*/
func Cloc(
	byFile bool, skipDuplicated bool,
	sortTag string, excludeExt string,
	includeLang string,
	match string, notMatch string,
	matchDir string, notMatchDir string,
	paths ...string,
) (r string) {
	if len(paths) == 0 {
		return
	}
	clocOpts := gocloc.NewClocOptions()

	// setup option for exclude extensions
	for _, ext := range strings.Split(excludeExt, ",") {
		e, ok := gocloc.Exts[ext]
		if ok {
			clocOpts.ExcludeExts[e] = struct{}{}
		} else {
			clocOpts.ExcludeExts[ext] = struct{}{}
		}
	}

	// directory and file matching options
	if match != "" {
		clocOpts.ReMatch = regexp.MustCompile(match)
	}
	if notMatch != "" {
		clocOpts.ReNotMatch = regexp.MustCompile(notMatch)
	}
	if matchDir != "" {
		clocOpts.ReMatchDir = regexp.MustCompile(matchDir)
	}
	if notMatchDir != "" {
		clocOpts.ReNotMatchDir = regexp.MustCompile(notMatchDir)
	}

	// value for language result
	languages := gocloc.NewDefinedLanguages()

	// setup option for include languages
	for _, lang := range strings.Split(includeLang, ",") {
		if _, ok := languages.Langs[lang]; ok {
			clocOpts.IncludeLangs[lang] = struct{}{}
		}
	}

	clocOpts.SkipDuplicated = skipDuplicated

	processor := gocloc.NewProcessor(languages, clocOpts)
	result, err := processor.Analyze(paths)
	if err != nil {
		fmt.Printf("fail gocloc analyze. error: %v\n", err)
		return
	}
	r = PrepareJsonResult(result, byFile, sortTag)
	return
}

func PrepareJsonResult(result *gocloc.Result, byFile bool, sortTag string) (r string) {
	if !byFile {
		clocLangs := result.Languages
		total := result.Total
		var sortedLanguages gocloc.Languages
		for _, language := range clocLangs {
			if len(language.Files) != 0 {
				sortedLanguages = append(sortedLanguages, *language)
			}
		}
		switch sortTag {
		case "name":
			sortedLanguages.SortByName()
		case "files":
			sortedLanguages.SortByFiles()
		case "comment":
			sortedLanguages.SortByComments()
		case "blank":
			sortedLanguages.SortByBlanks()
		default:
			sortedLanguages.SortByCode()
		}

		jsonResult := gocloc.NewJSONLanguagesResultFromCloc(total, sortedLanguages)
		buf, _ := json.Marshal(jsonResult)
		r = string(buf)
	} else {
		clocFiles := result.Files
		total := result.Total

		var sortedFiles gocloc.ClocFiles
		for _, file := range clocFiles {
			sortedFiles = append(sortedFiles, *file)
		}
		switch sortTag {
		case "name":
			sortedFiles.SortByName()
		case "comment":
			sortedFiles.SortByComments()
		case "blank":
			sortedFiles.SortByBlanks()
		default:
			sortedFiles.SortByCode()
		}

		jsonResult := gocloc.NewJSONFilesResultFromCloc(total, sortedFiles)
		buf, _ := json.Marshal(jsonResult)
		r = string(buf)
	}
	return
}
