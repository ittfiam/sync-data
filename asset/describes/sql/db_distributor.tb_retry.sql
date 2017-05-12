drop table if exists tb_retry;
CREATE TABLE `tb_retry` (
  `id` bigint(20)  AUTO_INCREMENT COMMENT '重试ID',
  `event` text  COMMENT '重试消息内容',
  `status` int(11)  COMMENT '重试状态 0,1：等待重试,2:不再重试',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=738 DEFAULT CHARSET=utf8