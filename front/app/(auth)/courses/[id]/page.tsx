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
import { Metadata } from "next";
import { env } from "process";

const getCourse = async (id: string): Promise<Course> => {
  const res = await fetch(`${process.env.SERVER_HOST}/api/courses/${id}`, {
    cache: "no-cache",
  });

  if (!res.ok) {
    throw new Error("Failed to fetch course!");
  }

  return res.json();
};

export async function generateMetadata({
  params,
}: {
  params: { id: string };
}): Promise<Metadata> {
  const course = await getCourse(params.id);
  return {
    title: course.name,
    description: course.description,
  };
}

export default async function CoursePage({
  params,
}: {
  params: { id: string };
}) {
  const course = await getCourse(params.id);

  return (
    <div>
      <h1 className="text-3xl font-bold tracking-tight mt-6">{course.name}</h1>

      <div className="items-start justify-center gap-6 rounded-lg pt-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
        {course.sections.map((item, key) => (
          <Card
            key={item.id}
            className="h-[280px] flex flex-col place-content-between"
          >
            <CardHeader>
              <div className="flex justify-between">
                <div className="ml-2">
                  <CardTitle>{item.name}</CardTitle>
                  <CardDescription className="text-justify">
                    {item.description}
                  </CardDescription>
                </div>
                <span className="text-6xl font-bold">{key + 1}</span>
              </div>
            </CardHeader>
            <CardFooter className="flex flex-col">
              <div className="w-full mb-3">
                <Badge variant="outline">{`${item.itens.length} Videos`}</Badge>
              </div>

              <Button className="w-full" asChild>
                <a href={`/courses/sections/${item.id}`}>ACESSAR</a>
              </Button>
            </CardFooter>
          </Card>
        ))}
      </div>
    </div>
  );
}
