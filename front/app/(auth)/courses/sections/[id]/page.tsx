import { nextAuthOptions } from "@/app/api/auth/[...nextauth]/route";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Section } from "@/types/course";
import { Metadata } from "next";
import { getServerSession } from "next-auth/next";

const getSection = async (id: string): Promise<Section> => {
  const session = await getServerSession(nextAuthOptions);

  const res = await fetch(`${process.env.SERVER_HOST}/sections/${id}`, {
    headers: { Authorization: `Bearer ${session?.token}` },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch sections!");
  }

  return res.json();
};

export async function generateMetadata({
  params,
}: {
  params: { id: string };
}): Promise<Metadata> {
  const section = await getSection(params.id);
  return {
    title: section.name,
    description: section.description,
  };
}

export default async function CoursePage({
  params,
}: {
  params: { id: string };
}) {
  const section = await getSection(params.id);

  return (
    <div>
      <h1 className="text-3xl font-bold tracking-tight mt-6">{section.name}</h1>

      <div className="items-start justify-center gap-6 rounded-lg pt-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
        {section.itens.map((item, key) => (
          <Card
            key={item.id}
            className="h-[280px] flex flex-col place-content-between"
          >
            <CardHeader>
              <div className="flex justify-between">
                <div className="ml-2">
                  <CardTitle>{item.title}</CardTitle>
                  <CardDescription className="text-justify">
                    {item.description.slice(0, 150)}
                    {item.description.length > 150 && "..."}
                  </CardDescription>
                </div>
                <span className="text-6xl font-bold">{key + 1}</span>
              </div>
            </CardHeader>
            <CardFooter>
              <Button className="w-full" asChild>
                <a href={`/courses/itens/${item.id}`}>ASSISTIR</a>
              </Button>
            </CardFooter>
          </Card>
        ))}
      </div>
    </div>
  );
}
