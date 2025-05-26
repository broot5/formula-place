import { Button } from "@/components/ui/button";
import { createFileRoute, Link } from "@tanstack/react-router";
import { ArrowRight } from "lucide-react";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  return (
    <div className="flex flex-col items-center justify-center max-w-4xl mx-auto py-12 gap-4">
      <h1 className="text-3xl md:text-4xl font-bold text-center tracking-tight">
        Formula Place
      </h1>
      <p className="text-muted-foreground text-center">
        A place to create, store, and share mathematical formulas
      </p>
      <Link to="/formulas" className="flex items-center">
        <Button>
          Get Started <ArrowRight className="h-4 w-4" />
        </Button>
      </Link>
    </div>
  );
}
