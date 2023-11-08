import VideoPlayer from "@/components/video-player";
import CoursesServiceGrpc from "@/services/courses";
import FilesServiceGrpc from "@/services/files";
import { Item } from "@/types/course";
import { Video } from "@/types/video";
import { Metadata } from "next";

const getItem = async (id: string): Promise<Item> => {
  const service = new CoursesServiceGrpc();
  const { error, response } = await service.GetItem({ id });

  if (error || !response) throw new Error("Failed to fetch item!");

  return response;
};

const getVideo = async (id: string): Promise<Video> => {
  const service = new FilesServiceGrpc();
  const { error, response } = await service.GetVideo(id);

  if (error || !response) throw new Error("Failed to fetch video!");

  return response;
};

export async function generateMetadata({
  params,
}: {
  params: { id: string };
}): Promise<Metadata> {
  const item = await getItem(params.id);

  return {
    title: item.title,
    description: item.description,
  };
}

export default async function ItemPage({ params }: { params: { id: string } }) {
  const item = await getItem(params.id);
  const video = await getVideo(item.videoId);

  return (
    <div>
      <div className="mt-8">
        <VideoPlayer
          source={`http://localhost:3002/api/${video.url}`}
          qualities={[1080, 720, 480]}
        />
      </div>
      <h1 className="text-3xl font-bold tracking-tight mt-6">{item.title}</h1>
      <h2 className="text-xl text-justify mt-6">{item.description}</h2>

      <div className="mt-12"></div>
    </div>
  );
}
