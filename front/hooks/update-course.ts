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
  update(data: any): void;
};

export default function useUpdateCourse(): TUseUpdateCourse {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [course, setCourse] = useState<any>(null);

  const update = () => {
    (async () => {
      const response = await fetch(`${process.env.NEXT_PUBLIC_SERVER_HOST}/api/courses`, {
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
    })();
  };

  return { loading, error, course, update };
}
