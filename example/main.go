package main

import (
	"fmt"
	"os"

	"github.com/skaji/go-github-release"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	release := &github.Release{
		Owner:      "skaji",
		Repository: "relocatable-perl",
	}
	tag, err := release.GetLatestTag()
	if err != nil {
		return err
	}
	fmt.Println("latest tag is:", tag)

	assets, err := release.GetLatestAssets()
	if err != nil {
		return err
	}
	fmt.Println("latest assets are:")
	for _, a := range assets {
		fmt.Println("*", a)
	}
	return nil
}
