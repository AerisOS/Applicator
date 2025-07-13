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

func HandlePermissions(Manifest Types.Manifest) (bool, []string) {
	if Manifest.Permissions == nil || len(Manifest.Permissions) == 0 {
		return true, []string{}
	}

	for _, permission := range Manifest.Permissions {
		switch permission.Permission {
		case "FILE_ACCESS":
			Properties := permission.Data.(map[string]interface{})

			if Properties["paths"] != nil {
				paths := Properties["paths"].([]interface{})
				var CustomPermissions []string

				for _, path := range paths {
					if !slices.Contains(BlacklistedPaths, path.(string)) {
						pathStr := strings.Replace(path.(string), "~", os.Getenv("HOME"), -1)

						if AskForPermission(Manifest.Application.Name, permission, pathStr) {
							CustomPermissions = append(CustomPermissions, "--bind", pathStr, pathStr)
						}
					}
				}

				return true, CustomPermissions
			} else {
				return false, []string{"No paths specified for FILE_ACCESS permission"}
			}
		default:
			return false, []string{"Unknown permission requested: " + permission.Permission}
		}

	}

	return false, []string{}
}

func AskForPermission(ApplicationName string, Permission Types.Permission, SpecificData string) bool {
	var Message string

	switch Permission.Permission {
	case "FILE_ACCESS":
		Message = "'" + ApplicationName + "' is requesting access to the following folder:\n\n" + SpecificData + "\n\nDo you want to allow this application to access this folder?"
	}

	return dialog.Message(Message).
		Title("An Application is requesting a permission").
		YesNo()
}
