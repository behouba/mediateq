
-- media_repository table store data about uploaded media files
CREATE TABLE IF NOT EXISTS "media_repository" (
    -- Unique identifier of the media 
    -- We use the md5 hash of the file as id to ensure integrity and uniquness of files
    "media_id" TEXT NOT NULL UNIQUE,
    -- Mime type of the file
    "content_type" TEXT NOT NULL,
    -- Origin domain from which the file was uploaded
    "origin" TEXT NOT NULL,
    -- Public URL to download the media file
    "url" TEXT NOT NULL UNIQUE,
    -- base64 encoding string representation of a SHA-256 hash sum of the file data.
    "base64hash" TEXT NOT NULL,
    -- The UNIX timestamp when the file was uploaded
    "timestamp" BIGINT NOT NULL,
    -- The size of the file in bytes
    "size_bytes" BIGINT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS media_repository_index ON media_repository (media_id);
