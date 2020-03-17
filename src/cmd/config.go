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
	"strings"

	ini "github.com/dlintw/goconf"
)

func parseDumperConfig(file string) (*common.Args, error) {
	args := &common.Args{
		Wheres: make(map[string]string, 0),
	}

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
	database, _ := cfg.GetString("mysql", "database")
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

	// Options
	if err := loadOptions(cfg, "where", args.Wheres); err != nil {
		return nil, err
	}

	var selects []string
	if selects, err = cfg.GetOptions("select"); err != nil {
		return nil, err
	}
	for _, tblcol := range selects {
		var table, column string
		split := strings.Split(tblcol, ".")
		table = split[0]
		column = split[1]

		if args.Selects == nil {
			args.Selects = make(map[string]map[string]string)
		}
		if args.Selects[table] == nil {
			args.Selects[table] = make(map[string]string, 0)
		}
		if args.Selects[table][column], err = cfg.GetString("select", tblcol); err != nil {
			return nil, err
		}
	}

	database_regexp, _ := cfg.GetString("database", "regexp")
	database_invert_regexp, err := cfg.GetBool("database", "invert_regexp")
	if err != nil {
		database_invert_regexp = false
	}

	args.Address = fmt.Sprintf("%s:%d", host, port)
	args.User = user
	args.Password = password
	args.Database = database
	args.DatabaseRegexp = database_regexp
	args.DatabaseInvertRegexp = database_invert_regexp
	args.Table = table
	args.Outdir = outdir
	args.ChunksizeInMB = chunksizemb
	args.SessionVars = sessionVars
	args.Threads = 16
	args.StmtSize = 1000000
	args.IntervalMs = 10 * 1000
	return args, nil
}

func loadOptions(cfg *ini.ConfigFile, section string, optMap map[string]string) error {
	var err error
	var opts []string

	if opts, err = cfg.GetOptions(section); err != nil {
		return err
	}

	for _, key := range opts {
		if optMap[key], err = cfg.GetString(section, key); err != nil {
			return err
		}
	}
	return nil
}
