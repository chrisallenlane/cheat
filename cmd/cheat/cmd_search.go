package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheet"
	"github.com/cheat/cheat/internal/sheets"
)

// searchResult holds the result of searching a sheet
type searchResult struct {
	sheet     sheet.Sheet
	index     int // original position for maintaining order
	pathIndex int // which cheatpath this came from
}

// cmdSearch searches for strings in cheatsheets.
func cmdSearch(opts map[string]interface{}, conf config.Config) {

	phrase := opts["--search"].(string)

	// load the cheatsheets
	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list cheatsheets: %v\n", err)
		os.Exit(1)
	}

	// filter cheatcheats by tag if --tag was provided
	if opts["--tag"] != nil {
		cheatsheets = sheets.Filter(
			cheatsheets,
			strings.Split(opts["--tag"].(string), ","),
		)
	}

	// prepare the search pattern
	pattern := "(?i)" + phrase

	// unless --regex is provided, in which case we pass the regex unaltered
	if opts["--regex"] == true {
		pattern = phrase
	}

	// compile the regex once, outside the loop
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to compile regexp: %s, %v\n", pattern, err)
		os.Exit(1)
	}

	// determine number of workers (use number of CPUs)
	numWorkers := runtime.NumCPU()
	
	// create channels for work distribution and result collection
	jobs := make(chan searchResult, 100)
	results := make(chan searchResult, 100)
	
	// start workers
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go searchWorker(&wg, jobs, results, reg, opts, conf)
	}
	
	// send jobs to workers
	go func() {
		jobIndex := 0
		for pathIndex, pathcheats := range cheatsheets {
			// sort the cheatsheets alphabetically
			sorted := sheets.Sort(pathcheats)
			for _, sheet := range sorted {
				jobs <- searchResult{
					sheet:     sheet,
					index:     jobIndex,
					pathIndex: pathIndex,
				}
				jobIndex++
			}
		}
		close(jobs)
	}()
	
	// wait for all workers to complete and close results
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// collect all results
	var allResults []searchResult
	for result := range results {
		if result.sheet.Text != "" {
			allResults = append(allResults, result)
		}
	}
	
	// sort results by original order to maintain deterministic output
	sort.Slice(allResults, func(i, j int) bool {
		if allResults[i].pathIndex != allResults[j].pathIndex {
			return allResults[i].pathIndex < allResults[j].pathIndex
		}
		return allResults[i].index < allResults[j].index
	})
	
	// build output string
	out := ""
	for _, result := range allResults {
		sheet := result.sheet
		
		// display the cheatsheet body
		out += fmt.Sprintf(
			"%s %s\n%s\n",
			// append the cheatsheet title
			sheet.Title,
			// append the cheatsheet path
			display.Faint(fmt.Sprintf("(%s)", sheet.CheatPath), conf),
			// indent each line of content
			display.Indent(sheet.Text),
		)
	}

	// trim superfluous newlines
	out = strings.TrimSpace(out)

	// display the output
	// NB: resist the temptation to call `display.Write` multiple times in the
	// loop above. That will not play nicely with the paginator.
	display.Write(out, conf)
}

// searchWorker processes search jobs concurrently
func searchWorker(wg *sync.WaitGroup, jobs <-chan searchResult, results chan<- searchResult, 
	reg *regexp.Regexp, opts map[string]interface{}, conf config.Config) {
	
	defer wg.Done()
	
	for job := range jobs {
		sheet := job.sheet
		
		// if <cheatsheet> was provided, constrain the search only to
		// matching cheatsheets
		if opts["<cheatsheet>"] != nil && sheet.Title != opts["<cheatsheet>"] {
			continue
		}
		
		// `Search` will return text entries that match the search terms.
		// We're using it here to overwrite the prior cheatsheet Text,
		// filtering it to only what is relevant.
		sheet.Text = sheet.Search(reg)
		
		// if the sheet did not match the search, ignore it and move on
		if sheet.Text == "" {
			continue
		}
		
		// if colorization was requested, apply it here
		if conf.Color(opts) {
			sheet.Colorize(conf)
		}
		
		// send result back
		results <- searchResult{
			sheet:     sheet,
			index:     job.index,
			pathIndex: job.pathIndex,
		}
	}
}