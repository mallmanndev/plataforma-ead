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
import CoursesServiceGrpc from "@/services/courses";
import { Course } from "@/types/course";

export const metadata = {
  title: "Home",
};

const getCourses = async () => {
  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Get({
    id: "",
    user_id: "",
    visible: true,
  });

  if (error) {
    throw new Error("Failed to fetch course!");
  }

  return response;
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

      <div className="items-start justify-center gap-6 rounded-lg pt-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
        {courses?.map((course) => (
          <Card>
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
    </>
  );
}
