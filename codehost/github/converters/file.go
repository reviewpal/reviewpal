package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func CommitFilesFromGithubCommitFiles(fs []*github.CommitFile) []target.CommitFile {
	files := make([]target.CommitFile, len(fs))
	for i, f := range fs {
		files[i] = CommitFileFromGithubCommitFile(f)
	}

	return files
}

func CommitFileFromGithubCommitFile(f *github.CommitFile) target.CommitFile {
	return target.CommitFile{
		Filename: f.GetFilename(),
		Patch:    f.GetPatch(),
	}
}
