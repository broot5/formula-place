import type { FormulaResponse } from "@/types/formula";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { LatexRenderer } from "@/components/LatexRenderer";
import { useNavigate } from "@tanstack/react-router";
import { Clock } from "lucide-react";
import { formatDistanceToNow } from "date-fns";

interface FormulaCardProps {
  formula: FormulaResponse;
}

export function FormulaCard({ formula }: FormulaCardProps) {
  const navigate = useNavigate();

  return (
    <Card
      onClick={() => navigate({ to: `/formulas/${formula.id}` })}
      className="cursor-pointer"
    >
      <CardHeader>
        <CardTitle className="truncate">{formula.title}</CardTitle>
        <CardDescription className="line-clamp-1">
          {formula.description}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <LatexRenderer content={formula.content} />
      </CardContent>
      <CardFooter className="text-xs text-muted-foreground gap-1">
        <Clock className="h-4 w-4" />
        <span>
          {formatDistanceToNow(new Date(formula.updated_at), {
            addSuffix: true,
          })}
        </span>
      </CardFooter>
    </Card>
  );
}
