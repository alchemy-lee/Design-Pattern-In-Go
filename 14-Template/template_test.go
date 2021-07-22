package template

import "testing"

func TestTemplate(t *testing.T) {
	t.Log("HTTP Downloader")
	httpDownloader := NewHTTPDownloader()
	httpDownloader.Download("test.txt")

	t.Log("FTP Downloader")
	ftpDownloader := NewFTPDownloader()
	ftpDownloader.Download("test.txt")
}
