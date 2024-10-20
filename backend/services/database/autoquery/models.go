package autoquery

type LoaderParams struct {
	LoadType       string                 `json:"loadType"`
	MaxId          int                    `json:"maxId"`
	MinId          int                    `json:"minId"`
	ActiveColumns  []string               `json:"activeColumns"`
	FilterModels   map[string]FilterModel `json:"filterModels"`
	OrderBy        string                 `json:"orderBy"`
	OrderDirection string                 `json:"orderDirection"`
	PageId         int                    `json:"pageId"`
	PageSize       int                    `json:"pageSize"`
}

type FilterModel struct {
	Key        string   `json:"key"`
	FilterType string   `json:"filterType"`
	Operator   string   `json:"operator"`
	Value      []string `json:"value"`
}

const (
	OperatorEqual              OperatorType = "equal"
	OperatorNotEqual           OperatorType = "not_equal"
	OperatorContains           OperatorType = "contains"
	OperatorNotContains        OperatorType = "not_contains"
	OperatorStartsWith         OperatorType = "starts_with"
	OperatorEndsWith           OperatorType = "ends_with"
	OperatorIsNull             OperatorType = "is_null"
	OperatorIsNotNull          OperatorType = "is_not_null"
	OperatorBetween            OperatorType = "between"
	OperatorNotBetween         OperatorType = "not_between"
	OperatorGreaterThan        OperatorType = "greater_than"
	OperatorLessThan           OperatorType = "less_than"
	OperatorGreaterThanOrEqual OperatorType = "greater_than_or_equal"
	OperatorLessThanOrEqual    OperatorType = "less_than_or_equal"
)

type OperatorType = string
