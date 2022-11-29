# TSM

Terminal Services Manager, where you can store calls into Terminal.

## Overview

- Use local database Sqlite as storage node information
- Node information is encrypted with an RSA 2048-bit key
- Fast storage/access to service nodes
- Use Cobra to provide a friendly terminal interaction experience
- Use tablewriter to provide friendly display effect

## Usage

### 1. Quick Use

If you are using linux or mac, you can execute the following command to get the compiled program file

```shell
wget https://mirrors.host900.com/https://raw.githubusercontent.com/muqiuren/tsm/releases/download/v1.0.0
```

### 2. Build

```shell
git clone https://github.com/muqiuren/tsm.git
cd tsm && go mod tidy
go build .
```

## License

@暮秋人, 2022~time.Now

Released under the [MIT License](https://github.com/muqiuren/tsm/blob/master/License)