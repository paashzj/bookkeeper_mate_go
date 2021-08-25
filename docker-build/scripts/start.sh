#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
bash -x $DIR/start-daemon.sh
if [ $CLUSTER_INIT == "true" ]; then
    $BOOKKEEPER_HOME/mate/bookkeeper_mate 2>&1
else
  tail -f /dev/null
fi
