//Copyright (c) 2019 Romano (Viacoin developer)
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
package binaries

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"os"
	"os/exec"
	"time"
)

// Download binaries from github
func Download() {

	url := "https://github.com/viacoin/viacoin/releases/download/v0.16.3/viacoin-0.16.3-x86_64-linux-gnu.tar.gz"

	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("Downloading Viacoin binaries (%.2f%%)\n",
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Download saved to ./%v \n", resp.Filename)

}

func Untar() {
	exec.Command("/bin/sh", "untar.sh").Run()
}
