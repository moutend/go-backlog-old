package backlog

import (
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

	client, _ = New("spaceName", "XXXXXXXX")
	client.root = root

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

func TestGetProjects(t *testing.T) {
	_, err := client.GetProjects(nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetIssues(t *testing.T) {
	_, err := client.GetIssues(nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetIssue(t *testing.T) {
	_, err := client.GetIssue(12345)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestDeleteIssue(t *testing.T) {
	_, err := client.DeleteIssue(12345)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestSetIssue(t *testing.T) {
	_, err := client.SetIssue(12345, nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetIssuesCount(t *testing.T) {
	_, err := client.GetIssuesCount(nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetStatuses(t *testing.T) {
	_, err := client.GetStatuses()
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetIssueTypes(t *testing.T) {
	_, err := client.GetIssueTypes(12345)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestGetPriorities(t *testing.T) {
	_, err := client.GetPriorities()
	if err != nil {
		t.Fatal(err)
	}
	return
}
