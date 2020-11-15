create table country
(
    name         char(100),
    acronym      char(100),
    capital      char(100),
    calling_code char(100),
    latitude     char(100),
    longitude    char(100),
    primary key (name)
);

create table currency
(
    id         SERIAL,
    country_name char(100),
    code       char(100),
    name       char(100),
    primary key (id),
    foreign key (country_name) references country (name) on delete cascade
);

