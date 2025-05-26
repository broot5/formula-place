import { AlertCircle } from "lucide-react";

interface ErrorStateProps {
  error: string;
}

export function ErrorState({ error }: ErrorStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-8 text-center">
      <div className="relative">
        <div className="rounded-full bg-destructive/20 p-6">
          <AlertCircle className="h-12 w-12 text-destructive" />
        </div>
      </div>
      <h3 className="text-lg font-medium mt-4">Error</h3>
      <p className="text-muted-foreground mt-2 max-w-md">{error}</p>
    </div>
  );
}
