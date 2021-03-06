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
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type iamProjectsCreateOpts struct {
	globalOpts
	name  string
	store store.ProjectCreator
}

func (opts *iamProjectsCreateOpts) init() error {
	if opts.OrgID() == "" {
		return errMissingOrgID
	}

	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *iamProjectsCreateOpts) Run() error {
	projects, err := opts.store.CreateProject(opts.name, opts.OrgID())

	if err != nil {
		return err
	}

	return json.PrettyPrint(projects)
}

// mongocli iam project(s) create name [--orgId orgId]
func IAMProjectsCreateBuilder() *cobra.Command {
	opts := &iamProjectsCreateOpts{}
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: description.CreateProject,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.orgID, flags.OrgID, "", usage.OrgID)

	return cmd
}
