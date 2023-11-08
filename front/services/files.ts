import { filesGrpcClient } from "@/configs/grpc-client";
import { TResponse } from "./response";
import { Video } from "@/types/video";

export default class FilesServiceGrpc {
  private service: any;

  constructor() {
    this.service = filesGrpcClient;
  }

  public async GetVideo(id: string): Promise<TResponse<Video>> {
    return new Promise((resolve) => {
      this.service.GetVideo({ id }, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }
}
