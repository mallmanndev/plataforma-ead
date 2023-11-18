import { coursesClient } from "@/configs/grpc-client";
import { TResponse } from "./response";
import {
  TCreateCourseItemData,
  TCreateSectionData,
  TDeleteCourseItemData,
  TDeleteSectionData,
  TUpdateCourseItemData,
  TUpdateSectionData,
} from "@/contracts/course";
import { Course, Item, Section } from "@/types/course";

type TInstructor = {
  id: string;
  name: string;
  type: string;
};

type TCreateData = {
  name: string;
  description: string;
  instructor: TInstructor;
};

type TUpdateData = {
  id: string;
  name: string;
  description: string;
  instructor: TInstructor;
};

type TCourse = {
  id: string;
  name: string;
  description: string;
};

type TGetCourseFilters = {
  id: string;
  user_id: string;
  visible: boolean;
};

type TDeleteCourseInput = {
  id: string;
  userId: string;
};

class CoursesServiceGrpc {
  private service: any;

  constructor() {
    this.service = coursesClient;
  }

  public async Create(data: TCreateData): Promise<TResponse<TCourse>> {
    return new Promise((resolve) => {
      this.service.Create(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async Update(data: TUpdateData): Promise<TResponse<TCourse>> {
    return new Promise((resolve) => {
      this.service.Update(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async Get(filters: TGetCourseFilters): Promise<TResponse<Course[]>> {
    return new Promise((resolve, _) => {
      this.service.Get(filters, (err: any, response: { courses: Course[] }) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response.courses });
      });
    });
  }

  public async Delete({
    id,
    userId,
  }: TDeleteCourseInput): Promise<TResponse<{ message: string }>> {
    return new Promise((resolve, _) => {
      this.service.Delete(
        { course_id: id, user_id: userId },
        (err: any, response: any) => {
          if (err) {
            return resolve({ error: { code: err.code, message: err.details } });
          }
          return resolve({ response: response });
        }
      );
    });
  }

  public async CreateSection(
    data: TCreateSectionData
  ): Promise<TResponse<TCourse>> {
    return new Promise((resolve, _) => {
      this.service.CreateSection(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async UpdateSection(
    data: TUpdateSectionData
  ): Promise<TResponse<TCourse>> {
    return new Promise((resolve, _) => {
      this.service.UpdateSection(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async DeleteSection(
    data: TDeleteSectionData
  ): Promise<TResponse<TCourse>> {
    return new Promise((resolve, _) => {
      this.service.DeleteSection(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async GetSection(data: { id: string }): Promise<TResponse<Section>> {
    return new Promise((resolve, _) => {
      this.service.GetSection(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async CreateItem(
    data: TCreateCourseItemData
  ): Promise<TResponse<TCourse>> {
    return new Promise((resolve, _) => {
      this.service.CreateItem(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async UpdateItem(
    data: TUpdateCourseItemData
  ): Promise<TResponse<TCourse>> {
    return new Promise((resolve, _) => {
      this.service.UpdateItem(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async DeleteItem(
    data: TDeleteCourseItemData
  ): Promise<TResponse<TCourse>> {
    return new Promise((resolve, _) => {
      this.service.DeleteItem(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async GetItem(data: { id: string }): Promise<TResponse<Item>> {
    return new Promise((resolve, _) => {
      this.service.GetItem(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async MakeVisible(data: {
    id: string;
    user_id: string;
  }): Promise<TResponse<boolean>> {
    return new Promise((resolve, _) => {
      this.service.MakeVisible(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }

  public async MakeInvisible(data: {
    id: string;
    user_id: string;
  }): Promise<TResponse<boolean>> {
    return new Promise((resolve, _) => {
      this.service.MakeInvisible(data, (err: any, response: any) => {
        if (err) {
          return resolve({ error: { code: err.code, message: err.details } });
        }
        return resolve({ response: response });
      });
    });
  }
}

export default CoursesServiceGrpc;
