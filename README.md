# mediateq

mediateq is a file storage REST API microservice that allows users to upload and download files, as well as resize images on the fly.

![CI](https://github.com/behouba/mediateq/actions/workflows/main.yml/badge.svg)

## Prerequisites

- libvips 8.3+ (8.8+ recommended)
- PostgreSQL 14+

## Installation

To install mediateq, clone the project repository:

```bash
git clone https://github.com/behouba/mediateq.git
```

Mediateq depend on libvips. Run the following script from [bimg](https://github.com/h2non/bimg) as sudo (supports OSX, Debian/Ubuntu, Redhat, Fedora, Amazon Linux):

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

### /info (GET)

Retrieves information about the server.

Example Request

```
GET /info
```

Example Response

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "version": "v0",
    "domain": "http://localhost:8080",
    "port": 8080,
    "startTime": "2022-12-25T14:07:54.6703027+03:00",
    "allowedContentTypes": [
    "image/jpeg",
    "image/png",
    "image/bimg",
    "image/webp"
    ]
}
```

### /upload (POST)

Uploads a new media file to the server.

Example Request

```
curl -X POST -H "Content-Type: multipart/form-data" -F "file=@herman.jpeg" http://localhost:8080/mediateq/v0/upload
```

Example Response

```json
Copy code
HTTP/1.1 200 OK
Content-Type: application/json

{
    "media": {
        "id": 3,
        "base64Hash": "RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo",
        "url": "http://localhost:8080/mediateq/v0/download/RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo",
        "filePath": "upload/2022/12/RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo",
        "origin": "::1",
        "contentType": "image/jpeg",
        "size": 235807,
        "tmestamp": 1671898018
    }
}
```

### /download/{mediaId} (GET)

Downloads a media file from the server.

You can use `width` query parameter to get resized version for images files.

Example Request

```
GET /download/RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo?width=400
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
        "id": 1,
        "base64Hash": "4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
        "url": "http://localhost:8080/mediateq/v0/download/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
        "filePath": "upload/2022/12/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
        "origin": "::1",
        "contentType": "image/jpeg",
        "size": 92998,
        "tmestamp": 1671895940
    },
    {
        "id": 2,
        "base64Hash": "OSxgpSK--4-S9UlcQS9HqlJxQKssGppuawP57HTXKX8",
        "url": "http://localhost:8080/mediateq/v0/download/OSxgpSK--4-S9UlcQS9HqlJxQKssGppuawP57HTXKX8",
        "filePath": "upload/2022/12/OSxgpSK--4-S9UlcQS9HqlJxQKssGppuawP57HTXKX8",
        "origin": "::1",
        "contentType": "image/jpeg",
        "size": 112816,
        "tmestamp": 1671896197
    },
    {
        "id": 3,
        "base64Hash": "RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo",
        "url": "http://localhost:8080/mediateq/v0/download/RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo",
        "filePath": "upload/2022/12/RhoW8IwG1qO_nP_yua5VfYbdSI_wiNAZ2BWcsogMAVo",
        "origin": "::1",
        "contentType": "image/jpeg",
        "size": 235807,
        "tmestamp": 1671898018
    },
    {
        "id": 4,
        "base64Hash": "_n_0rkoX659AMNyTAlnoyOfI3nvpSZMjn57GNI3A2oE",
        "url": "http://localhost:8080/mediateq/v0/download/_n_0rkoX659AMNyTAlnoyOfI3nvpSZMjn57GNI3A2oE",
        "filePath": "upload/2022/12/_n_0rkoX659AMNyTAlnoyOfI3nvpSZMjn57GNI3A2oE",
        "origin": "::1",
        "contentType": "image/png",
        "size": 747191,
        "tmestamp": 1671906045
    },
    {
        "id": 5,
        "base64Hash": "wGNrFrL53Nr6pNlIyF4q57jquUAkKiCK3q_WwSt52II",
        "url": "http://localhost:8080/mediateq/v0/download/wGNrFrL53Nr6pNlIyF4q57jquUAkKiCK3q_WwSt52II",
        "filePath": "upload/2022/12/wGNrFrL53Nr6pNlIyF4q57jquUAkKiCK3q_WwSt52II",
        "origin": "::1",
        "contentType": "image/jpeg",
        "size": 237301,
        "tmestamp": 1671906079
    }
]
}
```

### /media/{mediaId} (GET)

Retrieves information about a specific media file.

Example Request

```

GET /media/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY
```

Example Response

```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "media": {
        "id": 1,
        "base64Hash": "4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
        "url": "http://localhost:8080/mediateq/v0/download/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
        "filePath": "upload/2022/12/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
        "origin": "::1",
        "contentType": "image/jpeg",
        "size": 92998,
        "tmestamp": 1671895940
    }
}
```

### /media/{mediaId} (DELETE)

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

TODO
