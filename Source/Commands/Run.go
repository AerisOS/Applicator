package Commands

import (
	"context"
	"fmt"
	"github.com/AerisHQ/Applicator/Source/Utilities"
	"github.com/alexellis/go-execute/v2"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	// Global Variables
	verbose       bool
	appPath       string
	ignorePrompts bool

	// Local Variables
	ManifestPath             string
	ManifestFound            bool
	PermissionsGrantedString string
)

var Run = &cobra.Command{
	Use:     "run",
	Aliases: []string{"execute", "start"},
	Short:   "Run the specified application in a sandboxed environment",
	Long:    "Run the specified application in a sandboxed environment with limited permissions and resources. You may specify if you want to give the application access to certain system resources or file paths.",
	Run: func(cmd *cobra.Command, args []string) {
		Logger := log.NewWithOptions(os.Stdout, log.Options{
			Level:        log.InfoLevel,
			TimeFormat:   "2006-01-02 15:04:05",
			Prefix:       "Sandboxer",
			ReportCaller: true,
		})

		/* Variable Checks */
		if verbose {
			Logger.SetLevel(log.DebugLevel)
			Logger.Debug("Verbose logging enabled")
		}

		if ignorePrompts {
			Logger.Warn("Ignoring permission prompts, all permissions will be granted automatically")
		}

		/* Validating the application path */
		IsApplication, ApplicationPath := Utilities.IsApplication(appPath)

		if !IsApplication {
			Logger.Fatal("The specified path is not a valid application", "path", appPath)
		}

		/* Reading the application directory */
		DirectoryOutput, DirectoryReadError := os.ReadDir(ApplicationPath)

		Logger.Debugf("Reading application directory: %s", ApplicationPath)
		Logger.Debugf("Directory contents: %v", DirectoryOutput)

		if DirectoryReadError != nil {
			Logger.Fatal("Failed to read application directory", "error", DirectoryReadError)
		}

		/* Validating that the manifest file exists */
		if ManifestFound, ManifestPath = Utilities.ManifestExists(ApplicationPath); !ManifestFound {
			Logger.Fatal("Manifest file not found in the specified application directory", "path", ApplicationPath)
		}

		/* Parsing the manifest file */
		Logger.Debug("Reading manifest file", "file", ManifestPath)
		Manifest, ManifestParseError := Utilities.ParseManifest(ManifestPath)

		if ManifestParseError != nil {
			Logger.Fatal("Failed to parse manifest file", "file", ManifestPath, "error", ManifestParseError)
		}

		Logger.Debugf("Manifest file contents: %s", Manifest)
		Logger.Debug("Parsing permissions from manifest", "permissions", Manifest.Permissions)
		DidUserGrantPermissions, CLIArguments, PermissionsGranted := Utilities.HandlePermissions(*Manifest, ignorePrompts)

		fmt.Println("Permissions Granted:", PermissionsGranted)
		fmt.Println(Manifest.Permissions)

		if !DidUserGrantPermissions {
			Logger.Debug("User did not grant the required permissions for the application to run", "application", Manifest.Application.Name)
		} else {
			Logger.Debug("User granted permissions", "permissions", PermissionsGranted)
		}

		if len(PermissionsGranted) > 0 {
			/* We can't use ',' as a separator here because bubblewrap will error out */
			PermissionsGrantedString = strings.Join(PermissionsGranted, "%")
		} else {
			PermissionsGrantedString = "..."
		}

		Args := []string{
			"--setenv", "USER", os.Getenv("USER"),
			"--setenv", "HOME", os.Getenv("HOME"),
			"--setenv", "PERMISSIONS_GRANTED", PermissionsGrantedString,
			"--bind", ApplicationPath + "/Data", "/Data",
			"--ro-bind", ApplicationPath, "/Application",
			"--ro-bind /usr /usr",
			"--ro-bind /bin /bin",
			"--ro-bind", "/usr/lib64 /lib64",
			"--dev /dev",
			"--unshare-pid",
			"--new-session",
			"--",
			"/Application/" + Manifest.Application.Entrypoint,
		}

		Args = append(CLIArguments, Args...)

		Command := execute.ExecTask{
			Command: "bwrap",
			Args:    Args,
			Shell:   true,
		}

		Output, _ := Command.Execute(context.Background())
		fmt.Printf("stdout: %s\nstderr: %s\nexit-code: %d\n", Output.Stdout, Output.Stderr, Output.ExitCode)
	},
}

func init() {
	Run.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	Run.Flags().StringVarP(&appPath, "app", "a", appPath, "Path to the application to run")
	Run.Flags().BoolVarP(&ignorePrompts, "ignoreprompts", "i", false, "Ignore permission prompts and automatically grant all permission (default: false)")
	err := Run.MarkFlagRequired("app")

	if err != nil {
		return
	}
}
