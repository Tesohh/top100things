package form

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Tesohh/top100things/auth"
	"github.com/Tesohh/top100things/model"
	"github.com/Tesohh/top100things/sb"
	"github.com/nfnt/resize"
)

func resizeImage(r io.Reader) ([]byte, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	resizedImg := resize.Thumbnail(100, 100, img, resize.Lanczos2)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resizedImg, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func encodeBase64(img []byte, header *multipart.FileHeader) string {
	filetype := header.Header.Get("Content-Type")
	b64 := fmt.Sprintf("data:%s;base64,", filetype)
	b64 += base64.StdEncoding.EncodeToString(img)

	return b64
}

func processImage(r *http.Request) (string, error) {
	rawImage, header, err := r.FormFile("image")
	if err != nil {
		return "", err
	}
	defer rawImage.Close()

	img, err := resizeImage(rawImage)
	if err != nil {
		return "", err
	}

	return encodeBase64(img, header), nil
}

func NewThing(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	desc := r.FormValue("description")
	img, err := processImage(r)
	if err != nil {
		fmt.Fprint(w, err)
	}

	user := auth.User(r)
	if user == nil {
		fmt.Fprint(w, "you need to be logged in!")
		return
	}

	newthing := map[string]any{
		"name":      title,
		"image_url": img,
		"author":    user.ID,
	}
	var newThings []model.Thing
	err = sb.SB.DB.From("things").Insert(newthing).Execute(&newThings)
	if err != nil {
		fmt.Fprint(w, "thing: "+err.Error())
		return
	}

	var myRankings []model.Ranking
	sb.SB.DB.From("rankings").Select("user_id").Eq("user_id", user.ID).Execute(&myRankings)

	newranking := map[string]any{
		"user_id":     user.ID,
		"thing_id":    newThings[0].ID,
		"position":    len(myRankings) + 1,
		"description": desc,
	}
	iDontCare := make([]map[string]any, 0)
	err = sb.SB.DB.From("rankings").Insert(newranking).Execute(&iDontCare)
	if err != nil {
		fmt.Fprint(w, "ranking: "+err.Error())
	}

	http.Redirect(w, r, "/things/"+user.ID, http.StatusSeeOther)

}
