import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import UsersServiceGrpc from "@/services/users-service";
import grpcStatusToHttp from "@/lib/grpc-status-to-http";
import jwt from "jsonwebtoken";

export async function POST(request: Request) {
  const data = await request.json();
  const service = new UsersServiceGrpc();
  const { error, response } = await service.Login(data);
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
  const token = jwt.sign(response, process.env.JWT_SECRET as string, {
    expiresIn: "1d",
  });
  cookies().set("token", token);
  return NextResponse.json({ ok: true });
}
