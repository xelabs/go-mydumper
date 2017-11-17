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
	flag_overwrite_tables                                                                        bool
	flag_threads, flag_port, flag_2port, flag_stmt_size                                          int
	flag_user, flag_passwd, flag_host, flag_2user, flag_2passwd, flag_2host, flag_db, flag_table string

	log = xlog.NewStdLog(xlog.Level(xlog.INFO))
)

func init() {
	flag.StringVar(&flag_user, "u", "", "Upstream username with privileges to run the streamer")
	flag.StringVar(&flag_passwd, "p", "", "Upstream user password")
	flag.StringVar(&flag_host, "h", "", "The upstream host to connect to")
	flag.IntVar(&flag_port, "P", 3306, "Upstream TCP/IP port to connect to")
	flag.StringVar(&flag_2user, "2u", "", "Downstream username with privileges to run the streamer")
	flag.StringVar(&flag_2passwd, "2p", "", "Downstream user password")
	flag.StringVar(&flag_2host, "2h", "", "The downstream host to connect to")
	flag.IntVar(&flag_2port, "2P", 3306, "Downstream TCP/IP port to connect to")
	flag.StringVar(&flag_db, "db", "", "Database to stream")
	flag.StringVar(&flag_table, "table", "", "Table to stream")
	flag.IntVar(&flag_threads, "t", 16, "Number of threads to use")
	flag.IntVar(&flag_stmt_size, "s", 1000000, "Attempted size of INSERT statement in bytes")
	flag.BoolVar(&flag_overwrite_tables, "o", false, "Drop tables if they already exist")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -2h [DOWNSTREAM-HOST] -2P [DOWNSTREAM-PORT] -2u [DOWNSTREAM-USER] -2p [DOWNSTREAM-PASSWORD] -db [DATABASE] [-o]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	if flag_host == "" || flag_user == "" || flag_passwd == "" || flag_db == "" || flag_2host == "" || flag_2user == "" || flag_2passwd == "" {
		usage()
		os.Exit(0)
	}

	args := &common.Args{
		User:            flag_user,
		Password:        flag_passwd,
		Address:         fmt.Sprintf("%s:%d", flag_host, flag_port),
		ToUser:          flag_2user,
		ToPassword:      flag_2passwd,
		ToAddress:       fmt.Sprintf("%s:%d", flag_2host, flag_2port),
		Database:        flag_db,
		Table:           flag_table,
		Threads:         flag_threads,
		StmtSize:        flag_stmt_size,
		IntervalMs:      10 * 1000,
		OverwriteTables: flag_overwrite_tables,
	}
	common.Streamer(log, args)
}
