import { TCreateCourseItemData, TCreateSectionData } from "@/contracts/course";
import { Course } from "@/types/course";
import { useState } from "react";

type TUseCreateItem = {
  loading: boolean;
  error: string;
  course: Course;
  create(data: Omit<TCreateCourseItemData, "user_id">): void;
};

export default function useCreateItem(): TUseCreateItem {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>("");
  const [course, setCourse] = useState<any>(null);

  const create = (data: Omit<TCreateCourseItemData, "user_id">) => {
    (async () => {
      setLoading(true);
      const response = await fetch(`/api/itens`, {
        method: "POST",
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

  return { loading, error, course, create };
}
