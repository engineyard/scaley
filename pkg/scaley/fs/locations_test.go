package fs

import (
	"testing"

	"github.com/ess/testscope"
)

type paths struct {
	created []string
}

func includesPath(paths []string, query string) bool {
	for _, path := range paths {
		if path == query {
			return true
		}
	}

	return false
}

func (p *paths) includes(path string) bool {
	return includesPath(p.created, path)
}

func (p *paths) create(path string) {
	p.created = append(p.created, path)
}

func mockingCreateDir(doit func(*paths)) {
	createdDirs := make([]string, 0)
	p := &paths{createdDirs}
	origCreateDir := CreateDir
	CreateDir = func(path string) {
		p.create(path)
	}

	defer func() { CreateDir = origCreateDir }()

	doit(p)
}

func TestGroupConfigs(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	t.Run("is the group config path", func(t *testing.T) {
		mockingCreateDir(func(garbage *paths) {
			result := GroupConfigs()

			if result != "/etc/scaley/groups" {
				t.Errorf("Expected to receive the group config path")
			}
		})
	})

	t.Run("creates the overall config path", func(t *testing.T) {
		mockingCreateDir(func(created *paths) {
			GroupConfigs()

			if !created.includes(configDir()) {
				t.Errorf("Expected the config dir to be created")
			}
		})
	})

	t.Run("creates the groups path", func(t *testing.T) {
		mockingCreateDir(func(created *paths) {
			result := GroupConfigs()

			if !created.includes(result) {
				t.Errorf("Expected the group configs path to be created")
			}
		})
	})
}

func TestLocks(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	t.Run("is the locks path", func(t *testing.T) {
		mockingCreateDir(func(garbage *paths) {
			result := Locks()

			if result != "/var/run/scaley/lock" {
				t.Errorf("Expected to receive the lock state path")
			}
		})
	})

	t.Run("creates the overall state path", func(t *testing.T) {
		mockingCreateDir(func(created *paths) {
			Locks()

			if !created.includes(stateDir()) {
				t.Errorf("Expected the state dir to be created")
			}
		})
	})

	t.Run("creates the lock state path", func(t *testing.T) {
		mockingCreateDir(func(created *paths) {
			result := Locks()

			if !created.includes(result) {
				t.Errorf("Expected the lock state path to be created")
			}
		})
	})
}

func TestDataDir(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	t.Run("is the data path", func(t *testing.T) {
		mockingCreateDir(func(garbage *paths) {
			result := DataDir()

			if result != "/var/lib/scaley" {
				t.Errorf("Expected to receive the data path, got %s", result)
			}
		})
	})

	t.Run("creates the data path", func(t *testing.T) {
		mockingCreateDir(func(created *paths) {
			result := DataDir()

			if !created.includes(result) {
				t.Errorf("Expected the data to be created")
			}
		})
	})
}

func TestFileExists(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	t.Run("when the path exists", func(t *testing.T) {
		result := FileExists("locations_test.go")
		if !result {
			t.Errorf("Expected an affirmative response")
		}
	})

	t.Run("when the file does not exist", func(t *testing.T) {
		result := FileExists("thisisatotallynonexistentfile")
		if result {
			t.Errorf("Expected a negative response")
		}
	})
}
