import {
  createSectionSchema,
  deleteSectionSchema,
  updateSectionSchema,
} from "@/contracts/course";
import grpcStatusToHttp from "@/lib/grpc-status-to-http";
import validateToken from "@/lib/validate-token";
import CoursesServiceGrpc from "@/services/courses";
import { NextResponse } from "next/server";

export async function POST(request: Request) {
  const data = await request.json();
  const user = validateToken();
  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const validatedData = createSectionSchema.safeParse({
    user_id: user.id,
    ...data,
  });
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.CreateSection({
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

export async function PUT(request: Request) {
  const data = await request.json();
  const user = validateToken();

  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const validatedData = updateSectionSchema.safeParse({
    user_id: user.id,
    ...data,
  });
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.UpdateSection({
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

export async function DELETE(request: Request) {
  const data = await request.json();
  const user = validateToken();

  if (!user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const validatedData = deleteSectionSchema.safeParse({
    user_id: user.id,
    ...data,
  });
  if (!validatedData.success) {
    return NextResponse.json({ message: validatedData.error }, { status: 400 });
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.DeleteSection({
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
