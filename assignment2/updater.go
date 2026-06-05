package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	sconstructPattern = regexp.MustCompile(`point=\d+`)
	versionPattern    = regexp.MustCompile(`ADLMSDK_VERSION_POINT=(\s*)\d+`)
)

// UpdateBuildFiles updates the build version in SConstruct and VERSION.
func UpdateBuildFiles(sourcePath, buildNum string) error {
	if sourcePath == "" {
		return fmt.Errorf("source path is required")
	}
	if buildNum == "" {
		return fmt.Errorf("build number is required")
	}

	root := filepath.Join(sourcePath, "develop", "global", "src")
	if err := replaceInFile(filepath.Join(root, "SConstruct"), sconstructPattern, "point="+buildNum); err != nil {
		return err
	}
	if err := replaceInFile(filepath.Join(root, "VERSION"), versionPattern, "ADLMSDK_VERSION_POINT=${1}"+buildNum); err != nil {
		return err
	}

	return nil
}

func replaceInFile(path string, pattern *regexp.Regexp, replacement string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	input, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	updated := pattern.ReplaceAllString(string(input), replacement)
	if updated == string(input) {
		return fmt.Errorf("no matching version field found in %s", path)
	}

	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, []byte(updated), info.Mode()); err != nil {
		return err
	}

	if err := os.Rename(tmpPath, path); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}

	return nil
}
