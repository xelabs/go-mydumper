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
	"flag"
	"fmt"
	"os"

	"github.com/xelabs/go-mysqlstack/xlog"
)

var (
	flagConfig string

	log = xlog.NewStdLog(xlog.Level(xlog.INFO))
)

func init() {
	flag.StringVar(&flagConfig, "c", "", "config file")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -c conf/mydumper.ini.sample")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	if flagConfig == "" {
		usage()
		os.Exit(0)
	}

	args, err := parseConfig(flagConfig)
	common.AssertNil(err)

	if _, err := os.Stat(args.Outdir); os.IsNotExist(err) {
		x := os.MkdirAll(args.Outdir, 0777)
		common.AssertNil(x)
	}

	common.Dumper(log, args)
}
