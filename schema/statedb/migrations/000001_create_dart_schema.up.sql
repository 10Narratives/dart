create schema dart;

CREATE TABLE
  dart.projects (
    name TEXT PRIMARY KEY,
    display_name TEXT NOT NULL,
    description TEXT,
    create_time TIMESTAMP
    WITH
      TIME ZONE DEFAULT NOW (),
      update_time TIMESTAMP
    WITH
      TIME ZONE DEFAULT NOW ()
  );

CREATE TABLE
  dart.resources (
    name TEXT PRIMARY KEY,
    compute_capabilities JSONB NOT NULL,
    memory_capabilities JSONB NOT NULL,
    network_capabilities JSONB NOT NULL,
    task_execution_capabilities JSONB NOT NULL,
    create_time TIMESTAMP
    WITH
      TIME ZONE DEFAULT NOW (),
      address TEXT NOT NULL
  );

CREATE TABLE
  dart.agents (
    name TEXT PRIMARY KEY,
    compute_capabilities JSONB NOT NULL,
    memory_capabilities JSONB NOT NULL,
    network_capabilities JSONB NOT NULL,
    task_execution_capabilities JSONB NOT NULL,
    parameters_schema TEXT,
    parameters_descriptions JSONB,
    create_time TIMESTAMP
    WITH
      TIME ZONE DEFAULT NOW ()
  );

CREATE TABLE
  dart.task_states (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
  );

INSERT INTO
  dart.task_states (id, name, description)
VALUES
  (
    1,
    'PENDING',
    'Task is queued but not yet started'
  ),
  (2, 'RUNNING', 'Task is currently executing'),
  (3, 'FAILED', 'Task execution failed'),
  (
    4,
    'EXPIRED',
    'Task exceeded its TTL before execution'
  ),
  (5, 'COMPLETED', 'Task completed successfully'),
  (6, 'CANCELED', 'Task was explicitly canceled');

CREATE TABLE
  dart.tasks (
    name TEXT PRIMARY KEY,
    assigned_resource TEXT REFERENCES dart.resources (name) ON DELETE SET NULL,
    ttl_seconds BIGINT,
    create_time TIMESTAMP
    WITH
      TIME ZONE DEFAULT NOW (),
      start_time TIMESTAMP
    WITH
      TIME ZONE,
      end_time TIMESTAMP
    WITH
      TIME ZONE,
      parameters TEXT NOT NULL,
      state_id BIGINT NOT NULL REFERENCES dart.task_states (id) DEFAULT 1
  );