/*
 * go-mydumper
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package common

import (
	"testing"

	"github.com/XeLabs/go-mysqlstack/driver"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/XeLabs/go-mysqlstack/xlog"
	"github.com/stretchr/testify/assert"
)

func TestLoader(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.DEBUG))
	fakedbs := driver.NewTestHandler(log)
	server, err := driver.MockMysqlServer(log, fakedbs)
	assert.Nil(t, err)
	defer server.Close()
	address := server.Addr()

	// fakedbs.
	{
		fakedbs.AddQueryPattern("use .*", &sqltypes.Result{})
		fakedbs.AddQuery("create database if not exists `test`", &sqltypes.Result{})
		fakedbs.AddQuery("create table `t1` (`a` int(11) default null,`b` varchar(100) default null) engine=innodb", &sqltypes.Result{})
		fakedbs.AddQueryPattern("insert into .*", &sqltypes.Result{})
	}

	args := &Args{
		Outdir:     "/tmp/dumpertest",
		User:       "mock",
		Password:   "mock",
		Threads:    16,
		Address:    address,
		IntervalMs: 500,
	}
	// Loader.
	{
		Loader(log, args)
	}
}
