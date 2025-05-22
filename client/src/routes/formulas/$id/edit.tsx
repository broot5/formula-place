import {
  formulaUpdateSchema,
  type formulaUpdateFormData,
} from "@/schemas/formulaSchema";
import { getFormula, updateFormula } from "@/services/formulaService";
import { zodResolver } from "@hookform/resolvers/zod";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import type { UUID } from "crypto";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";

export const Route = createFileRoute("/formulas/$id/edit")({
  component: EditFormulaPage,
});

function EditFormulaPage() {
  const { id } = Route.useParams();
  const navigate = useNavigate();

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors, isSubmitting, isValid },
  } = useForm({
    mode: "onBlur",
    resolver: zodResolver(formulaUpdateSchema),
  });

  useEffect(() => {
    const fetchFormula = async () => {
      try {
        const data = await getFormula(id as UUID);

        reset({
          title: data.title,
          description: data.description,
          content: data.content,
        });
      } catch (err) {
        setError("Failed to get formula");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchFormula();
  }, [id, reset]);

  const onSubmit = async (data: formulaUpdateFormData) => {
    try {
      await updateFormula(id as UUID, data);
      navigate({ to: `/formulas/${id}` });
    } catch (err) {
      setError("Failed to update formula");
      console.error(err);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      {error && <div>{error}</div>}
      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <label htmlFor="title">Title</label>
          <input type="text" id="title" {...register("title")} />
          {errors.title && <p>{errors.title.message}</p>}
        </div>

        <div>
          <label htmlFor="description">Description</label>
          <input type="text" id="description" {...register("description")} />
          {errors.description && <p>{errors.description.message}</p>}
        </div>

        <div>
          <label htmlFor="content">Content</label>
          <textarea id="content" rows={10} {...register("content")}></textarea>
          {errors.content && <p>{errors.content.message}</p>}
        </div>

        <div>
          <button
            type="button"
            onClick={() => navigate({ to: `/formulas/${id}` })}
          >
            Cancel
          </button>
          <button type="submit" disabled={isSubmitting || !isValid}>
            {isSubmitting ? "Saving..." : "Save"}
          </button>
        </div>
      </form>
    </div>
  );
}
