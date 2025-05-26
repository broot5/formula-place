import {
  deleteFormula,
  getFormula,
  updateFormula,
} from "@/services/formulaService";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import type { UUID } from "crypto";
import { useEffect, useState } from "react";
import type { FormulaResponse } from "@/types/formula";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowLeft, CalendarClock, CalendarDays, Trash } from "lucide-react";
import type { FormulaFormData } from "@/schemas/formulaSchema";
import { FormulaForm } from "@/components/FormulaForm";
import { LoadingState } from "@/components/LoadingState";
import { ErrorState } from "@/components/ErrorState";
import { Separator } from "@/components/ui/separator";
import { PageHeader } from "@/components/PageHeader";
import {
  AlertDialog,
  AlertDialogHeader,
  AlertDialogContent,
  AlertDialogTrigger,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogCancel,
  AlertDialogAction,
} from "@/components/ui/alert-dialog";

export const Route = createFileRoute("/formulas/$id")({
  component: FormulaPage,
});

function FormulaPage() {
  const { id } = Route.useParams();
  const navigate = useNavigate();

  const [formula, setFormula] = useState<FormulaResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchFormula = async () => {
      try {
        const data = await getFormula(id as UUID);
        setFormula(data);
      } catch (err) {
        setError("Failed to get formula.");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchFormula();
  }, [id]);

  const handleDelete = async () => {
    try {
      await deleteFormula(id as UUID);
      navigate({ to: "/formulas" });
    } catch (err) {
      setError("Failed to delete formula.");
      console.error(err);
    }
  };

  const handleSubmit = async (data: Partial<FormulaFormData>) => {
    try {
      await updateFormula(id as UUID, data);

      // Update the local formula data
      if (formula) {
        setFormula({
          ...formula,
          ...data,
          updated_at: new Date(),
        });
      }
    } catch (err) {
      setError("Failed to update formula.");
      console.error(err);
    }
  };

  return (
    <div className="max-w-4xl mx-auto flex flex-col gap-6">
      <PageHeader
        pageName="Edit Formula"
        action={
          <div className="flex gap-4">
            <Button
              variant="outline"
              onClick={() => navigate({ to: "/formulas" })}
            >
              <ArrowLeft className="h-4 w-4" />
              Back to Formulas
            </Button>

            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button variant="destructive" disabled={loading || !formula}>
                  <Trash className="h-4 w-4" />
                  Delete Formula
                </Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Delete Formula</AlertDialogTitle>
                  <AlertDialogDescription>
                    Are you sure you want to delete this formula? This action
                    cannot be undone.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction
                    onClick={handleDelete}
                    className="bg-destructive hover:bg-destructive/50"
                  >
                    Delete
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          </div>
        }
      />

      <Separator />

      <Card>
        {loading ? (
          <LoadingState />
        ) : !formula ? (
          <ErrorState error="Formula not found" />
        ) : error ? (
          <ErrorState error={error} />
        ) : (
          <>
            <CardHeader>
              <div className="flex justify-between mb-4">
                <CardTitle className="flex items-center text-2xl">
                  {formula.title}
                </CardTitle>

                <div className="flex flex-col gap-1 text-sm text-muted-foreground">
                  <div className="flex items-center justify-end gap-1">
                    <CalendarDays className="h-4 w-4" />
                    <span>
                      Created: {new Date(formula.created_at).toLocaleString()}
                    </span>
                  </div>
                  <div className="flex items-center justify-end gap-1">
                    <CalendarClock className="h-4 w-4" />
                    <span>
                      Updated: {new Date(formula.updated_at).toLocaleString()}
                    </span>
                  </div>
                </div>
              </div>
            </CardHeader>
            <CardContent>
              <FormulaForm
                defaultValues={{
                  title: formula.title,
                  description: formula.description,
                  content: formula.content,
                }}
                onSubmit={handleSubmit}
                error={error}
                isEditMode={true}
                submitButtonText="Save Changes"
                onCancel={() => navigate({ to: "/formulas" })}
              />
            </CardContent>
          </>
        )}
      </Card>
    </div>
  );
}
