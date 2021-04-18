package service

import (
	"bufio"
	"encoding/base64"
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTempAttachment(t *testing.T) {
	// given
	tables := []struct {
		attachment      string
		expectedPattern string
		error           error
	}{
		{
			attachment:      "./../../resources/testAttachment.png",
			expectedPattern: "./*.png",
			error:           nil,
		},
		{
			attachment:      "",
			expectedPattern: "",
			error:           errors.New("provided attachment empty"),
		},
	}

	for _, table := range tables {
		// and
		var testFile string
		if table.error == nil {
			testFile = base64.StdEncoding.EncodeToString(createTestFile(t, table.attachment))
		}

		// when
		actual, err := createTempAttachment("./", testFile)

		// then
		assert.Equal(t, table.error, err)
		if table.error == nil {
			match, _ := filepath.Match(table.expectedPattern, actual.Name())
			assert.True(t, match)
		}
	}
}

func createTestFile(t *testing.T, fileName string) []byte {
	t.Helper()
	attachment, err := os.Open(fileName)
	reader := bufio.NewReader(attachment)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		assert.Fail(t, "failed to create test file", err)
	}
	t.Cleanup(func() {
		files, _ := filepath.Glob("./*.png")
		for _, file := range files {
			os.Remove(file)
		}
	})
	return content
}
