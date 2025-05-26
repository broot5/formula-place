import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";
import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

export const Route = createRootRoute({
  component: () => (
    <div className="min-h-screen flex flex-col bg-background text-foreground">
      <Header />
      <main className="w-full max-w-7xl mx-auto px-4 md:px-8 py-4 flex-1">
        <Outlet />
      </main>
      <Footer />
      <TanStackRouterDevtools />
    </div>
  ),
});
