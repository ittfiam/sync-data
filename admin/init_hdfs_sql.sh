find /data/soft/sync-mysql/asset/describes/mysql2hdfs/sql -name "*.sql" -exec /data/soft/hive-1.2.2/bin/hive -S -f {} \;