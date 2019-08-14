/*
 * go-mydumper
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package main

import (
	"common"
	"fmt"

	ini "github.com/dlintw/goconf"
)

func parseDumperConfig(file string) (*common.Args, error) {
	args := &common.Args{}

	cfg, err := ini.ReadConfigFile(file)
	if err != nil {
		return nil, err
	}

	host, err := cfg.GetString("mysql", "host")
	if err != nil {
		return nil, err
	}
	port, err := cfg.GetInt("mysql", "port")
	if err != nil {
		return nil, err
	}
	user, err := cfg.GetString("mysql", "user")
	if err != nil {
		return nil, err
	}
	password, err := cfg.GetString("mysql", "password")
	if err != nil {
		return nil, err
	}
	database, err := cfg.GetString("mysql", "database")
	if err != nil {
		return nil, err
	}
	outdir, err := cfg.GetString("mysql", "outdir")
	if err != nil {
		return nil, err
	}
	sessionVars, err := cfg.GetString("mysql", "vars")
	if err != nil {
		return nil, err
	}
	chunksizemb, err := cfg.GetInt("mysql", "chunksize")
	if err != nil {
		return nil, err
	}
	table, _ := cfg.GetString("mysql", "table")

	args.Address = fmt.Sprintf("%s:%d", host, port)
	args.User = user
	args.Password = password
	args.Database = database
	args.Table = table
	args.Outdir = outdir
	args.ChunksizeInMB = chunksizemb
	args.SessionVars = sessionVars
	args.Threads = 16
	args.StmtSize = 1000000
	args.IntervalMs = 10 * 1000
	return args, nil
}
