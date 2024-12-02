package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"fmt"
	"os"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

// ----------------------------------------------------------------
func main() {
	fmt.Println("Video Web Server")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	xip := fmt.Sprintf("%s", GetOutboundIP())
	port := "8080"
	exefile := ""
	exefilea := ""

	drive := "c"
	wdir := "/tmp/"
	switch runtime.GOOS {
	case "windows":
		exefile = "/ffmpeg/bin/ffmpeg.exe"
		exefilea = "/ffmpeg/bin/ffprobe.exe"
		wdir = "/dwhelper/"

	case "linux":
		exefile = "ffmpeg"
		exefilea = "ffprobe"
		wdir = "/media/dave/Elements/dwhelper/"

	}

	pgsize := 10
	maxsel := 1000
	display := 0
	subdir := true
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:

		fmt.Println("Not")

		//-------------------------------------------------------------
	default:

		fmt.Println("Server running....")
		fmt.Println("Listening on " + xip + ":" + port)

		fmt.Println("")
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			xdata := InitPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page Handler
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := AboutPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------Dymnamic Display Page Handler
		http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
			page := r.URL.Query().Get("page")
			sdir := r.URL.Query().Get("sdir")
			xdata := DisplayPage(subdir, xip, port, page, sdir, exefile, exefilea, drive, wdir, pgsize, maxsel, display)
			fmt.Fprint(w, xdata)

		})

		//------------------------------------------------Play Video Page Handler
		http.HandleFunc("/playvideo", func(w http.ResponseWriter, r *http.Request) {
			video := r.URL.Query().Get("video")
			sdir := r.URL.Query().Get("sdir")
			xdata := PlayVideoPage(xip, port, video, exefile, exefilea, drive, wdir, sdir)
			fmt.Fprint(w, xdata)

		})

		//------------------------------------------------Move Video Page Handler
		http.HandleFunc("/movevideo", func(w http.ResponseWriter, r *http.Request) {
			video := r.URL.Query().Get("video")
			sdir := r.URL.Query().Get("sdir")
			xdata := MoveVideoPage(xip, port, video, exefile, exefilea, drive, wdir, sdir, subdir)
			fmt.Fprint(w, xdata)

		})

		//------------------------------------------------Move Video Page Handler
		http.HandleFunc("/movevideocomplete", func(w http.ResponseWriter, r *http.Request) {
			video := r.URL.Query().Get("video")
			sdir := r.URL.Query().Get("sdir")
			ddir := r.URL.Query().Get("ddir")
			xdata := MoveVideoCompletePage(xip, port, video, exefile, exefilea, drive, wdir, sdir, ddir)
			fmt.Fprint(w, xdata)

		})

		//------------------------------------------------ Tag Video Page Handler
		http.HandleFunc("/tagvideo", func(w http.ResponseWriter, r *http.Request) {
			video := r.URL.Query().Get("video")
			sdir := r.URL.Query().Get("sdir")
			xdata := TagVideoPage(xip, port, video, exefile, exefilea, drive, wdir, sdir)
			fmt.Fprint(w, xdata)

		})

		//------------------------------------------------- Static Handler Handler
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
		//------------------------------------------------- Start Server
		Openbrowser(xip + ":" + port)
		if err := http.ListenAndServe(xip+":"+port, nil); err != nil {
			panic(err)
		}
	}
}

// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start msedge"}

	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func DateTimeDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm' : 'am';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtdt').innerHTML = day+', '+month+' '+dm+', '+yr+' - '+h + ':' + m + ':' + s+' '+ampm;"

	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}
func LoopDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startLoop() {"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtloop').innerHTML = Math.random();"
	xdata = xdata + "  var t = setTimeout(startLoop, 500);"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func fixFileName(fileName string) string {
	newName := ""
	chr := ""
	ascval := 0

	for x := 0; x < len(fileName); x++ {
		chr = fileName[x : x+1]
		ascval = asciistring.StringToASCII(chr)
		switch {
		case ascval < 45:
		case ascval == 64:
		case ascval == 92:
		case ascval == 96:
		case ascval > 122:
		default:
			newName = newName + chr
		}
	}

	err := os.Rename(fileName, newName)
	if err != nil {
		fmt.Println("Error renaming file:", err)
	}
	return newName
}

func ValidFileType(fileExt string) bool {
	rtn := false
	switch {
	case fileExt == ".mp4":
		rtn = true
	case fileExt == ".avi":
		rtn = true
	case fileExt == ".wmv":
		rtn = true
	case fileExt == ".asf":
		rtn = true
	}
	return rtn
}

func ParseFrameRate(data string) string {
	rtn := ""
	chr := ""
	do := false
	pass := 1
	v1 := ""
	v2 := ""
	add := true
	ascval := 0
	for x := 0; x < len(data); x++ {
		chr = data[x : x+1]
		add = true
		ascval = asciistring.StringToASCII(chr)
		if ascval == 13 {
			add = false
			numerator, _ := strconv.Atoi(v1)
			denominator, _ := strconv.Atoi(v2)
			if denominator > 0 {
				fps := numerator / denominator
				rtn = strconv.Itoa(fps)
			}
		}
		if ascval == 10 {
			add = false
			pass = 1
			do = false
		}
		if chr == "," {
			do = true
			add = false
		}
		if chr == "/" {
			pass = 2
			add = false
		}
		if do {
			if add {
				if pass == 1 {
					v1 = v1 + chr
				}
				if pass == 2 {
					v2 = v2 + chr
				}
			}
		}
	}
	return rtn
}

func ParseBitRate(data string) string {
	rtn := ""
	chr := ""
	do := false
	add := true
	pass := 1
	ascval := 0
	for x := 0; x < len(data); x++ {
		chr = data[x : x+1]
		add = true
		ascval = asciistring.StringToASCII(chr)
		if ascval == 13 {
			add = false
		}
		if ascval == 10 {
			add = false
			do = false
			pass = 2
		}
		if chr == "," {
			do = true
			add = false
		}
		if do {
			if add {
				if pass == 2 {
					rtn = rtn + chr
				}
			}
		}
	}
	return rtn
}

func FileData(exefilea string, tnfile string, fileName string) string {
	xdata := ""
	bfile := ""
	switch runtime.GOOS {
	case "windows":
		bfile = "tmp.bat"
	case "linux":
		bfile = "tmp.sh"

	}

	bdata := []byte(exefilea + " -i " + tnfile + " -show_entries stream=width,height -of csv=" + fmt.Sprintf("%q", "p=0") + ">tmp.csv")
	err := os.WriteFile(bfile, bdata, 0777)
	if runtime.GOOS == "linux" {
		bfile = "./" + bfile
	}
	cmd := exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat := []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata := string(dat)
	tmp := strings.Split(tdata, ",")
	xdata = xdata + "Frame width " + tmp[0] + "<BR>"
	xdata = xdata + "Frame height " + tmp[1] + "<BR>"

	//-------------------------------------------------------------------------------------------------
	bdata = []byte(exefilea + " -i " + tnfile + " -show_entries format=duration -v quiet -of csv >tmp.csv")
	err = os.WriteFile(bfile, bdata, 0644)
	cmd = exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat = []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata = string(dat)
	tmp = strings.Split(tdata, ",")
	tmpa := strings.Split(tmp[1], ".")
	t := tmpa[0]
	i, _ := strconv.Atoi(t)
	mc := 0
	m := 0
	sc := 0
	for x := 0; x < i; x++ {
		mc++
		sc++
		if mc > 59 {
			m++
			mc = 0
			sc = 0
		}

	}
	xdata = xdata + "Length  " + strconv.Itoa(m) + ":" + strconv.Itoa(sc) + " <BR>"
	//-------------------------------------------------------------------------------------------------
	bdata = []byte(exefilea + " -i " + tnfile + " -show_entries stream=r_frame_rate  -of csv" + ">tmp.csv")

	err = os.WriteFile(bfile, bdata, 0644)
	cmd = exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat = []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata = string(dat)
	fr := ParseFrameRate(tdata)
	xdata = xdata + "Frames per second  " + fr + " <BR>"

	//-------------------------------------------------------------------------------------------------
	bdata = []byte(exefilea + " -i " + tnfile + "  -show_entries stream=bit_rate -v quiet -of csv >tmp.csv")
	err = os.WriteFile(bfile, bdata, 0644)
	cmd = exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat = []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata = string(dat)
	br := ParseBitRate(tdata)
	xdata = xdata + "Bit Rate " + br + " <BR>"
	xdata = xdata + "<BR><BR>"

	return xdata
}

func BasicDisplay(exefile string, exefilea string, tnfile string, fileName string, fc int, pfc int, fn string, xip string, sdir string) string {
	xdata := ""
	tp1 := TimePosition(exefilea, tnfile, 1)
	tp2 := TimePosition(exefilea, tnfile, 2)
	tp3 := TimePosition(exefilea, tnfile, 3)
	tp4 := TimePosition(exefilea, tnfile, 4)
	tp5 := TimePosition(exefilea, tnfile, 5)
	tp6 := TimePosition(exefilea, tnfile, 6)

	xdata = xdata + "File # " + strconv.Itoa(fc) + "<BR>"

	e := os.Remove("static/" + strconv.Itoa(pfc) + "1.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "2.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "3.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "4.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "5.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "6.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}

	cmd := exec.Command(exefile, "-ss", tp1, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"1.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}

	cmd = exec.Command(exefile, "-ss", tp2, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"2.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp3, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"3.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp4, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"4.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp5, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"5.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp6, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"6.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}

	xdata = xdata + " <A HREF='http://" + xip + ":8080/playvideo?video=" + fn + "&sdir=" + sdir + "'>  [ " + fn + " ] <BR> <IMG SRC=static/" + strconv.Itoa(pfc) + "1.png ALT=test123> <IMG SRC=static/" + strconv.Itoa(pfc) + "2.png  ALT=xerror> <IMG SRC=static/" + strconv.Itoa(pfc) + "3.png  ALT=error> <IMG SRC=static/" + strconv.Itoa(pfc) + "4.png  ALT=error><IMG SRC=static/" + strconv.Itoa(pfc) + "5.png  ALT=error> <IMG SRC=static/" + strconv.Itoa(pfc) + "6.png  ALT=error></A><BR> "

	return xdata
}

func ImageScrollDisplay(exefile string, tnfile string, fileName string, fc int, pfc int, fn string, xip string, sdir string) string {
	xdata := ""

	e := os.Remove("static/" + strconv.Itoa(pfc) + "1.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "2.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "3.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	cmd := exec.Command(exefile, "-ss", "00:00:01", "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"1.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", "00:00:10", "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"2.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", "00:00:20", "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"3.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	//xdata = xdata + "  <A HREF='file:///" + tnfile + "'>  [ " + fileName + " ] <BR> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "1.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "2.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "3.png" + "  ALT=error> </A><BR> "

	xdata = xdata + "<div class='scroll-container'>"
	xdata = xdata + "<img src= " + "static/" + strconv.Itoa(pfc) + "1.png  allt='test' width='600' height='400'>"
	xdata = xdata + "<img src= " + "static/" + strconv.Itoa(pfc) + "2.png  allt='test' width='600' height='400'>"
	xdata = xdata + "<img src= " + "static/" + strconv.Itoa(pfc) + "3.png  allt='test' width='600' height='400'>"

	xdata = xdata + "</div>"

	//-------------------------------------------------------------------------------------------------
	return xdata
}

func MoveDisplay(exefile string, exefilea string, tnfile string, fileName string, pfc int, fn string, xip string, sdir string) string {
	xdata := ""
	tp1 := TimePosition(exefilea, tnfile, 1)
	tp2 := TimePosition(exefilea, tnfile, 2)
	tp3 := TimePosition(exefilea, tnfile, 3)
	tp4 := TimePosition(exefilea, tnfile, 4)
	tp5 := TimePosition(exefilea, tnfile, 5)
	tp6 := TimePosition(exefilea, tnfile, 6)

	e := os.Remove("static/" + strconv.Itoa(pfc) + "1.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "2.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "3.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "4.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "5.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	e = os.Remove("static/" + strconv.Itoa(pfc) + "6.png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	cmd := exec.Command(exefile, "-ss", tp1, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"1.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	fmt.Println(tp1)
	cmd = exec.Command(exefile, "-ss", tp2, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"2.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp3, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"3.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp4, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"4.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp5, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"5.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", tp6, "-i", tnfile, "-vframes", "100", "-s", "128x96", "static/"+strconv.Itoa(pfc)+"6.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}

	xdata = xdata + "  <IMG SRC=static/" + strconv.Itoa(pfc) + "1.png ALT=test123> <IMG SRC=static/" + strconv.Itoa(pfc) + "2.png  ALT=xerror><BR> <IMG SRC=static/" + strconv.Itoa(pfc) + "3.png  ALT=error> <IMG SRC=static/" + strconv.Itoa(pfc) + "4.png  ALT=error><BR><IMG SRC=static/" + strconv.Itoa(pfc) + "5.png  ALT=error> <IMG SRC=static/" + strconv.Itoa(pfc) + "6.png  ALT=error><BR> "

	return xdata
}

func TimePosition(exefilea string, tnfile string, ctl int) string {
	xdata := ""
	bfile := ""
	switch runtime.GOOS {
	case "windows":
		bfile = "tmp.bat"
	case "linux":
		bfile = "tmp.sh"

	}
	//-------------------------------------------------------------------------------------------------
	bdata := []byte(exefilea + " -i " + tnfile + " -show_entries format=duration -v quiet -of csv >tmp.csv")
	err := os.WriteFile(bfile, bdata, 0777)
	if runtime.GOOS == "linux" {
		bfile = "./" + bfile
	}
	cmd := exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat := []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata := string(dat)
	tmp := strings.Split(tdata, ",")
	tmpa := strings.Split(tmp[1], ".")
	t := tmpa[0]
	i, _ := strconv.Atoi(t)
	mc := 0
	m := 0
	sc := 0
	for x := 0; x < i; x++ {
		mc++
		sc++
		if mc > 59 {
			m++
			mc = 0
			sc = 0
		}

	}
	if m < 2 {
		xdata = "00:00:10"
	} else {
		if m < 60 {
			nt := m / 2
			switch {
			case ctl == 1:
				ntt := nt / 2
				nttt := ntt / 2
				ntttt := nttt / 2
				nt = nt - ntt
				nt = nt - nttt
				nt = nt - ntttt
			case ctl == 2:
				ntt := nt / 2
				nttt := ntt / 2
				nt = nt - ntt
				nt = nt - nttt
			case ctl == 3:
				ntt := nt / 2
				nt = nt - ntt
			case ctl == 4:
				ntt := nt / 2
				nt = nt + ntt
			case ctl == 5:
				ntt := nt / 2
				nttt := ntt / 2
				nt = nt + ntt
				nt = nt + nttt
			case ctl == 6:
				ntt := nt / 2
				nttt := ntt / 2
				ntttt := nttt / 2
				nt = nt + ntt
				nt = nt + nttt
				nt = nt + ntttt
			}
			if nt > 9 {
				xdata = "00:" + strconv.Itoa(nt) + ":00"
			} else {
				xdata = "00:0" + strconv.Itoa(nt) + ":00"
			}
		} else {
			xdata = "00:59:00"
		}
	}

	return xdata
}

func CheckforFile(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func InitPage(xip string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xxip := ""
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Video Web Server</title>"
	xdata = DateTimeDisplay(xdata)
	//------------------------------------------------------------------------
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------

	xdata = xdata + "<body>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>Video Web Server</H1>"
	//---------
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xxip = fmt.Sprintf("%s", ipv4)
		}
	}
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<div id='txtdt'></div>"

	xdata = xdata + "Host Port IP : " + xip
	xdata = xdata + "<BR> Machine IP : " + xxip + "</p>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080/about'> [ About ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/display?page=1'> [ Display ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/static/index.html'> [ Static Index ] </A>  "
	xdata = xdata + "<BR><BR>Video Web Server"

	//------------------------------------------------------------------------

	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// ----------------------------------------------------------------
func AboutPage(xip string) string {
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>About Page</title>"
	xdata = LoopDisplay(xdata)
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: lightblue;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>Video Web Server</p>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/Video-Web-Server'> [ Video Web Server GitHub Repository ] </A>  "
	xdata = xdata + "<BR><BR>"
	//------------------------------------------------------------------------
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"

	xdata = xdata + "Video Web Server"
	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func DisplayPage(subdir bool, xip string, port string, page string, sdir string, exefile string, exefilea string, drive string, wdir string, pgsize int, maxsel int, display int) string {
	//---------------------------------------------------------------------------
	pgselect := ""
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	switch {
	case display == 1:
		xdata = xdata + "<meta name='viewport' content='width=device-width, initial-scale=1'>"
		xdata = xdata + "<style>"
		xdata = xdata + "div.scroll-container {"
		xdata = xdata + "background-color: #333;"
		xdata = xdata + "overflow: auto;"
		xdata = xdata + "white-space: nowrap;"
		xdata = xdata + "padding: 10px;"
		xdata = xdata + "}"
		xdata = xdata + "div.scroll-container img {"
		xdata = xdata + "padding: 10px;"
		xdata = xdata + "}"
		xdata = xdata + "</style>"
	}
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>Video Display Page</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<p2>Video Web Server</p2>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p1>Page " + page + "</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	pg, _ := strconv.Atoi(page)
	fctl := false
	nsd := ""
	switch runtime.GOOS {
	case "windows":
		if _, err := os.Stat(exefile); err == nil {
			nsd = wdir + sdir + "/"
			fctl = true
		} else {
			fmt.Println(err)
		}

	case "linux":
		nsd = wdir + sdir + "/"
		fctl = true
	}
	if fctl {
		files, err := ioutil.ReadDir(nsd)
		if err != nil {

			log.Fatal(err)
		}

		fc := 0
		pfc := 0
		pgcnt := 0
		for _, file := range files {
			//fmt.Println(file.Name())
			if ValidFileType(strings.ToLower(path.Ext(file.Name()))) {
				pfc++
				fc++
				if pfc > pgsize {
					pfc = 0
					pgcnt++

				}
			}
		}
		pp := true
		np := true
		for x := 0; x < pgcnt+1; x++ {
			if pg > 1 {
				if pp {
					pgselect = pgselect + "  <A HREF='http://" + xip + ":8080/display?page=" + strconv.Itoa(pg-1) + "&sdir=" + sdir + "'> <<  </A>  "
					pp = false
				}
			}
			if pg == x+1 {
				//pgselect = pgselect + "..."
				pgselect = pgselect + "[<B>" + strconv.Itoa(pg) + "</B>]"
			} else {

				if x < maxsel {
					pgselect = pgselect + "  <A HREF='http://" + xip + ":8080/display?page=" + strconv.Itoa(x+1) + "&sdir=" + sdir + "'> [ " + strconv.Itoa(x+1) + " ] </A>  "
				}
			}
			if pgcnt > 1 {
				if x == pgcnt {
					if np {
						pgselect = pgselect + "  <A HREF='http://" + xip + ":8080/display?page=" + strconv.Itoa(pg+1) + "&sdir=" + sdir + "'> >>  </A><BR>"
						np = false
					}
				}
			}
		}
		xdata = xdata + "<center>"
		xdata = xdata + "File Count " + strconv.Itoa(fc)
		xdata = xdata + "<BR>"
		xdata = xdata + pgselect
		xdata = xdata + "<BR>"
		xdata = xdata + "</center>"
		pfc = 0
		xdata = xdata + "<center>"
		xdata = xdata + "<TABLE>"
		xdata = xdata + "<TD valign='top' with='160'>"
		xdata = xdata + "<center>"
		xdata = xdata + "<FIELDSET>"
		xdata = xdata + "<LEGEND>"
		xdata = xdata + " Folders "
		xdata = xdata + "</LEGEND>"
		if subdir {
			entries, err := os.ReadDir(wdir + "./")
			if err != nil {
				log.Fatal(err)
			}
			fcol := 0
			xdata = xdata + "<TABLE>"
			xdata = xdata + "<TD valign='top' with='50'>"
			xdata = xdata + "<center>"
			for _, e := range entries {
				fcol++
				if fcol == 1 {
					xdata = xdata + "  <A HREF='http://" + xip + ":8080/display?page=1&sdir=" + e.Name() + "'>  [" + e.Name() + "]  </A><BR>  "
				}
				if fcol > 3 {
					fcol = 0
				}
			}
			xdata = xdata + "</center>"
			xdata = xdata + "</TD>"
			xdata = xdata + "<TD valign='top' with='50'>"
			xdata = xdata + "<center>"
			fcol = 0
			for _, e := range entries {
				fcol++
				if fcol == 2 {
					xdata = xdata + "  <A HREF='http://" + xip + ":8080/display?page=1&sdir=" + e.Name() + "'>  [" + e.Name() + "]  </A><BR>  "
				}
				if fcol > 3 {
					fcol = 0
				}
			}

			xdata = xdata + "</center>"
			xdata = xdata + "</TD>"
			xdata = xdata + "<TD valign='top' with='50'>"
			xdata = xdata + "<center>"
			fcol = 0
			for _, e := range entries {
				fcol++
				if fcol == 3 {
					xdata = xdata + "  <A HREF='http://" + xip + ":8080/display?page=1&sdir=" + e.Name() + "'>  [" + e.Name() + "]  </A><BR>  "
				}
				if fcol > 3 {
					fcol = 0
				}
			}
			xdata = xdata + "</center>"
			xdata = xdata + "</TD>"
			xdata = xdata + "<TD valign='top' with='50'>"
			xdata = xdata + "<center>"
			fcol = 0
			for _, e := range entries {
				fcol++
				if fcol == 4 {
					xdata = xdata + "  <A HREF='http://" + xip + ":8080/display?page=1&sdir=" + e.Name() + "'>  [" + e.Name() + "]  </A><BR>  "
				}
				if fcol > 3 {
					fcol = 0
				}
			}
			xdata = xdata + "</center>"
			xdata = xdata + "</TD>"
			xdata = xdata + "</TABLE>"

		}
		xdata = xdata + "</FIELDSET>"
		xdata = xdata + "</center>"
		xdata = xdata + "</TD>"
		xdata = xdata + "<TD with='100'>"
		xdata = xdata + "</TD>"
		xdata = xdata + "<TD valign='top' with='300'>"
		xdata = xdata + "<center>"
		xdata = xdata + "<FIELDSET>"
		xdata = xdata + "<LEGEND>"
		xdata = xdata + sdir
		xdata = xdata + "</LEGEND>"
		pgcnt = 1
		fc = 0
		for _, file := range files {

			if ValidFileType(strings.ToLower(path.Ext(file.Name()))) {
				pfc++
				fc++
				if pfc > pgsize {
					pfc = 0
					pgcnt++
				}
				if pgcnt == pg {
					tfile := nsd + file.Name()
					tnfile := fixFileName(tfile)
					switch {
					case display == 0:
						xdata = xdata + BasicDisplay(exefile, exefilea, tnfile, file.Name(), fc, pfc, file.Name(), xip, sdir)
					case display == 1:
						xdata = xdata + ImageScrollDisplay(exefile, tnfile, file.Name(), fc, pfc, file.Name(), xip, sdir)
					}
					//-------------------------------------------------------------------------------------------------
					xdata = xdata + "<TABLE>"
					xdata = xdata + "<TD with='200'>"
					xdata = xdata + "<center>"
					xdata = xdata + " <A HREF='http://" + xip + ":8080/tagvideo?video=" + file.Name() + "&sdir=" + sdir + "'> [ Video Tags ] </A>  "
					xdata = xdata + "</center>"
					xdata = xdata + "</TD>"
					xdata = xdata + "<TD with='200'>"
					xdata = xdata + " -------------- "
					xdata = xdata + "</TD>"
					xdata = xdata + "<TD with='400'>"
					xdata = xdata + "<center>"
					xdata = xdata + FileData(exefilea, tnfile, file.Name())
					xdata = xdata + "</center>"
					xdata = xdata + "</TD>"
					xdata = xdata + "<TD with='200'>"
					xdata = xdata + " -------------- "
					xdata = xdata + "</TD>"
					xdata = xdata + "<TD with='200'>"
					xdata = xdata + "<center>"
					xdata = xdata + " <A HREF='http://" + xip + ":8080/movevideo?video=" + file.Name() + "&sdir=" + sdir + "'> [ Move Video ] </A>  "
					xdata = xdata + "</center>"
					xdata = xdata + "</TD>"
					xdata = xdata + "</TABLE>"
					if pfc < pgsize {
						xdata = xdata + "<HR>"
					}
				}
			}
		}
		xdata = xdata + "</FIELDSET>"
		xdata = xdata + "</center>"
		xdata = xdata + "</TD>"
		xdata = xdata + "</TABLE>"
		xdata = xdata + "</center>"

	}
	//------------------------------------------------------------------------
	xdata = xdata + "<center>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + pgselect
	xdata = xdata + "<BR><BR>"
	xdata = xdata + " <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func PlayVideoPage(xip string, port string, video string, exefile string, exefilea string, drive string, wdir string, sdir string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Play Video Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: white;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Play Video Page</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p2>Video Web Server</p2>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p1>Video " + video + "</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	//------------------------------------------------------------------------
	tfile := wdir + sdir + "/" + video
	tnfile := fixFileName(tfile)
	e := os.Remove("static/" + video + ".png")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	cmd := exec.Command(exefile, "-ss", "00:00:01", "-i", tnfile, "-vframes", "100", "-s", "600x400", "static/"+video+".png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	xdata = xdata + "<center>"
	xdata = xdata + " <A HREF='http://" + xip + ":8080/static/tmp.mp4'>  [ " + video + " ] <BR> <IMG SRC=static/" + video + ".png ALT=test123>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	//------------------------------------------------------------------------
	e = os.Remove("static/tmp.mp4")
	if e != nil {
		fmt.Printf("Delete Error: %s\n", e)
	}
	if strings.ToLower(path.Ext(tnfile)) == ".mp4" {

		_, _ = copy(tnfile, "static/tmp.mp4")
		//fmt.Printf("%d %e \n", c, ce)

	} else {

		cmd = exec.Command(exefile, "-i", tnfile, "-strict", "-2", "static/tmp.mp4")
		if err := cmd.Run(); err != nil {
			fmt.Printf("Command %s \n Error: %s\n", cmd, err)
		}
	}
	//------------------------------------------------------------------------

	return xdata

}

// ----------------------------------------------------------------
func TagVideoPage(xip string, port string, video string, exefile string, exefilea string, drive string, wdir string, sdir string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Tag Video Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: white;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Tag Video Page</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p2>Video Web Server</p2>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p1>Video " + video + "</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	//------------------------------------------------------------------------

	return xdata

}

// ----------------------------------------------------------------
func MoveVideoPage(xip string, port string, video string, exefile string, exefilea string, drive string, wdir string, sdir string, subdir bool) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Move Video Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: white;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Move Video Page</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p2>Video Web Server</p2>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p1>Video " + video + "</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	tfile := wdir + sdir + "/" + video
	tnfile := fixFileName(tfile)
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<center>"
	xdata = xdata + "<TABLE>"
	xdata = xdata + "<TD valign='top' with='200'>"
	xdata = xdata + "<center>"
	xdata = xdata + MoveDisplay(exefile, exefilea, tnfile, tfile, 1, tfile, xip, sdir)
	xdata = xdata + "</center>"
	xdata = xdata + "</TD>"
	xdata = xdata + "<TD valign='top' with='200'>"
	xdata = xdata + "<center>"
	xdata = xdata + FileData(exefilea, tnfile, tfile)
	xdata = xdata + "</center>"
	xdata = xdata + "</TD>"
	xdata = xdata + "<TD with='200'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<FIELDSET>"
	xdata = xdata + "<LEGEND>"
	xdata = xdata + "Move to Folder"
	xdata = xdata + "</LEGEND>"

	if subdir {
		entries, err := os.ReadDir(wdir + "./")
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range entries {
			xdata = xdata + "  <A HREF='http://" + xip + ":8080/movevideocomplete?video=" + video + "&sdir=" + sdir + "&ddir=" + e.Name() + "'> " + e.Name() + " </A><BR>  "

		}
	}
	xdata = xdata + "</FIELDSET>"

	xdata = xdata + "</center>"

	xdata = xdata + "</TD>"
	xdata = xdata + "</TABLE>"

	xdata = xdata + "</center>"

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	//------------------------------------------------------------------------

	return xdata

}

// ----------------------------------------------------------------
func MoveVideoCompletePage(xip string, port string, video string, exefile string, exefilea string, drive string, wdir string, sdir string, ddir string) string {
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Move Video Complete Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: white;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Move Complete Video Page</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p2>Video Web Server</p2>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p1>Video " + video + "</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	//------------------------------------------------------------------------
	tfile := wdir + sdir + "/" + video
	srcfile := fixFileName(tfile)
	tfile = wdir + ddir + "/" + video
	dstfile := fixFileName(tfile)
	xdata = xdata + "<center><BR>"

	xdata = xdata + "Souce :" + srcfile + "<BR>"
	xdata = xdata + "Destination :" + dstfile + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	//------------------------------------------------------------------------
	err := os.Rename(srcfile, dstfile)
	if err != nil {
		fmt.Println(err)
		xdata = xdata + "<center><B>"
		xdata = xdata + "<BR><BR>Rename Error <BR><I>" + err.Error() + "</I><BR>"
		xdata = xdata + "</B></center>"
		return xdata
	}
	//---------------------------------------------------------------------------

	return xdata

}
