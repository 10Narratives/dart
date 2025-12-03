-- Create schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS dart;

-- Create projects table
CREATE TABLE IF NOT EXISTS dart.projects (
    name TEXT PRIMARY KEY,
    display_name TEXT NOT NULL,
    description TEXT,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    update_time TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create resources table
CREATE TABLE IF NOT EXISTS dart.resources (
    name TEXT PRIMARY KEY,
    compute_capabilities JSONB NOT NULL DEFAULT '{}',
    memory_capabilities JSONB NOT NULL DEFAULT '{}',
    network_capabilities JSONB NOT NULL DEFAULT '{}',
    task_execution_capabilities JSONB NOT NULL DEFAULT '{}',
    address TEXT NOT NULL,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create agents table
CREATE TABLE IF NOT EXISTS dart.agents (
    name TEXT PRIMARY KEY,
    compute_capabilities JSONB NOT NULL DEFAULT '{}',
    memory_capabilities JSONB NOT NULL DEFAULT '{}',
    network_capabilities JSONB NOT NULL DEFAULT '{}',
    task_execution_capabilities JSONB NOT NULL DEFAULT '{}',
    parameters_schema TEXT,
    parameters_descriptions JSONB,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create task states table with constraints
CREATE TABLE IF NOT EXISTS dart.task_states (
    id INTEGER PRIMARY KEY CHECK (id >= 1 AND id <= 999),
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

-- Insert task states with ON CONFLICT to handle re-runs
INSERT INTO dart.task_states (id, name, description)
VALUES
    (1, 'PENDING', 'Task is queued but not yet started'),
    (2, 'RUNNING', 'Task is currently executing'),
    (3, 'FAILED', 'Task execution failed'),
    (4, 'EXPIRED', 'Task exceeded its TTL before execution'),
    (5, 'COMPLETED', 'Task completed successfully'),
    (6, 'CANCELED', 'Task was explicitly canceled')
ON CONFLICT (id) DO NOTHING;

-- Create tasks table
CREATE TABLE IF NOT EXISTS dart.tasks (
    name TEXT PRIMARY KEY,
    assigned_resource TEXT REFERENCES dart.resources (name) ON DELETE SET NULL,
    ttl_seconds BIGINT,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    parameters TEXT NOT NULL,
    state_id INTEGER NOT NULL REFERENCES dart.task_states (id) DEFAULT 1,
    CHECK (end_time IS NULL OR start_time IS NOT NULL),
    CHECK (end_time IS NULL OR end_time >= start_time)
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_tasks_state ON dart.tasks (state_id);
CREATE INDEX IF NOT EXISTS idx_tasks_create_time ON dart.tasks (create_time);
CREATE INDEX IF NOT EXISTS idx_tasks_assigned_resource ON dart.tasks (assigned_resource);
CREATE INDEX IF NOT EXISTS idx_projects_create_time ON dart.projects (create_time);

-- Add comments for documentation
COMMENT ON TABLE dart.projects IS 'Project entities in the system';
COMMENT ON TABLE dart.resources IS 'Resource entities with capability information';
COMMENT ON TABLE dart.agents IS 'Agent entities with execution capabilities';
COMMENT ON TABLE dart.task_states IS 'Valid states for tasks';
COMMENT ON TABLE dart.tasks IS 'Task execution records';

-- Create update trigger function with proper syntax
CREATE OR REPLACE FUNCTION dart.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger with proper syntax
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger 
        WHERE tgname = 'update_projects_updated_at' 
        AND tgrelid = 'dart.projects'::regclass
    ) THEN
        CREATE TRIGGER update_projects_updated_at 
            BEFORE UPDATE ON dart.projects 
            FOR EACH ROW 
            EXECUTE FUNCTION dart.update_updated_at_column();
    END IF;
END $$;