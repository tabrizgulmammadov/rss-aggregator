DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'feeds' AND column_name = 'last_fetched_at'
    ) THEN
        ALTER TABLE feeds DROP COLUMN last_fetched_at;
    END IF;
END $$;
