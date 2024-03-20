package controller

import (
	"bytes"
	"errors"
	"fmt"
	"image-processing-app/bootstrap"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type ResizeController struct {
	Env *bootstrap.Env
}

func (cc *ResizeController) Resize(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	filename := file.Filename
	err = c.SaveUploadedFile(file, "./tmp/"+filename)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	nameSplit := strings.Split(filename, ".")
	newFilename := nameSplit[0] + "-resized." + nameSplit[1]

	buffer := bytes.NewBuffer(nil)

	width := c.PostForm("width")
	if width == "" {
		width = "-1"
	}

	height := c.PostForm("height")
	if height == "" {
		height = "-1"
	}

	if width != "-1" || height != "-1" {
		err := ffmpeg.Input("./tmp/"+filename).
			Output("./tmp/"+newFilename, ffmpeg.KwArgs{"vf": fmt.Sprintf("scale=%s:%s", width, height)}).
			Run()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		byteArr, err := os.ReadFile("./tmp/" + newFilename)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		buffer = bytes.NewBuffer(byteArr)

		err = os.Remove("./tmp/" + filename)
		if err != nil {
			c.Error(err)
		}

		err = os.Remove("./tmp/" + newFilename)
		if err != nil {
			c.Error(err)
		}
	} else {
		c.Error(errors.New("width and height properties unsupplied, original image will be returned."))

		byteArr, err := os.ReadFile("./tmp/" + filename)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		buffer = bytes.NewBuffer(byteArr)

		err = os.Remove("./tmp/" + filename)
		if err != nil {
			c.Error(err)
		}
	}

	c.Header("Content-Disposition", "attachment; filename="+newFilename)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())
}
