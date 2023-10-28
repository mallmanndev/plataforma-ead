import { usersService } from "@/configs/grpc-client";
import { TResponse } from "./response";

type TCreateData = {
  name: string;
  email: string;
  phone: string;
  password: string;
};

type TCreateResponse = {
  id: string;
  name: string;
  email: string;
  phone: string;
  password: string;
};

type TLoginData = {
  email: string;
  password: string;
};

class UsersServiceGrpc {
  private service: any;

  constructor() {
    this.service = usersService;
  }

  public async Create(data: TCreateData): Promise<TResponse<TCreateResponse>> {
    return new Promise((resolve, _) => {
      this.service.Create(data, (err: any, response: TCreateResponse) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async Login(data: TLoginData): Promise<TResponse<TCreateResponse>> {
    return new Promise((resolve, _) => {
      this.service.Login(data, (err: any, response: TCreateResponse) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }
}

export default UsersServiceGrpc;
