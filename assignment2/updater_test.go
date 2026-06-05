package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUpdateBuildFiles(t *testing.T) {
	dir := t.TempDir()
	root := filepath.Join(dir, "develop", "global", "src")
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatal(err)
	}

	sconstructPath := filepath.Join(root, "SConstruct")
	versionPath := filepath.Join(root, "VERSION")

	if err := os.WriteFile(sconstructPath, []byte("point=123,\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(versionPath, []byte("ADLMSDK_VERSION_POINT= 123\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := UpdateBuildFiles(dir, "456"); err != nil {
		t.Fatalf("UpdateBuildFiles returned error: %v", err)
	}

	gotSConstruct, err := os.ReadFile(sconstructPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(gotSConstruct) != "point=456,\n" {
		t.Fatalf("SConstruct = %q, want %q", string(gotSConstruct), "point=456,\n")
	}

	gotVersion, err := os.ReadFile(versionPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(gotVersion) != "ADLMSDK_VERSION_POINT= 456\n" {
		t.Fatalf("VERSION = %q, want %q", string(gotVersion), "ADLMSDK_VERSION_POINT= 456\n")
	}
}
