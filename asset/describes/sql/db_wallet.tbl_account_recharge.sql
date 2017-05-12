drop table if exists tbl_account_recharge;
CREATE TABLE `tbl_account_recharge` (
  `operator_name` varchar(64)  COMMENT '操作人',
  `operate_key` varchar(64)  COMMENT '操作key，用于去重',
  `account_id` varchar(64)  COMMENT '充值账户id',
  `create_time` varchar(32)  COMMENT '充值时间',
  `money` double  COMMENT '充值金额',
  `mark` varchar(256)  COMMENT '备注',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`operate_key`),
  KEY `id_time` (`account_id`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8