package steps

import (
	"fmt"
	"strings"

	"github.com/engineyard/scaley/pkg/basher"
	"github.com/engineyard/scaley/pkg/group"
)

var mygroup *group.Group

func generateGroup() *group.Group {
	mygroup = &group.Group{
		Name: "mygroup",
		ScalingServers: []*group.Server{
			&group.Server{ID: "i-00000001"},
			&group.Server{ID: "i-00000002"},
		},
		ScalingScript: "/bin/decider",
		Strategy:      "legion",
	}

	return mygroup
}

func stubBasher(direction int) {
	basher.Run = func(command string) int {
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
