# NB
搞事情
创建价格监控表 语句：
```hive
 create table price_log_monitor 
(	
	id string,
	skuno string,
	areacode string,
	type string,
	node string,
	time bigint,
	isbatch bigint,
	status bigint,
	action bigint,
	result bigint
) 
PARTITIONED BY(dt string)
ROW FORMAT DELIMITED ‘\n’
FIELDS TERMINATED BY '\t'
STORED AS ORC;
```
