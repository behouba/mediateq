
CREATE TABLE IF NOT EXISTS "media" (
    "id" BIGSERIAL PRIMARY KEY,
    "hash" TEXT NOT NULL UNIQUE,
    "file_path" TEXT NOT NULL,
    "content_type" TEXT NOT NULL,
    "origin" TEXT NOT NULL,
    "url" TEXT NOT NULL UNIQUE,
    "timestamp" BIGINT NOT NULL,
    "size" BIGINT NOT NULL
)