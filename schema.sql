create table menu(
    id bigserial primary key ,
    name text not null ,
    cost int not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz,
    active bool not null default true,
    category_id bigint references menu_categories
);

create table menu_categories(
    id bigserial primary key,
    name text not null
);



create table roles(
    id bigserial primary key,
    name text not null
);


create table employers(
    id bigserial primary key,
    role_id bigint references roles,
    name text not null,
    fines float not null default 0,
    allowances float not null default 0,
    salary float not null default 1,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz,
    active bool not null default true,
    login text not null,
    password text not null
);

create table tokens(
    id bigserial primary key ,
    token text not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    expiration_time timestamptz not null default current_timestamp + interval '1day',
    employers_id bigint references employers
);


create table orders(
    id bigserial primary key,
    waiter_id bigint references employers,
    type_order bool not null default false,
    cost float not null default 0,
    total_cost float not null default 0,
    place bigint references hall_map,
    pay bool default false,
    service_charge bigint references service_charges default 1
);

create table orders_items(
    id bigserial primary key,
    order_id bigint references orders,
    menu_id bigint references menu,
    status bigint references ready_status
);

create table hall_map(
    id bigserial primary key,
    table_number int not null,
    zone_id bigint references zones,
    reserved bool not null default false
);


create table ready_status(
    id bigserial primary key,
    status text
);

create table zones(
    id bigserial primary key,
    name text
);

create table service_charges (
    id bigserial primary key ,
    percent float check (percent >= 0)
);