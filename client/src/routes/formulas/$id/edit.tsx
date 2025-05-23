import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
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

  const form = useForm({
    mode: "onBlur",
    resolver: zodResolver(formulaUpdateSchema),
  });

  useEffect(() => {
    const fetchFormula = async () => {
      try {
        const data = await getFormula(id as UUID);

        form.reset({
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
  }, [id, form]);

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
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Title</FormLabel>
                <FormControl>
                  <input type="text" {...field} />
                </FormControl>
                <FormDescription>This is formula's title.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Description</FormLabel>
                <FormControl>
                  <input type="text" {...field} />
                </FormControl>
                <FormDescription>This is formula's description</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="content"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Content</FormLabel>
                <FormControl>
                  <textarea {...field} />
                </FormControl>
                <FormDescription>This is formula's content.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <div>
            <Button
              type="button"
              onClick={() => navigate({ to: `/formulas/${id}` })}
            >
              Cancel
            </Button>
            <Button
              type="submit"
              disabled={form.formState.isSubmitting || !form.formState.isValid}
            >
              {form.formState.isSubmitting ? "Submitting..." : "Submit"}
            </Button>
          </div>
        </form>
      </Form>
    </div>
  );
}
