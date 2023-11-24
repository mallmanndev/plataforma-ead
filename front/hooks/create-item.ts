import { TCreateCourseItemData } from "@/contracts/course";
import { Course } from "@/types/course";
import { getSession } from "next-auth/react";
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
      const session = await getSession();

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/itens`,
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

  return { loading, error, course, create };
}
