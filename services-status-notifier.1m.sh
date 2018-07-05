#!/usr/bin/env bash

# <bitbar.title>Service Status Notifier</bitbar.title>
# <bitbar.author>Martins Irbe</bitbar.author>
# <bitbar.author.github>martinsirbe</bitbar.author.github>
# <bitbar.dependencies>Go</bitbar.dependencies>
# <bitbar.desc>Checks specified services statuspage and reports on unhealthy service status</bitbar.desc>

export GOPATH=$HOME/dev/go_workspace # change me!
export PATH=$HOME/bin:/usr/local/bin:/usr/bin/:$GOPATH/bin

services-status-notifier-bitbar
