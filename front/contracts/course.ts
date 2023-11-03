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

export const createCourseSectionSchema = z.object({
  name: z.string({ required_error }),
  description: z.string({ required_error }),
});
export type TCreateCourseSectionData = z.infer<
  typeof createCourseSectionSchema
>;

export const createSectionSchema = z.object({
  course_id: z.string(),
  user_id: z.string(),
  name: z.string(),
  description: z.string(),
});
export type TCreateSectionData = z.infer<typeof createSectionSchema>;

export const updateSectionSchema = z.object({
  id: z.string(),
  user_id: z.string(),
  name: z.string(),
  description: z.string(),
});
export type TUpdateSectionData = z.infer<typeof updateSectionSchema>;

export const deleteSectionSchema = z.object({
  id: z.string(),
  user_id: z.string(),
});
export type TDeleteSectionData = z.infer<typeof deleteSectionSchema>;

export const createCourseItemSchema = z.object({
  section_id: z.string(),
  user_id: z.string(),
  title: z.string(),
  description: z.string(),
  video_id: z.string(),
});
export type TCreateCourseItemData = z.infer<typeof createCourseItemSchema>;

export const updateCourseItemSchema = z.object({
  id: z.string(),
  user_id: z.string(),
  title: z.string(),
  description: z.string(),
});
export type TUpdateCourseItemData = z.infer<typeof updateCourseItemSchema>;

export const deleteCourseItemSchema = z.object({
  id: z.string(),
  user_id: z.string(),
});
export type TDeleteCourseItemData = z.infer<typeof deleteCourseItemSchema>;
