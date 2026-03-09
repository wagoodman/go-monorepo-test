package lib

// version constants that can be tweaked for testing new releases
const (
	Version      = "0.5.0"
	BuildDate    = "2024-01-01"
	CommitHash   = "dev"
	ReleaseNotes = "7 stuff!!"
)

// Hello returns a greeting message with version info.
func Hello() string {
	return "Hello, World! (v" + Version + ")"
}

// VersionInfo returns detailed version information.
func VersionInfo() string {
	return "Version: " + Version + "\nBuild Date: " + BuildDate + "\nCommit: " + CommitHash
}
