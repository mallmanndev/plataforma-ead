"use client";

import { Section } from "@/types/course";
import { getSession } from "next-auth/react";
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
    const session = await getSession();

    const fetchData = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/sections/${id}`,
      { headers: { Authorization: `Bearer ${session?.token}` } }
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
