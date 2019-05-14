package steps

import (
	"fmt"
	//"io/ioutil"

	"github.com/ess/eygo"
	"github.com/ess/kennel"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/v2/pkg/scaley"
	"github.com/engineyard/scaley/v2/pkg/scaley/eycore"
	"github.com/engineyard/scaley/v2/pkg/scaley/fs"
)

type Group struct {
	model *scaley.Group
	api   *eygo.MockDriver
}

func (steps *Group) showReqs() {
	fmt.Println("get requests:", steps.api.Requests("get"))
	fmt.Println("post requests:", steps.api.Requests("post"))
	fmt.Println("put requests:", steps.api.Requests("put"))
	fmt.Println("patch requests:", steps.api.Requests("patch"))
	fmt.Println("delete reuests:", steps.api.Requests("delete"))
}

func (steps *Group) writeGroup() error {
	data, err := yaml.Marshal(steps.model)
	if err != nil {
		return err
	}

	err = fs.Root.MkdirAll("/etc/scaley/groups", 0755)
	if err != nil {
		return fmt.Errorf("Could not create scaley config")
	}

	return afero.WriteFile(
		fs.Root,
		"/etc/scaley/groups/"+steps.model.Name+".yml",
		data,
		0644,
	)

}

func (steps *Group) scriptBase() string {
	return "/bin/"
}

func (steps *Group) stubEnvironment() {
	steps.api.AddResponse(
		"get",
		"environments?id=1",
		steps.constructResponse(`{"environments": [{"id" : 1}]}`, nil),
	)

	//steps.api.AddResponse(
	//"get",
	//"/environments&?id=1&page=2&per_page=100",
	//steps.constructResponse(`{"environments": []}`, nil),
	//)
}

func (steps *Group) stubServer(id int, provisionedId string, state string) {
	steps.api.AddResponse(
		"get",
		fmt.Sprintf("servers?provisioned_id=%s", provisionedId),
		steps.constructResponse(fmt.Sprintf(`{"servers" : [{"id" : %d, "provisioned_id" : "%s", "state" : "%s", "private_hostname": "server%d", "environment": "/1"}]}`, id, provisionedId, state, id), nil),
	)

	//steps.api.AddResponse(
	//"get",
	//fmt.Sprintf("/servers?page=2&per_page=100&provisioned_id=%s", provisionedId),
	//steps.constructResponse(`{"servers" : []}`, nil),
	//)
}

func (steps *Group) unstubServer(id int, provisionedId string) {
	steps.api.RemoveResponse(
		"get",
		fmt.Sprintf("servers?provisioned_id=%s", provisionedId),
	)

	//steps.api.RemoveResponse(
	//"get",
	//fmt.Sprintf("/servers?page=2&per_page=100&provisioned_id=%s", provisionedId),
	//)

	steps.api.AddResponse(
		"get",
		fmt.Sprintf("servers?provisioned_id=%s", provisionedId),
		steps.constructResponse(`{"servers" : []}`, nil),
	)
}

func (steps *Group) constructResponse(data string, err error) eygo.Response {
	return eygo.Response{
		Pages: [][]byte{[]byte(data)},
		Error: err,
	}
}

func (steps *Group) stubStart(id int, provisionedId string, success bool) {
	method := "put"
	path := fmt.Sprintf("servers/%d/start", id)
	response := fmt.Sprintf(`{"request" : {"type" : "start_server", "id" : "%s", "finished_at" : "2018-05-14T00:00:00+00:00", "successful" : %t}}`, provisionedId, success)

	steps.api.RemoveResponse(method, path)
	steps.api.AddResponse(method, path, steps.constructResponse(response, nil))
}

func (steps *Group) stubStop(id int, provisionedId string, success bool) {
	method := "put"
	path := fmt.Sprintf("servers/%d/stop", id)
	response := fmt.Sprintf(`{"request" : {"type" : "stop_server", "id" : "%s", "finished_at" : "2018-05-14T00:00:00+00:00", "successful" : %t}}`, provisionedId, success)

	steps.api.RemoveResponse(method, path)
	steps.api.AddResponse(method, path, steps.constructResponse(response, nil))
}

func (steps *Group) stubChef(success bool) {
	method := "post"
	path := "environments/1/apply"
	response := fmt.Sprintf(`{"request" : {"type" : "configure_environment", "id" : "lolchefrun", "finished_at" : "finished", "successful" : %t}}`, success)

	steps.api.RemoveResponse(method, path)
	steps.api.AddResponse(method, path, steps.constructResponse(response, nil))
}

func (steps *Group) serversStarted() []string {
	found := make([]string, 0)

	for _, id := range []string{"0", "1"} {
		for _, path := range steps.api.Requests("put") {
			if path == "servers/"+id+"/start" {
				found = append(found, id)
			}
		}
	}

	return found
}

func (steps *Group) serversStopped() []string {
	found := make([]string, 0)

	for _, id := range []string{"0", "1"} {
		for _, path := range steps.api.Requests("put") {
			if path == "servers/"+id+"/stop" {
				found = append(found, id)
			}
		}
	}

	return found
}

func (steps *Group) StepUp(s kennel.Suite) {
	s.Step(`^I have a group named mygroup$`, func() error {
		steps.model = generateGroup("")

		return steps.writeGroup()
	})

	s.Step(`^my group is configured to use the legion strategy`, func() error {
		steps.model = generateGroup("legion")

		return steps.writeGroup()
	})

	s.Step(`^my group is configured to ignore stop script errors$`, func() error {
		steps.model.IgnoreStopScriptErrors = true

		return steps.writeGroup()
	})

	s.Step(`^my group is configured to unlock on failures$`, func() error {
		steps.model.UnlockOnFailure = true

		return steps.writeGroup()
	})

	s.Step(`^my group is configured to use the individual strategy$`, func() error {
		steps.model = generateGroup("individual")

		return steps.writeGroup()
	})

	s.Step(`^I have a script that determines if I should scale up or down$`, func() error {

		// Create a dummy scripts location in the fake FS
		err := fs.Root.MkdirAll(steps.scriptBase(), 0755)
		if err != nil {
			return err
		}

		err = afero.WriteFile(fs.Root, steps.scriptBase()+"decider", []byte(""), 0755)

		if err != nil {
			return fmt.Errorf("Could not write scaling script")
		}

		return nil
	})

	s.Step(`^the scaling script for my group does not exist$`, func() error {
		return fs.Root.Remove(steps.scriptBase() + "decider")
	})

	s.Step(`^there is capacity for the group to upscale$`, func() error {
		steps.stubEnvironment()

		for i, server := range steps.model.ScalingServers {
			steps.stubServer(i, server, "stopped")
			steps.stubStart(i, server, true)
		}

		steps.stubChef(true)

		return nil
	})

	s.Step(`^the API is erroring on server start requests$`, func() error {
		for i := range steps.model.ScalingServers {
			method := "put"
			path := fmt.Sprintf("servers/%d/start", i)
			steps.api.RemoveResponse(method, path)
		}

		return nil
	})

	s.Step(`^the API is erroring on server stop requests$`, func() error {
		for i := range steps.model.ScalingServers {
			method := "put"
			path := fmt.Sprintf("servers/%d/stop", i)
			steps.api.RemoveResponse(method, path)
		}

		return nil
	})

	s.Step(`^the servers cannot be started successfully$`, func() error {
		for i, server := range steps.model.ScalingServers {
			steps.stubStart(i, server, false)
		}

		return nil
	})

	s.Step(`^the servers cannot be stopped successfully$`, func() error {
		for i, server := range steps.model.ScalingServers {
			steps.stubStop(i, server, false)
		}

		return nil
	})

	s.Step(`^the API is erroring on environment configuration requests$`, func() error {
		steps.api.RemoveResponse("post", "environments/1/apply")

		return nil
	})

	s.Step(`^the environment cannot run chef successfully$`, func() error {
		steps.stubChef(false)

		return nil
	})

	s.Step(`^there is capacity for the group to downscale$`, func() error {
		steps.stubEnvironment()

		for i, server := range steps.model.ScalingServers {
			steps.stubServer(i, server, "running")
			steps.stubStop(i, server, true)
		}

		steps.stubChef(true)

		return nil
	})

	s.Step(`^the group is scaled up$`, func() error {
		if len(steps.serversStarted()) != 2 {
			return fmt.Errorf("At least one server was not scaled up")
		}

		return nil
	})

	s.Step(`^all of the servers in the group are started$`, func() error {
		if len(steps.serversStarted()) != len(steps.model.ScalingServers) {
			return fmt.Errorf("At least one server was not started")
		}

		return nil
	})

	s.Step(`^only one server in the group is started$`, func() error {
		if len(steps.serversStarted()) > 1 {
			return fmt.Errorf("More than one server was started")
		}

		return nil
	})

	s.Step(`^the group is scaled down$`, func() error {
		if len(steps.serversStopped()) != 2 {
			return fmt.Errorf("At least one server was not scaled down")
		}

		return nil
	})

	s.Step(`^all of the servers in the group are stopped$`, func() error {
		if len(steps.serversStopped()) != len(steps.model.ScalingServers) {
			return fmt.Errorf("At least one server was not stopped")
		}

		return nil
	})

	s.Step(`^only one server in the group is stopped$`, func() error {
		if len(steps.serversStopped()) > 1 {
			return fmt.Errorf("More than one server was stopped")
		}

		return nil
	})

	s.Step(`^all applicable servers but the first server are stopped$`, func() error {
		found := steps.serversStopped()

		for _, id := range found {
			if id == "0" {
				return fmt.Errorf("Expected the first server not to be stopped")
			}
		}

		if len(found) != len(steps.model.ScalingServers)-1 {
			return fmt.Errorf("Expected all servers but the first to be stopped")
		}

		return nil
	})

	s.Step(`^there is not capacity for the group to upscale$`, func() error {
		steps.stubEnvironment()

		for i, server := range steps.model.ScalingServers {
			steps.stubServer(i, server, "running")
		}

		return nil
	})

	s.Step(`^there is not capacity for the group to downscale$`, func() error {
		steps.stubEnvironment()

		for i, server := range steps.model.ScalingServers {
			steps.stubServer(i, server, "stopped")
		}

		return nil
	})

	s.Step(`^no changes are made$`, func() error {
		found := 0

		for _, id := range []string{"0", "1"} {
			for _, path := range steps.api.Requests("put") {
				if path == "servers/"+id+"/start" {
					found += 1
				}
			}
		}

		if found != 0 {
			return fmt.Errorf("No servers should have been started")
		}

		return nil

	})

	s.Step(`^my group does not use a custom stop script$`, func() error {
		steps.model.StopScript = ""

		return steps.writeGroup()
	})

	s.Step(`^my group uses a custom stop script that always succeeds$`, func() error {
		steps.model.StopScript = "stop_script_good"

		return steps.writeGroup()
	})

	s.Step(`^my group uses a custom stop script that fails for the first server$`, func() error {
		steps.model.StopScript = "stop_script_bad"

		return steps.writeGroup()
	})

	s.Step(`^my group lacks a scaling script$`, func() error {
		steps.model.ScalingScript = ""

		return steps.writeGroup()
	})

	s.Step(`^my group has no scaling servers$`, func() error {
		steps.model.ScalingServers = make([]string, 0)

		return steps.writeGroup()
	})

	s.Step(`^my group has a scaling server that doesn't exist$`, func() error {
		for i, server := range steps.model.ScalingServers {
			steps.unstubServer(i, server)
		}

		return nil
	})

	s.BeforeScenario(func(interface{}) {
		stubBasher(0)
		steps.api = eygo.NewMockDriver()
		steps.api.AddResponse("get", "/users/current", steps.constructResponse(`{"user": {}}`, nil))

		eycore.Driver = steps.api
	})
}

func init() {
	kennel.Register(new(Group))
}
