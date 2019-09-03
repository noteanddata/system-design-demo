
create keyspace tinyurl with replication = {'class':'SimpleStrategy', 'replication_factor':'3'};

use tinyurl;

create table sequences (id  UUID primary key, seq bigint );
insert into sequences (id, seq) values (uuid(), 1);


create table urls(short_url ascii primary key , long_url ascii );