package irc

var (
	// Version release version
	Version = "1.5.2"

	// Build will be overwritten automatically by the build system
	Build = "-dev"

	// GitCommit will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

// FullVersion display the full version and build
func FullVersion() string {
	return Version + Build + " (" + GitCommit + ")"
}
