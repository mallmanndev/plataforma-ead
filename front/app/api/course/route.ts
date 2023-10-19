import {NextResponse} from "next/server";
import {coursesGrpcClient} from "@/configs/grpc-client";
import validateToken from "@/lib/validate-token";

export async function POST(request: Request) {
    const data = await request.json()
    const user = validateToken()

    const req = new Promise((resolve, _) => {
        coursesGrpcClient.Create({
            name: data.name,
            description: data.description,
            instructor: user
        }, (err: any, response: any) => {
            console.log(err, response)
            resolve({err, res: response})
        })
    })
    console.log(req)


    console.log(data)

    return NextResponse.json({ok: true})
}