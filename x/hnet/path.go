package hnet

import (
	"net/http"
	"strconv"
	"time"

	"github.com/periaate/blume/er"
	"github.com/periaate/blume/options"
)

type PathReader struct {
	r       *http.Request
	Nerrors []er.Net
}

func (pr PathReader) Int(name string, conds ...options.Condition[int]) (res int) {
	val, err := strconv.Atoi(pr.r.PathValue(name))
	if err != nil {
		nerr := er.Free(http.StatusBadRequest,
			"path item", name,
			"in request", pr.r.RequestURI,
			"expected", "int",
			"got", pr.r.PathValue(name),
			"error", err.Error())
		pr.Nerrors = append(pr.Nerrors, nerr)
		return
	}
	res = val

	for _, opt := range conds {
		if ans := opt(res); ans != nil {
			nerr := er.Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Name,
				"because", ans.Reason,
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}

	return
}

func (pr PathReader) String(name string, conds ...options.Condition[string]) (res string) {
	res = pr.r.PathValue(name)
	if res == "" {
		nerr := er.Free(http.StatusBadRequest,
			"path item", name,
			"in request", pr.r.RequestURI,
			"expected", "string",
			"received empty string")
		pr.Nerrors = append(pr.Nerrors, nerr)
	}

	for _, opt := range conds {
		if ans := opt(res); ans != nil {
			nerr := er.Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Name,
				"because", ans.Reason,
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}

	return
}

func (pr PathReader) URL(name string, conds ...options.Condition[string]) (res string) {
	res = pr.r.PathValue(name)
	if res == "" {
		nerr := er.Free(http.StatusBadRequest,
			"path item", name,
			"in request", pr.r.RequestURI,
			"expected", "URL",
			"received empty string")
		pr.Nerrors = append(pr.Nerrors, nerr)
	}

	res = URL(res)

	for _, opt := range conds {
		if ans := opt(res); ans != nil {
			nerr := er.Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Name,
				"because", ans.Reason,
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}

	return
}

func (pr PathReader) Duration(name string, conds ...options.Condition[time.Duration]) (res time.Duration) {
	res = time.Duration(pr.Int(name))
	for _, opt := range conds {
		if ans := opt(res); ans != nil {
			nerr := er.Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Name,
				"because", ans.Reason,
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}
	return
}

func PathValue(r *http.Request) PathReader {
	return PathReader{r, []er.Net{}}
}
