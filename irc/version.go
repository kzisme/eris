package irc

var (
	//PackageName package name
	Package = "eris"

	// Version release version
	Version = "1.6.0"

	// Build will be overwritten automatically by the build system
	Build = "-dev"

	// GitCommit will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

// FullVersion display the full version and build
func FullVersion() string {
	return Package + " v" + Version + Build + " (" + GitCommit + ")"
}
