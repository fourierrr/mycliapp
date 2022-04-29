[![Go Report Card](https://goreportcard.com/badge/github.com/fourierrr/mycliapp)](https://goreportcard.com/report/github.com/fourierrr/mycliapp)
[![GitHub license](https://img.shields.io/github/license/fourierrr/mycliapp)](https://github.com/fourierrr/mycliapp/blob/master/LICENSE)


`mycliapp`  is a small cli app that could show current time , just for testing the cobra package


## Installing

### Pre-built binaries

See the [release](https://github.com/fourierrr/mycliapp/releases) page for the full list of pre-built assets.

The commands here show `amd64` versions, `386` versions are available in the releases page.

**Linux**

```bash
export release=v1.0.0
curl -L -o mycliapp.tar.gz https://github.com/fourierrr/mycliapp/releases/download/${release}/mycliapp_${release}_Linux_arm64.tar.gz
tar -xvf mycliapp.tar.gz
cp mycliapp /usr/local/bin/mycliapp
```

**OSX**

```bash
export release=v1.0.0
curl -L -o mycliapp.tar.gz https://github.com/fourierrr/mycliapp/releases/download/${release}/mycliapp_${release}_Darwin_x86_64.tar.gz
tar -xvf mycliapp.tar.gz
mv mycliapp /usr/local/bin/mycliapp
```


**Windows**

In PowerShell v5+
```powershell
$url = "https://github.com/fourierrr/mycliapp/releases/download/v1.0.0/mycliapp_1.0.0_Windows_x86_64.tar.gz"
$output = "$PSScriptRoot\mycliapp.zip"

Invoke-WebRequest -Uri $url -OutFile $output
Expand-Archive "$PSScriptRoot\mycliapp.zip" -DestinationPath "$PSScriptRoot\mycliapp"
```



### Source

Using go modules, you can build mycliapp at any git tag:
```
$ GO111MODULE=on go get github.com/fourierrr/mycliapp/cmd/mycliapp@latest
```

This will download and compile `mycliapp` , note that you will need to be on a recent version of go which supports go modules.

## Usage
just a small app to show current time on 2 ways

A cli-app will show current time, just for test the package 'cobra', and try releas this app in github
```
Usage:
  mycliapp [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  show        A brief description of your command
  time        show current time

Flags:
  -h, --help     help for mycliapp
  -t, --toggle   Help message for toggle

Use "mycliapp [command] --help" for more information about a command.
```
