# Image Processing Service with Golang
A simple back-end service providing basic image processing functionalities implemented with Golang and Gin framework. 
<br><br>

## Main Dependencies Used
- [**gin**](https://gin-gonic.com) : A golang-based web framework to simplify web application development workflow
- [**ffmpeg-go**](https://github.com/u2takey/ffmpeg-go) : A golang based FFmpeg wrapper for image processing tasks
- [**testify**](https://github.com/stretchr/testify) : A testing library that provides various assertion functions and utilities for testing purposes
<br>

## Getting Started
### Requirements
Prior installation, the followings should be installed on your local machine :
  - Golang
    - Download the latest stable version from the [main site](https://go.dev/dl/)
    - Follow the installation [instructions](https://go.dev/doc/install) according to your platform
    - Verify the installation via terminal and make sure it outputs the installed version correctly :
  ```shell
  go version
  ```
  - FFmpeg
    - Download and install the latest build from the [main site](https://ffmpeg.org/download.html)
    - Verify the installation via terminal and make sure it outputs the installed version correctly :
  ```shell
  ffmpeg -version
  ```

### Installation
1. Clone this repository to your local directory and ``cd`` into the project dir
```shell
git clone https://github.com/hafizhH/image-processing-backend.git
cd image-processing-backend
```
2. Install all dependencies as listed in ``go.mod`` file
```shell
go mod download
```
3. Create a new ``.env`` file and specify the configuration of your service. The file should contain basic configuration similar to example below :  
```
APP_ENV=development
SERVER_ADDRESS=localhost:8080
CONTEXT_TIMEOUT=5
```
- ``APP_ENV`` : Switch between ``development`` and ``release`` environment
- ``SERVER_ADDRESS`` : The address and port on which the server will listen
- ``CONTEXT_TIMEOUT`` : Maximum allowed processing time before client request timed out 

### Run
Execute the command below to run the server with previously configured environment variables :
```shell
go run ./cmd/main.go
```

Execute the command below to run API integration test and package unit tests with test coverage report :
```shell
go test ./... -v -coverpkg=./...
```
<br>

## API Documentation
### 1. Convert Endpoint (``/convert``)
Upload a ``.png`` image as input, convert it to ``.jpg``, and return as output

- Request : ``POST /convert``  
```
Header :
  Content-Type : multipart/form-data

Body :
  image (form-file)
```

- Response :
```
Status : 200

Header :
  Content-Disposition : attachment; filename=[filename.ext]
  Content-Type : application/octet-stream

Body :
  image
```

### 2. Resize Endpoint (``/resize``)
Upload an image as input with additional width and/or height properties, resize the image to specified dimension, and return as output.  
If both of width and height properties supplied, the resulting image will be stretched to fit the dimension.
If only either one supplied, the image will be resized by maintaining the aspect ratio.

- Request : ``POST /resize``  
```
Header :
  Content-Type : multipart/form-data

Body :
  width (form-data)
  height (form-data)
  image (form-file)
```

- Response :
```
Status : 200

Header :
  Content-Disposition : attachment; filename=[filename.ext]
  Content-Type : application/octet-stream

Body :
  image
```

### 3. Compress Endpoint (``/compress``)
Upload an image as input, compress with specified quality parameter, and return the resulting image as output.
The quality parameter will be supplied as ffmpeg parameter.

- Request : ``POST /compress``  
```
Header :
  Content-Type : multipart/form-data

Body :
  quality (form-data)
  image (form-file)
```

- Response :
```
Status : 200

Header :
  Content-Disposition : attachment; filename=[filename.ext]
  Content-Type : application/octet-stream

Body :
  image
```

## Project Structure
```
├── go.mod
├── go.sum
├── .env
├── api
│   ├── controller
│   │   ├── compress_controller.go
│   │   ├── convert_controller.go
│   │   └── resize_controller.go
│   ├── route
│   │   ├── route.go
│   │   ├── compress_route.go
│   │   ├── convert_route.go
│   │   └── resize_route.go
│   └── middleware
│       └── error_handler_middleware.go
├── bootstrap
│   ├── app.go
│   ├── env.go
│   └── bootstrap_test.go
├── cmd
│   └── main.go
└── tests
    ├── init.go
    └── api_integration_test.go
```

## License
This project is licensed under the terms of the MIT license.
