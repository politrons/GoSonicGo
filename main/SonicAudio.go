package main

/**
This program is possible thanks to [faiface] Beep library
*/
import (
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

/**
[speaker.Init] Wwo arguments: the sample rate, and the buffer size.
The sample rate argument simply tells the speaker how quickly it should push the samples to the output.
*/
func startAudio() {
	audioFile, err := os.Open("audio/sonic_audio.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(audioFile)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
