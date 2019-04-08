package steps

import (
	"fmt"
	"strings"

	"github.com/engineyard/scaley/v2/pkg/scaley"
	"github.com/engineyard/scaley/v2/pkg/scaley/bash"
)

var mygroup *scaley.Group

func generateGroup(strategy string) *scaley.Group {
	if len(strategy) == 0 {
		strategy = "legion"
	}

	mygroup = &scaley.Group{
		Name: "mygroup",
		ScalingServers: []string{
			"i-00000001",
			"i-00000002",
		},
		ScalingScript: "/bin/decider",
		Strategy:      strategy,
	}

	return mygroup
}

func stubBasher(direction int) {
	bash.Run = func(command string) int {
		if command == "/bin/decider" {
			return direction
		}

		if strings.HasPrefix(command, "stop_script") {
			base := "STOP_SCRIPT"
			contexts := strings.Split(command, " ")

			switch contexts[0] {
			case "stop_script_good":
				fmt.Println(base+":", command)
			case "stop_script_bad":
				if contexts[1] == "server0" {
					fmt.Println(base+"_ERROR:", command)
					return 1
				} else {
					fmt.Println(base+":", command)
				}
			}
		}

		return 0
	}
}
