drop table if exists tb_identity_bank_auth;
CREATE TABLE `tb_identity_bank_auth` (
  `identity_card` varchar(32)  COMMENT '身份证号',
  `bank_card` varchar(32)  COMMENT '银行卡号',
  `username` varchar(32)  COMMENT '用户真实姓名',
  `auth_result` varchar(12)  COMMENT '认证结果. 0: 信息一致, 1: 信息不一致',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`identity_card`,`bank_card`,`username`),
  KEY `index_identity_card` (`identity_card`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='身份认证记录表'