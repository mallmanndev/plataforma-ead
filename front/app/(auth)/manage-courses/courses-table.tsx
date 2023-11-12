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
import useGetCourses from "@/hooks/get-courses";
import formatDate from "@/utils/formatDate";
import { EmptyCourses } from "./components/empty-courses";
import { useToast } from "@/components/ui/use-toast";
import { useEffect } from "react";
import CourseOptions from "./components/course-options";
import { Course } from "@/types/course";

const videos = (course: Course) =>
  course.sections.reduce((acc, cur) => {
    if (cur.itens) acc += cur.itens.length;
    return acc;
  }, 0);

export default function CoursesTable({ userId }: { userId: string }) {
  const { toast } = useToast();
  const { loading, courses, error, refetch } = useGetCourses({
    user_id: userId,
  });

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível buscar os cursos",
        description: error,
      });
    }
  }, [toast, error]);

  if (loading) {
    return <p>Carregando cursos...</p>;
  }

  if (courses.length < 1) {
    return <EmptyCourses />;
  }

  return (
    <Table>
      <TableCaption>Gerencie seus cursos</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Nome</TableHead>
          <TableHead>Descrição</TableHead>
          <TableHead>Seções</TableHead>
          <TableHead>Videos</TableHead>
          <TableHead className="w-[150px]">Criado em</TableHead>
          <TableHead className="w-[150px]">Visibilidade</TableHead>
          <TableHead className="w-[50px]"></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {courses.map((course: any) => (
          <TableRow key={course.id}>
            <TableCell className="font-medium">{course.name}</TableCell>
            <TableCell>{course.description}</TableCell>
            <TableCell>{course.sections.length}</TableCell>
            <TableCell>{videos(course)}</TableCell>
            <TableCell>{formatDate(course.createdAt)}</TableCell>
            <TableCell className="text-right">
              <Select>
                <SelectTrigger className="h-8">
                  <SelectValue placeholder="Visibilidade" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Visibilidade</SelectLabel>
                    <SelectItem value="public">Público</SelectItem>
                    <SelectItem value="private">Privado</SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </TableCell>
            <TableCell>
              <CourseOptions id={course.id} onDelete={refetch} />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
