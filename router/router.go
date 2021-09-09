package router

import (
	"encoding/json"
	"fc-javascript/data"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/exec"
)

type Router struct {
}

func (r Router) Exec(c *fiber.Ctx) error {

	var Src data.Codes
	json.Unmarshal([]byte(c.Body()), &Src)

	fmt.Print("Ini isi body : ", c.Body(), "\n Ini isi struct nya : ", Src.Code)

	cmd1 := exec.Command("node", "-e", Src.Code)
	cmd1.Stdin = os.Stdin
	stdout, err := cmd1.CombinedOutput()

	if err != nil {
		c.Status(500).SendString(err.Error())
	}

	response, _ := json.Marshal(string(stdout))
	fmt.Print("\nIni out : ", string(response))
	return c.Send(response)

}
