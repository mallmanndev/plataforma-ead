import { z } from "zod";

const required_error = "Este campo é obrigatório.";
export const createCourseSchema = z.object({
  name: z.string({ required_error }),
  description: z.string({ required_error }),
});

export type TCreateCourseData = z.infer<typeof createCourseSchema>;

export const updateCourseSchema = z.object({
  name: z.string({ required_error }),
  description: z.string({ required_error }),
});

export type TUpdateCourseData = z.infer<typeof updateCourseSchema>;
