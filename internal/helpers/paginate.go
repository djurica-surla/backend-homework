package helpers

import (
	"fmt"
	"net/url"
	"strconv"
)

// Paginate extracts page and page_size query values, validates them
// and returns page and offset necessary for db query.
func Paginate(query url.Values) (int, int, error) {
	// If both queries are empty, return defaul values (1 for page and 50 for page size).
	if query.Get("page") == "" && query.Get("page_size") == "" {
		page := 1
		pageSize := 50
		offset := (page - 1) * pageSize
		return pageSize, offset, nil
	}

	// If at least one is not empty, validate both of them.
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		return 0, 0, fmt.Errorf("error converting page query param to number: %w", err)
	}

	if page == 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(query.Get("page_size"))

	if err != nil {
		return 0, 0, fmt.Errorf("error converting page_size query param to number: %w", err)
	}

	switch {
	case pageSize > 50:
		pageSize = 50
	case pageSize == 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	return pageSize, offset, nil
}
