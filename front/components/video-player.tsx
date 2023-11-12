"use client";

import React, { useEffect, useRef, useState } from "react";
import Plyr from "plyr";
import "plyr/dist/plyr.css";
import Hls from "hls.js";

type TVideoPlayer = {
  source: string;
  qualities: number[];
};

const quality_indexes: any = {
  1080: 2,
  720: 1,
  480: 0,
  5000: -1,
};

const controls = [
  "play-large",
  "restart",
  "rewind",
  "play",
  "fast-forward",
  "progress",
  "current-time",
  "duration",
  "mute",
  "volume",
  "settings",
  "pip",
  "airplay",
  "fullscreen",
];

export default function VideoPlayer({ source, qualities }: TVideoPlayer) {
  const videoRef = useRef<any>(null);
  const hlsRef = useRef<Hls | null>(null);
  const [quality, setQuality] = useState(-1);

  qualities.push(5000);

  useEffect(() => {
    if (!Hls.isSupported()) {
      videoRef.current.src = source;
      return;
    }

    new Plyr(videoRef.current, {
      controls,
      quality: {
        default: 5000,
        options: qualities.sort((a, b) => (a < b ? 1 : -1)),
        forced: true,
        onChange: (e) => {
          setQuality(quality_indexes[e]);
        },
      },
      i18n: {
        qualityLabel: { 5000: "Auto", 1080: "1080p", 720: "720p", 480: "480p" },
      },
    });

    hlsRef.current = new Hls();
    hlsRef.current.loadSource(source);
    hlsRef.current.attachMedia(videoRef.current);

    return () => {
      hlsRef.current?.destroy();
    };
  }, [qualities, source]);

  useEffect(() => {
    if (hlsRef.current) hlsRef.current.currentLevel = quality;
  }, [quality]);

  return <video width={"100%"} ref={videoRef} controls />;
}
