package Utilities

import (
	"github.com/AerisHQ/Applicator/Source/Types"
	"github.com/sqweek/dialog"
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

func HandlePermissions(Manifest Types.Manifest) (bool, []string, []string) {
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

						if AskForPermission(Manifest.Application.Name, permission, pathStr) {
							CLIArguments = append(CLIArguments, "--bind", pathStr, pathStr)
							PermissionsGranted = append(PermissionsGranted, "FILE_ACCESS:"+pathStr)
						}
					}
				}
			} else {
				return false, []string{}, []string{}
			}
		case "SYSTEM_PROCESSES":
			if AskForPermission(Manifest.Application.Name, permission, "System processes") {
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

func AskForPermission(ApplicationName string, Permission Types.Permission, SpecificData string) bool {
	var Message string

	switch Permission.Permission {
	case "FILE_ACCESS":
		Message = "'" + ApplicationName + "' is requesting access to the following folder:\n\n" + SpecificData + "\n\nDo you want to allow this application to access this folder?"
	case "SYSTEM_PROCESSES":
		Message = "'" + ApplicationName + "' is requesting access to system processes.\n\nThis will allow the application to interact with system processes, which may include reading process information or sending data to processes.\n\nDo you want to allow this application to access system processes?"
	}

	return dialog.Message(Message).
		Title("An Application is requesting a permission").
		YesNo()
}
