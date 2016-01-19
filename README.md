# Honeycomb-ServerTeam
The readme of Python version is inside server_python_version directory

## Team member:
    Ruijian Wang
  	Zhoucheng Li
  	Handi Xu
  	Long He

## Usage Instruction [Golang]:
###### Dependency
```
// export GOROOT and GOPATH first
go get github.com/gocql/gocql
```  	
  	

###### On BIC Cluster
```
CREATE KEYSPACE Honey
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE honey.data (
  id int PRIMARY KEY,
  path varchar
);

CREATE TABLE honey.result (
  id int PRIMARY KEY,
  path varchar
);
```

###### Inside source directory
```
export GOROOT=/home/honeycomb/Go/go
export GOPATH=/home/honeycomb/HoneyBuzzard
go build srunner.go
./srunner &
```
   
