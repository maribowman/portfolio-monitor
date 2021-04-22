package repository

import (
	"github.com/stretchr/testify/assert"
	"maribowman/portfolio-monitor/app/model"
	"testing"
)

func TestPush(t *testing.T) {
	message := model.Message{
		Message:     "test from code",
		Sender:      "+4915226264500",
		Recipients:  []string{"+4915226264500"},
		Attachments: nil,
	}
	err := NewSignalClient().Push(model.Holding{}, message)
	assert.NoError(t, err)
}
