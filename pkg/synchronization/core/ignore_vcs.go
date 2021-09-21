package core

// DefaultVCSIgnores is the default set of ignores to use when ignoring VCS
// directories.
var DefaultVCSIgnores = []string{
	".git/",
	".git/**/*",
	".svn/",
	".svn/**/*",
	".hg/",
	".hg/**/*",
	".bzr/",
	".bzr/**/*",
	"_darcs/",
	"_darcs/**/*",
}
