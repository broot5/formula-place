import type { UUID } from "crypto";
import type { FormulaFormData } from "@/schemas/formulaSchema";

export type FormulaRequest = FormulaFormData;

export interface FormulaResponse {
  id: UUID;
  title: string;
  description: string;
  content: string;
  created_at: Date;
  updated_at: Date;
}
