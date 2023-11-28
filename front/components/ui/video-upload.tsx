"use client";

import React, { ChangeEvent, useCallback, useRef, useState } from "react";
import { Progress } from "@/components/ui/progress";
import { Button } from "./button";
import { X } from "lucide-react";

const MEGABYTES = 1048576;

type TVideoUpload = {
  uploadProgress: number;
  uploadCompleted: boolean;
  onFileChange: (file: File) => void;
  onRemoveFile: () => void;
};

const VideoUpload = ({
  uploadProgress,
  onFileChange,
  onRemoveFile,
  uploadCompleted,
}: TVideoUpload) => {
  const fileInputRef = useRef<any>(null);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    if (e.target.files) {
      const file = e.target.files[0];
      setSelectedFile(file);
      onFileChange(file);
    }
  };

  const onRemoveClick = () => {
    setSelectedFile(null);
    onRemoveFile();
  };

  const fileSize = useCallback(() => {
    if (!selectedFile) return null;
    return (selectedFile.size / MEGABYTES).toFixed(2);
  }, [selectedFile]);

  const handleDrop = (e: any) => {
    e.preventDefault(); // Impede o comportamento padrão do navegador

    // Obtenha o arquivo do evento de soltura
    const file = e.dataTransfer.files[0];

    setSelectedFile(file);
    onFileChange(file);
  };

  return (
    <div>
      {!selectedFile && (
        <label
          htmlFor="dropzone-file"
          onDrop={handleDrop}
          onDragOver={(e) => e.preventDefault()}
        >
          <div className="flex h-[150px] shrink-0 items-center justify-center rounded-md border border-dashed">
            <div className="mx-auto flex max-w-[420px] flex-col items-center justify-center text-center">
              <h3 className="mt-4 text-lg font-semibold">
                Nenhum video selecionado
              </h3>
              <p className="mb-4 mt-2 text-sm text-muted-foreground">
                Arraste e solte um arquivo aqui ou clique aqui para escolher um
                arquivo! <br />
                (Formato permitido: <b>.mp4</b>)
              </p>
            </div>
          </div>
          <input
            ref={fileInputRef}
            id="dropzone-file"
            type="file"
            accept="video/mp4"
            className="hidden"
            onChange={handleFileChange}
          />
        </label>
      )}
      {selectedFile && (
        <div className="rounded-lg border bg-card text-card-foreground shadow-sm w-[100%] p-2">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-lg font-bold">{selectedFile.name}</p>
              <p className="text-sm text-zinc-500 font-bold">
                {fileSize()}MB •{" "}
                {!uploadCompleted ? `${uploadProgress}%` : "Upload concluído"}
              </p>
            </div>
            <div>
              <Button variant="outline" size="icon" onClick={onRemoveClick}>
                <X className="h-4 w-4" />
              </Button>
            </div>
          </div>

          {uploadProgress > 0 && !uploadCompleted && (
            <div className="mt-4">
              <Progress value={uploadProgress} />
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default VideoUpload;
