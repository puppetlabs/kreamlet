package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	//"path/filepath"
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func DownloadFile(url string, destName string) error {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server return non-200 status: %s\n", resp.Status)
		return err
	}

	size := resp.ContentLength

	// create dest
	dest, err := os.Create(destName)
	if err != nil {
		fmt.Printf("Can't create %s: %v\n", destName, err)
		return err
	}
	defer dest.Close()

	p := mpb.New(
		mpb.WithWidth(60),
		mpb.WithFormat("[=>-|"),
		mpb.WithRefreshRate(180*time.Millisecond),
	)

	bar := p.AddBar(size,
		mpb.PrependDecorators(
			decor.CountersKibiByte("% 6.1f / % 6.1f"),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_MMSS, 1024*8),
			decor.Name(" ] "),
			decor.AverageSpeed(decor.UnitKiB, "% .2f"),
		),
	)

	// create proxy reader
	reader := bar.ProxyReader(resp.Body)

	// and copy from reader, ignoring errors
	io.Copy(dest, reader)

	p.Wait()

	return nil

}
