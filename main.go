package main

import (
	"fmt"
	"strings"

	"github.com/rockaut/g2z"
)

func main() {
	fmt.Println("Say hello HPE 3Par")
}

func Echo(request *g2z.AgentRequest) (string, error) {
	return strings.Join(request.Params, " "), nil
}

func init() {
	g2z.RegisterStringItem("hpe3par.echo", "Hello world!", Echo)
}
