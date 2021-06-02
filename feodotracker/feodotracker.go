package feodotracker

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// FeodoTrackerRules implements ruler.Ruler interface
type FeodoTrackerRuler struct {
	Url         string
	Lines       []string
	Rules       []string
	writer      *bufio.Writer
	Description string
}

// GetRules pointer receiver
func (f *FeodoTrackerRuler) GetRules() error {
	f_url := f.Url

	r, err := http.Get(f_url)
	if err != nil {
		fmt.Printf("URL:%s is not reachable\n", f_url)
		return err
	}
	defer r.Body.Close()

	sc := r.StatusCode
	if sc/100 != 2 {
		fmt.Printf("Error reaching (url=%s, status code=%v)\n", f_url, sc)
		return fmt.Errorf("Error - apigeeids: status code %v when requesting Url=%s", sc, f_url)
	}

	scanner := bufio.NewScanner(r.Body)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "alert") {
			f.Lines = append(f.Lines, l)
		}
	}
	return nil
}

// SetRules pointer receiver
func (f *FeodoTrackerRuler) SetRules() error {
	for _, l := range f.Lines {
		f.Rules = append(f.Rules, l)
	}
	return nil
}

// WriteRules pointer receiver
// func: WriteLines to write line per line on a file
func (f *FeodoTrackerRuler) WriteRules() error {
	for _, r := range f.Rules {
		fmt.Fprintln(f.writer, r)
	}
	return nil
}

//  func: New(url string, w *bufio.Writer, desc string)
func New(url string, w *bufio.Writer, desc string) (*FeodoTrackerRuler, error) {
	var f *FeodoTrackerRuler
	// instantiate a GoogleAPIRequester
	f = &FeodoTrackerRuler{
		Url:         url,
		Lines:       make([]string, 0, 0),
		Rules:       make([]string, 0, 0),
		writer:      w,
		Description: desc,
	}
	currentTime := time.Now()
	date := fmt.Sprintf("# %s - %s", desc, currentTime.Format("2006-01-02 15:04:05"))
	f.Lines = append(f.Lines, date)
	return f, nil
}
