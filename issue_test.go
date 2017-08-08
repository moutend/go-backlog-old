package backlog

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const testdata = "testdata"

func Test(t *testing.T) {
	var err error
	var file []byte
	var issue Issue
	if file, err = ioutil.ReadFile(filepath.Join(testdata, "issue.json")); err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(file, &issue); err != nil {
		t.Fatal(err)
	}
	return
}
