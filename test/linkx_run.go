// +build !nacl
// run

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Run the linkx test.

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	test(" ") // old deprecated syntax
	test("=") // new syntax
}

func test(sep string) {
	// Successful run
	cmd := exec.Command("go", "run", "-ldflags=-X main.tbd"+sep+"hello -X main.overwrite"+sep+"trumped -X main.nosuchsymbol"+sep+"neverseen", "linkx.go")
	var out, errbuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errbuf
	err := cmd.Run()
	if err != nil {
		fmt.Println(errbuf.String())
		fmt.Println(out.String())
		fmt.Println(err)
		os.Exit(1)
	}

	want := "hello\ntrumped\n"
	got := out.String()
	if got != want {
		fmt.Printf("got %q want %q\n", got, want)
		os.Exit(1)
	}

	// Issue 8810
	cmd = exec.Command("go", "run", "-ldflags=-X main.tbd", "linkx.go")
	_, err = cmd.CombinedOutput()
	if err == nil {
		fmt.Println("-X linker flag should not accept keys without values")
		os.Exit(1)
	}
}
