DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'feeds' AND column_name = 'last_fetched_at'
    ) THEN
        ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;
    END IF;
END $$;
