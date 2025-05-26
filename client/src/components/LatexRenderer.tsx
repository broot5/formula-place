import { MathJax } from "better-react-mathjax";
import { cn } from "@/lib/utils";

interface LatexRendererProps {
  content: string;
  className?: string;
}

export function LatexRenderer({ content, className }: LatexRendererProps) {
  return (
    <div
      className={cn("border rounded-md bg-muted/50 overflow-x-auto", className)}
    >
      <MathJax>{`$$${content}$$`}</MathJax>
    </div>
  );
}
