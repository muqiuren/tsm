# TSM

Terminal Services Manager, where you can store calls into Terminal.

## Overview

- Use local database Sqlite as storage node information
- Node information is encrypted with an RSA 2048-bit key
- Fast storage/access to service nodes
- Use Cobra to provide a friendly terminal interaction experience
- Use tablewriter to provide friendly display effect

## Getting Started

### 1. Quick Use

If you are using linux or mac, you can execute the following command to get the compiled program file

```shell
wget https://github.com/muqiuren/tsm/releases/download/v1.0.0/tsm
```

### 2. Build



```shell
git clone https://github.com/muqiuren/tsm.git
cd tsm && go mod tidy
go build .
```

## Usage

1. create terminal node

```shell
tsm create -n server1 -s 127.0.0.1 -p 22 -u root -a 'admin123'
```

2. list terminal node

```shell
tsm list
```
Execute the above command, you can see the following output

```
+----+-----------+---------------+
| ID |  公网IP   |   节点名称    |
+----+-----------+---------------+
| 1  | *.*.*.190 | 测试          |
| 2  | *.*.*.41  | xx系统测试服   |
+----+-----------+---------------+
```
You can quickly enter the node according to the id corresponding to the node

3. go to terminal server

```shell
tsm go 1
```

4. remove terminal node

```shell
tsm remove 1
```

5. more to help

```shell
tsm help
```

## License

@暮秋人, 2022~time.Now

Released under the [MIT License](https://github.com/muqiuren/tsm/blob/master/License)