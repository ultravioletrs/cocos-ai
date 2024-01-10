// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package notifications

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errFailedToCreateNotification = errors.New("failed to create notification on server")

type service struct {
	service   string
	serverUrl string
}

type Service interface {
	SendNotification(event, computationId, status string, details json.RawMessage) error
}

func New(svc, serverUrl string) Service {
	return &service{
		service:   svc,
		serverUrl: serverUrl,
	}
}

func (s *service) SendNotification(event, computationId, status string, details json.RawMessage) error {
	body := struct {
		EventType     string          `json:"event_type"`
		Timestamp     time.Time       `json:"timestamp"`
		ComputationID string          `json:"computation_id,omitempty"`
		Details       json.RawMessage `json:"details,omitempty"`
		Originator    string          `json:"originator"`
		Status        string          `json:"status,omitempty"`
	}{
		EventType:     event,
		Timestamp:     time.Now(),
		ComputationID: computationId,
		Originator:    s.service,
		Status:        status,
		Details:       details,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/computations/events", s.serverUrl), bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errFailedToCreateNotification
	}
	return nil
}
