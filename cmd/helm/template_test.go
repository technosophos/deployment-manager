/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"testing"
)

var chartPath = "./../../pkg/chartutil/testdata/subpop/charts/subchart1"

func TestTemplateCmd(t *testing.T) {
	tests := []cmdTestCase{
		{
			name:   "check name",
			cmd:    "template " + chartPath,
			golden: "output/template.txt",
		},
		{
			name:   "check set name", //
			cmd:    "template --set service.name=apache " + chartPath,
			golden: "output/template-set.txt",
		},
		{
			name:   "check release name",
			cmd:    "template --name test " + chartPath,
			golden: "output/template-name.txt",
		},
		{
			name:   "check notes",
			cmd:    "template --notes " + chartPath,
			golden: "output/template-notes.txt",
		},
		{
			name:   "check values files",
			cmd:    "template --values " + chartPath + "/charts/subchartA/values.yaml " + chartPath,
			golden: "output/template-values-files.txt",
		},
		{
			name:   "check name template",
			cmd:    `template --name-template='foobar-{{ b64enc "abc" }}-baz' ` + chartPath,
			golden: "output/template-name-template.txt",
		},
		{
			name:   "check kube version",
			cmd:    "template --kube-version 1.6 " + chartPath,
			golden: "output/template-kube-version.txt",
		},
		{
			name:      "check no args",
			cmd:       "template",
			wantError: true,
			golden:    "output/template-no-args.txt",
		},
	}
	runTestCmd(t, tests)
}
