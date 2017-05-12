drop table if exists tbl_fix_money;
CREATE TABLE `tbl_fix_money` (
  `uid` varchar(64)  COMMENT '业务员id',
  `symbol` varchar(4)  COMMENT '+/-',
  `operate_key` varchar(64)  COMMENT '操作key',
  `operator_name` varchar(64)  COMMENT '操作人',
  `operate_type` int(11)  COMMENT '操作类型',
  `create_time` varchar(32)  COMMENT '时间',
  `money` double  COMMENT '金额',
  `mark` varchar(128)  COMMENT '备注',
  `extend` mediumtext COMMENT '数据信息',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`operate_key`),
  KEY `uid` (`uid`),
  KEY `uid_create_time` (`uid`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8