import { TCreateSectionData } from "@/contracts/course";
import { Section } from "@/types/course";
import { getSession } from "next-auth/react";
import { useState } from "react";

type TUseCreateSection = {
  loading: boolean;
  error: string | null;
  course: Section;
  createSection(data: Omit<TCreateSectionData, "user_id">): void;
};

const useCreateSection = (): TUseCreateSection => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [course, setCourse] = useState<any>(null);

  const createSection = (data: Omit<TCreateSectionData, "user_id">) => {
    (async () => {
      setLoading(true);
      const session = await getSession();

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/sections`,
        {
          method: "POST",
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

  return { loading, error, course, createSection };
};

export default useCreateSection;
