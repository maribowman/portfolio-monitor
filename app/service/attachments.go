package service

import (
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"os"
)

func createTempAttachment(directory, attachment string) (*os.File, error) {
	if len(attachment) == 0 {
		return nil, errors.New("provided attachment empty")
	}
	decoded, err := base64.StdEncoding.DecodeString(attachment)
	if err != nil {
		return nil, err
	}

	fileType, err := filetype.Get(decoded)
	if err != nil {
		return nil, err
	}
	tempFilePath := directory + uuid.New().String() + "." + fileType.Extension
	//tempFilePath := "/tempAttachments/" + uuid.New().String() + "." + fileType.Extension
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()
	if _, err := tempFile.Write(decoded); err != nil {
		os.Remove(directory)
		return nil, err
	}
	if err := tempFile.Sync(); err != nil {
		os.Remove(directory)
		return nil, err
	}
	tempFile.Close()
	return tempFile, nil
}
