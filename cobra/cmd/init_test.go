package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestInitCmd(t *testing.T) {
	viper.Set("license", "apache")
	viper.Set("year", 2017)

	defer os.RemoveAll("testproject")
	rootCmd.SetArgs([]string{"init", "github.com/spf13/testproject"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Error during execution: %v", err)
	}

	tests := []struct {
		name string
		expected string
	}{
		{"testproject/cmd/root.go", "8b0f6cbb314ae60c44e6155323f76ea3"},
		{"testproject/main.go", "cd6be5e8173223447b0b6e7d0e351f19"},
	}

	for _, v := range tests {
		// get md5sum from generated file
		file, err := os.Open(v.name)
		if err != nil {
			t.Fatalf("could not open file for reading: %v", err)
		}

		hash := md5.New()

		if _, err := io.Copy(hash, file); err != nil {
			t.Fatalf("could not copy file into hash interface: %v", err)
		}

		hashBytes := hash.Sum(nil)
		sum := hex.EncodeToString(hashBytes)

		if sum != v.expected {
			t.Errorf("md5 sums not equal\nhave: %s got: %s\n", v.expected, sum)
		}

		file.Close()
	}

}

// TestGoldenInitCmd initializes the project "github.com/spf13/testproject"
// in GOPATH and compares the content of files in initialized project with
// appropriate golden files ("testdata/*.golden").
// Use -update to update existing golden files.
/*
func TestGoldenInitCmd(t *testing.T) {
	projectName := "github.com/spf13/testproject"
	project := NewProject(projectName)
	defer os.RemoveAll(project.AbsPath())

	viper.Set("author", "NAME HERE <EMAIL ADDRESS>")
	viper.Set("license", "apache")
	viper.Set("year", 2017)
	defer viper.Set("author", nil)
	defer viper.Set("license", nil)
	defer viper.Set("year", nil)

	os.Args = []string{"cobra", "init", projectName}
	if err := rootCmd.Execute(); err != nil {
		t.Fatal("Error by execution:", err)
	}

	expectedFiles := []string{".", "cmd", "LICENSE", "main.go", "cmd/root.go"}
	gotFiles := []string{}

	// Check project file hierarchy and compare the content of every single file
	// with appropriate golden file.
	err := filepath.Walk(project.AbsPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Make path relative to project.AbsPath().
		// E.g. path = "/home/user/go/src/github.com/spf13/testproject/cmd/root.go"
		// then it returns just "cmd/root.go".
		relPath, err := filepath.Rel(project.AbsPath(), path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)
		gotFiles = append(gotFiles, relPath)
		goldenPath := filepath.Join("testdata", filepath.Base(path)+".golden")

		switch relPath {
		// Known directories.
		case ".", "cmd":
			return nil
		// Known files.
		case "LICENSE", "main.go", "cmd/root.go":
			if *update {
				got, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				if err := ioutil.WriteFile(goldenPath, got, 0644); err != nil {
					t.Fatal("Error while updating file:", err)
				}
			}
			return compareFiles(path, goldenPath)
		}
		// Unknown file.
		return errors.New("unknown file: " + path)
	})
	if err != nil {
		t.Fatal(err)
	}

	// Check if some files lack.
	if err := checkLackFiles(expectedFiles, gotFiles); err != nil {
		t.Fatal(err)
	}
}
*/
