drop table if exists tbl_income_weekly;
CREATE TABLE `tbl_income_weekly` (
  `uid` varchar(64)  COMMENT '业务员id',
  `operate_key` varchar(64)  COMMENT '操作key',
  `time_start` varchar(32)  COMMENT '开始时间',
  `time_end` varchar(32)  COMMENT '结束时间',
  `money` double  COMMENT '金额',
  `mark` varchar(128)  COMMENT '备注',
  `extend` mediumtext COMMENT '数据信息',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`operate_key`),
  KEY `uid` (`uid`),
  KEY `uid_time_start_time_end` (`uid`,`time_start`,`time_end`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8