import { deleteSectionSchema, updateSectionSchema } from "@/contracts/course";
import grpcStatusToHttp from "@/lib/grpc-status-to-http";
import validateToken from "@/lib/validate-token";
import CoursesServiceGrpc from "@/services/courses";
import { NextResponse } from "next/server";

export async function PUT(
  request: Request,
  { params }: { params: { id: string } }
) {
  const data = await request.json();
  const user = validateToken();

  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const validatedData = updateSectionSchema.safeParse({
    user_id: user.id,
    id: params.id,
    name: data.name,
    description: data.description,
  });
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.UpdateSection(validatedData.data);

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

export async function DELETE(
  _: Request,
  { params }: { params: { id: string } }
) {
  const user = validateToken();

  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const validatedData = deleteSectionSchema.safeParse({
    user_id: user.id,
    id: params.id,
  });
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.DeleteSection(validatedData.data);

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
