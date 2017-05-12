drop table if exists tbl_transaction_business;
CREATE TABLE `tbl_transaction_business` (
  `operate_person` varchar(64)  DEFAULT '' COMMENT '操作人',
  `business_time` varchar(32)  COMMENT '业务发生时间',
  `business_id` varchar(64)  COMMENT '业务id，用于去重',
  `detail_id` varchar(64)  COMMENT '明细id，用于去重',
  `pipeline_id` varchar(64)  COMMENT '流水号',
  `update_time` varchar(32)  DEFAULT '' COMMENT '业务最后更新时间',
  `status` int(11)  COMMENT '状态：0，待扣款；1，成功；2，余额不足；3，扣款失败；4，扣款成功；5，人工补偿',
  `money_from_id` varchar(64)  COMMENT '付款方id',
  `money_from_name` varchar(128)  COMMENT '付款方名',
  `money_to_id` varchar(64)  COMMENT '收款方id',
  `money_to_name` varchar(128)  COMMENT '收款方名',
  `money` double  COMMENT '业务金额',
  `fee_type` varchar(128)  COMMENT '费项',
  `mark` varchar(256)  COMMENT '备注',
  `extend` mediumtext COMMENT '数据信息',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`business_id`,`detail_id`),
  KEY `from_time` (`money_from_id`,`business_time`),
  KEY `to_time` (`money_to_id`,`business_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8