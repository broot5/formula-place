import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { getAllFormulas } from "@/services/formulaService";
import type { FormulaResponse } from "@/types/formula";

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
        setError("Failed to load the list of formulas");
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div>
      <h1>Formula List</h1>
      {formulas.length === 0 ? (
        <div>No formulas registered.</div>
      ) : (
        <ul>
          {formulas.map((formula) => (
            <li
              key={formula.id}
              onClick={() => navigate({ to: `/formulas/${formula.id}` })}
            >
              <h2>{formula.title}</h2>
              <p>{formula.description}</p>
              <pre>{formula.content}</pre>
              <div>
                Created: {new Date(formula.created_at).toLocaleString()}
                <br />
                Updated: {new Date(formula.updated_at).toLocaleString()}
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
