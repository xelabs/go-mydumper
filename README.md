[![Build Status](https://travis-ci.org/XeLabs/go-mydumper.png)](https://travis-ci.org/XeLabs/go-mydumper) [![codecov.io](https://codecov.io/gh/XeLabs/go-mydumper/graphs/badge.svg)](https://codecov.io/gh/XeLabs/go-mydumper/branch/master)

# go-mydumper

***go-mydumper*** is a multi-threaded MySQL backup and restore tool, and it is compatible with [maxbube/mydumper](https://github.com/maxbube/mydumper) in the layout.


## Build

```
$git clone https://github.com/XeLabs/go-mydumper
$cd go-mydumper
$make
$./bin/mydumper --help
$./bin/myloader --help
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

Examples:
$./bin/mydumper -h 127.0.0.1 -P 3306 -u mock -p mock -db test  -o test.sql
 2017/09/07 11:44:21.867412 dumper.go:58: 	  [INFO]  	dumping.database[test].schema...
 2017/09/07 11:44:21.867572 dumper.go:68: 	  [INFO]  	dumping.table[test.t1].schema...
 2017/09/07 11:44:21.867598 dumper.go:191: 	  [INFO]  	dumping.table[test.t1].datas.thread[1]...
 2017/09/07 11:44:21.867674 dumper.go:68: 	  [INFO]  	dumping.table[test.t2].schema...
 2017/09/07 11:44:21.867726 dumper.go:191: 	  [INFO]  	dumping.table[test.t2].datas.thread[2]...
 2017/09/07 11:44:21.957461 dumper.go:128: 	  [INFO]  	dumping.table[test.t1].rows[29960].bytes[1MB].part[1].thread[1]
 2017/09/07 11:44:22.059075 dumper.go:128: 	  [INFO]  	dumping.table[test.t1].rows[59920].bytes[2MB].part[2].thread[1]
 2017/09/07 11:44:22.062807 dumper.go:128: 	  [INFO]  	dumping.table[test.t2].rows[29960].bytes[1MB].part[1].thread[2]
 2017/09/07 11:44:22.150842 dumper.go:128: 	  [INFO]  	dumping.table[test.t1].rows[89880].bytes[3MB].part[3].thread[1]
 2017/09/07 11:44:22.202011 dumper.go:128: 	  [INFO]  	dumping.table[test.t2].rows[59920].bytes[2MB].part[2].thread[2]
 2017/09/07 11:44:22.275228 dumper.go:128: 	  [INFO]  	dumping.table[test.t1].rows[119840].bytes[4MB].part[4].thread[1]
 2017/09/07 11:44:22.299726 dumper.go:128: 	  [INFO]  	dumping.table[test.t2].rows[89880].bytes[3MB].part[3].thread[2]
 2017/09/07 11:44:22.368031 dumper.go:205: 	  [INFO]  	dumping.allbytes[8MB].allrows[248681].time[0.50sec].rates[15.99MB/sec]...
 2017/09/07 11:44:22.380971 dumper.go:128: 	  [INFO]  	dumping.table[test.t1].rows[149800].bytes[5MB].part[5].thread[1]
 2017/09/07 11:44:22.436246 dumper.go:128: 	  [INFO]  	dumping.table[test.t2].rows[119840].bytes[4MB].part[4].thread[2]
 2017/09/07 11:44:22.485945 dumper.go:128: 	  [INFO]  	dumping.table[test.t1].rows[179760].bytes[6MB].part[6].thread[1]
 2017/09/07 11:44:22.549684 dumper.go:128: 	  [INFO]  	dumping.table[test.t2].rows[149800].bytes[5MB].part[5].thread[2]
 2017/09/07 11:44:22.556675 dumper.go:145: 	  [INFO]  	dumping.table[test.t1].done.allrows[201710].allbytes[6MB].thread[1]...
 2017/09/07 11:44:22.556698 dumper.go:193: 	  [INFO]  	dumping.table[test.t1].datas.thread[1].done...
 2017/09/07 11:44:22.606178 dumper.go:128: 	  [INFO]  	dumping.table[test.t2].rows[179760].bytes[6MB].part[6].thread[2]
 2017/09/07 11:44:22.653431 dumper.go:145: 	  [INFO]  	dumping.table[test.t2].done.allrows[201710].allbytes[6MB].thread[2]...
 2017/09/07 11:44:22.653474 dumper.go:193: 	  [INFO]  	dumping.table[test.t2].datas.thread[2].done...
 2017/09/07 11:44:22.653529 dumper.go:211: 	  [INFO]  	dumping.all.done.cost[0.79sec].allrows[403420].allbytes[14119700].rate[16.54MB/s]
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
$./bin/myloader -h 127.0.0.1 -P 3306 -u mock -p mock -d test.sql
 2017/09/07 11:44:23.499584 loader.go:93: 	  [INFO]  	restoring.database[test]
 2017/09/07 11:44:23.499730 loader.go:117: 	  [INFO]  	restoring.schema[test.t1]
 2017/09/07 11:44:23.499907 loader.go:117: 	  [INFO]  	restoring.schema[test.t2]
 2017/09/07 11:44:23.499995 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00006].thread[2]
 2017/09/07 11:44:23.500087 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00001].thread[7]
 2017/09/07 11:44:23.500185 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00004].thread[8]
 2017/09/07 11:44:23.500454 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00002].thread[12]
 2017/09/07 11:44:23.500317 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00005].thread[3]
 2017/09/07 11:44:23.500579 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00003].thread[14]
 2017/09/07 11:44:23.502278 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00003].thread[4]
 2017/09/07 11:44:23.500549 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00002].thread[13]
 2017/09/07 11:44:23.500463 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00007].thread[5]
 2017/09/07 11:44:23.500333 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00005].thread[9]
 2017/09/07 11:44:23.500358 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00001].thread[10]
 2017/09/07 11:44:23.500673 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00007].thread[15]
 2017/09/07 11:44:23.500379 loader.go:133: 	  [INFO]  	restoring.tables[t1].parts[00006].thread[11]
 2017/09/07 11:44:23.500516 loader.go:133: 	  [INFO]  	restoring.tables[t2].parts[00004].thread[6]
 2017/09/07 11:44:23.878522 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00007].thread[15].done...
 2017/09/07 11:44:23.956317 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00001].thread[10].done...
 2017/09/07 11:44:23.958765 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00001].thread[7].done...
 2017/09/07 11:44:23.967950 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00006].thread[11].done...
 2017/09/07 11:44:23.994385 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00005].thread[3].done...
 2017/09/07 11:44:24.000229 loader.go:199: 	  [INFO]  	restoring.allbytes[5MB].time[0.50sec].rates[10.00MB/sec]...
 2017/09/07 11:44:24.020923 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00002].thread[12].done...
 2017/09/07 11:44:24.069449 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00003].thread[14].done...
 2017/09/07 11:44:24.103210 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00006].thread[2].done...
 2017/09/07 11:44:24.127812 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00007].thread[5].done...
 2017/09/07 11:44:24.151496 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00002].thread[13].done...
 2017/09/07 11:44:24.197424 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00005].thread[9].done...
 2017/09/07 11:44:24.198938 loader.go:149: 	  [INFO]  	restoring.tables[t2].parts[00004].thread[6].done...
 2017/09/07 11:44:24.204801 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00004].thread[8].done...
 2017/09/07 11:44:24.205811 loader.go:149: 	  [INFO]  	restoring.tables[t1].parts[00003].thread[4].done...
 2017/09/07 11:44:24.205916 loader.go:205: 	  [INFO]  	restoring.all.done.cost[0.71sec].allbytes[14.00MB].rate[19.83MB/s]
```
