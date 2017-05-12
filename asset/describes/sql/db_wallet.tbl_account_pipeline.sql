drop table if exists tbl_account_pipeline;
CREATE TABLE `tbl_account_pipeline` (
  `business_id` varchar(64)  COMMENT '业务id，用于去重',
  `detail_id` varchar(64)  COMMENT '明细id，用于去重',
  `pipeline_id` varchar(64)  COMMENT '流水号',
  `money_from_id` varchar(64)  COMMENT '付款方id',
  `money_from_name` varchar(128)  COMMENT '付款方名',
  `money_to_id` varchar(64)  COMMENT '收款方id',
  `money_to_name` varchar(128)  COMMENT '收款方名',
  `account_id` varchar(64)  COMMENT '账户id',
  `account_name` varchar(128)  COMMENT '账户名',
  `create_time` varchar(32)  COMMENT '流水产生时间',
  `symbol` varchar(4)  COMMENT '+/-',
  `money` double  COMMENT '金额',
  `fee_type` varchar(128)  COMMENT '费项',
  `mark` varchar(128)  COMMENT '备注',
  `extend` mediumtext COMMENT '数据信息',
  PRIMARY KEY (`pipeline_id`,`account_id`),
  KEY `account_time` (`account_id`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8