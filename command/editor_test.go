package command

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveFunctionality(t *testing.T) {
	tests := []struct {
		name string
		path string
		body string
	}{
		{"Save empty file", "test.md", ""},
		{"Save test file", "../testing.md", "Current Week: {{ .Week }}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			edit := Editor{Args: []string{"cat"}, IDE: "vim"}
			want := []byte(tt.body)
			_, err := edit.SaveTemplateFile(tt.path, bytes.NewBufferString(tt.body))
			if err != nil {
				t.Fatalf("Unexpected error during save: %v", err)
			}

			if _, err := os.Stat(tt.path); err != nil {
				t.Fatalf("Template file doesn't exist: %s", tt.path)
			}

			if got, err := ioutil.ReadFile(tt.path); err != nil || !bytes.Equal(want, got) {
				t.Errorf("Want: %s, got: %s", want, got)
			}
		})
	}
}

func TestCopyContent(t *testing.T) {
	path := "test.txt"
	content := "test contents"
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	f.Close()

	err = copyContents(f, bytes.NewBufferString(content))
	if err == nil {
		t.Fatalf("ExpectedError copying contents: %v", err)
	}
}

func TestSaveFileReturnsError(t *testing.T) {
	edit := Editor{Args: []string{"cat"}, IDE: "vim"}
	path := "\000Test.md"
	body := "0000"

	_, err := edit.SaveTemplateFile(path, bytes.NewBufferString(body))
	if err == nil {
		t.Fatalf("ExpectedError copying contents: %v", err)
	}
}

func TestEditor(t *testing.T) {
	edit := NewEditor()
	if edit.Args[0] != "vim" {
		t.Errorf("Error expected 'vim' got %s", edit.Args[0])
	}
}
