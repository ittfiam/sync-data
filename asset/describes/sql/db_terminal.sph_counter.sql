drop table if exists sph_counter;
CREATE TABLE `sph_counter` (
  `counter_id` int(11) ,
  `max_id` int(11) ,
  PRIMARY KEY (`counter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8