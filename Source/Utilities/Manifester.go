package Utilities

import (
	"github.com/AerisHQ/Applicator/Source/Types"
	"github.com/goccy/go-yaml"
	"os"
	"strings"
)

func ParseManifest(manifestFile string) (*Types.Manifest, error) {
	var manifest Types.Manifest

	/* Read the manifest file */
	Content, FileReadError := os.ReadFile(manifestFile)

	if FileReadError != nil {
		return nil, FileReadError
	}

	/* Unmarshal the content into a Manifest struct */
	YAMLError := yaml.Unmarshal(Content, &manifest)

	if YAMLError != nil {
		return nil, YAMLError
	}

	return &manifest, nil
}

func ManifestExists(ApplicationFolder string) (bool, string) {
	DirectoryList, DirectoryReadError := os.ReadDir(ApplicationFolder)

	if DirectoryReadError != nil {
		return false, ""
	}

	for _, file := range DirectoryList {
		Name := strings.ToLower(file.Name())

		if Name == "manifest.yaml" || Name == "manifest.yml" {
			return true, ApplicationFolder + "/" + file.Name()
		}
	}

	return false, ""
}
