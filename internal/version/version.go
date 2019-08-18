package version

import "fmt"

var (
	// Version is the build version eg: v1.0.0
	Version string
	// MinVersion is the latest commit hash on building
	MinVersion string
	// BuildTime is the unix timestamp when the build was done
	BuildTime string
)

// PrintVersion will print the version details
func PrintVersion() {
	fmt.Println("Version :", Version)
	fmt.Println("MinVersion :", MinVersion)
	fmt.Println("BuildTime :", BuildTime)
}
