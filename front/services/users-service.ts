import {UsersService} from "@/configs/grpc-client";
import * as grpc from '@grpc/grpc-js'

type TCreateData = {
    name: string
    email: string
    phone: string
    password: string
}

type TCreateResponse = {
    id: string
    name: string
    email: string
    phone: string
    password: string
}

type TLoginData = {
    email: string
    password: string
}

type TResponse<T> = {
    error?: { code: number, message: string },
    response?: T
}

class UsersServiceGrpc {
    // @ts-ignore
    private service: UsersService

    constructor() {
        this.service = new UsersService("service-core:3000", grpc.credentials.createInsecure());
    }

    public async Create(data: TCreateData): Promise<TResponse<TCreateResponse>> {
        return new Promise((resolve, _) => {
            this.service.Create(data, (err: any, response: TCreateResponse) => {
                if (err) {
                    return resolve({error: {code: err.code, message: err.details}})
                }
                return resolve({response: response})
            })
        })
    }

    public async Login(data: TLoginData): Promise<TResponse<TCreateResponse>> {
        return new Promise((resolve, _) => {
            this.service.Login(data, (err: any, response: TCreateResponse) => {
                if (err) {
                    return resolve({error: {code: err.code, message: err.details}})
                }
                return resolve({response: response})
            })
        })

    }
}

export default UsersServiceGrpc