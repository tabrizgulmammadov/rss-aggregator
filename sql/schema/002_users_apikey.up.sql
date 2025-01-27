DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users' AND column_name = 'api_key'
    ) THEN
        ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
            encode(sha256(random()::text::bytea), 'hex')
        );
    END IF;
END $$;
