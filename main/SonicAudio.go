package main

/**
This program is possible thanks to [faiface] Beep library
*/
import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func startIntroTheme() {
	playAudio("audio/sonic_intro.mp3")
}

func startMainTheme() {
	playAudio("audio/sonic_audio.mp3")
}

/**
[speaker.Init] Wwo arguments: the sample rate, and the buffer size.
The sample rate argument simply tells the speaker how quickly it should push the samples to the output.
*/
func playAudio(audioPath string) {
	audioFile, err := os.Open(audioPath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(audioFile)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	loop := beep.Loop(100, streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))
	select {}
}
