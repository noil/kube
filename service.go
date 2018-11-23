package main

import (
	"net/http"
	"github.com/noil/kube/version"
	"github.com/noil/kube/state"
	json2 "encoding/json"
)

type Service interface {
	CheckState(w http.ResponseWriter, r *http.Request)
	CheckHealth(w http.ResponseWriter, r *http.Request)
}

type kubeService struct {
	State state.State
}

func (s *kubeService) CheckState(w http.ResponseWriter, r *http.Request) {
	info := s.State.Stats(version.RELEASE, version.REPO, version.COMMIT)
	json, err := json2.Marshal(info)
	if err != nil {
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	log.Info(r.RemoteAddr, info)
}

func (s *kubeService) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
	log.Info(r.RemoteAddr, http.StatusOK)
}

func NewService() Service {
	return &kubeService{State: state.NewState()}
}
