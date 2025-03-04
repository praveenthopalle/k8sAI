/*
Copyright 2023 The k8sAI Authors.
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

package filters

import (
	"github.com/spf13/cobra"
)

var FiltersCmd = &cobra.Command{
	Use:     "filters",
	Aliases: []string{"filter"},
	Short:   "Manage filters for analyzing Kubernetes resources",
	Long: `The filters command allows you to manage filters that are used to analyze Kubernetes resources.
	You can list available filters to analyze resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
	},
}

func init() {
	FiltersCmd.AddCommand(listCmd)
	FiltersCmd.AddCommand(addCmd)
	FiltersCmd.AddCommand(removeCmd)
}
