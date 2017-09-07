/*
 * go-mydumper
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package mydumper

import (
	"common"
	"flag"
	"fmt"
	"os"

	"github.com/XeLabs/go-mysqlstack/xlog"
)

var (
	flag_chunksize, flag_threads, flag_port, flag_stmt_size          int
	flag_user, flag_passwd, flag_host, flag_db, flag_table, flag_dir string

	log = xlog.NewStdLog(xlog.Level(xlog.INFO))
)

func init() {
	flag.StringVar(&flag_user, "u", "", "Username with privileges to run the dump")
	flag.StringVar(&flag_passwd, "p", "", "User password")
	flag.StringVar(&flag_host, "h", "", "The host to connect to")
	flag.IntVar(&flag_port, "P", 3306, "TCP/IP port to connect to")
	flag.StringVar(&flag_db, "db", "", "Database to dump")
	flag.StringVar(&flag_table, "table", "", "Table to dump")
	flag.StringVar(&flag_dir, "o", "", "Directory to output files to")
	flag.IntVar(&flag_chunksize, "F", 128, "Split tables into chunks of this output file size. This value is in MB")
	flag.IntVar(&flag_threads, "t", 16, "Number of threads to use")
	flag.IntVar(&flag_stmt_size, "s", 1000000, "Attempted size of INSERT statement in bytes")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -db [DATABASE] -o [OUTDIR]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	if flag_host == "" || flag_user == "" || flag_passwd == "" || flag_db == "" {
		usage()
		os.Exit(0)
	}

	if _, err := os.Stat(flag_dir); os.IsNotExist(err) {
		x := os.MkdirAll(flag_dir, 0777)
		common.AssertNil(x)
	}

	args := &common.Args{
		User:          flag_user,
		Password:      flag_passwd,
		Address:       fmt.Sprintf("%s:%d", flag_host, flag_port),
		Database:      flag_db,
		Table:         flag_table,
		Outdir:        flag_dir,
		ChunksizeInMB: flag_chunksize,
		Threads:       flag_threads,
		StmtSize:      flag_stmt_size,
		IntervalMs:    10 * 1000,
	}

	common.Dumper(log, args)
}
