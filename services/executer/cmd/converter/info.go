package converter

import (
	"fmt"
	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"path/filepath"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Посмотреть информацию о файлах в папке",
	Long:  `Посмотреть информацию о файлах в папке`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath.Walk(args[0],
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {

					//ext := filepath.Ext(path)
					//convertedName := strings.Replace(path, ext, fmt.Sprintf("__convert__%s", ext), -1)
					//fmt.Print(convertedName)
					//ffmpeg.
					//	Input(path).
					//	OverWriteOutput().
					//	Output(convertedName,
					//		ffmpeg.KwArgs{"metadata": fmt.Sprintf("title=%s", info.Name())},
					//		ffmpeg.KwArgs{"c": "copy"},
					//	).
					//	ErrorToStdOut().
					//	Run()
					//os.Remove(path)
					//os.Rename(convertedName, path)
					probe, _ := ffmpeg.Probe(path)
					fmt.Printf("File: %s", probe)
					os.Exit(0)
				}
				return err
			})
	},
}

func init() {
	ConvertCmd.AddCommand(infoCmd)
}
