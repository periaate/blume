package hnet

import (
	"net/http"
	"strconv"
	"time"

	"github.com/periaate/blume/gen"
)

type PathReader struct {
	r       *http.Request
	Nerrors []NetErr
}

func (pr PathReader) Int(name string, conds ...gen.Condition[int]) (res int) {
	val, err := strconv.Atoi(pr.r.PathValue(name))
	if err != nil {
		nerr := Free(http.StatusBadRequest,
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
			nerr := Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Value(),
				"because", ans.Reason(),
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}

	return
}

func (pr PathReader) String(name string, conds ...gen.Condition[string]) (res string) {
	res = pr.r.PathValue(name)
	if res == "" {
		nerr := Free(http.StatusBadRequest,
			"path item", name,
			"in request", pr.r.RequestURI,
			"expected", "string",
			"received empty string")
		pr.Nerrors = append(pr.Nerrors, nerr)
	}

	for _, opt := range conds {
		if ans := opt(res); ans != nil {
			nerr := Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Value(),
				"because", ans.Reason(),
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}

	return
}

func (pr PathReader) URL(name string, conds ...gen.Condition[string]) (res URL) {
	res = URL(pr.r.PathValue(name))
	if res == "" {
		nerr := Free(http.StatusBadRequest,
			"path item", name,
			"in request", pr.r.RequestURI,
			"expected", "URL",
			"received empty string")
		pr.Nerrors = append(pr.Nerrors, nerr)
	}

	res = URL(res)

	for _, opt := range conds {
		if ans := opt(string(res)); ans != nil {
			nerr := Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Value(),
				"because", ans.Reason(),
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}

	return
}

func (pr PathReader) Duration(name string, conds ...gen.Condition[time.Duration]) (res time.Duration) {
	res = time.Duration(pr.Int(name))
	for _, opt := range conds {
		if ans := opt(res); ans != nil {
			nerr := Free(http.StatusBadRequest,
				"path item", name,
				"in request", pr.r.RequestURI,
				"failed condition", ans.Value(),
				"because", ans.Reason(),
			)
			pr.Nerrors = append(pr.Nerrors, nerr)
			return
		}
	}
	return
}

func PathValue(r *http.Request) PathReader { return PathReader{r, []NetErr{}} }
