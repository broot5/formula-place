import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { getAllFormulas } from "@/services/formulaService";
import type { FormulaResponse } from "@/types/formula";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Skeleton } from "@/components/ui/skeleton";
import { FileText, PlusCircle } from "lucide-react";
import { EmptyState } from "@/components/EmptyState";
import { ErrorState } from "@/components/ErrorState";
import { PageHeader } from "@/components/PageHeader";
import { FormulaCard } from "@/components/FormulaCard";

export const Route = createFileRoute("/formulas/")({
  component: FormulasPage,
});

function FormulasPage() {
  const navigate = useNavigate();

  const [formulas, setFormulas] = useState<FormulaResponse[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    getAllFormulas()
      .then((data) => {
        setFormulas(data);
        setLoading(false);
      })
      .catch(() => {
        setError("Failed to load the list of formulas.");
        setLoading(false);
      });
  }, []);

  return (
    <div className="flex flex-col gap-6">
      <PageHeader
        pageName="Formulas"
        action={
          !loading &&
          !error && (
            <Button
              type="button"
              onClick={() => navigate({ to: "/formulas/new" })}
            >
              <PlusCircle className="h-4 w-4" />
              Add Formula
            </Button>
          )
        }
      />

      <Separator />

      {/* Loading state */}
      {loading && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3, 4, 5, 6].map((i) => (
            <Card key={i}>
              <CardHeader>
                <Skeleton className="h-6 w-3/4" />
                <Skeleton className="h-4 w-1/2" />
              </CardHeader>
              <CardContent>
                <Skeleton className="h-36 w-full" />
              </CardContent>
              <CardFooter>
                <Skeleton className="h-4 w-1/3" />
              </CardFooter>
            </Card>
          ))}
        </div>
      )}

      {/* Error state */}
      {!loading && error && <ErrorState error={error} />}

      {/* Empty state */}
      {!loading && !error && formulas.length === 0 && (
        <EmptyState
          icon={FileText}
          title="No formulas yet"
          description="Start by creating your first formula."
          action={
            <Button onClick={() => navigate({ to: "/formulas/new" })}>
              <PlusCircle className="h-4 w-4" />
              Create your first formula
            </Button>
          }
        />
      )}

      {/* Data loaded successfully */}
      {!loading && !error && formulas.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {formulas.map((formula) => (
            <FormulaCard key={formula.id} formula={formula} />
          ))}
        </div>
      )}
    </div>
  );
}
