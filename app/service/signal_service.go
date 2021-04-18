package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"maribowman/signal-transmitter/app/model"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func SendMessage(context *gin.Context) {
	var messageDTA model.Message
	if err := context.ShouldBindJSON(&messageDTA); err != nil {
		context.JSON(http.StatusBadRequest, "could not bind dta")
		return
	}

	cmdline := NewCmdLineBuilder("./").
		from(messageDTA.Sender).
		sendMessage(messageDTA.Text).
		to(messageDTA.Recipient).
		withAttachment(messageDTA.Attachment).
		build()

	log.Println(strings.Join(cmdline, " "))

	cmd := exec.Command("/home/mari/Dev/signal-cli-0.8.1/bin/signal-cli", cmdline...)
	if err := cmd.Start(); err != nil {
		log.Println("signal-cli:", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	log.Println("waiting...")
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	log.Println("still waiting...")

	time.Sleep(20 * time.Second)
	context.Status(http.StatusNoContent)
}
