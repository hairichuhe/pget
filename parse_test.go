package pget

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts_of_ready(t *testing.T) {

	// begin test
	fmt.Fprintf(os.Stdout, "Testing parse_test\n")
	url := "http://example.com/filename.tar.gz"

	args := []string{
		"pget",
		"-p",
		"2",
		url,
		"--trace",
	}

	p := New()
	var opts Options
	if err := p.parseOptions(&opts, args); err != nil {
		t.Errorf("failed to parse command line args: %s", err)
	}

	assert.Equal(t, true, opts.Trace, "failed to parse arguments of trace")
	assert.Equal(t, opts.Procs, 2, "failed to parse arguments of procs")

	if err := p.parseURLs(); err != nil {
		t.Errorf("failed to parse of url: %s", err)
	}

	filename := p.URLFileName(p.TargetDir, url)
	p.SetFileName(filename)
	assert.Equal(t, p.FileName(), "filename.tar.gz", "failed to get of filename from url")

	fmt.Fprintf(os.Stdout, "parse_test Done\n\n")
}

func TestShowhelp(t *testing.T) {
	// begin test
	fmt.Fprintf(os.Stdout, "Testing showhelp_test\n")

	args := []string{
		"pget",
		"-h",
	}

	p := New()
	opts := Options{}
	err := p.parseOptions(&opts, args)
	assert.NotNil(t, err)

	args = []string{
		"pget",
		"--help",
	}

	p = New()
	opts = Options{}
	err = p.parseOptions(&opts, args)
	assert.NotNil(t, err)

	fmt.Fprintf(os.Stdout, "showhelp_test Done\n\n")
}

func TestShowversion(t *testing.T) {
	// begin test
	fmt.Fprintf(os.Stdout, "Testing showversion_test\n")

	args := []string{
		"pget",
		"-v",
	}

	p := New()
	opts := Options{}
	err := p.parseOptions(&opts, args)
	assert.NotNil(t, err)

	args = []string{
		"pget",
		"--version",
	}

	p = New()
	opts = Options{}
	err = p.parseOptions(&opts, args)
	assert.NotNil(t, err)

	fmt.Fprintf(os.Stdout, "showversion_test Done\n\n")
}

func TestShowisupdate(t *testing.T) {
	// begin test
	fmt.Fprintf(os.Stdout, "Testing showversion_test\n")

	args := []string{
		"pget",
		"--check-update",
	}

	p := New()
	opts := Options{}
	err := p.parseOptions(&opts, args)
	assert.NotNil(t, err)

	fmt.Fprintf(os.Stdout, "showversion_test Done\n\n")
}
