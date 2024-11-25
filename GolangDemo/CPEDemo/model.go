package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type CPEStaticRequest struct {
	Keyword string `json:"keyword"`
	TimeLTE int64  `json:"time_lte"`
	TimeGTE int64  `json:"time_gte"`
}

func (req CPEStaticRequest) validate() error {
	req.Keyword = strings.TrimSpace(req.Keyword)
	if req.TimeGTE < 0 || req.TimeLTE < 0 {
		return errors.New("invalid time")
	}

	if req.TimeGTE > 0 && req.TimeLTE == 0 {
		req.TimeLTE = time.Now().UTC().UnixMilli()
	}

	if req.TimeGTE > 0 && req.TimeLTE < req.TimeGTE {
		return errors.New("invalid time")
	}

	return nil
}

func (req CPEStaticRequest) genAggsQuery() map[string]interface{} {
	var mustQuery []map[string]interface{}
	mustQuery = append(mustQuery, map[string]interface{}{
		"nested": map[string]interface{}{
			"path": "attribute",
			"query": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"wildcard": map[string]interface{}{
							"attribute.vendor": fmt.Sprintf("*%s*", req.Keyword),
						},
					},
					{
						"wildcard": map[string]interface{}{
							"attribute.version": fmt.Sprintf("*%s*", req.Keyword),
						},
					},
					{
						"wildcard": map[string]interface{}{
							"attribute.product": fmt.Sprintf("*%s*", req.Keyword),
						},
					},
				},
			},
		},
	})

	lastModified := map[string]interface{}{}
	if req.TimeGTE != 0 {
		lastModified["gte"] = req.TimeGTE
	}
	if req.TimeLTE != 0 {
		lastModified["lte"] = req.TimeLTE
	}
	if req.TimeGTE != 0 || req.TimeLTE != 0 {
		dateRange := map[string]interface{}{
			"range": map[string]interface{}{
				"modified": lastModified,
			},
		}
		mustQuery = append(mustQuery, dateRange)
	}

	var query map[string]interface{}
	query["query"] = map[string]interface{}{
		"bool": map[string]interface{}{
			"must": mustQuery,
		},
	}
	query["size"] = 0
	query["aggs"] = map[string]interface{}{
		"nested_attributes": map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "attribute",
			},
			"aggs": map[string]interface{}{
				"part": map[string]interface{}{
					"filters": map[string]interface{}{
						"filters": map[string]interface{}{
							"a": map[string]interface{}{
								"terms": map[string]interface{}{
									"attribute.part": "a",
								},
							},
							"o": map[string]interface{}{
								"terms": map[string]interface{}{
									"attribute.part": "o",
								},
							},
							"h": map[string]interface{}{
								"terms": map[string]interface{}{
									"attribute.part": "h",
								},
							},
						},
					},
				},
			},
		},
	}

	return query
}
