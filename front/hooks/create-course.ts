import { TCreateCourseData } from "@/contracts/course";
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

  return { loading, error, course, createCourse };
};

export default useCreateCourse;
