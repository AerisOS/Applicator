package Utilities

import (
	"fmt"
	"github.com/AerisHQ/Applicator/Source/Types"
	"github.com/sqweek/dialog"
	"os"
	"strings"
)

func PermissionPopup(ApplicationName string, Permission Types.Permission, SpecificData string) bool {
	var Message string

	switch Permission.Permission {
	case "FILE_ACCESS":
		Message = "'" + ApplicationName + "' is requesting access to the following folder:\n\n" + SpecificData + "\n\nDo you want to allow this application to access this folder?"
	case "SYSTEM_PROCESSES":
		Message = "'" + ApplicationName + "' is requesting access to system processes.\n\nThis will allow the application to interact with system processes, which may include reading process information or sending data to processes.\n\nDo you want to allow this application to access system processes?"
	}

	if os.Getenv("DISPLAY") != "" || os.Getenv("WAYLAND_DISPLAY") != "" {
		return dialog.Message(Message).
			Title("An Application is requesting a permission").
			YesNo()
	}

	/* Fallback to terminal input if there is no display is available */
	var UserResponse string

	fmt.Println(Message)
	fmt.Print("Type 'yes' to allow or 'no' to deny (Default: no): ")
	_, err := fmt.Scanln(&UserResponse)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}

	UserResponse = strings.ToLower(strings.TrimSpace(UserResponse))

	if UserResponse == "yes" || UserResponse == "y" {
		fmt.Println("================")
		return true
	} else {
		fmt.Println("================")
		return false
	}
}
