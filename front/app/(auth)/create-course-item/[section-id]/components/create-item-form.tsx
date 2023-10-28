"use client";

import VideoUpload from "@/components/ui/video-upload";
import { useState } from "react";

export default function CreateItemForm() {
  const [videoId, setVideoId] = useState(null);
  const [uploadProgress, setUploadProgress] = useState<number>(0);

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
    setVideoId(null);
  };

  return (
    <>
      <VideoUpload
        onFileChange={onFileSelect}
        uploadProgress={uploadProgress}
        onRemoveFile={onRemoveFile}
        uploadCompleted={!!videoId}
      />
    </>
  );
}
