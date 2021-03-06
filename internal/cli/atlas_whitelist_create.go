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
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	cidrBlock        = "cidrBlock"
	ipAddress        = "ipAddress"
	awsSecurityGroup = "awsSecurityGroup"
)

type atlasWhitelistCreateOpts struct {
	globalOpts
	entry     string
	entryType string
	comment   string
	store     store.ProjectIPWhitelistCreator
}

func (opts *atlasWhitelistCreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *atlasWhitelistCreateOpts) Run() error {
	entry := opts.newWhitelist()
	result, err := opts.store.CreateProjectIPWhitelist(entry)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasWhitelistCreateOpts) newWhitelist() *atlas.ProjectIPWhitelist {
	projectIPWhitelist := &atlas.ProjectIPWhitelist{
		GroupID: opts.ProjectID(),
		Comment: opts.comment,
	}
	switch opts.entryType {
	case cidrBlock:
		projectIPWhitelist.CIDRBlock = opts.entry
	case ipAddress:
		projectIPWhitelist.IPAddress = opts.entry
	case awsSecurityGroup:
		projectIPWhitelist.AwsSecurityGroup = opts.entry
	}
	return projectIPWhitelist
}

// mongocli atlas whitelist(s) create [entry] --type cidrBlock|ipAddress [--comment comment] [--projectId projectId]
func AtlasWhitelistCreateBuilder() *cobra.Command {
	opts := &atlasWhitelistCreateOpts{}
	cmd := &cobra.Command{
		Use:   "create [entry]",
		Short: description.CreateWhitelist,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.entry = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.entryType, flags.Type, ipAddress, usage.WhitelistType)
	cmd.Flags().StringVar(&opts.comment, flags.Comment, "", usage.Comment)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
