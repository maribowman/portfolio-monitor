package service

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"maribowman/signal-transmitter/app/model"
	"strings"
	"testing"
)

func TestCmdLineBuilder(t *testing.T) {
	// given
	tables := []struct {
		sender     string
		recipient  model.Recipient
		message    string
		attachment string
		expected   []string
	}{
		{
			sender: "mari",
			recipient: model.Recipient{
				IsGroup:   false,
				GroupID:   "",
				Receivers: []string{"everyone"},
			},
			message:    "I love bacon!",
			attachment: "",
			expected:   []string{"-u mari send -m \"I love bacon!\" everyone"},
		},
		{
			sender: "mari",
			recipient: model.Recipient{
				IsGroup:   true,
				GroupID:   "<group-id>",
				Receivers: []string{},
			},
			message:    "I love bacon!",
			attachment: "",
			expected:   []string{"-u mari send -m \"I love bacon!\" -g <group-id>"},
		},
		{
			sender: "mari",
			recipient: model.Recipient{
				IsGroup:   false,
				GroupID:   "",
				Receivers: []string{"everyone"},
			},
			message:    "I love bacon!",
			attachment: "./../../resources/testAttachment.png",
			expected:   []string{"-u mari send -m \"I love bacon!\" everyone -a ", ".png"},
		},
	}

	for _, table := range tables {
		var testFile string
		if len(table.attachment) != 0 {
			testFile = base64.StdEncoding.EncodeToString(createTestFile(t, table.attachment))
		}

		// when
		signalCmd := NewCmdLineBuilder("./").
			from(table.sender).
			sendMessage(table.message).
			to(table.recipient).
			withAttachment(testFile).
			build()
		actual := strings.Join(signalCmd, " ")

		// then
		if len(table.attachment) != 0 {
			assert.True(t, strings.HasPrefix(actual, table.expected[0]))
			assert.True(t, strings.HasSuffix(actual, table.expected[1]))
		} else {
			assert.Equal(t, table.expected[0], actual)
		}
	}
}
