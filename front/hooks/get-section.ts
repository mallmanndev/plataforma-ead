"use client";

import { Section } from "@/types/course";
import { useEffect, useState } from "react";

type TUseGetSection = {
  loading: boolean;
  error?: string;
  section?: Section;
  refetch: () => void;
};

export default function useGetSection(id: string): TUseGetSection {
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>();
  const [section, setSection] = useState<Section>();

  useEffect(() => {
    refetch();
  }, []);

  const refetch = async () => {
    setLoading(true);

    const fetchData = await fetch(
      `${process.env.NEXT_PUBLIC_SERVER_HOST}/api/sections/${id}`
    );
    if (!fetchData.ok) {
      setError("Não foi possível buscar o curso.");
      return setLoading(false);
    }

    const section = await fetchData.json();
    setSection(section);
    setLoading(false);
  };

  return { error, section, loading, refetch };
}
