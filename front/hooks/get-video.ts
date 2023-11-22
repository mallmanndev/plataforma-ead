"use client";

import { Video } from "@/types/video";
import { getSession } from "next-auth/react";
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
    const session = await getSession();

    const fetchData = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/videos/${id}`,
      { headers: { Authorization: `Bearer ${session?.token}` } }
    );
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
