package tests

import (
	"bytes"
	"image-processing-backend/api/route"
	"image-processing-backend/bootstrap"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConvertEndpoint(t *testing.T) {
	router := SetupEnvironment()

	resRecorder := httptest.NewRecorder()

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	LoadFileToMultipart("./sample/", "test.png", buffer, writer)

	writer.Close()

	req, _ := http.NewRequest("POST", "/convert", buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	router.ServeHTTP(resRecorder, req)

	assert.Equal(t, 200, resRecorder.Code, "Response status code should be 200")
	assert.Equal(t, "attachment; filename=test.jpg", resRecorder.Result().Header["Content-Disposition"][0], "Response header should contain output filename")
	assert.NotNil(t, resRecorder.Result().Body, "Response body should not be empty")
}

func TestResizeEndpoint(t *testing.T) {
	router := SetupEnvironment()

	resRecorder := httptest.NewRecorder()

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	LoadFileToMultipart("./sample/", "test.png", buffer, writer)
	AppendMultipartField("width", "512", buffer, writer)
	AppendMultipartField("height", "128", buffer, writer)

	writer.Close()

	req, _ := http.NewRequest("POST", "/resize", buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	router.ServeHTTP(resRecorder, req)

	assert.Equal(t, 200, resRecorder.Code, "Response status code should be 200")
	assert.Equal(t, "attachment; filename=test-resized.png", resRecorder.Result().Header["Content-Disposition"][0], "Response header should contain output filename")
	assert.NotNil(t, resRecorder.Result().Body, "Response body should not be empty")
}

func TestCompressEndpoint(t *testing.T) {
	router := SetupEnvironment()

	resRecorder := httptest.NewRecorder()

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	LoadFileToMultipart("./sample/", "test.jpg", buffer, writer)
	AppendMultipartField("quality", "5", buffer, writer)

	writer.Close()

	req, _ := http.NewRequest("POST", "/compress", buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	router.ServeHTTP(resRecorder, req)

	assert.Equal(t, 200, resRecorder.Code, "Response status code should be 200")
	assert.Equal(t, "attachment; filename=test-compressed.jpg", resRecorder.Result().Header["Content-Disposition"][0], "Response header should contain output filename")
	assert.NotNil(t, resRecorder.Result().Body, "Response body should not be empty")
}

func SetupEnvironment() *gin.Engine {
	// env := bootstrap.NewEnv("../../.env")
	env, _ := bootstrap.NewEnv()
	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin.SetMode("release")
	router := gin.New()
	router.Use(gin.Recovery())
	route.SetupRouter(env, timeout, router)

	return router
}

func LoadFileToMultipart(path string, filename string, buffer *bytes.Buffer, writer *multipart.Writer) {
	byteArr, _ := os.ReadFile(path + filename)
	reader := bytes.NewReader(byteArr)

	part, _ := writer.CreateFormFile("image", filename)
	io.Copy(part, reader)
}

func AppendMultipartField(key string, value string, buffer *bytes.Buffer, writer *multipart.Writer) {
	reader := bytes.NewReader([]byte(value))

	part, _ := writer.CreateFormField(key)
	io.Copy(part, reader)
}
