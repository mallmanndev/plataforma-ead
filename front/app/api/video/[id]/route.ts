import grpcStatusToHttp from "@/lib/grpc-status-to-http";
import FilesServiceGrpc from "@/services/files";
import { NextResponse } from "next/server";

export async function GET(_: Request, { params }: { params: { id: string } }) {
  const service = new FilesServiceGrpc();
  const { error, response } = await service.GetVideo(params.id);

  if (error) {
    return NextResponse.json(
      { message: error.message },
      { status: grpcStatusToHttp(error.code) }
    );
  }
  if (!response) {
    return NextResponse.json(
      { message: "Não foi possível logar!" },
      { status: 400 }
    );
  }

  return NextResponse.json(response);
}
