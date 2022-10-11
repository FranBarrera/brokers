// Copyright 2022 TriggerMesh Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"

	"github.com/alecthomas/kong"
	"go.uber.org/zap"

	"github.com/triggermesh/brokers/cmd/redis-broker/cmd"
	pkgcmd "github.com/triggermesh/brokers/pkg/cmd"
)

type Gateway struct {
	pkgcmd.Globals

	Start cmd.StartCmd `cmd:"" help:"Starts the TriggerMesh gateway."`
}

func main() {

	// TODO configure logger
	// zl, err := zap.NewProduction()
	zl, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	cli := Gateway{
		Globals: pkgcmd.Globals{
			Logger:  zl.Sugar(),
			Context: context.Background(),
		},
	}

	instance, err := os.Hostname()
	if err != nil {
		zl.Panic("error retrieving the host name", zap.Error(err))
	}

	kc := kong.Parse(&cli,
		kong.Vars{
			"instance_name": instance,
		})
	err = kc.Run(&cli.Globals)
	kc.FatalIfErrorf(err)
}
