package main

import "fmt"

type MediaPlayer interface {
	Play(audioType string, fileName string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

/*AdvancedMediaPlayer*/

type VlcPlayer struct {
}

func (*VlcPlayer) PlayVlc(fileName string) {
	fmt.Printf("Playing vlc file. Name: %v \n", fileName)
}

func (*VlcPlayer) PlayMp4(fileName string) {
	fmt.Printf("not working, fileName: %v \n ", fileName)
}

type Mp4Player struct {
}

func (*Mp4Player) PlayVlc(fileName string) {
	fmt.Printf("not working, fileName: %v \n", fileName)
}

func (*Mp4Player) PlayMp4(fileName string) {
	fmt.Printf("Playing mp4 file. Name: %v \n", fileName)
}

/*MediaPlayer*/

type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) MediaAdapter(audioType string) {
	switch audioType {
	case "vlc":
		m.advancedMusicPlayer = &VlcPlayer{}
	case "mp4":
		m.advancedMusicPlayer = &Mp4Player{}
	}
}

func (m *MediaAdapter) Play(audioType string, fileName string) {
	switch audioType {
	case "vlc":
		m.advancedMusicPlayer.PlayVlc(fileName)
	case "mp4":
		m.advancedMusicPlayer.PlayMp4(fileName)
	}
}

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (a *AudioPlayer) Play(audioType string, fileName string) {
	switch audioType {
	case "mp3":
		fmt.Printf("Playing mp3 file. Name: %v \n", fileName)
	case "vlc", "mp4":
		a.mediaAdapter.MediaAdapter(audioType)
		a.mediaAdapter.Play(audioType, fileName)
	default:
		fmt.Printf("Invalid media. %v format not supported \n", audioType)
	}
}
