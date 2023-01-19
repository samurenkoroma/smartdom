.DEFAULT_GOAL := build-convert
build-convert:
	 go build  -o ~/bin/convert  services/convertor/cmd/app.go
#ffmpeg -i 05.\ История\ в\ игре.mp4 -vcodec libx265 -crf 25 -vf "scale=iw/2:ih/2" 05.mp4