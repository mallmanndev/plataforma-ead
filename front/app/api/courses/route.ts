import { NextRequest, NextResponse } from "next/server";
import validateToken from "@/lib/validate-token";
import { createCourseSchema } from "@/contracts/course";
import CoursesServiceGrpc from "@/services/courses";
import grpcStatusToHttp from "@/lib/grpc-status-to-http";

export async function POST(request: Request) {
  const data = await request.json();
  const user = validateToken();
  const validatedData = createCourseSchema.safeParse(data);
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Create({
    ...data,
    instructor: user,
  });

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

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Get({
    id: searchParams.get("id") as string,
    user_id: searchParams.get("user_id") as string,
    visible: searchParams.get("visible") === "1",
  });
  if (error) {
    return NextResponse.json(
      { message: error.message },
      { status: grpcStatusToHttp(error.code) }
    );
  }

  return NextResponse.json(response?.courses);
}
