package backlog

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	client   *Client
	testdata string
)

func setup() (*httptest.Server, error) {
	testdata = filepath.Join("testdata", "api", "v2")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(
			testdata,
			strings.Replace(r.URL.Path, "/", string(os.PathSeparator), -1),
			r.Method+".json")

		file, err := ioutil.ReadFile(path)
		if err != nil {
			panic(path + " not found")
		}

		w.Write(file)
		return
	}))

	root, err := url.Parse(server.URL)
	if err != nil {
		return nil, err
	}

	client = &Client{
		root: root,
	}

	return server, nil
}

func teardown(server *httptest.Server) error {
	server.Close()

	return nil
}

func TestMain(m *testing.M) {
	server, err := setup()
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()

	err = teardown(server)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(code)

	return
}

func TestGetIssuesContext(t *testing.T) {
	ctx := context.Background()
	_, err := client.GetIssuesContext(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetIssueContext(t *testing.T) {
	ctx := context.Background()
	_, err := client.GetIssueContext(ctx, 12345)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestSetIssueContext(t *testing.T) {
	ctx := context.Background()
	_, err := client.SetIssueContext(ctx, 12345, nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetIssuesCountContext(t *testing.T) {
	ctx := context.Background()
	_, err := client.GetIssuesCountContext(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetStatusesContext(t *testing.T) {
	ctx := context.Background()
	_, err := client.GetStatusesContext(ctx)
	if err != nil {
		t.Fatal(err)
	}
	return
}
