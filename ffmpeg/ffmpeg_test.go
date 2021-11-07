package ffmpeg

import (
	"fmt"
	"testing"
)

func TestFightingErrors(t *testing.T) {
	cfg, err := NewAutoConfig(Flags{Progress: true})
	if err != nil {
		t.Error(err)
	}
	// Create instance and fill
	instance := New(cfg).Input("test").Output("test").SkipMetadata()
	// Start
	p, err := instance.Start(Options{"-y"})
	if err != nil {
		t.Error(err)
	}
	if p == nil {
		t.Error("waiting progress chain!")
	}
	for tick := range p {
		fmt.Println(tick)
	}
	if err = instance.Error(); err == nil {
		t.Error("waiting for error!")
	}
}
