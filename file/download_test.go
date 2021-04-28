package file

import (
	"github.com/LinkinStars/pan/conf"
	"testing"
)

func TestDownload(t *testing.T) {
	fileDownloader := NewDownloaderWithFsID(conf.TestData.AccessToken, conf.TestData.FsID, conf.TestData.LocalFilePath)
	err := fileDownloader.Download()
	if err != nil {
		t.Fail()
	} else {
		t.Logf("TestDownload Success")
	}
}
