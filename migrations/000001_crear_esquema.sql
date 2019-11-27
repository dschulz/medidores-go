

-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS medidor (
    codigo bigint NOT NULL,
    creado timestamp without time zone NOT NULL,
    creado_por character varying(255) NOT NULL,
    modificado_por character varying(255) NOT NULL,
    ultima_modificacion timestamp without time zone NOT NULL,
    uuid uuid NOT NULL,
    version bigint NOT NULL,
    numero integer NOT NULL,
    altitud double precision,
    latitud double precision,
    longitud double precision,
    organizacion_codigo bigint NOT NULL
);


CREATE TABLE IF NOT EXISTS lectura (
    codigo bigint NOT NULL,
    creado timestamp without time zone NOT NULL,
    creado_por character varying(255) NOT NULL,
    modificado_por character varying(255) NOT NULL,
    ultima_modificacion timestamp without time zone NOT NULL,
    uuid uuid NOT NULL,
    version bigint NOT NULL,
    contador integer,
    altitud double precision,
    latitud double precision,
    longitud double precision,
    medidor_codigo bigint NOT NULL
);


CREATE TABLE IF NOT EXISTS organizacion (
    codigo bigint NOT NULL,
    creado timestamp without time zone NOT NULL,
    creado_por character varying(255) NOT NULL,
    modificado_por character varying(255) NOT NULL,
    ultima_modificacion timestamp without time zone NOT NULL,
    uuid uuid NOT NULL,
    version bigint NOT NULL,
    email character varying(255),
    habilitado boolean NOT NULL,
    nombre character varying(50)
);

CREATE TABLE IF NOT EXISTS organizacion_usuario (
    codigo_organizacion bigint NOT NULL,
    codigo_usuario bigint NOT NULL
);


CREATE TABLE IF NOT EXISTS permiso (
    codigo bigint NOT NULL,
    creado timestamp without time zone NOT NULL,
    creado_por character varying(255) NOT NULL,
    modificado_por character varying(255) NOT NULL,
    ultima_modificacion timestamp without time zone NOT NULL,
    uuid uuid NOT NULL,
    version bigint NOT NULL,
    descripcion character varying(255),
    nombre character varying(255)
);


CREATE TABLE IF NOT EXISTS usuario (
    codigo bigint NOT NULL,
    creado timestamp without time zone NOT NULL,
    creado_por character varying(255) NOT NULL,
    modificado_por character varying(255) NOT NULL,
    ultima_modificacion timestamp without time zone NOT NULL,
    uuid uuid NOT NULL,
    version bigint NOT NULL,
    documento character varying(255),
    email character varying(255) NOT NULL,
    habilitado boolean NOT NULL,
    nombre character varying(50)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS medidor;
DROP TABLE IF EXISTS lectura;
DROP TABLE IF EXISTS organizacion_usuario;
DROP TABLE IF EXISTS usuario;
DROP TABLE IF EXISTS organizacion;
DROP TABLE IF EXISTS permiso;
