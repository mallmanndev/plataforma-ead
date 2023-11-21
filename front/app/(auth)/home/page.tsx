import { nextAuthOptions } from "@/app/api/auth/[...nextauth]/route";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Course } from "@/types/course";
import { getServerSession } from "next-auth/next";

export const metadata = {
  title: "Home",
};

const getCourses = async () => {
  const session = await getServerSession(nextAuthOptions);

  const req = await fetch(`${process.env.SERVER_HOST}/courses?visible=1`, {
    headers: { Authorization: `Bearer ${session?.token}` },
  });

  if (!req.ok) {
    throw new Error("Failed to fetch courses!");
  }

  return req.json();
};

export default async function Home() {
  const courses = await getCourses();

  const videos = (course: Course) =>
    course.sections.reduce((acc, cur) => {
      if (cur.itens) acc += cur.itens.length;
      return acc;
    }, 0);

  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight pt-4">Cursos</h2>

      {courses?.length === 0 && (
        <div className="flex items-center w-full h-24">
          <p className="text-lg font-bold text-gray-500">
            Nenhum curso disponível
          </p>
        </div>
      )}

      {courses && courses.length > 0 && (
        <div className="items-start justify-center gap-6 rounded-lg pt-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
          {courses?.map((course: any) => (
            <Card key={course.id}>
              <CardHeader>
                <CardTitle>{course.name}</CardTitle>
                <CardDescription>{course.description}</CardDescription>
              </CardHeader>
              <CardContent>
                <Badge variant="outline">{`${course.sections.length} Seções`}</Badge>
                <Badge variant="outline">{`${videos(course)} Videos`}</Badge>
              </CardContent>
              <CardFooter>
                <Button className="w-full" asChild>
                  <a href={`/courses/${course.id}`}>ACESSAR</a>
                </Button>
              </CardFooter>
            </Card>
          ))}
        </div>
      )}
    </>
  );
}
