import { updateCourseSchema } from "@/contracts/course";
import grpcStatusToHttp from "@/lib/grpc-status-to-http";
import validateToken from "@/lib/validate-token";
import CoursesServiceGrpc from "@/services/courses";
import { NextResponse } from "next/server";

export async function GET(_: Request, params: { params: { id: string } }) {
  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Get({
    id: params.params.id || "",
    visible: false,
    instructor_id: "",
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

  return NextResponse.json(response.courses[0]);
}

export async function DELETE(
  _: Request,
  { params }: { params: { id: string } }
) {
  const user = validateToken();
  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Delete({
    id: params.id,
    userId: user?.id as string,
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

export async function PUT(
  request: Request,
  { params }: { params: { id: string } }
) {
  const data = await request.json();
  const user = validateToken();
  const validatedData = updateCourseSchema.safeParse(data);
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Update({
    ...data,
    course_id: params.id,
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
