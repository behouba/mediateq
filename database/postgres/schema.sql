
CREATE TABLE IF NOT EXISTS "media" (
    "id" BIGSERIAL PRIMARY KEY,
    "base64_hash" TEXT NOT NULL,
    "file_path" TEXT NOT NULL UNIQUE,
    "content_type" TEXT NOT NULL,
    "origin" TEXT NOT NULL,
    "url" TEXT NOT NULL UNIQUE,
    "timestamp" BIGINT NOT NULL,
    "size" BIGINT NOT NULL
);

CREATE INDEX idx_media_hash ON media ("base64_hash");
