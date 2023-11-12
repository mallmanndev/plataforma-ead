"use client";

import ItemForm from "@/components/forms/item-form";
import { toast } from "@/components/ui/use-toast";
import useUpdateItem from "@/hooks/update-item";
import { Course } from "@/types/course";
import { useRouter } from "next/navigation";
import { useCallback, useEffect } from "react";

export default function UpdateItemForm({
  id,
  initialData,
}: {
  id: string;
  initialData: {
    title: string;
    description: string;
  };
}) {
  const { push } = useRouter();
  const { loading, error, course, update } = useUpdateItem();

  const getSectionId = useCallback((course: Course, itemId: string): string => {
    for (const section of course.sections) {
      for (const item of section.itens) {
        if (item.id === itemId) {
          return section.id;
        }
      }
    }
    return "";
  }, []);

  useEffect(() => {
    if (course) {
      toast({ title: "Item criado com sucesso." });
      const sectionId = getSectionId(course, id);
      push(`/manage-itens/${sectionId}`);
    }
  }, [push, getSectionId, course, id]);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível alterar o item.",
        description: error,
      });
    }
  }, [error]);

  return (
    <ItemForm
      loading={loading}
      error={error}
      buttonText="Alterar item"
      defaultValues={initialData}
      onSubmit={(data) => update({ ...data, id })}
    />
  );
}
