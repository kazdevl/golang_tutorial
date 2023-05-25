package main

import "testing"

func Test_CreateFile(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		content string
	}{
		{
			name:    "sample",
			file:    "sample.txt",
			content: "sample",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFile(tt.file, tt.content); err != nil {
				t.Errorf("CreateFile() error = %v", err)
			}
		})
	}
}
