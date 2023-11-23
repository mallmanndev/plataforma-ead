"use client";

import CourseForm from "@/components/forms/course-form";
import { useToast } from "@/components/ui/use-toast";
import useUpdateCourse from "@/hooks/update-course";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

type UpdateCourseForm = {
  id: string;
  initialData: {
    name: string;
    description: string;
    discord_url: string;
  };
};

export default function UpdateCourseForm({
  id,
  initialData,
}: UpdateCourseForm) {
  const { push } = useRouter();
  const { toast } = useToast();
  const { loading, error, course, update } = useUpdateCourse();

  useEffect(() => {
    if (course) {
      toast({ title: "Curso alterado com sucesso" });
      push("/manage-courses");
    }
  }, [push, toast, course]);

  const handleSubmit = (data: any) => {
    update({ id, ...data });
  };

  return (
    <CourseForm
      loading={loading}
      error={error}
      buttonText="Alterar curso"
      defaultValues={initialData}
      onSubmit={handleSubmit}
    />
  );
}
