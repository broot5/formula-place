import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/formulas')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/formulas"!</div>
}
