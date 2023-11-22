import { nextAuthOptions } from "@/app/api/auth/[...nextauth]/route";
import VideoPlayer from "@/components/video-player";
import { Item } from "@/types/course";
import { Video } from "@/types/video";
import { Metadata } from "next";
import { getServerSession } from "next-auth/next";

const getItem = async (id: string): Promise<Item> => {
  const session = await getServerSession(nextAuthOptions);

  const res = await fetch(`${process.env.SERVER_HOST}/itens/${id}`, {
    headers: { Authorization: `Bearer ${session?.token}` },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch item!");
  }

  return res.json();
};

const getVideo = async (id: string): Promise<Video> => {
  const session = await getServerSession(nextAuthOptions);

  const res = await fetch(`${process.env.SERVER_HOST}/videos/${id}`, {
    headers: { Authorization: `Bearer ${session?.token}` },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch video!");
  }

  return res.json();
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
          source={`${process.env.NEXT_PUBLIC_SERVER_HOST}/api/${video.url}`}
          qualities={video.resolutions.map((item) => parseInt(item.resolution))}
        />
      </div>
      <h1 className="text-3xl font-bold tracking-tight mt-6">{item.title}</h1>
      <h2 className="text-xl text-justify mt-6">{item.description}</h2>

      <div className="mt-12"></div>
    </div>
  );
}
