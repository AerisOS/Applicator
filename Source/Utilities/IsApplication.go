package Utilities

import (
	"os"
)

func IsApplication(ID string) (bool, string) {
	if ID == "" {
		return false, ""
	}

	IsLocalApplication, LocalError := os.Stat("./" + ID)
	IsUserApplication, UserError := os.Stat(os.ExpandEnv("$HOME/Applications/" + ID))
	IsSystemApplication, SystemError := os.Stat("/Applications/" + ID)

	if LocalError == nil && IsLocalApplication.IsDir() {
		return true, "./" + ID
	}

	if UserError == nil && IsUserApplication.IsDir() {
		return true, os.ExpandEnv("$HOME/Applications/" + ID)
	}

	if SystemError == nil && IsSystemApplication.IsDir() {
		return true, "/Applications/" + ID
	}

	return false, ""
}
