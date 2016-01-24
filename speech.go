package main

import (
	"errors"
	"time"

	"github.com/gophergala2016/papercast/mactts"
)

const (
	speakTimeout = 2 * time.Minute
)

var voice *mactts.VoiceSpec

func init() {
	initVoice()
}

func initVoice() {
	voice, err := mactts.GetVoice(3) // hopefully Alex :) FIXME
	if err != nil {
		panic(err)
	}
	vd, err := voice.Description()
	if err != nil {
		panic(err)
	}
	println("Voice: " + vd.Name())
}

func speak(article string) (*ResponseBuffer, error) {
	var audioBuffer ResponseBuffer

	af, err := mactts.NewOutputAACFile(&audioBuffer, float64(22050), 1, 16)
	if err != nil {
		return nil, err
	}
	defer af.Close()

	eaf, err := af.ExtAudioFile()
	if err != nil {
		return nil, err
	}
	defer eaf.Close()

	channel, err := mactts.NewChannel(voice)
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	err = channel.SetExtAudioFile(eaf)
	if err != nil {
		return nil, err
	}

	done := make(chan int)
	err = channel.SetDone(func() {
		done <- 1
		close(done)
	})
	if err != nil {
		return nil, err
	}

	err = channel.SpeakString(article)
	if err != nil {
		return nil, err
	}

	select {
	case <-done:
	case <-time.After(speakTimeout):
		return nil, errors.New("Timed out speaking")
	}

	return &audioBuffer, nil
}
