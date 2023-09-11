package pkg

import (
	"encoding/json"
	"fmt"
	"golang.org/x/mod/semver"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// KnowThyEnemy struct for this mod's data
type KnowThyEnemy struct {
	Latest  string    `json:"latest"`
	History []History `json:"history"`
}

// main function for knowthyenemy
func modKnowthyenemy() {
	// first, check if release exists in database
	latest, err := latest_KTE()
	if err != nil {
		log.Println("modKnowthyenemy.latest_KTE:", err)
		return
	}

	// is latest newer than latest in database?
	if semver.Compare(fmt.Sprintf("v%s", latest), fmt.Sprintf("v%s", modData.KnowThyEnemy.Latest)) == 1 {
		log.Println("KTE: new version available")
	} else {
		log.Println("KTE: no new version available")
		return
	}

	// set latest to latest version in database
	modData.KnowThyEnemy.Latest = latest

	// append to modData history
	modData.KnowThyEnemy.History = append(modData.KnowThyEnemy.History, History{
		Version:  latest,
		Checksum: "",
	})

	// download latest version
	err = download_KTE(latest)
	if err != nil {
		log.Println("modKnowthyenemy.download_KTE:", err)
		return
	}

	// write data file
	err = writeData()
	if err != nil {
		log.Println("modKnowthyenemy.writeData:", err)
		return
	}
}

// gets latest tag from github
func latest_KTE() (string, error) {
	releases := "https://api.github.com/repos/typedeck0/Know-thy-enemy/releases/latest"

	// get releases and parse out tags
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(releases)
	if err != nil {
		log.Println("TestGithub.Get:", err)
		return "", err
	}
	defer resp.Body.Close()

	// parse out tags
	var release GithubReleaseStruct
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		log.Println("TestGithub.Decode:", err)
		return "", err
	}

	return release.TagName, nil
}

// downloads latest version of kte to it's place
func download_KTE(version string) error {
	url := fmt.Sprintf("https://github.com/typedeck0/Know-thy-enemy/releases/download/%s/know-thy-enemy.dll", version)
	outfile := fmt.Sprintf("mods/knowthyenemy/%s-know-thy-enemy.dll", version)

	// download file to "mods/knowthyenemy"
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

type GithubReleaseStruct struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		Label    any    `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
	Body       string `json:"body"`
}
