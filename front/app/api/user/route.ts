import {NextResponse} from "next/server";

export async function GET(request: Request) {
    return NextResponse.json({"success": true})
}

export async function POST(request: Request) {
    const body = await request.json()
    console.log(body)

    return NextResponse.json({"success": true})
}