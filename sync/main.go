package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	// 이거 경로 찾는 부분 제거하고 image 폴더로 가리키
	sr := data.ScanResultRetrieve(id)
	hscans1, err := loadHscans(path, int(sr.B_scans), scan_ext)

	if err != nil {
		logger.Errorf("loadHscans() error : %v", err)
	}

	vscans, _ := processScanVertical(hscans1, sr)
	for i, vscan := range vscans {
		saveImage(vscan, path+fmt.Sprintf("/v%03d"+scan_ext, i))
	}

	diff := getBreathingDiff(id, path, scan_ext)

	_, segments, err := loadAndSplitSegments(path+"/layers"+data_ext, int(sr.B_scans), 7)
	if err != nil {
		logger.Errorf("loadAndSplitSegments() error : %v", err)
		return
	}
	rpe := segments[RPE]

	rpeW := (*rpe).Bounds().Dx()
	rpeH := (*rpe).Bounds().Dy()

	if len(diff) < rpeH {
		logger.Errorf("Invalid Breathing Process: DIFF=%v, rpeH=%v", len(diff), rpeH)
		return
	}

	tempdir_path := filepath.Join(path, "temp")
	os.MkdirAll(tempdir_path, os.ModePerm)

	//logger.Debugf("height: %v", rpeH)
	for x := 0; x < rpeH; x++ {
		hScanName := fmt.Sprintf("%03d"+scan_ext, x)
		adjustBreathingHscan(path, hScanName, diff, x)
		if err != nil {
			logger.Errorf("adjustBreathingHscan error = %v", err)
			return
		}
	}

	err = adjustBreathingLayer(id, path, diff)
	if err != nil {
		logger.Errorf("adjustBreathingLayer error = %v", err)
		return
	}

	hscans2, err := loadHscans(tempdir_path, int(sr.B_scans), scan_ext)
	if err != nil {
		logger.Errorf("loadHscans() error : %v", err)
	}

	tileWalls, err := processScanTiling(hscans2, sr, path)
	if err == nil {
		for i, tile := range *tileWalls {
			saveImage(tile, path+fmt.Sprintf("/tile_%d"+data_ext, i))
		}
	}

	RemoveContents(tempdir_path)
	os.Remove(tempdir_path)

	for x := 0; x < rpeW; x++ {
		vScanName := fmt.Sprintf("v%03d%s", x, scan_ext)
		os.Remove(path + "/" + vScanName)
	}

	f, err := os.Create(markerFile)
	if err != nil {
		logger.Errorf("Error creating marker file: %v", err)
		return
	}
	defer f.Close()

	today := time.Now().Format("2006-01-02 15:04:05")
	if _, err := f.WriteString(today); err != nil {
		logger.Errorf("Error writing to marker file: %v", err)
		return
	}
}
