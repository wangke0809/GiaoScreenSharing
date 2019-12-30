# GiaoScreenSharing

[![Build Status](https://github.com/wangke0809/GiaoScreenSharing/workflows/Go/badge.svg)](https://github.com/wangke0809/GiaoScreenSharing/commits/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/wangke0809/screensharing)](https://goreportcard.com/report/github.com/wangke0809/screensharing)
[![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/wangke0809/GiaoScreenSharing?include_prereleases)](https://github.com/wangke0809/GiaoScreenSharing/releases/latest)
![GitHub](https://img.shields.io/github/license/wangke0809/GiaoScreenSharing)
[![GitHub All Releases](https://img.shields.io/github/downloads/wangke0809/GiaoScreenSharing/total)](https://github.com/wangke0809/GiaoScreenSharing/releases/latest)

A Cross Platform Screen Sharing tool based on UDP Multicast written in Go supporting Windows, Linux, macOS.

[中文说明]()

<p align="center">
  <img src="https://raw.githubusercontent.com/wangke0809/giaoscreensharing/master/docs/screenshot.png"/>
</p>

## Installation

### Binary file

Download the binary file form the [release page](https://github.com/wangke0809/GiaoScreenSharing/releases/latest).

The `server` will share screen to `client`.

### Building from source

- golang: 1.13+

```sh
git clone https://github.com/wangke0809/GiaoScreenSharing.git
mkdir dist && cd dist && go build ../cmd/client/ && go build -v ../cmd/server/
```

## Usage

### server

```sh
> ./server
> start using 224.0.0.1:9999
> FPS: 12.3, Send Block:   4, Capture:  76 (ms), Diff&JPEG Compress:   3 (ms) Send: 0.259 (ms)
```

See more configuration:

```sh
> ./server -h
> Usage of ./server:
    -block int
          screen transfer block size (default 150)
    -display int
          screen display index
    -host string
          udp host ip (default "224.0.0.1")
    -port int
          udp listen port (default 9999)
    -quality int
          jpeg compress quality (default 75)
```

### Client

```sh
> ./client
```

See more configuration:

```sh
> ./client -h
> Usage of ./client:
    -host string
          udp host ip (default "224.0.0.1")
    -port int
          udp listen port (default 9999)
```

