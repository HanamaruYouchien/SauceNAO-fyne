package saucenao

type Response struct {
	Header  Header   `json:"header"`
	Results []Result `json:"results"`
}

type Header struct {
	UserId           int `json:"user_id,string"`
	AccountType      int `json:"account_type,string"`
	ShortLimit       int `json:"short_limit,string"`
	LongLimit        int `json:"long_limit,string"`
	LongRemaining    int `json:"long_remaining"`
	ShortRemaining   int `json:"short_remaining"`
	Status           int `json:"status"`
	ResultsRequested int `json:"results_requested,string"`
	Index            map[int]struct {
		Status   int `json:"status"`
		ParentId int `json:"parent_id"`
		Id       int `json:"id"`
		Results  int `json:"results"`
	} `json:"index"`
	SearchDepth       int     `json:"search_depth,string"`
	MinimumSimilarity float64 `json:"minimum_similarity"`
	QueryImageDisplay string  `json:"query_image_display"`
	QueryImage        string  `json:"query_image"`
	ResultsReturned   int     `json:"results_returned"`
}

type ResultData map[string]interface{}

type Result struct {
	Header struct {
		Similarity float64 `json:"similarity,string"`
		Thumbnail  string  `json:"thumbnail"`
		IndexId    int     `json:"index_id"`
		IndexName  string  `json:"index_name"`
		Dupes      int     `json:"dupes"`
		Hidden     int     `json:"hidden"`
	} `json:"header"`
	Data ResultData `json:"data"`
}
