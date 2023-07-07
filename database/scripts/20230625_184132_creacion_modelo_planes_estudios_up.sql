-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.0.4
-- PostgreSQL version: 14.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: new_database | type: DATABASE --
-- DROP DATABASE IF EXISTS new_database;
--CREATE DATABASE new_database;
-- ddl-end --


-- object: plan_estudios | type: SCHEMA --
-- DROP SCHEMA IF EXISTS plan_estudios CASCADE;
--CREATE SCHEMA plan_estudios;
-- ddl-end --
ALTER SCHEMA plan_estudios OWNER TO postgres;
-- ddl-end --

--SET search_path TO pg_catalog,public,plan_estudios;
-- ddl-end --

-- object: plan_estudios.plan_estudio | type: TABLE --
-- DROP TABLE IF EXISTS plan_estudios.plan_estudio CASCADE;
CREATE TABLE plan_estudios.plan_estudio (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	codigo_abreaviacion character varying(50),
	codigo character varying(10) NOT NULL,
	proyecto_academico_id integer NOT NULL,
	total_creditos numeric(6) NOT NULL,
	numero_semestres numeric(4) NOT NULL,
	numero_resolucion numeric(10) NOT NULL,
	ano_resolucion numeric(4) NOT NULL,
	espacios_semestre_distribucion json,
	resumen_plan_estudios json,
	soporte_documental json,
	observacion character varying(300),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	plan_estudio_padre_id integer,
	estado_aprobacion_id integer,
	CONSTRAINT plan_estudio_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.codigo IS E'Campo para el manejo del codigo del plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.codigo_abreaviacion IS E'Campo para el manejo del codigo de abreaviacion para el plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.proyecto_academico_id IS E'campo para la relacion del proyecto academico del plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.total_creditos IS E'Campo para el registro del total de creditos';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.numero_semestres IS E'Campo para almacenar el numero de semestres del plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.numero_resolucion IS E'Campo para almacenar el numero de la resolucion del plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.ano_resolucion IS E'Campo para el manejo del ano de la resolucion del plan de estuidos';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.espacios_semestre_distribucion IS E'Campo para almancenar la distribucion de los espacios academicos para el plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.resumen_plan_estudios IS E'Campo para almancenar el resumen del plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.soporte_documental IS E'Campo para la relacion de los documentos del plan de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.observacion IS E'Campo para el registro de las observaciones de los planes de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.activo IS E'campo para el estado del registro';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.fecha_creacion IS E'Campo para el registro de la fecha de creacion del registro';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.fecha_modificacion IS E'Campo para el registro de la fecha de modificacion del registro';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.plan_estudio_padre_id IS E'Campo para la relacione entre planes de estudios';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio.estado_aprobacion_id IS E'Campo para el registro del estado del plan de estudio';
-- ddl-end --
ALTER TABLE plan_estudios.plan_estudio OWNER TO postgres;
-- ddl-end --

-- object: plan_estudios.estado_aprobacion | type: TABLE --
-- DROP TABLE IF EXISTS plan_estudios.estado_aprobacion CASCADE;
CREATE TABLE plan_estudios.estado_aprobacion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(50),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT estado_aprobacion_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.id IS E'Campo serial';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.nombre IS E'Campo para el registro del nombre del estado';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.descripcion IS E'Campo para la descrición del registro';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.codigo_abreviacion IS E'Campo para el manejo del codigo de abreviacion';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.activo IS E'Campo para el estado del registro';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.fecha_creacion IS E'Campo para el registo de la fecha de creacion';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.estado_aprobacion.fecha_modificacion IS E'Campo para el registro de la fecha de modificacion';
-- ddl-end --
ALTER TABLE plan_estudios.estado_aprobacion OWNER TO postgres;
-- ddl-end --

-- object: plan_estudios.plan_estudio_proyecto_academico | type: TABLE --
-- DROP TABLE IF EXISTS plan_estudios.plan_estudio_proyecto_academico CASCADE;
CREATE TABLE plan_estudios.plan_estudio_proyecto_academico (
	id serial NOT NULL,
	plan_estudio_id integer NOT NULL,
	proyecto_academico_id integer NOT NULL,
	orden_proyecto json NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT plan_estudio_proyecto_academico_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio_proyecto_academico.plan_estudio_id IS E'Campo para la referecia del plan de estudio';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio_proyecto_academico.proyecto_academico_id IS E'campo para el registro del proyecto academico';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio_proyecto_academico.orden_proyecto IS E'Campos para el registro del orden los proyectos en un plan de estudio compartido';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio_proyecto_academico.activo IS E'Campo para el registro del estado';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio_proyecto_academico.fecha_creacion IS E'Campo para el registro de la fecha de creacion';
-- ddl-end --
COMMENT ON COLUMN plan_estudios.plan_estudio_proyecto_academico.fecha_modificacion IS E'Campo para el registro de la fecha de modificacion';
-- ddl-end --
ALTER TABLE plan_estudios.plan_estudio_proyecto_academico OWNER TO postgres;
-- ddl-end --

-- object: rel_estado_aprobacion_plan_estudio | type: CONSTRAINT --
-- ALTER TABLE plan_estudios.plan_estudio DROP CONSTRAINT IF EXISTS rel_estado_aprobacion_plan_estudio CASCADE;
ALTER TABLE plan_estudios.plan_estudio ADD CONSTRAINT rel_estado_aprobacion_plan_estudio FOREIGN KEY (estado_aprobacion_id)
REFERENCES plan_estudios.estado_aprobacion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: rel_plan_estudio_plan_estudio | type: CONSTRAINT --
-- ALTER TABLE plan_estudios.plan_estudio DROP CONSTRAINT IF EXISTS rel_plan_estudio_plan_estudio CASCADE;
ALTER TABLE plan_estudios.plan_estudio ADD CONSTRAINT rel_plan_estudio_plan_estudio FOREIGN KEY (plan_estudio_padre_id)
REFERENCES plan_estudios.plan_estudio (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: rel_plan_estudio_plan_estudio_proyecto_academico | type: CONSTRAINT --
-- ALTER TABLE plan_estudios.plan_estudio_proyecto_academico DROP CONSTRAINT IF EXISTS rel_plan_estudio_plan_estudio_proyecto_academico CASCADE;
ALTER TABLE plan_estudios.plan_estudio_proyecto_academico ADD CONSTRAINT rel_plan_estudio_plan_estudio_proyecto_academico FOREIGN KEY (plan_estudio_id)
REFERENCES plan_estudios.plan_estudio (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


