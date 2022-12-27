
-- media_repository table store data about uploaded media files
CREATE TABLE IF NOT EXISTS "media_repository" (
    -- Unique identifier of the media 
    "media_id" TEXT NOT NULL UNIQUE,
    -- Mime type of the file
    "content_type" TEXT NOT NULL,
    -- Origin domain from which the file was uploaded
    "origin" TEXT NOT NULL,
    -- Public URL to download the media file
    "url" TEXT NOT NULL,
    -- base64 encoding string representation of a SHA-256 hash sum of the file data.
    "base64hash" TEXT NOT NULL,
    -- The UNIX timestamp when the file was uploaded
    "timestamp" BIGINT NOT NULL,
    -- The size of the file in bytes
    "size_bytes" BIGINT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS media_repository_index ON media_repository (media_id);


-- thumbnails table store data about generated thumbnails
CREATE TABLE IF NOT EXISTS "thumbnail" (
    -- Unique identifier of the media 
    "media_id" TEXT NOT NULL,
    -- Mime type of the file
    "content_type" TEXT NOT NULL,
    -- Origin domain from which the file was uploaded
    "origin" TEXT NOT NULL,
    -- Public URL to download the media file
    "url" TEXT NOT NULL,
    -- base64 encoding string representation of a SHA-256 hash sum of the file data.
    "base64hash" TEXT NOT NULL,
    -- The UNIX timestamp when the file was uploaded
    "timestamp" BIGINT NOT NULL,
    -- The size of the file in bytes
    "size_bytes" BIGINT NOT NULL,
    -- The width of the thumbnail
    "width" INTEGER NOT NULL,
    -- The height of the thumbnail
    "height" INTEGER NOT NULL,
    -- The resize method used to generate the thumbnail. Can be crop or scale.
    "crop" BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS thumbnail_index ON "thumbnail" (media_id, width, height, crop);
