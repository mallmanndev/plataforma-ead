"use client";

import { Course } from "@/types/course";
import { getSession } from "next-auth/react";
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
    const session = await getSession();

    const fetchData = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/courses/${id}`,
      { headers: { Authorization: `Bearer ${session?.token}` } }
    );
    if (!fetchData.ok) {
      setError("Não foi possível buscar o curso.");
      return setLoading(false);
    }

    const courses = await fetchData.json();

    setCourse(courses);
    setLoading(false);
  };

  return { error, course, loading, refetch };
};

export default useGetCourse;
