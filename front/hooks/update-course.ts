import { useState } from "react";

type TCourse = {
  id: string;
  name: string;
  description: string;
};

type TUseUpdateCourse = {
  loading: boolean;
  error: string | null;
  course: TCourse;
  update(data: TCourse): void;
};

export default function useUpdateCourse(): TUseUpdateCourse {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [course, setCourse] = useState<any>(null);

  const update = (data: TCourse) => {
    (async () => {
      setLoading(true);
      const response = await fetch(`/api/courses/${data.id}`, {
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
}
