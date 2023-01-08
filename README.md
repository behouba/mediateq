# mediateq

mediateq is a file storage REST API microservice that allows users to upload and download files, as well as resize images on the fly.

![CI](https://github.com/behouba/mediateq/actions/workflows/main.yml/badge.svg)

## Prerequisites

- libvips 8.3+
- PostgreSQL 14+

## Installation

To install mediateq, clone the project repository:

```bash
git clone https://github.com/behouba/mediateq.git
```

Mediateq depend on libvips.

Install libvips on Debian based Linux distributions:

```bash
sudo apt install -y libvips-dev
```

Run the following script from [bimg](https://github.com/h2non/bimg) as sudo (supports OSX, Debian/Ubuntu, Redhat, Fedora, Amazon Linux):

```bash
curl -s https://raw.githubusercontent.com/h2non/bimg/master/preinstall.sh | sudo bash -
```

The install script requires curl and pkg-config

If the above script is not working, please follow libvips installation instructions:

https://libvips.github.io/libvips/install.html

You can then build mediateq by running the `build.sh` script:

```bash
./build.sh
```

You will also need to setup a database. Currently mediateq only support postgreSQL.

Read the instructions to setup a postgreSQL database [here](database/postgres/README.md)

To run mediateq, use the following command:

```bash
./bin/mediateq -config=mediateq.yaml
```

You can also run tests for mediateq by using the run_tests.sh script:

```bash
./run_tests.sh
```

## API specification

The API specification of mediateq can be found [here](docs/mediateq-0.0.1.yaml)

All endpoints should be prefixed with: `/mediateq/{version}`

Note: mediateq uses base64 hash of files as filename to avoid duplications.

### /info (GET)

Retrieves information about the server.

Example Request

```
GET /info
```

Example Response

```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "version": "v0",
    "domain": "http://localhost:8080",
    "port": 8080,
    "database": "postgres",
    "storage": "localdisk",
    "allowedContentTypes": [
    "image/jpeg",
    "image/png",
    "image/bimg",
    "image/webp"
    ],
    "uptime": 35
}
```

### /upload (POST)

Uploads a new media file to the server.

Example Request

```
curl -X POST -H "Content-Type: multipart/form-data" -F "file=@herman.jpeg" http://localhost:8080/mediateq/v0/upload
```

Example Response

```
Copy code
HTTP/1.1 200 OK
Content-Type: application/json

{
    "media": {
        "id": "2K2dScz76xXsc0D3u29M3w1iKcw",
        "url": "http://localhost:8080/mediateq/v0/download/2K2dScz76xXsc0D3u29M3w1iKcw",
        "origin": "::1",
        "contentType": "image/jpeg",
        "sizeBytes": 103417,
        "tmestamp": 1673176612,
        "base64Hash": "zKZqFTxVdiIWGf-4otBjjBS46aPqs-Q6W0mefJmVhmo"
    }
}
```

### /download/{base64Hash} (GET)

Downloads a media file from the server.

Example Request

```
GET /download/2K2dScz76xXsc0D3u29M3w1iKcw?width=400
Example Response

Copy code
HTTP/1.1 200 OK
Content-Type: video/mp4
[binary data]
```

### /thumbnail/{base64Hash} (GET)

Downloads thumbnail version of images files.

The `width` and `height` of the image are required queries paramters.

Example Request

```

GET /download/2K2dScz76xXsc0D3u29M3w1iKcw?width=640&height=480
Example Response

Copy code
HTTP/1.1 200 OK
Content-Type: video/mp4
[binary data]

```

### /media (GET)

Retrieves a paginated list of all media files on the server.

Use query parameters `limit` and `offset` to paginate the results.

Example Request

GET /media?limit=4&offset=1

Example Response

```

HTTP/1.1 200 OK
Content-Type: application/json

{
    "mediaList": [
        {
            "id": "2JVvl82araY9MfG9O2keI1dSZDy",
            "url": "http://localhost:8080/mediateq/v0/download/2JVvl82araY9MfG9O2keI1dSZDy",
            "origin": "::1",
            "contentType": "image/jpeg",
            "sizeBytes": 103417,
            "tmestamp": 1672176212,
            "base64Hash": "zKZqFTxVdiIWGf-4otBjjBS46aPqs-Q6W0mefJmVhmo"
        },
        {
            "id": "2K2dScz76xXsc0D3u29M3w1iKcw",
            "url": "http://localhost:8080/mediateq/v0/download/2K2dScz76xXsc0D3u29M3w1iKcw",
            "origin": "::1",
            "contentType": "image/jpeg",
            "sizeBytes": 103417,
            "tmestamp": 1673176612,
            "base64Hash": "zKZqFTxVdiIWGf-4otBjjBS46aPqs-Q6W0mefJmVhmo"
        }
    ]
}

```

### /media/{base64Hash} (GET)

Retrieves information about a specific media file.

Example Request

```

GET /media/2K2dScz76xXsc0D3u29M3w1iKcw

```

Example Response

```

HTTP/1.1 200 OK
Content-Type: application/json

{
    "media": {
        "id": "2K2dScz76xXsc0D3u29M3w1iKcw",
        "url": "http://localhost:8080/mediateq/v0/download/2K2dScz76xXsc0D3u29M3w1iKcw",
        "origin": "::1",
        "contentType": "image/jpeg",
        "sizeBytes": 103417,
        "tmestamp": 1673176612,
        "base64Hash": "zKZqFTxVdiIWGf-4otBjjBS46aPqs-Q6W0mefJmVhmo"
    }
}

```

### /media/{base64Hash} (DELETE)

Delete a specific media file by it base64 hash identifier.

Example Request

```

DELETE /media/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY

```

Example Response

```

HTTP/1.1 200 OK
Content-Type: application/json

{
    "message": "media 4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY deleted"
}

```

### Docker container

[TODO]
