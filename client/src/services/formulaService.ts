import type { UUID } from "crypto";
import apiClient from "@/lib/axios";
import type { FormulaRequest, FormulaResponse } from "@/types/formula";

export const createFormula = async (
  data: FormulaRequest
): Promise<FormulaResponse> => {
  const response = await apiClient.post<FormulaResponse>("/formulas", data);
  return response.data;
};

export const getFormula = async (id: UUID): Promise<FormulaResponse> => {
  const response = await apiClient.get<FormulaResponse>(`/formulas/${id}`);
  return response.data;
};

export const updateFormula = async (
  id: UUID,
  data: Partial<FormulaRequest>
): Promise<FormulaResponse> => {
  const response = await apiClient.patch<FormulaResponse>(
    `/formulas/${id}`,
    data
  );
  return response.data;
};

export const deleteFormula = async (id: UUID): Promise<void> => {
  await apiClient.delete(`/formulas/${id}`);
};

export const getAllFormulas = async (
  title?: string
): Promise<FormulaResponse[]> => {
  const response = await apiClient.get<FormulaResponse[]>(
    title ? `/formulas?title=${title}` : `/formulas`
  );
  return response.data;
};
