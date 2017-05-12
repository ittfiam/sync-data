drop table if exists tbl_extract_money;
CREATE TABLE `tbl_extract_money` (
  `uid` varchar(64)  COMMENT '业务员id',
  `symbol` varchar(4)  COMMENT '+/-',
  `operate_key` varchar(64)  COMMENT '操作key',
  `operate_type` int(11)  COMMENT '操作类型',
  `create_time` varchar(32)  COMMENT '时间',
  `update_time` varchar(32) DEFAULT '' COMMENT '最后更新时间',
  `money` double  COMMENT '金额',
  `mark` varchar(128)  COMMENT '备注',
  `bank_name` varchar(64) DEFAULT '' COMMENT '银行行名',
  `bank_code` varchar(64) DEFAULT '' COMMENT '银行行别码',
  `bank_card_id` varchar(32) DEFAULT '' COMMENT '银行卡卡号',
  `bank_card_user_name` varchar(32) DEFAULT '' COMMENT '银行卡用户名',
  `extract_time_start` varchar(32) DEFAULT '' COMMENT '提取起始时间',
  `extract_time_end` varchar(32) DEFAULT '' COMMENT '提取结束时间',
  `has_extract_return_flag` int(11)  COMMENT '余额是否包含提现返还',
  `status` int(11)  COMMENT '状态',
  `extend` mediumtext COMMENT '数据信息',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`operate_key`),
  KEY `uid` (`uid`),
  KEY `uid_create_time` (`uid`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8