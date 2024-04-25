// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file

package target

import (
	"fmt"
	"regexp"
)

type Patch map[string]*File

type CommitFile struct {
	Filename string `json:"filename,omitempty"`
	Patch    string `json:"patch,omitempty"`
}

type File struct {
	Repr *CommitFile
	Diff []*diffBlock
}

func (f *File) AppendToDiff(
	isContext bool,
	oldStart, oldEnd, newStart, newEnd int64,
	oldLine, newLine string,
) {
	f.Diff = append(f.Diff, &diffBlock{
		IsContext: isContext,
		Old: &diffSpan{
			oldStart,
			oldEnd,
		},
		New: &diffSpan{
			newStart,
			newEnd,
		},
		oldLine: oldLine,
		newLine: newLine,
	})
}

func NewFile(file *CommitFile) (*File, error) {
	diffBlocks, err := parseFilePatch(file.Patch)
	if err != nil {
		return nil, fmt.Errorf("error in file patch %s: %v", file.Filename, err)
	}

	return &File{
		Repr: file,
		Diff: diffBlocks,
	}, nil
}

func (f *File) Query(expr string) (bool, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return false, fmt.Errorf("query: compile error %v", err)
	}

	for _, block := range f.Diff {
		if !block.IsContext {
			if r.Match([]byte(block.newLine)) {
				return true, nil
			}
		}
	}
	return false, nil
}
