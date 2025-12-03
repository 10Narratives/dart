-- Create schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS code;

-- Create project functions
CREATE OR REPLACE FUNCTION code.create_project(
    p_name TEXT,
    p_display_name TEXT,
    p_description TEXT DEFAULT NULL
) RETURNS VOID AS $$
BEGIN
    INSERT INTO dart.projects (
        name,
        display_name,
        description,
        create_time,
        update_time
    )
    VALUES (
        p_name,
        p_display_name,
        p_description,
        NOW(),
        NOW()
    );
EXCEPTION
    WHEN unique_violation THEN
        RAISE EXCEPTION 'Project with name % already exists', p_name;
    WHEN check_violation THEN
        RAISE EXCEPTION 'Check constraint violation: %', SQLERRM;
    WHEN OTHERS THEN
        RAISE;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION code.update_project(
    p_name TEXT,
    p_display_name TEXT,
    p_description TEXT DEFAULT NULL
) RETURNS dart.projects AS $$
DECLARE
    updated_project dart.projects;
BEGIN
    UPDATE dart.projects
    SET
        display_name = p_display_name,
        description = p_description,
        update_time = NOW()
    WHERE name = p_name
    RETURNING * INTO updated_project;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Project with name % not found', p_name;
    END IF;

    RETURN updated_project;
EXCEPTION
    WHEN OTHERS THEN
        RAISE;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION code.get_project(p_name TEXT) RETURNS dart.projects AS $$
DECLARE
    project_record dart.projects;
BEGIN
    SELECT * INTO project_record
    FROM dart.projects
    WHERE name = p_name;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Project with name % not found', p_name;
    END IF;

    RETURN project_record;
EXCEPTION
    WHEN OTHERS THEN
        RAISE;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION code.list_projects(
    p_page_size INTEGER DEFAULT 50,
    p_page_token TEXT DEFAULT NULL
) RETURNS TABLE (
    name TEXT,
    display_name TEXT,
    description TEXT,
    create_time TIMESTAMP WITH TIME ZONE,
    update_time TIMESTAMP WITH TIME ZONE,
    next_page_token TEXT
) AS $$
DECLARE
    last_create_time TIMESTAMP WITH TIME ZONE;
    last_name TEXT;
    has_more BOOLEAN := FALSE;
    temp_record RECORD;
    i INTEGER := 0;
BEGIN
    -- Validate page size
    IF p_page_size <= 0 OR p_page_size > 1000 THEN
        p_page_size := 50;
    END IF;

    -- Decode page token if provided (format: base64_encoded_timestamp:name)
    IF p_page_token IS NOT NULL AND p_page_token != '' THEN
        BEGIN
            -- Decode base64 and split
            SELECT
                (regexp_split_to_array(
                    convert_from(decode(p_page_token, 'base64'), 'UTF8'),
                    ':'
                ))[1]::TIMESTAMP WITH TIME ZONE,
                (regexp_split_to_array(
                    convert_from(decode(p_page_token, 'base64'), 'UTF8'),
                    ':'
                ))[2]
            INTO last_create_time, last_name;
        EXCEPTION
            WHEN OTHERS THEN
                RAISE EXCEPTION 'Invalid page token format';
        END;
    END IF;

    FOR temp_record IN
        SELECT
            p.name,
            p.display_name,
            p.description,
            p.create_time,
            p.update_time
        FROM dart.projects p
        WHERE
            CASE
                WHEN p_page_token IS NOT NULL THEN (
                    p.create_time > last_create_time
                    OR (p.create_time = last_create_time AND p.name > last_name)
                )
                ELSE TRUE
            END
        ORDER BY p.create_time ASC, p.name ASC
        LIMIT (p_page_size + 1)
    LOOP
        i := i + 1;
        IF i <= p_page_size THEN
            name := temp_record.name;
            display_name := temp_record.display_name;
            description := temp_record.description;
            create_time := temp_record.create_time;
            update_time := temp_record.update_time;
            RETURN NEXT;
        ELSE
            has_more := TRUE;
        END IF;
    END LOOP;

    -- Set next_page_token if there are more records
    IF has_more THEN
        -- Get the last record that would have been returned
        SELECT
            p.name,
            p.create_time
        INTO last_name, last_create_time
        FROM dart.projects p
        WHERE
            CASE
                WHEN p_page_token IS NOT NULL THEN (
                    p.create_time > last_create_time
                    OR (p.create_time = last_create_time AND p.name > last_name)
                )
                ELSE TRUE
            END
        ORDER BY p.create_time ASC, p.name ASC
        LIMIT 1 OFFSET (p_page_size - 1);

        next_page_token := encode(
            convert_to(
                last_create_time::TEXT || ':' || last_name, 'UTF8'
            ), 'base64'
        );
    ELSE
        next_page_token := NULL;
    END IF;

    RETURN;
EXCEPTION
    WHEN OTHERS THEN
        RAISE;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION code.delete_project(p_name TEXT) RETURNS dart.projects AS $$
DECLARE
    deleted_project dart.projects;
BEGIN
    DELETE FROM dart.projects
    WHERE name = p_name
    RETURNING * INTO deleted_project;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Project with name % not found', p_name;
    END IF;

    RETURN deleted_project;
EXCEPTION
    WHEN OTHERS THEN
        RAISE;
END;
$$ LANGUAGE plpgsql;