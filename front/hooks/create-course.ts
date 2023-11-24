import { TCreateCourseData } from "@/contracts/course";
import { getSession } from "next-auth/react";
import { useState } from "react";

type TCourse = {
  id: string;
  name: string;
  description: string;
};

type TUseCreateCourse = {
  loading: boolean;
  error: string | null;
  course: TCourse;
  createCourse(data: TCreateCourseData): void;
};

const useCreateCourse = (): TUseCreateCourse => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [course, setCourse] = useState<any>(null);

  const createCourse = (data: TCreateCourseData) => {
    (async () => {
      const session = await getSession();

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/courses`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${session?.token}`,
          },
          body: JSON.stringify(data),
        }
      );

      const json = await response.json();

      if (response.ok) {
        setCourse(json);
      } else {
        setError(json.message);
      }
    })();
  };

  return { loading, error, course, createCourse };
};

export default useCreateCourse;
