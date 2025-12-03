DROP TRIGGER IF EXISTS update_projects_updated_at ON dart.projects;

DROP FUNCTION IF EXISTS dart.update_updated_at_column();

DROP TABLE IF EXISTS dart.tasks;
DROP TABLE IF EXISTS dart.task_states;
DROP TABLE IF EXISTS dart.agents;
DROP TABLE IF EXISTS dart.resources;
DROP TABLE IF EXISTS dart.projects;

DROP SCHEMA IF EXISTS dart;