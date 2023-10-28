"use client";

import CourseForm from "@/components/forms/course-form";
import useUpdateCourse from "@/hooks/update-course";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

type UpdateCourseForm = {
  initialData: {
    name: string;
    description: string;
  };
};

export default function UpdateCourseForm({ initialData }: UpdateCourseForm) {
  const { push } = useRouter();
  const { loading, error, course, update } = useUpdateCourse();

  useEffect(() => {
    if (course) push("/my-courses");
  }, [course]);

  return (
    <CourseForm
      loading={loading}
      error={error}
      buttonText="Alterar curso"
      defaultValues={initialData}
      onSubmit={update}
    />
  );
}
