import { createFileRoute } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { getAllFormulas } from "@/services/formulaService";
import type { FormulaResponse } from "@/types/formula";

export const Route = createFileRoute("/formulas/")({
  component: FormulasPage,
});

function FormulasPage() {
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

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-4">Formula List</h2>
      {formulas.length === 0 ? (
        <div>No formulas registered.</div>
      ) : (
        <ul className="space-y-4">
          {formulas.map((formula) => (
            <li key={formula.id} className="border rounded p-4">
              <h3 className="text-lg font-semibold">{formula.title}</h3>
              <p className="text-gray-700 mb-2">{formula.description}</p>
              <pre className="bg-gray-100 p-2 rounded overflow-x-auto text-sm">
                {formula.content}
              </pre>
              <div className="text-xs text-gray-500 mt-2">
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
