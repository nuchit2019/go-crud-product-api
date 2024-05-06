DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'janawat_db') THEN
        CREATE DATABASE janawat_db;
    END IF;
END $$;
