import { filesGrpcClient } from "@/configs/grpc-client";
import { TResponse } from "./response";

type TVideo = {
  resolutions: {
    resolution: string;
    url: string;
  }[];
  id: string;
  type: string;
  status: string;
  size: number;
  createdAt: Date;
  updatedAt: Date;
  url: string;
};

export default class FilesServiceGrpc {
  private service: any;

  constructor() {
    this.service = filesGrpcClient;
  }

  public async GetVideo(id: string): Promise<TResponse<{ courses: TVideo }>> {
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
