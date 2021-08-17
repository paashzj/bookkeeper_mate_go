#!/bin/bash

if [ $CLUSTER_INIT == "true" ]; then
    $BOOKKEEPER_HOME/mate/bookkeeper_mate 2>&1
else
    nohup $BOOKKEEPER_HOME/mate/bookkeeper_mate >$BOOKKEEPER_HOME/bookkeeper_mate.stdout.log 2>$BOOKKEEPER_HOME/bookkeeper_mate.stderr.log
fi