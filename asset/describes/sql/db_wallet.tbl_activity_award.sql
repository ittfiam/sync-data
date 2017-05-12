drop table if exists tbl_activity_award;
CREATE TABLE `tbl_activity_award` (
  `uid` varchar(64)  COMMENT '业务员id',
  `operate_key` varchar(64)  COMMENT '操作key',
  `time` varchar(32)  COMMENT '时间',
  `award_type` int(11)  COMMENT '奖励类型',
  `money` double  COMMENT '新人奖金额',
  `flag` int(11)  COMMENT '标记:0，未处理；1，已处理',
  `extend` mediumtext COMMENT '数据信息',
  `extend1` varchar(64) DEFAULT '' COMMENT '扩展数据',
  `extend2` varchar(64) DEFAULT '' COMMENT '扩展数据',
  `extend3` varchar(64) DEFAULT '' COMMENT '扩展数据',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`operate_key`),
  KEY `uid` (`uid`),
  KEY `uid_time` (`uid`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8