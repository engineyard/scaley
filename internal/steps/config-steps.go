package steps

import (
	"fmt"

	"github.com/ess/kennel"
	"github.com/spf13/afero"
	"github.com/spf13/viper"

	"github.com/engineyard/scaley/v2/pkg/scaley/fs"
)

type Config struct{}

func (steps *Config) writeConfig() error {
	config := `---
token: supersekrat
reporting_url: "https://example.com/reporting/1234"`

	err := fs.Root.MkdirAll("/etc/scaley", 0755)
	if err != nil {
		return fmt.Errorf("Could not create scaley config")
	}

	err = afero.WriteFile(
		fs.Root,
		"/etc/scaley/config.yml",
		[]byte(config),
		0644,
	)

	//data, _ := afero.ReadFile(
	//common.Root,
	//)

	return err
}

func (steps *Config) StepUp(s kennel.Suite) {
	s.Step(`^I have a scaley config$`, func() error {
		return steps.writeConfig()
	})

	s.BeforeScenario(func(interface{}) {
		fs.Root = afero.NewMemMapFs()
		viper.SetFs(fs.Root)
	})
}

func init() {
	kennel.Register(new(Config))
}
