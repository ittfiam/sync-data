drop table if exists tb_user_basic;
CREATE TABLE `tb_user_basic` (
  `uid` varchar(32)  COMMENT '用户ID',
  `user_type` varchar(16)  COMMENT '用户类型',
  `phone` varchar(32)  COMMENT '手机号码',
  `email` varchar(128)  COMMENT '电子邮箱',
  `password` varchar(64)  COMMENT '用户密码',
  `jpush_id` varchar(64)  COMMENT '极光推送ID',
  `wechat_openid` varchar(64)  COMMENT '微信openid',
  `status` int(11)  DEFAULT '0' COMMENT '用户状态: 0=正常，1=删除，2=禁用',
  `pic_key` varchar(128)  DEFAULT '' COMMENT '用户图片key',
  `register_time` datetime  COMMENT '添加时间',
  `create_time` datetime  COMMENT '创建记录时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新记录时间',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `phone` (`phone`,`user_type`),
  KEY `email` (`email`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='基础用户信息表'