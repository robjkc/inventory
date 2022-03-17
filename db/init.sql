create table if not exists items (
    id serial primary key,
    item_type varchar(10) not null,
    title varchar(100) not null,
    price decimal not null,
    title varchar(100) not null,
    author varchar(100) not null,
    director varchar(100) not null,
    year int not null,
    minutes decimal not null,
    tracks jsonb,
    chapters jsonb
);

