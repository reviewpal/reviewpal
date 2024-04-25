// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package engine

import (
	"reflect"
)

const (
	SILENT_MODE  string = "silent"
	VERBOSE_MODE string = "verbose"
)

type PadImport struct {
	Url string `yaml:"url"`
}

func (p PadImport) equals(o PadImport) bool {
	return p.Url == o.Url
}

type PadRule struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Spec        string `yaml:"spec"`
}

func (p PadRule) equals(o PadRule) bool {
	if p.Name != o.Name {
		return false
	}

	if p.Description != o.Description {
		return false
	}

	if p.Spec != o.Spec {
		return false
	}

	return true
}

var kinds = []string{"patch", "author"}

type PadWorkflowRule struct {
	Rule string `yaml:"rule"`
}

func (p PadWorkflowRule) equals(o PadWorkflowRule) bool {
	return p.Rule == o.Rule
}

type PadLabel struct {
	Name        string `yaml:"name"`
	Color       string `yaml:"color"`
	Description string `yaml:"description"`
}

func (p PadLabel) equals(o PadLabel) bool {
	if p.Name != o.Name {
		return false
	}

	if p.Color != o.Color {
		return false
	}

	if p.Description != o.Description {
		return false
	}

	return true
}

type PadWorkflow struct {
	Name                 string                `yaml:"name"`
	Description          string                `yaml:"description"`
	Rules                []PadWorkflowRule     `yaml:"-"`
	Runs                 []PadWorkflowRunBlock `yaml:"-"`
	NonNormalizedRules   any                   `yaml:"if"`
	NonNormalizedActions any                   `yaml:"then"`
	NonNormalizedElse    any                   `yaml:"else"`
	NonNormalizedRun     any                   `yaml:"run"`
}

type PadWorkflowRunBlock struct {
	If      []PadWorkflowRule
	Then    []PadWorkflowRunBlock
	Else    []PadWorkflowRunBlock
	Actions []string
	ForEach *PadWorkflowRunForEachBlock
}

type PadWorkflowRunForEachBlock struct {
	Key   string
	Value string
	In    string
	Do    []PadWorkflowRunBlock
}

func (p PadWorkflow) equals(o PadWorkflow) bool {
	if p.Name != o.Name {
		return false
	}

	if p.Description != o.Description {
		return false
	}

	if len(p.Rules) != len(o.Rules) {
		return false
	}

	for i, pP := range p.Rules {
		oP := o.Rules[i]
		if !pP.equals(oP) {
			return false
		}
	}

	return true
}

type PadGroup struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Kind        string `yaml:"kind"`
	Type        string `yaml:"type"`
	Spec        string `yaml:"spec"`
	Param       string `yaml:"param"`
	Where       string `yaml:"where"`
}

func (p PadGroup) equals(o PadGroup) bool {
	if p.Name != o.Name {
		return false
	}

	if p.Description != o.Description {
		return false
	}

	if p.Kind != o.Kind {
		return false
	}

	if p.Type != o.Type {
		return false
	}

	if p.Spec != o.Spec {
		return false
	}

	if p.Where != o.Where {
		return false
	}

	return true
}

type ReviewpadFile struct {
	Mode           string              `yaml:"mode"`
	IgnoreErrors   *bool               `yaml:"ignore-errors"`
	MetricsOnMerge *bool               `yaml:"metrics-on-merge"`
	Imports        []PadImport         `yaml:"imports"`
	Extends        []string            `yaml:"extends"`
	Groups         []PadGroup          `yaml:"groups"`
	Checks         map[string]PadCheck `yaml:"checks"`
	Rules          []PadRule           `yaml:"rules"`
	Labels         map[string]PadLabel `yaml:"labels"`
	Workflows      []PadWorkflow       `yaml:"workflows"`
	Pipelines      []PadPipeline       `yaml:"pipelines"`
	Recipes        map[string]*bool    `yaml:"recipes"`
	Dictionaries   []PadDictionary     `yaml:"dictionaries"`
}

type PadCheck struct {
	Severity   string                 `yaml:"severity"`
	Activation string                 `yaml:"activation"`
	Parameters map[string]interface{} `yaml:"parameters"`
}

func (p PadCheck) equals(o PadCheck) bool {
	if p.Severity != o.Severity {
		return false
	}

	if len(p.Parameters) != len(o.Parameters) {
		return false
	}

	for key, value := range p.Parameters {
		if o.Parameters[key] != value {
			return false
		}
	}

	return true
}

type PadDictionary struct {
	Name string            `yaml:"name"`
	Spec map[string]string `yaml:"spec"`
}

func (p PadDictionary) equals(o PadDictionary) bool {
	if p.Name != o.Name {
		return false
	}

	if len(p.Spec) != len(o.Spec) {
		return false
	}

	for key, value := range p.Spec {
		if o.Spec[key] != value {
			return false
		}
	}

	return true
}

type PadPipeline struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Trigger     string     `yaml:"trigger"`
	Stages      []PadStage `yaml:"stages"`
}

type PadStage struct {
	Actions              []string `yaml:"-"`
	NonNormalizedActions any      `yaml:"actions"`
	Until                string   `yaml:"until"`
}

func (p PadPipeline) equals(o PadPipeline) bool {
	if p.Name != o.Name {
		return false
	}

	if p.Description != o.Description {
		return false
	}

	if p.Trigger != o.Trigger {
		return false
	}

	for i, pS := range p.Stages {
		oS := o.Stages[i]
		if !pS.equals(oS) {
			return false
		}
	}

	return true
}

func (p PadStage) equals(o PadStage) bool {
	if p.Until != o.Until {
		return false
	}

	for i, pA := range p.Actions {
		oA := o.Actions[i]
		if pA != oA {
			return false
		}
	}

	return true
}

func (r *ReviewpadFile) equals(o *ReviewpadFile) bool {
	if r.Mode != o.Mode {
		return false
	}

	if r.IgnoreErrors != o.IgnoreErrors {
		return false
	}

	if r.MetricsOnMerge != o.MetricsOnMerge {
		return false
	}

	if len(r.Imports) != len(o.Imports) {
		return false
	}
	for i, rI := range r.Imports {
		oI := o.Imports[i]
		if !rI.equals(oI) {
			return false
		}
	}

	if len(r.Extends) != len(o.Extends) {
		return false
	}
	for i, rE := range r.Extends {
		oE := o.Extends[i]
		if rE != oE {
			return false
		}
	}

	if len(r.Rules) != len(o.Rules) {
		return false
	}
	for i, rR := range r.Rules {
		oR := o.Rules[i]
		if !rR.equals(oR) {
			return false
		}
	}

	if len(r.Labels) != len(o.Labels) {
		return false
	}

	for i, rL := range r.Labels {
		oL := o.Labels[i]
		if !rL.equals(oL) {
			return false
		}
	}

	if len(r.Workflows) != len(o.Workflows) {
		return false
	}
	for i, rP := range r.Workflows {
		oP := o.Workflows[i]
		if !rP.equals(oP) {
			return false
		}
	}

	if len(r.Groups) != len(o.Groups) {
		return false
	}
	for i, rG := range r.Groups {
		oG := o.Groups[i]
		if !rG.equals(oG) {
			return false
		}
	}

	if len(r.Dictionaries) != len(o.Dictionaries) {
		return false
	}
	for i, rD := range r.Dictionaries {
		oD := o.Dictionaries[i]
		if !rD.equals(oD) {
			return false
		}
	}

	if len(r.Pipelines) != len(o.Pipelines) {
		return false
	}
	for i, rP := range r.Pipelines {
		oP := o.Pipelines[i]
		if !rP.equals(oP) {
			return false
		}
	}

	if len(r.Checks) != len(o.Checks) {
		return false
	}
	for i, rC := range r.Checks {
		oC := o.Checks[i]
		if !rC.equals(oC) {
			return false
		}
	}

	return reflect.DeepEqual(r.Recipes, o.Recipes)
}

func findGroup(groups []PadGroup, name string) (*PadGroup, bool) {
	for _, group := range groups {
		if group.Name == name {
			return &group, true
		}
	}

	return nil, false
}

func findRule(rules []PadRule, name string) (*PadRule, bool) {
	for _, rule := range rules {
		if rule.Name == name {
			return &rule, true
		}
	}

	return nil, false
}
