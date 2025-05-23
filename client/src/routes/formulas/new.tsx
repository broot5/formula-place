import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import { createFormula } from "@/services/formulaService";
import { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { formulaSchema, type FormulaFormData } from "@/schemas/formulaSchema";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Button } from "@/components/ui/button";

export const Route = createFileRoute("/formulas/new")({
  component: NewFormulaPage,
});

function NewFormulaPage() {
  const navigate = useNavigate();

  const [error, setError] = useState<string | null>(null);

  const form = useForm({
    mode: "onBlur",
    resolver: zodResolver(formulaSchema),
  });

  const onSubmit = async (data: FormulaFormData) => {
    try {
      await createFormula(data);
      navigate({ to: "/formulas" });
    } catch (err) {
      setError("Failed to create formula");
      console.error(err);
    }
  };

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
                <FormDescription>
                  This is formula's description.
                </FormDescription>
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
          <Button
            type="submit"
            disabled={form.formState.isSubmitting || !form.formState.isValid}
          >
            {form.formState.isSubmitting ? "Submitting..." : "Submit"}
          </Button>
        </form>
      </Form>
    </div>
  );
}
