drop table if exists tb_goods_sync_flag;
CREATE TABLE `tb_goods_sync_flag` (
  `id` int(11) ,
  `flag` int(11) ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8