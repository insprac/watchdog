package state

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/insprac/watchdog/config"
	"gopkg.in/yaml.v3"
)

type Timestamp struct {
	Value float64   `yaml:"value"`
	Time  time.Time `yaml:"time"`
}

type State map[string]Timestamp

var (
	mu    sync.Mutex
	state State
)

func init() {
	state = make(State)
	filename := config.GetStateFile()
	if _, err := os.Stat(filename); err == nil {
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(fileContent, &state)
		if err != nil {
			panic(err)
		}
	}
}

func Get(name string) *Timestamp {
	mu.Lock()
	defer mu.Unlock()

	if ts, ok := state[name]; ok {
		return &ts
	}

	return nil
}

func Set(name string, value float64) error {
	mu.Lock()
	defer mu.Unlock()

	state[name] = Timestamp{
		Value: value,
		Time:  time.Now(),
	}

	filename := config.GetStateFile()
	fileContent, err := yaml.Marshal(state)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, fileContent, 0644)
}
