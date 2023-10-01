import {NextResponse} from "next/server";
import UsersServiceGrpc from "@/services/users-service";
import grpcStatusToHttp from "@/lib/grpc-status-to-http";

export async function GET(request: Request) {
    return NextResponse.json({"success": true})
}

export async function POST(request: Request) {
    const body = await request.json()
    const service = new UsersServiceGrpc()
    const {error, response} = await service.Create(body)
    if (error) {
        return NextResponse.json({message: error.message}, {status: grpcStatusToHttp(error.code)})
    }
    return NextResponse.json(response)
}