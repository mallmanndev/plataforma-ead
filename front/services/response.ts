export type TResponse<T> = {
  error?: { code: number; message: string };
  response?: T;
};
