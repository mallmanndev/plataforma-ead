"use client";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import formatDate from "@/utils/formatDate";
import { useToast } from "@/components/ui/use-toast";
import { useEffect } from "react";
import useGetCourse from "@/hooks/get-course";
import EmptySections from "../components/empty-sections";
import SectionOptions from "../components/section-options";

export default function SectionsTable({ courseId }: { courseId: string }) {
  const { toast } = useToast();
  const { loading, course, error, refetch } = useGetCourse(courseId);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível buscar as seções",
        description: error,
      });
    }
  }, [error]);

  if (loading) {
    return <p>Carregando seções...</p>;
  }

  if (!course || course.sections.length < 1) {
    return <EmptySections courseId={courseId} />;
  }

  return (
    <Table>
      <TableCaption>Gerencie seus cursos</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Nome</TableHead>
          <TableHead>Descrição</TableHead>
          <TableHead>Videos</TableHead>
          <TableHead className="w-[150px]">Criado em</TableHead>
          <TableHead className="w-[50px]"></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {course.sections.map((section: any) => (
          <TableRow key={section.id}>
            <TableCell className="font-medium">{section.name}</TableCell>
            <TableCell>{section.description}</TableCell>
            <TableCell>{section.itens.length}</TableCell>
            <TableCell>{formatDate(section.createdAt)}</TableCell>
            <TableCell>
              <SectionOptions id={section.id} onDelete={refetch} />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
