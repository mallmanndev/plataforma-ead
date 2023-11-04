"use client";

import { Course } from "@/types/course";
import { useEffect, useState } from "react";

type TUseGetCourses = {
  loading: boolean;
  error?: string;
  course?: Course;
  refetch: () => void;
};

const useGetCourse = (id: string): TUseGetCourses => {
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>();
  const [course, setCourse] = useState<Course>();

  useEffect(() => {
    refetch();
  }, []);

  const refetch = async () => {
    setLoading(true);

    const fetchData = await fetch(
      `${process.env.NEXT_PUBLIC_SERVER_HOST}/api/courses/${id}`
    );
    if (!fetchData.ok) {
      setError("Não foi possível buscar o curso.");
      return setLoading(false);
    }

    const courses = await fetchData.json();

    console.log(courses)

    setCourse(courses);
    setLoading(false);
  };

  return { error, course, loading, refetch };
};

export default useGetCourse;
