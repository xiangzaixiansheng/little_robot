package send_util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func SendMsg(msg string) string {

	hookUrl := os.Getenv("HOOK_URL")
	strCommand := fmt.Sprintf("curl '%v' -H \"Content-Type: application/json\" -d '{\"msgtype\": \"text\",\"text\": {\"content\": \"%v\"}'", hookUrl, msg)
	fmt.Println(strCommand)
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}
	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()
	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return strings.TrimSpace(string(out_bytes))
}
