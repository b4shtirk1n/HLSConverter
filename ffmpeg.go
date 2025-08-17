package main

import (
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func (m Model) Convert() {
	os.Mkdir(m.OutPath, 0750)
	if m.Err = os.Chdir(m.OutPath); m.Err != nil {
		p.Send(ThrowErrorMsg{})
		return
	}

	if m.IsGPU {
		cmd = exec.Command("ffmpeg", "-y", "-i",
			m.SelectedFile,
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-c:v", "hevc_nvenc", "-tag:v", "hvc1", "-crf", "26", "-c:a", "aac", "-ar", "48000",
			"-filter:v:0", "scale=w=640:h=360", "-b:v", "1200k", "-minrate:v:0", "1200k", "-b:a:0", "128k",
			"-filter:v:1", "scale=w=854:h=480", "-b:v", "2000k", "-minrate:v:1", "2000k", "-b:a:1", "128k",
			"-filter:v:2", "scale=w=1280:h=720", "-b:v", "6000k", "-minrate:v:2", "6000k", "-b:a:2", "128k",
			"-filter:v:3", "scale=w=1920:h=1080", "-b:v", "10000k", "-minrate:v:3", "10000k", "-b:a:3", "192k",
			"-filter:v:4", "scale=w=2560:h=1440", "-b:v", "18000k", "-minrate:v:4", "18000k", "-b:a:4", "192k",
			"-filter:v:5", "scale=w=3840:h=2160", "-b:v", "25000k", "-minrate:v:5", "25000k", "-b:a:5", "192k",
			"-var_stream_map", "v:0,a:0,name:360p v:1,a:1,name:480p v:2,a:2,name:720p "+
				"v:3,a:3,name:1080p v:4,a:4,name:1440p v:5,a:5,name:2160p",
			"-hls_list_size", "0", "-f", "hls", "-hls_time", "5", "-hls_playlist_type", "vod",
			"-master_pl_name", m.OutPath+"-pl.m3u8", m.OutPath+"-%v.m3u8")
	} else {
		cmd = exec.Command("ffmpeg", "-y", "-i",
			m.SelectedFile,
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-map", "0:v:0", "-map", "0:a:0",
			"-c:v", "libx265", "-crf", "26", "-c:a", "aac", "-ar", "48000",
			"-filter:v:0", "scale=w=640:h=360", "-b:v", "1200k", "-minrate:v:0", "1200k", "-b:a:0", "128k",
			"-filter:v:1", "scale=w=854:h=480", "-b:v", "2000k", "-minrate:v:1", "2000k", "-b:a:1", "128k",
			"-filter:v:2", "scale=w=1280:h=720", "-b:v", "6000k", "-minrate:v:2", "6000k", "-b:a:2", "128k",
			"-filter:v:3", "scale=w=1920:h=1080", "-b:v", "10000k", "-minrate:v:3", "10000k", "-b:a:3", "192k",
			"-filter:v:4", "scale=w=2560:h=1440", "-b:v", "18000k", "-minrate:v:4", "18000k", "-b:a:4", "192k",
			"-filter:v:5", "scale=w=3840:h=2160", "-b:v", "25000k", "-minrate:v:5", "25000k", "-b:a:5", "192k",
			"-var_stream_map", "v:0,a:0,name:360p v:1,a:1,name:480p v:2,a:2,name:720p "+
				"v:3,a:3,name:1080p v:4,a:4,name:1440p v:5,a:5,name:2160p",
			"-hls_list_size", "0", "-f", "hls", "-hls_time", "5", "-hls_playlist_type", "vod",
			"-master_pl_name", m.OutPath+"-pl.m3u8", m.OutPath+"-%v.m3u8")
	}

	m.Err = cmd.Run()
	os.Chdir("../")

	if m.Err != nil {
		os.Remove(m.OutPath)
		p.Send(ThrowErrorMsg{})
		return
	}
	p.Send(CompleteMsg{})
}
