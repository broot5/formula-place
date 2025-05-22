import { getFormula } from "@/services/formulaService";
import { createFileRoute } from "@tanstack/react-router";
import type { UUID } from "crypto";
import { useEffect, useState } from "react";
import type { FormulaResponse } from "@/types/formula";

export const Route = createFileRoute("/formulas/$id/")({
  component: FormulaPage,
});

function FormulaPage() {
  const { id } = Route.useParams();

  const [formula, setFormula] = useState<FormulaResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchFormula = async () => {
      try {
        const data = await getFormula(id as UUID);
        setFormula(data);
      } catch (err) {
        setError("Failed to get formula");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchFormula();
  }, [id]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error || !formula) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h1>{formula.title}</h1>

      <div>
        <p>Created: {new Date(formula.created_at).toLocaleString()}</p>
        <p>Updated: {new Date(formula.updated_at).toLocaleString()}</p>
      </div>

      <div>
        <h2>Description</h2>
        <p>{formula.description}</p>
      </div>

      <div>
        <h2>Content</h2>
        <pre>{formula.content}</pre>
      </div>
    </div>
  );
}
