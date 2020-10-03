package golib

import (
	"regexp"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/gogf/gf/errors/gerror"
)

func IsDigOrAlpha(s string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(s)
}

func DlFileWithProgress(url string, dlDst string, progressCallBack func(dlProgress int32, dlSpeed float64)) error {

	client := grab.NewClient()
	req, err := grab.NewRequest(dlDst, url)
	if err != nil {
		return gerror.Wrap(err, "failed to create request")
	}

	req.NoResume = true
	resp := client.Do(req)
	if err != nil {
		return gerror.Wrap(err, "failed to do the request")
	}

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			progressCallBack(int32(100*resp.Progress()), resp.BytesPerSecond())
		case <-resp.Done:
			if resp.Err() != nil {
				return gerror.Wrap(err, "failed to download")
			}
			progressCallBack(100, resp.BytesPerSecond())
			break Loop
		}
	}
	return nil
}
