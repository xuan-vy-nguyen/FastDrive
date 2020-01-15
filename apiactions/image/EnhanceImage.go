package image

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

// EnhancementImage is used for testing
func EnhancementImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enhancement Image")
	jwtStr := r.Header["AccessToken"][0]
	filename := r.Header["File-Name"][0]
	style := r.Header["Style"][0]

	// defer
	var returnImage []byte
	w.Header().Set("Content-Type", "multipart/form-data")
	defer func() {
		w.Write(returnImage)
		fmt.Println("")
	}()

	// check style
	if !(style == "mosaic" || style == "candy" || style == "rain_princess" || style == "udnie") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// find userMail
	UserInfor, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get image data
	file, _, err2 := r.FormFile("Image")
	if err2 != nil {
		fmt.Println(err2.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Read the file into memory
	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get nameImage
	nameImage := UserInfor.Mail + "_" + filename

	// write param to args.txt
	f, errWriteParams := os.OpenFile("args.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if errWriteParams != nil {
		fmt.Println(err)
		return
	}
	f.WriteString("image/input/" + nameImage + "\n")
	f.WriteString("image/output/" + nameImage + "\n")
	f.WriteString("fast_neural_style/saved_models/" + style + ".pth\n")
	f.Close()

	// write Image on disk
	img, _, errWriteonDisk := image.Decode(bytes.NewReader(data))
	if errWriteonDisk != nil {
		log.Fatalln(errWriteonDisk)
	}

	out, _ := os.Create("image/input/" + nameImage)

	var opts jpeg.Options
	opts.Quality = 1

	errEncodeImage := png.Encode(out, img)
	if errEncodeImage != nil {
		log.Println(err)
	}
	out.Close()

	// exec python
	cmd := exec.Command("python3.6", "fast_neural_style/neural_style/neural_style.py")
	if err := cmd.Start(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	// read image & respone
	infile, errReadImage := os.Open("image/output/" + nameImage)

	if errReadImage != nil {
		fmt.Println(errReadImage.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer infile.Close()

	// decode image
	src, errDecode := png.Decode(infile)
	if errDecode != nil {
		fmt.Println(errDecode.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	// convert to byte
	buf := new(bytes.Buffer)
	errCVByte := png.Encode(buf, src)
	if errCVByte != nil {
		fmt.Println(errCVByte.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	returnImage = buf.Bytes()

	// return ok
	w.WriteHeader(http.StatusCreated)
}
