[![Build Status](https://travis-ci.org/xelabs/go-mydumper.png)](https://travis-ci.org/xelabs/go-mydumper) [![Go Report Card](https://goreportcard.com/badge/github.com/xelabs/go-mydumper)](https://goreportcard.com/report/github.com/xelabs/go-mydumper) [![codecov.io](https://codecov.io/gh/xelabs/go-mydumper/graphs/badge.svg)](https://codecov.io/gh/xelabs/go-mydumper/branch/master)

# go-mydumper

***go-mydumper*** is a multi-threaded MySQL backup and restore tool, and it is compatible with [maxbube/mydumper](https://github.com/maxbube/mydumper) in the layout.


## Build

```
$git clone https://github.com/xelabs/go-mydumper
$cd go-mydumper
$make build
$./bin/mydumper   -h
$./bin/myloader   -h
$./bin/mystreamer -h
```

## Test

```
$make test
```

## Usage

### mydumper

```
./bin/mydumper --help
Usage: ./bin/mydumper -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -db [DATABASE] -o [OUTDIR]
  -F int
    	Split tables into chunks of this output file size. This value is in MB (default 128)
  -P int
    	TCP/IP port to connect to (default 3306)
  -db string
    	Database to dump
  -h string
    	The host to connect to
  -o string
    	Directory to output files to
  -p string
    	User password
  -s int
    	Attempted size of INSERT statement in bytes (default 1000000)
  -t int
    	Number of threads to use (default 16)
  -table string
    	Table to dump
  -u string
    	Username with privileges to run the dump
  -vars string
    	Session variables

Examples:
$./bin/mydumper -h 192.168.0.1 -P 3306 -u mock -p mock -db sbtest  -o sbtest.sql
 2017/10/25 13:12:52.933391 dumper.go:35:         [INFO]        dumping.database[sbtest].schema...
 2017/10/25 13:12:52.937743 dumper.go:45:         [INFO]        dumping.table[sbtest.benchyou0].schema...
 2017/10/25 13:12:52.937791 dumper.go:168:        [INFO]        dumping.table[sbtest.benchyou0].datas.thread[1]...
 2017/10/25 13:12:52.939008 dumper.go:45:         [INFO]        dumping.table[sbtest.benchyou1].schema...
 2017/10/25 13:12:52.939055 dumper.go:168:        [INFO]        dumping.table[sbtest.benchyou1].datas.thread[2]...
 2017/10/25 13:12:55.611905 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou0].rows[633987].bytes[128MB].part[1].thread[1]
 2017/10/25 13:12:55.765127 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou1].rows[633987].bytes[128MB].part[1].thread[2]
 2017/10/25 13:12:58.146093 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou0].rows[1266050].bytes[256MB].part[2].thread[1]
 2017/10/25 13:12:58.253219 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou1].rows[1266054].bytes[256MB].part[2].thread[2]

 ...
 [stripped]
 ...

 2017/10/25 13:13:02.939278 dumper.go:182:        [INFO]        dumping.allbytes[1024MB].allrows[5054337].time[10.01sec].rates[102.34MB/sec]...
 2017/10/25 13:13:35.496439 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou1].rows[11345659].bytes[2304MB].part[18].thread[2]
 2017/10/25 13:13:37.627178 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou0].rows[11974624].bytes[2432MB].part[19].thread[1]
 2017/10/25 13:13:37.753966 dumper.go:105:        [INFO]        dumping.table[sbtest.benchyou1].rows[11974630].bytes[2432MB].part[19].thread[2]
 2017/10/25 13:13:39.453430 dumper.go:122:        [INFO]        dumping.table[sbtest.benchyou0].done.allrows[12486842].allbytes[2536MB].thread[1]...
 2017/10/25 13:13:39.453462 dumper.go:170:        [INFO]        dumping.table[sbtest.benchyou0].datas.thread[1].done...
 2017/10/25 13:13:39.622390 dumper.go:122:        [INFO]        dumping.table[sbtest.benchyou1].done.allrows[12484135].allbytes[2535MB].thread[2]...
 2017/10/25 13:13:39.622423 dumper.go:170:        [INFO]        dumping.table[sbtest.benchyou1].datas.thread[2].done...
 2017/10/25 13:13:39.622454 dumper.go:188:        [INFO]        dumping.all.done.cost[46.69sec].allrows[24970977].allbytes[5318557708].rate[108.63MB/s]
```

The dump files:
```
$ ls sbtest.sql/
metadata                    sbtest.benchyou0.00009.sql  sbtest.benchyou0.00018.sql   sbtest.benchyou1.00006.sql  sbtest.benchyou1.00015.sql
sbtest.benchyou0.00001.sql  sbtest.benchyou0.00010.sql  sbtest.benchyou0.00019.sql   sbtest.benchyou1.00007.sql  sbtest.benchyou1.00016.sql
sbtest.benchyou0.00002.sql  sbtest.benchyou0.00011.sql  sbtest.benchyou0.00020.sql   sbtest.benchyou1.00008.sql  sbtest.benchyou1.00017.sql
sbtest.benchyou0.00003.sql  sbtest.benchyou0.00012.sql  sbtest.benchyou0-schema.sql  sbtest.benchyou1.00009.sql  sbtest.benchyou1.00018.sql
sbtest.benchyou0.00004.sql  sbtest.benchyou0.00013.sql  sbtest.benchyou1.00001.sql   sbtest.benchyou1.00010.sql  sbtest.benchyou1.00019.sql
sbtest.benchyou0.00005.sql  sbtest.benchyou0.00014.sql  sbtest.benchyou1.00002.sql   sbtest.benchyou1.00011.sql  sbtest.benchyou1.00020.sql
sbtest.benchyou0.00006.sql  sbtest.benchyou0.00015.sql  sbtest.benchyou1.00003.sql   sbtest.benchyou1.00012.sql  sbtest.benchyou1-schema.sql
sbtest.benchyou0.00007.sql  sbtest.benchyou0.00016.sql  sbtest.benchyou1.00004.sql   sbtest.benchyou1.00013.sql  sbtest-schema-create.sql
sbtest.benchyou0.00008.sql  sbtest.benchyou0.00017.sql  sbtest.benchyou1.00005.sql   sbtest.benchyou1.00014.sql
```

### myloader

```
$ ./bin/myloader --help
Usage: ./bin/myloader -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -d  [DIR]
  -P int
    	TCP/IP port to connect to (default 3306)
  -d string
    	Directory of the dump to import
  -h string
    	The host to connect to
  -p string
    	User password
  -t int
    	Number of threads to use (default 16)
  -u string
    	Username with privileges to run the loader

Examples:
$./bin/myloader -h 192.168.0.2 -P 3306 -u mock -p mock -d sbtest.sql
 2017/10/25 13:04:17.396002 loader.go:75:         [INFO]        restoring.database[sbtest]
 2017/10/25 13:04:17.458076 loader.go:99:         [INFO]        restoring.schema[sbtest.benchyou0]
 2017/10/25 13:04:17.516236 loader.go:99:         [INFO]        restoring.schema[sbtest.benchyou1]
 2017/10/25 13:04:17.516389 loader.go:115:        [INFO]        restoring.tables[benchyou0].parts[00015].thread[1]
 2017/10/25 13:04:17.516456 loader.go:115:        [INFO]        restoring.tables[benchyou0].parts[00005].thread[2]

...
[stripped]
...

 2017/10/25 13:05:27.783560 loader.go:131:        [INFO]        restoring.tables[benchyou1].parts[00005].thread[9].done...
 2017/10/25 13:05:36.133758 loader.go:181:        [INFO]        restoring.allbytes[4087MB].time[78.62sec].rates[51.99MB/sec]...
 2017/10/25 13:05:44.759183 loader.go:131:        [INFO]        restoring.tables[benchyou0].parts[00001].thread[3].done...
 2017/10/25 13:05:46.133728 loader.go:181:        [INFO]        restoring.allbytes[4216MB].time[88.62sec].rates[47.58MB/sec]...
 2017/10/25 13:05:46.567156 loader.go:131:        [INFO]        restoring.tables[benchyou1].parts[00016].thread[6].done...
 2017/10/25 13:05:50.612200 loader.go:131:        [INFO]        restoring.tables[benchyou0].parts[00008].thread[10].done...
 2017/10/25 13:05:51.131155 loader.go:131:        [INFO]        restoring.tables[benchyou0].parts[00014].thread[2].done...
 2017/10/25 13:05:51.185629 loader.go:131:        [INFO]        restoring.tables[benchyou0].parts[00011].thread[1].done...
 2017/10/25 13:05:51.836354 loader.go:131:        [INFO]        restoring.tables[benchyou1].parts[00004].thread[0].done...
 2017/10/25 13:05:52.286931 loader.go:131:        [INFO]        restoring.tables[benchyou1].parts[00006].thread[11].done...
 2017/10/25 13:05:52.602444 loader.go:131:        [INFO]        restoring.tables[benchyou0].parts[00019].thread[8].done...
 2017/10/25 13:05:52.602573 loader.go:187:        [INFO]        restoring.all.done.cost[95.09sec].allbytes[5120.00MB].rate[53.85MB/s]
```

### mystreamer

Streaming mode, dumps datas from upstream to downstream in parallel instead of dumping to the out directory.
```
$./bin/mystreamer
Usage: ./bin/mystreamer -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -db [DATABASE] -2h [DOWNSTREAM-HOST] -2P [DOWNSTREAM-PORT] -2u [DOWNSTREAM-USER] -2p [DOWNSTREAM-PASSWORD] [-2db DOWNSTREAM-DATABASE] [-o]
  -2P int
    	Downstream TCP/IP port to connect to (default 3306)
  -2db string
    	Downstream database, default is same as upstream db
  -2h string
    	The downstream host to connect to
  -2p string
    	Downstream user password
  -2u string
    	Downstream username with privileges to run the streamer
  -P int
    	Upstream TCP/IP port to connect to (default 3306)
  -db string
    	Database to stream
  -h string
    	The upstream host to connect to
  -o	Drop tables if they already exist
  -p string
    	Upstream user password
  -s int
    	Attempted size of INSERT statement in bytes (default 1000000)
  -t int
    	Number of threads to use (default 16)
  -table string
    	Table to stream
  -u string
    	Upstream username with privileges to run the streamer

Examples:
$./bin/mystreamer -h 192.168.0.2 -P 3306 -u mock -p mock -2h 192.168.0.3 -2P 3306 -2u mock -2p mock -db sbtest
 2017/11/17 23:52:46.424107 streamer.go:30:       [INFO]        streaming.database[sbtest].schema...
 2017/11/17 23:52:46.458807 streamer.go:50:       [INFO]        streaming.table[sbtest.benchyou0].schema...
 2017/11/17 23:52:46.458892 streamer.go:165:      [INFO]        streaming.table[sbtest.benchyou0].datas.thread[1]...
 2017/11/17 23:52:46.509911 streamer.go:50:       [INFO]        streaming.table[sbtest.benchyou1].schema...
 2017/11/17 23:52:46.510008 streamer.go:165:      [INFO]        streaming.table[sbtest.benchyou1].datas.thread[2]...
 2017/11/17 23:52:46.557443 streamer.go:50:       [INFO]        streaming.table[sbtest.benchyou2].schema...
 2017/11/17 23:52:46.557499 streamer.go:165:      [INFO]        streaming.table[sbtest.benchyou2].datas.thread[3]...
 2017/11/17 23:52:46.608464 streamer.go:50:       [INFO]        streaming.table[sbtest.benchyou3].schema...
 2017/11/17 23:52:46.608526 streamer.go:165:      [INFO]        streaming.table[sbtest.benchyou3].datas.thread[4]...

 ...
 [stripped]
 ...

 2017/11/17 23:52:53.056923 streamer.go:167:      [INFO]        streaming.table[sbtest.benchyou4].datas.thread[5].done...
 2017/11/17 23:52:53.319846 streamer.go:121:      [INFO]        streaming.table[sbtest.benchyou5].done.allrows[134490].allbytes[28MB].thread[6]...
 2017/11/17 23:52:53.319881 streamer.go:167:      [INFO]        streaming.table[sbtest.benchyou5].datas.thread[6].done...
 2017/11/17 23:52:53.338505 streamer.go:121:      [INFO]        streaming.table[sbtest.benchyou6].done.allrows[134035].allbytes[28MB].thread[7]...
 2017/11/17 23:52:53.338553 streamer.go:167:      [INFO]        streaming.table[sbtest.benchyou6].datas.thread[7].done...
 2017/11/17 23:52:53.465916 streamer.go:121:      [INFO]        streaming.table[sbtest.benchyou7].done.allrows[133240].allbytes[28MB].thread[8]...
 2017/11/17 23:52:53.465958 streamer.go:167:      [INFO]        streaming.table[sbtest.benchyou7].datas.thread[8].done...
 2017/11/17 23:52:53.465999 streamer.go:185:      [INFO]        streaming.all.done.cost[7.04sec].allrows[1070285].allbytes[240555976].rate[32.52MB/s]
```


## License

go-mydumper is released under the GPLv3. See LICENSE
