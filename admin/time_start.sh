#!/bin/sh

./sync-mysql-schedule > /data/soft/sync-mysql/log.`date +\%Y\%m\%d` 2>&1 &