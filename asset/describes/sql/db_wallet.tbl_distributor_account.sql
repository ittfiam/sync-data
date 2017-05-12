drop table if exists tbl_distributor_account;
CREATE TABLE `tbl_distributor_account` (
  `id` varchar(64)  COMMENT '账户id',
  `description` varchar(256)  COMMENT '账户说明',
  `create_time` varchar(32)  COMMENT '账户创建时间',
  `update_time` varchar(32)  COMMENT '账户最后更新时间',
  `total_money` decimal(16,2) DEFAULT '0.00' COMMENT '总余额',
  `credit_limit_money` double DEFAULT '0' COMMENT '信用额度',
  `warning_money` double DEFAULT '0' COMMENT '预警额度',
  `can_overdraw_money` double DEFAULT '0' COMMENT '可透支额度',
  `status` int(11)  COMMENT '账户状态：0，其它；1，信用透支；',
  `flag` int(11)  COMMENT '标记：1，使用中；0，已废弃',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8