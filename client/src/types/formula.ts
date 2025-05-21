import type { UUID } from "crypto";
import type {
  FormulaFormData,
  formulaUpdateFormData,
} from "@/schemas/formulaSchema";

export type CreateFormulaRequest = FormulaFormData;
export type UpdateFormulaRequest = formulaUpdateFormData;

export interface FormulaResponse {
  id: UUID;
  title: string;
  content: string;
  description: string;
  created_at: Date;
  updated_at: Date;
}
