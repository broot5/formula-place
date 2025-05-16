import type { UUID } from "crypto";

export interface CreateFormulaRequest {
  title: string;
  content: string;
  description?: string;
}

export interface UpdateFormulaRequest {
  title?: string;
  content?: string;
  description?: string;
}

export interface FormulaResponse {
  id: UUID;
  title: string;
  content: string;
  description: string;
  created_at: Date;
  updated_at: Date;
}
