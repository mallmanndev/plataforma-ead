import { TUpdateSectionData } from "@/contracts/course";
import { Section } from "@/types/course";
import { useState } from "react";

type TUseUpdateSection = {
  loading: boolean;
  error: string | null;
  course: Section;
  update(data: Omit<TUpdateSectionData, "user_id">): void;
};

const useUpdateSection = (): TUseUpdateSection => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [course, setCourse] = useState<any>(null);

  const update = (data: Omit<TUpdateSectionData, "user_id">) => {
    (async () => {
      setLoading(true);
      const response = await fetch(`/api/sections/${data.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        const course = await response.json();
        setCourse(course);
      } else {
        const errorData = await response.json();
        setError(errorData.message);
      }
      setLoading(false);
    })();
  };

  return { loading, error, course, update };
};

export default useUpdateSection;
