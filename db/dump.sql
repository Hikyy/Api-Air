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

create sequence conditions_id_seq
    as integer;

alter sequence conditions_id_seq owner to postgres;

create table "user_groups"
(
    group_name    varchar(255) not null
        constraint "user_group_pk"
            primary key
);
INSERT INTO "user_groups" (group_name)
VALUES ('user'), ('administrator');

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
    room_key   varchar(255)                                            not null,
    floor_id    integer                                                not null
        constraint rooms_floor_id_fkey
            references "floors"
);

alter table "rooms"
    owner to postgres;

alter sequence rooms_room_id_seq owned by "rooms".room_id;

create table "sensors"
(
    id   serial
        primary key,
    sensor_id   INT         not null,
    sensor_name varchar(50) not null,
    sensor_type varchar(50) not null,
    room_id     integer
        constraint "sensors_rooms_room_id_fk"
            references "rooms"
);

alter table "sensors"
    owner to postgres;


CREATE TABLE "actuators"
(
    id                serial primary key,
    actuator_name     varchar(255) not null,
    actuator_command  integer not null,
    data_key          varchar(25)  not null,
    destination_address varchar(255) not null ,
    room_id           integer not null,
    sensor_id         integer not null,
    constraint "actuators_rooms_room_id_fk" foreign key (room_id) references "rooms" (room_id),
    constraint "actuators_sensors_sensor_id_fk" foreign key (sensor_id) references "sensors" (id)
);

alter table "actuators"
    owner to postgres;

CREATE SEQUENCE sensor_events_event_id_seq;

CREATE TABLE "public"."sensor_events" (
  "id" integer DEFAULT nextval('sensor_events_event_id_seq') NOT NULL,
  "event_timestamp" timestamp(0),
  "event_data" jsonb,
  "sensor_id" integer NOT NULL,
  CONSTRAINT "sensor_events_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "sensor_events" OWNER TO postgres;

ALTER TABLE sensor_events
    ADD CONSTRAINT sensor_events_sensor_id_fkey
    FOREIGN KEY (sensor_id)
    REFERENCES sensors (id);

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
    id        serial
        primary key,
    log_timestamp timestamp(0) with time zone default CURRENT_TIMESTAMP not null,
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
    state           bool,
    last_updated_at timestamp(0) with time zone default CURRENT_TIMESTAMP
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

CREATE TABLE "conditions"
(
    id             serial primary key,
    automation_name varchar(255) not null,
    sensor_id      integer not null references "sensors",
    data_key       varchar(255) not null,
    operator       varchar(10) not null,
    value          double precision not null,
    actuator_id   integer not null,                                      -- Ajout de la colonne activator_id
    constraint "conditions_actuator_actuator_id_fk" foreign key (actuator_id) references "actuators" (id)  -- Ajout de la clé étrangère
);

alter table "conditions"
    owner to postgres;

alter sequence conditions_id_seq owned by "conditions".id;

CREATE OR REPLACE FUNCTION notify_event() RETURNS TRIGGER AS $$

    DECLARE
data json;
        notification json;

BEGIN

        -- Convert the old or new row to JSON, based on the kind of action.
        -- Action = DELETE?             -> OLD row
        -- Action = INSERT or UPDATE?   -> NEW row
        IF (TG_OP = 'CREATE') THEN
            data = row_to_json(OLD);
ELSE
            data = row_to_json(NEW);
END IF;

        -- Contruct the notification as a JSON string.
        notification = json_build_object(
                          'table',TG_TABLE_NAME,
                          'action', TG_OP,
                          'data', data);


        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('events',notification::text);

        -- Result is ignored since this is an AFTER trigger
RETURN NULL;
END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER products_notify_event
    AFTER INSERT OR UPDATE OR DELETE ON sensor_events
    FOR EACH ROW EXECUTE PROCEDURE notify_event();

INSERT INTO "buildings" (building_id)
VALUES ('A');

INSERT INTO "floors" (floor_number, building_id)
VALUES (1, 'A');

INSERT INTO "rooms" (room_number, room_key, floor_id)
    VALUES 
    ('105', '1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b', 1),
    ('106', 'a95cec4a-8aaf-4204-9fa2-b6c4aa8779e7', 1),
    ('107', '5e178fd2-5321-4cf5-b04c-4c6a8a827d88', 1);


    INSERT INTO sensors(sensor_id, sensor_name,sensor_type,room_id)
VALUES
    (102, 'ac', 'climatiseur',1),
    (116, 'atmospheric_pressure', 'pression_atmo',1),
    (100, 'persons', 'capteur_presence', 1),
    (119, 'level', 'niveau',1),
    (118, 'lux', 'capteur_luminosite',1),
    (128, 'kwh', 'consommation_elec',1),
    (101, 'heat', 'chaleur',1),
    (131, 'co2', 'capteur_co2',1),
    (114, 'humidity', 'capteur_humidite',1),
    (136, 'adc', 'analog_to_digital_converter',1),
    (112, 'temperature', 'temperature',1),
    (103, 'humidity', 'capteur_humidite',1),
    (115, 'motion', 'mouvement',1),
    (104, 'light', 'lumiere',1),
    (102, 'ac', 'climatiseur',2),
    (116, 'atmospheric_pressure', 'pression_atmo',2),
    (100, 'persons', 'capteur_presence', 2),
    (119, 'level', 'niveau',2),
    (118, 'lux', 'capteur_luminosite',2),
    (128, 'kwh', 'consommation_elec',2),
    (101, 'heat', 'chaleur',2),
    (131, 'co2', 'capteur_co2',2),
    (114, 'humidity', 'capteur_humidite',2),
    (136, 'adc', 'analog_to_digital_converter',2),
    (112, 'temperature', 'temperature',2),
    (103, 'humidity', 'capteur_humidite',2),
    (115, 'motion', 'mouvement',2),
    (104, 'light', 'lumiere',2),
    (102, 'ac', 'climatiseur',3),
    (116, 'atmospheric_pressure', 'pression_atmo',3),
    (100, 'persons', 'capteur_presence', 3),
    (119, 'level', 'niveau',3),
    (118, 'lux', 'capteur_luminosite',3),
    (128, 'kwh', 'consommation_elec',3),
    (101, 'heat', 'chaleur',3),
    (131, 'co2', 'capteur_co2',3),
    (114, 'humidity', 'capteur_humidite',3),
    (136, 'adc', 'analog_to_digital_converter',3),
    (112, 'temperature', 'temperature',3),
    (103, 'humidity', 'capteur_humidite',3),
    (115, 'motion', 'mouvement',3),
    (104, 'light', 'lumiere',3);

INSERT INTO users (firstname, lastname, email, password, group_name)
VALUES('Admin', 'Admin', 'admin@admin.fr', '$2a$10$hFZcsuSzOOgXNlPLVhY4WOnigHa0FQwVqUl9VG4UyHcYY9sg/faxO', 'administrator');

       INSERT INTO actuators(actuator_name, actuator_command, data_key , destination_address , room_id, sensor_id) 
VALUES ('HEATER_UP', '201', 'heat', '48edcb72-2940-4764-8a30-f1f266c9a697', 1 , 7),
       ('HEATER_DOWN', '202', 'heat','48edcb72-2940-4764-8a30-f1f266c9a697', 1 , 7),
       ('AC_UP', '203', 'ac','67d02c51-d07c-4588-a3cc-fa1c37e40480', 1, 1 ),
       ('AC_DOWN', '204', 'ac','67d02c51-d07c-4588-a3cc-fa1c37e40480', 1, 1),
       ('VENT_UP', '205', 'vent','d3470761-ebe4-4160-9c8e-0f64d241b57c', 1 , 11),
       ('VENT_DOWN', '206', 'vent','d3470761-ebe4-4160-9c8e-0f64d241b57c', 1 ,11),
       ('LIGHT_ON', '207', 'light','c73b859b-21ac-48b0-82f1-871e69d640a2', 1 , 14),
       ('LIGHT_OFF', '208', 'light','c73b859b-21ac-48b0-82f1-871e69d640a2', 1 , 14),
       ('HEATER_UP', '201','heat','54ecef7c-50a8-491d-9e62-ba61ce1639fa', 2 , 21),
       ('HEATER_DOWN', '202', 'heat','54ecef7c-50a8-491d-9e62-ba61ce1639fa', 2 , 21 ),
       ('AC_UP', '203', 'ac','a43e9218-379a-4e43-af55-5d66ba493b2e', 2 , 15),
       ('AC_DOWN', '204', 'ac','a43e9218-379a-4e43-af55-5d66ba493b2e', 2 , 15),
       ('VENT_UP', '205', 'vent','c7d34a93-e440-4a11-a745-9d7cb8f93e07', 2 , 25),
       ('VENT_DOWN', '206', 'vent','c7d34a93-e440-4a11-a745-9d7cb8f93e07', 2 , 25),
       ('LIGHT_ON', '207', 'light','3933a900-4f73-4716-9a97-42f5353c9254', 2 , 28),
       ('LIGHT_OFF', '208', 'light','3933a900-4f73-4716-9a97-42f5353c9254', 2 , 28),
       ('HEATER_UP', '201', 'heat','2543e375-1594-404f-b53c-5f3e38642a30', 3 , 35),
       ('HEATER_DOWN', '202', 'heat','2543e375-1594-404f-b53c-5f3e38642a30', 3 , 35),
       ('AC_UP', '203', 'ac','5424f807-56c6-406e-94af-f8ec8562522f', 3 , 29),
       ('AC_DOWN', '204', 'ac','5424f807-56c6-406e-94af-f8ec8562522f', 3 , 29),
       ('VENT_UP', '205', 'vent','cb124d1f-2e5d-4d15-a5f3-846682dded7d', 3 , 39),
       ('VENT_DOWN', '206', 'vent','cb124d1f-2e5d-4d15-a5f3-846682dded7d', 3 , 39),
       ('LIGHT_ON', '207', 'light','db0b2380-acf0-4688-b219-04ad29c369f3', 3 , 42),
       ('LIGHT_OFF', '208', 'light','db0b2380-acf0-4688-b219-04ad29c369f3', 3 , 42);