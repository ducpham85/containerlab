// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package host

import (
	"context"

	"github.com/srl-labs/containerlab/nodes"
	"github.com/srl-labs/containerlab/types"
)

var kindnames = []string{"host"}

func init() {
	nodes.Register(kindnames, func() nodes.Node {
		return new(host)
	})
}

type host struct {
	nodes.DefaultNode
}

func (s *host) Init(cfg *types.NodeConfig, opts ...nodes.NodeOption) error {
	s.Cfg = cfg
	for _, o := range opts {
		o(s)
	}

	return nil
}
func (*host) Deploy(_ context.Context) error { return nil }
func (*host) GetImages() map[string]string   { return map[string]string{} }

func (*host) WithMgmtNet(*types.MgmtNet) {}

func (*host) Delete(_ context.Context) error {
	return nil
}
