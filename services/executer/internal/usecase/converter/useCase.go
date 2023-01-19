package converter

import (
	"fmt"
	"github.com/goccy/go-json"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"path/filepath"
	"smartdom/pkg/tools/systools"
	"smartdom/services/executer/internal/domain"
)

func Run(source string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {

		if filepath.Ext(source) == ".mp4" {

			if isCandidate(source) {
				outFile := fmt.Sprintf("/tmp/%s", info.Name())

				if err = convert(source, outFile); err != nil {
					fmt.Printf("convert Error %s", err)
					return err
				}
				if err = systools.Move(outFile, source, info.Mode()); err != nil {
					fmt.Printf("mv Error %s", err)
					return err
				}
			}
		}
	}
	return err
}

func isCandidate(source string) bool {
	probe, _ := ffmpeg.Probe(source)
	file := domain.VideoInfo{}

	json.Unmarshal([]byte(probe), &file)

	res := domain.CodecH264 == file.Streams[0].CodecName

	//size, _ := strconv.Atoi(file.Format.Size)
	//duration, _ := strconv.ParseFloat(file.Format.Duration, 32)
	//if res {
	//	fmt.Printf("%s | %.2fmin | %dMb |\n", file.Format.Filename, duration/60, size/1024/1024)
	//}

	return res
}

func convert(source string, dest string) error {
	return ffmpeg.Input(source).
		Output(dest,
			ffmpeg.KwArgs{"c:v": "libx265"},
			ffmpeg.KwArgs{"crf": 25},
			ffmpeg.KwArgs{"vf": "scale=960x540"},
		).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

}
