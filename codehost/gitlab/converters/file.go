package converters

import (
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func CommitFilesFromGitlabMergeRequestDiffs(fs []*gitlab.MergeRequestDiff) []target.CommitFile {
	files := make([]target.CommitFile, len(fs))
	for i, f := range fs {
		files[i] = CommitFileFromGitlabMergeRequestDiff(f)
	}

	return files
}

func CommitFileFromGitlabMergeRequestDiff(f *gitlab.MergeRequestDiff) target.CommitFile {
	return target.CommitFile{
		Filename: f.NewPath,
		Patch:    f.Diff,
	}
}
