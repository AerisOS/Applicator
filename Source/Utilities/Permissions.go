package Utilities

import (
	"github.com/AerisHQ/Applicator/Source/Types"
	"os"
	"slices"
	"strings"
)

var BlacklistedPaths = []string{
	"/",
	"/sys",
	"/bin",
	"/boot",
	"/dev",
	"/etc",
	"/home",
	"/lib",
	"/lib64",
	"/media",
	"/mnt",
	"/opt",
	"/proc",
	"/usr",
}

/* The PermissionPopup function is located in Source/Utilities/PermissionPopup.go */

func HandlePermissions(Manifest Types.Manifest, IgnorePrompts bool) (bool, []string, []string) {
	if Manifest.Permissions == nil || len(Manifest.Permissions) == 0 {
		return true, []string{}, []string{}
	}

	var CLIArguments []string
	var PermissionsGranted []string

	for _, permission := range Manifest.Permissions {
		switch permission.Permission {
		case "FILE_ACCESS":
			Properties := permission.Data.(map[string]interface{})

			if Properties["paths"] != nil {
				paths := Properties["paths"].([]interface{})

				for _, path := range paths {
					if !slices.Contains(BlacklistedPaths, path.(string)) {
						pathStr := strings.Replace(path.(string), "~", os.Getenv("HOME"), -1)

						if PermissionPopup(Manifest.Application.Name, permission, pathStr, IgnorePrompts) {
							CLIArguments = append(CLIArguments, "--bind", pathStr, pathStr)
							PermissionsGranted = append(PermissionsGranted, "FILE_ACCESS:"+pathStr)
						}
					}
				}
			} else {
				return false, []string{}, []string{}
			}
		case "SYSTEM_PROCESSES":
			if PermissionPopup(Manifest.Application.Name, permission, "System processes", IgnorePrompts) {
				CLIArguments = append(CLIArguments, "--proc", "/proc")
				PermissionsGranted = append(PermissionsGranted, "SYSTEM_PROCESSES")
			}
		default:
			return false, []string{}, []string{}
		}
	}

	if len(PermissionsGranted) > 0 {
		return true, CLIArguments, PermissionsGranted
	} else {
		return false, []string{}, []string{}
	}
}
