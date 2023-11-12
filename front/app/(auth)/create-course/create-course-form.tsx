"use client";

import CourseForm from "@/components/forms/course-form";
import useCreateCourse from "@/hooks/create-course";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function CreateCourseForm() {
  const { push } = useRouter();
  const { loading, error, course, createCourse } = useCreateCourse();

  useEffect(() => {
    if (course) push("/manage-courses");
  }, [push, course]);

  const initialData = {
    name: "",
    description: "",
  };

  return (
    <CourseForm
      loading={loading}
      error={error}
      buttonText="Criar curso"
      defaultValues={initialData}
      onSubmit={createCourse}
    />
  );
}
