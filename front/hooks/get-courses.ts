"use client";

import { Course } from "@/types/course";
import { getSession, useSession } from "next-auth/react";
import { useEffect, useState } from "react";

type TUseGetCoursesFilter = {
  user_id: string;
};

type TUseGetCourses = {
  loading: boolean;
  error?: string;
  courses: Course[];
  refetch: () => void;
};

const useGetCourses = (filters: TUseGetCoursesFilter): TUseGetCourses => {
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>();
  const [courses, setCourses] = useState<Course[]>([]);

  useEffect(() => {
    refetch();
  }, []);

  const refetch = async () => {
    const session = await getSession();
    setLoading(true);

    console.log(session);
    const fetchData = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/courses?user_id=${filters.user_id}`,
      { headers: { Authorization: `Bearer ${session?.token}` } }
    );
    if (!fetchData.ok) {
      setError("Não foi possível buscar os cursos.");
      return setLoading(false);
    }

    const courses = await fetchData.json();

    setCourses(courses);
    setLoading(false);
  };

  return { error, courses, loading, refetch };
};

export default useGetCourses;
