"use client";

import ItemForm from "@/components/forms/item-form";
import { toast } from "@/components/ui/use-toast";
import VideoUpload from "@/components/ui/video-upload";
import useCreateItem from "@/hooks/create-item";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function CreateItemForm({ sectionId }: { sectionId: string }) {
  const { push } = useRouter();
  const { loading, error, course, create } = useCreateItem();
  const [videoId, setVideoId] = useState<string>();
  const [uploadProgress, setUploadProgress] = useState<number>(0);

  useEffect(() => {
    if (course) {
      toast({ title: "Item criado com sucesso." });
      push(`/manage-itens/${sectionId}`);
    }
  }, [course]);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível criar o item.",
        description: error,
      });
    }
  }, [error]);

  const onFileSelect = async (file: File) => {
    const data = new FormData();
    data.set("file", file);

    const request = new XMLHttpRequest();
    request.upload.onprogress = (event: any) => {
      if (event.lengthComputable) {
        setUploadProgress(Math.floor((event.loaded / event.total) * 100));
      }
    };

    request.onload = () => {
      const json = JSON.parse(request.response);
      setVideoId(json.id);
    };

    request.open("POST", "/api/video-upload", true);
    request.send(data);
  };

  const onRemoveFile = () => {
    setUploadProgress(0);
    setVideoId("");
  };

  return (
    <div>
      <div className="mb-6">
        <VideoUpload
          onFileChange={onFileSelect}
          uploadProgress={uploadProgress}
          onRemoveFile={onRemoveFile}
          uploadCompleted={!!videoId}
        />
      </div>

      {videoId && (
        <ItemForm
          loading={loading}
          error={error}
          buttonText="Criar item"
          defaultValues={{ title: "", description: "" }}
          onSubmit={(data) =>
            create({ ...data, video_id: videoId, section_id: sectionId })
          }
        />
      )}
    </div>
  );
}
