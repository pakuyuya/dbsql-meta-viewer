package textdata

import (
	"fmt"
	"regexp"
)

type searchTextdataOpt struct {
	Query      string `form:"query"`
	WordOnly   bool   `form:"wordOnly"`
	IgnoreCase bool   `form:"ignoreCase"`
	Regex      bool   `form:"regex"`
	BodyOnly   bool   `form:"bodyOnly"`
	ReadBody   bool   `form:"readBody"`
}

// repository on memory
var textDataRepository *[]Textdata

// SwitchRepository is function that do switch repository datas to argunent slices
func SwitchRepository(src *[]Textdata) {
	datas := make([]Textdata, len(*src))
	copy(datas, *src)

	textDataRepository = &datas
}

// Search is a function search textdata in shared repository
func SearchRepository(query string, with ...func(opt *searchTextdataOpt) searchTextdataOpt) ([]Textdata, error) {

	// search from shared memory(singleton, no exclusive locks)
	refRepos := textDataRepository
	if refRepos == nil {
		msg := "Repository is not initialized."
		return nil, fmt.Errorf(msg)
	}

	datas := make([]Textdata, 0, 0)
	if query == "" {
		datas = append(datas, *refRepos...)
		return datas, nil
	}

	// init & applyOptions
	opt := searchTextdataOpt{
		Query:      query,
		WordOnly:   false,
		IgnoreCase: true,
		Regex:      false,
		BodyOnly:   false,
	}
	for _, optFunc := range with {
		opt = optFunc(&opt)
	}

	regex, err := initRegexMatcher(&opt)
	if err != nil {
		return nil, err
	}

	for _, data := range *refRepos {
		if !opt.BodyOnly {
			if regex.MatchString(data.Caption + "/" + data.Namespace) {
				datas = append(datas, data)
				continue
			}
		}
		if regex.MatchString(data.Body) {
			datas = append(datas, data)
			continue
		}
	}

	return datas, nil
}

func initRegexMatcher(opt *searchTextdataOpt) (*regexp.Regexp, error) {
	s := opt.Query

	if !opt.Regex {
		s = regexp.QuoteMeta(s)
	}

	if opt.WordOnly {
		s = "(^|\\W)(" + s + ")($|\\W)"
	}

	if opt.IgnoreCase {
		s = "(?i)" + s
	}

	return regexp.Compile(s)
}

// WithRegex is a function that set option for Search() function
func WithRegex(b bool) func(opt *searchTextdataOpt) searchTextdataOpt {
	return func(opt *searchTextdataOpt) searchTextdataOpt {
		opt.Regex = b
		return *opt
	}
}

// WithWordOnly is a function that set option for Search() function
func WithWordOnly(b bool) func(opt *searchTextdataOpt) searchTextdataOpt {
	return func(opt *searchTextdataOpt) searchTextdataOpt {
		opt.WordOnly = b
		return *opt
	}
}

// WithIgnoreCase is a function that set option for Search() function
func WithIgnoreCase(b bool) func(opt *searchTextdataOpt) searchTextdataOpt {
	return func(opt *searchTextdataOpt) searchTextdataOpt {
		opt.IgnoreCase = b
		return *opt
	}
}

// WithBodyOnly is a function that set option for Search() function
func WithBodyOnly(b bool) func(opt *searchTextdataOpt) searchTextdataOpt {
	return func(opt *searchTextdataOpt) searchTextdataOpt {
		opt.BodyOnly = b
		return *opt
	}
}
