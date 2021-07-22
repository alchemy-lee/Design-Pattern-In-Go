package template

import "fmt"

type Downloader interface {
	Download(uri string)
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

func (t *template) Download(uri string) {
	t.uri = uri
	t.implement.download()
	t.implement.save()
}

func (t *template) save() {
	fmt.Printf("Default save.\n")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := &template{implement: downloader}
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("Download %s via HTTP.\n", d.uri)
}

func (d *HTTPDownloader) save() {
	fmt.Printf("HTTP save.\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := &template{implement: downloader}
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("Download %s via FTP.\n", d.uri)
}
