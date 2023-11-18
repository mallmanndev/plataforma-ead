import grpcStatusToHttp from "@/lib/grpc-status-to-http";
import validateToken from "@/lib/validate-token";
import CoursesServiceGrpc from "@/services/courses";
import { NextResponse } from "next/server";

export async function PATCH(
  request: Request,
  { params }: { params: { id: string } }
) {
  const data = await request.json();
  const user = validateToken();

  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const service = new CoursesServiceGrpc();

  let error, response;
  if (data.visibility === "public") {
    const req = await service.MakeVisible({
      id: params.id,
      user_id: user.id,
    });
    error = req.error;
    response = req.response;
  } else if (data.visibility === "private") {
    const req = await service.MakeInvisible({
      id: params.id,
      user_id: user.id,
    });
    error = req.error;
    response = req.response;
  }

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
