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

create table "UserGroups"
(
    group_name    varchar(255) not null
        constraint "UserGroup_pk"
            primary key,
    group_options jsonb        not null
);

alter table "UserGroups"
    owner to postgres;

create table "Buildings"
(
    building_id varchar(1) not null
        constraint "Building_pk"
            primary key
);

alter table "Buildings"
    owner to postgres;

create table "Floors"
(
    floor_id     integer default nextval('floors_floor_id_seq'::regclass) not null
        constraint floors_pkey
            primary key,
    floor_number integer                                                  not null,
    building_id  varchar(1)                                               not null
        constraint floors_building_id_fkey
            references "Buildings"
);

alter table "Floors"
    owner to postgres;

alter sequence floors_floor_id_seq owned by "Floors".floor_id;

create table "Rooms"
(
    room_id     integer default nextval('rooms_room_id_seq'::regclass) not null
        constraint rooms_pkey
            primary key,
    room_number integer                                                not null,
    floor_id    integer                                                not null
        constraint rooms_floor_id_fkey
            references "Floors"
);

alter table "Rooms"
    owner to postgres;

alter sequence rooms_room_id_seq owned by "Rooms".room_id;

create table "Sensors"
(
    sensor_id   serial
        primary key,
    sensor_name varchar(50) not null,
    sensor_type varchar(50) not null,
    room_id     integer
        constraint "Sensors_Rooms_room_id_fk"
            references "Rooms"
);

alter table "Sensors"
    owner to postgres;

create table "Actuators"
(
    actuator_id   serial
        primary key,
    actuator_name varchar(50) not null,
    actuator_type varchar(50) not null,
    room_id       integer
        constraint "Actuators_Rooms_room_id_fk"
            references "Rooms"
);

alter table "Actuators"
    owner to postgres;

create table "SensorEvents"
(
    event_id        serial
        primary key,
    event_timestamp timestamp with time zone not null,
    event_data      jsonb,
    sensor_id       integer                  not null
        references "Sensors"
);

alter table "SensorEvents"
    owner to postgres;

create table "Users"
(
    user_id              serial
        primary key,
    user_firstname       varchar(255)                                               not null,
    user_lastname        varchar(255)                                               not null,
    user_email           varchar(255)                                               not null,
    user_password        varchar(1000),
    user_created_at      timestamp with time zone default CURRENT_TIMESTAMP,
    user_last_updated_at timestamp with time zone default CURRENT_TIMESTAMP,
    group_name           varchar(255)             default 'user'::character varying not null
        constraint "User_UserGroup_name_fk"
        references "UserGroups"
);

alter table "Users"
    owner to postgres;

create unique index "User_email_uindex"
    on "Users" (user_email);

create table "UserLogs"
(
    log_id        serial
        primary key,
    log_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    log_data      jsonb,
    user_id       integer                                            not null
        references "Users"
);

alter table "UserLogs"
    owner to postgres;

create table "ActuatorStates"
(
    actuator_id     integer not null
        primary key
        references "Actuators",
    state           jsonb,
    last_updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table "ActuatorStates"
    owner to postgres;

create table "Scenes"
(
    scene_id   integer default nextval('scenes_scene_id_seq'::regclass) not null
        primary key,
    scene_name varchar(255)                                             not null
);

alter table "Scenes"
    owner to postgres;

create table "Actions"
(
    action_id   integer default nextval('actions_action_id_seq'::regclass) not null
        primary key,
    action_name varchar(255)                                               not null,
    actuator_id integer                                                    not null
        references "Actuators",
    state       jsonb
);

alter table "Actions"
    owner to postgres;

create table "SceneActions"
(
    scene_id  integer not null
        references "Scenes",
    action_id integer not null
        references "Actions",
    primary key (scene_id, action_id)
);

alter table "SceneActions"
    owner to postgres;

create table "Automations"
(
    automation_id   integer default nextval('automations_automation_id_seq'::regclass) not null
        primary key,
    automation_name varchar(255)                                                       not null
);

alter table "Automations"
    owner to postgres;

alter sequence automations_automation_id_seq owned by "Automations".automation_id;

create table "AutomationActions"
(
    automation_id integer not null
        references "Automations",
    action_id     integer not null
        references "Actions",
    primary key (automation_id, action_id)
);

alter table "AutomationActions"
    owner to postgres;

create table "Conditions"
(
    condition_id   integer default nextval('conditions_condition_id_seq'::regclass) not null
        primary key,
    condition_name varchar(255)                                                     not null,
    sensor_id      integer                                                          not null
        references "Sensors",
    data_key       varchar(255)                                                     not null,
    operator       varchar(10)                                                      not null,
    value          double precision                                                 not null
);

alter table "Conditions"
    owner to postgres;

alter sequence conditions_condition_id_seq owned by "Conditions".condition_id;

create table "AutomationConditions"
(
    automation_id integer not null
        references "Automations",
    condition_id  integer not null
        references "Conditions",
    primary key (automation_id, condition_id)
);

alter table "AutomationConditions"
    owner to postgres;