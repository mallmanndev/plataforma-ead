import { NextResponse } from "next/server";
import { filesGrpcClient } from "@/configs/grpc-client";
import validateToken from "@/lib/validate-token";

export async function POST(request: Request) {
  const data = await request.formData();
  const user = validateToken();
  const file = data.get("file") as unknown as File;

  if(!user){
    return NextResponse.json({error: "Unauthorized"}, {status: 401})
  }

  if (!file) {
    return NextResponse.json({
      error: "Não foi possível fazer o upload do arquivo.",
    });
  }

  const upload = new Promise<any>((resolve, _) => {
    const call = filesGrpcClient.VideoUpload((error: any, stats: any) => {
      if (error) {
        resolve({ response: null, error: error.message });
      }

      resolve({ response: stats, error: null });
    });

    call.write({ info: { type: "mp4", user_id: user.id, size: file.size } });

    const wStream = new WritableStream({
      write: (chunk) => call.write({ chunk: chunk }),
      close: () => call.end(),
    });

    file.stream().pipeTo(wStream);
  });

  const { response, error } = await upload;
  if (error) {
    return NextResponse.json(error);
  }

  return NextResponse.json(response);
}
