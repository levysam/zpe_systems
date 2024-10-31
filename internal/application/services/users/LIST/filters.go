package users

import "go-skeleton/internal/application/providers/filters"

func (r *Request) SetFiltersRules() error {
	parseErr := r.Filters.Parse(
		map[string]string{},
		map[string]filters.FilterData{},
	)

	if parseErr != nil {
		return parseErr
	}
	return nil
}
