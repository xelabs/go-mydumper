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

	"github.com/XeLabs/go-mysqlstack/xlog"
)

var (
	flag_overwrite_tables                       bool
	flag_port, flag_threads                     int
	flag_user, flag_passwd, flag_host, flag_dir string

	log = xlog.NewStdLog(xlog.Level(xlog.INFO))
)

func init() {
	flag.StringVar(&flag_user, "u", "", "Username with privileges to run the loader")
	flag.StringVar(&flag_passwd, "p", "", "User password")
	flag.StringVar(&flag_host, "h", "", "The host to connect to")
	flag.IntVar(&flag_port, "P", 3306, "TCP/IP port to connect to")
	flag.StringVar(&flag_dir, "d", "", "Directory of the dump to import")
	flag.IntVar(&flag_threads, "t", 16, "Number of threads to use")
	flag.BoolVar(&flag_overwrite_tables, "o", false, "Drop tables if they already exist")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -d [DIR] [-o]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	if flag_host == "" || flag_user == "" || flag_passwd == "" || flag_dir == "" {
		usage()
		os.Exit(0)
	}

	args := &common.Args{
		User:            flag_user,
		Password:        flag_passwd,
		Address:         fmt.Sprintf("%s:%d", flag_host, flag_port),
		Outdir:          flag_dir,
		Threads:         flag_threads,
		IntervalMs:      10 * 1000,
		OverwriteTables: flag_overwrite_tables,
	}
	common.Loader(log, args)
}
