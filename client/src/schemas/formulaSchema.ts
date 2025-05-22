import { z } from "zod";

export const formulaSchema = z.object({
  title: z
    .string()
    .min(1, "Title is required")
    .max(255, "Title must be less than 255 characters"),
  description: z.string().optional(),
  content: z.string().min(1, "Content is required"),
});

export type FormulaFormData = z.infer<typeof formulaSchema>;

export const formulaUpdateSchema = formulaSchema.partial();
export type formulaUpdateFormData = z.infer<typeof formulaUpdateSchema>;
