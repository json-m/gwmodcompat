package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ArcDps struct {
	Latest  time.Time `json:"latest"`
	History []History `json:"history"`
}

// main function for arcdps
func modArcdps() {
	latest, chk, err := latest_arcdps()
	if err != nil {
		log.Println("modArcdps.latest_arcdps:", err)
		return
	}

	// is latest newer than latest in database?
	if latest.After(modData.ArcDps.Latest) {
		//log.Println("ArcDps: new version available")
	} else {
		//log.Println("ArcDps: no new version available")
		return
	}

	// set latest
	modData.ArcDps.Latest = latest

	// append to modData history
	modData.ArcDps.History = append(modData.ArcDps.History, History{
		Version:  latest.String(),
		Checksum: chk,
	})

	// download latest version
	//fmt.Println("downloading arcdps", strconv.FormatInt(latest.Unix(), 10))
	err = download_arcdps(strconv.FormatInt(latest.Unix(), 10))
	if err != nil {
		log.Println("modArcdps.download_arcdps:", err)
		return
	}

	// write data
	err = writeData()
	if err != nil {
		log.Println("modArcdps.writeData:", err)
		return
	}

	return
}

// gets latest version of arcdps
func latest_arcdps() (time.Time, string, error) {
	dllURL := "https://www.deltaconnected.com/arcdps/x64/d3d11.dll"
	checksumURL := "https://www.deltaconnected.com/arcdps/x64/d3d11.dll.md5sum"

	// make head request to url, parse out last-modified header
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("HEAD", dllURL, nil)
	if err != nil {
		return time.Time{}, "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, "", err
	}
	defer resp.Body.Close()

	// parse out time from last-modified header
	//< last-modified: Sat, 09 Sep 2023 14:30:02 GMT
	//fmt.Println(resp.Header.Get("last-modified"))
	t, err := time.Parse(time.RFC1123, resp.Header.Get("last-modified"))
	if err != nil {
		return time.Time{}, "", err
	}

	// get checksum url
	resp, err = client.Get(checksumURL)
	if err != nil {
		return time.Time{}, "", err
	}

	// read in body
	buf := make([]byte, 32)
	_, err = resp.Body.Read(buf)
	if err != nil {
		return time.Time{}, "", err
	}

	// convert to string
	b := string(buf)

	return t, strings.Split(b, " ")[0], nil
}

// downloads latest version of arcdps to it's place
func download_arcdps(version string) error {
	// todo: check if version requested is version at server
	url := "https://www.deltaconnected.com/arcdps/x64/d3d11.dll"
	outfile := fmt.Sprintf("mods/arcdps/%s-d3d11.dll", version)

	// download file to "mods/arcdps"
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()

	// make request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// write file contents
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
