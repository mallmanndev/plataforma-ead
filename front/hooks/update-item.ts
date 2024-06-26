import { TUpdateCourseItemData } from "@/contracts/course";
import { Course } from "@/types/course";
import { getSession } from "next-auth/react";
import { useState } from "react";

type TUpdateData = Omit<TUpdateCourseItemData, "user_id">;

type TUseUpdateItem = {
  loading: boolean;
  error: string;
  course?: Course;
  update(data: TUpdateData): void;
};

export default function useUpdateItem(): TUseUpdateItem {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>("");
  const [course, setCourse] = useState<Course>();

  const update = (data: TUpdateData) => {
    (async () => {
      setLoading(true);
      const session = await getSession();

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/itens/${data.id}`,
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
}
