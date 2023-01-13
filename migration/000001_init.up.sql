CREATE TABLE proxy_list (
    id serial not null,
    types varchar(10),
    ip varchar(20),
    port int,
    speed int,
    anonlvl text,
    city text,
    country text,
    last_check timestamp
);
