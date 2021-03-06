// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/atmcfg"
)

type opsManagerClustersStartupOpts struct {
	globalOpts
	name    string
	confirm bool
	store   store.AutomationPatcher
}

func (opts *opsManagerClustersStartupOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *opsManagerClustersStartupOpts) Run() error {
	if !opts.confirm {
		return nil
	}
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	if !search.ClusterExists(current, opts.name) {
		return fmt.Errorf("cluster '%s' doesn't exist", opts.name)
	}

	atmcfg.Startup(current, opts.name)

	if err := opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(deploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

func (opts *opsManagerClustersStartupOpts) Confirm() error {
	if opts.confirm {
		return nil
	}
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to startup: %s", opts.name),
	}
	return survey.AskOne(prompt, &opts.confirm)
}

// mongocli cloud-manager cluster(s) startup [name] --projectId projectId [--force]
func OpsManagerClustersStartupBuilder() *cobra.Command {
	opts := &opsManagerClustersStartupOpts{}
	cmd := &cobra.Command{
		Use:   "startup [name]",
		Short: description.StartUpCluster,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.initStore); err != nil {
				return err
			}
			opts.name = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
