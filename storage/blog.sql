
create table blogs (
    id serial primary key not null,
    title varchar not null,
    description text not null,
    author varchar not null,
    created_at timestamp default current_timestamp
);