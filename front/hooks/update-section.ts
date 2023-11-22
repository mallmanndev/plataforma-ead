import { TUpdateSectionData } from "@/contracts/course";
import { Section } from "@/types/course";
import { getSession } from "next-auth/react";
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
      const session = await getSession();

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/sections/${data.id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${session?.token}`,
          },
          body: JSON.stringify(data),
        }
      );

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
