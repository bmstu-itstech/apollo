create table if not exists departments (
    id integer primary key,
    name text not null,
    description text not null
);

create table if not exists disciplines (
    id integer primary key,
    name text not null
);

create table if not exists materials (
    uuid text primary key,
    name text not null,
    description text not null,
    url text not null,
    author text,
    views integer not null, 
    department_id integer references departments(id),
    discipline_id integer references disciplines(id),
    created_at timestamp not null,
);
