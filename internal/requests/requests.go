package requests

type FormData struct {
	URL          string `json:"request_url"`
	RequestCount int    `json:"request_count"`
	Parallel     bool   `json:"parallel"`
	RequestType  string `json:"request_type"`
	JSONData     string `json:"json_data"`
	Randomize    bool   `json:"randomize"`
}
