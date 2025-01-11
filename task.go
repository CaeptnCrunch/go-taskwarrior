// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// Implementation for taskwarrior's Task entries.

package taskwarrior

// Task representation.
type Task struct {
	Description string  `json:"description"`
	Project     string  `json:"project,omitempty"`
	Status      string  `json:"status,omitempty"`
	Uuid        string  `json:"uuid,omitempty"`
	Urgency     float32 `json:"urgency,omitempty"`
	Priority    string  `json:"priority,omitempty"`
	Due         string  `json:"due,omitempty"`
	End         string  `json:"end,omitempty"`
	Entry       string  `json:"entry,omitempty"`
	Modified    string  `json:"modified,omitempty"`
    Depends   []string  `json:"depends"`
}
