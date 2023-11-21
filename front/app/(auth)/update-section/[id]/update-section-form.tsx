"use client";

import SectionForm from "@/components/forms/section-form";
import { toast } from "@/components/ui/use-toast";
import useUpdateSection from "@/hooks/update-section";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function UpdateSectionForm({
  id,
  initialData,
}: {
  id: string;
  initialData: {
    name: string;
    description: string;
  };
}) {
  const { push } = useRouter();
  const { loading, error, course, update } = useUpdateSection();

  useEffect(() => {
    if (course) {
      toast({ title: "Seção alterada com sucesso." });
      push(`/manage-sections/${course.id}`);
    }
  }, [push, course]);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível alterar a seção",
        description: error,
      });
    }
  }, [error]);

  return (
    <SectionForm
      loading={loading}
      error={error}
      buttonText="Alterar seção"
      defaultValues={initialData}
      onSubmit={(data) => update({ ...data, id })}
    />
  );
}
