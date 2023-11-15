"use client";

import { Video } from "@/types/video";
import { useEffect, useState } from "react";

type TUseGetVideo = {
  loading: boolean;
  error?: string;
  video?: Video;
  refetch: () => void;
};

export default function useGetVideo(id: string): TUseGetVideo {
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>();
  const [video, setVideo] = useState<Video>();

  useEffect(() => {
    refetch();
  }, []);

  const refetch = async () => {
    setLoading(true);

    const fetchData = await fetch(`/api/video/${id}`);
    if (!fetchData.ok) {
      setError("Não foi possível buscar o curso.");
      return setLoading(false);
    }

    const video = await fetchData.json();

    setVideo(video);
    setLoading(false);
  };

  return { error, video, loading, refetch };
}
