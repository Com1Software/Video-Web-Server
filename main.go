package main

import (
	"net"
	"net/http"
	"os/exec"
	"runtime"

	"fmt"
	"os"
)

//----------------------------------------------------------------
func main() {
	fmt.Println("Video Web Server")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	xip := fmt.Sprintf("%s", GetOutboundIP())
	port := "8080"
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
			xdata := DisplayPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Generate Static Site Information Handler
		http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
			xdata := Generate(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Generate Static Site Information Handler
		http.HandleFunc("/generatesite", func(w http.ResponseWriter, r *http.Request) {
			xdata := GenerateSite(xip)
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
	xdata = xdata + "<H1>Video Web Server.</H1>"
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
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/display'> [ Dynamic Display ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/static/index.html'> [ Static Index ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/generate'> [ Generate Site Info] </A>  "
	xdata = xdata + "<BR><BR>Video Web Server...."

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

//----------------------------------------------------------------
func AboutPage(xip string) string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>About Page</title>"
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

//----------------------------------------------------------------
func Generate(xip string) string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Generate Static Site Info</title>"
	//------------------------------------------------------------------------

	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: grey;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
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
	xdata = xdata + "<body>"
	xdata = xdata + "<H1>Generate Static Site Info</H1>"
	//---------
	xdata = xdata + "<center>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080/generatesite'> [ Start Generate ] </A> <BR><BR> "

	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  <BR><BR>"

	xdata = xdata + "<p1>Video Web Server</p1>"
	xdata = xdata + "</center>"

	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

//----------------------------------------------------------------
func GenerateSite(xip string) string {
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Generate Static Site Complete</title>"
	//------------------------------------------------------------------------

	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: green;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
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
	xdata = xdata + "<body>"
	xdata = xdata + "<H1>Generate Static Site Complete</H1>"
	//---------
	xdata = xdata + "<center>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  <BR><BR>"

	xdata = xdata + "<p1>Video Web Server</p1>"
	xdata = xdata + "</center>"

	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"

	return xdata

}

//----------------------------------------------------------------
func DisplayPage(xip string) string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Display Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: black;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
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
	xdata = xdata + "<H1>Display Page</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p1>Go Web Server</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p2>Go Web Server</p2>"
	xdata = xdata + "</center>"

	//------------------------------------------------------------------------
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

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
	xdata = xdata + "  document.getElementById('txtdt').innerHTML =  '+day+', '+month+' '+dm+', '+yr' - '+h' + ':' + m + ':' + s+' '+ampm+;"
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

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
