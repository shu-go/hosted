Windows HOSTS file editor command

[![Go Report Card](https://goreportcard.com/badge/github.com/shu-go/hosted)](https://goreportcard.com/report/github.com/shu-go/hosted)
![MIT License](https://img.shields.io/badge/License-MIT-blue)


# Usage 

## Subcommands

```
hosted - edit Windows HOSTS file(0.1.0)

Sub commands:
  list, ls
  add          add an entry with args: {IP} {Host} [Comment]
  delete, del  delete an entry with  --ip 192.168.1.200 and/or --host oldserver
  enable, o    change comment -> normal entry if exists
  disable, x   change normal entry -> comment if exists

Options:
  --hosts PATH   (default: C:\Windows\System32\Drivers\etc\hosts)

Usage:
  RUN AS ADMINISTRATOR

  # ADD new server "server01" as 192.168.1.201
  hosted add --ip 192.168.1.201 --host server01 --comment "new server"
  hosted add 192.168.1.201 server01 new server

  # REMOVE a wrong entry
  hosted delete --ip 192.168.1.210

  # COMMENT-OUT (disable) oldserver
  hosted disable --host oldserver
  hosted x --host oldserver

  # COMMENT-IN (enable) oldserver back
  hosted enable --host oldserver
  hosted o --host oldserver
```
