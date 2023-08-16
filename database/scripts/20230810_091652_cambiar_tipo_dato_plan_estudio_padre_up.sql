ALTER TABLE plan_estudios.plan_estudio DROP CONSTRAINT IF EXISTS rel_plan_estudio_plan_estudio CASCADE;
ALTER TABLE plan_estudios.plan_estudio ALTER COLUMN plan_estudio_padre_id SET DATA TYPE bool USING (plan_estudio_padre_id::int::bool);
ALTER TABLE plan_estudios.plan_estudio RENAME COLUMN plan_estudio_padre_id TO es_plan_estudio_padre;