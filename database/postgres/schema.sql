
CREATE TABLE IF NOT EXISTS "media" (
    "nid" BIGSERIAL PRIMARY KEY,
    "id" TEXT NOT NULL UNIQUE,
    "content_type" TEXT NOT NULL,
    "origin" TEXT NOT NULL,
    "url" TEXT NOT NULL UNIQUE,
    "timestamp" BIGINT NOT NULL,
    "size_bytes" BIGINT NOT NULL
)