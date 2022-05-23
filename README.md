# stash
Stash is a file storage REST API server. 

## Dependencies
- libvips 8.3+

## Installation

### Install libvips
for mac:

```bash
brew install vips
```

for Linux (Ubuntu):

```bash
sudo apt-get update -y
```

```bash
sudo apt-get install -y libvips
```

For other operating systems check [libvips install](https://libvips.github.io/libvips/install.html) instructions


### build stashd

To build stashd you can run the build.sh script
```bash
./build.sh
```

### Run stashd

```bash
./bin/stashd
```
