package videoconversion

import "fmt"

type videoPorcess struct{}

func NewVideoProcess() *videoPorcess {
	return &videoPorcess{}
}

func (v *videoPorcess) ProcessVideo(name string, format string) {
	v.downloadVideo(name)
	v.encodeVideo(name, format)
}
func (v *videoPorcess) downloadVideo(name string) {
	fmt.Printf("Downloading Video: %s\n", name)
}
func (v *videoPorcess) encodeVideo(name string, format string) {
	fmt.Printf("Encoding the video:%s to the format: %s\n", name, format)
}
