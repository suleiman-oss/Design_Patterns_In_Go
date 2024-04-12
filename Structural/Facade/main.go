package main

import videoconversion "github.com/suleiman-oss/Design_Patterns_in_Go/Structural/Facade/VideoConversion"

func main() {
	v := videoconversion.NewVideoProcess()
	v.ProcessVideo("name.mp4", "ACC4")
}
