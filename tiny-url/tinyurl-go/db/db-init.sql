-- https://medium.com/@adamhooper/in-mysql-never-use-utf8-use-utf8mb4-11761243e434
create database tinyurl CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
create user tinyurl_user identified by 'Adeg*#%23f';  
grant insert, select, update, delete on tinyurl.* to tinyurl_user;
  
-- https://stackoverflow.com/questions/417142/what-is-the-maximum-length-of-a-url-in-different-browsers  
create table tinyurl.urls (
  id int,
  full_url varchar(2048),
  short_url_key varchar(8),
  primary key (id),
  unique key (short_url_key)
) ENGINE=INNODB;


create table tinyurl.sequences (
  sequence_no int auto_increment, 
  PRIMARY KEY (sequence_no)
) ENGINE=INNODB;

