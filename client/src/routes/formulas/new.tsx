import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import { createFormula } from "@/services/formulaService";
import { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { formulaSchema, type FormulaFormData } from "@/schemas/formulaSchema";

export const Route = createFileRoute("/formulas/new")({
  component: NewFormulaPage,
});

function NewFormulaPage() {
  const navigate = useNavigate();
  const [submitError, setSubmitError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting, isValid },
  } = useForm({
    mode: "onBlur",
    resolver: zodResolver(formulaSchema),
  });

  const onSubmit = async (data: FormulaFormData) => {
    try {
      setSubmitError(null);
      await createFormula(data);
      navigate({ to: "/formulas" });
    } catch (error) {
      console.error(error);
      setSubmitError("Failed to create formula. Please try again.");
    }
  };

  return (
    <div>
      {submitError && <div>{submitError}</div>}
      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <label htmlFor="title">Title</label>
          <input type="text" id="title" {...register("title")} />
          {errors.title && <p>{errors.title.message}</p>}
        </div>

        <div>
          <label htmlFor="content">Content</label>
          <textarea id="content" rows={10} {...register("content")}></textarea>
          {errors.content && <p>{errors.content.message}</p>}
        </div>

        <div>
          <label htmlFor="description">Description</label>
          <input type="text" id="description" {...register("description")} />
          {errors.description && <p>{errors.description.message}</p>}
        </div>

        <button type="submit" disabled={isSubmitting || !isValid}>
          {isSubmitting ? "Saving..." : "Save"}
        </button>
      </form>
    </div>
  );
}
