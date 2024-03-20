package controller

import (
	"bytes"
	"errors"
	"image-processing-app/bootstrap"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type CompressController struct {
	Env *bootstrap.Env
}

func (cc *CompressController) Compress(c *gin.Context) {
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

	name_split := strings.Split(filename, ".")
	newFilename := name_split[0] + "-compressed." + name_split[1]

	buffer := bytes.NewBuffer(nil)

	quality := 4
	input_quality := c.PostForm("quality")
	if input_quality == "" {
		c.Error(errors.New("quality property unsupplied, using default value"))
	} else {
		quality, err = strconv.Atoi(input_quality)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	quality = max(0, min(31, quality))

	err = ffmpeg.Input("./tmp/"+filename).
		Output("./tmp/"+newFilename, ffmpeg.KwArgs{"qscale:v": strconv.Itoa(quality)}).
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
