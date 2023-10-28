import VideoPlayer from "@/components/video-player";
import { Metadata } from "next";
import Sections from "./components/sections";
import { env } from "process";

const getCourse = async (id: string) => {
  const res = await fetch(`${env.SERVER_HOST}/api/course/${id}`, {
    cache: "no-cache",
  });

  if (!res.ok) throw new Error("Failed to fetch course!");

  return res.json();
};

const getVideo = async (id: string) => {
  const res = await fetch(`${env.SERVER_HOST}/api/video/${id}`, {
    cache: "no-cache",
  });

  if (!res.ok) throw new Error("Failed to fetch video!");

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

const getItem = (course: any, itemId?: string) => {
  if (!itemId) return course.sections[0].itens[0].videoId;

  for (const sections of course.sections) {
    for (const item of sections.itens) {
      if (item.id === itemId) return item.videoId;
    }
  }

  return null;
};

export default async function CoursePage({
  params,
  searchParams,
}: {
  params: { id: string };
  searchParams: { item: string };
}) {
  const course = await getCourse(params.id);
  const item = getItem(course, searchParams.item);

  const video = await getVideo(item.videoId);

  return (
    <div>
      <h2 className="text-3xl font-bold tracking-tight mt-6">{course.name}</h2>

      <div className="flex flex-wrap mt-8">
        <div className="w-full md:w-2/3 lg:w-3/4">
          <VideoPlayer
            source="http://localhost:3002/api/videos/11851ead-d665-4c54-a481-bc5a44b2e39b/playlist.m3u8"
            qualities={[1080, 720, 480]}
          />
        </div>

        <div className="w-full md:w-1/3 lg:w-1/4">
          <h1 className="text-2xl mb-4 px-4 font-bold">Conteúdo do curso</h1>

          <div className="px-4">
            <Sections sections={course.sections} />
          </div>
        </div>
      </div>
    </div>
  );
}
