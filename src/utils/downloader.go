package utils

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type Download struct {
	Track TrackInfo
}

// ZipResult contains information about the ZIP creation process
type ZipResult struct {
	ZipPath      string
	SuccessCount int
	Errors       []error
}

func NewDownload(track TrackInfo) *Download {
	return &Download{Track: track}
}

func (d *Download) Process() (string, []byte, error) {
	switch {
	case d.Track.CdnURL == "":
		return "", nil, errMissingCDNURL
	case d.Track.Platform == "spotify":
		return d.processSpotify()
	default:
		return d.processDirectDL()
	}
}

func (d *Download) processDirectDL() (string, []byte, error) {
	track := d.Track

	// Check for Telegram URL pattern (e.g., https://t.me/username/123) for *YouTube*
	if regexp.MustCompile(`^https:\/\/t\.me\/([a-zA-Z0-9_]{5,})\/(\d+)$`).MatchString(track.CdnURL) {
		coverData, err := getCover(track.Cover)
		return track.CdnURL, coverData, err
	}

	filePath, err := downloadFile(context.Background(), track.CdnURL, "", false)
	if err != nil {
		return "", nil, fmt.Errorf("failed to download file: %w", err)
	}

	coverData, err := getCover(track.Cover)
	return filePath, coverData, err
}

// ZipTracks creates a ZIP archive containing all tracks from PlatformTracks
func ZipTracks(tracks *PlatformTracks) (*ZipResult, error) {
	zipFilename := generateRandomZipName()
	result := &ZipResult{ZipPath: zipFilename}
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to create zip file: %v", err)
	}

	defer func(zipFile *os.File) {
		_ = zipFile.Close()
	}(zipFile)

	zipWriter := zip.NewWriter(zipFile)

	defer func(zipWriter *zip.Writer) {
		_ = zipWriter.Close()
	}(zipWriter)

	for _, track := range tracks.Results {
		err := processTrack(zipWriter, track)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Errorf("track %s: %v", track.ID, err))
			continue
		}
		result.SuccessCount++
	}

	if absPath, err := filepath.Abs(zipFilename); err == nil {
		result.ZipPath = absPath
	}

	if result.SuccessCount == 0 {
		return result, fmt.Errorf("no tracks were successfully added to the zip")
	}

	return result, nil
}

// processTrack handles downloading and adding a single track to the ZIP
func processTrack(zipWriter *zip.Writer, track MusicTrack) error {
	// Get the track data
	apiData := NewApiData(track.URL)
	trackData, err := apiData.GetTrack()
	if err != nil {
		return fmt.Errorf("failed to get track info: %v", err)
	}

	filename, _, err := NewDownload(*trackData).Process()
	if err != nil {
		return fmt.Errorf("failed to download track: %v", err)
	}

	audioFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open downloaded file: %v", err)
	}
	defer func(audioFile *os.File) {
		_ = audioFile.Close()
	}(audioFile)

	baseName := filepath.Base(filename)
	zipEntry, err := zipWriter.Create(baseName)
	if err != nil {
		return fmt.Errorf("failed to create zip entry: %v", err)
	}

	if _, err := io.Copy(zipEntry, audioFile); err != nil {
		return fmt.Errorf("failed to write to zip: %v", err)
	}

	defer func() {
		_ = os.Remove(filename)
	}()

	return nil
}
