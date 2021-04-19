package service

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"maribowman/signal-transmitter/app/model"
	"net/http"
	"os/exec"
	"time"
)

func SendMessage(context *gin.Context) {
	var dto model.Message
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, "could not bind dta")
		return
	}

	cmdline := NewCmdLineBuilder("./").
		from(dto.Sender).
		sendMessage(dto.Text).
		to(dto.Recipient).
		withAttachment(dto.Attachment).
		build()

	cmd := exec.Command("signal-cli", cmdline...)
	var errBuffer bytes.Buffer
	cmd.Stderr = &errBuffer
	if err := cmd.Start(); err != nil {
		log.Println("signal-cli:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "operation timed out. killing process."})
		return
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(20 * time.Second):
		cmd.Process.Kill()
		context.JSON(http.StatusInternalServerError, gin.H{"error": "operation timed out. killing process."})
		return
	case err := <-done:
		if err != nil {
			log.Println("signal-cli:", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("signal-cli:%s", errBuffer.String())})
			return
		}
	}
	context.Status(http.StatusNoContent)
}
