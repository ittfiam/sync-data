drop table if exists tbl_transaction_business_total;
CREATE TABLE `tbl_transaction_business_total` (
  `business_time` varchar(32)  COMMENT '业务发生时间',
  `business_id` varchar(64)  COMMENT '业务id，用于去重',
  `extend` mediumtext COMMENT '数据信息',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`business_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8