package main

import (
	"testing"
	"os/exec"
)

func TestMain(t *testing.T) {
	cmd := exec.Command("go" , "run" , "main.go", "Hello There!", "shadow" )
	outputByte, err:= cmd.Output()
	if err != nil {
		t.Errorf("Failed to run program: %v" , err)
	}
	output := string(outputByte)

	expected := 
`                                                                                         
_|    _|          _| _|                _|_|_|_|_| _|                                  _| 
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| 
                                                                                         
                                                                                         
`
	
	if output != expected {
		t.Errorf("Expected %q, but got %q" , expected , output)
	}

}