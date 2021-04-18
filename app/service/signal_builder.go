package service

import (
	"fmt"
	"maribowman/signal-transmitter/app/model"
)

type cmdLineBuilder struct {
	builder       []string
	attachmentDir string
}

func NewCmdLineBuilder(attachmentDir string) *cmdLineBuilder {
	clb := &cmdLineBuilder{
		builder:       []string{"./signal-cli"},
		attachmentDir: attachmentDir,
	}
	return clb
}

func (clb *cmdLineBuilder) from(sender string) *cmdLineBuilder {
	clb.builder = append(clb.builder, "-u")
	clb.builder = append(clb.builder, sender)
	return clb
}

func (clb *cmdLineBuilder) sendMessage(msg string) *cmdLineBuilder {
	clb.builder = append(clb.builder, fmt.Sprintf("send -m \"%s\"", msg))
	return clb
}

func (clb *cmdLineBuilder) to(recipient model.Recipient) *cmdLineBuilder {
	if recipient.IsGroup {
		clb.builder = append(clb.builder, "-g", recipient.GroupID)
	} else {
		clb.builder = append(clb.builder, recipient.Receivers...)
	}
	return clb
}

func (clb *cmdLineBuilder) withAttachment(base64Attachment string) *cmdLineBuilder {
	if len(base64Attachment) == 0 {
		return clb
	}
	attachment, err := createTempAttachment(clb.attachmentDir, base64Attachment)
	if err != nil {
		return clb
	}
	clb.builder = append(clb.builder, "-a", attachment.Name())
	return clb
}

func (clb *cmdLineBuilder) build() []string {
	return clb.builder
}
