import { useState } from "react";
import { useForm, useFormState, useWatch } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { type FormulaFormData, formulaSchema } from "@/schemas/formulaSchema";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Eye } from "lucide-react";
import { LatexRenderer } from "@/components/LatexRenderer";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { ErrorState } from "@/components/ErrorState";

interface FormulaFormProps {
  defaultValues?: {
    title: string;
    description: string;
    content: string;
  };
  onSubmit: (data: Partial<FormulaFormData>) => Promise<void>;
  error?: string | null;
  isEditMode?: boolean;
  submitButtonText?: string;
  onCancel?: () => void;
}

export function FormulaForm({
  defaultValues = {
    title: "",
    description: "",
    content: "",
  },
  onSubmit,
  error,
  isEditMode = false,
  submitButtonText = "Create Formula",
  onCancel,
}: FormulaFormProps) {
  const [showPreview, setShowPreview] = useState(true);

  const form = useForm({
    mode: "onBlur",
    resolver: zodResolver(formulaSchema),
    defaultValues,
  });

  const { dirtyFields, isSubmitting, isValid } = useFormState({
    control: form.control,
  });

  const content = useWatch({
    control: form.control,
    name: "content",
    defaultValue: defaultValues.content,
  });

  const handleSubmit = async (data: FormulaFormData) => {
    try {
      if (isEditMode) {
        const dirtyData = Object.keys(dirtyFields).reduce((acc, key) => {
          if (dirtyFields[key as keyof FormulaFormData]) {
            acc[key as keyof FormulaFormData] =
              data[key as keyof FormulaFormData];
          }
          return acc;
        }, {} as Partial<FormulaFormData>);

        if (Object.keys(dirtyData).length > 0) {
          await onSubmit(dirtyData);
        }
      } else {
        await onSubmit(data);
      }
    } catch (error) {
      console.error("Form submission error:", error);
    }
  };

  return (
    <>
      {error && <ErrorState error={error} />}

      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(handleSubmit)}
          className="flex flex-col gap-4"
        >
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Title</FormLabel>
                <FormControl>
                  <Input placeholder="Enter formula title" {...field} />
                </FormControl>
                <FormDescription>
                  A clear, descriptive title for your formula.
                </FormDescription>
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
                  <Input placeholder="Enter a brief description" {...field} />
                </FormControl>
                <FormDescription>
                  A short description explaining the formula's purpose.
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
                <FormLabel>Formula Content</FormLabel>
                <FormControl>
                  <Textarea
                    placeholder="Enter formula content (LaTeX format supported)"
                    className="min-h-[100px] font-mono resize-y"
                    {...field}
                  />
                </FormControl>
                <FormDescription>
                  Enter your mathematical formula using LaTeX notation (e.g., E
                  = mc^2).
                </FormDescription>
                <FormMessage />

                <div>
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    onClick={() => setShowPreview(!showPreview)}
                    className="flex items-center gap-1 -ml-2 my-1"
                  >
                    <Eye className="h-4 w-4" />
                    {showPreview ? "Hide Preview" : "Show Preview"}
                  </Button>

                  {showPreview && content && (
                    <div className="p-4 mt-2">
                      <LatexRenderer content={content} />
                    </div>
                  )}
                </div>
              </FormItem>
            )}
          />

          <div className="flex justify-end gap-4">
            {onCancel && (
              <Button type="button" variant="outline" onClick={onCancel}>
                Cancel
              </Button>
            )}
            <Button
              type="submit"
              disabled={
                isSubmitting ||
                !isValid ||
                (isEditMode && Object.keys(dirtyFields).length === 0)
              }
            >
              {isSubmitting ? "Submitting..." : submitButtonText}
            </Button>
          </div>
        </form>
      </Form>
    </>
  );
}
