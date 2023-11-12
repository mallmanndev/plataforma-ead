"use client";

import SectionForm from "@/components/forms/section-form";
import { toast } from "@/components/ui/use-toast";
import useCreateSection from "@/hooks/create-section";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function CreateSectionForm({ courseId }: { courseId: string }) {
  const { push } = useRouter();
  const { loading, error, course, createSection } = useCreateSection();

  useEffect(() => {
    if (course) {
      toast({ title: "Seção criada com sucesso." });
      push(`/manage-sections/${course.id}`);
    }
  }, [push, course]);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível criar seção",
        description: error,
      });
    }
  }, [error]);

  const initialData = {
    name: "",
    description: "",
  };

  return (
    <SectionForm
      loading={loading}
      error={error}
      buttonText="Criar seção"
      defaultValues={initialData}
      onSubmit={(data) => {
        createSection({ ...data, course_id: courseId });
      }}
    />
  );
}
