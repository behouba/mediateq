
CREATE TABLE IF NOT EXISTS "media" (
    "id" BIGSERIAL PRIMARY KEY,
    "hash" TEXT NOT NULL,
    "full_path" TEXT NOT NULL,
    "content_type" TEXT NOT NULL,
    "origin" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "timestamp" BIGINT NOT NULL,
    "size" BIGINT NOT NULL
)