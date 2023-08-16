ALTER TABLE plan_estudios.plan_estudio ALTER COLUMN es_plan_estudio_padre SET DATA TYPE integer;
ALTER TABLE plan_estudios.plan_estudio RENAME COLUMN es_plan_estudio_padre TO plan_estudio_padre_id;
ALTER TABLE plan_estudios.plan_estudio ADD CONSTRAINT rel_plan_estudio_plan_estudio FOREIGN KEY (plan_estudio_padre_id)
    REFERENCES plan_estudios.plan_estudio (id) MATCH SIMPLE
    ON DELETE NO ACTION ON UPDATE NO ACTION;
