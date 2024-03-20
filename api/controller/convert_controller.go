package controller

import (
	"bytes"
	"errors"
	"image-processing-app/bootstrap"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type ConvertController struct {
	Env *bootstrap.Env
}

func (cc *ConvertController) Convert(c *gin.Context) {
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
	if nameSplit[1] != "png" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Only *.png image is accepted"))
		return
	}

	newFilename := nameSplit[0] + ".jpg"

	buffer := bytes.NewBuffer(nil)
	err = ffmpeg.Input("./tmp/" + filename).
		Output("./tmp/" + newFilename).
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

	c.Header("Content-Disposition", "attachment; filename="+newFilename)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())
}
