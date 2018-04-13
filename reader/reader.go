package reader

import (
	wof_reader "github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HTTPReader struct {
	wof_reader.Reader
	root *url.URL
}

func NewHTTPReader(root string) (wof_reader.Reader, error) {

	root_url, err := url.Parse(root)

	if err != nil {
		return nil, err
	}

	r := HTTPReader{
		root: root_url,
	}

	return &r, nil
}

func (r *HTTPReader) Read(key string) (io.ReadCloser, error) {

	url := r.root.String() + key

	if !strings.HasSuffix(r.root.String(), "/") && !strings.HasPrefix(key, "/") {
		url = r.root.String() + "/" + key
	}

	rsp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return rsp.Body, nil
}
