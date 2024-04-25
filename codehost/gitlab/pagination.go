package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

func paginatedRequest(
	initFn func() interface{},
	reqFn func(interface{}, int) (interface{}, *gitlab.Response, error),
) (interface{}, error) {
	page := 1
	results, resp, err := reqFn(initFn(), page)
	if err != nil {
		return nil, err
	}

	numPages := resp.TotalPages
	for page <= numPages && resp.NextPage > page {
		page++
		results, resp, err = reqFn(results, page)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
