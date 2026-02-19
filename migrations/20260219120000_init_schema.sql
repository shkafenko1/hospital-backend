-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.patient
(
    id bigserial NOT NULL,
    passport character varying(9) NOT NULL,
    first_name character varying(20) NOT NULL,
    middle_name character varying(20) NOT NULL,
    last_name character varying(20) NOT NULL,
    phone character varying(13) NOT NULL,
    birth_date date NOT NULL,
    sex character varying,
    address character varying(50) NOT NULL,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS public.doctor
(
    id bigserial NOT NULL,
    first_name character varying(20) NOT NULL,
    middle_name character varying(20) NOT NULL,
    last_name character varying(20) NOT NULL,
    specialization character varying(20) NOT NULL,
    office bigint NOT NULL,
    license_mumber bigint NOT NULL,
    phone character varying(13) NOT NULL,
    experience bigint NOT NULL,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS public.ward
(
    id bigserial NOT NULL,
    "number" bigint NOT NULL,
    floor bigint NOT NULL,
    capacity bigint NOT NULL,
    price bigint,
    comfort character varying(20) NOT NULL,
    busy boolean NOT NULL,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS public.medicine
(
    id bigserial NOT NULL,
    article bigint NOT NULL,
    name character varying(30) NOT NULL,
    active_substance character varying(30) NOT NULL,
    form character varying(30) NOT NULL,
    contraindications character varying(50) NOT NULL,
    prescription boolean NOT NULL,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS public.hospitalization
(
    id serial NOT NULL,
    type character varying(20) NOT NULL,
    status character varying(20) NOT NULL,
    beginning_date date NOT NULL,
    primary_diagnosis character varying(100) NOT NULL,
    discharge_date date,
    discharge_diagnosis character varying(100),
    patient_id bigint NOT NULL,
    doctor_id bigint NOT NULL,
    ward_id bigint NOT NULL,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS public.visit
(
    id bigserial NOT NULL,
    date_time timestamp with time zone NOT NULL,
    symptoms character varying(100) NOT NULL,
    diagnosis character varying(100) NOT NULL,
    advise character varying(100) NOT NULL,
    price numeric,
    patient_id bigint NOT NULL,
    doctor_id bigint NOT NULL,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS public.hospitalization_medicine
(
    id bigserial NOT NULL,
    hospitalization_id integer NOT NULL,
    medicine_id bigint NOT NULL,
    PRIMARY KEY (id)
    );

-- Foreign Keys
ALTER TABLE IF EXISTS public.hospitalization
    ADD CONSTRAINT patient_fk FOREIGN KEY (patient_id)
    REFERENCES public.patient (id);

ALTER TABLE IF EXISTS public.hospitalization
    ADD CONSTRAINT doctor_fk FOREIGN KEY (doctor_id)
    REFERENCES public.doctor (id);

ALTER TABLE IF EXISTS public.hospitalization
    ADD CONSTRAINT ward_fk FOREIGN KEY (ward_id)
    REFERENCES public.ward (id);

ALTER TABLE IF EXISTS public.visit
    ADD CONSTRAINT doctor_fk FOREIGN KEY (doctor_id)
    REFERENCES public.doctor (id);

ALTER TABLE IF EXISTS public.visit
    ADD CONSTRAINT patient_fk FOREIGN KEY (patient_id)
    REFERENCES public.patient (id);

ALTER TABLE IF EXISTS public.hospitalization_medicine
    ADD CONSTRAINT hospitalization_fk FOREIGN KEY (hospitalization_id)
    REFERENCES public.hospitalization (id);

ALTER TABLE IF EXISTS public.hospitalization_medicine
    ADD CONSTRAINT medicine_fk FOREIGN KEY (medicine_id)
    REFERENCES public.medicine (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Удаление таблиц в обратном порядке зависимостей
DROP TABLE IF EXISTS public.hospitalization_medicine;
DROP TABLE IF EXISTS public.visit;
DROP TABLE IF EXISTS public.hospitalization;
DROP TABLE IF EXISTS public.medicine;
DROP TABLE IF EXISTS public.ward;
DROP TABLE IF EXISTS public.doctor;
DROP TABLE IF EXISTS public.patient;
-- +goose StatementEnd