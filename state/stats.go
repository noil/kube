package state

import (
	"os"
	"runtime"
	"fmt"
)

type State interface {
	Stats(version, repo, commit string) ServiceState
}

type ServiceState struct {
	Host    string       `json:"host"`
	Runtime *RuntimeInfo `json:"runtime"`
	Version string       `json:"version"`
	Repo    string       `json:"repo"`
	Commit  string       `json:"commit"`
}

// RuntimeInfo defines runtime part of service information
type RuntimeInfo struct {
	Compiler   string `json:"compilier"`
	CPU        int    `json:"cpu"`
	Memory     string `json:"memory"`
	Goroutines int    `json:"goroutines"`
}

type kubeState struct {
}

func NewState() State {
	return &kubeState{}
}

func (s *kubeState) Stats(version, repo, commit string) ServiceState {
	host, _ := os.Hostname()
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)

	rt := &RuntimeInfo{
		CPU:        runtime.NumCPU(),
		Memory:     fmt.Sprintf("%.2fMB", float64(m.Alloc)/(1<<(10*2))),
		Goroutines: runtime.NumGoroutine(),
	}

	return ServiceState{
		Host:    host,
		Runtime: rt,
		Version: version,
		Repo:    repo,
		Commit:  commit,
	}
}

func (s ServiceState) String() string {
	return fmt.Sprintf(` host: %s, compilier: %s, cpu: %v, memory: %s, goroutines: %v, version: %s, repo: %s, commit: %s`, s.Host, s.Runtime.Compiler, s.Runtime.CPU, s.Runtime.Memory, s.Runtime.Goroutines, s.Version, s.Repo, s.Commit)
}
