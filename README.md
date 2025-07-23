
---

###  README.md

# VekoLoad - Lightning Fast Load Testing

[![Build Status](https://img.shields.io/github/actions/workflow/status/vekoload/vekoload/build.yml?style=flat-square)](https://github.com/vekoload/vekoload/actions)
[![GitHub release](https://img.shields.io/github/v/release/vekoload/vekoload?style=flat-square)](https://github.com/vekoload/vekoload/releases)

Single-binary load testing tool for HTTP, WebSocket, and gRPC services. 
Designed for high-performance scenarios with minimal resource usage.

## Features

-  10x faster than k6 (tested at 50k RPS)
-  Single binary - no dependencies
-  Real-time metrics & HTML reports
-  WebSocket & gRPC support
-  TOML configuration files

## Installation

```bash
# Linux/macOS
curl -sL https://get.vekoload.io | bash

# Windows
iwr -useb https://get.vekoload.io/win | iex
