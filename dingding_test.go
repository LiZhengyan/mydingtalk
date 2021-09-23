package mydingtalk

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var webHook string = "https://oapi.dingtalk.com/robot/send?access_token=a5a26bff5d7836b532bfc1d8a934f515a5cd39004e14eb3da7cce4a5a580934e"
	var secret string = ""
	dt := New(webHook, WithSecret(secret))
	dt.initClient()
	markdownTitle := "test"
	markdownText := "#### test"
	if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		fmt.Println(err.Error())
	}
}
