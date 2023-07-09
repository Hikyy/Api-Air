CREATE ROLE postgres;

\c postgres;

create sequence floors_floor_id_seq
    as integer;

alter sequence floors_floor_id_seq owner to postgres;

create sequence rooms_room_id_seq
    as integer;

alter sequence rooms_room_id_seq owner to postgres;

create sequence scenes_scene_id_seq
    as integer;

alter sequence scenes_scene_id_seq owner to postgres;

create sequence actions_action_id_seq
    as integer;

alter sequence actions_action_id_seq owner to postgres;

create sequence automations_automation_id_seq
    as integer;

alter sequence automations_automation_id_seq owner to postgres;

create sequence conditions_condition_id_seq
    as integer;

alter sequence conditions_condition_id_seq owner to postgres;

create table "user_groups"
(
    group_name    varchar(255) not null
        constraint "user_group_pk"
            primary key
);

INSERT INTO "user_groups" (group_name)
VALUES
    ('user'),
    ('admin'),
    ('owner');

alter table "user_groups"
    owner to postgres;

create table "buildings"
(
    building_id varchar(1) not null
        constraint "building_pk"
            primary key
);

alter table "buildings"
    owner to postgres;

create table "floors"
(
    floor_id     integer default nextval('floors_floor_id_seq'::regclass) not null
        constraint floors_pkey
            primary key,
    floor_number integer                                                  not null,
    building_id  varchar(1)                                               not null
        constraint floors_building_id_fkey
            references "buildings"
);

alter table "floors"
    owner to postgres;

alter sequence floors_floor_id_seq owned by "floors".floor_id;

create table "rooms"
(
    room_id     integer default nextval('rooms_room_id_seq'::regclass) not null
        constraint rooms_pkey
            primary key,
    room_number integer                                                not null,
    floor_id    integer                                                not null
        constraint rooms_floor_id_fkey
            references "floors"
);

alter table "rooms"
    owner to postgres;

alter sequence rooms_room_id_seq owned by "rooms".room_id;

create table "sensors"
(
    sensor_id   serial
        primary key,
    sensor_name varchar(50) not null,
    sensor_type varchar(50) not null,
    room_id     integer
        constraint "sensors_rooms_room_id_fk"
            references "rooms"
);

alter table "sensors"
    owner to postgres;

create table "actuators"
(
    actuator_id   serial
        primary key,
    actuator_name varchar(50) not null,
    actuator_type varchar(50) not null,
    room_id       integer
        constraint "actuators_rooms_room_id_fk"
            references "rooms"
);

alter table "actuators"
    owner to postgres;

create table "sensor_events"
(
    event_id        serial
        primary key,
    event_timestamp timestamp with time zone not null,
    event_data      jsonb,
    sensor_id       integer                  not null
        references "sensors"
);

alter table "sensor_events"
    owner to postgres;

CREATE SEQUENCE IF NOT EXISTS users_id_seq;

create table "users"
(
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "firstname" character varying(255) NOT NULL,
    "lastname" character varying(255) NOT NULL,
    "email" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    group_name           varchar(255)             default 'user'::character varying not null
        constraint "User_user_group_name_fk"
        references "user_groups",
    CONSTRAINT "users_id_key" UNIQUE ("id")
);

alter table "users"
    owner to postgres;

create unique index "user_email_uindex"
    on "users" (email);

create table "user_logs"
(
    log_id        serial
        primary key,
    log_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    log_data      jsonb,
    user_id       integer                                            not null
        references "users"("id")
);

alter table "user_logs"
    owner to postgres;

create table "actuator_states"
(
    actuator_id     integer not null
        primary key
        references "actuators",
    state           jsonb,
    last_updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table "actuator_states"
    owner to postgres;

create table "scenes"
(
    scene_id   integer default nextval('scenes_scene_id_seq'::regclass) not null
        primary key,
    scene_name varchar(255)                                             not null
);

alter table "scenes"
    owner to postgres;

create table "actions"
(
    action_id   integer default nextval('actions_action_id_seq'::regclass) not null
        primary key,
    action_name varchar(255)                                               not null,
    actuator_id integer                                                    not null
        references "actuators",
    state       jsonb
);

alter table "actions"
    owner to postgres;

create table "scene_actions"
(
    scene_id  integer not null
        references "scenes",
    action_id integer not null
        references "actions",
    primary key (scene_id, action_id)
);

alter table "scene_actions"
    owner to postgres;

create table "automations"
(
    automation_id   integer default nextval('automations_automation_id_seq'::regclass) not null
        primary key,
    automation_name varchar(255)                                                       not null
);

alter table "automations"
    owner to postgres;

alter sequence automations_automation_id_seq owned by "automations".automation_id;

create table "automation_actions"
(
    automation_id integer not null
        references "automations",
    action_id     integer not null
        references "actions",
    primary key (automation_id, action_id)
);

alter table "automation_actions"
    owner to postgres;

create table "conditions"
(
    condition_id   integer default nextval('conditions_condition_id_seq'::regclass) not null
        primary key,
    condition_name varchar(255)                                                     not null,
    sensor_id      integer                                                          not null
        references "sensors",
    data_key       varchar(255)                                                     not null,
    operator       varchar(10)                                                      not null,
    value          double precision                                                 not null
);

alter table "conditions"
    owner to postgres;

alter sequence conditions_condition_id_seq owned by "conditions".condition_id;

create table "automation_conditions"
(
    automation_id integer not null
        references "automations",
    condition_id  integer not null
        references "conditions",
    primary key (automation_id, condition_id)
);

alter table "automation_conditions"
    owner to postgres;