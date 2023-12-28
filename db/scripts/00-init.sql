create table if not exists todos(
    id            serial,
    title         varchar(64) not null,
    descr         text,
    done          boolean default false
);
