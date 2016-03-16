package ghdownloads

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var Client = http.DefaultClient

func Count(repo string) (int, error) {
	return count("https://api.github.com/repos/" + repo + "/releases")
}

func count(url string) (int, error) {
	resp, err := Client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("bad response from GitHub: %s", resp.Status)
	}
	var releases []struct {
		Assets []struct {
			DownloadCount int `json:"download_count"`
		} `json:"assets"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return 0, err
	}
	count := 0
	for _, release := range releases {
		for _, asset := range release.Assets {
			count += asset.DownloadCount
		}
	}
	if len(resp.Header.Get("Link")) > 0 {
		fmt.Println("WARNING: Ignoring extra items:", resp.Header.Get("Link"))
		// TODO: load 'rel="next"' link and get those too
	}
	return count, nil
}
