DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users' AND column_name = 'api_key'
    ) THEN
        ALTER TABLE users DROP COLUMN api_key;
    END IF;
END $$;
