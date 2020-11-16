create table country
(
    name         varchar(100),
    acronym      varchar(100),
    capital      varchar(100),
    calling_code varchar(100),
    latitude     float,
    longitude    float,
    primary key (name)
);

create table currency
(
    id         SERIAL,
    name       varchar(100),
    code       varchar(100),
    country_name varchar(100),
    primary key (id),
    foreign key (country_name) references country (name) on delete cascade
);

