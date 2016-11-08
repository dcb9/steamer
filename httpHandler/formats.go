package httpHandler

type request struct {
	Route string
	Uid string
}

type searchRequest struct {
	request
	Url string
}

type downloadRequest struct {
	request
	Id     int64
	Url    string
	SIndex string `json:"stream_index"`
}

type response struct {
	Uid string `json:"uid"`
	Success int `json:"success"`
	Data interface{} `json:"data"`
}

func (r *request) isSearchRequest() bool {
	return r.Route == "/search"
}

func (r *request) isDownloadRequest() bool {
	return r.Route == "/download"
}
