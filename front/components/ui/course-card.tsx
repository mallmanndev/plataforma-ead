"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import React from "react";
import { Badge } from "./badge";
import { Course } from "@/types/course";
import { useRouter } from "next/navigation";

type TCourseCardProps = {
  course: Course;
};

export default function CourseCard({ course }: TCourseCardProps) {
  const { push } = useRouter();
  const videos = course.sections.reduce((acc, cur) => {
    if (cur.itens) acc += cur.itens.length;
    return acc;
  }, 0);

  return (
    <Card>
      <CardHeader>
        <CardTitle>{course.name}</CardTitle>
        <CardDescription>{course.description}</CardDescription>
      </CardHeader>
      <CardContent>
        <Badge variant="outline">{`${course.sections.length} Seções`}</Badge>
        <Badge variant="outline">{`${videos} Videos`}</Badge>
      </CardContent>
      <CardFooter>
        <Button
          className="w-full"
          onClick={() => push(`/courses/${course.id}`)}
        >
          ACESSAR
        </Button>
      </CardFooter>
    </Card>
  );
}
