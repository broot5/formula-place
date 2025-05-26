import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { createFormula } from "@/services/formulaService";
import { useState } from "react";
import type { FormulaFormData } from "@/schemas/formulaSchema";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { FormulaForm } from "@/components/FormulaForm";
import { PageHeader } from "@/components/PageHeader";
import { ArrowLeft } from "lucide-react";

export const Route = createFileRoute("/formulas/new")({
  component: NewFormulaPage,
});

function NewFormulaPage() {
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (data: Partial<FormulaFormData>) => {
    try {
      if (!data.title || !data.content) {
        setError("Title and content are required.");
        return;
      }
      await createFormula({
        title: data.title,
        content: data.content,
        description: data.description ?? "",
      });
      navigate({ to: "/formulas" });
    } catch (err) {
      setError("Failed to create formula.");
      console.error(err);
    }
  };

  return (
    <div className="max-w-4xl mx-auto flex flex-col gap-6">
      <PageHeader
        pageName="Create New Formula"
        action={
          <Button
            variant="outline"
            onClick={() => navigate({ to: "/formulas" })}
          >
            <ArrowLeft className="h-4 w-4" />
            Back to Formulas
          </Button>
        }
      />

      <Separator />

      <Card>
        <CardContent>
          <FormulaForm
            onSubmit={handleSubmit}
            error={error}
            submitButtonText="Create Formula"
            onCancel={() => navigate({ to: "/formulas" })}
          />
        </CardContent>
      </Card>
    </div>
  );
}
