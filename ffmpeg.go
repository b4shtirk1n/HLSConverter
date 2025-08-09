package main

import (
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func (m Model) Convert() {
	args := "libx265"

	if m.IsGPU {
		args = "hevc_nvenc tag:v hvc1 profile fast"
	}

	os.Mkdir(m.OutPath, 0750)
	if m.Err = os.Chdir(m.OutPath); m.Err != nil {
		p.Send(ThrowErrorMsg{})
		return
	}

	cmd = exec.Command("ffmpeg", "-y", "-i",
		m.SelectedFile,
		"-map", "0:v:0", "-map", "0:a:0",
		"-map", "0:v:0", "-map", "0:a:0",
		"-map", "0:v:0", "-map", "0:a:0",
		"-map", "0:v:0", "-map", "0:a:0",
		"-map", "0:v:0", "-map", "0:a:0",
		"-map", "0:v:0", "-map", "0:a:0",
		"-c:v", args, "-crf", "26", "-c:a", "aac", "-ar", "48000",
		"-filter:v:0", "scale=w=640:h=360", "-maxrate:v:0", "900k", "-b:a:0", "64k",
		"-filter:v:1", "scale=w=854:h=480", "-maxrate:v:1", "1600k", "-b:a:1", "128k",
		"-filter:v:2", "scale=w=1280:h=720", "-maxrate:v:2", "4400k", "-b:a:2", "128k",
		"-filter:v:3", "scale=w=1920:h=1080", "-maxrate:v:3", "7400k", "-b:a:3", "192k",
		"-filter:v:4", "scale=w=2560:h=1440", "-maxrate:v:4", "10000k", "-b:a:4", "192k",
		"-filter:v:5", "scale=w=3840:h=2160", "-maxrate:v:5", "18000k", "-b:a:5", "192k",
		"-var_stream_map", "v:0,a:0,name:360p v:1,a:1,name:480p v:2,a:2,name:720p "+
			"v:3,a:3,name:1080p v:4,a:4,name:1440p v:5,a:5,name:2160p",
		"-hls_list_size", "0", "-f", "hls", "-hls_time", "5", "-hls_playlist_type", "vod",
		"-master_pl_name", m.OutPath+"-pl.m3u8", m.OutPath+"-%v.m3u8")

	m.Err = cmd.Run()
	os.Chdir("../")

	if m.Err != nil {
		os.Remove(m.OutPath)
		p.Send(ThrowErrorMsg{})
		return
	}
	p.Send(CompleteMsg{})
}
