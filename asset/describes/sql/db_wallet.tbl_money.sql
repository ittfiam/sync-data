drop table if exists tbl_money;
CREATE TABLE `tbl_money` (
  `uid` varchar(64)  COMMENT '业务员id',
  `total_money` decimal(16,2) DEFAULT '0.00' COMMENT '总金额',
  `can_extract_money` decimal(16,2) DEFAULT '0.00' COMMENT '可提现金额',
  `update_date_time` varchar(32)  COMMENT '更新时间',
  `has_extract_return_flag` int(11)  COMMENT '余额是否包含提现返还',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8